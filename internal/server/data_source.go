package server

import (
	"fmt"
	"net/http"
	"terraform-provider-haproxy/internal/transaction"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func DataSourceHaproxyServer() *schema.Resource {
	return &schema.Resource{
		Read: DataSourceHaproxyServerRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"parent_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"parent_type": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func DataSourceHaproxyServerRead(d *schema.ResourceData, m interface{}) error {
	serverName := d.Get("name").(string)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	configMap := m.(map[string]interface{})
	serverConfig := configMap["server"].(*ConfigServer)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return serverConfig.GetAServerConfiguration(serverName, transactionID, parentName, parentType)
	})

	if err != nil {
		fmt.Println("Error updating Server configuration:", err)
		return err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return fmt.Errorf("error creating Server configuration: %s", resp.Status)
	}

	d.SetId(serverName)
	return nil
}
