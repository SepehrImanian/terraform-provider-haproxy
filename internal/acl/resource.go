package acl

import (
	"encoding/json"
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
				Type:     schema.TypeString,
				Required: true,
			},
			"parent_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"index": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"criterion": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
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

	payloadJSON, err := json.Marshal(payload)
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

	payloadJSON, err := json.Marshal(payload)
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
