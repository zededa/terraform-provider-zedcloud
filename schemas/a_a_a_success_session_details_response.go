package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAASuccessSessionDetailsResponseModel(d *schema.ResourceData) *models.AAASuccessSessionDetailsResponse {
	var original *models.OpaqueToken64 // OpaqueToken64
	originalInterface, originalIsSet := d.GetOk("original")
	if originalIsSet && originalInterface != nil {
		originalMap := originalInterface.([]interface{})
		if len(originalMap) > 0 {
			original = OpaqueToken64ModelFromMap(originalMap[0].(map[string]interface{}))
		}
	}
	var policies []*models.Policy // []*Policy
	policiesInterface, policiesIsSet := d.GetOk("policies")
	if policiesIsSet {
		var items []interface{}
		if listItems, isList := policiesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = policiesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := PolicyConfigModelFromMap(v.(map[string]interface{}))
			policies = append(policies, m)
		}
	}
	var user *models.DetailedUser // DetailedUser
	userInterface, userIsSet := d.GetOk("user")
	if userIsSet && userInterface != nil {
		userMap := userInterface.([]interface{})
		if len(userMap) > 0 {
			user = DetailedUserModelFromMap(userMap[0].(map[string]interface{}))
		}
	}
	return &models.AAASuccessSessionDetailsResponse{
		Original: original,
		Policies: policies,
		User:     user,
	}
}

func AAASuccessSessionDetailsResponseModelFromMap(m map[string]interface{}) *models.AAASuccessSessionDetailsResponse {
	var original *models.OpaqueToken64 // OpaqueToken64
	originalInterface, originalIsSet := m["original"]
	if originalIsSet && originalInterface != nil {
		originalMap := originalInterface.([]interface{})
		if len(originalMap) > 0 {
			original = OpaqueToken64ModelFromMap(originalMap[0].(map[string]interface{}))
		}
	}
	//
	var policies []*models.Policy // []*Policy
	policiesInterface, policiesIsSet := m["policies"]
	if policiesIsSet {
		var items []interface{}
		if listItems, isList := policiesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = policiesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := PolicyConfigModelFromMap(v.(map[string]interface{}))
			policies = append(policies, m)
		}
	}
	var user *models.DetailedUser // DetailedUser
	userInterface, userIsSet := m["user"]
	if userIsSet && userInterface != nil {
		userMap := userInterface.([]interface{})
		if len(userMap) > 0 {
			user = DetailedUserModelFromMap(userMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.AAASuccessSessionDetailsResponse{
		Original: original,
		Policies: policies,
		User:     user,
	}
}

func SetAAASuccessSessionDetailsResponseResourceData(d *schema.ResourceData, m *models.AAASuccessSessionDetailsResponse) {
	d.Set("original", SetOpaqueToken64SubResourceData([]*models.OpaqueToken64{m.Original}))
	d.Set("policies", SetPolicyConfigSubResourceData(m.Policies))
	d.Set("user", SetDetailedUserSubResourceData([]*models.DetailedUser{m.User}))
}

func SetAAASuccessSessionDetailsResponseSubResourceData(m []*models.AAASuccessSessionDetailsResponse) (d []*map[string]interface{}) {
	for _, AAASuccessSessionDetailsResponseModel := range m {
		if AAASuccessSessionDetailsResponseModel != nil {
			properties := make(map[string]interface{})
			properties["original"] = SetOpaqueToken64SubResourceData([]*models.OpaqueToken64{AAASuccessSessionDetailsResponseModel.Original})
			properties["policies"] = SetPolicyConfigSubResourceData(AAASuccessSessionDetailsResponseModel.Policies)
			properties["user"] = SetDetailedUserSubResourceData([]*models.DetailedUser{AAASuccessSessionDetailsResponseModel.User})
			d = append(d, &properties)
		}
	}
	return
}

func AAASuccessSessionDetailsResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"original": {
			Description: ``,
			Type:        schema.TypeList, //GoType: OpaqueToken64
			Elem: &schema.Resource{
				Schema: OpaqueToken64Schema(),
			},
			Optional: true,
		},

		"policies": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*Policy
			Elem: &schema.Resource{
				Schema: Policy(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"user": {
			Description: ``,
			Type:        schema.TypeList, //GoType: DetailedUser
			Elem: &schema.Resource{
				Schema: DetailedUserSchema(),
			},
			Optional: true,
		},
	}
}

func GetAAASuccessSessionDetailsResponsePropertyFields() (t []string) {
	return []string{
		"original",
		"policies",
		"user",
	}
}
