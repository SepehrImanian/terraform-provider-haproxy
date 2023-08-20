package bind

import (
	"fmt"
	"io"
	"net/http"
	"terraform-provider-haproxy/internal/utils"
)

// GetABindConfiguration returns the configuration of a Bind.
func (c *ConfigBind) GetABindConfiguration(BindName string, TransactionID string, parentName string, parentType string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/binds/%s?transaction_id=%s&parent_name=%s&parent_type=%s", c.BaseURL, BindName, TransactionID, parentName, parentType)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := utils.HTTPRequest("GET", url, nil, headers, c.Username, c.Password)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	fmt.Println("**************** 404 err bind ****************", string(body))
	defer resp.Body.Close()
	return resp, nil
}

// AddBindConfiguration adds a Bind configuration.
func (c *ConfigBind) AddBindConfiguration(payload []byte, TransactionID string, parentName string, parentType string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/binds?transaction_id=%s&parent_name=%s&parent_type=%s", c.BaseURL, TransactionID, parentName, parentType)
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

	fmt.Println("**************** response body ****************", string(body))
	return resp, nil
}

// UpdateBindConfiguration updates a Bind configuration.
func (c *ConfigBind) UpdateBindConfiguration(BindName string, payload []byte, TransactionID string, parentName string, parentType string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/binds/%s?transaction_id=%s&parent_name=%s&parent_type=%s", c.BaseURL, BindName, TransactionID, parentName, parentType)
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

// DeleteBindConfiguration deletes a Bind configuration.
func (c *ConfigBind) DeleteBindConfiguration(BindName string, TransactionID string, parentName string, parentType string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/binds/%s?transaction_id=%s&parent_name=%s&parent_type=%s", c.BaseURL, BindName, TransactionID, parentName, parentType)
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
