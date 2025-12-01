package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZksInstanceStatusSummariesModel(d *schema.ResourceData) *models.ZksInstanceStatusSummaries {
	var statuses *models.ZksInstanceStatusSummary // ZksInstanceStatusSummary
	statusesInterface, statusesIsSet := d.GetOk("statuses")
	if statusesIsSet && statusesInterface != nil {
		statusesMap := statusesInterface.([]interface{})
		if len(statusesMap) > 0 {
			statuses = ZksInstanceStatusSummaryModelFromMap(statusesMap[0].(map[string]interface{}))
		}
	}
	return &models.ZksInstanceStatusSummaries{
		Statuses: statuses,
	}
}

func ZksInstanceStatusSummariesModelFromMap(m map[string]interface{}) *models.ZksInstanceStatusSummaries {
	var statuses *models.ZksInstanceStatusSummary // ZksInstanceStatusSummary
	statusesInterface, statusesIsSet := m["statuses"]
	if statusesIsSet && statusesInterface != nil {
		statusesMap := statusesInterface.([]interface{})
		if len(statusesMap) > 0 {
			statuses = ZksInstanceStatusSummaryModelFromMap(statusesMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.ZksInstanceStatusSummaries{
		Statuses: statuses,
	}
}

func SetZksInstanceStatusSummariesResourceData(d *schema.ResourceData, m *models.ZksInstanceStatusSummaries) {
	d.Set("statuses", SetZksInstanceStatusSummarySubResourceData([]*models.ZksInstanceStatusSummary{m.Statuses}))
}

func SetZksInstanceStatusSummariesSubResourceData(m []*models.ZksInstanceStatusSummaries) (d []*map[string]interface{}) {
	for _, ZksInstanceStatusSummariesModel := range m {
		if ZksInstanceStatusSummariesModel != nil {
			properties := make(map[string]interface{})
			properties["statuses"] = SetZksInstanceStatusSummarySubResourceData([]*models.ZksInstanceStatusSummary{ZksInstanceStatusSummariesModel.Statuses})
			d = append(d, &properties)
		}
	}
	return
}

func ZksInstanceStatusSummariesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"statuses": {
			Description: `Summary of the zks instance status`,
			Type:        schema.TypeList, //GoType: ZksInstanceStatusSummary
			Elem: &schema.Resource{
				Schema: ZksInstanceStatusSummarySchema(),
			},
			Optional: true,
		},
	}
}

func GetZksInstanceStatusSummariesPropertyFields() (t []string) {
	return []string{
		"statuses",
	}
}
