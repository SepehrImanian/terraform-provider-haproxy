package bind

// Config defines variable for haproxy configuration
type ConfigBind struct {
	Username string
	Password string
	BaseURL  string
	SSL      bool
}
