package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate EdgeviewPolicy resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func EdgeviewPolicyModel(d *schema.ResourceData) *models.EdgeviewPolicy {
	accessAllowChange := d.Get("access_allow_change").(bool)
	edgeviewAllow := d.Get("edgeview_allow").(bool)
	var edgeviewcfg *models.EdgeviewCfg // EdgeviewCfg
	edgeviewcfgInterface, edgeviewcfgIsSet := d.GetOk("edgeviewcfg")
	if edgeviewcfgIsSet {
		edgeviewcfgMap := edgeviewcfgInterface.([]interface{})[0].(map[string]interface{})
		edgeviewcfg = EdgeviewCfgModelFromMap(edgeviewcfgMap)
	}
	maxExpireSec := int64(d.Get("max_expire_sec").(int))
	maxInst := int64(d.Get("max_inst").(int))
	return &models.EdgeviewPolicy{
		AccessAllowChange: accessAllowChange,
		EdgeviewAllow:     &edgeviewAllow, // bool true false false
		Edgeviewcfg:       edgeviewcfg,
		MaxExpireSec:      maxExpireSec,
		MaxInst:           maxInst,
	}
}

func EdgeviewPolicyModelFromMap(m map[string]interface{}) *models.EdgeviewPolicy {
	accessAllowChange := m["access_allow_change"].(bool)
	edgeviewAllow := m["edgeview_allow"].(bool)
	var edgeviewcfg *models.EdgeviewCfg // EdgeviewCfg
	edgeviewcfgInterface, edgeviewcfgIsSet := m["edgeviewcfg"]
	if edgeviewcfgIsSet {
		edgeviewcfgMap := edgeviewcfgInterface.([]interface{})[0].(map[string]interface{})
		edgeviewcfg = EdgeviewCfgModelFromMap(edgeviewcfgMap)
	}
	//
	maxExpireSec := int64(m["max_expire_sec"].(int)) // int64 false false false
	maxInst := int64(m["max_inst"].(int))            // int64 false false false
	return &models.EdgeviewPolicy{
		AccessAllowChange: accessAllowChange,
		EdgeviewAllow:     &edgeviewAllow,
		Edgeviewcfg:       edgeviewcfg,
		MaxExpireSec:      maxExpireSec,
		MaxInst:           maxInst,
	}
}

// Update the underlying EdgeviewPolicy resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetEdgeviewPolicyResourceData(d *schema.ResourceData, m *models.EdgeviewPolicy) {
	d.Set("access_allow_change", m.AccessAllowChange)
	d.Set("edgeview_allow", m.EdgeviewAllow)
	d.Set("edgeviewcfg", SetEdgeviewCfgSubResourceData([]*models.EdgeviewCfg{m.Edgeviewcfg}))
	d.Set("max_expire_sec", m.MaxExpireSec)
	d.Set("max_inst", m.MaxInst)
}

// Iterate throught and update the EdgeviewPolicy resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetEdgeviewPolicySubResourceData(m []*models.EdgeviewPolicy) (d []*map[string]interface{}) {
	for _, EdgeviewPolicyModel := range m {
		if EdgeviewPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["access_allow_change"] = EdgeviewPolicyModel.AccessAllowChange
			properties["edgeview_allow"] = EdgeviewPolicyModel.EdgeviewAllow
			properties["edgeviewcfg"] = SetEdgeviewCfgSubResourceData([]*models.EdgeviewCfg{EdgeviewPolicyModel.Edgeviewcfg})
			properties["max_expire_sec"] = EdgeviewPolicyModel.MaxExpireSec
			properties["max_inst"] = EdgeviewPolicyModel.MaxInst
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the EdgeviewPolicy resource defined in the Terraform configuration
func EdgeviewPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access_allow_change": {
			Description: `Allow inherit instance to change access policy`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"edgeview_allow": {
			Description: `Allow device to enable Edgeview in this project`,
			Type:        schema.TypeBool,
			Required:    true,
		},

		"edgeviewcfg": {
			Description: `Edgeview configuration and policies`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},

		"max_expire_sec": {
			Description: `Maximum seconds allowed for Edgeview session`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"max_inst": {
			Description: `Maximum instances allowed for Edgeview`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the EdgeviewPolicy resource
func GetEdgeviewPolicyPropertyFields() (t []string) {
	return []string{
		"access_allow_change",
		"edgeview_allow",
		"edgeviewcfg",
		"max_expire_sec",
		"max_inst",
	}
}
