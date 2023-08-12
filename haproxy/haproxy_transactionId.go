package haproxy

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

type TransactionResponse struct {
	Version int    `json:"_version"`
	ID      string `json:"id"`
	Status  string `json:"status"`
}

var configMutex sync.Mutex

// getCurrentConfigurationVersion get current haproxy configuration version
func (c *Config) getCurrentConfigurationVersion() (int, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/version", c.BaseURL)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := HTTPRequest("GET", url, nil, headers, c.Username, c.Password)
	if err != nil {
		return 0, fmt.Errorf("error sending request: %v", err)
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("error reading response body: %v", err)
	}

	var version int
	_, err = fmt.Sscanf(string(bodyBytes), "%d", &version)
	if err != nil {
		return 0, fmt.Errorf("error parsing version: %v", err)
	}

	fmt.Println("--------------getCurrentConfigurationVersion---------", string(bodyBytes))
	return version, nil
}

// createTransactionID for create a transaction ID
func (c *Config) createTransactionID(version int) (string, error) {
	versionStr := strconv.Itoa(version)

	url := fmt.Sprintf("%s/v2/services/haproxy/transactions?version=%s", c.BaseURL, versionStr)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := HTTPRequest("POST", url, nil, headers, c.Username, c.Password)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	var responseData TransactionResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return "", fmt.Errorf("error decoding response JSON: %v", err)
	}

	fmt.Println("--------------create TransactionID---------", string(body))

	return responseData.ID, nil
}

// persistTransactionID to persist a transaction ID from memory into haproxy config file
func (c *Config) persistTransactionID(TransactionID string) error {
	url := fmt.Sprintf("%s/v2/services/haproxy/transactions/%s", c.BaseURL, TransactionID)

	fmt.Println("----------url----------", url)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := HTTPRequest("PUT", url, nil, headers, c.Username, c.Password)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	fmt.Println(string(body))
	return nil
}

// Transaction  encapsulate function to ensure that it's executed within a locked context
func (c *Config) Transaction(fn func(transactionID string) (*http.Response, error)) (*http.Response, error) {
	configMutex.Lock()
	defer configMutex.Unlock()

	version, err := c.getCurrentConfigurationVersion()
	if err != nil {
		return nil, err
	}

	id, err := c.createTransactionID(version)
	if err != nil {
		return nil, err
	}

	// Call the provided function (fn) within the locked context and pass transactionID
	resp, err := fn(id)
	if err != nil {
		return nil, err
	}

	err = c.persistTransactionID(id)
	if err != nil {
		return nil, err
	}

	return resp, err
}
