package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AppProfileReadModel(d *schema.ResourceData) *models.AppProfileRead {
	var appPolicies []*models.ProfileAppPolicy // []*ProfileAppPolicy
	appPoliciesInterface, appPoliciesIsSet := d.GetOk("app_policies")
	if appPoliciesIsSet {
		var items []interface{}
		if listItems, isList := appPoliciesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = appPoliciesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ProfileAppPolicyModelFromMap(v.(map[string]interface{}))
			appPolicies = append(appPolicies, m)
		}
	}
	id, _ := d.Get("id").(string)
	latestVersionInt, _ := d.Get("latest_version").(int)
	latestVersion := int64(latestVersionInt)
	name, _ := d.Get("name").(string)
	var networkPolicies []*models.ProfileNetworkPolicy // []*ProfileNetworkPolicy
	networkPoliciesInterface, networkPoliciesIsSet := d.GetOk("network_policies")
	if networkPoliciesIsSet {
		var items []interface{}
		if listItems, isList := networkPoliciesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = networkPoliciesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ProfileNetworkPolicyModelFromMap(v.(map[string]interface{}))
			networkPolicies = append(networkPolicies, m)
		}
	}
	var status *models.AppProfileStatusType // AppProfileStatusType
	statusInterface, statusIsSet := d.GetOk("status")
	if statusIsSet {
		statusModel := statusInterface.(string)
		status = models.NewAppProfileStatusType(models.AppProfileStatusType(statusModel))
	}
	title, _ := d.Get("title").(string)
	var volumePolicies []*models.ProfileVolumePolicy // []*ProfileVolumePolicy
	volumePoliciesInterface, volumePoliciesIsSet := d.GetOk("volume_policies")
	if volumePoliciesIsSet {
		var items []interface{}
		if listItems, isList := volumePoliciesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = volumePoliciesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ProfileVolumePolicyModelFromMap(v.(map[string]interface{}))
			volumePolicies = append(volumePolicies, m)
		}
	}
	return &models.AppProfileRead{
		AppPolicies:     appPolicies,
		ID:              id,
		LatestVersion:   latestVersion,
		Name:            name,
		NetworkPolicies: networkPolicies,
		Status:          status,
		Title:           title,
		VolumePolicies:  volumePolicies,
	}
}

func AppProfileReadModelFromMap(m map[string]interface{}) *models.AppProfileRead {
	var appPolicies []*models.ProfileAppPolicy // []*ProfileAppPolicy
	appPoliciesInterface, appPoliciesIsSet := m["app_policies"]
	if appPoliciesIsSet {
		var items []interface{}
		if listItems, isList := appPoliciesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = appPoliciesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ProfileAppPolicyModelFromMap(v.(map[string]interface{}))
			appPolicies = append(appPolicies, m)
		}
	}
	id := m["id"].(string)
	latestVersion := int64(m["latest_version"].(int)) // int64
	name := m["name"].(string)
	var networkPolicies []*models.ProfileNetworkPolicy // []*ProfileNetworkPolicy
	networkPoliciesInterface, networkPoliciesIsSet := m["network_policies"]
	if networkPoliciesIsSet {
		var items []interface{}
		if listItems, isList := networkPoliciesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = networkPoliciesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ProfileNetworkPolicyModelFromMap(v.(map[string]interface{}))
			networkPolicies = append(networkPolicies, m)
		}
	}
	var status *models.AppProfileStatusType // AppProfileStatusType
	statusInterface, statusIsSet := m["status"]
	if statusIsSet {
		statusModel := statusInterface.(string)
		status = models.NewAppProfileStatusType(models.AppProfileStatusType(statusModel))
	}
	title := m["title"].(string)
	var volumePolicies []*models.ProfileVolumePolicy // []*ProfileVolumePolicy
	volumePoliciesInterface, volumePoliciesIsSet := m["volume_policies"]
	if volumePoliciesIsSet {
		var items []interface{}
		if listItems, isList := volumePoliciesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = volumePoliciesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ProfileVolumePolicyModelFromMap(v.(map[string]interface{}))
			volumePolicies = append(volumePolicies, m)
		}
	}
	return &models.AppProfileRead{
		AppPolicies:     appPolicies,
		ID:              id,
		LatestVersion:   latestVersion,
		Name:            name,
		NetworkPolicies: networkPolicies,
		Status:          status,
		Title:           title,
		VolumePolicies:  volumePolicies,
	}
}

func SetAppProfileReadResourceData(d *schema.ResourceData, m *models.AppProfileRead) {
	d.Set("app_policies", SetProfileAppPolicySubResourceData(m.AppPolicies))
	d.Set("id", m.ID)
	d.Set("latest_version", m.LatestVersion)
	d.Set("name", m.Name)
	d.Set("network_policies", SetProfileNetworkPolicySubResourceData(m.NetworkPolicies))
	d.Set("status", m.Status)
	d.Set("title", m.Title)
	d.Set("volume_policies", SetProfileVolumePolicySubResourceData(m.VolumePolicies))
}

func SetAppProfileReadSubResourceData(m []*models.AppProfileRead) (d []*map[string]interface{}) {
	for _, AppProfileReadModel := range m {
		if AppProfileReadModel != nil {
			properties := make(map[string]interface{})
			properties["app_policies"] = SetProfileAppPolicySubResourceData(AppProfileReadModel.AppPolicies)
			properties["id"] = AppProfileReadModel.ID
			properties["latest_version"] = AppProfileReadModel.LatestVersion
			properties["name"] = AppProfileReadModel.Name
			properties["network_policies"] = SetProfileNetworkPolicySubResourceData(AppProfileReadModel.NetworkPolicies)
			properties["status"] = AppProfileReadModel.Status
			properties["title"] = AppProfileReadModel.Title
			properties["volume_policies"] = SetProfileVolumePolicySubResourceData(AppProfileReadModel.VolumePolicies)
			d = append(d, &properties)
		}
	}
	return
}

func AppProfileReadSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_policies": {
			Description: `list of app instance policies`,
			Type:        schema.TypeList, //GoType: []*ProfileAppPolicy
			Elem: &schema.Resource{
				Schema: ProfileAppPolicySchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"id": {
			Description: `unique id for an App profile`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"latest_version": {
			Description: `latest version of app profile`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"name": {
			Description: `user defined name for the App profile`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"network_policies": {
			Description: `list of network instance policies`,
			Type:        schema.TypeList, //GoType: []*ProfileNetworkPolicy
			Elem: &schema.Resource{
				Schema: ProfileNetworkPolicySchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"status": {
			Description: `status of the app profile`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"title": {
			Description: `user defined title for the app profile`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"volume_policies": {
			Description: `list of volume instance policies`,
			Type:        schema.TypeList, //GoType: []*ProfileVolumePolicy
			Elem: &schema.Resource{
				Schema: ProfileVolumePolicySchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
	}
}

func GetAppProfileReadPropertyFields() (t []string) {
	return []string{
		"app_policies",
		"id",
		"latest_version",
		"name",
		"network_policies",
		"status",
		"title",
		"volume_policies",
	}
}
