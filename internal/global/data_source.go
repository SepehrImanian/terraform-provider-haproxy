package global

import (
	"fmt"
	"net/http"
	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func DataSourceHaproxyGlobal() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHaproxyGlobalRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the Global. It must be unique and cannot be changed.",
			},
		},
	}
}

func dataSourceHaproxyGlobalRead(d *schema.ResourceData, m interface{}) error {
	configMap := m.(map[string]interface{})
	GlobalConfig := configMap["global"].(*ConfigGlobal)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return GlobalConfig.GetAGlobalConfiguration(transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError("global", "error reading Global configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId("global")
	return nil
}
