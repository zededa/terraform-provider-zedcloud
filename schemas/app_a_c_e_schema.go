package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate AppACE resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AppACEModel(d *schema.ResourceData) *models.AppACE {
	actions := d.Get("actions").([]*models.AppACEAction) // []*AppACEAction
	id := int32(d.Get("id").(int))
	matches := d.Get("matches").([]*models.AppACEMatch) // []*AppACEMatch
	name := d.Get("name").(string)
	return &models.AppACE{
		Actions: actions,
		ID:      &id, // int32 true false false
		Matches: matches,
		Name:    &name, // string true false false
	}
}

func AppACEModelFromMap(m map[string]interface{}) *models.AppACE {
	actions := m["actions"].([]*models.AppACEAction) // []*AppACEAction
	id := int32(m["id"].(int))                       // int32 true false false
	matches := m["matches"].([]*models.AppACEMatch)  // []*AppACEMatch
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

// Iterate throught and update the AppACE resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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
func AppACESchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"actions": {
			Description: `app ACE actions`,
			Type:        schema.TypeList, //GoType: []*AppACEAction
			Elem: &schema.Resource{
				Schema: AppACEActionSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Required:   true,
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
			ConfigMode: schema.SchemaConfigModeAttr,
			Required:   true,
		},

		"name": {
			Description: `User defined name of the app ACE, unique across the enterprise. Once object is created, name can’t be changed`,
			Type:        schema.TypeString,
			Required:    true,
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
