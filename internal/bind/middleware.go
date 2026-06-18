package bind

import (
	"fmt"
	"io"
	"net/http"
	"terraform-provider-haproxy/internal/utils"
)

// bindURL builds the binds endpoint URL. Data Plane API v3 nests binds
// under their parent (e.g. /frontends/{parent_name}/binds) instead of the
// flat /binds?parent_name=&parent_type= query-param style used in v2.
func (c *ConfigBind) bindURL(bindName string, transactionID string, parentName string, parentType string) string {
	if c.APIVersion == "v3" {
		base := fmt.Sprintf("%s/%s/services/haproxy/configuration/%ss/%s/binds", c.BaseURL, c.APIVersion, parentType, parentName)
		if bindName != "" {
			base = fmt.Sprintf("%s/%s", base, bindName)
		}
		return fmt.Sprintf("%s?transaction_id=%s", base, transactionID)
	}
	base := fmt.Sprintf("%s/%s/services/haproxy/configuration/binds", c.BaseURL, c.APIVersion)
	if bindName != "" {
		base = fmt.Sprintf("%s/%s", base, bindName)
	}
	return fmt.Sprintf("%s?transaction_id=%s&parent_name=%s&parent_type=%s", base, transactionID, parentName, parentType)
}

// GetABindConfiguration returns the configuration of a Bind.
func (c *ConfigBind) GetABindConfiguration(BindName string, TransactionID string, parentName string, parentType string) (*http.Response, error) {
	url := c.bindURL(BindName, TransactionID, parentName, parentType)
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

// AddBindConfiguration adds a Bind configuration.
func (c *ConfigBind) AddBindConfiguration(payload []byte, TransactionID string, parentName string, parentType string) (*http.Response, error) {
	url := c.bindURL("", TransactionID, parentName, parentType)
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

// UpdateBindConfiguration updates a Bind configuration.
func (c *ConfigBind) UpdateBindConfiguration(BindName string, payload []byte, TransactionID string, parentName string, parentType string) (*http.Response, error) {
	url := c.bindURL(BindName, TransactionID, parentName, parentType)
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

// DeleteBindConfiguration deletes a Bind configuration.
func (c *ConfigBind) DeleteBindConfiguration(BindName string, TransactionID string, parentName string, parentType string) (*http.Response, error) {
	url := c.bindURL(BindName, TransactionID, parentName, parentType)
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
