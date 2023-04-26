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
	id := int32(d.Get("id").(int))
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
		ID:      &id, // int32 true false false
		Matches: matches,
		Name:    &name, // string true false false
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
	id := int32(m["id"].(int))
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

// Update the underlying AppACE resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppACEResourceData(d *schema.ResourceData, m *models.AppACE) {
	d.Set("actions", SetAppACEActionSubResourceData(m.Actions))
	d.Set("id", m.ID)
	d.Set("matches", SetAppACEMatchSubResourceData(m.Matches))
	d.Set("name", m.Name)
}

// Iterate through and update the AppACE resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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

// Schema mapping representing the AppACE resource defined in the Terraform configuration
func AppACE() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"actions": {
			Description: `app ACE actions`,
			Type:        schema.TypeList, //GoType: []*AppACEAction
			Elem: &schema.Resource{
				Schema: AppACEActionSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"id": {
			Description: `app ACE id`,
			Type:        schema.TypeInt,
			Computed:    true,
		},

		"matches": {
			Description: `app ACE match`,
			Type:        schema.TypeList, //GoType: []*AppACEMatch
			Elem: &schema.Resource{
				Schema: AppACEMatchSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"name": {
			Description: `User defined name of the app ACE, unique across the enterprise. Once object is created, name can’t be changed`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the AppACE resource
func GetAppACEPropertyFields() (t []string) {
	return []string{
		"actions",
		"id",
		"matches",
		"name",
	}
}
