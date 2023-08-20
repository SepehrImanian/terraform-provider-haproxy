package resolvers

import (
	"fmt"
	"io"
	"net/http"
	"terraform-provider-haproxy/internal/utils"
)

// GetAResolversConfiguration returns the configuration of a Resolvers.
func (c *ConfigResolvers) GetAResolversConfiguration(ResolversName string, TransactionID string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/resolvers/%s?transaction_id=%s", c.BaseURL, ResolversName, TransactionID)
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

// AddResolversConfiguration adds a Resolvers configuration.
func (c *ConfigResolvers) AddResolversConfiguration(payload []byte, TransactionID string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/resolvers?transaction_id=%s", c.BaseURL, TransactionID)
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

	fmt.Println("**************** Response Resolvers ****************", string(body))
	return resp, nil
}

// DeleteResolversConfiguration deletes a Resolvers configuration.
func (c *ConfigResolvers) DeleteResolversConfiguration(ResolversName string, TransactionID string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/resolvers/%s?transaction_id=%s", c.BaseURL, ResolversName, TransactionID)
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

// UpdateResolversConfiguration updates a Resolvers configuration.
func (c *ConfigResolvers) UpdateResolversConfiguration(ResolversName string, payload []byte, TransactionID string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/resolvers/%s?transaction_id=%s", c.BaseURL, ResolversName, TransactionID)
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
