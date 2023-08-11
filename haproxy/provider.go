package haproxy

import (
	"fmt"
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
			//"haproxy_frontend":  resourceHaproxyFrontend(),
			"haproxy_backend": resourceHaproxyBackend(),
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

	transactionID, err := createTransactionID(baseurl, username, password)
	if err != nil {
		fmt.Println("Error createTransactionID:", err)
		return nil, err
	}

	resp, err := persistTransactionID(baseurl, username, password, transactionID)
	if err != nil {
		fmt.Println("Error persistTransactionID:", err)
		return "", err
	}

	fmt.Println("-----transactionID-----", transactionID)
	fmt.Println("-------resp------------", resp)

	config := &Config{
		Username:      username,
		Password:      password,
		BaseURL:       baseurl,
		TransactionID: transactionID,
	}
	return &config, nil
}
