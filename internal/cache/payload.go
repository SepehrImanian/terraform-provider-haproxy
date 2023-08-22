package cache

// Config defines variable for haproxy configuration
type ConfigCache struct {
	Username string
	Password string
	BaseURL  string
	SSL      bool
}

type CachePayload struct {
	Name                string `json:"name"`
	MaxAge              int    `json:"max_age"`
	MaxObjectSize       int    `json:"max_object_size"`
	MaxSecondaryEntries int    `json:"max_secondary_entries"`
	ProcessVary         bool   `json:"process_vary"`
	TotalMaxSize        int    `json:"total_max_size"`
}
