// provider.go
package haproxy

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
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
