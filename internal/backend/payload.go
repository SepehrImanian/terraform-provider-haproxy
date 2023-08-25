package backend

// Config defines variable for backend configuration
type ConfigBackend struct {
	Username string
	Password string
	BaseURL  string
	SSL      bool
}

type BackendPayload struct {
	Name                      string `json:"name"`
	Mode                      string `json:"mode"`
	AdvCheck                  string `json:"adv_check"`
	HttpConnectionMode        string `json:"http_connection_mode"`
	ServerTimeout             int    `json:"server_timeout"`
	CheckTimeout              int    `json:"check_timeout"`
	ConnectTimeout            int    `json:"connect_timeout"`
	QueueTimeout              int    `json:"queue_timeout"`
	TunnelTimeout             int    `json:"tunnel_timeout"`
	TarpitTimeout             int    `json:"tarpit_timeout"`
	CheckCache                string `json:"checkcache"` //bool
	Balance                   Balance
	HttpchkParams             HttpchkParams
	Forwardfor                Forwardfor
}

type Balance struct {
	Algorithm string `json:"algorithm"`
}

type HttpchkParams struct {
	Method  string `json:"method"`
	Uri     string `json:"uri"`
	Version string `json:"version"`
}

type Forwardfor struct {
	Enabled string `json:"enabled"` //bool
}
