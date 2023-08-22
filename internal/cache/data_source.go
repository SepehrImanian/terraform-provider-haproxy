package cache

import (
	"fmt"
	"net/http"
	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func DataSourceHaproxyCache() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHaproxyCacheRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: "The name of the Cache. It must be unique and cannot be changed.",
			},
		},
	}
}

func dataSourceHaproxyCacheRead(d *schema.ResourceData, m interface{}) error {
	CacheName := d.Get("name").(string)

	configMap := m.(map[string]interface{})
	CacheConfig := configMap["cache"].(*ConfigCache)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return CacheConfig.GetACacheConfiguration(CacheName, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(CacheName, "error reading Cache configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(CacheName)
	return nil
}
