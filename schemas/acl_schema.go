package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate ACL resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ACLModel(d *schema.ResourceData) *models.ACL {
	actions := d.Get("actions").([]*models.ACLAction) // []*ACLAction
	matches := d.Get("matches").([]*models.Match)     // []*Match
	name := d.Get("name").(string)
	return &models.ACL{
		Actions: actions,
		Matches: matches,
		Name:    name,
	}
}

func ACLModelFromMap(m map[string]interface{}) *models.ACL {
	actions := m["actions"].([]*models.ACLAction) // []*ACLAction
	matches := m["matches"].([]*models.Match)     // []*Match
	name := m["name"].(string)
	return &models.ACL{
		Actions: actions,
		Matches: matches,
		Name:    name,
	}
}

// Update the underlying ACL resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetACLResourceData(d *schema.ResourceData, m *models.ACL) {
	d.Set("actions", SetACLActionSubResourceData(m.Actions))
	d.Set("matches", SetMatchSubResourceData(m.Matches))
	d.Set("name", m.Name)
}

// Iterate throught and update the ACL resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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

// Schema mapping representing the ACL resource defined in the Terraform configuration
func ACLSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"actions": {
			Description: `Chain of actions to be taken on matching network traffic`,
			Type:        schema.TypeList, //GoType: []*ACLAction
			Elem: &schema.Resource{
				Schema: ACLActionSchema(),
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
