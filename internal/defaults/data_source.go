package defaults

import (
	"fmt"
	"net/http"
	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func DataSourceHaproxyDefaults() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHaproxyDefaultsRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: "The name of the defaults. It must be unique and cannot be changed.",
			},
		},
	}
}

func dataSourceHaproxyDefaultsRead(d *schema.ResourceData, m interface{}) error {
	defaultsName := d.Get("name").(string)

	configMap := m.(map[string]interface{})
	DefaultsConfig := configMap["defaults"].(*ConfigDefaults)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return DefaultsConfig.GetADefaultsConfiguration(defaultsName, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(defaultsName, "error reading Defaults configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(defaultsName)
	return nil
}
