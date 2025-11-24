package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func GitRepoMetadataModel(d *schema.ResourceData) *models.GitRepoMetadata {
	description, _ := d.Get("description").(string)
	name, _ := d.Get("name").(string)
	title, _ := d.Get("title").(string)
	return &models.GitRepoMetadata{
		Description: description,
		Name:        &name, // string
		Title:       title,
	}
}

func GitRepoMetadataModelFromMap(m map[string]interface{}) *models.GitRepoMetadata {
	description := m["description"].(string)
	name := m["name"].(string)
	title := m["title"].(string)
	return &models.GitRepoMetadata{
		Description: description,
		Name:        &name,
		Title:       title,
	}
}

func SetGitRepoMetadataResourceData(d *schema.ResourceData, m *models.GitRepoMetadata) {
	d.Set("description", m.Description)
	d.Set("name", m.Name)
	d.Set("title", m.Title)
}

func SetGitRepoMetadataSubResourceData(m []*models.GitRepoMetadata) (d []*map[string]interface{}) {
	for _, GitRepoMetadataModel := range m {
		if GitRepoMetadataModel != nil {
			properties := make(map[string]interface{})
			properties["description"] = GitRepoMetadataModel.Description
			properties["name"] = GitRepoMetadataModel.Name
			properties["title"] = GitRepoMetadataModel.Title
			d = append(d, &properties)
		}
	}
	return
}

func GitRepoMetadataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Description: `Description of the GitRepo configuration`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"name": {
			Description: `Name of the GitRepo configuration (required)`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"title": {
			Description: `Display title of the GitRepo configuration`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetGitRepoMetadataPropertyFields() (t []string) {
	return []string{
		"description",
		"name",
		"title",
	}
}
