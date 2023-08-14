package haproxy

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

// Config defines variable for haproxy configuration
type Config struct {
	Username string
	Password string
	BaseURL  string
	SSL      bool
}

var testAccProviders map[string]func() (*schema.Provider, error)
var testAccProvider *schema.Provider
