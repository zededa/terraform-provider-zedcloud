package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ProfileDeploymentFilterModel(d *schema.ResourceData) *models.ProfileDeploymentFilter {
	namePattern, _ := d.Get("name_pattern").(string)
	profileDeploymentID, _ := d.Get("profile_deployment_id").(string)
	return &models.ProfileDeploymentFilter{
		NamePattern:         namePattern,
		ProfileDeploymentID: profileDeploymentID,
	}
}

func ProfileDeploymentFilterModelFromMap(m map[string]interface{}) *models.ProfileDeploymentFilter {
	namePattern := m["name_pattern"].(string)
	profileDeploymentID := m["profile_deployment_id"].(string)
	return &models.ProfileDeploymentFilter{
		NamePattern:         namePattern,
		ProfileDeploymentID: profileDeploymentID,
	}
}

func SetProfileDeploymentFilterResourceData(d *schema.ResourceData, m *models.ProfileDeploymentFilter) {
	d.Set("name_pattern", m.NamePattern)
	d.Set("profile_deployment_id", m.ProfileDeploymentID)
}

func SetProfileDeploymentFilterSubResourceData(m []*models.ProfileDeploymentFilter) (d []*map[string]interface{}) {
	for _, ProfileDeploymentFilterModel := range m {
		if ProfileDeploymentFilterModel != nil {
			properties := make(map[string]interface{})
			properties["name_pattern"] = ProfileDeploymentFilterModel.NamePattern
			properties["profile_deployment_id"] = ProfileDeploymentFilterModel.ProfileDeploymentID
			d = append(d, &properties)
		}
	}
	return
}

func ProfileDeploymentFilterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name_pattern": {
			Description: `deployment name pattern`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"profile_deployment_id": {
			Description: `Unique id of the asset, unique across the enterprise.`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetProfileDeploymentFilterPropertyFields() (t []string) {
	return []string{
		"name_pattern",
		"profile_deployment_id",
	}
}
