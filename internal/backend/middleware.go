package backend

import (
	"fmt"
	"io"
	"net/http"
	"terraform-provider-haproxy/internal/utils"
)

// GetABackendConfiguration returns the configuration of a backend.
func (c *ConfigBackend) GetABackendConfiguration(backendName string, TransactionID string) (*http.Response, error) {

	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/backends/%s?transaction_id=%s", c.BaseURL, backendName, TransactionID)
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

// AddBackendConfiguration adds a backend configuration.
func (c *ConfigBackend) AddBackendConfiguration(payload []byte, TransactionID string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/backends?transaction_id=%s", c.BaseURL, TransactionID)
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	fmt.Println("**************** url 400 *******************", url)
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

	fmt.Println("****************response response 400 ****************", string(body))
	//fmt.Println("^^^^^^^AddBackendConfiguration TransactionID^^^^^^^", TransactionID)

	return resp, nil
}

// DeleteBackendConfiguration deletes a backend configuration.
func (c *ConfigBackend) DeleteBackendConfiguration(backendName string, TransactionID string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/backends/%s?transaction_id=%s", c.BaseURL, backendName, TransactionID)
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

// UpdateBackendConfiguration updates a backend configuration.
func (c *ConfigBackend) UpdateBackendConfiguration(backendName string, payload []byte, TransactionID string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/backends/%s?transaction_id=%s", c.BaseURL, backendName, TransactionID)
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
