package httpcheck

// Config defines variable for haproxy configuration
type ConfigHttpCheck struct {
	Username string
	Password string
	BaseURL  string
	SSL      bool
}

type HttpCheckPayload struct {
	Index           int    `json:"index"`
	Addr            string `json:"addr"`
	Alpn            string `json:"alpn"`
	Body            string `json:"body"`
	BodyLogFormat   string `json:"body_log_format"`
	CheckComment    string `json:"check_comment"`
	Default         bool   `json:"default"`
	ErrorStatus     string `json:"error_status"`
	ExclamationMark bool   `json:"exclamation_mark"`
	Headers         []struct {
		Name string `json:"name"`
		Fmt  string `json:"fmt"`
	} `json:"headers"`
	Linger       bool   `json:"linger"`
	Match        string `json:"match"`
	Method       string `json:"method"`
	MinRecv      int    `json:"min_recv"`
	OkStatus     string `json:"ok_status"`
	OnError      string `json:"on_error"`
	OnSuccess    string `json:"on_success"`
	Pattern      string `json:"pattern"`
	Port         int    `json:"port"`
	PortString   string `json:"port_string"`
	Proto        string `json:"proto"`
	SendProxy    bool   `json:"send_proxy"`
	Sni          string `json:"sni"`
	Ssl          bool   `json:"ssl"`
	StatusCode   string `json:"status-code"`
	Type         string `json:"type"`
	Uri          string `json:"uri"`
	UriLogFormat string `json:"uri_log_format"`
	VarExpr      string `json:"var_expr"`
	VarName      string `json:"var_name"`
	VarFormat    string `json:"var_format"`
	VarScope     string `json:"var_scope"`
	Version      string `json:"version"`
	ViaSocks4    bool   `json:"via_socks4"`
}
