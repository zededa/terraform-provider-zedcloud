package schemas

import (
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ProfilePolicyCommonModel(d *schema.ResourceData) *models.ProfilePolicyCommon {
	id, _ := d.Get("id").(string)
	if id == "" {
		id, _ = uuid.GenerateUUID()
	}
	name, _ := d.Get("name").(string)
	tags := map[string]string{}
	tagsInterface, tagsIsSet := d.GetOk("tags")
	if tagsIsSet {
		tagsMap := tagsInterface.(map[string]interface{})
		for k, v := range tagsMap {
			if v == nil {
				continue
			}
			tags[k] = v.(string)
		}
	}

	title, _ := d.Get("title").(string)
	return &models.ProfilePolicyCommon{
		ID:    id,
		Name:  name,
		Tags:  tags,
		Title: title,
	}
}

func ProfilePolicyCommonModelFromMap(m map[string]interface{}) *models.ProfilePolicyCommon {
	id := m["id"].(string)
	if id == "" {
		id, _ = uuid.GenerateUUID()
	}
	name := m["name"].(string)
	tags := map[string]string{}
	tagsInterface, tagsIsSet := m["tags"]
	if tagsIsSet {
		tagsMap := tagsInterface.(map[string]interface{})
		for k, v := range tagsMap {
			if v == nil {
				continue
			}
			tags[k] = v.(string)
		}
	}

	title := m["title"].(string)
	return &models.ProfilePolicyCommon{
		ID:    id,
		Name:  name,
		Tags:  tags,
		Title: title,
	}
}

func SetProfilePolicyCommonResourceData(d *schema.ResourceData, m *models.ProfilePolicyCommon) {
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("tags", m.Tags)
	d.Set("title", m.Title)
}

func SetProfilePolicyCommonSubResourceData(m []*models.ProfilePolicyCommon) (d []*map[string]interface{}) {
	for _, ProfilePolicyCommonModel := range m {
		if ProfilePolicyCommonModel != nil {
			properties := make(map[string]interface{})
			properties["id"] = ProfilePolicyCommonModel.ID
			properties["name"] = ProfilePolicyCommonModel.Name
			properties["tags"] = ProfilePolicyCommonModel.Tags
			properties["title"] = ProfilePolicyCommonModel.Title
			d = append(d, &properties)
		}
	}
	return
}

func ProfilePolicyCommonSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Description: `unique id for an policy`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: `user defined policy name`,
			Type:        schema.TypeString,
			Optional:    true,
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

func GetProfilePolicyCommonPropertyFields() (t []string) {
	return []string{
		"id",
		"name",
		"tags",
		"title",
	}
}
