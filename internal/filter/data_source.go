package filter

import (
	"fmt"
	"net/http"

	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func DataSourceHaproxyFilter() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHaproxyFilterRead,
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
				Description: "The index of the Filter in the parent object starting at 0",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the Filter.",
			},
		},
	}
}

func dataSourceHaproxyFilterRead(d *schema.ResourceData, m interface{}) error {
	FilterName := d.Get("name").(string)
	indexName := d.Get("index").(int)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	configMap := m.(map[string]interface{})
	FilterConfig := configMap["filter"].(*ConfigFilter)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return FilterConfig.GetAFilterConfiguration(indexName, transactionID, parentName, parentType)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(FilterName, "error reading Filter configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(FilterName)
	return nil
}
