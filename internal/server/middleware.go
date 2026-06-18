package server

import (
	"fmt"
	"io"
	"net/http"
	"terraform-provider-haproxy/internal/utils"
)

// serverURL builds the servers endpoint URL. Data Plane API v3 nests servers
// under their parent (e.g. /backends/{parent_name}/servers) instead of the
// flat /servers?parent_name=&parent_type= query-param style used in v2.
func (c *ConfigServer) serverURL(serverName string, transactionID string, parentName string, parentType string) string {
	if c.APIVersion == "v3" {
		base := fmt.Sprintf("%s/%s/services/haproxy/configuration/%ss/%s/servers", c.BaseURL, c.APIVersion, parentType, parentName)
		if serverName != "" {
			base = fmt.Sprintf("%s/%s", base, serverName)
		}
		return fmt.Sprintf("%s?transaction_id=%s", base, transactionID)
	}
	base := fmt.Sprintf("%s/%s/services/haproxy/configuration/servers", c.BaseURL, c.APIVersion)
	if serverName != "" {
		base = fmt.Sprintf("%s/%s", base, serverName)
	}
	return fmt.Sprintf("%s?transaction_id=%s&parent_name=%s&parent_type=%s", base, transactionID, parentName, parentType)
}

// GetAServerConfiguration returns the configuration of a Server.
func (c *ConfigServer) GetAServerConfiguration(ServerName string, TransactionID string, parentName string, parentType string) (*http.Response, error) {
	url := c.serverURL(ServerName, TransactionID, parentName, parentType)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := utils.HTTPRequest("GET", url, nil, headers, c.Username, c.Password)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, nil
}

// AddServerConfiguration adds a Server configuration.
func (c *ConfigServer) AddServerConfiguration(payload []byte, TransactionID string, parentName string, parentType string) (*http.Response, error) {
	url := c.serverURL("", TransactionID, parentName, parentType)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := utils.HTTPRequest("POST", url, payload, headers, c.Username, c.Password)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	return resp, nil
}

// UpdateServerConfiguration updates a Server configuration.
func (c *ConfigServer) UpdateServerConfiguration(ServerName string, payload []byte, TransactionID string, parentName string, parentType string) (*http.Response, error) {
	url := c.serverURL(ServerName, TransactionID, parentName, parentType)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := utils.HTTPRequest("PUT", url, payload, headers, c.Username, c.Password)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, nil
}

// DeleteServerConfiguration deletes a Server configuration.
func (c *ConfigServer) DeleteServerConfiguration(ServerName string, TransactionID string, parentName string, parentType string) (*http.Response, error) {
	url := c.serverURL(ServerName, TransactionID, parentName, parentType)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := utils.HTTPRequest("DELETE", url, nil, headers, c.Username, c.Password)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, nil
}
