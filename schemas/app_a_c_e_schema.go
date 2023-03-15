package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AppACEModel(d *schema.ResourceData) *models.AppACE {
	var actions []*models.AppACEAction // []*AppACEAction
	actionsInterface, actionsIsSet := d.GetOk("actions")
	if actionsIsSet {
		var items []interface{}
		if listItems, isList := actionsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = actionsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := AppACEActionModelFromMap(v.(map[string]interface{}))
			actions = append(actions, m)
		}
	}
	idInt, _ := d.Get("id").(int)
	id := int32(idInt)
	var matches []*models.AppACEMatch // []*AppACEMatch
	matchesInterface, matchesIsSet := d.GetOk("matches")
	if matchesIsSet {
		var items []interface{}
		if listItems, isList := matchesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = matchesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := AppACEMatchModelFromMap(v.(map[string]interface{}))
			matches = append(matches, m)
		}
	}
	name, _ := d.Get("name").(string)
	return &models.AppACE{
		Actions: actions,
		ID:      &id, // int32
		Matches: matches,
		Name:    &name, // string
	}
}

func AppACEModelFromMap(m map[string]interface{}) *models.AppACE {
	var actions []*models.AppACEAction // []*AppACEAction
	actionsInterface, actionsIsSet := m["actions"]
	if actionsIsSet {
		var items []interface{}
		if listItems, isList := actionsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = actionsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := AppACEActionModelFromMap(v.(map[string]interface{}))
			actions = append(actions, m)
		}
	}
	id := int32(m["id"].(int))        // int32
	var matches []*models.AppACEMatch // []*AppACEMatch
	matchesInterface, matchesIsSet := m["matches"]
	if matchesIsSet {
		var items []interface{}
		if listItems, isList := matchesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = matchesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := AppACEMatchModelFromMap(v.(map[string]interface{}))
			matches = append(matches, m)
		}
	}
	name := m["name"].(string)
	return &models.AppACE{
		Actions: actions,
		ID:      &id,
		Matches: matches,
		Name:    &name,
	}
}

func SetAppACEResourceData(d *schema.ResourceData, m *models.AppACE) {
	d.Set("actions", SetAppACEActionSubResourceData(m.Actions))
	d.Set("id", m.ID)
	d.Set("matches", SetAppACEMatchSubResourceData(m.Matches))
	d.Set("name", m.Name)
}

func SetAppACESubResourceData(m []*models.AppACE) (d []*map[string]interface{}) {
	for _, AppACEModel := range m {
		if AppACEModel != nil {
			properties := make(map[string]interface{})
			properties["actions"] = SetAppACEActionSubResourceData(AppACEModel.Actions)
			properties["id"] = AppACEModel.ID
			properties["matches"] = SetAppACEMatchSubResourceData(AppACEModel.Matches)
			properties["name"] = AppACEModel.Name
			d = append(d, &properties)
		}
	}
	return
}

func AppACE() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"actions": {
			Description: `app ACE actions`,
			Type:        schema.TypeList, //GoType: []*AppACEAction
			Elem: &schema.Resource{
				Schema: AppACEActionSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Required: true,
		},

		"id": {
			Description: `app ACE id`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"matches": {
			Description: `app ACE match`,
			Type:        schema.TypeList, //GoType: []*AppACEMatch
			Elem: &schema.Resource{
				Schema: AppACEMatchSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Required: true,
		},

		"name": {
			Description: `User defined name of the app ACE, unique across the enterprise. Once object is created, name can’t be changed`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func GetAppACEPropertyFields() (t []string) {
	return []string{
		"actions",
		"id",
		"matches",
		"name",
	}
}
