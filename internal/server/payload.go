package server

// Config defines variable for haproxy configuration
type ConfigServer struct {
	Username string
	Password string
	BaseURL  string
	SSL      bool
}

type ServerPayload struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Port      int    `json:"port"`
	SendProxy string `json:"send-proxy"`
	Check     string `json:"check"`
	Inter     int    `json:"inter"`
	Rise      int    `json:"rise"`
	Fall      int    `json:"fall"`
}
