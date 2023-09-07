package ServerTemplate

import (
	"fmt"
	"net/http"

	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ResourceHaproxyServerTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceHaproxyServerTemplateCreate,
		Read:   resourceHaproxyServerTemplateRead,
		Update: resourceHaproxyServerTemplateUpdate,
		Delete: resourceHaproxyServerTemplateDelete,

		Schema: map[string]*schema.Schema{
			"prefix": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the ServerTemplate.",
			},
			"backend": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The backend of the ServerTemplate.",
			},
			"fqdn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The address of the ServerTemplate",
			},
			"num_or_range": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The number or range of the ServerTemplate",
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The port of the ServerTemplate",
			},
		},
	}
}

func resourceHaproxyServerTemplateRead(d *schema.ResourceData, m interface{}) error {
	prefix := d.Get("prefix").(string)
	backend := d.Get("backend").(string)

	configMap := m.(map[string]interface{})
	ServerTemplateConfig := configMap["ServerTemplate"].(*ConfigServerTemplate)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, _, err := tranConfig.TransactionWithData(func(transactionID string) (*http.Response, []byte, error) {
		return ServerTemplateConfig.GetAServerTemplatesConfiguration(prefix, transactionID, backend)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(prefix, "error reading ServerTemplate configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(prefix)
	return nil
}

func resourceHaproxyServerTemplateCreate(d *schema.ResourceData, m interface{}) error {
	prefix := d.Get("prefix").(string)
	backend := d.Get("backend").(string)

	payload := ServerTemplatePayload{
		Prefix:     prefix,
		Fqdn:       d.Get("fqdn").(string),
		NumOrRange: d.Get("num_or_range").(string),
		Port:       d.Get("port").(int),
	}

	payloadJSON, err := utils.MarshalNonZeroFields(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	ServerTemplateConfig := configMap["ServerTemplate"].(*ConfigServerTemplate)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, _, err := tranConfig.TransactionWithData(func(transactionID string) (*http.Response, []byte, error) {
		return ServerTemplateConfig.AddServerTemplatesConfiguration(payloadJSON, transactionID, backend)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(prefix, "error creating ServerTemplate configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(prefix)
	return nil
}

func resourceHaproxyServerTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	prefix := d.Get("prefix").(string)
	backend := d.Get("backend").(string)

	payload := ServerTemplatePayload{
		Prefix:     prefix,
		Fqdn:       d.Get("fqdn").(string),
		NumOrRange: d.Get("num_or_range").(string),
		Port:       d.Get("port").(int),
	}

	payloadJSON, err := utils.MarshalNonZeroFields(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	ServerTemplateConfig := configMap["ServerTemplate"].(*ConfigServerTemplate)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, _, err := tranConfig.TransactionWithData(func(transactionID string) (*http.Response, []byte, error) {
		return ServerTemplateConfig.UpdateServerTemplatesConfiguration(prefix, payloadJSON, transactionID, backend)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(prefix, "error creating ServerTemplate configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(prefix)
	return nil
}

func resourceHaproxyServerTemplateDelete(d *schema.ResourceData, m interface{}) error {
	prefix := d.Get("prefix").(string)
	backend := d.Get("backend").(string)

	configMap := m.(map[string]interface{})
	ServerTemplateConfig := configMap["ServerTemplate"].(*ConfigServerTemplate)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, _, err := tranConfig.TransactionWithData(func(transactionID string) (*http.Response, []byte, error) {
		return ServerTemplateConfig.DeleteServerTemplatesConfiguration(prefix, transactionID, backend)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(prefix, "error deleting ServerTemplate configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId("")
	return nil
}
