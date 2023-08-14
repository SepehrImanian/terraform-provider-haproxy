package frontend

// Config defines variable for haproxy configuration
type ConfigFrontend struct {
	Username string
	Password string
	BaseURL  string
	SSL      bool
}
