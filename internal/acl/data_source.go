package acl

import (
	"fmt"
	"net/http"

	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func DataSourceHaproxyAcl() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHaproxyAclRead,
		Schema: map[string]*schema.Schema{
			"parent_name": {
				Type:        schema.TypeString,
				Description: "The name of the parent object",
				Required:    true,
			},
			"parent_type": {
				Type:        schema.TypeString,
				Description: "The type of the parent object",
				Required:    true,
			},
			"index": {
				Type:        schema.TypeInt,
				Description: "The index of the acl in the parent object starting at 0",
				Optional:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the acl. It must be unique and cannot be changed.",
			},
		},
	}
}

func dataSourceHaproxyAclRead(d *schema.ResourceData, m interface{}) error {
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
