package server

// Config defines variable for haproxy configuration
type ConfigServer struct {
	Username string
	Password string
	BaseURL  string
	SSL      bool
}
