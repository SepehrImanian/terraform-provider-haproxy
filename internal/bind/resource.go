package bind

import (
	"fmt"
	"net/http"

	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ResourceHaproxyBind() *schema.Resource {
	return &schema.Resource{
		Create: resourceHaproxyBindCreate,
		Read:   resourceHaproxyBindRead,
		Update: resourceHaproxyBindUpdate,
		Delete: resourceHaproxyBindDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the bind. It must be unique and cannot be changed.",
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
			"port": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The port of the bind",
			},
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The address of the bind",
			},
			"mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "http, tcp",
			},
			"maxconn": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The max connections of the bind",
			},
			"user": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The user of the bind",
			},
			"group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The group of the bind",
			},
			"ssl": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable SSL termination on the bind (equivalent to the `ssl` keyword on the HAProxy `bind` line).",
			},
			"ssl_certificate": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Path to the SSL certificate (or directory of certificates) to load, equivalent to the `crt` keyword on the HAProxy `bind` line.",
			},
			"ssl_cafile": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Path to the CA certificate file used to verify client certificates, equivalent to the `ca-file` keyword on the HAProxy `bind` line.",
			},
			"verify": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Client certificate verification mode: none, optional, or required.",
			},
		},
	}
}

func resourceHaproxyBindRead(d *schema.ResourceData, m interface{}) error {
	bindName := d.Get("name").(string)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	configMap := m.(map[string]interface{})
	BindConfig := configMap["bind"].(*ConfigBind)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return BindConfig.GetABindConfiguration(bindName, transactionID, parentName, parentType)
	})

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(bindName, "error reading Bind configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(bindName)
	return nil
}

func resourceHaproxyBindCreate(d *schema.ResourceData, m interface{}) error {
	bindName := d.Get("name").(string)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	payload := BindPayload{
		Name:           bindName,
		Address:        d.Get("address").(string),
		Port:           d.Get("port").(int),
		Maxconn:        d.Get("maxconn").(int),
		User:           d.Get("user").(string),
		Group:          d.Get("group").(string),
		Mode:           d.Get("mode").(string),
		Ssl:            d.Get("ssl").(bool),
		SslCertificate: d.Get("ssl_certificate").(string),
		SslCafile:      d.Get("ssl_cafile").(string),
		Verify:         d.Get("verify").(string),
	}

	payloadJSON, err := utils.MarshalNonZeroFields(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	BindConfig := configMap["bind"].(*ConfigBind)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return BindConfig.AddBindConfiguration(payloadJSON, transactionID, parentName, parentType)
	})

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(bindName, "error creating Bind configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(bindName)
	return nil
}

func resourceHaproxyBindUpdate(d *schema.ResourceData, m interface{}) error {
	bindName := d.Get("name").(string)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	payload := BindPayload{
		Name:           bindName,
		Address:        d.Get("address").(string),
		Port:           d.Get("port").(int),
		Maxconn:        d.Get("maxconn").(int),
		User:           d.Get("user").(string),
		Group:          d.Get("group").(string),
		Mode:           d.Get("mode").(string),
		Ssl:            d.Get("ssl").(bool),
		SslCertificate: d.Get("ssl_certificate").(string),
		SslCafile:      d.Get("ssl_cafile").(string),
		Verify:         d.Get("verify").(string),
	}

	payloadJSON, err := utils.MarshalNonZeroFields(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	BindConfig := configMap["bind"].(*ConfigBind)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return BindConfig.UpdateBindConfiguration(bindName, payloadJSON, transactionID, parentName, parentType)
	})

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(bindName, "error updating Bind configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(bindName)
	return nil
}

func resourceHaproxyBindDelete(d *schema.ResourceData, m interface{}) error {
	bindName := d.Get("name").(string)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	configMap := m.(map[string]interface{})
	BindConfig := configMap["bind"].(*ConfigBind)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return BindConfig.DeleteBindConfiguration(bindName, transactionID, parentName, parentType)
	})

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(bindName, "error deleting Bind configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId("")
	return nil
}
