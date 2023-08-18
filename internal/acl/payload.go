package acl

// Config defines variable for haproxy configuration
type ConfigAcl struct {
	Username string
	Password string
	BaseURL  string
	SSL      bool
}

type AclPayload struct {
	AclName   string `json:"acl_name"`
	Criterion string `json:"criterion"`
	Index     int    `json:"index"`
	Value     string `json:"value"`
}
