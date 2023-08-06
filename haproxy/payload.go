package haproxy

// BackendConfig is a struct that contains the configuration for the HAProxy backend.
type BackendConfig struct {
	BaseURL       string
	Username      string
	Password      string
	TransactionID string
}
