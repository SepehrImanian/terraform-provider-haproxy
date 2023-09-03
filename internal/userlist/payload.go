package userlist

// Config defines variable for haproxy configuration
type ConfigUserlist struct {
	Username string
	Password string
	BaseURL  string
	SSL      bool
}

type UserlistPayload struct {
	Name string `json:"name"`
}
