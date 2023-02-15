package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func MatchModel(d *schema.ResourceData) *models.Match {
	typeVar, _ := d.Get("type").(string)
	value, _ := d.Get("value").(string)
	return &models.Match{
		Type:  typeVar,
		Value: value,
	}
}

func MatchModelFromMap(m map[string]interface{}) *models.Match {
	typeVar := m["type"].(string)
	value := m["value"].(string)
	return &models.Match{
		Type:  typeVar,
		Value: value,
	}
}

func SetMatchResourceData(d *schema.ResourceData, m *models.Match) {
	d.Set("type", m.Type)
	d.Set("value", m.Value)
}

func SetMatchSubResourceData(m []*models.Match) (d []*map[string]interface{}) {
	for _, MatchModel := range m {
		if MatchModel != nil {
			properties := make(map[string]interface{})
			properties["type"] = MatchModel.Type
			properties["value"] = MatchModel.Value
			d = append(d, &properties)
		}
	}
	return
}

func MatchSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"type": {
			Description: `Type of Match (Required)`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"value": {
			Description: `Value of match (Required)`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the Match resource
func GetMatchPropertyFields() (t []string) {
	return []string{
		"type",
		"value",
	}
}

func CompareMatchList(a, b []*models.Match) bool {
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
			if oldList.Type != newList.Type {
				continue
			}
			if oldList.Value != newList.Value {
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
			if oldList.Type != newList.Type {
				continue
			}
			if oldList.Value != newList.Value {
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
