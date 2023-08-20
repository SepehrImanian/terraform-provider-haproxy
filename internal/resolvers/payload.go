package resolvers

import "github.com/stretchr/testify/mock"

// Config defines variable for haproxy configuration
type ConfigResolvers struct {
	Username string
	Password string
	BaseURL  string
	SSL      bool
}

type ResolversPayload struct {
	Name                string `json:"name"`
	AcceptedPayloadSize int    `json:"accepted_payload_size"`
	HoldNx              int    `json:"hold_nx"`
	HoldOther           int    `json:"hold_other"`
	HoldRefused         int    `json:"hold_refused"`
	HoldTimeout         int    `json:"hold_timeout"`
	HoldValid           int    `json:"hold_valid"`
	ParseResolvConf     bool   `json:"parse-resolv-conf"`
	ResolveRetries      int    `json:"resolve_retries"`
	TimeoutResolve      int    `json:"timeout_resolve"`
	TimeoutRetry        int    `json:"timeout_retry"`
}

// Mock implementation for ConfigResolvers and ConfigTransaction for unit testing

type MockConfigResolvers struct {
	mock.Mock
}

type MockConfigTransaction struct {
	mock.Mock
}
