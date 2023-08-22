package global

import (
	"fmt"
	"net/http"
	"terraform-provider-haproxy/internal/utils"
)

// GetAGlobalConfiguration returns the configuration of a Global.
func (c *ConfigGlobal) GetAGlobalConfiguration(TransactionID string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/global?transaction_id=%s", c.BaseURL, TransactionID)
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

// UpdateGlobalConfiguration updates a Global configuration.
func (c *ConfigGlobal) UpdateGlobalConfiguration(payload []byte, TransactionID string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/global?transaction_id=%s", c.BaseURL, TransactionID)
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
