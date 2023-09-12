package filter

// Config defines variable for haproxy configuration
type ConfigFilter struct {
	Username string
	Password string
	BaseURL  string
	SSL      bool
}

type FilterPayload struct {
	Index              int    `json:"index"`
	AppName            string `json:"app_name"`
	BandwidthLimitName string `json:"bandwidth_limit_name"`
	CacheName          string `json:"cache_name"`
	DefaultLimit       int    `json:"default_limit"`
	DefaultPeriod      int    `json:"default_period"`
	Key                string `json:"key"`
	Limit              int    `json:"limit"`
	MinSize            int    `json:"min_size"`
	SpoeConfig         string `json:"spoe_config"`
	SpoeEngine         string `json:"spoe_engine"`
	Table              string `json:"table"`
	TraceHexdump       bool   `json:"trace_hexdump"`
	TraceName          string `json:"trace_name"`
	TraceRndForwarding bool   `json:"trace_rnd_forwarding"`
	TraceRndParsing    bool   `json:"trace_rnd_parsing"`
	Type               string `json:"type"`
}
