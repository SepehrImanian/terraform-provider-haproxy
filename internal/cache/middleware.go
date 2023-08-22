package cache

import (
	"fmt"
	"io"
	"net/http"
	"terraform-provider-haproxy/internal/utils"
)

// GetACacheConfiguration returns the configuration of a Cache.
func (c *ConfigCache) GetACacheConfiguration(CacheName string, TransactionID string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/caches/%s?transaction_id=%s", c.BaseURL, CacheName, TransactionID)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := utils.HTTPRequest("GET", url, nil, headers, c.Username, c.Password)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	fmt.Println("**************** Response Cache ****************", string(body))

	return resp, nil
}

// AddCacheConfiguration adds a Cache configuration.
func (c *ConfigCache) AddCacheConfiguration(payload []byte, TransactionID string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/caches?transaction_id=%s", c.BaseURL, TransactionID)
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

	fmt.Println("**************** Response Cache ****************", string(body))
	return resp, nil
}

// DeleteCacheConfiguration deletes a Cache configuration.
func (c *ConfigCache) DeleteCacheConfiguration(CacheName string, TransactionID string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/caches/%s?transaction_id=%s", c.BaseURL, CacheName, TransactionID)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := utils.HTTPRequest("DELETE", url, nil, headers, c.Username, c.Password)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	fmt.Println("**************** Response Cache ****************", string(body))
	return resp, nil
}

// UpdateCacheConfiguration updates a Cache configuration.
func (c *ConfigCache) UpdateCacheConfiguration(CacheName string, payload []byte, TransactionID string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/caches/%s?transaction_id=%s", c.BaseURL, CacheName, TransactionID)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := utils.HTTPRequest("PUT", url, payload, headers, c.Username, c.Password)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	fmt.Println("**************** Response Cache ****************", string(body))
	return resp, nil
}
