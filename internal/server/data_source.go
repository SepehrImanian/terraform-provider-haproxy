package server

import (
	"fmt"
	"net/http"
	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func DataSourceHaproxyServer() *schema.Resource {
	return &schema.Resource{
		Read: DataSourceHaproxyServerRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the server. It must be unique and cannot be changed.",
			},
			"parent_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the parent object",
			},
			"parent_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of the parent object",
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

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(serverName, "error reading Server configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(serverName)
	return nil
}
