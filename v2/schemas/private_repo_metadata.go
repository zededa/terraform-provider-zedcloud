package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func PrivateRepoMetadataModel(d *schema.ResourceData) *models.PrivateRepoMetadata {
	description, _ := d.Get("description").(string)
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
	return &models.PrivateRepoMetadata{
		Description: description,
		Name:        &name, // string
		Tags:        tags,
		Title:       title,
	}
}

func PrivateRepoMetadataModelFromMap(m map[string]interface{}) *models.PrivateRepoMetadata {
	description := m["description"].(string)
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
	return &models.PrivateRepoMetadata{
		Description: description,
		Name:        &name,
		Tags:        tags,
		Title:       title,
	}
}

func SetPrivateRepoMetadataResourceData(d *schema.ResourceData, m *models.PrivateRepoMetadata) {
	d.Set("description", m.Description)
	d.Set("name", m.Name)
	d.Set("tags", m.Tags)
	d.Set("title", m.Title)
}

func SetPrivateRepoMetadataSubResourceData(m []*models.PrivateRepoMetadata) (d []*map[string]interface{}) {
	for _, PrivateRepoMetadataModel := range m {
		if PrivateRepoMetadataModel != nil {
			properties := make(map[string]interface{})
			properties["description"] = PrivateRepoMetadataModel.Description
			properties["name"] = PrivateRepoMetadataModel.Name
			properties["tags"] = PrivateRepoMetadataModel.Tags
			properties["title"] = PrivateRepoMetadataModel.Title
			d = append(d, &properties)
		}
	}
	return
}

func PrivateRepoMetadataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Description: `Description of the private repository configuration`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"name": {
			Description: `Name of the private repository configuration (required)`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"tags": {
			Description: `Tags associated with the private repository`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"title": {
			Description: `Display title of the private repository configuration`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetPrivateRepoMetadataPropertyFields() (t []string) {
	return []string{
		"description",
		"name",
		"tags",
		"title",
	}
}
