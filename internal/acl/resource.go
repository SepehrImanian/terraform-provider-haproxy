package acl

import (
	"fmt"
	"net/http"

	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ResourceHaproxyAcl() *schema.Resource {
	return &schema.Resource{
		Create: resourceHaproxyAclCreate,
		Read:   resourceHaproxyAclRead,
		Update: resourceHaproxyAclUpdate,
		Delete: resourceHaproxyAclDelete,

		Schema: map[string]*schema.Schema{
			"parent_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the parent object",
			},
			"parent_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of the parent object",
			},
			"index": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The index of the acl in the parent object starting at 0",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the acl. It must be unique and cannot be changed.",
			},
			"criterion": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The criterion of the acl",
			},
			"value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The value of the acl.",
			},
		},
	}
}

func resourceHaproxyAclRead(d *schema.ResourceData, m interface{}) error {
	aclName := d.Get("name").(string)
	indexName := d.Get("index").(int)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	configMap := m.(map[string]interface{})
	AclConfig := configMap["acl"].(*ConfigAcl)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return AclConfig.GetAAclConfiguration(indexName, transactionID, parentName, parentType)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(aclName, "error reading Acl configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(aclName)
	return nil
}

func resourceHaproxyAclCreate(d *schema.ResourceData, m interface{}) error {
	aclName := d.Get("name").(string)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	payload := AclPayload{
		AclName:   aclName,
		Criterion: d.Get("criterion").(string),
		Index:     d.Get("index").(int),
		Value:     d.Get("value").(string),
	}

	payloadJSON, err := utils.MarshalNonZeroFields(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	AclConfig := configMap["acl"].(*ConfigAcl)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return AclConfig.AddAclConfiguration(payloadJSON, transactionID, parentName, parentType)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(aclName, "error creating Acl configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(aclName)
	return nil
}

func resourceHaproxyAclUpdate(d *schema.ResourceData, m interface{}) error {
	aclName := d.Get("name").(string)
	indexName := d.Get("index").(int)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	payload := AclPayload{
		AclName:   aclName,
		Criterion: d.Get("criterion").(string),
		Index:     d.Get("index").(int),
		Value:     d.Get("value").(string),
	}

	payloadJSON, err := utils.MarshalNonZeroFields(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	AclConfig := configMap["acl"].(*ConfigAcl)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return AclConfig.UpdateAclConfiguration(indexName, payloadJSON, transactionID, parentName, parentType)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(aclName, "error updating Acl configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(aclName)
	return nil
}

func resourceHaproxyAclDelete(d *schema.ResourceData, m interface{}) error {
	aclName := d.Get("name").(string)
	indexName := d.Get("index").(int)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	configMap := m.(map[string]interface{})
	AclConfig := configMap["acl"].(*ConfigAcl)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return AclConfig.DeleteAclConfiguration(indexName, transactionID, parentName, parentType)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(aclName, "error deleting Acl configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId("")
	return nil
}
