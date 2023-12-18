package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func RoleModel(d *schema.ResourceData) *models.Role {
	description, _ := d.Get("description").(string)
	id, _ := d.Get("id").(string)
	name, _ := d.Get("name").(string)
	projectTags, _ := d.Get("project_tags").(string)
	var scopes []*models.Scope // []*Scope
	scopesInterface, scopesIsSet := d.GetOk("scopes")
	if scopesIsSet {
		var items []interface{}
		if listItems, isList := scopesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = scopesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ScopeModelFromMap(v.(map[string]interface{}))
			scopes = append(scopes, m)
		}
	}
	var state *models.RoleState // RoleState
	stateInterface, stateIsSet := d.GetOk("state")
	if stateIsSet {
		stateModel := stateInterface.(string)
		state = models.NewRoleState(models.RoleState(stateModel))
	}
	title, _ := d.Get("title").(string)
	var typeVar *models.UserRole // UserRole
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewUserRole(models.UserRole(typeModel))
	}
	return &models.Role{
		Description: description,
		ID:          id,
		Name:        &name, // string
		ProjectTags: projectTags,
		Scopes:      scopes,
		State:       state,
		Title:       &title, // string
		Type:        typeVar,
	}
}

func RoleModelFromMap(m map[string]interface{}) *models.Role {
	description := m["description"].(string)
	id := m["id"].(string)
	name := m["name"].(string)
	projectTags := m["project_tags"].(string)
	var scopes []*models.Scope // []*Scope
	scopesInterface, scopesIsSet := m["scopes"]
	if scopesIsSet {
		var items []interface{}
		if listItems, isList := scopesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = scopesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ScopeModelFromMap(v.(map[string]interface{}))
			scopes = append(scopes, m)
		}
	}
	var state *models.RoleState // RoleState
	stateInterface, stateIsSet := m["state"]
	if stateIsSet {
		stateModel := stateInterface.(string)
		state = models.NewRoleState(models.RoleState(stateModel))
	}
	title := m["title"].(string)
	var typeVar *models.UserRole // UserRole
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewUserRole(models.UserRole(typeModel))
	}
	return &models.Role{
		Description: description,
		ID:          id,
		Name:        &name,
		ProjectTags: projectTags,
		Scopes:      scopes,
		State:       state,
		Title:       &title,
		Type:        typeVar,
	}
}

func SetRoleResourceData(d *schema.ResourceData, m *models.Role) {
	d.Set("description", m.Description)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("project_tags", m.ProjectTags)
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("scopes", SetScopeSubResourceData(m.Scopes))
	d.Set("state", m.State)
	d.Set("title", m.Title)
	d.Set("type", m.Type)
}

func SetRoleSubResourceData(m []*models.Role) (d []*map[string]interface{}) {
	for _, RoleModel := range m {
		if RoleModel != nil {
			properties := make(map[string]interface{})
			properties["description"] = RoleModel.Description
			properties["id"] = RoleModel.ID
			properties["name"] = RoleModel.Name
			properties["project_tags"] = RoleModel.ProjectTags
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{RoleModel.Revision})
			properties["scopes"] = SetScopeSubResourceData(RoleModel.Scopes)
			properties["state"] = RoleModel.State
			properties["title"] = RoleModel.Title
			properties["type"] = RoleModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func RoleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Description: `Detailed description of the role`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: `Unique system defined role ID`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: `User defined name of the role. Name cannot be changed once created`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"project_tags": {
			Description: `Map of project tags filter`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"revision": {
			Description: `System defined info`,
			Type:        schema.TypeList, //GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevision(),
			},
			Computed: true,
		},

		"scopes": {
			Description: `Scopes/Permissions associated with the role`,
			Type:        schema.TypeList, //GoType: []*Scope
			Elem: &schema.Resource{
				Schema: ScopeSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Required: true,
		},

		"state": {
			Description: `State of the role`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"title": {
			Description: `User defined title of the role. Title can be changed anytime`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"type": {
			Description: `Type of the role`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func GetRolePropertyFields() (t []string) {
	return []string{
		"description",
		"id",
		"name",
		"project_tags",
		"scopes",
		"state",
		"title",
		"type",
	}
}
