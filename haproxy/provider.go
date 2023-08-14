package haproxy

import (
	backend "terraform-provider-haproxy/internal/backend"
	frontend "terraform-provider-haproxy/internal/frontend"
	server "terraform-provider-haproxy/internal/server"
	transaction "terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

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

		DataSourcesMap: map[string]*schema.Resource{
			// "haproxy_acl":      dataSourceHaproxyAcl(),
			"haproxy_frontend": frontend.DataSourceHaproxyFrontend(),
			"haproxy_backend":  backend.DataSourceHaproxyBackend(),
		},

		ResourcesMap: map[string]*schema.Resource{
			//"haproxy_global":    resourceHaproxyGlobal(),
			//"haproxy_defaults":  resourceHaproxyDefaults(),
			//"haproxy_dashboard": resourceHaproxyDashboard(),
			//"haproxy_acl":       resourceHaproxyAcl(),
			"haproxy_frontend": frontend.ResourceHaproxyFrontend(),
			"haproxy_backend":  backend.ResourceHaproxyBackend(),
			"haproxy_server":   server.ResourceHaproxyServer(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(data *schema.ResourceData) (interface{}, error) {
	commonConfig := utils.Configuration{
		Username: data.Get("username").(string),
		Password: data.Get("password").(string),
		BaseURL:  data.Get("url").(string),
	}

	// Create backend config
	backendConfig := &backend.ConfigBackend{}
	utils.SetConfigValues(backendConfig, commonConfig)

	// Create frontend config
	frontendConfig := &frontend.ConfigFrontend{}
	utils.SetConfigValues(frontendConfig, commonConfig)

	// Create server config
	serverConfig := &server.ConfigServer{}
	utils.SetConfigValues(serverConfig, commonConfig)

	// Create transaction config
	transactionConfig := &transaction.ConfigTransaction{}
	utils.SetConfigValues(transactionConfig, commonConfig)

	return map[string]interface{}{
		"backend":     backendConfig,
		"frontend":    frontendConfig,
		"server":      serverConfig,
		"transaction": transactionConfig,
	}, nil
}
