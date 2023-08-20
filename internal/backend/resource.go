package backend

import (
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
				Description: "The name of the backend. It must be unique and cannot be changed.",
			},
			"mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "http",
				Description: "The mode of the backend. It must be one of the following: http or tcp",
			},
			"balance_algorithm": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "roundrobin",
				Description: "The balance algorithm of the backend. It must be one of the following: roundrobin, static-rr, leastconn, first, source, uri, url_param, hdr, random, rdp-cookie, hash",
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
	mode := d.Get("mode").(string)
	balanceAlgorithm := d.Get("balance_algorithm").(string)

	payload := []byte(fmt.Sprintf(`
	{
	  "adv_check": "httpchk",
	  "balance": {
		"algorithm": "%s"
	  },
	  "forwardfor": {
		"enabled": "enabled"
	  },
	  "httpchk_params": {
		"method": "GET",
		"uri": "/check",
		"version": "HTTP/1.1"
	  },
	  "mode": "%s",
	  "name": "%s"
	}
	`, balanceAlgorithm, mode, backendName))

	configMap := m.(map[string]interface{})
	backendConfig := configMap["backend"].(*ConfigBackend)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return backendConfig.AddBackendConfiguration(payload, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(backendName, "error creating Backend configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(backendName)
	return nil
}

func resourceHaproxyBackendUpdate(d *schema.ResourceData, m interface{}) error {
	backendName := d.Get("name").(string)
	mode := d.Get("mode").(string)
	balanceAlgorithm := d.Get("balance_algorithm").(string)

	payload := []byte(fmt.Sprintf(`
	{
	  "adv_check": "httpchk",
	  "balance": {
		"algorithm": "%s"
	  },
	  "forwardfor": {
		"enabled": "enabled"
	  },
	  "httpchk_params": {
		"method": "GET",
		"uri": "/check",
		"version": "HTTP/1.1"
	  },
	  "mode": "%s",
	  "name": "%s"
	}
	`, balanceAlgorithm, mode, backendName))

	configMap := m.(map[string]interface{})
	backendConfig := configMap["backend"].(*ConfigBackend)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return backendConfig.UpdateBackendConfiguration(backendName, payload, transactionID)
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
