package frontend

// Config defines variable for haproxy configuration
type ConfigFrontend struct {
	Username string
	Password string
	BaseURL  string
	SSL      bool
}

type FrontendPayload struct {
	Name                     string `json:"name"`
	DefaultBackend           string `json:"default_backend"`
	HttpConnectionMode       string `json:"http_connection_mode"`
	AcceptInvalidHttpRequest string `json:"accept_invalid_http_request"` //bool
	MaxConn                  int    `json:"maxconn"`
	Mode                     string `json:"mode"`
	Backlog                  int    `json:"backlog"`
	HttpKeepAliveTimeout     int    `json:"http_keep_alive_timeout"`
	HttpRequestTimeout       int    `json:"http_request_timeout"`
	HttpUseProxyHeader       string `json:"http_use_proxy_header"`
	HttpLog                  bool   `json:"httplog"`
	HttpsLog                 string `json:"httpslog"` //bool
	ErrorLogFormat           string `json:"error_log_format"`
	LogFormat                string `json:"log_format"`
	LogFormatSd              string `json:"log_format_sd"`
	MonitorUri               string `json:"monitor_uri"`
	TcpLog                   bool   `json:"tcplog"`
	Compression              Compression
	Forwardfor               Forwardfor
}

type Compression struct {
	Algorithms []string `json:"algorithms"`
	Offload    bool     `json:"offload"`
	Types      []string `json:"types"`
}

type Forwardfor struct {
	Enabled string `json:"enabled"` //bool
	Except  string `json:"except"`
	Header  string `json:"header"`
	Ifnone  bool   `json:"ifnone"`
}
