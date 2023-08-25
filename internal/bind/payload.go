package bind

// Config defines variable for haproxy configuration
type ConfigBind struct {
	Username string
	Password string
	BaseURL  string
	SSL      bool
}

type BindPayload struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Port    int    `json:"port"`
	Maxconn int    `json:"maxconn"`
	User    string `json:"user"`
	Group   string `json:"group"`
	Mode    string `json:"mode"`
}
