package httpcheck

import (
	"fmt"
	"net/http"

	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func DataSourceHaproxyHttpcheck() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHaproxyHttpCheckRead,
		Schema: map[string]*schema.Schema{
			"parent_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the parent object",
			},
			"parent_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of the parent object",
			},
			"index": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The index of the HttpCheck in the parent object starting at 0",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of the check",
			},
		},
	}
}

func dataSourceHaproxyHttpCheckRead(d *schema.ResourceData, m interface{}) error {
	indexName := d.Get("index").(int)
	indexNameStr := fmt.Sprintf("%d", indexName)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	configMap := m.(map[string]interface{})
	HttpCheckConfig := configMap["httpcheck"].(*ConfigHttpCheck)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return HttpCheckConfig.GetAHttpCheckConfiguration(indexName, transactionID, parentName, parentType)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(indexNameStr, "error reading HttpCheck configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(indexNameStr)
	return nil
}
