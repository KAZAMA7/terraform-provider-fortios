// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: CA certificate.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceVpnCertificateCa() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnCertificateCaCreate,
		Read:   resourceVpnCertificateCaRead,
		Update: resourceVpnCertificateCaUpdate,
		Delete: resourceVpnCertificateCaDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Required:     true,
			},
			"ca": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusted": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scep_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"auto_update_days": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auto_update_days_warning": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_updated": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceVpnCertificateCaCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectVpnCertificateCa(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnCertificateCa resource while getting object: %v", err)
	}

	o, err := c.CreateVpnCertificateCa(obj)

	if err != nil {
		return fmt.Errorf("Error creating VpnCertificateCa resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnCertificateCa")
	}

	return resourceVpnCertificateCaRead(d, m)
}

func resourceVpnCertificateCaUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectVpnCertificateCa(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateCa resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnCertificateCa(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateCa resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnCertificateCa")
	}

	return resourceVpnCertificateCaRead(d, m)
}

func resourceVpnCertificateCaDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteVpnCertificateCa(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting VpnCertificateCa resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnCertificateCaRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadVpnCertificateCa(mkey)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateCa resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnCertificateCa(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateCa resource from API: %v", err)
	}
	return nil
}

func flattenVpnCertificateCaName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCaCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCaRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCaSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCaTrusted(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCaScepUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCaAutoUpdateDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCaAutoUpdateDaysWarning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCaSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateCaLastUpdated(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnCertificateCa(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenVpnCertificateCaName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("ca", flattenVpnCertificateCaCa(o["ca"], d, "ca")); err != nil {
		if !fortiAPIPatch(o["ca"]) {
			return fmt.Errorf("Error reading ca: %v", err)
		}
	}

	if err = d.Set("range", flattenVpnCertificateCaRange(o["range"], d, "range")); err != nil {
		if !fortiAPIPatch(o["range"]) {
			return fmt.Errorf("Error reading range: %v", err)
		}
	}

	if err = d.Set("source", flattenVpnCertificateCaSource(o["source"], d, "source")); err != nil {
		if !fortiAPIPatch(o["source"]) {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("trusted", flattenVpnCertificateCaTrusted(o["trusted"], d, "trusted")); err != nil {
		if !fortiAPIPatch(o["trusted"]) {
			return fmt.Errorf("Error reading trusted: %v", err)
		}
	}

	if err = d.Set("scep_url", flattenVpnCertificateCaScepUrl(o["scep-url"], d, "scep_url")); err != nil {
		if !fortiAPIPatch(o["scep-url"]) {
			return fmt.Errorf("Error reading scep_url: %v", err)
		}
	}

	if err = d.Set("auto_update_days", flattenVpnCertificateCaAutoUpdateDays(o["auto-update-days"], d, "auto_update_days")); err != nil {
		if !fortiAPIPatch(o["auto-update-days"]) {
			return fmt.Errorf("Error reading auto_update_days: %v", err)
		}
	}

	if err = d.Set("auto_update_days_warning", flattenVpnCertificateCaAutoUpdateDaysWarning(o["auto-update-days-warning"], d, "auto_update_days_warning")); err != nil {
		if !fortiAPIPatch(o["auto-update-days-warning"]) {
			return fmt.Errorf("Error reading auto_update_days_warning: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenVpnCertificateCaSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("last_updated", flattenVpnCertificateCaLastUpdated(o["last-updated"], d, "last_updated")); err != nil {
		if !fortiAPIPatch(o["last-updated"]) {
			return fmt.Errorf("Error reading last_updated: %v", err)
		}
	}

	return nil
}

func flattenVpnCertificateCaFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnCertificateCaName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCaCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCaRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCaSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCaTrusted(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCaScepUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCaAutoUpdateDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCaAutoUpdateDaysWarning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCaSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCaLastUpdated(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnCertificateCa(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandVpnCertificateCaName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("ca"); ok {
		t, err := expandVpnCertificateCaCa(d, v, "ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca"] = t
		}
	}

	if v, ok := d.GetOk("range"); ok {
		t, err := expandVpnCertificateCaRange(d, v, "range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok {
		t, err := expandVpnCertificateCaSource(d, v, "source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	if v, ok := d.GetOk("trusted"); ok {
		t, err := expandVpnCertificateCaTrusted(d, v, "trusted")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusted"] = t
		}
	}

	if v, ok := d.GetOk("scep_url"); ok {
		t, err := expandVpnCertificateCaScepUrl(d, v, "scep_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-url"] = t
		}
	}

	if v, ok := d.GetOk("auto_update_days"); ok {
		t, err := expandVpnCertificateCaAutoUpdateDays(d, v, "auto_update_days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-update-days"] = t
		}
	}

	if v, ok := d.GetOk("auto_update_days_warning"); ok {
		t, err := expandVpnCertificateCaAutoUpdateDaysWarning(d, v, "auto_update_days_warning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-update-days-warning"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandVpnCertificateCaSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("last_updated"); ok {
		t, err := expandVpnCertificateCaLastUpdated(d, v, "last_updated")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last-updated"] = t
		}
	}

	return &obj, nil
}
