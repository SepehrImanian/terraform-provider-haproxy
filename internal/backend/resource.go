package backend

import (
	"encoding/json"
	"fmt"
	"net/http"

	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ResourceHaproxyBackend() *schema.Resource {
	return &schema.Resource{
		Create: resourceHaproxyBackendCreate,
		Read:   resourceHaproxyBackendRead,
		Update: resourceHaproxyBackendUpdate,
		Delete: resourceHaproxyBackendDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the backend. It must be unique and cannot be changed.",
			},
			"mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The mode of the backend. It must be one of the following: http or tcp",
			},
			"adv_check": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The adv_check of the backend. It must be one of the following: ssl-hello-chk, smtpchk, ldap-check, mysql-check, pgsql-check, tcp-check, redis-check, httpchk",
			},
			"http_connection_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The http_connection_mode of the backend. It must be one of the following: http-keep-alive, httpclose, http-server-close",
			},
			"server_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The server_timeout of the backend.",
			},
			"check_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The check_timeout of the backend.",
			},
			"connect_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The connect_timeout of the backend.",
			},
			"queue_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The queue_timeout of the backend.",
			},
			"tunnel_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The tunnel_timeout of the backend.",
			},
			"tarpit_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The tarpit_timeout of the backend.",
			},
			"check_cache": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The check_cache of the backend.",
			},
			"balance": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The balance of the backend.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"algorithm": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "roundrobin",
							Description: "The algorithm of the balance.",
						},
					},
				},
			},
			"httpchk_params": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The httpchk_params of the backend.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"method": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The method of the httpchk_params.",
						},
						"uri": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The uri of the httpchk_params.",
						},
						"version": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The version of the httpchk_params.",
						},
					},
				},
			},
			"forwardfor": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The forwardfor of the backend.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The enabled of the forwardfor.",
						},
					},
				},
			},
		},
	}
}

func resourceHaproxyBackendRead(d *schema.ResourceData, m interface{}) error {
	backendName := d.Get("name").(string)

	configMap := m.(map[string]interface{})
	backendConfig := configMap["backend"].(*ConfigBackend)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return backendConfig.GetABackendConfiguration(backendName, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(backendName, "error reading Backend configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(backendName)
	return nil
}

func resourceHaproxyBackendCreate(d *schema.ResourceData, m interface{}) error {
	backendName := d.Get("name").(string)
	checkCache := d.Get("check_cache").(bool)

	// Read values for balance
	balance := d.Get("balance").(*schema.Set).List()
	algorithm := balance[0].(map[string]interface{})["algorithm"].(string)

	// Read values for httpchk_params
	httpchk_params := d.Get("httpchk_params").(*schema.Set).List()
	version := httpchk_params[0].(map[string]interface{})["version"].(string)
	uri := httpchk_params[0].(map[string]interface{})["uri"].(string)
	method := httpchk_params[0].(map[string]interface{})["method"].(string)

	// Read values for forwardfor
	forwardfor := d.Get("forwardfor").(*schema.Set).List()
	enabled := forwardfor[0].(map[string]interface{})["enabled"].(bool)

	payload := BackendPayload{
		Name:               backendName,
		Mode:               d.Get("mode").(string),
		AdvCheck:           d.Get("adv_check").(string),
		HttpConnectionMode: d.Get("http_connection_mode").(string),
		ServerTimeout:      d.Get("server_timeout").(int),
		CheckTimeout:       d.Get("check_timeout").(int),
		ConnectTimeout:     d.Get("connect_timeout").(int),
		QueueTimeout:       d.Get("queue_timeout").(int),
		TunnelTimeout:      d.Get("tunnel_timeout").(int),
		TarpitTimeout:      d.Get("tarpit_timeout").(int),
		CheckCache:         utils.BoolToStr(checkCache),
		Balance: Balance{
			Algorithm: algorithm,  // Access the nested attribute
		},
		HttpchkParams: HttpchkParams{
			Method: method,
			Uri:     uri,
			Version: version,
		},
		Forwardfor: Forwardfor{
			Enabled: utils.BoolToStr(enabled),
		},
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	backendConfig := configMap["backend"].(*ConfigBackend)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return backendConfig.AddBackendConfiguration(payloadJSON, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(backendName, "error creating Backend configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(backendName)
	return nil
}

func resourceHaproxyBackendUpdate(d *schema.ResourceData, m interface{}) error {
	backendName := d.Get("name").(string)
	checkCache := d.Get("check_cache").(bool)

	// Read values for balance
	balance := d.Get("balance").(*schema.Set).List()
	algorithm := balance[0].(map[string]interface{})["algorithm"].(string)

	// Read values for httpchk_params
	httpchk_params := d.Get("httpchk_params").(*schema.Set).List()
	version := httpchk_params[0].(map[string]interface{})["version"].(string)
	uri := httpchk_params[0].(map[string]interface{})["uri"].(string)
	method := httpchk_params[0].(map[string]interface{})["method"].(string)

	// Read values for forwardfor
	forwardfor := d.Get("forwardfor").(*schema.Set).List()
	enabled := forwardfor[0].(map[string]interface{})["enabled"].(bool)

	payload := BackendPayload{
		Name:               backendName,
		Mode:               d.Get("mode").(string),
		AdvCheck:           d.Get("adv_check").(string),
		HttpConnectionMode: d.Get("http_connection_mode").(string),
		ServerTimeout:      d.Get("server_timeout").(int),
		CheckTimeout:       d.Get("check_timeout").(int),
		ConnectTimeout:     d.Get("connect_timeout").(int),
		QueueTimeout:       d.Get("queue_timeout").(int),
		TunnelTimeout:      d.Get("tunnel_timeout").(int),
		TarpitTimeout:      d.Get("tarpit_timeout").(int),
		CheckCache:         utils.BoolToStr(checkCache),
		Balance: Balance{
			Algorithm: algorithm,  // Access the nested attribute
		},
		HttpchkParams: HttpchkParams{
			Method: method,
			Uri:     uri,
			Version: version,
		},
		Forwardfor: Forwardfor{
			Enabled: utils.BoolToStr(enabled),
		},
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	backendConfig := configMap["backend"].(*ConfigBackend)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return backendConfig.UpdateBackendConfiguration(backendName, payloadJSON, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(backendName, "error updating Backend configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(backendName)
	return nil
}

func resourceHaproxyBackendDelete(d *schema.ResourceData, m interface{}) error {
	backendName := d.Get("name").(string)

	configMap := m.(map[string]interface{})
	backendConfig := configMap["backend"].(*ConfigBackend)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return backendConfig.DeleteBackendConfiguration(backendName, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(backendName, "error deleting Backend configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId("")
	return nil
}
