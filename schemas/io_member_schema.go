package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate IoMember resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func IoMemberModel(d *schema.ResourceData) *models.IoMember {
	assigngrp, _ := d.Get("assigngrp").(string)
	cbattr, _ := d.Get("cbattr").(map[string]string) // map[string]string
	costInt, _ := d.Get("cost").(int)
	cost := int64(costInt)
	logicallabel, _ := d.Get("logicallabel").(string)
	phyaddrs, _ := d.Get("phyaddrs").(map[string]string) // map[string]string
	phylabel, _ := d.Get("phylabel").(string)
	usageModel, _ := d.Get("usage").(models.AdapterUsage) // AdapterUsage
	usage := &usageModel
	if !ok {
		usage = nil
	}
	usagePolicy, _ := d.Get("usage_policy").(map[string]bool) // map[string]bool
	ztypeModel, _ := d.Get("ztype").(models.IoType)           // IoType
	ztype := &ztypeModel
	if !ok {
		ztype = nil
	}
	return &models.IoMember{
		Assigngrp:    &assigngrp, // string true false false
		Cbattr:       cbattr,
		Cost:         &cost,         // int64 true false false
		Logicallabel: &logicallabel, // string true false false
		Phyaddrs:     phyaddrs,
		Phylabel:     &phylabel, // string true false false
		Usage:        usage,
		UsagePolicy:  usagePolicy,
		Ztype:        ztype,
	}
}

func IoMemberModelFromMap(m map[string]interface{}) *models.IoMember {
	assigngrp := m["assigngrp"].(string)
	cbattr := m["cbattr"].(map[string]string)
	cost := int64(m["cost"].(int)) // int64 true false false
	logicallabel := m["logicallabel"].(string)
	phyaddrs := m["phyaddrs"].(map[string]string)
	phylabel := m["phylabel"].(string)
	usage := m["usage"].(*models.AdapterUsage) // AdapterUsage
	usagePolicy := m["usage_policy"].(map[string]bool)
	ztype := m["ztype"].(*models.IoType) // IoType
	return &models.IoMember{
		Assigngrp:    &assigngrp,
		Cbattr:       cbattr,
		Cost:         &cost,
		Logicallabel: &logicallabel,
		Phyaddrs:     phyaddrs,
		Phylabel:     &phylabel,
		Usage:        usage,
		UsagePolicy:  usagePolicy,
		Ztype:        ztype,
	}
}

// Update the underlying IoMember resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetIoMemberResourceData(d *schema.ResourceData, m *models.IoMember) {
	d.Set("assigngrp", m.Assigngrp)
	d.Set("cbattr", m.Cbattr)
	d.Set("cost", m.Cost)
	d.Set("logicallabel", m.Logicallabel)
	d.Set("phyaddrs", m.Phyaddrs)
	d.Set("phylabel", m.Phylabel)
	d.Set("usage", m.Usage)
	d.Set("usage_policy", m.UsagePolicy)
	d.Set("ztype", m.Ztype)
}

// Iterate through and update the IoMember resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetIoMemberSubResourceData(m []*models.IoMember) (d []*map[string]interface{}) {
	for _, IoMemberModel := range m {
		if IoMemberModel != nil {
			properties := make(map[string]interface{})
			properties["assigngrp"] = IoMemberModel.Assigngrp
			properties["cbattr"] = IoMemberModel.Cbattr
			properties["cost"] = IoMemberModel.Cost
			properties["logicallabel"] = IoMemberModel.Logicallabel
			properties["phyaddrs"] = IoMemberModel.Phyaddrs
			properties["phylabel"] = IoMemberModel.Phylabel
			properties["usage"] = IoMemberModel.Usage
			properties["usage_policy"] = IoMemberModel.UsagePolicy
			properties["ztype"] = IoMemberModel.Ztype
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the IoMember resource defined in the Terraform configuration
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

		"ztype": {
			Description: `Z Type`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

// Retrieve property field names for updating the IoMember resource
func GetIoMemberPropertyFields() (t []string) {
	return []string{
		"assigngrp",
		"cbattr",
		"cost",
		"logicallabel",
		"phyaddrs",
		"phylabel",
		"usage",
		"usage_policy",
		"ztype",
	}
}
