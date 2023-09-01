package cache

import (
	"fmt"
	"net/http"
	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ResourceHaproxyCache() *schema.Resource {
	return &schema.Resource{
		Create: resourceHaproxyCacheCreate,
		Read:   resourceHaproxyCacheRead,
		Update: resourceHaproxyCacheUpdate,
		Delete: resourceHaproxyCacheDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the Cache. It must be unique and cannot be changed.",
			},
			"max_age": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The max age of the Cache",
			},
			"max_object_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The max object size of the Cache",
			},
			"max_secondary_entries": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "the max secondary entries of the Cache",
			},
			"process_vary": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The process vary of the Cache",
			},
			"total_max_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The total max size of the Cache",
			},
		},
	}
}

func resourceHaproxyCacheRead(d *schema.ResourceData, m interface{}) error {
	cacheName := d.Get("name").(string)

	configMap := m.(map[string]interface{})
	CacheConfig := configMap["cache"].(*ConfigCache)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return CacheConfig.GetACacheConfiguration(cacheName, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(cacheName, "error reading Cache configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(cacheName)
	return nil
}

func resourceHaproxyCacheCreate(d *schema.ResourceData, m interface{}) error {
	cacheName := d.Get("name").(string)

	payload := CachePayload{
		Name:                d.Get("name").(string),
		MaxAge:              d.Get("max_age").(int),
		MaxObjectSize:       d.Get("max_object_size").(int),
		MaxSecondaryEntries: d.Get("max_secondary_entries").(int),
		ProcessVary:         d.Get("process_vary").(bool),
		TotalMaxSize:        d.Get("total_max_size").(int),
	}

	payloadJSON, err := utils.MarshalNonZeroFields(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	CacheConfig := configMap["cache"].(*ConfigCache)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return CacheConfig.AddCacheConfiguration(payloadJSON, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(cacheName, "error creating Cache configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(cacheName)
	return nil
}

func resourceHaproxyCacheUpdate(d *schema.ResourceData, m interface{}) error {
	cacheName := d.Get("name").(string)

	payload := CachePayload{
		Name:                d.Get("name").(string),
		MaxAge:              d.Get("max_age").(int),
		MaxObjectSize:       d.Get("max_object_size").(int),
		MaxSecondaryEntries: d.Get("max_secondary_entries").(int),
		ProcessVary:         d.Get("process_vary").(bool),
		TotalMaxSize:        d.Get("total_max_size").(int),
	}

	payloadJSON, err := utils.MarshalNonZeroFields(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	CacheConfig := configMap["cache"].(*ConfigCache)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return CacheConfig.UpdateCacheConfiguration(cacheName, payloadJSON, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(cacheName, "error updating Cache configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(cacheName)
	return nil
}

func resourceHaproxyCacheDelete(d *schema.ResourceData, m interface{}) error {
	cacheName := d.Get("name").(string)

	configMap := m.(map[string]interface{})
	CacheConfig := configMap["cache"].(*ConfigCache)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return CacheConfig.DeleteCacheConfiguration(cacheName, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(cacheName, "error deleting Cache configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId("")
	return nil
}
