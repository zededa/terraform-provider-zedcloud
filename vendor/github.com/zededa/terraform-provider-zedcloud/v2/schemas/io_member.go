package schemas

import (
	"slices"
	"cmp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func IoMemberModel(d *schema.ResourceData) *models.IoMember {
	assigngrp, _ := d.Get("assigngrp").(string)
	cbattr := map[string]string{}
	cbattrInterface, cbattrIsSet := d.GetOk("cbattr")
	if cbattrIsSet {
		cbattrMap := cbattrInterface.(map[string]interface{})
		for k, v := range cbattrMap {
			if v == nil {
				continue
			}
			cbattr[k] = v.(string)
		}
	}

	costInt, _ := d.Get("cost").(int)
	cost := int64(costInt)
	logicallabel, _ := d.Get("logicallabel").(string)
	parentassigngrp, _ := d.Get("parentassigngrp").(string)
	phyaddrs := map[string]string{}
	phyaddrsInterface, phyaddrsIsSet := d.GetOk("phyaddrs")
	if phyaddrsIsSet {
		phyaddrsMap := phyaddrsInterface.(map[string]interface{})
		for k, v := range phyaddrsMap {
			if v == nil {
				continue
			}
			phyaddrs[k] = v.(string)
		}
	}

	phylabel, _ := d.Get("phylabel").(string)
	var usage *models.AdapterUsage // AdapterUsage
	usageInterface, usageIsSet := d.GetOk("usage")
	if usageIsSet {
		usageModel := usageInterface.(string)
		if usageModel == "" {
			usageModel = "ADAPTER_USAGE_UNSPECIFIED"
		}
		usage = models.NewAdapterUsage(models.AdapterUsage(usageModel))
	} 
	usagePolicy := map[string]bool{} // map[string]bool
	usagePolicyInterface, usagePolicyIsSet := d.GetOk("phyaddrs")
	if usagePolicyIsSet {
		usagePolicyMap := usagePolicyInterface.(map[string]interface{})
		for k, v := range usagePolicyMap {
			if v == nil {
				continue
			}
			usagePolicy[k] = v.(bool)
		}
	}

	var vfs *models.Vfs // Vfs
	vfsInterface, vfsIsSet := d.GetOk("vfs")
	if vfsIsSet && vfsInterface != nil {
		vfsMap := vfsInterface.([]interface{})
		if len(vfsMap) > 0 {
			vfs = VfsModelFromMap(vfsMap[0].(map[string]interface{}))
		}
	}
	var ztype *models.IoType // IoType
	ztypeInterface, ztypeIsSet := d.GetOk("ztype")
	if ztypeIsSet {
		ztypeModel := ztypeInterface.(string)
		ztype = models.NewIoType(models.IoType(ztypeModel))
	}
	return &models.IoMember{
		Assigngrp:       &assigngrp, // string
		Cbattr:          cbattr,
		Cost:            &cost,         // int64
		Logicallabel:    &logicallabel, // string
		Parentassigngrp: parentassigngrp,
		Phyaddrs:        phyaddrs,
		Phylabel:        &phylabel, // string
		Usage:           usage,
		UsagePolicy:     usagePolicy,
		Vfs:             vfs,
		Ztype:           ztype,
	}
}

func IoMemberModelFromMap(m map[string]interface{}) *models.IoMember {
	assigngrp := m["assigngrp"].(string)
	cbattr := map[string]string{}
	cbattrInterface, cbattrIsSet := m["cbattr"]
	if cbattrIsSet {
		cbattrMap := cbattrInterface.(map[string]interface{})
		for k, v := range cbattrMap {
			if v == nil {
				continue
			}
			cbattr[k] = v.(string)
		}
	}

	cost := int64(m["cost"].(int)) // int64
	logicallabel := m["logicallabel"].(string)
	parentassigngrp := m["parentassigngrp"].(string)
	phyaddrs := map[string]string{}
	phyaddrsInterface, phyaddrsIsSet := m["phyaddrs"]
	if phyaddrsIsSet {
		phyaddrsMap := phyaddrsInterface.(map[string]interface{})
		for k, v := range phyaddrsMap {
			if v == nil {
				continue
			}
			phyaddrs[k] = v.(string)
		}
	}

	phylabel := m["phylabel"].(string)
	var usage *models.AdapterUsage // AdapterUsage
	usageInterface, usageIsSet := m["usage"]
	if usageIsSet {
		usageModel := usageInterface.(string)
		if usageModel == "" {
			usageModel = "ADAPTER_USAGE_UNSPECIFIED"
		}
		usage = models.NewAdapterUsage(models.AdapterUsage(usageModel))
	} 
	usagePolicy := map[string]bool{} // map[string]bool
	usagePolicyInterface, usagePolicyIsSet := m["usage_policy"]
	if usagePolicyIsSet {
		usagePolicyMap := usagePolicyInterface.(map[string]interface{})
		for k, v := range usagePolicyMap {
			if v == nil {
				continue
			}
			usagePolicy[k] = v.(bool)
		}
	}

	var vfs *models.Vfs // Vfs
	vfsInterface, vfsIsSet := m["vfs"]
	if vfsIsSet && vfsInterface != nil {
		vfsMap := vfsInterface.([]interface{})
		if len(vfsMap) > 0 {
			vfs = VfsModelFromMap(vfsMap[0].(map[string]interface{}))
		}
	}
	//
	var ztype *models.IoType // IoType
	ztypeInterface, ztypeIsSet := m["ztype"]
	if ztypeIsSet {
		ztypeModel := ztypeInterface.(string)
		ztype = models.NewIoType(models.IoType(ztypeModel))
	}
	return &models.IoMember{
		Assigngrp:       &assigngrp,
		Cbattr:          cbattr,
		Cost:            &cost,
		Logicallabel:    &logicallabel,
		Parentassigngrp: parentassigngrp,
		Phyaddrs:        phyaddrs,
		Phylabel:        &phylabel,
		Usage:           usage,
		UsagePolicy:     usagePolicy,
		Vfs:             vfs,
		Ztype:           ztype,
	}
}

