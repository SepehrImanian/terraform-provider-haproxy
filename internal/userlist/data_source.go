package userlist

import (
	"fmt"
	"net/http"
	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func DataSourceHaproxyUserlist() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHaproxyUserlistRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the Userlist. It must be unique",
			},
		},
	}
}

func dataSourceHaproxyUserlistRead(d *schema.ResourceData, m interface{}) error {
	userlistName := d.Get("name").(string)

	configMap := m.(map[string]interface{})
	UserlistConfig := configMap["userlist"].(*ConfigUserlist)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return UserlistConfig.GetAUserlistConfiguration(userlistName, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(userlistName, "error reading Userlist configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(userlistName)
	return nil
}
