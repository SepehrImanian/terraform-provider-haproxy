package nameserver

import (
	"fmt"
	"net/http"

	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ResourceHaproxyNameserver() *schema.Resource {
	return &schema.Resource{
		Create: resourceHaproxyNameserverCreate,
		Read:   resourceHaproxyNameserverRead,
		Update: resourceHaproxyNameserverUpdate,
		Delete: resourceHaproxyNameserverDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the Nameserver. It must be unique and cannot be changed.",
			},
			"resolver": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The resolver of the Nameserver",
			},
			"port": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The port of the Nameserver",
			},
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The address of the Nameserver",
			},
		},
	}
}

func resourceHaproxyNameserverRead(d *schema.ResourceData, m interface{}) error {
	nameserverName := d.Get("name").(string)
	resolver := d.Get("resolver").(string)

	configMap := m.(map[string]interface{})
	NameserverConfig := configMap["nameserver"].(*ConfigNameserver)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, _, err := tranConfig.TransactionWithData(func(transactionID string) (*http.Response, []byte, error) {
		return NameserverConfig.GetANameserversConfiguration(nameserverName, transactionID, resolver)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(nameserverName, "error reading Nameserver configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(nameserverName)
	return nil
}

func resourceHaproxyNameserverCreate(d *schema.ResourceData, m interface{}) error {
	nameserverName := d.Get("name").(string)
	resolver := d.Get("resolver").(string)

	payload := NameserverPayload{
		Name:    nameserverName,
		Address: d.Get("address").(string),
		Port:    d.Get("port").(int),
	}

	payloadJSON, err := utils.MarshalNonZeroFields(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	NameserverConfig := configMap["nameserver"].(*ConfigNameserver)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, _, err := tranConfig.TransactionWithData(func(transactionID string) (*http.Response, []byte, error) {
		return NameserverConfig.AddNameserversConfiguration(payloadJSON, transactionID, resolver)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(nameserverName, "error creating Nameserver configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(nameserverName)
	return nil
}

func resourceHaproxyNameserverUpdate(d *schema.ResourceData, m interface{}) error {
	NameserverName := d.Get("name").(string)
	resolver := d.Get("resolver").(string)

	payload := NameserverPayload{
		Name:    NameserverName,
		Address: d.Get("address").(string),
		Port:    d.Get("port").(int),
	}

	payloadJSON, err := utils.MarshalNonZeroFields(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	NameserverConfig := configMap["nameserver"].(*ConfigNameserver)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, _, err := tranConfig.TransactionWithData(func(transactionID string) (*http.Response, []byte, error) {
		return NameserverConfig.UpdateNameserversConfiguration(NameserverName, payloadJSON, transactionID, resolver)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(NameserverName, "error updating Nameserver configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(NameserverName)
	return nil
}

func resourceHaproxyNameserverDelete(d *schema.ResourceData, m interface{}) error {
	NameserverName := d.Get("name").(string)
	resolver := d.Get("resolver").(string)

	configMap := m.(map[string]interface{})
	NameserverConfig := configMap["nameserver"].(*ConfigNameserver)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, _, err := tranConfig.TransactionWithData(func(transactionID string) (*http.Response, []byte, error) {
		return NameserverConfig.DeleteNameserversConfiguration(NameserverName, transactionID, resolver)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(NameserverName, "error deleting Nameserver configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId("")
	return nil
}
