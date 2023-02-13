package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate PolicyCommon resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PolicyCommonModel(d *schema.ResourceData) *models.PolicyCommon {
	id, _ := d.Get("id").(string)
	name, _ := d.Get("name").(string)
	policyTargetCondition, _ := d.Get("policy_target_condition").(map[string]string) // map[string]string
	var revision *models.ObjectRevision                                              // ObjectRevision
	revisionInterface, revisionIsSet := d.GetOk("revision")
	if revisionIsSet {
		revisionMap := revisionInterface.([]interface{})[0].(map[string]interface{})
		revision = ObjectRevisionModelFromMap(revisionMap)
	}
	tags, _ := d.Get("tags").(map[string]string) // map[string]string
	title, _ := d.Get("title").(string)
	return &models.PolicyCommon{
		ID:                    id,
		Name:                  name,
		PolicyTargetCondition: policyTargetCondition,
		Revision:              revision,
		Tags:                  tags,
		Title:                 title,
	}
}

func PolicyCommonModelFromMap(m map[string]interface{}) *models.PolicyCommon {
	id := m["id"].(string)
	name := m["name"].(string)
	policyTargetCondition := m["policy_target_condition"].(map[string]string)
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := m["revision"]
	if revisionIsSet {
		revisionMap := revisionInterface.([]interface{})[0].(map[string]interface{})
		revision = ObjectRevisionModelFromMap(revisionMap)
	}
	//
	tags := m["tags"].(map[string]string)
	title := m["title"].(string)
	return &models.PolicyCommon{
		ID:                    id,
		Name:                  name,
		PolicyTargetCondition: policyTargetCondition,
		Revision:              revision,
		Tags:                  tags,
		Title:                 title,
	}
}

// Update the underlying PolicyCommon resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPolicyCommonResourceData(d *schema.ResourceData, m *models.PolicyCommon) {
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("policy_target_condition", m.PolicyTargetCondition)
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("tags", m.Tags)
	d.Set("title", m.Title)
}

// Iterate through and update the PolicyCommon resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPolicyCommonSubResourceData(m []*models.PolicyCommon) (d []*map[string]interface{}) {
	for _, PolicyCommonModel := range m {
		if PolicyCommonModel != nil {
			properties := make(map[string]interface{})
			properties["id"] = PolicyCommonModel.ID
			properties["name"] = PolicyCommonModel.Name
			properties["policy_target_condition"] = PolicyCommonModel.PolicyTargetCondition
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{PolicyCommonModel.Revision})
			properties["tags"] = PolicyCommonModel.Tags
			properties["title"] = PolicyCommonModel.Title
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the PolicyCommon resource defined in the Terraform configuration
func PolicyCommonSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Description: `system generated unique id for an policy`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: `user defined policy name`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"policy_target_condition": {
			Description: `user defined key-value pairs of a policy that will be used for targeting by devices`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"revision": {
			Description: `object revision details`,
			Type:        schema.TypeList, //GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevision(),
			},
			Optional: true,
		},

		"tags": {
			Description: `user defined key-value pairs of a policy`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"title": {
			Description: `user defined policy title`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the PolicyCommon resource
func GetPolicyCommonPropertyFields() (t []string) {
	return []string{
		"id",
		"name",
		"policy_target_condition",
		"revision",
		"tags",
		"title",
	}
}
