package bind

import (
	"fmt"
	"net/http"

	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func DataSourceHaproxyBind() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHaproxyBindRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the bind. It must be unique and cannot be changed.",
			},
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
		},
	}
}

func dataSourceHaproxyBindRead(d *schema.ResourceData, m interface{}) error {
	bindName := d.Get("name").(string)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	configMap := m.(map[string]interface{})
	bindConfig := configMap["bind"].(*ConfigBind)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return bindConfig.GetABindConfiguration(bindName, transactionID, parentName, parentType)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(bindName, "error reading Bind configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(bindName)
	return nil
}
