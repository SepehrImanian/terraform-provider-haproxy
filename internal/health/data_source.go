package health

import (
	"fmt"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func DataSourceHaproxyHealth() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHaproxyHealthRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the health check",
			},
			"health": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceHaproxyHealthRead(d *schema.ResourceData, m interface{}) error {
	healthName := d.Get("name").(string)

	configMap := m.(map[string]interface{})
	healthConfig := configMap["health"].(*ConfigHealth)

	check, resp, err := healthConfig.GetAHealth()

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError("health", "error reading health configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.Set("health", check)

	d.SetId(healthName)
	return nil
}
