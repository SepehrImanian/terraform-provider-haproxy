package nameserver

import (
	"fmt"
	"net/http"

	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func DataSourceHaproxyNameserver() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHaproxyNameserverRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the Nameserver. It must be unique",
			},
			"resolver": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the parent object",
			},
		},
	}
}

func dataSourceHaproxyNameserverRead(d *schema.ResourceData, m interface{}) error {
	nameserverName := d.Get("name").(string)
	resolver := d.Get("resolver").(string)

	configMap := m.(map[string]interface{})
	NameserverConfig := configMap["nameserver"].(*ConfigNameserver)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, _, err := tranConfig.TransactionWithData(func(transactionID string) (*http.Response, []byte, error) {
		return NameserverConfig.GetANameserversConfiguration(nameserverName, transactionID, resolver)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(nameserverName, "error reading Nameserver configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(nameserverName)
	return nil
}
