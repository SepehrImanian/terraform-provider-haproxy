package global

import (
	"fmt"
	"net/http"
	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ResourceHaproxyGlobal() *schema.Resource {
	return &schema.Resource{
		Read:   resourceHaproxyGlobalRead,
		Create: resourceHaproxyGlobalUpdate,
		Update: resourceHaproxyGlobalUpdate,
		Delete: resourceHaproxyGlobalDelete,

		Schema: map[string]*schema.Schema{
			"chroot": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Chroot directory",
			},
			"user": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User name",
			},
			"group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Group name",
			},
			"master_worker": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Master worker mode",
			},
			"process": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Process name",
			},
			"cpu_set": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "CPU set",
			},
			"daemon": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Run in daemon mode",
			},
			"maxcompcpuusage": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Maximum CPU usage in percent",
			},
			"maxpipes": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Maximum number of pipes",
			},
			"maxsslconn": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Maximum number of SSL connections",
			},
			"maxconn": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Maximum number of connections",
			},
			"nbproc": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Number of processes",
			},
			"nbthread": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Number of threads",
			},
			"pidfile": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "PID file",
			},
			"ulimit_n": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Ulimit number",
			},
			"crt_base": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Certificate base directory",
			},
			"ca_base": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "CA base directory",
			},
			"stats_maxconn": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Stats maximum number of connections",
			},
			"stats_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Stats timeout",
			},
			"ssl_default_bind_ciphers": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "SSL default bind ciphers",
			},
			"ssl_default_bind_options": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "SSL default bind options",
			},
		},
	}
}

func resourceHaproxyGlobalRead(d *schema.ResourceData, m interface{}) error {
	configMap := m.(map[string]interface{})
	GlobalConfig := configMap["global"].(*ConfigGlobal)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return GlobalConfig.GetAGlobalConfiguration(transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError("global", "error reading Global configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId("global")
	return nil
}

func resourceHaproxyGlobalUpdate(d *schema.ResourceData, m interface{}) error {
	daemon := d.Get("daemon").(bool)

	// Convert bool to string
	daemonStr := utils.BoolToStr(daemon)

	payload := GlobalPayload{
		User:   d.Get("user").(string),
		Group:  d.Get("group").(string),
		Chroot: d.Get("chroot").(string),
		CpuMaps: CpuMaps{
			Process: d.Get("process").(string),
			CpuSet:  d.Get("cpu_set").(string),
		},
		Daemon:                daemonStr,
		MasterWorker:          d.Get("master_worker").(bool),
		MaxCompCpuUsage:       d.Get("maxcompcpuusage").(int),
		MaxPipes:              d.Get("maxpipes").(int),
		MaxSslConn:            d.Get("maxsslconn").(int),
		MaxConn:               d.Get("maxconn").(int),
		NbProc:                d.Get("nbproc").(int),
		NbThread:              d.Get("nbthread").(int),
		PidFile:               d.Get("pidfile").(string),
		UlimitN:               d.Get("ulimit_n").(int),
		CrtBase:               d.Get("crt_base").(string),
		CaBase:                d.Get("ca_base").(string),
		StatsMaxConn:          d.Get("stats_maxconn").(int),
		StatsTimeOut:          d.Get("stats_timeout").(int),
		SslDefaultBindCiphers: d.Get("ssl_default_bind_ciphers").(string),
		SslDefaultBindOptions: d.Get("ssl_default_bind_options").(string),
	}

	payloadJSON, err := utils.MarshalNonZeroFields(payload)
	if err != nil {
		return err
	}

	configMap := m.(map[string]interface{})
	GlobalConfig := configMap["global"].(*ConfigGlobal)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return GlobalConfig.UpdateGlobalConfiguration(payloadJSON, transactionID)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError("global", "error updating Global configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId("global")
	return nil
}

func resourceHaproxyGlobalDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
