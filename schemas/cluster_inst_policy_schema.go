package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate ClusterInstPolicy resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ClusterInstPolicyModel(d *schema.ResourceData) *models.ClusterInstPolicy {
	id := d.Get("id").(string)
	name := d.Get("name").(string)
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := d.GetOk("revision")
	if revisionIsSet {
		revisionMap := revisionInterface.([]interface{})[0].(map[string]interface{})
		revision = ObjectRevisionModelFromMap(revisionMap)
	}
	title := d.Get("title").(string)
	return &models.ClusterInstPolicy{
		ID:       id,
		Name:     name,
		Revision: revision,
		Title:    title,
	}
}

func ClusterInstPolicyModelFromMap(m map[string]interface{}) *models.ClusterInstPolicy {
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
	return &models.ClusterInstPolicy{
		ID:       id,
		Name:     name,
		Revision: revision,
		Title:    title,
	}
}

// Update the underlying ClusterInstPolicy resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetClusterInstPolicyResourceData(d *schema.ResourceData, m *models.ClusterInstPolicy) {
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("title", m.Title)
}

// Iterate throught and update the ClusterInstPolicy resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetClusterInstPolicySubResourceData(m []*models.ClusterInstPolicy) (d []*map[string]interface{}) {
	for _, ClusterInstPolicyModel := range m {
		if ClusterInstPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["id"] = ClusterInstPolicyModel.ID
			properties["name"] = ClusterInstPolicyModel.Name
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{ClusterInstPolicyModel.Revision})
			properties["title"] = ClusterInstPolicyModel.Title
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ClusterInstPolicy resource defined in the Terraform configuration
func ClusterInstPolicySchema() map[string]*schema.Schema {
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
			Type:        schema.TypeList, //GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevisionSchema(),
			},
			Optional: true,
		},

		"title": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the ClusterInstPolicy resource
func GetClusterInstPolicyPropertyFields() (t []string) {
	return []string{
		"id",
		"name",
		"revision",
		"title",
	}
}
