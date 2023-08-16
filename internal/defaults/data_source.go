package defaults

import (
	"fmt"
	"net/http"
	"terraform-provider-haproxy/internal/transaction"

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

	if err != nil {
		fmt.Println("Error updating Defaults configuration:", err)
		return err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return fmt.Errorf("error reading Defaults configuration: %s", resp.Status)
	}

	d.SetId(defaultsName)
	return nil
}
