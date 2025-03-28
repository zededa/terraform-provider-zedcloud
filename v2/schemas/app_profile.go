package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AppProfileModel(d *schema.ResourceData) *models.AppProfile {
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
	return &models.AppProfile{
		AppPolicies:     appPolicies,
		ID:              id,
		Name:            name,
		NetworkPolicies: networkPolicies,
		Title:           title,
		VolumePolicies:  volumePolicies,
	}
}

func AppProfileModelFromMap(m map[string]interface{}) *models.AppProfile {
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
	return &models.AppProfile{
		AppPolicies:     appPolicies,
		ID:              id,
		Name:            name,
		NetworkPolicies: networkPolicies,
		Title:           title,
		VolumePolicies:  volumePolicies,
	}
}

func SetAppProfileResourceData(d *schema.ResourceData, m *models.AppProfile) {
	d.Set("app_policies", SetProfileAppPolicySubResourceData(m.AppPolicies))
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("network_policies", SetProfileNetworkPolicySubResourceData(m.NetworkPolicies))
	d.Set("title", m.Title)
	d.Set("volume_policies", SetProfileVolumePolicySubResourceData(m.VolumePolicies))
}

func SetAppProfileSubResourceData(m []*models.AppProfile) (d []*map[string]interface{}) {
	for _, AppProfileModel := range m {
		if AppProfileModel != nil {
			properties := make(map[string]interface{})
			properties["app_policies"] = SetProfileAppPolicySubResourceData(AppProfileModel.AppPolicies)
			properties["id"] = AppProfileModel.ID
			properties["name"] = AppProfileModel.Name
			properties["network_policies"] = SetProfileNetworkPolicySubResourceData(AppProfileModel.NetworkPolicies)
			properties["title"] = AppProfileModel.Title
			properties["volume_policies"] = SetProfileVolumePolicySubResourceData(AppProfileModel.VolumePolicies)
			d = append(d, &properties)
		}
	}
	return
}

func AppProfileSchema() map[string]*schema.Schema {
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

func GetAppProfilePropertyFields() (t []string) {
	return []string{
		"app_policies",
		"id",
		"name",
		"network_policies",
		"title",
		"volume_policies",
	}
}
