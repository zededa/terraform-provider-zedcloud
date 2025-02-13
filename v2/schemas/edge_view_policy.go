package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func EdgeViewPolicyModel(d *schema.ResourceData) *models.EdgeViewPolicy {
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
	var metaData *models.PolicyCommon // PolicyCommon
	metaDataInterface, metaDataIsSet := d.GetOk("meta_data")
	if metaDataIsSet && metaDataInterface != nil {
		metaDataMap := metaDataInterface.([]interface{})
		if len(metaDataMap) > 0 {
			metaData = PolicyCommonModelFromMap(metaDataMap[0].(map[string]interface{}))
		}
	}
	return &models.EdgeViewPolicy{
		AccessAllowChange: accessAllowChange,
		EdgeviewAllow:     &edgeviewAllow, // bool
		Edgeviewcfg:       edgeviewcfg,
		MaxExpireSec:      maxExpireSec,
		MaxInst:           maxInst,
		MetaData:          metaData,
	}
}

func EdgeViewPolicyModelFromMap(m map[string]interface{}) *models.EdgeViewPolicy {
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
	var metaData *models.PolicyCommon                // PolicyCommon
	metaDataInterface, metaDataIsSet := m["meta_data"]
	if metaDataIsSet && metaDataInterface != nil {
		metaDataMap := metaDataInterface.([]interface{})
		if len(metaDataMap) > 0 {
			metaData = PolicyCommonModelFromMap(metaDataMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.EdgeViewPolicy{
		AccessAllowChange: accessAllowChange,
		EdgeviewAllow:     &edgeviewAllow,
		Edgeviewcfg:       edgeviewcfg,
		MaxExpireSec:      maxExpireSec,
		MaxInst:           maxInst,
		MetaData:          metaData,
	}
}

func SetEdgeViewPolicyResourceData(d *schema.ResourceData, m *models.EdgeViewPolicy) {
	d.Set("access_allow_change", m.AccessAllowChange)
	d.Set("edgeview_allow", m.EdgeviewAllow)
	d.Set("edgeviewcfg", SetEdgeviewCfgSubResourceData([]*models.EdgeviewCfg{m.Edgeviewcfg}))
	d.Set("max_expire_sec", m.MaxExpireSec)
	d.Set("max_inst", m.MaxInst)
	d.Set("meta_data", SetPolicyCommonSubResourceData([]*models.PolicyCommon{m.MetaData}))
}

func SetEdgeViewPolicySubResourceData(m []*models.EdgeViewPolicy) (d []*map[string]interface{}) {
	for _, EdgeViewPolicyModel := range m {
		if EdgeViewPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["access_allow_change"] = EdgeViewPolicyModel.AccessAllowChange
			properties["edgeview_allow"] = EdgeViewPolicyModel.EdgeviewAllow
			properties["edgeviewcfg"] = SetEdgeviewCfgSubResourceData([]*models.EdgeviewCfg{EdgeViewPolicyModel.Edgeviewcfg})
			properties["max_expire_sec"] = EdgeViewPolicyModel.MaxExpireSec
			properties["max_inst"] = EdgeViewPolicyModel.MaxInst
			properties["meta_data"] = SetPolicyCommonSubResourceData([]*models.PolicyCommon{EdgeViewPolicyModel.MetaData})
			d = append(d, &properties)
		}
	}
	return
}

func EdgeViewPolicy() map[string]*schema.Schema {
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

		"meta_data": {
			Description: `all the required metadata for a policy like id, name, different types of tags`,
			Type:        schema.TypeList, //GoType: PolicyCommon
			Elem: &schema.Resource{
				Schema: PolicyCommon(),
			},
			Optional: true,
		},
	}
}

func GetEdgeViewPolicyPropertyFields() (t []string) {
	return []string{
		"access_allow_change",
		"edgeview_allow",
		"edgeviewcfg",
		"max_expire_sec",
		"max_inst",
		"meta_data",
	}
}