func SetIoMemberResourceData(d *schema.ResourceData, m *models.IoMember) {
	d.Set("assigngrp", m.Assigngrp)
	d.Set("cbattr", m.Cbattr)
	d.Set("cost", m.Cost)
	d.Set("logicallabel", m.Logicallabel)
	d.Set("parentassigngrp", m.Parentassigngrp)
	d.Set("phyaddrs", m.Phyaddrs)
	d.Set("phylabel", m.Phylabel)
	d.Set("usage", m.Usage)
	d.Set("usage_policy", m.UsagePolicy)
	d.Set("vfs", SetVfsSubResourceData([]*models.Vfs{m.Vfs}))
	d.Set("ztype", m.Ztype)
}

func SetIoMemberSubResourceData(m []*models.IoMember) (d []*map[string]interface{}) {
	for _, IoMemberModel := range m {
		if IoMemberModel != nil {
			properties := make(map[string]interface{})
			properties["assigngrp"] = IoMemberModel.Assigngrp
			properties["cbattr"] = IoMemberModel.Cbattr
			properties["cost"] = IoMemberModel.Cost
			properties["logicallabel"] = IoMemberModel.Logicallabel
			properties["parentassigngrp"] = IoMemberModel.Parentassigngrp
			properties["phyaddrs"] = IoMemberModel.Phyaddrs
			properties["phylabel"] = IoMemberModel.Phylabel
			properties["usage"] = IoMemberModel.Usage
			properties["usage_policy"] = IoMemberModel.UsagePolicy
			properties["vfs"] = SetVfsSubResourceData([]*models.Vfs{IoMemberModel.Vfs})
			properties["ztype"] = IoMemberModel.Ztype
			d = append(d, &properties)
		}
	}
	return
}

func IoMemberSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"assigngrp": {
			Description: `Assign Group`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"cbattr": {
			Description: `attributes`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"cost": {
			Description: `cost of using this ioMember. Default is 0.`,
			Type:        schema.TypeInt,
			Required:    true,
		},

		"logicallabel": {
			Description: `Logical Label`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"parentassigngrp": {
			Description: `Parent group for a IoMember. Can be empty if there is no parent group.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"phyaddrs": {
			Description: `Map of Physical Addresses`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Required: true,
		},

		"phylabel": {
			Description: `Physical Label`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"usage": {
			Description: `Adopter Usage`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"usage_policy": {
			Description: ``,
			Type:        schema.TypeMap, //GoType: map[string]bool
			Elem: &schema.Schema{
				Type: schema.TypeBool,
			},
			Optional: true,
		},

		"vfs": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Vfs
			Elem: &schema.Resource{
				Schema: VfsSchema(),
			},
			Optional: true,
		},

		"ztype": {
			Description: `Z Type`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func GetIoMemberPropertyFields() (t []string) {
	return []string{
		"assigngrp",
		"cbattr",
		"cost",
		"logicallabel",
		"parentassigngrp",
		"phyaddrs",
		"phylabel",
		"usage",
		"usage_policy",
		"vfs",
		"ztype",
	}
}

func CompareIoMemberLists(a, b []*models.IoMember) bool {
	if len(a) != len(b) {
		return false
	}
	slices.SortStableFunc(a, func(left, right *models.IoMember) int {
		return cmp.Compare(*left.Logicallabel, *right.Logicallabel)
	})
	slices.SortStableFunc(b, func(left, right *models.IoMember) int {
		return cmp.Compare(*left.Logicallabel, *right.Logicallabel)
	})
	for i, v := range a {
		if cmp.Compare(*v.Logicallabel, *b[i].Logicallabel) != 0 {
			return false
		}
		if cmp.Compare(*v.Phylabel, *b[i].Phylabel) != 0 {
			return false
		}
		if cmp.Compare(*v.Ztype, *b[i].Ztype) != 0 {
			return false
		}
		if cmp.Compare(*v.Cost, *b[i].Cost) != 0 {
			return false
		}
		if cmp.Compare(*v.Assigngrp, *b[i].Assigngrp) != 0 {
			return false
		}
		if cmp.Compare(v.Parentassigngrp, b[i].Parentassigngrp) != 0 {
			return false
		}
		if cmp.Compare(*v.Usage, *b[i].Usage) != 0 {
			return false
		}
	}
	return true
}