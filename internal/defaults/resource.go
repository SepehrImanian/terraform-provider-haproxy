package defaults

import (
	"encoding/json"
	"fmt"
	"net/http"
	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ResourceHaproxyDefaults() *schema.Resource {
	return &schema.Resource{
		Create: resourceHaproxyDefaultsCreate,
		Read:   resourceHaproxyDefaultsRead,
		Update: resourceHaproxyDefaultsUpdate,
		Delete: resourceHaproxyDefaultsDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"backlog": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"httplog": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"httpslog": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"tcplog": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"retries": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"check_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"client_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"connect_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"http_keep_alive_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"http_request_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"queue_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"server_fin_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"maxconn": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceHaproxyDefaultsRead(d *schema.ResourceData, m interface{}) error {
	defaultsName := d.Get("name").(string)

	configMap := m.(map[string]interface{})
	DefaultsConfig := configMap["defaults"].(*ConfigDefaults)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return DefaultsConfig.GetADefaultsConfiguration(defaultsName, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(defaultsName, "error reading Defaults configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(defaultsName)
	return nil
}

func resourceHaproxyDefaultsCreate(d *schema.ResourceData, m interface{}) error {
	defaultsName := d.Get("name").(string)
	HTTPSLog := d.Get("httpslog").(bool)

	// Convert bool to string
	HTTPSLogStr := utils.BoolToStr(HTTPSLog)

	payload := DefaultsPayload{
		Name:                 defaultsName,
		Mode:                 d.Get("mode").(string),
		Backlog:              d.Get("backlog").(int),
		HTTPLog:              d.Get("httplog").(bool),
		HTTPSLog:             HTTPSLogStr,
		TCPLog:               d.Get("tcplog").(bool),
		Retries:              d.Get("retries").(int),
		CheckTimeout:         d.Get("check_timeout").(int),
		ClientTimeout:        d.Get("client_timeout").(int),
		ConnectTimeout:       d.Get("connect_timeout").(int),
		HTTPKeepAliveTimeout: d.Get("http_keep_alive_timeout").(int),
		HTTPRequestTimeout:   d.Get("http_request_timeout").(int),
		QueueTimeout:         d.Get("queue_timeout").(int),
		ServerTimeout:        d.Get("server_timeout").(int),
		ServerFinTimeout:     d.Get("server_fin_timeout").(int),
		MaxConn:              d.Get("maxconn").(int),
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	DefaultsConfig := configMap["defaults"].(*ConfigDefaults)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return DefaultsConfig.AddDefaultsConfiguration(payloadJSON, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(defaultsName, "error creating Defaults configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(defaultsName)
	return nil
}

func resourceHaproxyDefaultsUpdate(d *schema.ResourceData, m interface{}) error {
	defaultsName := d.Get("name").(string)

	HTTPSLog := d.Get("httpslog").(bool)

	// Convert bool to string
	HTTPSLogStr := utils.BoolToStr(HTTPSLog)

	payload := DefaultsPayload{
		Name:                 defaultsName,
		Mode:                 d.Get("mode").(string),
		Backlog:              d.Get("backlog").(int),
		HTTPLog:              d.Get("httplog").(bool),
		HTTPSLog:             HTTPSLogStr,
		TCPLog:               d.Get("tcplog").(bool),
		Retries:              d.Get("retries").(int),
		CheckTimeout:         d.Get("check_timeout").(int),
		ClientTimeout:        d.Get("client_timeout").(int),
		ConnectTimeout:       d.Get("connect_timeout").(int),
		HTTPKeepAliveTimeout: d.Get("http_keep_alive_timeout").(int),
		HTTPRequestTimeout:   d.Get("http_request_timeout").(int),
		QueueTimeout:         d.Get("queue_timeout").(int),
		ServerTimeout:        d.Get("server_timeout").(int),
		ServerFinTimeout:     d.Get("server_fin_timeout").(int),
		MaxConn:              d.Get("maxconn").(int),
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	DefaultsConfig := configMap["defaults"].(*ConfigDefaults)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return DefaultsConfig.UpdateDefaultsConfiguration(defaultsName, payloadJSON, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(defaultsName, "error updating Defaults configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(defaultsName)
	return nil
}

func resourceHaproxyDefaultsDelete(d *schema.ResourceData, m interface{}) error {
	defaultsName := d.Get("name").(string)

	configMap := m.(map[string]interface{})
	DefaultsConfig := configMap["defaults"].(*ConfigDefaults)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return DefaultsConfig.DeleteDefaultsConfiguration(defaultsName, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(defaultsName, "error deleting Defaults configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId("")
	return nil
}
