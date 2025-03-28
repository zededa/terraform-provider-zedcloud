package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ProfileDeploymentResourceDetailedStatusModel(d *schema.ResourceData) *models.ProfileDeploymentResourceDetailedStatus {
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
	name, _ := d.Get("name").(string)
	var objectStatus []*models.ProfileDeploymentResourceObjectStatus // []*ProfileDeploymentResourceObjectStatus
	objectStatusInterface, objectStatusIsSet := d.GetOk("object_status")
	if objectStatusIsSet {
		var items []interface{}
		if listItems, isList := objectStatusInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = objectStatusInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ProfileDeploymentResourceObjectStatusModelFromMap(v.(map[string]interface{}))
			objectStatus = append(objectStatus, m)
		}
	}
	projectID, _ := d.Get("project_id").(string)
	var targetAssetGroup *models.TargetAssetGroup // TargetAssetGroup
	targetAssetGroupInterface, targetAssetGroupIsSet := d.GetOk("target_asset_group")
	if targetAssetGroupIsSet && targetAssetGroupInterface != nil {
		targetAssetGroupMap := targetAssetGroupInterface.([]interface{})
		if len(targetAssetGroupMap) > 0 {
			targetAssetGroup = TargetAssetGroupModelFromMap(targetAssetGroupMap[0].(map[string]interface{}))
		}
	}
	title, _ := d.Get("title").(string)
	return &models.ProfileDeploymentResourceDetailedStatus{
		AppProfileInfo:   appProfileInfo,
		Description:      description,
		ID:               id,
		Name:             name,
		ObjectStatus:     objectStatus,
		ProjectID:        projectID,
		TargetAssetGroup: targetAssetGroup,
		Title:            title,
	}
}

func ProfileDeploymentResourceDetailedStatusModelFromMap(m map[string]interface{}) *models.ProfileDeploymentResourceDetailedStatus {
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
	var objectStatus []*models.ProfileDeploymentResourceObjectStatus // []*ProfileDeploymentResourceObjectStatus
	objectStatusInterface, objectStatusIsSet := m["object_status"]
	if objectStatusIsSet {
		var items []interface{}
		if listItems, isList := objectStatusInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = objectStatusInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ProfileDeploymentResourceObjectStatusModelFromMap(v.(map[string]interface{}))
			objectStatus = append(objectStatus, m)
		}
	}
	projectID := m["project_id"].(string)
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
	return &models.ProfileDeploymentResourceDetailedStatus{
		AppProfileInfo:   appProfileInfo,
		Description:      description,
		ID:               id,
		Name:             name,
		ObjectStatus:     objectStatus,
		ProjectID:        projectID,
		TargetAssetGroup: targetAssetGroup,
		Title:            title,
	}
}

func SetProfileDeploymentResourceDetailedStatusResourceData(d *schema.ResourceData, m *models.ProfileDeploymentResourceDetailedStatus) {
	d.Set("app_profile_info", SetAppProfileInfoSubResourceData([]*models.AppProfileInfo{m.AppProfileInfo}))
	d.Set("description", m.Description)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("object_status", SetProfileDeploymentResourceObjectStatusSubResourceData(m.ObjectStatus))
	d.Set("project_id", m.ProjectID)
	d.Set("target_asset_group", SetTargetAssetGroupSubResourceData([]*models.TargetAssetGroup{m.TargetAssetGroup}))
	d.Set("title", m.Title)
}

func SetProfileDeploymentResourceDetailedStatusSubResourceData(m []*models.ProfileDeploymentResourceDetailedStatus) (d []*map[string]interface{}) {
	for _, ProfileDeploymentResourceDetailedStatusModel := range m {
		if ProfileDeploymentResourceDetailedStatusModel != nil {
			properties := make(map[string]interface{})
			properties["app_profile_info"] = SetAppProfileInfoSubResourceData([]*models.AppProfileInfo{ProfileDeploymentResourceDetailedStatusModel.AppProfileInfo})
			properties["description"] = ProfileDeploymentResourceDetailedStatusModel.Description
			properties["id"] = ProfileDeploymentResourceDetailedStatusModel.ID
			properties["name"] = ProfileDeploymentResourceDetailedStatusModel.Name
			properties["object_status"] = SetProfileDeploymentResourceObjectStatusSubResourceData(ProfileDeploymentResourceDetailedStatusModel.ObjectStatus)
			properties["project_id"] = ProfileDeploymentResourceDetailedStatusModel.ProjectID
			properties["target_asset_group"] = SetTargetAssetGroupSubResourceData([]*models.TargetAssetGroup{ProfileDeploymentResourceDetailedStatusModel.TargetAssetGroup})
			properties["title"] = ProfileDeploymentResourceDetailedStatusModel.Title
			d = append(d, &properties)
		}
	}
	return
}

func ProfileDeploymentResourceDetailedStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_profile_info": {
			Description: `app profile id`,
			Type:        schema.TypeList, //GoType: AppProfileInfo
			Elem: &schema.Resource{
				Schema: AppProfileInfoSchema(),
			},
			Optional: true,
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
			Optional:    true,
		},

		"object_status": {
			Description: `list of status of objects created by deployment`,
			Type:        schema.TypeList, //GoType: []*ProfileDeploymentResourceObjectStatus
			Elem: &schema.Resource{
				Schema: ProfileDeploymentResourceObjectStatusSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"project_id": {
			Description: `project id`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"target_asset_group": {
			Description: `asset group info`,
			Type:        schema.TypeList, //GoType: TargetAssetGroup
			Elem: &schema.Resource{
				Schema: TargetAssetGroupSchema(),
			},
			Optional: true,
		},

		"title": {
			Description: `User defined title of the deployment. Title can be changed at any time.`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetProfileDeploymentResourceDetailedStatusPropertyFields() (t []string) {
	return []string{
		"app_profile_info",
		"description",
		"id",
		"name",
		"object_status",
		"project_id",
		"target_asset_group",
		"title",
	}
}
