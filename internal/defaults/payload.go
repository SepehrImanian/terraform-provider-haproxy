package defaults

// Config defines variable for haproxy configuration
type ConfigDefaults struct {
	Username string
	Password string
	BaseURL  string
	SSL      bool
}

type DefaultsPayload struct {
	Name                 string `json:"name"`
	Mode                 string `json:"mode"`
	Backlog              int    `json:"backlog"`
	HTTPLog              bool   `json:"httplog"`
	HTTPSLog             string `json:"httpslog"`
	TCPLog               bool   `json:"tcplog"`
	Retries              int    `json:"retries"`
	CheckTimeout         int    `json:"check_timeout"`
	ClientTimeout        int    `json:"client_timeout"`
	ConnectTimeout       int    `json:"connect_timeout"`
	HTTPKeepAliveTimeout int    `json:"http_keep_alive_timeout"`
	HTTPRequestTimeout   int    `json:"http_request_timeout"`
	QueueTimeout         int    `json:"queue_timeout"`
	ServerTimeout        int    `json:"server_timeout"`
	ServerFinTimeout     int    `json:"server_fin_timeout"`
	MaxConn              int    `json:"maxconn"`
}
