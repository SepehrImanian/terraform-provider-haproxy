package transaction

import (
	"fmt"
	"io"
	"terraform-provider-haproxy/internal/utils"
)

// persistTransactionID to persist a transaction ID from memory into haproxy config file
func (c *ConfigTransaction) commitTransactionID(TransactionID string) error {
	url := fmt.Sprintf("%s/v2/services/haproxy/transactions/%s", c.BaseURL, TransactionID)

	fmt.Println("----------url----------", url)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := utils.HTTPRequest("PUT", url, nil, headers, c.Username, c.Password)
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
