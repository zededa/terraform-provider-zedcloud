package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func PolicyVersionListModel(d *schema.ResourceData) *models.PolicyVersionList {
	var list []*models.PolicyVersion // []*PolicyVersion
	listInterface, listIsSet := d.GetOk("list")
	if listIsSet {
		var items []interface{}
		if listItems, isList := listInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = listInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := PolicyVersionModelFromMap(v.(map[string]interface{}))
			list = append(list, m)
		}
	}
	return &models.PolicyVersionList{
		List: list,
	}
}

func PolicyVersionListModelFromMap(m map[string]interface{}) *models.PolicyVersionList {
	var list []*models.PolicyVersion // []*PolicyVersion
	listInterface, listIsSet := m["list"]
	if listIsSet {
		var items []interface{}
		if listItems, isList := listInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = listInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := PolicyVersionModelFromMap(v.(map[string]interface{}))
			list = append(list, m)
		}
	}
	return &models.PolicyVersionList{
		List: list,
	}
}

func SetPolicyVersionListResourceData(d *schema.ResourceData, m *models.PolicyVersionList) {
	d.Set("list", SetPolicyVersionSubResourceData(m.List))
}

func SetPolicyVersionListSubResourceData(m []*models.PolicyVersionList) (d []*map[string]interface{}) {
	for _, PolicyVersionListModel := range m {
		if PolicyVersionListModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetPolicyVersionSubResourceData(PolicyVersionListModel.List)
			d = append(d, &properties)
		}
	}
	return
}

func PolicyVersionListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*PolicyVersion
			Elem: &schema.Resource{
				Schema: PolicyVersionSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
	}
}

func GetPolicyVersionListPropertyFields() (t []string) {
	return []string{
		"list",
	}
}
