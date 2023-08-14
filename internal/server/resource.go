package server

import (
	"fmt"
	"net/http"

	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ResourceHaproxyServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceHaproxyServerCreate,
		Read:   resourceHaproxyServerRead,
		Update: resourceHaproxyServerUpdate,
		Delete: resourceHaproxyServerDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"parent_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"parent_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"send_proxy": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"check": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"inter": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rise": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fall": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceHaproxyServerRead(d *schema.ResourceData, m interface{}) error {
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

func resourceHaproxyServerCreate(d *schema.ResourceData, m interface{}) error {
	serverName := d.Get("name").(string)
	port := d.Get("port").(int)
	address := d.Get("address").(string)
	sendProxy := d.Get("send_proxy").(bool)
	check := d.Get("check").(bool)
	inter := d.Get("inter").(int)
	rise := d.Get("rise").(int)
	fall := d.Get("fall").(int)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	// change bool to enabled/disabled
	sendProxyStr := utils.BoolToStr(sendProxy)
	checkStr := utils.BoolToStr(check)

	payload := []byte(fmt.Sprintf(`
	{
		"name": "%s",
		"address": "%s",
		"port": %d,
		"send-proxy": "%s",
		"check": "%s",
		"inter": %d,
		"rise": %d,
		"fall": %d
	}
	`, serverName, address, port, sendProxyStr, checkStr, inter, rise, fall))

	configMap := m.(map[string]interface{})
	serverConfig := configMap["server"].(*ConfigServer)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return serverConfig.AddServerConfiguration(payload, transactionID, parentName, parentType)
	})

	if err != nil {
		fmt.Println("Error creating Server configuration:", err)
		return err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return fmt.Errorf("error creating Server configuration: %s", resp.Status)
	}

	d.SetId(serverName)
	return nil
}

func resourceHaproxyServerUpdate(d *schema.ResourceData, m interface{}) error {
	serverName := d.Get("name").(string)
	port := d.Get("port").(int)
	address := d.Get("address").(string)
	sendProxy := d.Get("send_proxy").(bool)
	check := d.Get("check").(bool)
	inter := d.Get("inter").(int)
	rise := d.Get("rise").(int)
	fall := d.Get("fall").(int)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	// change bool to enabled/disabled
	sendProxyStr := utils.BoolToStr(sendProxy)
	checkStr := utils.BoolToStr(check)

	payload := []byte(fmt.Sprintf(`
	{
		"name": "%s",
		"address": "%s",
		"port": %d,
		"send-proxy": "%s",
		"check": "%s",
		"inter": %d,
		"rise": %d,
		"fall": %d
	}
	`, serverName, address, port, sendProxyStr, checkStr, inter, rise, fall))

	configMap := m.(map[string]interface{})
	serverConfig := configMap["server"].(*ConfigServer)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return serverConfig.UpdateServerConfiguration(serverName, payload, transactionID, parentName, parentType)
	})

	if err != nil {
		fmt.Println("Error creating Server configuration:", err)
		return err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return fmt.Errorf("error creating Server configuration: %s", resp.Status)
	}

	d.SetId(serverName)
	return nil
}

func resourceHaproxyServerDelete(d *schema.ResourceData, m interface{}) error {
	serverName := d.Get("name").(string)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	configMap := m.(map[string]interface{})
	serverConfig := configMap["server"].(*ConfigServer)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return serverConfig.DeleteServerConfiguration(serverName, transactionID, parentName, parentType)
	})

	if err != nil {
		fmt.Println("Error updating Server configuration:", err)
		return err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return fmt.Errorf("error creating Server configuration: %s", resp.Status)
	}

	d.SetId("")
	return nil
}
