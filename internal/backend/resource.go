package backend

import (
	"fmt"
	"net/http"

	"terraform-provider-haproxy/internal/transaction"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ResourceHaproxyBackend() *schema.Resource {
	return &schema.Resource{
		Create: resourceHaproxyBackendCreate,
		Read:   resourceHaproxyBackendRead,
		Update: resourceHaproxyBackendUpdate,
		Delete: resourceHaproxyBackendDelete,

		Schema: map[string]*schema.Schema{
			"backend_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "http",
			},
			"balance_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "roundrobin",
			},
		},
	}
}

func resourceHaproxyBackendRead(d *schema.ResourceData, m interface{}) error {
	backendName := d.Get("backend_name").(string)

	configMap := m.(map[string]interface{})
	backendConfig := configMap["backend"].(*ConfigBackend)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return backendConfig.GetABackendConfiguration(backendName, transactionID)
	})

	if err != nil {
		fmt.Println("Error updating backend configuration:", err)
		return err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return fmt.Errorf("error creating backend configuration: %s", resp.Status)
	}

	d.SetId(backendName)
	return nil
}

func resourceHaproxyBackendCreate(d *schema.ResourceData, m interface{}) error {
	backendName := d.Get("backend_name").(string)
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

	if err != nil {
		fmt.Println("Error creating backend configuration:", err)
		return err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return fmt.Errorf("error creating backend configuration: %s", resp.Status)
	}

	d.SetId(backendName)
	return nil
}

func resourceHaproxyBackendUpdate(d *schema.ResourceData, m interface{}) error {
	backendName := d.Get("backend_name").(string)
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

	if err != nil {
		fmt.Println("Error updating backend configuration:", err)
		return err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return fmt.Errorf("error creating backend configuration: %s", resp)
	}
	d.SetId(backendName)
	return nil
}

func resourceHaproxyBackendDelete(d *schema.ResourceData, m interface{}) error {
	backendName := d.Get("backend_name").(string)

	configMap := m.(map[string]interface{})
	backendConfig := configMap["backend"].(*ConfigBackend)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return backendConfig.DeleteBackendConfiguration(backendName, transactionID)
	})

	if err != nil {
		fmt.Println("Error updating backend configuration:", err)
		return err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return fmt.Errorf("error creating backend configuration: %s", resp.Status)
	}

	d.SetId("")
	return nil
}
