package frontend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ResourceHaproxyFrontend() *schema.Resource {
	return &schema.Resource{
		Create: resourceHaproxyFrontendCreate,
		Read:   resourceHaproxyFrontendRead,
		Update: resourceHaproxyFrontendUpdate,
		Delete: resourceHaproxyFrontendDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the frontend. It must be unique and cannot be changed.",
			},
			"backend": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the backend.",
			},
			"http_connection_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The http connection mode of the frontend. It can be one of the following values: httpclose, http-server-close, http-keep-alive",
			},
			"accept_invalid_http_request": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The accept invalid http request of the frontend.",
			},
			"maxconn": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The max connection of the frontend.",
			},
			"mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The mode of the frontend. It can be one of the following values: http, tcp",
			},
			"backlog": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The backlog of the frontend.",
			},
			"http_keep_alive_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The http keep alive timeout of the frontend.",
			},
			"http_request_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The http request timeout of the frontend.",
			},
			"http_use_proxy_header": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The http use proxy header of the frontend.",
			},
			"httplog": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The http log of the frontend.",
			},
			"httpslog": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The https log of the frontend.",
			},
			"error_log_format": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The error log format of the frontend.",
			},
			"log_format": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The log format of the frontend.",
			},
			"log_format_sd": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The log format sd of the frontend.",
			},
			"monitor_uri": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The monitor uri of the frontend.",
			},
			"tcplog": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The tcp log of the frontend.",
			},
			"compression": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The compression of the frontend.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"algorithms": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The algorithms of the compression.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"offload": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The offload of the compression.",
						},
						"types": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The types of the compression.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"forwardfor": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The forwardfor of the frontend.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The enabled of the forwardfor.",
						},
						"except": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The except of the forwardfor.",
						},
						"header": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The header of the forwardfor.",
						},
						"ifnone": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The ifnone of the forwardfor.",
						},
					},
				},
			},
		},
	}
}

