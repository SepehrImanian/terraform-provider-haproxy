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
			"source": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "	",
			},
		},
	}
}

func resourceHaproxyBackendCreate(d *schema.ResourceData, m interface{}) error {
	backendName := d.Get("backend_name").(string)
	mode := d.Get("mode").(string)
	balanceAlgorithm := d.Get("balance_algorithm").(string)
	source := d.Get("source").(string)

	payload := []byte(fmt.Sprintf(`
	{
		"mode": "%s",
		"balance_algorithm": "%s",
		"source": "%s"
	}
	`, mode, balanceAlgorithm, source))

	config := m.(*BackendConfig)
	resp, err := config.AddBackendConfiguration(payload)
	if err != nil {
		fmt.Println("Error creating backend configuration:", err)
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("Error creating backend configuration: %s", resp.Status)
	}

	d.SetId(backendName)
	return nil
}

func resourceHaproxyBackendUpdate(d *schema.ResourceData, m interface{}) error {
	backendName := d.Get("backend_name").(string)
	mode := d.Get("mode").(string)
	balanceAlgorithm := d.Get("balance_algorithm").(string)
	source := d.Get("source").(string)

	payload := []byte(fmt.Sprintf(`
	{
		"mode": "%s",
		"balance_algorithm": "%s",
		"source": "%s"
	}
	`, mode, balanceAlgorithm, source))

	config := m.(*BackendConfig)
	resp, err := config.UpdateBackendConfiguration(backendName, payload)
	if err != nil {
		fmt.Println("Error updating backend configuration:", err)
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("Error updating backend configuration: %s", resp.Status)
	}

	d.SetId(backendName)
	return nil
}

func resourceHaproxyBackendDelete(d *schema.ResourceData, m interface{}) error {
	backendName := d.Get("backend_name").(string)
	config := m.(*BackendConfig)
	resp, err := config.DeleteBackendConfiguration(backendName)
	if err != nil {
		fmt.Println("Error updating backend configuration:", err)
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("Error updating backend configuration: %s", resp.Status)
	}
	d.SetId("")
	return nil
}

// GetABackendConfiguration returns the configuration of a backend.
func (c *BackendConfig) GetABackendConfiguration(backendName string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/backends/%s?transaction_id=%s", c.BaseURL, backendName, c.TransactionID)
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Basic " + encodeCredentials(c.Username, c.Password),
	}
	resp, err := HTTPRequest("GET", url, nil, headers)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(resp.Body)
	return resp, nil
}

// AddBackendConfiguration adds a backend configuration.
func (c *BackendConfig) AddBackendConfiguration(payload []byte) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/backends?transaction_id=%s", c.BaseURL, c.TransactionID)
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Basic " + encodeCredentials(c.Username, c.Password),
	}
	resp, err := HTTPRequest("POST", url, payload, headers)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(resp.Body)
	return resp, nil
}

// DeleteBackendConfiguration deletes a backend configuration.
func (c *BackendConfig) DeleteBackendConfiguration(backendName string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/backends/%s?transaction_id=%s", c.BaseURL, backendName, c.TransactionID)
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Basic " + encodeCredentials(c.Username, c.Password),
	}
	resp, err := HTTPRequest("DELETE", url, nil, headers)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(resp.Body)
	return resp, nil
}

// UpdateBackendConfiguration updates a backend configuration.
func (c *BackendConfig) UpdateBackendConfiguration(backendName string, payload []byte) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/backends/%s?transaction_id=%s", c.BaseURL, backendName, c.TransactionID)
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Basic " + encodeCredentials(c.Username, c.Password),
	}
	resp, err := HTTPRequest("PUT", url, payload, headers)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(resp.Body)
	return resp, nil
}
