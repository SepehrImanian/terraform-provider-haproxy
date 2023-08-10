package haproxy

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type TransactionResponse struct {
	Version int    `json:"_version"`
	ID      string `json:"id"`
	Status  string `json:"status"`
}

func getCurrentConfigurationVersion(BaseURL string, Username string, Password string) (int, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/version", BaseURL)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := HTTPRequest("GET", url, nil, headers, Username, Password)
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

	_ = &Config{
		ConfigurationVersion: version,
	}
	return version, nil
}

// createTransactionID for create a transaction ID
func createTransactionID(BaseURL string, Username string, Password string) (string, error) {
	configurationVersion, err := getCurrentConfigurationVersion(BaseURL, Username, Password)

	versionStr := strconv.Itoa(configurationVersion)
	//versionStr := configurationVersion + 1

	//fmt.Println("====================versionStr=================", versionStr)
	url := fmt.Sprintf("%s/v2/services/haproxy/transactions?version=%s", BaseURL, versionStr)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := HTTPRequest("POST", url, nil, headers, Username, Password)
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

	return responseData.ID, nil
}

// persistTransactionID to persist a transaction ID from memory into haproxy config file
func persistTransactionID(transactionID string, BaseURL string, Username string, Password string) (*http.Response, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/transactions/%s", BaseURL, transactionID)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := HTTPRequest("PUT", url, nil, headers, Username, Password)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()
	return resp, nil
}
