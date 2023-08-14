package backend

import (
	"fmt"
	"net/http"
	"terraform-provider-haproxy/internal/transaction"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func DataSourceHaproxyBackend() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHaproxyABackendRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceHaproxyABackendRead(d *schema.ResourceData, m interface{}) error {
	backendName := d.Get("name").(string)

	configMap := m.(map[string]interface{})
	backendConfig := configMap["backend"].(*ConfigBackend)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return backendConfig.GetABackendConfiguration(backendName, transactionID)
	})

	if err != nil {
		fmt.Println("Error updating backend configuration:", err)
		return err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return fmt.Errorf("error creating backend configuration: %s", resp.Status)
	}

	d.SetId(backendName)
	return nil
}
