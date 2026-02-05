package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func PrivateRepoListResponseModel(d *schema.ResourceData) *models.PrivateRepoListResponse {
	var list []*models.PrivateRepoGetByIDResponse // []*PrivateRepoGetByIDResponse
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
			m := PrivateRepoGetByIDResponseModelFromMap(v.(map[string]interface{}))
			list = append(list, m)
		}
	}
	totalInt, _ := d.Get("total").(int)
	total := int32(totalInt)
	return &models.PrivateRepoListResponse{
		List:  list,
		Total: total,
	}
}

func PrivateRepoListResponseModelFromMap(m map[string]interface{}) *models.PrivateRepoListResponse {
	var list []*models.PrivateRepoGetByIDResponse // []*PrivateRepoGetByIDResponse
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
			m := PrivateRepoGetByIDResponseModelFromMap(v.(map[string]interface{}))
			list = append(list, m)
		}
	}
	total := int32(m["total"].(int)) // int32
	return &models.PrivateRepoListResponse{
		List:  list,
		Total: total,
	}
}

func SetPrivateRepoListResponseResourceData(d *schema.ResourceData, m *models.PrivateRepoListResponse) {
	d.Set("list", SetPrivateRepoGetByIDResponseSubResourceData(m.List))
	d.Set("total", m.Total)
}

func SetPrivateRepoListResponseSubResourceData(m []*models.PrivateRepoListResponse) (d []*map[string]interface{}) {
	for _, PrivateRepoListResponseModel := range m {
		if PrivateRepoListResponseModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetPrivateRepoGetByIDResponseSubResourceData(PrivateRepoListResponseModel.List)
			properties["total"] = PrivateRepoListResponseModel.Total
			d = append(d, &properties)
		}
	}
	return
}

func PrivateRepoListResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: `List of private repositories`,
			Type:        schema.TypeList, //GoType: []*PrivateRepoGetByIDResponse
			Elem: &schema.Resource{
				Schema: PrivateRepoGetByIDResponseSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"total": {
			Description: `Total number of private repositories`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetPrivateRepoListResponsePropertyFields() (t []string) {
	return []string{
		"list",
		"total",
	}
}
