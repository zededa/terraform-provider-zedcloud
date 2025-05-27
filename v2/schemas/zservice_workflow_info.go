package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZserviceWorkflowInfoModel(d *schema.ResourceData) *models.ZserviceWorkflowInfo {
	deviceid, _ := d.Get("deviceid").(string)
	workflowid, _ := d.Get("workflowid").(string)
	return &models.ZserviceWorkflowInfo{
		Deviceid:   deviceid,
		Workflowid: workflowid,
	}
}

func ZserviceWorkflowInfoModelFromMap(m map[string]interface{}) *models.ZserviceWorkflowInfo {
	deviceid := m["deviceid"].(string)
	workflowid := m["workflowid"].(string)
	return &models.ZserviceWorkflowInfo{
		Deviceid:   deviceid,
		Workflowid: workflowid,
	}
}

func SetZserviceWorkflowInfoResourceData(d *schema.ResourceData, m *models.ZserviceWorkflowInfo) {
	d.Set("deviceid", m.Deviceid)
	d.Set("workflowid", m.Workflowid)
}

func SetZserviceWorkflowInfoSubResourceData(m []*models.ZserviceWorkflowInfo) (d []*map[string]interface{}) {
	for _, ZserviceWorkflowInfoModel := range m {
		if ZserviceWorkflowInfoModel != nil {
			properties := make(map[string]interface{})
			properties["deviceid"] = ZserviceWorkflowInfoModel.Deviceid
			properties["workflowid"] = ZserviceWorkflowInfoModel.Workflowid
			d = append(d, &properties)
		}
	}
	return
}

func ZserviceWorkflowInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"deviceid": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"workflowid": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetZserviceWorkflowInfoPropertyFields() (t []string) {
	return []string{
		"deviceid",
		"workflowid",
	}
}
