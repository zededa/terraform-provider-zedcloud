package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AppProfileConfigSummaryModel(d *schema.ResourceData) *models.AppProfileConfigSummary {
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
	description, _ := d.Get("description").(string)
	id, _ := d.Get("id").(string)
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
	return &models.AppProfileConfigSummary{
		AppPolicies:     appPolicies,
		Description:     description,
		ID:              id,
		Name:            name,
		NetworkPolicies: networkPolicies,
		Title:           title,
		VolumePolicies:  volumePolicies,
	}
}

func AppProfileConfigSummaryModelFromMap(m map[string]interface{}) *models.AppProfileConfigSummary {
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
	description := m["description"].(string)
	id := m["id"].(string)
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
	return &models.AppProfileConfigSummary{
		AppPolicies:     appPolicies,
		Description:     description,
		ID:              id,
		Name:            name,
		NetworkPolicies: networkPolicies,
		Title:           title,
		VolumePolicies:  volumePolicies,
	}
}

func SetAppProfileConfigSummaryResourceData(d *schema.ResourceData, m *models.AppProfileConfigSummary) {
	d.Set("app_policies", SetProfileAppPolicySubResourceData(m.AppPolicies))
	d.Set("description", m.Description)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("network_policies", SetProfileNetworkPolicySubResourceData(m.NetworkPolicies))
	d.Set("title", m.Title)
	d.Set("volume_policies", SetProfileVolumePolicySubResourceData(m.VolumePolicies))
}

func SetAppProfileConfigSummarySubResourceData(m []*models.AppProfileConfigSummary) (d []*map[string]interface{}) {
	for _, AppProfileConfigSummaryModel := range m {
		if AppProfileConfigSummaryModel != nil {
			properties := make(map[string]interface{})
			properties["app_policies"] = SetProfileAppPolicySubResourceData(AppProfileConfigSummaryModel.AppPolicies)
			properties["description"] = AppProfileConfigSummaryModel.Description
			properties["id"] = AppProfileConfigSummaryModel.ID
			properties["name"] = AppProfileConfigSummaryModel.Name
			properties["network_policies"] = SetProfileNetworkPolicySubResourceData(AppProfileConfigSummaryModel.NetworkPolicies)
			properties["title"] = AppProfileConfigSummaryModel.Title
			properties["volume_policies"] = SetProfileVolumePolicySubResourceData(AppProfileConfigSummaryModel.VolumePolicies)
			d = append(d, &properties)
		}
	}
	return
}

func AppProfileConfigSummarySchema() map[string]*schema.Schema {
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

		"description": {
			Description: `Detailed description of the app profile.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: `unique id for app profile`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: `user defined app profile name`,
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

		"title": {
			Description: `user defined app profile title`,
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

func GetAppProfileConfigSummaryPropertyFields() (t []string) {
	return []string{
		"app_policies",
		"description",
		"id",
		"name",
		"network_policies",
		"title",
		"volume_policies",
	}
}
