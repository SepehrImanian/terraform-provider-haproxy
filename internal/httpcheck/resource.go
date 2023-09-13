package httpcheck

import (
	"fmt"
	"net/http"

	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ResourceHaproxyHttpcheck() *schema.Resource {
	return &schema.Resource{
		Create: resourceHaproxyHttpCheckCreate,
		Read:   resourceHaproxyHttpCheckRead,
		Update: resourceHaproxyHttpCheckUpdate,
		Delete: resourceHaproxyHttpCheckDelete,

		Schema: map[string]*schema.Schema{
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
			"index": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The index of the HttpCheck in the parent object starting at 0",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of the check",
			},
			"addr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The address to connect to",
			},
			"alpn": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ALPN protocol to use",
			},
			"body": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The body to send",
			},
			"body_log_format": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The log format to use for the body",
			},
			"check_comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The comment to add to the check",
			},
			"default": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Set this check as the default one",
			},
			"error_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status to return on error",
			},
			"exclamation_mark": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Add an exclamation mark to the check",
			},
			"headers": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The headers to send",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the header",
						},
						"fmt": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The format of the header",
						},
					},
				},
			},
			"linger": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable linger",
			},
			"match": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The match to use",
			},
			"method": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The method to use",
			},
			"min_recv": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The minimum number of bytes to receive",
			},
			"ok_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status to return on success",
			},
			"on_error": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The action to take on error",
			},
			"on_success": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The action to take on success",
			},
			"pattern": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The pattern to use",
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The port to connect to",
			},
			"port_string": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The port to connect to as a string",
			},
			"proto": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The protocol to use",
			},
			"send_proxy": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Send the proxy protocol",
			},
			"sni": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The SNI to use",
			},
			"ssl": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable SSL",
			},
			"status_code": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status code to use",
			},
			"uri": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The URI to use",
			},
			"uri_log_format": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The log format to use for the URI",
			},
			"var_expr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The expression to use for the variable",
			},
			"var_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the variable",
			},
			"var_format": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The format of the variable",
			},
			"var_scope": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The scope of the variable",
			},
			"version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The version to use",
			},
			"via_socks4": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Use SOCKS4",
			},
		},
	}
}

