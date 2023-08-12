package haproxy

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"io"
	"net/http"
)

func resourceHaproxyFrontend() *schema.Resource {
	return &schema.Resource{
		Create: resourceHaproxyFrontendCreate,
		Read:   resourceHaproxyFrontendRead,
		Update: resourceHaproxyFrontendUpdate,
		Delete: resourceHaproxyFrontendDelete,

		Schema: map[string]*schema.Schema{
			"Frontend_name": {
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

func resourceHaproxyFrontendRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceHaproxyFrontendCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceHaproxyFrontendUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceHaproxyFrontendDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

// GetAFrontendConfiguration returns the configuration of a Frontend.
func (c *Config) GetAFrontendConfiguration(FrontendName string, TransactionID string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/Frontends/%s?transaction_id=%s", c.BaseURL, FrontendName, TransactionID)
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

// AddFrontendConfiguration adds a Frontend configuration.
func (c *Config) AddFrontendConfiguration(payload []byte, TransactionID string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/Frontends?transaction_id=%s", c.BaseURL, TransactionID)
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
	//fmt.Println("^^^^^^^AddFrontendConfiguration TransactionID^^^^^^^", TransactionID)

	return resp, nil
}

// DeleteFrontendConfiguration deletes a Frontend configuration.
func (c *Config) DeleteFrontendConfiguration(FrontendName string, TransactionID string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/Frontends/%s?transaction_id=%s", c.BaseURL, FrontendName, TransactionID)
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

// UpdateFrontendConfiguration updates a Frontend configuration.
func (c *Config) UpdateFrontendConfiguration(FrontendName string, payload []byte, TransactionID string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/Frontends/%s?transaction_id=%s", c.BaseURL, FrontendName, TransactionID)
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
