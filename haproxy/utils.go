package haproxy

import (
	"bytes"
	"encoding/base64"
	"net/http"
)

func HTTPRequest(method, url string, body []byte, headers map[string]string, Username string, Password string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	if Username != "" && Password != "" {
		req.SetBasicAuth(Username, Password)
	}
	return http.DefaultClient.Do(req)
}

func encodeCredentials(username, password string) string {
	credentials := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(credentials))
}
