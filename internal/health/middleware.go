package health

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"terraform-provider-haproxy/internal/utils"
)

// GetAHealth returns haproxy Health.
func (c *ConfigHealth) GetAHealth() (health bool, resp *http.Response, err error) {
	url := fmt.Sprintf("%s/v2/health", c.BaseURL)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err = utils.HTTPRequest("GET", url, nil, headers, c.Username, c.Password)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return false, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, nil, fmt.Errorf("error reading response body: %v", err)
	}

	var healthData struct {
		Haproxy string `json:"haproxy"`
	}

	err = json.Unmarshal(body, &healthData)
	if err != nil {
		return false, nil, fmt.Errorf("error decoding response JSON: %v", err)
	}

	return healthData.Haproxy == "up", resp, nil
}
