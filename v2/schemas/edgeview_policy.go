package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func EdgeviewPolicyModel(d *schema.ResourceData) *models.EdgeviewPolicy {
	accessAllowChange, _ := d.Get("access_allow_change").(bool)
	edgeviewAllow, _ := d.Get("edgeview_allow").(bool)
	var edgeviewcfg *models.EdgeviewCfg // EdgeviewCfg
	edgeviewcfgInterface, edgeviewcfgIsSet := d.GetOk("edgeviewcfg")
	if edgeviewcfgIsSet && edgeviewcfgInterface != nil {
		edgeviewcfgMap := edgeviewcfgInterface.([]interface{})
		if len(edgeviewcfgMap) > 0 {
			edgeviewcfg = EdgeviewCfgModelFromMap(edgeviewcfgMap[0].(map[string]interface{}))
		}
	}
	maxExpireSecInt, _ := d.Get("max_expire_sec").(int)
	maxExpireSec := int64(maxExpireSecInt)
	maxInstInt, _ := d.Get("max_inst").(int)
	maxInst := int64(maxInstInt)
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
	if edgeviewcfgIsSet && edgeviewcfgInterface != nil {
		edgeviewcfgMap := edgeviewcfgInterface.([]interface{})
		if len(edgeviewcfgMap) > 0 {
			edgeviewcfg = EdgeviewCfgModelFromMap(edgeviewcfgMap[0].(map[string]interface{}))
		}
	}
	//
	maxExpireSec := int64(m["max_expire_sec"].(int)) // int64
	maxInst := int64(m["max_inst"].(int))            // int64
	return &models.EdgeviewPolicy{
		AccessAllowChange: accessAllowChange,
		EdgeviewAllow:     &edgeviewAllow,
		Edgeviewcfg:       edgeviewcfg,
		MaxExpireSec:      maxExpireSec,
		MaxInst:           maxInst,
	}
}

func SetEdgeviewPolicyResourceData(d *schema.ResourceData, m *models.EdgeviewPolicy) {
	d.Set("access_allow_change", m.AccessAllowChange)
	d.Set("edgeview_allow", m.EdgeviewAllow)
	d.Set("edgeviewcfg", SetEdgeviewCfgSubResourceData([]*models.EdgeviewCfg{m.Edgeviewcfg}))
	d.Set("max_expire_sec", m.MaxExpireSec)
	d.Set("max_inst", m.MaxInst)
}

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

func EdgeviewPolicy() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access_allow_change": {
			Description: `Allow inherit instance to change access policy`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"edgeview_allow": {
			Description: `Allow device to enable Edgeview in this project`,
			Type:        schema.TypeBool,
			// Required:    true,
			Optional: true,
		},

		"edgeviewcfg": {
			Description: `Edgeview configuration and policies`,
			Type:        schema.TypeList, //GoType: EdgeviewCfg
			Elem: &schema.Resource{
				Schema: EdgeView(),
			},
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

func GetEdgeviewPolicyPropertyFields() (t []string) {
	return []string{
		"access_allow_change",
		"edgeview_allow",
		"edgeviewcfg",
		"max_expire_sec",
		"max_inst",
	}
}
