package backend

import "net/http"

// Backend defines interface for backend configuration
type Backend interface {
	GetABackendConfiguration(backendName string, TransactionID string) (*http.Response, error)
	AddBackendConfiguration(payload []byte, TransactionID string) (*http.Response, error)
	DeleteBackendConfiguration(backendName string, TransactionID string) (*http.Response, error)
	UpdateBackendConfiguration(backendName string, payload []byte, TransactionID string) (*http.Response, error)
}

// Config defines variable for backend configuration
type ConfigBackend struct {
	Username string
	Password string
	BaseURL  string
	SSL      bool
}
