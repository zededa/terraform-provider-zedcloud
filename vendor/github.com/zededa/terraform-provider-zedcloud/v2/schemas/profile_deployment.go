package schemas

import (
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ProfileDeploymentModel(d *schema.ResourceData) *models.ProfileDeployment {
	var appProfileInfo *models.AppProfileInfo // AppProfileInfo
	appProfileInfoInterface, appProfileInfoIsSet := d.GetOk("app_profile_info")
	if appProfileInfoIsSet && appProfileInfoInterface != nil {
		appProfileInfoMap := appProfileInfoInterface.([]interface{})
		if len(appProfileInfoMap) > 0 {
			appProfileInfo = AppProfileInfoModelFromMap(appProfileInfoMap[0].(map[string]interface{}))
		}
	}
	description, _ := d.Get("description").(string)
	id, _ := d.Get("id").(string)
	if id == "" {
		id, _ = uuid.GenerateUUID()
	}
	name, _ := d.Get("name").(string)
	projectID, _ := d.Get("project_id").(string)
	status, _ := d.Get("status").(string)
	var targetAssetGroup *models.TargetAssetGroup // TargetAssetGroup
	targetAssetGroupInterface, targetAssetGroupIsSet := d.GetOk("target_asset_group")
	if targetAssetGroupIsSet && targetAssetGroupInterface != nil {
		targetAssetGroupMap := targetAssetGroupInterface.([]interface{})
		if len(targetAssetGroupMap) > 0 {
			targetAssetGroup = TargetAssetGroupModelFromMap(targetAssetGroupMap[0].(map[string]interface{}))
		}
	}
	title, _ := d.Get("title").(string)
	return &models.ProfileDeployment{
		AppProfileInfo:   appProfileInfo,
		Description:      description,
		ID:               id,
		Name:             &name,      // string
		ProjectID:        &projectID, // string
		Status:           status,
		TargetAssetGroup: targetAssetGroup,
		Title:            title,
	}
}

func ProfileDeploymentModelFromMap(m map[string]interface{}) *models.ProfileDeployment {
	var appProfileInfo *models.AppProfileInfo // AppProfileInfo
	appProfileInfoInterface, appProfileInfoIsSet := m["app_profile_info"]
	if appProfileInfoIsSet && appProfileInfoInterface != nil {
		appProfileInfoMap := appProfileInfoInterface.([]interface{})
		if len(appProfileInfoMap) > 0 {
			appProfileInfo = AppProfileInfoModelFromMap(appProfileInfoMap[0].(map[string]interface{}))
		}
	}
	//
	description := m["description"].(string)
	id := m["id"].(string)
	name := m["name"].(string)
	projectID := m["project_id"].(string)
	status := m["status"].(string)
	var targetAssetGroup *models.TargetAssetGroup // TargetAssetGroup
	targetAssetGroupInterface, targetAssetGroupIsSet := m["target_asset_group"]
	if targetAssetGroupIsSet && targetAssetGroupInterface != nil {
		targetAssetGroupMap := targetAssetGroupInterface.([]interface{})
		if len(targetAssetGroupMap) > 0 {
			targetAssetGroup = TargetAssetGroupModelFromMap(targetAssetGroupMap[0].(map[string]interface{}))
		}
	}
	//
	title := m["title"].(string)
	return &models.ProfileDeployment{
		AppProfileInfo:   appProfileInfo,
		Description:      description,
		ID:               id,
		Name:             &name,
		ProjectID:        &projectID,
		Status:           status,
		TargetAssetGroup: targetAssetGroup,
		Title:            title,
	}
}

func SetProfileDeploymentResourceData(d *schema.ResourceData, m *models.ProfileDeployment) {
	d.Set("app_profile_info", SetAppProfileInfoSubResourceData([]*models.AppProfileInfo{m.AppProfileInfo}))
	d.Set("description", m.Description)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("project_id", m.ProjectID)
	d.Set("status", m.Status)
	d.Set("target_asset_group", SetTargetAssetGroupSubResourceData([]*models.TargetAssetGroup{m.TargetAssetGroup}))
	d.Set("title", m.Title)
}

func SetProfileDeploymentSubResourceData(m []*models.ProfileDeployment) (d []*map[string]interface{}) {
	for _, ProfileDeploymentModel := range m {
		if ProfileDeploymentModel != nil {
			properties := make(map[string]interface{})
			properties["app_profile_info"] = SetAppProfileInfoSubResourceData([]*models.AppProfileInfo{ProfileDeploymentModel.AppProfileInfo})
			properties["description"] = ProfileDeploymentModel.Description
			properties["id"] = ProfileDeploymentModel.ID
			properties["name"] = ProfileDeploymentModel.Name
			properties["project_id"] = ProfileDeploymentModel.ProjectID
			properties["status"] = ProfileDeploymentModel.Status
			properties["target_asset_group"] = SetTargetAssetGroupSubResourceData([]*models.TargetAssetGroup{ProfileDeploymentModel.TargetAssetGroup})
			properties["title"] = ProfileDeploymentModel.Title
			d = append(d, &properties)
		}
	}
	return
}

func ProfileDeploymentSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_profile_info": {
			Description: `app profile id`,
			Type:        schema.TypeList, //GoType: AppProfileInfo
			Elem: &schema.Resource{
				Schema: AppProfileInfoSchema(),
			},
			Required: true,
		},

		"description": {
			Description: `Detailed description of the deployment.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: `unique Id of the deployment.`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: `User defined name of the deployment, unique across the enterprise. Once deployment is created, name can’t be changed.`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"project_id": {
			Description: `project id`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"status": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"target_asset_group": {
			Description: `asset group info`,
			Type:        schema.TypeList, //GoType: TargetAssetGroup
			Elem: &schema.Resource{
				Schema: TargetAssetGroupSchema(),
			},
			Required: true,
		},

		"title": {
			Description: `User defined title of the deployment. Title can be changed at any time.`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetProfileDeploymentPropertyFields() (t []string) {
	return []string{
		"app_profile_info",
		"description",
		"id",
		"name",
		"project_id",
		"status",
		"target_asset_group",
		"title",
	}
}
