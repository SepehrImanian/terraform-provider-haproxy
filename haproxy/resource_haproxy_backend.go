package haproxy

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"io"
	"net/http"
)

func resourceHaproxyBackend() *schema.Resource {
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

	config := m.(**Config)
	conf := *config

	resp, err := conf.Transaction(func(transactionID string) (*http.Response, error) {
		return conf.GetABackendConfiguration(backendName, transactionID)
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

	config := m.(**Config)
	conf := *config

	resp, err := conf.Transaction(func(transactionID string) (*http.Response, error) {
		return conf.AddBackendConfiguration(payload, transactionID)
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

	config := m.(**Config)
	conf := *config

	resp, err := conf.Transaction(func(transactionID string) (*http.Response, error) {
		return conf.UpdateBackendConfiguration(backendName, payload, transactionID)
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

	config := m.(**Config)
	conf := *config

	resp, err := conf.Transaction(func(transactionID string) (*http.Response, error) {
		return conf.DeleteBackendConfiguration(backendName, transactionID)
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

// GetABackendConfiguration returns the configuration of a backend.
func (c *Config) GetABackendConfiguration(backendName string, TransactionID string) (*http.Response, error) {

	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/backends/%s?transaction_id=%s", c.BaseURL, backendName, TransactionID)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := HTTPRequest("GET", url, nil, headers, c.Username, c.Password)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()
	return resp, nil
}

// AddBackendConfiguration adds a backend configuration.
func (c *Config) AddBackendConfiguration(payload []byte, TransactionID string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/backends?transaction_id=%s", c.BaseURL, TransactionID)
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	fmt.Println("**************** url 400 *******************", url)
	resp, err := HTTPRequest("POST", url, payload, headers, c.Username, c.Password)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	fmt.Println("****************response response 400 ****************", string(body))
	//fmt.Println("^^^^^^^AddBackendConfiguration TransactionID^^^^^^^", TransactionID)

	return resp, nil
}

// DeleteBackendConfiguration deletes a backend configuration.
func (c *Config) DeleteBackendConfiguration(backendName string, TransactionID string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/backends/%s?transaction_id=%s", c.BaseURL, backendName, TransactionID)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := HTTPRequest("DELETE", url, nil, headers, c.Username, c.Password)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()
	return resp, nil
}

// UpdateBackendConfiguration updates a backend configuration.
func (c *Config) UpdateBackendConfiguration(backendName string, payload []byte, TransactionID string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/backends/%s?transaction_id=%s", c.BaseURL, backendName, TransactionID)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := HTTPRequest("PUT", url, payload, headers, c.Username, c.Password)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()
	return resp, nil
}
