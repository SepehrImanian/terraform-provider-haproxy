package filter

import (
	"fmt"
	"net/http"

	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ResourceHaproxyFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceHaproxyFilterCreate,
		Read:   resourceHaproxyFilterRead,
		Update: resourceHaproxyFilterUpdate,
		Delete: resourceHaproxyFilterDelete,

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
				Description: "The index of the Filter in the parent object starting at 0",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the Filter.",
			},
			"bandwidth_limit_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the bandwidth limit to use.",
			},
			"cache_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the cache to use.",
			},
			"default_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The default limit to use.",
			},
			"default_period": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The default period to use.",
			},
			"key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The key to use.",
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The limit to use.",
			},
			"min_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The minimum size to use.",
			},
			"spoe_config": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The spoe config to use.",
			},
			"spoe_engine": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The spoe engine to use.",
			},
			"table": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The table to use.",
			},
			"trace_hexdump": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The trace hexdump to use.",
			},
			"trace_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The trace name to use.",
			},
			"trace_rnd_forwarding": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The trace rnd forwarding to use.",
			},
			"trace_rnd_parsing": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The trace rnd parsing to use.",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type to use.",
			},
		},
	}
}

func resourceHaproxyFilterRead(d *schema.ResourceData, m interface{}) error {
	FilterName := d.Get("name").(string)
	indexName := d.Get("index").(int)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	configMap := m.(map[string]interface{})
	FilterConfig := configMap["filter"].(*ConfigFilter)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return FilterConfig.GetAFilterConfiguration(indexName, transactionID, parentName, parentType)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(FilterName, "error reading Filter configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(FilterName)
	return nil
}

func resourceHaproxyFilterCreate(d *schema.ResourceData, m interface{}) error {
	FilterName := d.Get("name").(string)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	payload := FilterPayload{
		Index:              d.Get("index").(int),
		AppName:            FilterName,
		BandwidthLimitName: d.Get("bandwidth_limit_name").(string),
		CacheName:          d.Get("cache_name").(string),
		DefaultLimit:       d.Get("default_limit").(int),
		DefaultPeriod:      d.Get("default_period").(int),
		Key:                d.Get("key").(string),
		Limit:              d.Get("limit").(int),
		MinSize:            d.Get("min_size").(int),
		SpoeConfig:         d.Get("spoe_config").(string),
		SpoeEngine:         d.Get("spoe_engine").(string),
		Table:              d.Get("table").(string),
		TraceHexdump:       d.Get("trace_hexdump").(bool),
		TraceName:          d.Get("trace_name").(string),
		TraceRndForwarding: d.Get("trace_rnd_forwarding").(bool),
		TraceRndParsing:    d.Get("trace_rnd_parsing").(bool),
		Type:               d.Get("type").(string),
	}

	excludeFields := []string{"index"}
	payloadJSON, err := utils.MarshalExcludeFields(payload, excludeFields)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	FilterConfig := configMap["filter"].(*ConfigFilter)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return FilterConfig.AddFilterConfiguration(payloadJSON, transactionID, parentName, parentType)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(FilterName, "error creating Filter configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(FilterName)
	return nil
}

func resourceHaproxyFilterUpdate(d *schema.ResourceData, m interface{}) error {
	FilterName := d.Get("name").(string)
	indexName := d.Get("index").(int)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	payload := FilterPayload{
		Index:              d.Get("index").(int),
		AppName:            FilterName,
		BandwidthLimitName: d.Get("bandwidth_limit_name").(string),
		CacheName:          d.Get("cache_name").(string),
		DefaultLimit:       d.Get("default_limit").(int),
		DefaultPeriod:      d.Get("default_period").(int),
		Key:                d.Get("key").(string),
		Limit:              d.Get("limit").(int),
		MinSize:            d.Get("min_size").(int),
		SpoeConfig:         d.Get("spoe_config").(string),
		SpoeEngine:         d.Get("spoe_engine").(string),
		Table:              d.Get("table").(string),
		TraceHexdump:       d.Get("trace_hexdump").(bool),
		TraceName:          d.Get("trace_name").(string),
		TraceRndForwarding: d.Get("trace_rnd_forwarding").(bool),
		TraceRndParsing:    d.Get("trace_rnd_parsing").(bool),
		Type:               d.Get("type").(string),
	}
	excludeFields := []string{"index"}
	payloadJSON, err := utils.MarshalExcludeFields(payload, excludeFields)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	FilterConfig := configMap["filter"].(*ConfigFilter)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return FilterConfig.UpdateFilterConfiguration(indexName, payloadJSON, transactionID, parentName, parentType)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(FilterName, "error updating Filter configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(FilterName)
	return nil
}

func resourceHaproxyFilterDelete(d *schema.ResourceData, m interface{}) error {
	FilterName := d.Get("name").(string)
	indexName := d.Get("index").(int)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	configMap := m.(map[string]interface{})
	FilterConfig := configMap["filter"].(*ConfigFilter)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return FilterConfig.DeleteFilterConfiguration(indexName, transactionID, parentName, parentType)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(FilterName, "error deleting Filter configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId("")
	return nil
}
