package group

import (
	"fmt"
	"net/http"

	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ResourceHaproxyGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceHaproxyGroupCreate,
		Read:   resourceHaproxyGroupRead,
		Update: resourceHaproxyGroupUpdate,
		Delete: resourceHaproxyGroupDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The Groupname of the Group",
			},
			"userlist": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The userlist of the User",
			},
			"users": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The users of the Group",
			},
		},
	}
}

func resourceHaproxyGroupRead(d *schema.ResourceData, m interface{}) error {
	groupName := d.Get("name").(string)
	userlist := d.Get("userlist").(string)

	configMap := m.(map[string]interface{})
	GroupConfig := configMap["group"].(*ConfigGroup)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, _, err := tranConfig.TransactionWithData(func(transactionID string) (*http.Response, []byte, error) {
		return GroupConfig.GetAGroupsConfiguration(groupName, transactionID, userlist)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(groupName, "error reading Group configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(groupName)
	return nil
}

func resourceHaproxyGroupCreate(d *schema.ResourceData, m interface{}) error {
	groupName := d.Get("name").(string)
	userlist := d.Get("userlist").(string)
	users := d.Get("users").(string)

	payload := GroupPayload{
		Name:  groupName,
		Users: users,
	}

	payloadJSON, err := utils.MarshalNonZeroFields(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	GroupConfig := configMap["group"].(*ConfigGroup)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, _, err := tranConfig.TransactionWithData(func(transactionID string) (*http.Response, []byte, error) {
		return GroupConfig.AddGroupsConfiguration(payloadJSON, transactionID, userlist)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(groupName, "error creating Group configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(groupName)
	return nil
}

func resourceHaproxyGroupUpdate(d *schema.ResourceData, m interface{}) error {
	groupName := d.Get("name").(string)
	userlist := d.Get("userlist").(string)
	users := d.Get("users").(string)

	payload := GroupPayload{
		Name:  groupName,
		Users: users,
	}

	payloadJSON, err := utils.MarshalNonZeroFields(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	GroupConfig := configMap["group"].(*ConfigGroup)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, _, err := tranConfig.TransactionWithData(func(transactionID string) (*http.Response, []byte, error) {
		return GroupConfig.AddGroupsConfiguration(payloadJSON, transactionID, userlist)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(groupName, "error creating Group configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(groupName)
	return nil
}

func resourceHaproxyGroupDelete(d *schema.ResourceData, m interface{}) error {
	groupName := d.Get("name").(string)
	userlist := d.Get("userlist").(string)

	configMap := m.(map[string]interface{})
	GroupConfig := configMap["group"].(*ConfigGroup)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, _, err := tranConfig.TransactionWithData(func(transactionID string) (*http.Response, []byte, error) {
		return GroupConfig.DeleteGroupsConfiguration(groupName, transactionID, userlist)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(groupName, "error deleting Group configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId("")
	return nil
}
