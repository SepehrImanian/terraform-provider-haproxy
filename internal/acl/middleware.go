package acl

import (
	"fmt"
	"io"
	"net/http"
	"terraform-provider-haproxy/internal/utils"
)

// GetAAclConfiguration returns the configuration of a Acl.
func (c *ConfigAcl) GetAAclConfiguration(AclIndexName int, TransactionID string, parentName string, parentType string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/acls/%d?transaction_id=%s&parent_name=%s&parent_type=%s", c.BaseURL, AclIndexName, TransactionID, parentName, parentType)
	fmt.Println("**************** response ****************", url)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := utils.HTTPRequest("GET", url, nil, headers, c.Username, c.Password)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	fmt.Println("**************** response response ****************", resp)
	defer resp.Body.Close()
	return resp, nil
}

// AddAclConfiguration adds a Acl configuration.
func (c *ConfigAcl) AddAclConfiguration(payload []byte, TransactionID string, parentName string, parentType string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/acls?transaction_id=%s&parent_name=%s&parent_type=%s", c.BaseURL, TransactionID, parentName, parentType)
	fmt.Println("**************** response ****************", url)
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

	fmt.Println("**************** response response ****************", string(body))
	return resp, nil
}

// UpdateAclConfiguration updates a Acl configuration.
func (c *ConfigAcl) UpdateAclConfiguration(AclIndexName int, payload []byte, TransactionID string, parentName string, parentType string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/acls/%d?transaction_id=%s&parent_name=%s&parent_type=%s", c.BaseURL, AclIndexName, TransactionID, parentName, parentType)
	fmt.Println("**************** response response  ****************", url)
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

// DeleteAclConfiguration deletes a Acl configuration.
func (c *ConfigAcl) DeleteAclConfiguration(AclIndexName int, TransactionID string, parentName string, parentType string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/acls/%d?transaction_id=%s&parent_name=%s&parent_type=%s", c.BaseURL, AclIndexName, TransactionID, parentName, parentType)
	fmt.Println("**************** response response ****************", url)
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
