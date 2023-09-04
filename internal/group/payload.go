package group

// Config defines variable for haproxy configuration
type ConfigGroup struct {
	Username string
	Password string
	BaseURL  string
	SSL      bool
}

type GroupPayload struct {
	Name  string `json:"name"`
	Users string `json:"users"`
}
