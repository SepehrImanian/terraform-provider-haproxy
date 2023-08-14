package server

import (
	"fmt"
	"io"
	"net/http"
	"terraform-provider-haproxy/internal/utils"
)

// GetAServerConfiguration returns the configuration of a Server.
func (c *ConfigServer) GetAServerConfiguration(ServerName string, TransactionID string, parentName string, parentType string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/servers/%s?transaction_id=%s&parent_name=%s&parent_type=%s", c.BaseURL, ServerName, TransactionID, parentName, parentType)
	fmt.Println("****************response response 500 ****************", url)
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

// AddServerConfiguration adds a Server configuration.
func (c *ConfigServer) AddServerConfiguration(payload []byte, TransactionID string, parentName string, parentType string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/servers?transaction_id=%s&parent_name=%s&parent_type=%s", c.BaseURL, TransactionID, parentName, parentType)
	fmt.Println("****************response response 500 ****************", url)
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

	fmt.Println("****************response response 400 ****************", string(body))
	return resp, nil
}

// UpdateServerConfiguration updates a Server configuration.
func (c *ConfigServer) UpdateServerConfiguration(ServerName string, payload []byte, TransactionID string, parentName string, parentType string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/servers/%s?transaction_id=%s&parent_name=%s&parent_type=%s", c.BaseURL, ServerName, TransactionID, parentName, parentType)
	fmt.Println("****************response response 500 ****************", url)
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

// DeleteServerConfiguration deletes a Server configuration.
func (c *ConfigServer) DeleteServerConfiguration(ServerName string, TransactionID string, parentName string, parentType string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/servers/%s?transaction_id=%s&parent_name=%s&parent_type=%s", c.BaseURL, ServerName, TransactionID, parentName, parentType)
	fmt.Println("****************response response 500 ****************", url)
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
