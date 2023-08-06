// main.go
package main

import (
	"terraform-provider-haproxy/haproxy"

	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return haproxy.Provider()
		},
	})
}

//func main() {
//	baseURL := "http://103.75.196.148:5555"
//	username := "admin"
//	password := "adminpwd"
//	resp, err := haproxy.GetBackendConfiguration(baseURL, username, password)
//	if err != nil {
//		fmt.Println("Error getting backend configuration:", err)
//		return
//	}
//	fmt.Println(resp.Body)
//}
