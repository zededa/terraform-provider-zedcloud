package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func ClusterInstPolicyModel(d *schema.ResourceData) *models.ClusterInstPolicy {
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
	if revisionIsSet && revisionInterface != nil {
		revisionMap := revisionInterface.([]interface{})
		if len(revisionMap) > 0 {
			revision = ObjectRevisionModelFromMap(revisionMap[0].(map[string]interface{}))
		}
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

func SetClusterInstPolicyResourceData(d *schema.ResourceData, m *models.ClusterInstPolicy) {
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("title", m.Title)
}

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

func ClusterInstPolicy() map[string]*schema.Schema {
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

func GetClusterInstPolicyPropertyFields() (t []string) {
	return []string{
		"id",
		"name",
		"revision",
		"title",
	}
}
