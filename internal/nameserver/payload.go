package nameserver

// Config defines variable for haproxy configuration
type ConfigNameserver struct {
	Username string
	Password string
	BaseURL  string
	SSL      bool
}

type NameserverPayload struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Port    int    `json:"port"`
}
