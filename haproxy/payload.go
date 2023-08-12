package haproxy

// Config defines variable for haproxy configuration
type Config struct {
	Username string
	Password string
	BaseURL  string
	SSL      bool
}
