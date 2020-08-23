// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure IPv4 to IPv6 virtual IP groups.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallVipgrp46() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallVipgrp46Create,
		Read:   resourceFirewallVipgrp46Read,
		Update: resourceFirewallVipgrp46Update,
		Delete: resourceFirewallVipgrp46Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"color": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional: true,
			},
			"member": &schema.Schema{
				Type: schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceFirewallVipgrp46Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallVipgrp46(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallVipgrp46 resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallVipgrp46(obj)

	if err != nil {
		return fmt.Errorf("Error creating FirewallVipgrp46 resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallVipgrp46")
	}

	return resourceFirewallVipgrp46Read(d, m)
}

func resourceFirewallVipgrp46Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallVipgrp46(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVipgrp46 resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallVipgrp46(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVipgrp46 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallVipgrp46")
	}

	return resourceFirewallVipgrp46Read(d, m)
}

func resourceFirewallVipgrp46Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteFirewallVipgrp46(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallVipgrp46 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallVipgrp46Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadFirewallVipgrp46(mkey)
	if err != nil {
		return fmt.Errorf("Error reading FirewallVipgrp46 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallVipgrp46(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallVipgrp46 resource from API: %v", err)
	}
	return nil
}


func flattenFirewallVipgrp46Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipgrp46Uuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipgrp46Color(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipgrp46Comments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVipgrp46Member(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenFirewallVipgrp46MemberName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallVipgrp46MemberName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectFirewallVipgrp46(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("name", flattenFirewallVipgrp46Name(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallVipgrp46Uuid(o["uuid"], d, "uuid")); err != nil {
		if !fortiAPIPatch(o["uuid"]) {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("color", flattenFirewallVipgrp46Color(o["color"], d, "color")); err != nil {
		if !fortiAPIPatch(o["color"]) {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("comments", flattenFirewallVipgrp46Comments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

    if isImportTable() {
        if err = d.Set("member", flattenFirewallVipgrp46Member(o["member"], d, "member")); err != nil {
            if !fortiAPIPatch(o["member"]) {
                return fmt.Errorf("Error reading member: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("member"); ok {
            if err = d.Set("member", flattenFirewallVipgrp46Member(o["member"], d, "member")); err != nil {
                if !fortiAPIPatch(o["member"]) {
                    return fmt.Errorf("Error reading member: %v", err)
                }
            }
        }
    }


	return nil
}

func flattenFirewallVipgrp46FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandFirewallVipgrp46Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipgrp46Uuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipgrp46Color(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipgrp46Comments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipgrp46Member(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandFirewallVipgrp46MemberName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallVipgrp46MemberName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectFirewallVipgrp46(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallVipgrp46Name(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok {
		t, err := expandFirewallVipgrp46Uuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("color"); ok {
		t, err := expandFirewallVipgrp46Color(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandFirewallVipgrp46Comments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("member"); ok {
		t, err := expandFirewallVipgrp46Member(d, v, "member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}


	return &obj, nil
}
