package transaction

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"terraform-provider-haproxy/internal/utils"
)

// createTransactionID creates a new transaction ID
func (c *ConfigTransaction) createTransactionID(version int) (string, error) {
	versionStr := strconv.Itoa(version)

	url := fmt.Sprintf("%s/v2/services/haproxy/transactions?version=%s", c.BaseURL, versionStr)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := utils.HTTPRequest("POST", url, nil, headers, c.Username, c.Password)
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
