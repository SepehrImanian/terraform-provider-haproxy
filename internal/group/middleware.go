package group

import (
	"fmt"
	"io"
	"net/http"
	"terraform-provider-haproxy/internal/utils"
)

// GetAGroupsConfiguration returns the configuration of a Groups.
func (c *ConfigGroup) GetAGroupsConfiguration(GroupName string, TransactionID string, UserList string) (*http.Response, []byte, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/groups/%s?transaction_id=%s&userlist=%s", c.BaseURL, GroupName, TransactionID, UserList)
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

// AddGroupsConfiguration adds a Groups configuration.
func (c *ConfigGroup) AddGroupsConfiguration(payload []byte, TransactionID string, UserList string) (*http.Response, []byte, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/groups?transaction_id=%s&userlist=%s", c.BaseURL, TransactionID, UserList)
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

// DeleteGroupsConfiguration deletes a Groups configuration.
func (c *ConfigGroup) DeleteGroupsConfiguration(GroupName string, TransactionID string, UserList string) (*http.Response, []byte, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/groups/%s?transaction_id=%s&userlist=%s", c.BaseURL, GroupName, TransactionID, UserList)
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

// UpdateGroupsConfiguration updates a Groups configuration.
func (c *ConfigGroup) UpdateGroupsConfiguration(GroupName string, payload []byte, TransactionID string, UserList string) (*http.Response, []byte, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/groups/%s?transaction_id=%s&userlist=%s", c.BaseURL, GroupName, TransactionID, UserList)
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