func resourceHaproxyFrontendRead(d *schema.ResourceData, m interface{}) error {
	frontendName := d.Get("name").(string)

	configMap := m.(map[string]interface{})
	frontendConfig := configMap["frontend"].(*ConfigFrontend)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return frontendConfig.GetAFrontendConfiguration(frontendName, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(frontendName, "error reading Frontend configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(frontendName)
	return nil
}

func resourceHaproxyFrontendCreate(d *schema.ResourceData, m interface{}) error {
	frontendName := d.Get("name").(string)
	acceptInvalidHttpRequest := d.Get("accept_invalid_http_request").(bool)
	httpslog := d.Get("httpslog").(bool)
	httpUseProxyHeader := d.Get("http_use_proxy_header").(bool)

	var (
		compressionOffload    bool = false
		forwardforEnabled     bool = false
		forwardforExcept      string
		forwardforHeader      string
		forwardforIfnone      bool = false
		compressionAlgorithms []string
		compressionTypes      []string
		enabledStr            string = "enabled"
	)

	compressionItem := utils.GetFirstItemValue(d.Get, "compression")
	if compressionItem != nil {
		// Read the compression block
		compression := d.Get("compression").(*schema.Set).List()
		compressionData := compression[0].(map[string]interface{})
		compressionAlgorithmsRaw := compressionData["algorithms"].([]interface{})
		for _, algorithm := range compressionAlgorithmsRaw {
			compressionAlgorithms = append(compressionAlgorithms, algorithm.(string))
		}

		compressionOffload = compressionData["offload"].(bool)

		// Corrected handling of the 'types' attribute
		compressionTypesRaw := compressionData["types"].([]interface{})
		for _, t := range compressionTypesRaw {
			compressionTypes = append(compressionTypes, t.(string))
		}
	}

	forwardforItem := utils.GetFirstItemValue(d.Get, "forwardfor")
	if forwardforItem != nil {
		//Read the forwardfor block
		forwardfor := d.Get("forwardfor").(*schema.Set).List()
		forwardforEnabled = forwardfor[0].(map[string]interface{})["enabled"].(bool)
		forwardforExcept = forwardfor[0].(map[string]interface{})["except"].(string)
		forwardforHeader = forwardfor[0].(map[string]interface{})["header"].(string)
		forwardforIfnone = forwardfor[0].(map[string]interface{})["ifnone"].(bool)
		enabledStr = utils.BoolToStr(forwardforEnabled)
	}

	payload := FrontendPayload{
		Name:                     frontendName,
		DefaultBackend:           d.Get("backend").(string),
		HttpConnectionMode:       d.Get("http_connection_mode").(string),
		AcceptInvalidHttpRequest: utils.BoolToStr(acceptInvalidHttpRequest),
		MaxConn:                  d.Get("maxconn").(int),
		Mode:                     d.Get("mode").(string),
		Backlog:                  d.Get("backlog").(int),
		HttpKeepAliveTimeout:     d.Get("http_keep_alive_timeout").(int),
		HttpRequestTimeout:       d.Get("http_request_timeout").(int),
		HttpUseProxyHeader:       utils.BoolToStr(httpUseProxyHeader),
		HttpLog:                  d.Get("httplog").(bool),
		HttpsLog:                 utils.BoolToStr(httpslog),
		ErrorLogFormat:           d.Get("error_log_format").(string),
		LogFormat:                d.Get("log_format").(string),
		LogFormatSd:              d.Get("log_format_sd").(string),
		MonitorUri:               d.Get("monitor_uri").(string),
		TcpLog:                   d.Get("tcplog").(bool),
		Compression: Compression{
			Algorithms: compressionAlgorithms,
			Offload:    compressionOffload,
			Types:      compressionTypes,
		},
		Forwardfor: Forwardfor{
			Enabled: enabledStr,
			Except:  forwardforExcept,
			Header:  forwardforHeader,
			Ifnone:  forwardforIfnone,
		},
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	frontendConfig := configMap["frontend"].(*ConfigFrontend)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return frontendConfig.AddFrontendConfiguration(payloadJSON, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(frontendName, "error creating Frontend configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(frontendName)
	return nil
}

func resourceHaproxyFrontendUpdate(d *schema.ResourceData, m interface{}) error {
	frontendName := d.Get("name").(string)
	acceptInvalidHttpRequest := d.Get("accept_invalid_http_request").(bool)
	httpslog := d.Get("httpslog").(bool)
	httpUseProxyHeader := d.Get("http_use_proxy_header").(bool)

	var (
		compressionOffload    bool = false
		forwardforEnabled     bool = false
		forwardforExcept      string
		forwardforHeader      string
		forwardforIfnone      bool = false
		compressionAlgorithms []string
		compressionTypes      []string
		enabledStr            string = "enabled"
	)

	compressionItem := utils.GetFirstItemValue(d.Get, "compression")
	if compressionItem != nil {
		// Read the compression block
		compression := d.Get("compression").(*schema.Set).List()
		compressionData := compression[0].(map[string]interface{})
		compressionAlgorithmsRaw := compressionData["algorithms"].([]interface{})
		for _, algorithm := range compressionAlgorithmsRaw {
			compressionAlgorithms = append(compressionAlgorithms, algorithm.(string))
		}

		compressionOffload = compressionData["offload"].(bool)

		// Corrected handling of the 'types' attribute
		compressionTypesRaw := compressionData["types"].([]interface{})
		for _, t := range compressionTypesRaw {
			compressionTypes = append(compressionTypes, t.(string))
		}
	}

	forwardforItem := utils.GetFirstItemValue(d.Get, "forwardfor")
	if forwardforItem != nil {
		//Read the forwardfor block
		forwardfor := d.Get("forwardfor").(*schema.Set).List()
		forwardforEnabled = forwardfor[0].(map[string]interface{})["enabled"].(bool)
		forwardforExcept = forwardfor[0].(map[string]interface{})["except"].(string)
		forwardforHeader = forwardfor[0].(map[string]interface{})["header"].(string)
		forwardforIfnone = forwardfor[0].(map[string]interface{})["ifnone"].(bool)
		enabledStr = utils.BoolToStr(forwardforEnabled)
	}

	payload := FrontendPayload{
		Name:                     frontendName,
		DefaultBackend:           d.Get("backend").(string),
		HttpConnectionMode:       d.Get("http_connection_mode").(string),
		AcceptInvalidHttpRequest: utils.BoolToStr(acceptInvalidHttpRequest),
		MaxConn:                  d.Get("maxconn").(int),
		Mode:                     d.Get("mode").(string),
		Backlog:                  d.Get("backlog").(int),
		HttpKeepAliveTimeout:     d.Get("http_keep_alive_timeout").(int),
		HttpRequestTimeout:       d.Get("http_request_timeout").(int),
		HttpUseProxyHeader:       utils.BoolToStr(httpUseProxyHeader),
		HttpLog:                  d.Get("httplog").(bool),
		HttpsLog:                 utils.BoolToStr(httpslog),
		ErrorLogFormat:           d.Get("error_log_format").(string),
		LogFormat:                d.Get("log_format").(string),
		LogFormatSd:              d.Get("log_format_sd").(string),
		MonitorUri:               d.Get("monitor_uri").(string),
		TcpLog:                   d.Get("tcplog").(bool),
		Compression: Compression{
			Algorithms: compressionAlgorithms,
			Offload:    compressionOffload,
			Types:      compressionTypes,
		},
		Forwardfor: Forwardfor{
			Enabled: enabledStr,
			Except:  forwardforExcept,
			Header:  forwardforHeader,
			Ifnone:  forwardforIfnone,
		},
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	frontendConfig := configMap["frontend"].(*ConfigFrontend)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return frontendConfig.UpdateFrontendConfiguration(frontendName, payloadJSON, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(frontendName, "error updating Frontend configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(frontendName)
	return nil
}

func resourceHaproxyFrontendDelete(d *schema.ResourceData, m interface{}) error {
	frontendName := d.Get("name").(string)

	configMap := m.(map[string]interface{})
	frontendConfig := configMap["frontend"].(*ConfigFrontend)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return frontendConfig.DeleteFrontendConfiguration(frontendName, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(frontendName, "error deleting Frontend configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId("")
	return nil
}
