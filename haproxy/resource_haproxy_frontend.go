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

	config := m.(**Config)
	conf := *config

	resp, err := conf.Transaction(func(transactionID string) (*http.Response, error) {
		return conf.GetAFrontendConfiguration(frontendName, transactionID)
	})

	if err != nil {
		fmt.Println("Error updating frontend configuration:", err)
		return err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return fmt.Errorf("error creating frontend configuration: %s", resp.Status)
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

	config := m.(**Config)
	conf := *config

	resp, err := conf.Transaction(func(transactionID string) (*http.Response, error) {
		return conf.AddFrontendConfiguration(payload, transactionID)
	})

	if err != nil {
		fmt.Println("Error creating frontend configuration:", err)
		return err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return fmt.Errorf("error creating frontend configuration: %s", resp.Status)
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

	config := m.(**Config)
	conf := *config

	resp, err := conf.Transaction(func(transactionID string) (*http.Response, error) {
		return conf.UpdateFrontendConfiguration(frontendName, payload, transactionID)
	})

	if err != nil {
		fmt.Println("Error creating frontend configuration:", err)
		return err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return fmt.Errorf("error creating frontend configuration: %s", resp.Status)
	}

	d.SetId(frontendName)
	return nil
}

func resourceHaproxyFrontendDelete(d *schema.ResourceData, m interface{}) error {
	frontendName := d.Get("name").(string)

	config := m.(**Config)
	conf := *config

	resp, err := conf.Transaction(func(transactionID string) (*http.Response, error) {
		return conf.DeleteFrontendConfiguration(frontendName, transactionID)
	})

	if err != nil {
		fmt.Println("Error updating frontend configuration:", err)
		return err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return fmt.Errorf("error creating frontend configuration: %s", resp.Status)
	}

	d.SetId("")
	return nil
}

// GetAFrontendConfiguration returns the configuration of a Frontend.
func (c *Config) GetAFrontendConfiguration(FrontendName string, TransactionID string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/frontends/%s?transaction_id=%s", c.BaseURL, FrontendName, TransactionID)
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
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/frontends?transaction_id=%s", c.BaseURL, TransactionID)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
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
	return resp, nil
}

// DeleteFrontendConfiguration deletes a Frontend configuration.
func (c *Config) DeleteFrontendConfiguration(FrontendName string, TransactionID string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/frontends/%s?transaction_id=%s", c.BaseURL, FrontendName, TransactionID)
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
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/frontends/%s?transaction_id=%s", c.BaseURL, FrontendName, TransactionID)
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
