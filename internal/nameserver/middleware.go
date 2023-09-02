package nameserver

import (
	"fmt"
	"io"
	"net/http"
	"terraform-provider-haproxy/internal/utils"
)

// GetANameserversConfiguration returns the configuration of a Nameservers.
func (c *ConfigNameserver) GetANameserversConfiguration(NameserversName string, TransactionID string, ResolversName string) (*http.Response, []byte, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/nameservers/%s?transaction_id=%s&resolver=%s", c.BaseURL, NameserversName, TransactionID, ResolversName)
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

// AddNameserversConfiguration adds a Nameservers configuration.
func (c *ConfigNameserver) AddNameserversConfiguration(payload []byte, TransactionID string, ResolversName string) (*http.Response, []byte, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/nameservers?transaction_id=%s&resolver=%s", c.BaseURL, TransactionID, ResolversName)
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

// DeleteNameserversConfiguration deletes a Nameservers configuration.
func (c *ConfigNameserver) DeleteNameserversConfiguration(NameserversName string, TransactionID string, ResolversName string) (*http.Response, []byte, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/nameservers/%s?transaction_id=%s&resolver=%s", c.BaseURL, NameserversName, TransactionID, ResolversName)
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

// UpdateNameserversConfiguration updates a Nameservers configuration.
func (c *ConfigNameserver) UpdateNameserversConfiguration(NameserversName string, payload []byte, TransactionID string, ResolversName string) (*http.Response, []byte, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/nameservers/%s?transaction_id=%s&resolver=%s", c.BaseURL, NameserversName, TransactionID, ResolversName)
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
