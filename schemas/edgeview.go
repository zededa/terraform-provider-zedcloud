package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func EdgeViewModel(d *schema.ResourceData) *models.EdgeviewCfg {
	var appPolicy *models.AppAccessPolicy // AppAccessPolicy
	appPolicyInterface, appPolicyIsSet := d.GetOk("app_policy")
	if appPolicyIsSet {
		appPolicyMap := appPolicyInterface.([]interface{})[0].(map[string]interface{})
		appPolicy = AppAccessPolicyModelFromMap(appPolicyMap)
	}
	var devPolicy *models.DevAccessPolicy // DevAccessPolicy
	devPolicyInterface, devPolicyIsSet := d.GetOk("dev_policy")
	if devPolicyIsSet {
		devPolicyMap := devPolicyInterface.([]interface{})[0].(map[string]interface{})
		devPolicy = DevAccessPolicyModelFromMap(devPolicyMap)
	}
	var extPolicy *models.ExtAccessPolicy // ExtAccessPolicy
	extPolicyInterface, extPolicyIsSet := d.GetOk("ext_policy")
	if extPolicyIsSet {
		extPolicyMap := extPolicyInterface.([]interface{})[0].(map[string]interface{})
		extPolicy = ExtAccessPolicyModelFromMap(extPolicyMap)
	}
	generationIDInt, _ := d.Get("generation_id").(int)
	generationID := int64(generationIDInt)
	var jwtInfo *models.JWTInfo // JWTInfo
	jwtInfoInterface, jwtInfoIsSet := d.GetOk("jwt_info")
	if jwtInfoIsSet {
		jwtInfoMap := jwtInfoInterface.([]interface{})[0].(map[string]interface{})
		jwtInfo = JWTInfoModelFromMap(jwtInfoMap)
	}
	token, _ := d.Get("token").(string)
	return &models.EdgeviewCfg{
		AppPolicy:    appPolicy,
		DevPolicy:    devPolicy,
		ExtPolicy:    extPolicy,
		GenerationID: generationID,
		JwtInfo:      jwtInfo,
		Token:        token,
	}
}

func EdgeViewModelFromMap(m map[string]interface{}) *models.EdgeviewCfg {
	var appPolicy *models.AppAccessPolicy // AppAccessPolicy
	appPolicyInterface, appPolicyIsSet := m["app_policy"]
	if appPolicyIsSet {
		appPolicyMap := appPolicyInterface.([]interface{})[0].(map[string]interface{})
		appPolicy = AppAccessPolicyModelFromMap(appPolicyMap)
	}
	//
	var devPolicy *models.DevAccessPolicy // DevAccessPolicy
	devPolicyInterface, devPolicyIsSet := m["dev_policy"]
	if devPolicyIsSet {
		devPolicyMap := devPolicyInterface.([]interface{})[0].(map[string]interface{})
		devPolicy = DevAccessPolicyModelFromMap(devPolicyMap)
	}
	//
	var extPolicy *models.ExtAccessPolicy // ExtAccessPolicy
	extPolicyInterface, extPolicyIsSet := m["ext_policy"]
	if extPolicyIsSet {
		extPolicyMap := extPolicyInterface.([]interface{})[0].(map[string]interface{})
		extPolicy = ExtAccessPolicyModelFromMap(extPolicyMap)
	}
	//
	generationID := int64(m["generation_id"].(int)) // int64 false false false
	var jwtInfo *models.JWTInfo                     // JWTInfo
	jwtInfoInterface, jwtInfoIsSet := m["jwt_info"]
	if jwtInfoIsSet {
		jwtInfoMap := jwtInfoInterface.([]interface{})[0].(map[string]interface{})
		jwtInfo = JWTInfoModelFromMap(jwtInfoMap)
	}
	//
	token := m["token"].(string)
	return &models.EdgeviewCfg{
		AppPolicy:    appPolicy,
		DevPolicy:    devPolicy,
		ExtPolicy:    extPolicy,
		GenerationID: generationID,
		JwtInfo:      jwtInfo,
		Token:        token,
	}
}

func SetEdgeViewResourceData(d *schema.ResourceData, m *models.EdgeviewCfg) {
	d.Set("app_policy", SetAppAccessPolicySubResourceData([]*models.AppAccessPolicy{m.AppPolicy}))
	d.Set("dev_policy", SetDevAccessPolicySubResourceData([]*models.DevAccessPolicy{m.DevPolicy}))
	d.Set("ext_policy", SetExtAccessPolicySubResourceData([]*models.ExtAccessPolicy{m.ExtPolicy}))
	d.Set("generation_id", m.GenerationID)
	d.Set("jwt_info", SetJWTInfoSubResourceData([]*models.JWTInfo{m.JwtInfo}))
	d.Set("token", m.Token)
}

func SetEdgeViewSubResourceData(m []*models.EdgeviewCfg) (d []*map[string]interface{}) {
	for _, EdgeviewCfgModel := range m {
		if EdgeviewCfgModel != nil {
			properties := make(map[string]interface{})
			properties["app_policy"] = SetAppAccessPolicySubResourceData([]*models.AppAccessPolicy{EdgeviewCfgModel.AppPolicy})
			properties["dev_policy"] = SetDevAccessPolicySubResourceData([]*models.DevAccessPolicy{EdgeviewCfgModel.DevPolicy})
			properties["ext_policy"] = SetExtAccessPolicySubResourceData([]*models.ExtAccessPolicy{EdgeviewCfgModel.ExtPolicy})
			properties["generation_id"] = EdgeviewCfgModel.GenerationID
			properties["jwt_info"] = SetJWTInfoSubResourceData([]*models.JWTInfo{EdgeviewCfgModel.JwtInfo})
			properties["token"] = EdgeviewCfgModel.Token
			d = append(d, &properties)
		}
	}
	return
}

func Edgeview() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_policy": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AppAccessPolicy
			Elem: &schema.Resource{
				Schema: AppAccessPolicy(),
			},
			Optional: true,
		},

		"dev_policy": {
			Description: ``,
			Type:        schema.TypeList, //GoType: DevAccessPolicy
			Elem: &schema.Resource{
				Schema: DevAccessPolicy(),
			},
			Optional: true,
		},

		"ext_policy": {
			Description: ``,
			Type:        schema.TypeList, //GoType: ExtAccessPolicy
			Elem: &schema.Resource{
				Schema: ExtAccessPolicySchema(),
			},
			Optional: true,
		},

		"generation_id": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"jwt_info": {
			Description: ``,
			Type:        schema.TypeList, //GoType: JWTInfo
			Elem: &schema.Resource{
				Schema: JWTInfoSchema(),
			},
			Optional: true,
		},

		"token": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the EdgeviewCfg resource
func GetEdgeViewPropertyFields() (t []string) {
	return []string{
		"app_policy",
		"dev_policy",
		"ext_policy",
		"generation_id",
		"jwt_info",
		"token",
	}
}
