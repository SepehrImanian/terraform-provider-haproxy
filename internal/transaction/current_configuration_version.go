package transaction

import (
	"fmt"
	"io"
	"terraform-provider-haproxy/internal/utils"
)

// getCurrentConfigurationVersion get current haproxy configuration version
func (c *ConfigTransaction) getCurrentConfigurationVersion() (int, error) {
	url := fmt.Sprintf("%s/v2/services/haproxy/configuration/version", c.BaseURL)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := utils.HTTPRequest("GET", url, nil, headers, c.Username, c.Password)
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
