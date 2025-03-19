package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ProfileDeploymentResourceDetailedStatusRespModel(d *schema.ResourceData) *models.ProfileDeploymentResourceDetailedStatusResp {
	var next *models.Cursor // Cursor
	nextInterface, nextIsSet := d.GetOk("next")
	if nextIsSet && nextInterface != nil {
		nextMap := nextInterface.([]interface{})
		if len(nextMap) > 0 {
			next = CursorModelFromMap(nextMap[0].(map[string]interface{}))
		}
	}
	var resourceStatus *models.ProfileDeploymentResourceDetailedStatus // ProfileDeploymentResourceDetailedStatus
	resourceStatusInterface, resourceStatusIsSet := d.GetOk("resource_status")
	if resourceStatusIsSet && resourceStatusInterface != nil {
		resourceStatusMap := resourceStatusInterface.([]interface{})
		if len(resourceStatusMap) > 0 {
			resourceStatus = ProfileDeploymentResourceDetailedStatusModelFromMap(resourceStatusMap[0].(map[string]interface{}))
		}
	}
	return &models.ProfileDeploymentResourceDetailedStatusResp{
		Next:           next,
		ResourceStatus: resourceStatus,
	}
}

func ProfileDeploymentResourceDetailedStatusRespModelFromMap(m map[string]interface{}) *models.ProfileDeploymentResourceDetailedStatusResp {
	var next *models.Cursor // Cursor
	nextInterface, nextIsSet := m["next"]
	if nextIsSet && nextInterface != nil {
		nextMap := nextInterface.([]interface{})
		if len(nextMap) > 0 {
			next = CursorModelFromMap(nextMap[0].(map[string]interface{}))
		}
	}
	//
	var resourceStatus *models.ProfileDeploymentResourceDetailedStatus // ProfileDeploymentResourceDetailedStatus
	resourceStatusInterface, resourceStatusIsSet := m["resource_status"]
	if resourceStatusIsSet && resourceStatusInterface != nil {
		resourceStatusMap := resourceStatusInterface.([]interface{})
		if len(resourceStatusMap) > 0 {
			resourceStatus = ProfileDeploymentResourceDetailedStatusModelFromMap(resourceStatusMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.ProfileDeploymentResourceDetailedStatusResp{
		Next:           next,
		ResourceStatus: resourceStatus,
	}
}

func SetProfileDeploymentResourceDetailedStatusRespResourceData(d *schema.ResourceData, m *models.ProfileDeploymentResourceDetailedStatusResp) {
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("resource_status", SetProfileDeploymentResourceDetailedStatusSubResourceData([]*models.ProfileDeploymentResourceDetailedStatus{m.ResourceStatus}))
}

func SetProfileDeploymentResourceDetailedStatusRespSubResourceData(m []*models.ProfileDeploymentResourceDetailedStatusResp) (d []*map[string]interface{}) {
	for _, ProfileDeploymentResourceDetailedStatusRespModel := range m {
		if ProfileDeploymentResourceDetailedStatusRespModel != nil {
			properties := make(map[string]interface{})
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{ProfileDeploymentResourceDetailedStatusRespModel.Next})
			properties["resource_status"] = SetProfileDeploymentResourceDetailedStatusSubResourceData([]*models.ProfileDeploymentResourceDetailedStatus{ProfileDeploymentResourceDetailedStatusRespModel.ResourceStatus})
			d = append(d, &properties)
		}
	}
	return
}

func ProfileDeploymentResourceDetailedStatusRespSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"next": {
			Description: `Requested page details of filtered records`,
			Type:        schema.TypeList, //GoType: Cursor
			Elem: &schema.Resource{
				Schema: CursorSchema(),
			},
			Optional: true,
		},

		"resource_status": {
			Description: `List of status of deployment resources`,
			Type:        schema.TypeList, //GoType: ProfileDeploymentResourceDetailedStatus
			Elem: &schema.Resource{
				Schema: ProfileDeploymentResourceDetailedStatusSchema(),
			},
			Optional: true,
		},
	}
}

func GetProfileDeploymentResourceDetailedStatusRespPropertyFields() (t []string) {
	return []string{
		"next",
		"resource_status",
	}
}
