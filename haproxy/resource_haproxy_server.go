package haproxy

import (
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceHaproxyServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceHaproxyServerCreate,
		Read:   resourceHaproxyServerRead,
		Update: resourceHaproxyServerUpdate,
		Delete: resourceHaproxyServerDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"parent_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"parent_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"send_proxy": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"check": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"inter": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rise": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fall": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceHaproxyServerRead(d *schema.ResourceData, m interface{}) error {
	serverName := d.Get("name").(string)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	config := m.(**Config)
	conf := *config

	resp, err := conf.Transaction(func(transactionID string) (*http.Response, error) {
		return conf.GetAServerConfiguration(serverName, transactionID, parentName, parentType)
	})

	if err != nil {
		fmt.Println("Error updating Server configuration:", err)
		return err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return fmt.Errorf("error creating Server configuration: %s", resp.Status)
	}

	d.SetId(serverName)
	return nil
}

func resourceHaproxyServerCreate(d *schema.ResourceData, m interface{}) error {
	serverName := d.Get("name").(string)
	port := d.Get("port").(int)
	address := d.Get("address").(string)
	sendProxy := d.Get("send_proxy").(bool)
	check := d.Get("check").(bool)
	inter := d.Get("inter").(int)
	rise := d.Get("rise").(int)
	fall := d.Get("fall").(int)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	// change bool to enabled/disabled
	sendProxyStr := boolToStr(sendProxy)
	checkStr := boolToStr(check)

	payload := []byte(fmt.Sprintf(`
	{
		"name": "%s",
		"address": "%s",
		"port": %d,
		"send-proxy": "%s",
		"check": "%s",
		"inter": %d,
		"rise": %d,
		"fall": %d
	}
	`, serverName, address, port, sendProxyStr, checkStr, inter, rise, fall))

	config := m.(**Config)
	conf := *config

	resp, err := conf.Transaction(func(transactionID string) (*http.Response, error) {
		return conf.AddServerConfiguration(payload, transactionID, parentName, parentType)
	})

	if err != nil {
		fmt.Println("Error creating Server configuration:", err)
		return err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return fmt.Errorf("error creating Server configuration: %s", resp.Status)
	}

	d.SetId(serverName)
	return nil
}

func resourceHaproxyServerUpdate(d *schema.ResourceData, m interface{}) error {
	serverName := d.Get("name").(string)
	port := d.Get("port").(int)
	address := d.Get("address").(string)
	sendProxy := d.Get("send_proxy").(bool)
	check := d.Get("check").(bool)
	inter := d.Get("inter").(int)
	rise := d.Get("rise").(int)
	fall := d.Get("fall").(int)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	// change bool to enabled/disabled
	sendProxyStr := boolToStr(sendProxy)
	checkStr := boolToStr(check)

	payload := []byte(fmt.Sprintf(`
	{
		"name": "%s",
		"address": "%s",
		"port": %d,
		"send-proxy": "%s",
		"check": "%s",
		"inter": %d,
		"rise": %d,
		"fall": %d
	}
	`, serverName, address, port, sendProxyStr, checkStr, inter, rise, fall))

	config := m.(**Config)
	conf := *config

	resp, err := conf.Transaction(func(transactionID string) (*http.Response, error) {
		return conf.UpdateServerConfiguration(serverName, payload, transactionID, parentName, parentType)
	})

	if err != nil {
		fmt.Println("Error creating Server configuration:", err)
		return err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return fmt.Errorf("error creating Server configuration: %s", resp.Status)
	}

	d.SetId(serverName)
	return nil
}

func resourceHaproxyServerDelete(d *schema.ResourceData, m interface{}) error {
	serverName := d.Get("name").(string)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	config := m.(**Config)
	conf := *config

	resp, err := conf.Transaction(func(transactionID string) (*http.Response, error) {
		return conf.DeleteServerConfiguration(serverName, transactionID, parentName, parentType)
	})

	if err != nil {
		fmt.Println("Error updating Server configuration:", err)
		return err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return fmt.Errorf("error creating Server configuration: %s", resp.Status)
	}

	d.SetId("")
	return nil
}

// GetAServerConfiguration returns the configuration of a Server.
func (c *Config) GetAServerConfiguration(ServerName string, TransactionID string, parentName string, parentType string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/servers/%s?transaction_id=%s&parent_name=%s&parent_type=%s", c.BaseURL, ServerName, TransactionID, parentName, parentType)
	fmt.Println("****************response response 500 ****************", url)
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

// AddServerConfiguration adds a Server configuration.
func (c *Config) AddServerConfiguration(payload []byte, TransactionID string, parentName string, parentType string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/servers?transaction_id=%s&parent_name=%s&parent_type=%s", c.BaseURL, TransactionID, parentName, parentType)
	fmt.Println("****************response response 500 ****************", url)
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

// UpdateServerConfiguration updates a Server configuration.
func (c *Config) UpdateServerConfiguration(ServerName string, payload []byte, TransactionID string, parentName string, parentType string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/servers/%s?transaction_id=%s&parent_name=%s&parent_type=%s", c.BaseURL, ServerName, TransactionID, parentName, parentType)
	fmt.Println("****************response response 500 ****************", url)
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

// DeleteServerConfiguration deletes a Server configuration.
func (c *Config) DeleteServerConfiguration(ServerName string, TransactionID string, parentName string, parentType string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/servers/%s?transaction_id=%s&parent_name=%s&parent_type=%s", c.BaseURL, ServerName, TransactionID, parentName, parentType)
	fmt.Println("****************response response 500 ****************", url)
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
