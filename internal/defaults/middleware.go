package defaults

import (
	"fmt"
	"io"
	"net/http"
	"terraform-provider-haproxy/internal/utils"
)

// GetADefaultsConfiguration returns the configuration of a Defaults.
func (c *ConfigDefaults) GetADefaultsConfiguration(DefaultsName string, TransactionID string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/named_defaults/%s?transaction_id=%s", c.BaseURL, DefaultsName, TransactionID)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := utils.HTTPRequest("GET", url, nil, headers, c.Username, c.Password)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()
	return resp, nil
}

// AddDefaultsConfiguration adds a Defaults configuration.
func (c *ConfigDefaults) AddDefaultsConfiguration(payload []byte, TransactionID string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/named_defaults?transaction_id=%s", c.BaseURL, TransactionID)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := utils.HTTPRequest("POST", url, payload, headers, c.Username, c.Password)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	fmt.Println("**************** Response Defaults ****************", string(body))
	return resp, nil
}

// DeleteDefaultsConfiguration deletes a Defaults configuration.
func (c *ConfigDefaults) DeleteDefaultsConfiguration(DefaultsName string, TransactionID string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/named_defaults/%s?transaction_id=%s", c.BaseURL, DefaultsName, TransactionID)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := utils.HTTPRequest("DELETE", url, nil, headers, c.Username, c.Password)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()
	return resp, nil
}

// UpdateDefaultsConfiguration updates a Defaults configuration.
func (c *ConfigDefaults) UpdateDefaultsConfiguration(DefaultsName string, payload []byte, TransactionID string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/named_defaults/%s?transaction_id=%s", c.BaseURL, DefaultsName, TransactionID)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := utils.HTTPRequest("PUT", url, payload, headers, c.Username, c.Password)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()
	return resp, nil
}
