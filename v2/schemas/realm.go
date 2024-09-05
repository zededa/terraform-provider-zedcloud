package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func RealmModel(d *schema.ResourceData) *models.Realm {
	description, _ := d.Get("description").(string)
	enterpriseID, _ := d.Get("enterprise_id").(string)
	id, _ := d.Get("id").(string)
	name, _ := d.Get("name").(string)
	title, _ := d.Get("title").(string)
	return &models.Realm{
		Description:  description,
		EnterpriseID: &enterpriseID, // string
		ID:           id,
		Name:         &name,  // string
		Title:        &title, // string
	}
}

func RealmModelFromMap(m map[string]interface{}) *models.Realm {
	description := m["description"].(string)
	enterpriseID := m["enterprise_id"].(string)
	id := m["id"].(string)
	name := m["name"].(string)
	title := m["title"].(string)
	return &models.Realm{
		Description:  description,
		EnterpriseID: &enterpriseID,
		ID:           id,
		Name:         &name,
		Title:        &title,
	}
}

func SetRealmResourceData(d *schema.ResourceData, m *models.Realm) {
	d.Set("description", m.Description)
	d.Set("enterprise_id", m.EnterpriseID)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("title", m.Title)
}

func SetRealmSubResourceData(m []*models.Realm) (d []*map[string]interface{}) {
	for _, RealmModel := range m {
		if RealmModel != nil {
			properties := make(map[string]interface{})
			properties["description"] = RealmModel.Description
			properties["enterprise_id"] = RealmModel.EnterpriseID
			properties["id"] = RealmModel.ID
			properties["name"] = RealmModel.Name
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{RealmModel.Revision})
			properties["title"] = RealmModel.Title
			d = append(d, &properties)
		}
	}
	return
}

func RealmSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Description: `Detailed description of the realm`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"enterprise_id": {
			Description: `Enterprise ID of the enterprise where the realm is to be created`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"id": {
			Description: `Unique system defined realm ID`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: `User defined name of the realm. Name cannot be changed once created`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"revision": {
			Description: `System defined info`,
			Type:        schema.TypeList, //GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevision(),
			},
			Computed: true,
		},

		"title": {
			Description: `User defined title of the realm. Title can be changed anytime`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func GetRealmPropertyFields() (t []string) {
	return []string{
		"description",
		"enterprise_id",
		"id",
		"name",
		"title",
	}
}
