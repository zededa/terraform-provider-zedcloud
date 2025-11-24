package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func MetadataModel(d *schema.ResourceData) *models.Metadata {
	description, _ := d.Get("description").(string)
	name, _ := d.Get("name").(string)
	projectID, _ := d.Get("project_id").(string)
	title, _ := d.Get("title").(string)
	return &models.Metadata{
		Description: description,
		Name:        &name,      // string
		ProjectID:   &projectID, // string
		Title:       title,
	}
}

func MetadataModelFromMap(m map[string]interface{}) *models.Metadata {
	description := m["description"].(string)
	name := m["name"].(string)
	projectID := m["project_id"].(string)
	title := m["title"].(string)
	return &models.Metadata{
		Description: description,
		Name:        &name,
		ProjectID:   &projectID,
		Title:       title,
	}
}

func SetMetadataResourceData(d *schema.ResourceData, m *models.Metadata) {
	d.Set("description", m.Description)
	d.Set("name", m.Name)
	d.Set("project_id", m.ProjectID)
	d.Set("title", m.Title)
}

func SetMetadataSubResourceData(m []*models.Metadata) (d []*map[string]interface{}) {
	for _, MetadataModel := range m {
		if MetadataModel != nil {
			properties := make(map[string]interface{})
			properties["description"] = MetadataModel.Description
			properties["name"] = MetadataModel.Name
			properties["project_id"] = MetadataModel.ProjectID
			properties["title"] = MetadataModel.Title
			d = append(d, &properties)
		}
	}
	return
}

func MetadataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Description: `Description of the deployment. Example: 'Production NGINX deployment with custom configuration'`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"name": {
			Description: `Name of the deployment. Example: 'nginx-deployment'`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"project_id": {
			Description: `Project ID associated with the deployment`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"title": {
			Description: `Human-readable title of the deployment. Example: 'NGINX Web Server'`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetMetadataPropertyFields() (t []string) {
	return []string{
		"description",
		"name",
		"project_id",
		"title",
	}
}
