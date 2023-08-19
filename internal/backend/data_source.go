package backend

import (
	"fmt"
	"net/http"
	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func DataSourceHaproxyBackend() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHaproxyABackendRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Description: "The name of the backend. It must be unique and cannot be changed.",
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

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(backendName, "error reading Backend configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(backendName)
	return nil
}
