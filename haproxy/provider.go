package haproxy

import (
	acl "terraform-provider-haproxy/internal/acl"
	backend "terraform-provider-haproxy/internal/backend"
	bind "terraform-provider-haproxy/internal/bind"
	cache "terraform-provider-haproxy/internal/cache"
	defaults "terraform-provider-haproxy/internal/defaults"
	frontend "terraform-provider-haproxy/internal/frontend"
	global "terraform-provider-haproxy/internal/global"
	group "terraform-provider-haproxy/internal/group"
	health "terraform-provider-haproxy/internal/health"
	nameserver "terraform-provider-haproxy/internal/nameserver"
	resolvers "terraform-provider-haproxy/internal/resolvers"
	server "terraform-provider-haproxy/internal/server"
	ServerTemplate "terraform-provider-haproxy/internal/server_template"
	transaction "terraform-provider-haproxy/internal/transaction"
	user "terraform-provider-haproxy/internal/user"
	userlist "terraform-provider-haproxy/internal/userlist"
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
			"haproxy_acl":             acl.DataSourceHaproxyAcl(),
			"haproxy_defaults":        defaults.DataSourceHaproxyDefaults(),
			"haproxy_frontend":        frontend.DataSourceHaproxyFrontend(),
			"haproxy_backend":         backend.DataSourceHaproxyBackend(),
			"haproxy_server":          server.DataSourceHaproxyServer(),
			"haproxy_bind":            bind.DataSourceHaproxyBind(),
			"haproxy_resolvers":       resolvers.DataSourceHaproxyResolvers(),
			"haproxy_cache":           cache.DataSourceHaproxyCache(),
			"haproxy_global":          global.DataSourceHaproxyGlobal(),
			"haproxy_health":          health.DataSourceHaproxyHealth(),
			"haproxy_nameserver":      nameserver.DataSourceHaproxyNameserver(),
			"haproxy_userlist":        userlist.DataSourceHaproxyUserlist(),
			"haproxy_user":            user.DataSourceHaproxyUser(),
			"haproxy_group":           group.DataSourceHaproxyGroup(),
			"haproxy_server_template": ServerTemplate.DataSourceHaproxyServerTemplate(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"haproxy_acl":             acl.ResourceHaproxyAcl(),
			"haproxy_defaults":        defaults.ResourceHaproxyDefaults(),
			"haproxy_frontend":        frontend.ResourceHaproxyFrontend(),
			"haproxy_backend":         backend.ResourceHaproxyBackend(),
			"haproxy_server":          server.ResourceHaproxyServer(),
			"haproxy_bind":            bind.ResourceHaproxyBind(),
			"haproxy_resolvers":       resolvers.ResourceHaproxyResolvers(),
			"haproxy_cache":           cache.ResourceHaproxyCache(),
			"haproxy_global":          global.ResourceHaproxyGlobal(),
			"haproxy_nameserver":      nameserver.ResourceHaproxyNameserver(),
			"haproxy_userlist":        userlist.ResourceHaproxyUserlist(),
			"haproxy_user":            user.ResourceHaproxyUser(),
			"haproxy_group":           group.ResourceHaproxyGroup(),
			"haproxy_server_template": ServerTemplate.ResourceHaproxyServerTemplate(),
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

	// Create transaction config for cache
	cacheConfig := &cache.ConfigCache{}
	utils.SetConfigValues(cacheConfig, commonConfig)

	// Create transaction config for global
	globalConfig := &global.ConfigGlobal{}
	utils.SetConfigValues(globalConfig, commonConfig)

	// Create config for health
	healthConfig := &health.ConfigHealth{}
	utils.SetConfigValues(healthConfig, commonConfig)

	// Create config for nameserver
	nameserverConfig := &nameserver.ConfigNameserver{}
	utils.SetConfigValues(nameserverConfig, commonConfig)

	// Create config for userlist
	userlistConfig := &userlist.ConfigUserlist{}
	utils.SetConfigValues(userlistConfig, commonConfig)

	// Create config for user
	userConfig := &user.ConfigUser{}
	utils.SetConfigValues(userConfig, commonConfig)

	// Create config for group
	groupConfig := &group.ConfigGroup{}
	utils.SetConfigValues(groupConfig, commonConfig)

	// Create config for server template
	serverTemplateConfig := &ServerTemplate.ConfigServerTemplate{}
	utils.SetConfigValues(serverTemplateConfig, commonConfig)

	return map[string]interface{}{
		"backend":        backendConfig,
		"frontend":       frontendConfig,
		"server":         serverConfig,
		"transaction":    transactionConfig,
		"bind":           bindConfig,
		"defaults":       defaultsConfig,
		"acl":            aclConfig,
		"resolvers":      resolversConfig,
		"cache":          cacheConfig,
		"global":         globalConfig,
		"health":         healthConfig,
		"nameserver":     nameserverConfig,
		"userlist":       userlistConfig,
		"user":           userConfig,
		"group":          groupConfig,
		"ServerTemplate": serverTemplateConfig,
	}, nil
}
