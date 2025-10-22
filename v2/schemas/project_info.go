package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ProjectInfoModel(d *schema.ResourceData) *models.ProjectInfo {
	projectName, _ := d.Get("project_name").(string)
	var typeVar *models.TagType // TagType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewTagType(models.TagType(typeModel))
	}
	return &models.ProjectInfo{
		ProjectName: projectName,
		Type:        typeVar,
	}
}

func ProjectInfoModelFromMap(m map[string]interface{}) *models.ProjectInfo {
	projectName := m["project_name"].(string)
	var typeVar *models.TagType // TagType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewTagType(models.TagType(typeModel))
	}
	return &models.ProjectInfo{
		ProjectName: projectName,
		Type:        typeVar,
	}
}

func SetProjectInfoResourceData(d *schema.ResourceData, m *models.ProjectInfo) {
	d.Set("project_id", m.ProjectID)
	d.Set("project_name", m.ProjectName)
	d.Set("type", m.Type)
}

func SetProjectInfoSubResourceData(m []*models.ProjectInfo) (d []*map[string]interface{}) {
	for _, ProjectInfoModel := range m {
		if ProjectInfoModel != nil {
			properties := make(map[string]interface{})
			properties["project_id"] = ProjectInfoModel.ProjectID
			properties["project_name"] = ProjectInfoModel.ProjectName
			properties["type"] = ProjectInfoModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func ProjectInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"project_id": {
			Description: `system generated unique id for a project`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"project_name": {
			Description: `project name`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"type": {
			Description: `Resource group type`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetProjectInfoPropertyFields() (t []string) {
	return []string{
		"project_name",
		"type",
	}
}
