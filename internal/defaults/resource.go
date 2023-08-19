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
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the defaults. It must be unique and cannot be changed.",
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "http, tcp",
			},
			"backlog": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The backlog of the defaults, it can be true or false",
			},
			"httplog": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The httplog of the defaults, it can be true or false",
			},
			"httpslog": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The httpslog of the defaults, it can be true or false",
			},
			"tcplog": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The tcplog of the defaults, it can be true or false",
			},
			"retries": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The retries of the defaults, it can be integer or null",
			},
			"check_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The check timeout of the defaults",
			},
			"client_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The client timeout of the defaults",
			},
			"connect_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The connect timeout of the defaults",
			},
			"http_keep_alive_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The http keep alive timeout of the defaults",
			},
			"http_request_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The http request timeout of the defaults",
			},
			"queue_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The queue timeout of the defaults",
			},
			"server_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The server timeout of the defaults",
			},
			"server_fin_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The server fin timeout of the defaults",
			},
			"maxconn": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The max connections of the defaults",
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
