package haproxy

import "sync"

// Config defines variable for haproxy configuration
type Config struct {
	Username string
	Password string
	BaseURL  string
	SSL      bool
}

// TransactionResponse get response when Transaction create
type TransactionResponse struct {
	Version int    `json:"_version"`
	ID      string `json:"id"`
	Status  string `json:"status"`
}

var configMutex sync.Mutex
