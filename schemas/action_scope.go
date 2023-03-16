package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func ActionScopeModel(d *schema.ResourceData) *models.ActionScope {
	id, _ := d.Get("id").(string)
	var typeVar *models.ObjectType // ObjectType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewObjectType(models.ObjectType(typeModel))
	}
	return &models.ActionScope{
		ID:   id,
		Type: typeVar,
	}
}

func ActionScopeModelFromMap(m map[string]interface{}) *models.ActionScope {
	id := m["id"].(string)
	var typeVar *models.ObjectType // ObjectType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewObjectType(models.ObjectType(typeModel))
	}
	return &models.ActionScope{
		ID:   id,
		Type: typeVar,
	}
}

// Update the underlying ActionScope resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetActionScopeResourceData(d *schema.ResourceData, m *models.ActionScope) {
	d.Set("id", m.ID)
	d.Set("type", m.Type)
}

// Iterate through and update the ActionScope resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetActionScopeSubResourceData(m []*models.ActionScope) (d []*map[string]interface{}) {
	for _, ActionScopeModel := range m {
		if ActionScopeModel != nil {
			properties := make(map[string]interface{})
			properties["id"] = ActionScopeModel.ID
			properties["type"] = ActionScopeModel.Type
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ActionScope resource defined in the Terraform configuration
func ActionScopeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Description: `Unique id of the type (e.g. projectId in case of project type)`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"type": {
			Description: `Scope of the action (e.g. project in case certificate expired at project level)`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the ActionScope resource
func GetActionScopePropertyFields() (t []string) {
	return []string{
		"id",
		"type",
	}
}
