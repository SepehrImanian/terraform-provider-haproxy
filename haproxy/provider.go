package haproxy

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Haproxy Host and Port",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HAPROXY_ENDPOINT",
				}, nil),
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Haproxy User",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HAPROXY_USER",
				}, nil),
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Haproxy Password",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HAPROXY_PASSWORD",
				}, nil),
			},
			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Disable SSL certificate verification (default: false)",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"HAPROXY_INSECURE",
				}, nil),
			},
		},

		//DataSourcesMap: map[string]*schema.Resource{
		//	"haproxy_acl":      dataSourceHaproxyAcl(),
		//	"haproxy_frontend": dataSourceHaproxyFrontend(),
		//	"haproxy_backend":  dataSourceHaproxyBackend(),
		//},

		ResourcesMap: map[string]*schema.Resource{
			//"haproxy_global":    resourceHaproxyGlobal(),
			//"haproxy_defaults":  resourceHaproxyDefaults(),
			//"haproxy_dashboard": resourceHaproxyDashboard(),
			//"haproxy_acl":       resourceHaproxyAcl(),
			"haproxy_frontend": resourceHaproxyFrontend(),
			"haproxy_backend":  resourceHaproxyBackend(),
			"haproxy_server":   resourceHaproxyServer(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(data *schema.ResourceData) (interface{}, error) {
	var (
		username = data.Get("username").(string)
		password = data.Get("password").(string)
		baseurl  = data.Get("url").(string)
	)

	config := &Config{
		Username: username,
		Password: password,
		BaseURL:  baseurl,
	}

	return &config, nil
}
