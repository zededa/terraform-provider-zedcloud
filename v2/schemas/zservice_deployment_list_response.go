package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZserviceDeploymentListResponseModel(d *schema.ResourceData) *models.ZserviceDeploymentListResponse {
	var deployments []*models.ZserviceDeploymentReadRO // []*ZserviceDeploymentReadRO
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
			m := ZserviceDeploymentReadROModelFromMap(v.(map[string]interface{}))
			deployments = append(deployments, m)
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
	totalCountInt, _ := d.Get("total_count").(int)
	totalCount := int64(totalCountInt)
	return &models.ZserviceDeploymentListResponse{
		Deployments: deployments,
		Next:        next,
		TotalCount:  totalCount,
	}
}

func ZserviceDeploymentListResponseModelFromMap(m map[string]interface{}) *models.ZserviceDeploymentListResponse {
	var deployments []*models.ZserviceDeploymentReadRO // []*ZserviceDeploymentReadRO
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
			m := ZserviceDeploymentReadROModelFromMap(v.(map[string]interface{}))
			deployments = append(deployments, m)
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
	totalCount := int64(m["total_count"].(int)) // int64
	return &models.ZserviceDeploymentListResponse{
		Deployments: deployments,
		Next:        next,
		TotalCount:  totalCount,
	}
}

func SetZserviceDeploymentListResponseResourceData(d *schema.ResourceData, m *models.ZserviceDeploymentListResponse) {
	d.Set("deployments", SetZserviceDeploymentReadROSubResourceData(m.Deployments))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("total_count", m.TotalCount)
}

func SetZserviceDeploymentListResponseSubResourceData(m []*models.ZserviceDeploymentListResponse) (d []*map[string]interface{}) {
	for _, ZserviceDeploymentListResponseModel := range m {
		if ZserviceDeploymentListResponseModel != nil {
			properties := make(map[string]interface{})
			properties["deployments"] = SetZserviceDeploymentReadROSubResourceData(ZserviceDeploymentListResponseModel.Deployments)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{ZserviceDeploymentListResponseModel.Next})
			properties["total_count"] = ZserviceDeploymentListResponseModel.TotalCount
			d = append(d, &properties)
		}
	}
	return
}

func ZserviceDeploymentListResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"deployments": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*ZserviceDeploymentReadRO
			Elem: &schema.Resource{
				Schema: ZserviceDeploymentReadROSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"next": {
			Description: `Cursor`,
			Type:        schema.TypeList, //GoType: Cursor
			Elem: &schema.Resource{
				Schema: CursorSchema(),
			},
			Optional: true,
		},

		"total_count": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetZserviceDeploymentListResponsePropertyFields() (t []string) {
	return []string{
		"deployments",
		"next",
		"total_count",
	}
}
