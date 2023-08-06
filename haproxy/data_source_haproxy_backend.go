package haproxy

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"net/http"
)

func dataSourceHaproxyBackend() *schema.Resource {
	return nil
}

func GetBackendsConfiguration(baseURL string, username string, password string) (*http.Response, error) {
	URL := baseURL + "/v2/services/haproxy/configuration/backends"

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Basic " + encodeCredentials(username, password),
	}
	resp, err := HTTPRequest(http.MethodGet, URL, nil, headers)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, nil
}
