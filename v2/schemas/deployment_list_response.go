package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func DeploymentListResponseModel(d *schema.ResourceData) *models.DeploymentListResponse {
	var deployments []*models.DeploymentGetResponse // []*DeploymentGetResponse
	deploymentsInterface, deploymentsIsSet := d.GetOk("deployments")
	if deploymentsIsSet {
		var items []interface{}
		if listItems, isList := deploymentsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = deploymentsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := DeploymentGetResponseModelFromMap(v.(map[string]interface{}))
			deployments = append(deployments, m)
		}
	}
	var stateSummary *models.Summary // Summary
	stateSummaryInterface, stateSummaryIsSet := d.GetOk("state_summary")
	if stateSummaryIsSet && stateSummaryInterface != nil {
		stateSummaryMap := stateSummaryInterface.([]interface{})
		if len(stateSummaryMap) > 0 {
			stateSummary = SummaryModelFromMap(stateSummaryMap[0].(map[string]interface{}))
		}
	}
	totalCountInt, _ := d.Get("total_count").(int)
	totalCount := int64(totalCountInt)
	return &models.DeploymentListResponse{
		Deployments:  deployments,
		StateSummary: stateSummary,
		TotalCount:   totalCount,
	}
}

func DeploymentListResponseModelFromMap(m map[string]interface{}) *models.DeploymentListResponse {
	var deployments []*models.DeploymentGetResponse // []*DeploymentGetResponse
	deploymentsInterface, deploymentsIsSet := m["deployments"]
	if deploymentsIsSet {
		var items []interface{}
		if listItems, isList := deploymentsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = deploymentsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := DeploymentGetResponseModelFromMap(v.(map[string]interface{}))
			deployments = append(deployments, m)
		}
	}
	var stateSummary *models.Summary // Summary
	stateSummaryInterface, stateSummaryIsSet := m["state_summary"]
	if stateSummaryIsSet && stateSummaryInterface != nil {
		stateSummaryMap := stateSummaryInterface.([]interface{})
		if len(stateSummaryMap) > 0 {
			stateSummary = SummaryModelFromMap(stateSummaryMap[0].(map[string]interface{}))
		}
	}
	//
	totalCount := int64(m["total_count"].(int)) // int64
	return &models.DeploymentListResponse{
		Deployments:  deployments,
		StateSummary: stateSummary,
		TotalCount:   totalCount,
	}
}

func SetDeploymentListResponseResourceData(d *schema.ResourceData, m *models.DeploymentListResponse) {
	d.Set("deployments", SetDeploymentGetResponseSubResourceData(m.Deployments))
	d.Set("state_summary", SetSummarySubResourceData([]*models.Summary{m.StateSummary}))
	d.Set("total_count", m.TotalCount)
}

func SetDeploymentListResponseSubResourceData(m []*models.DeploymentListResponse) (d []*map[string]interface{}) {
	for _, DeploymentListResponseModel := range m {
		if DeploymentListResponseModel != nil {
			properties := make(map[string]interface{})
			properties["deployments"] = SetDeploymentGetResponseSubResourceData(DeploymentListResponseModel.Deployments)
			properties["state_summary"] = SetSummarySubResourceData([]*models.Summary{DeploymentListResponseModel.StateSummary})
			properties["total_count"] = DeploymentListResponseModel.TotalCount
			d = append(d, &properties)
		}
	}
	return
}

func DeploymentListResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"deployments": {
			Description: `List of deployments`,
			Type:        schema.TypeList, //GoType: []*DeploymentGetResponse
			Elem: &schema.Resource{
				Schema: DeploymentGetResponseSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"state_summary": {
			Description: `Summary of deployment states`,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},

		"total_count": {
			Description: `Total number of deployments`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetDeploymentListResponsePropertyFields() (t []string) {
	return []string{
		"deployments",
		"state_summary",
		"total_count",
	}
}
