package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func IntegrationPolicyModel(d *schema.ResourceData) *models.IntegrationPolicy {
	id, _ := d.Get("id").(string)
	name, _ := d.Get("name").(string)
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := d.GetOk("revision")
	if revisionIsSet && revisionInterface != nil {
		revisionMap := revisionInterface.([]interface{})
		if len(revisionMap) > 0 {
			revision = ObjectRevisionModelFromMap(revisionMap[0].(map[string]interface{}))
		}
	}
	title, _ := d.Get("title").(string)
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
	if revisionIsSet && revisionInterface != nil {
		revisionMap := revisionInterface.([]interface{})
		if len(revisionMap) > 0 {
			revision = ObjectRevisionModelFromMap(revisionMap[0].(map[string]interface{}))
		}
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

func SetIntegrationPolicyResourceData(d *schema.ResourceData, m *models.IntegrationPolicy) {
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("title", m.Title)
}

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

func IntegrationPolicy() map[string]*schema.Schema {
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
				Schema: ObjectRevision(),
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

func GetIntegrationPolicyPropertyFields() (t []string) {
	return []string{
		"id",
		"name",
		"revision",
		"title",
	}
}
