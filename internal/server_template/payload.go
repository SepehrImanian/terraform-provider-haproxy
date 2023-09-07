package ServerTemplate

// Config defines variable for haproxy configuration
type ConfigServerTemplate struct {
	Username string
	Password string
	BaseURL  string
	SSL      bool
}

type ServerTemplatePayload struct {
	Prefix     string `json:"prefix"`
	Fqdn       string `json:"fqdn"`
	NumOrRange string `json:"num_or_range"`
	Port       int    `json:"port"`
}
