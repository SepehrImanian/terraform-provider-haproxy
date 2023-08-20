package haproxy

import (
	acl "terraform-provider-haproxy/internal/acl"
	backend "terraform-provider-haproxy/internal/backend"
	bind "terraform-provider-haproxy/internal/bind"
	defaults "terraform-provider-haproxy/internal/defaults"
	frontend "terraform-provider-haproxy/internal/frontend"
	resolvers "terraform-provider-haproxy/internal/resolvers"
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
			"haproxy_acl":       acl.DataSourceHaproxyAcl(),
			"haproxy_defaults":  defaults.DataSourceHaproxyDefaults(),
			"haproxy_frontend":  frontend.DataSourceHaproxyFrontend(),
			"haproxy_backend":   backend.DataSourceHaproxyBackend(),
			"haproxy_server":    server.DataSourceHaproxyServer(),
			"haproxy_bind":      bind.DataSourceHaproxyBind(),
			"haproxy_resolvers": resolvers.DataSourceHaproxyResolvers(),
		},

		ResourcesMap: map[string]*schema.Resource{
			//"haproxy_global":    resourceHaproxyGlobal(),
			"haproxy_acl":       acl.ResourceHaproxyAcl(),
			"haproxy_defaults":  defaults.ResourceHaproxyDefaults(),
			"haproxy_frontend":  frontend.ResourceHaproxyFrontend(),
			"haproxy_backend":   backend.ResourceHaproxyBackend(),
			"haproxy_server":    server.ResourceHaproxyServer(),
			"haproxy_bind":      bind.ResourceHaproxyBind(),
			"haproxy_resolvers": resolvers.ResourceHaproxyResolvers(),
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

	// Create backend config for backend
	backendConfig := &backend.ConfigBackend{}
	utils.SetConfigValues(backendConfig, commonConfig)

	// Create frontend config for frontend
	frontendConfig := &frontend.ConfigFrontend{}
	utils.SetConfigValues(frontendConfig, commonConfig)

	// Create server config for server
	serverConfig := &server.ConfigServer{}
	utils.SetConfigValues(serverConfig, commonConfig)

	// Create transaction config for transaction
	transactionConfig := &transaction.ConfigTransaction{}
	utils.SetConfigValues(transactionConfig, commonConfig)

	// Create transaction config for bind
	bindConfig := &bind.ConfigBind{}
	utils.SetConfigValues(bindConfig, commonConfig)

	// Create transaction config for defaults
	defaultsConfig := &defaults.ConfigDefaults{}
	utils.SetConfigValues(defaultsConfig, commonConfig)

	// Create transaction config for acl
	aclConfig := &acl.ConfigAcl{}
	utils.SetConfigValues(aclConfig, commonConfig)

	// Create transaction config for resolvers
	resolversConfig := &resolvers.ConfigResolvers{}
	utils.SetConfigValues(resolversConfig, commonConfig)

	return map[string]interface{}{
		"backend":     backendConfig,
		"frontend":    frontendConfig,
		"server":      serverConfig,
		"transaction": transactionConfig,
		"bind":        bindConfig,
		"defaults":    defaultsConfig,
		"acl":         aclConfig,
		"resolvers":   resolversConfig,
	}, nil
}
