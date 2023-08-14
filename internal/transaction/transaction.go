package transaction

import (
	"net/http"
)

// Transaction encapsulate function to ensure that it's executed within a locked context
func (c *ConfigTransaction) Transaction(fn func(transactionID string) (*http.Response, error)) (*http.Response, error) {
	configMutex.Lock()
	defer configMutex.Unlock()

	version, err := c.getCurrentConfigurationVersion()
	if err != nil {
		return nil, err
	}

	id, err := c.createTransactionID(version)
	if err != nil {
		return nil, err
	}

	// Call the provided function (fn) within the locked context and pass transactionID
	resp, err := fn(id)
	if err != nil {
		return nil, err
	}

	err = c.commitTransactionID(id)
	if err != nil {
		return nil, err
	}

	return resp, err
}
