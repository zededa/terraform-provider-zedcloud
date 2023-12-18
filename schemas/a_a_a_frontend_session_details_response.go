package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AAAFrontendSessionDetailsResponseModel(d *schema.ResourceData) *models.AAAFrontendSessionDetailsResponse {
	var cause *models.AAAFrontendSessionDetailsResponseCause // AAAFrontendSessionDetailsResponseCause
	causeInterface, causeIsSet := d.GetOk("cause")
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewAAAFrontendSessionDetailsResponseCause(models.AAAFrontendSessionDetailsResponseCause(causeModel))
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
	return &models.AAAFrontendSessionDetailsResponse{
		Cause:    cause,
		Policies: policies,
		User:     user,
	}
}

func AAAFrontendSessionDetailsResponseModelFromMap(m map[string]interface{}) *models.AAAFrontendSessionDetailsResponse {
	var cause *models.AAAFrontendSessionDetailsResponseCause // AAAFrontendSessionDetailsResponseCause
	causeInterface, causeIsSet := m["cause"]
	if causeIsSet {
		causeModel := causeInterface.(string)
		cause = models.NewAAAFrontendSessionDetailsResponseCause(models.AAAFrontendSessionDetailsResponseCause(causeModel))
	}
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
	return &models.AAAFrontendSessionDetailsResponse{
		Cause:    cause,
		Policies: policies,
		User:     user,
	}
}

func SetAAAFrontendSessionDetailsResponseResourceData(d *schema.ResourceData, m *models.AAAFrontendSessionDetailsResponse) {
	d.Set("cause", m.Cause)
	d.Set("policies", SetPolicyConfigSubResourceData(m.Policies))
	d.Set("user", SetDetailedUserSubResourceData([]*models.DetailedUser{m.User}))
}

func SetAAAFrontendSessionDetailsResponseSubResourceData(m []*models.AAAFrontendSessionDetailsResponse) (d []*map[string]interface{}) {
	for _, AAAFrontendSessionDetailsResponseModel := range m {
		if AAAFrontendSessionDetailsResponseModel != nil {
			properties := make(map[string]interface{})
			properties["cause"] = AAAFrontendSessionDetailsResponseModel.Cause
			properties["policies"] = SetPolicyConfigSubResourceData(AAAFrontendSessionDetailsResponseModel.Policies)
			properties["user"] = SetDetailedUserSubResourceData([]*models.DetailedUser{AAAFrontendSessionDetailsResponseModel.User})
			d = append(d, &properties)
		}
	}
	return
}

func AAAFrontendSessionDetailsResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cause": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
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

func GetAAAFrontendSessionDetailsResponsePropertyFields() (t []string) {
	return []string{
		"cause",
		"policies",
		"user",
	}
}
