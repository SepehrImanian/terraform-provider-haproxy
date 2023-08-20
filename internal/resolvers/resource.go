package resolvers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ResourceHaproxyResolvers() *schema.Resource {
	return &schema.Resource{
		Create: resourceHaproxyResolversCreate,
		Read:   resourceHaproxyResolversRead,
		Update: resourceHaproxyResolversUpdate,
		Delete: resourceHaproxyResolversDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the Resolvers. It must be unique and cannot be changed.",
			},
			"accepted_payload_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The Accepted Payload Size of the Resolvers.",
			},
			"hold_nx": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The hold NX of the Resolvers.",
			},
			"hold_other": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The hold other of the Resolvers.",
			},
			"hold_refused": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The hold refused of the Resolvers.",
			},
			"hold_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The hold timeout of the Resolvers.",
			},
			"hold_valid": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The hold valid of the Resolvers.",
			},
			"parse_resolv_conf": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The pasre-resolv-conf of the Resolvers. it could be true or false",
			},
			"resolve_retries": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The retries of the Resolvers.",
			},
			"timeout_resolve": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The timeout resolve of the Resolvers.",
			},
			"timeout_retry": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The timeout retry of the Resolvers.",
			},
		},
	}
}

func resourceHaproxyResolversRead(d *schema.ResourceData, m interface{}) error {
	resolversName := d.Get("name").(string)

	configMap := m.(map[string]interface{})
	ResolversConfig := configMap["resolvers"].(*ConfigResolvers)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return ResolversConfig.GetAResolversConfiguration(resolversName, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(resolversName, "error reading Resolvers configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(resolversName)
	return nil
}

func resourceHaproxyResolversCreate(d *schema.ResourceData, m interface{}) error {
	resolversName := d.Get("name").(string)

	payload := ResolversPayload{
		Name:                resolversName,
		AcceptedPayloadSize: d.Get("accepted_payload_size").(int),
		HoldNx:              d.Get("hold_nx").(int),
		HoldOther:           d.Get("hold_other").(int),
		HoldRefused:         d.Get("hold_refused").(int),
		HoldTimeout:         d.Get("hold_timeout").(int),
		HoldValid:           d.Get("hold_valid").(int),
		ParseResolvConf:     d.Get("parse_resolv_conf").(bool),
		ResolveRetries:      d.Get("resolve_retries").(int),
		TimeoutResolve:      d.Get("timeout_resolve").(int),
		TimeoutRetry:        d.Get("timeout_retry").(int),
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	ResolversConfig := configMap["resolvers"].(*ConfigResolvers)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return ResolversConfig.AddResolversConfiguration(payloadJSON, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(resolversName, "error creating Resolvers configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(resolversName)
	return nil
}

func resourceHaproxyResolversUpdate(d *schema.ResourceData, m interface{}) error {
	resolversName := d.Get("name").(string)

	payload := ResolversPayload{
		Name:                resolversName,
		AcceptedPayloadSize: d.Get("accepted_payload_size").(int),
		HoldNx:              d.Get("hold_nx").(int),
		HoldOther:           d.Get("hold_other").(int),
		HoldRefused:         d.Get("hold_refused").(int),
		HoldTimeout:         d.Get("hold_timeout").(int),
		HoldValid:           d.Get("hold_valid").(int),
		ParseResolvConf:     d.Get("parse_resolv_conf").(bool),
		ResolveRetries:      d.Get("resolve_retries").(int),
		TimeoutResolve:      d.Get("timeout_resolve").(int),
		TimeoutRetry:        d.Get("timeout_retry").(int),
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	ResolversConfig := configMap["resolvers"].(*ConfigResolvers)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return ResolversConfig.UpdateResolversConfiguration(resolversName, payloadJSON, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(resolversName, "error updating Resolvers configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(resolversName)
	return nil
}

func resourceHaproxyResolversDelete(d *schema.ResourceData, m interface{}) error {
	ResolversName := d.Get("name").(string)

	configMap := m.(map[string]interface{})
	ResolversConfig := configMap["resolvers"].(*ConfigResolvers)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return ResolversConfig.DeleteResolversConfiguration(ResolversName, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(ResolversName, "error deleting Resolvers configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId("")
	return nil
}