func resourceHaproxyHttpCheckRead(d *schema.ResourceData, m interface{}) error {
	indexName := d.Get("index").(int)
	indexNameStr := fmt.Sprintf("%d", indexName)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	configMap := m.(map[string]interface{})
	HttpCheckConfig := configMap["httpcheck"].(*ConfigHttpCheck)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return HttpCheckConfig.GetAHttpCheckConfiguration(indexName, transactionID, parentName, parentType)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(indexNameStr, "error reading HttpCheck configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(indexNameStr)
	return nil
}

func resourceHaproxyHttpCheckCreate(d *schema.ResourceData, m interface{}) error {
	indexName := d.Get("index").(int)
	indexNameStr := fmt.Sprintf("%d", indexName)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	payload := HttpCheckPayload{
		Index:           indexName,
		Addr:            d.Get("addr").(string),
		Alpn:            d.Get("alpn").(string),
		Body:            d.Get("body").(string),
		BodyLogFormat:   d.Get("body_log_format").(string),
		CheckComment:    d.Get("check_comment").(string),
		Default:         d.Get("default").(bool),
		ErrorStatus:     d.Get("error_status").(string),
		ExclamationMark: d.Get("exclamation_mark").(bool),
		Linger:          d.Get("linger").(bool),
		Match:           d.Get("match").(string),
		Method:          d.Get("method").(string),
		MinRecv:         d.Get("min_recv").(int),
		OkStatus:        d.Get("ok_status").(string),
		OnError:         d.Get("on_error").(string),
		OnSuccess:       d.Get("on_success").(string),
		Pattern:         d.Get("pattern").(string),
		Port:            d.Get("port").(int),
		PortString:      d.Get("port_string").(string),
		Proto:           d.Get("proto").(string),
		SendProxy:       d.Get("send_proxy").(bool),
		Sni:             d.Get("sni").(string),
		Ssl:             d.Get("ssl").(bool),
		StatusCode:      d.Get("status_code").(string),
		Type:            d.Get("type").(string),
		Uri:             d.Get("uri").(string),
		UriLogFormat:    d.Get("uri_log_format").(string),
		VarExpr:         d.Get("var_expr").(string),
		VarName:         d.Get("var_name").(string),
		VarFormat:       d.Get("var_format").(string),
		VarScope:        d.Get("var_scope").(string),
		Version:         d.Get("version").(string),
		ViaSocks4:       d.Get("via_socks4").(bool),
	}

	payloadJSON, err := utils.MarshalNonZeroFields(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	HttpCheckConfig := configMap["httpcheck"].(*ConfigHttpCheck)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return HttpCheckConfig.AddHttpCheckConfiguration(payloadJSON, transactionID, parentName, parentType)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(indexNameStr, "error creating HttpCheck configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(indexNameStr)
	return nil
}

func resourceHaproxyHttpCheckUpdate(d *schema.ResourceData, m interface{}) error {
	indexName := d.Get("index").(int)
	indexNameStr := fmt.Sprintf("%d", indexName)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	payload := HttpCheckPayload{
		Index:           indexName,
		Addr:            d.Get("addr").(string),
		Alpn:            d.Get("alpn").(string),
		Body:            d.Get("body").(string),
		BodyLogFormat:   d.Get("body_log_format").(string),
		CheckComment:    d.Get("check_comment").(string),
		Default:         d.Get("default").(bool),
		ErrorStatus:     d.Get("error_status").(string),
		ExclamationMark: d.Get("exclamation_mark").(bool),
		Linger:          d.Get("linger").(bool),
		Match:           d.Get("match").(string),
		Method:          d.Get("method").(string),
		MinRecv:         d.Get("min_recv").(int),
		OkStatus:        d.Get("ok_status").(string),
		OnError:         d.Get("on_error").(string),
		OnSuccess:       d.Get("on_success").(string),
		Pattern:         d.Get("pattern").(string),
		Port:            d.Get("port").(int),
		PortString:      d.Get("port_string").(string),
		Proto:           d.Get("proto").(string),
		SendProxy:       d.Get("send_proxy").(bool),
		Sni:             d.Get("sni").(string),
		Ssl:             d.Get("ssl").(bool),
		StatusCode:      d.Get("status_code").(string),
		Type:            d.Get("type").(string),
		Uri:             d.Get("uri").(string),
		UriLogFormat:    d.Get("uri_log_format").(string),
		VarExpr:         d.Get("var_expr").(string),
		VarName:         d.Get("var_name").(string),
		VarFormat:       d.Get("var_format").(string),
		VarScope:        d.Get("var_scope").(string),
		Version:         d.Get("version").(string),
		ViaSocks4:       d.Get("via_socks4").(bool),
	}

	payloadJSON, err := utils.MarshalNonZeroFields(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	HttpCheckConfig := configMap["httpcheck"].(*ConfigHttpCheck)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return HttpCheckConfig.UpdateHttpCheckConfiguration(indexName, payloadJSON, transactionID, parentName, parentType)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(indexNameStr, "error updating HttpCheck configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(indexNameStr)
	return nil
}

func resourceHaproxyHttpCheckDelete(d *schema.ResourceData, m interface{}) error {
	indexName := d.Get("index").(int)
	indexNameStr := fmt.Sprintf("%d", indexName)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	configMap := m.(map[string]interface{})
	HttpCheckConfig := configMap["httpcheck"].(*ConfigHttpCheck)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return HttpCheckConfig.DeleteHttpCheckConfiguration(indexName, transactionID, parentName, parentType)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(indexNameStr, "error deleting HttpCheck configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId("")
	return nil
}
