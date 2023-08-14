package backend

import (
	"net/http"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func DataSourceHaproxyBackend() *schema.Resource {
	return nil
}

func GetBackendsConfiguration(baseURL string, username string, password string) (*http.Response, error) {
	URL := baseURL + "/v2/services/haproxy/configuration/backends"

	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := utils.HTTPRequest(http.MethodGet, URL, nil, headers, username, password)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, nil
}
