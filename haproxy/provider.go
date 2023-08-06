package haproxy

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"haproxy_server": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Haproxy Host and Port",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HAPROXY_ENDPOINT",
				}, nil),
			},
			"haproxy_user": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Haproxy User",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HAPROXY_USER",
				}, nil),
				ConflictsWith: []string{"minio_access_key"},
			},
			"haproxy_password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Haproxy Password",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HAPROXY_PASSWORD",
				}, nil),
				ConflictsWith: []string{"minio_secret_key"},
			},
			"haproxy_insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Disable SSL certificate verification (default: false)",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HAPROXY_INSECURE",
				}, nil),
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"haproxy_acl":      dataSourceHaproxyAcl(),
			"haproxy_frontend": dataSourceHaproxyFrontend(),
			"haproxy_backend":  dataSourceHaproxyBackend(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"haproxy_global":    resourceHaproxyGlobal(),
			"haproxy_defaults":  resourceHaproxyDefaults(),
			"haproxy_dashboard": resourceHaproxyDashboard(),
			"haproxy_acl":       resourceHaproxyAcl(),
			"haproxy_frontend":  resourceHaproxyFrontend(),
			"haproxy_backend":   resourceHaproxyBackend(),
		},
	}
}
