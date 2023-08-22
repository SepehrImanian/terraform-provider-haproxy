package global

// Config defines variable for haproxy configuration
type ConfigGlobal struct {
	Username string
	Password string
	BaseURL  string
	SSL      bool
}

// GlobalPayload defines the payload for global configuration
type GlobalPayload struct {
	User                  string `json:"user"`
	Group                 string `json:"group"`
	Chroot                string `json:"chroot"`
	CpuMaps               CpuMaps
	Daemon                string `json:"daemon"` //bool
	MasterWorker          bool   `json:"master-worker"`
	MaxCompCpuUsage       int    `json:"maxcompcpuusage"`
	MaxPipes              int    `json:"maxpipes"`
	MaxSslConn            int    `json:"maxsslconn"`
	MaxConn               int    `json:"maxconn"`
	NbProc                int    `json:"nbproc"`
	NbThread              int    `json:"nbthread"`
	PidFile               string `json:"pidfile"`
	UlimitN               int    `json:"ulimit_n"`
	CrtBase               string `json:"crt_base"`
	CaBase                string `json:"ca_base"`
	StatsMaxConn          int    `json:"stats_maxconn"`
	StatsTimeOut          int    `json:"stats_timeout"`
	SslDefaultBindCiphers string `json:"ssl_default_bind_ciphers"`
	SslDefaultBindOptions string `json:"ssl_default_bind_options"`
}

type CpuMaps struct {
	Process string `json:"process"`
	CpuSet  string `json:"cpu_set"`
}
