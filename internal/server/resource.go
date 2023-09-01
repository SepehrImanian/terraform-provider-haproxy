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
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the server. It must be unique and cannot be changed.",
			},
			"port": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The port of the server",
			},
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The address of the server",
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
			"send_proxy": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "To send a Proxy Protocol header to the backend server,",
			},
			"check": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "To enable health check for the server.",
			},
			"inter": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The inter value is the time interval in milliseconds between two consecutive health checks.",
			},
			"rise": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The rise value states that a server will be considered as operational after consecutive successful health checks.",
			},
			"fall": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The fall value states that a server will be considered as failed after consecutive unsuccessful health checks.",
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

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(serverName, "error reading Server configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(serverName)
	return nil
}

func resourceHaproxyServerCreate(d *schema.ResourceData, m interface{}) error {
	serverName := d.Get("name").(string)
	sendProxy := d.Get("send_proxy").(bool)
	check := d.Get("check").(bool)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	payload := ServerPayload{
		Name:    serverName,
		Address: d.Get("address").(string),
		Port:    d.Get("port").(int),
		Inter:   d.Get("inter").(int),
		Rise:    d.Get("rise").(int),
		Fall:    d.Get("fall").(int),
	}

	// Check sendProxy field
	if sendProxy {
		payload.SendProxy = utils.BoolToStr(sendProxy)
	}

	// Check check field
	if check {
		payload.Check = utils.BoolToStr(check)
	}

	payloadJSON, err := utils.MarshalNonZeroFields(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	serverConfig := configMap["server"].(*ConfigServer)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return serverConfig.AddServerConfiguration(payloadJSON, transactionID, parentName, parentType)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(serverName, "error creating Server configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(serverName)
	return nil
}

func resourceHaproxyServerUpdate(d *schema.ResourceData, m interface{}) error {
	serverName := d.Get("name").(string)
	sendProxy := d.Get("send_proxy").(bool)
	check := d.Get("check").(bool)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	payload := ServerPayload{
		Name:    serverName,
		Address: d.Get("address").(string),
		Port:    d.Get("port").(int),
		Inter:   d.Get("inter").(int),
		Rise:    d.Get("rise").(int),
		Fall:    d.Get("fall").(int),
	}

	// Check sendProxy field
	if sendProxy {
		payload.SendProxy = utils.BoolToStr(sendProxy)
	}

	// Check check field
	if check {
		payload.Check = utils.BoolToStr(check)
	}

	payloadJSON, err := utils.MarshalNonZeroFields(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	serverConfig := configMap["server"].(*ConfigServer)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return serverConfig.UpdateServerConfiguration(serverName, payloadJSON, transactionID, parentName, parentType)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(serverName, "error updating Server configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
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

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(serverName, "error deleting Server configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId("")
	return nil
}
