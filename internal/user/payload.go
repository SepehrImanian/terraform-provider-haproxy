package user

// Config defines variable for haproxy configuration
type ConfigUser struct {
	Username string
	Password string
	BaseURL  string
	SSL      bool
}

type UserPayload struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	SecurePassword bool   `json:"secure_password"`
	Groups         string `json:"groups"`
}
