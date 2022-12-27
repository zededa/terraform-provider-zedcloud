package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate IntegrationPolicy resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func IntegrationPolicyModel(d *schema.ResourceData) *models.IntegrationPolicy {
	id := d.Get("id").(string)
	name := d.Get("name").(string)
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := d.GetOk("revision")
	if revisionIsSet {
		revisionMap := revisionInterface.([]interface{})[0].(map[string]interface{})
		revision = ObjectRevisionModelFromMap(revisionMap)
	}
	title := d.Get("title").(string)
	return &models.IntegrationPolicy{
		ID:       id,
		Name:     name,
		Revision: revision,
		Title:    title,
	}
}

func IntegrationPolicyModelFromMap(m map[string]interface{}) *models.IntegrationPolicy {
	id := m["id"].(string)
	name := m["name"].(string)
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := m["revision"]
	if revisionIsSet {
		revisionMap := revisionInterface.([]interface{})[0].(map[string]interface{})
		revision = ObjectRevisionModelFromMap(revisionMap)
	}
	//
	title := m["title"].(string)
	return &models.IntegrationPolicy{
		ID:       id,
		Name:     name,
		Revision: revision,
		Title:    title,
	}
}

// Update the underlying IntegrationPolicy resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetIntegrationPolicyResourceData(d *schema.ResourceData, m *models.IntegrationPolicy) {
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("title", m.Title)
}

// Iterate throught and update the IntegrationPolicy resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetIntegrationPolicySubResourceData(m []*models.IntegrationPolicy) (d []*map[string]interface{}) {
	for _, IntegrationPolicyModel := range m {
		if IntegrationPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["id"] = IntegrationPolicyModel.ID
			properties["name"] = IntegrationPolicyModel.Name
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{IntegrationPolicyModel.Revision})
			properties["title"] = IntegrationPolicyModel.Title
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the IntegrationPolicy resource defined in the Terraform configuration
func IntegrationPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"revision": {
			Description: ``,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},

		"title": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the IntegrationPolicy resource
func GetIntegrationPolicyPropertyFields() (t []string) {
	return []string{
		"id",
		"name",
		"revision",
		"title",
	}
}
