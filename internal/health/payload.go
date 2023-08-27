package health

// Config defines variable for haproxy configuration
type ConfigHealth struct {
	Username string
	Password string
	BaseURL  string
	SSL      bool
}
