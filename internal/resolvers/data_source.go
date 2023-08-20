package resolvers

import (
	"fmt"
	"net/http"
	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func DataSourceHaproxyResolvers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHaproxyResolversRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the Resolvers. It must be unique and cannot be changed.",
			},
		},
	}
}

func dataSourceHaproxyResolversRead(d *schema.ResourceData, m interface{}) error {
	resolversName := d.Get("name").(string)

	configMap := m.(map[string]interface{})
	ResolversConfig := configMap["resolvers"].(*ConfigResolvers)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return ResolversConfig.GetAResolversConfiguration(resolversName, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(resolversName, "error reading Resolvers configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(resolversName)
	return nil
}
