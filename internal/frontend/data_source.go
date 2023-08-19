package frontend

import (
	"fmt"
	"net/http"
	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func DataSourceHaproxyFrontend() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHaproxyFrontendRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the frontend. It must be unique and cannot be changed.",
			},
		},
	}
}

func dataSourceHaproxyFrontendRead(d *schema.ResourceData, m interface{}) error {
	frontendName := d.Get("name").(string)

	configMap := m.(map[string]interface{})
	frontendConfig := configMap["frontend"].(*ConfigFrontend)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return frontendConfig.GetAFrontendConfiguration(frontendName, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(frontendName, "error reading Frontend configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(frontendName)
	return nil
}
