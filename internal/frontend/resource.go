package frontend

import (
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
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"backend": {
				Type:     schema.TypeString,
				Required: true,
			},
			"http_connection_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "http-keep-alive",
			},
			"max_connection": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
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
	backend := d.Get("backend").(string)
	httpConnectionMode := d.Get("http_connection_mode").(string)
	maxConnection := d.Get("max_connection").(int)
	mode := d.Get("mode").(string)

	payload := []byte(fmt.Sprintf(`
	{
		"default_backend": "%s",
		"http_connection_mode": "%s",
		"maxconn": %d,
		"mode": "%s",
		"name": "%s"
	}
	`, backend, httpConnectionMode, maxConnection, mode, frontendName))

	configMap := m.(map[string]interface{})
	frontendConfig := configMap["frontend"].(*ConfigFrontend)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return frontendConfig.AddFrontendConfiguration(payload, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(frontendName, "error creating Frontend configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(frontendName)
	return nil
}

func resourceHaproxyFrontendUpdate(d *schema.ResourceData, m interface{}) error {
	frontendName := d.Get("name").(string)
	backend := d.Get("backend").(string)
	httpConnectionMode := d.Get("http_connection_mode").(string)
	maxConnection := d.Get("max_connection").(int)
	mode := d.Get("mode").(string)

	//maxConnectionStr := strconv.Itoa(maxConnection)

	payload := []byte(fmt.Sprintf(`
	{
		"default_backend": "%s",
		"http_connection_mode": "%s",
		"maxconn": %d,
		"mode": "%s",
		"name": "%s"
	}
	`, backend, httpConnectionMode, maxConnection, mode, frontendName))

	configMap := m.(map[string]interface{})
	frontendConfig := configMap["frontend"].(*ConfigFrontend)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return frontendConfig.UpdateFrontendConfiguration(frontendName, payload, transactionID)
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
