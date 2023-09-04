package user

import (
	"fmt"
	"io"
	"net/http"
	"terraform-provider-haproxy/internal/utils"
)

// GetAUsersConfiguration returns the configuration of a Users.
func (c *ConfigUser) GetAUsersConfiguration(UsersName string, TransactionID string, UserList string) (*http.Response, []byte, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/users/%s?transaction_id=%s&userlist=%s", c.BaseURL, UsersName, TransactionID, UserList)
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

// AddUsersConfiguration adds a Users configuration.
func (c *ConfigUser) AddUsersConfiguration(payload []byte, TransactionID string, UserList string) (*http.Response, []byte, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/users?transaction_id=%s&userlist=%s", c.BaseURL, TransactionID, UserList)
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

// DeleteUsersConfiguration deletes a Users configuration.
func (c *ConfigUser) DeleteUsersConfiguration(UsersName string, TransactionID string, UserList string) (*http.Response, []byte, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/users/%s?transaction_id=%s&userlist=%s", c.BaseURL, UsersName, TransactionID, UserList)
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

// UpdateUsersConfiguration updates a Users configuration.
func (c *ConfigUser) UpdateUsersConfiguration(UsersName string, payload []byte, TransactionID string, UserList string) (*http.Response, []byte, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/users/%s?transaction_id=%s&userlist=%s", c.BaseURL, UsersName, TransactionID, UserList)
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
