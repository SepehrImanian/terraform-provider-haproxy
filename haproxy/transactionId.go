package haproxy

import (
	"encoding/json"
	"fmt"
	"io"
)

type TransactionResponse struct {
	Version int    `json:"_version"`
	ID      string `json:"id"`
	Status  string `json:"status"`
}

// CreateTransactionID for create a transaction ID
func CreateTransactionID(BaseURL string, Username string, Password string) (string, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/transactions?version=1", BaseURL)
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
