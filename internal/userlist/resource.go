package userlist

import (
	"fmt"
	"net/http"
	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ResourceHaproxyUserlist() *schema.Resource {
	return &schema.Resource{
		Create: resourceHaproxyUserlistCreate,
		Read:   resourceHaproxyUserlistRead,
		Update: resourceHaproxyUserlistCreate,
		Delete: resourceHaproxyUserlistDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the Userlist. It must be unique",
			},
		},
	}
}

func resourceHaproxyUserlistRead(d *schema.ResourceData, m interface{}) error {
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

func resourceHaproxyUserlistCreate(d *schema.ResourceData, m interface{}) error {
	userlistName := d.Get("name").(string)

	payload := UserlistPayload{
		Name: userlistName,
	}

	payloadJSON, err := utils.MarshalNonZeroFields(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	UserlistConfig := configMap["userlist"].(*ConfigUserlist)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return UserlistConfig.AddUserlistConfiguration(payloadJSON, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(userlistName, "error creating Userlist configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(userlistName)
	return nil
}

func resourceHaproxyUserlistDelete(d *schema.ResourceData, m interface{}) error {
	userlistName := d.Get("name").(string)

	configMap := m.(map[string]interface{})
	UserlistConfig := configMap["userlist"].(*ConfigUserlist)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return UserlistConfig.DeleteUserlistConfiguration(userlistName, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(userlistName, "error deleting Userlist configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId("")
	return nil
}
