package ServerTemplate

import (
	"fmt"
	"net/http"

	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func DataSourceHaproxyServerTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHaproxyServerTemplateRead,
		Schema: map[string]*schema.Schema{
			"prefix": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the ServerTemplate.",
			},
			"backend": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The backend of the ServerTemplate.",
			},
		},
	}
}

func dataSourceHaproxyServerTemplateRead(d *schema.ResourceData, m interface{}) error {
	prefix := d.Get("prefix").(string)
	backend := d.Get("backend").(string)

	configMap := m.(map[string]interface{})
	ServerTemplateConfig := configMap["ServerTemplate"].(*ConfigServerTemplate)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, _, err := tranConfig.TransactionWithData(func(transactionID string) (*http.Response, []byte, error) {
		return ServerTemplateConfig.GetAServerTemplatesConfiguration(prefix, transactionID, backend)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(prefix, "error reading ServerTemplate configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(prefix)
	return nil
}
