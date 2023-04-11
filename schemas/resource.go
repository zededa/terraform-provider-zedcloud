package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func ResourceModel(d *schema.ResourceData) *models.Resource {
	name, _ := d.Get("name").(string)
	value, _ := d.Get("value").(string)
	return &models.Resource{
		Name:  name,
		Value: value,
	}
}

func ResourceModelFromMap(m map[string]interface{}) *models.Resource {
	name := m["name"].(string)
	value := m["value"].(string)
	return &models.Resource{
		Name:  name,
		Value: value,
	}
}

func SetResourceResourceData(d *schema.ResourceData, m *models.Resource) {
	d.Set("name", m.Name)
	d.Set("value", m.Value)
}

func SetResourceSubResourceData(m []*models.Resource) (d []*map[string]interface{}) {
	for _, ResourceModel := range m {
		if ResourceModel != nil {
			properties := make(map[string]interface{})
			properties["name"] = ResourceModel.Name
			properties["value"] = ResourceModel.Value
			d = append(d, &properties)
		}
	}
	return
}

func Resource() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Description: `Name of the Resource (Required)`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"value": {
			Description: `Value of Resource (Required)`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the Resource resource
func GetResourcePropertyFields() (t []string) {
	return []string{
		"name",
		"value",
	}
}

func CompareResourceLists(a, b []*models.Resource) bool {
	if len(a) != len(b) {
		return false
	}
	// is each element of the new list in the old list?
	for _, newList := range b {
		if newList == nil {
			continue
		}

		found := false
		for _, oldList := range a {
			if oldList == nil {
				continue
			}
			if oldList.Value != newList.Value {
				continue
			}
			if oldList.Name != newList.Name {
				continue
			}
			found = true
			break
		}
		if !found {
			return false
		}
	}

	// is each element of the old list also in the new list?
	for _, oldList := range a {
		if oldList == nil {
			continue
		}

		found := false
		for _, newList := range b {
			if newList == nil {
				continue
			}
			if oldList.Value != newList.Value {
				continue
			}
			if oldList.Name != newList.Name {
				continue
			}
			found = true
			break
		}
		if !found {
			return false
		}
	}
	return true
}
