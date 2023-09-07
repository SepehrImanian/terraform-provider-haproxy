package ServerTemplate

import (
	"fmt"
	"io"
	"net/http"
	"terraform-provider-haproxy/internal/utils"
)

// GetAServerTemplatesConfiguration returns the configuration of a ServerTemplates.
func (c *ConfigServerTemplate) GetAServerTemplatesConfiguration(ServerTemplatesName string, TransactionID string, BackendName string) (*http.Response, []byte, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/server_templates/%s?transaction_id=%s&backend=%s", c.BaseURL, ServerTemplatesName, TransactionID, BackendName)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := utils.HTTPRequest("GET", url, nil, headers, c.Username, c.Password)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("error reading response body: %v", err)
	}

	return resp, body, nil
}

// AddServerTemplatesConfiguration adds a ServerTemplates configuration.
func (c *ConfigServerTemplate) AddServerTemplatesConfiguration(payload []byte, TransactionID string, BackendName string) (*http.Response, []byte, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/server_templates?transaction_id=%s&backend=%s", c.BaseURL, TransactionID, BackendName)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := utils.HTTPRequest("POST", url, payload, headers, c.Username, c.Password)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("error reading response body: %v", err)
	}

	fmt.Println("**************** Response Resolvers ****************", string(body))
	return resp, body, nil
}

// DeleteServerTemplatesConfiguration deletes a ServerTemplates configuration.
func (c *ConfigServerTemplate) DeleteServerTemplatesConfiguration(ServerTemplatesName string, TransactionID string, BackendName string) (*http.Response, []byte, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/server_templates/%s?transaction_id=%s&backend=%s", c.BaseURL, ServerTemplatesName, TransactionID, BackendName)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := utils.HTTPRequest("DELETE", url, nil, headers, c.Username, c.Password)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("error reading response body: %v", err)
	}

	return resp, body, nil
}

// UpdateServerTemplatesConfiguration updates a ServerTemplates configuration.
func (c *ConfigServerTemplate) UpdateServerTemplatesConfiguration(ServerTemplatesName string, payload []byte, TransactionID string, BackendName string) (*http.Response, []byte, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/server_templates/%s?transaction_id=%s&backend=%s", c.BaseURL, ServerTemplatesName, TransactionID, BackendName)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := utils.HTTPRequest("PUT", url, payload, headers, c.Username, c.Password)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("error reading response body: %v", err)
	}

	return resp, body, nil
}
