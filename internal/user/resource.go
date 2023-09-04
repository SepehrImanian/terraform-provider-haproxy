package user

import (
	"fmt"
	"net/http"

	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ResourceHaproxyUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceHaproxyUserCreate,
		Read:   resourceHaproxyUserRead,
		Update: resourceHaproxyUserUpdate,
		Delete: resourceHaproxyUserDelete,

		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The username of the User",
			},
			"userlist": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The userlist of the User",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The password of the User",
			},
			"secure_password": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The secure password of the User",
			},
			"groups": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The groups of the User",
			},
		},
	}
}

func resourceHaproxyUserRead(d *schema.ResourceData, m interface{}) error {
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

func resourceHaproxyUserCreate(d *schema.ResourceData, m interface{}) error {
	userName := d.Get("username").(string)
	userlist := d.Get("userlist").(string)

	payload := UserPayload{
		Username:       userName,
		Password:       d.Get("password").(string),
		SecurePassword: d.Get("secure_password").(bool),
		Groups:         d.Get("groups").(string),
	}

	payloadJSON, err := utils.MarshalNonZeroFields(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	UserConfig := configMap["user"].(*ConfigUser)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, _, err := tranConfig.TransactionWithData(func(transactionID string) (*http.Response, []byte, error) {
		return UserConfig.AddUsersConfiguration(payloadJSON, transactionID, userlist)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(userName, "error creating User configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(userName)
	return nil
}

func resourceHaproxyUserUpdate(d *schema.ResourceData, m interface{}) error {
	userName := d.Get("username").(string)
	userlist := d.Get("userlist").(string)

	payload := UserPayload{
		Username:       userName,
		Password:       d.Get("password").(string),
		SecurePassword: d.Get("secure_password").(bool),
		Groups:         d.Get("groups").(string),
	}

	payloadJSON, err := utils.MarshalNonZeroFields(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	UserConfig := configMap["user"].(*ConfigUser)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, _, err := tranConfig.TransactionWithData(func(transactionID string) (*http.Response, []byte, error) {
		return UserConfig.UpdateUsersConfiguration(userName, payloadJSON, transactionID, userlist)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(userName, "error updating User configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(userName)
	return nil
}

func resourceHaproxyUserDelete(d *schema.ResourceData, m interface{}) error {
	userName := d.Get("username").(string)
	userlist := d.Get("userlist").(string)

	configMap := m.(map[string]interface{})
	UserConfig := configMap["user"].(*ConfigUser)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, _, err := tranConfig.TransactionWithData(func(transactionID string) (*http.Response, []byte, error) {
		return UserConfig.DeleteUsersConfiguration(userName, transactionID, userlist)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(userName, "error deleting User configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId("")
	return nil
}
