package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func ACLModel(d *schema.ResourceData) *models.ACL {
	var actions []*models.ACLAction // []*ACLAction
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
			m := ACLActionModelFromMap(v.(map[string]interface{}))
			actions = append(actions, m)
		}
	}
	var matches []*models.Match // []*Match
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
			m := MatchModelFromMap(v.(map[string]interface{}))
			matches = append(matches, m)
		}
	}
	name, _ := d.Get("name").(string)
	return &models.ACL{
		Actions: actions,
		Matches: matches,
		Name:    name,
	}
}

func ACLModelFromMap(m map[string]interface{}) *models.ACL {
	var actions []*models.ACLAction // []*ACLAction
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
			m := ACLActionModelFromMap(v.(map[string]interface{}))
			actions = append(actions, m)
		}
	}
	var matches []*models.Match // []*Match
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
			m := MatchModelFromMap(v.(map[string]interface{}))
			matches = append(matches, m)
		}
	}
	name := m["name"].(string)
	return &models.ACL{
		Actions: actions,
		Matches: matches,
		Name:    name,
	}
}

func SetACLResourceData(d *schema.ResourceData, m *models.ACL) {
	d.Set("actions", SetACLActionSubResourceData(m.Actions))
	d.Set("matches", SetMatchSubResourceData(m.Matches))
	d.Set("name", m.Name)
}

func SetACLSubResourceData(m []*models.ACL) (d []*map[string]interface{}) {
	for _, ACLModel := range m {
		if ACLModel != nil {
			properties := make(map[string]interface{})
			properties["actions"] = SetACLActionSubResourceData(ACLModel.Actions)
			properties["matches"] = SetMatchSubResourceData(ACLModel.Matches)
			properties["name"] = ACLModel.Name
			d = append(d, &properties)
		}
	}
	return
}

func ACL() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"actions": {
			Description: `Chain of actions to be taken on matching network traffic`,
			Type:        schema.TypeList, //GoType: []*ACLAction
			Elem: &schema.Resource{
				Schema: ACLAction(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"matches": {
			Description: `Network traffic matching criteria consistngs of one or more of source IP address, destination IP address, protocol, source port and destination port`,
			Type:        schema.TypeList, //GoType: []*Match
			Elem: &schema.Resource{
				Schema: MatchSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"name": {
			Description: `Name of the Access Control List`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the ACL resource
func GetACLPropertyFields() (t []string) {
	return []string{
		"actions",
		"matches",
		"name",
	}
}

func CompareACLList(a, b []*models.ACL) bool {
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
			if oldList.Name != newList.Name {
				continue
			}
			if !CompareMatchList(oldList.Matches, newList.Matches) {
				continue
			}
			if !CompareACLActionList(oldList.Actions, newList.Actions) {
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
			if oldList.Name != newList.Name {
				continue
			}
			if !CompareMatchList(oldList.Matches, newList.Matches) {
				continue
			}
			if !CompareACLActionList(oldList.Actions, newList.Actions) {
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
