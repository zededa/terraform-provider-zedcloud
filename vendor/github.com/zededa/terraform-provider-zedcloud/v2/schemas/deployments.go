package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func DeploymentsModel(d *schema.ResourceData) *models.Deployments {
	var list []*models.DeploymentConfigSummary // []*DeploymentConfigSummary
	listInterface, listIsSet := d.GetOk("list")
	if listIsSet {
		var items []interface{}
		if listItems, isList := listInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = listInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := DeploymentConfigSummaryModelFromMap(v.(map[string]interface{}))
			list = append(list, m)
		}
	}
	var next *models.Cursor // Cursor
	nextInterface, nextIsSet := d.GetOk("next")
	if nextIsSet && nextInterface != nil {
		nextMap := nextInterface.([]interface{})
		if len(nextMap) > 0 {
			next = CursorModelFromMap(nextMap[0].(map[string]interface{}))
		}
	}
	var summaryByState *models.Summary // Summary
	summaryByStateInterface, summaryByStateIsSet := d.GetOk("summary_by_state")
	if summaryByStateIsSet && summaryByStateInterface != nil {
		summaryByStateMap := summaryByStateInterface.([]interface{})
		if len(summaryByStateMap) > 0 {
			summaryByState = SummaryModelFromMap(summaryByStateMap[0].(map[string]interface{}))
		}
	}
	return &models.Deployments{
		List:           list,
		Next:           next,
		SummaryByState: summaryByState,
	}
}

func DeploymentsModelFromMap(m map[string]interface{}) *models.Deployments {
	var list []*models.DeploymentConfigSummary // []*DeploymentConfigSummary
	listInterface, listIsSet := m["list"]
	if listIsSet {
		var items []interface{}
		if listItems, isList := listInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = listInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := DeploymentConfigSummaryModelFromMap(v.(map[string]interface{}))
			list = append(list, m)
		}
	}
	var next *models.Cursor // Cursor
	nextInterface, nextIsSet := m["next"]
	if nextIsSet && nextInterface != nil {
		nextMap := nextInterface.([]interface{})
		if len(nextMap) > 0 {
			next = CursorModelFromMap(nextMap[0].(map[string]interface{}))
		}
	}
	//
	var summaryByState *models.Summary // Summary
	summaryByStateInterface, summaryByStateIsSet := m["summary_by_state"]
	if summaryByStateIsSet && summaryByStateInterface != nil {
		summaryByStateMap := summaryByStateInterface.([]interface{})
		if len(summaryByStateMap) > 0 {
			summaryByState = SummaryModelFromMap(summaryByStateMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.Deployments{
		List:           list,
		Next:           next,
		SummaryByState: summaryByState,
	}
}

func SetDeploymentsResourceData(d *schema.ResourceData, m *models.Deployments) {
	d.Set("list", SetDeploymentConfigSummarySubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("summary_by_state", SetSummarySubResourceData([]*models.Summary{m.SummaryByState}))
}

func SetDeploymentsSubResourceData(m []*models.Deployments) (d []*map[string]interface{}) {
	for _, DeploymentsModel := range m {
		if DeploymentsModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetDeploymentConfigSummarySubResourceData(DeploymentsModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{DeploymentsModel.Next})
			properties["summary_by_state"] = SetSummarySubResourceData([]*models.Summary{DeploymentsModel.SummaryByState})
			d = append(d, &properties)
		}
	}
	return
}

func DeploymentsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `deployment instance summary list response`,
			Type:        schema.TypeList, //GoType: []*DeploymentConfigSummary
			Elem: &schema.Resource{
				Schema: DeploymentConfigSummarySchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"next": {
			Description: `cursor next`,
			Type:        schema.TypeList, //GoType: Cursor
			Elem: &schema.Resource{
				Schema: CursorSchema(),
			},
			Optional: true,
		},

		"summary_by_state": {
			Description: `deployment instance by summary`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},
	}
}

func GetDeploymentsPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"summary_by_state",
	}
}
