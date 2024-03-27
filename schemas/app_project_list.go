package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func AppProjectListModel(d *schema.ResourceData) *models.AppProjectList {
	var next *models.Cursor // Cursor
	nextInterface, nextIsSet := d.GetOk("next")
	if nextIsSet && nextInterface != nil {
		nextMap := nextInterface.([]interface{})
		if len(nextMap) > 0 {
			next = CursorModelFromMap(nextMap[0].(map[string]interface{}))
		}
	}
	var projectDetails []*models.ObjectDetail // []*ObjectDetail
	projectDetailsInterface, projectDetailsIsSet := d.GetOk("project_details")
	if projectDetailsIsSet {
		var items []interface{}
		if listItems, isList := projectDetailsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = projectDetailsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ObjectDetailModelFromMap(v.(map[string]interface{}))
			projectDetails = append(projectDetails, m)
		}
	}
	totalCountInt, _ := d.Get("total_count").(int)
	totalCount := int64(totalCountInt)
	return &models.AppProjectList{
		Next:           next,
		ProjectDetails: projectDetails,
		TotalCount:     totalCount,
	}
}

func AppProjectListModelFromMap(m map[string]interface{}) *models.AppProjectList {
	var next *models.Cursor // Cursor
	nextInterface, nextIsSet := m["next"]
	if nextIsSet && nextInterface != nil {
		nextMap := nextInterface.([]interface{})
		if len(nextMap) > 0 {
			next = CursorModelFromMap(nextMap[0].(map[string]interface{}))
		}
	}
	//
	var projectDetails []*models.ObjectDetail // []*ObjectDetail
	projectDetailsInterface, projectDetailsIsSet := m["project_details"]
	if projectDetailsIsSet {
		var items []interface{}
		if listItems, isList := projectDetailsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = projectDetailsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ObjectDetailModelFromMap(v.(map[string]interface{}))
			projectDetails = append(projectDetails, m)
		}
	}
	totalCount := int64(m["total_count"].(int)) // int64
	return &models.AppProjectList{
		Next:           next,
		ProjectDetails: projectDetails,
		TotalCount:     totalCount,
	}
}

// Update the underlying AppProjectList resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppProjectListResourceData(d *schema.ResourceData, m *models.AppProjectList) {
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("project_details", SetObjectDetailSubResourceData(m.ProjectDetails))
	d.Set("total_count", m.TotalCount)
}

// Iterate through and update the AppProjectList resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppProjectListSubResourceData(m []*models.AppProjectList) (d []*map[string]interface{}) {
	for _, AppProjectListModel := range m {
		if AppProjectListModel != nil {
			properties := make(map[string]interface{})
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{AppProjectListModel.Next})
			properties["project_details"] = SetObjectDetailSubResourceData(AppProjectListModel.ProjectDetails)
			properties["total_count"] = AppProjectListModel.TotalCount
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AppProjectList resource defined in the Terraform configuration
func AppProjectListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"next": {
			Description: `Responded page details of filtered records`,
			Type:        schema.TypeList, //GoType: Cursor
			Elem: &schema.Resource{
				Schema: CursorSchema(),
			},
			Optional: true,
		},

		"project_details": {
			Description: `List of projects which are associated with the given app`,
			Type:        schema.TypeList, //GoType: []*ObjectDetail
			Elem: &schema.Resource{
				Schema: ObjectDetailSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"total_count": {
			Description: `total number of records`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the AppProjectList resource
func GetAppProjectListPropertyFields() (t []string) {
	return []string{
		"next",
		"project_details",
		"total_count",
	}
}
