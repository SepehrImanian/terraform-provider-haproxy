package user

import (
	"fmt"
	"net/http"

	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func DataSourceHaproxyUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHaproxyUserRead,
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the User. It must be unique",
			},
			"userlist": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the parent object",
			},
		},
	}
}

func dataSourceHaproxyUserRead(d *schema.ResourceData, m interface{}) error {
	userName := d.Get("username").(string)
	userlist := d.Get("userlist").(string)

	configMap := m.(map[string]interface{})
	UserConfig := configMap["user"].(*ConfigUser)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, _, err := tranConfig.TransactionWithData(func(transactionID string) (*http.Response, []byte, error) {
		return UserConfig.GetAUsersConfiguration(userName, transactionID, userlist)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(userName, "error reading User configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(userName)
	return nil
}
