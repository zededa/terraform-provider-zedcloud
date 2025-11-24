package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZksInstanceSummaryModel(d *schema.ResourceData) *models.ZksInstanceSummary {
	description, _ := d.Get("description").(string)
	id, _ := d.Get("id").(string)
	nOfNodesInt, _ := d.Get("n_of_nodes").(int)
	nOfNodes := int32(nOfNodesInt)
	name, _ := d.Get("name").(string)
	projectID, _ := d.Get("project_id").(string)
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
	return &models.ZksInstanceSummary{
		Description: description,
		ID:          id,
		NOfNodes:    nOfNodes,
		Name:        name,
		ProjectID:   projectID,
		Tags:        tags,
		Title:       title,
	}
}

func ZksInstanceSummaryModelFromMap(m map[string]interface{}) *models.ZksInstanceSummary {
	description := m["description"].(string)
	id := m["id"].(string)
	nOfNodes := int32(m["n_of_nodes"].(int)) // int32
	name := m["name"].(string)
	projectID := m["project_id"].(string)
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
	return &models.ZksInstanceSummary{
		Description: description,
		ID:          id,
		NOfNodes:    nOfNodes,
		Name:        name,
		ProjectID:   projectID,
		Tags:        tags,
		Title:       title,
	}
}

func SetZksInstanceSummaryResourceData(d *schema.ResourceData, m *models.ZksInstanceSummary) {
	d.Set("description", m.Description)
	d.Set("id", m.ID)
	d.Set("n_of_nodes", m.NOfNodes)
	d.Set("name", m.Name)
	d.Set("project_id", m.ProjectID)
	d.Set("tags", m.Tags)
	d.Set("title", m.Title)
}

func SetZksInstanceSummarySubResourceData(m []*models.ZksInstanceSummary) (d []*map[string]interface{}) {
	for _, ZksInstanceSummaryModel := range m {
		if ZksInstanceSummaryModel != nil {
			properties := make(map[string]interface{})
			properties["description"] = ZksInstanceSummaryModel.Description
			properties["id"] = ZksInstanceSummaryModel.ID
			properties["n_of_nodes"] = ZksInstanceSummaryModel.NOfNodes
			properties["name"] = ZksInstanceSummaryModel.Name
			properties["project_id"] = ZksInstanceSummaryModel.ProjectID
			properties["tags"] = ZksInstanceSummaryModel.Tags
			properties["title"] = ZksInstanceSummaryModel.Title
			d = append(d, &properties)
		}
	}
	return
}

func ZksInstanceSummarySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Description: `Detailed description of the zks instance`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: `System defined universally unique Id of the zks instance`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"n_of_nodes": {
			Description: `Number of nodes in the zks instance`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"name": {
			Description: `User defined name of the zks instance, unique across the enterprise. Once zks instance is created, name can’t be changed`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"project_id": {
			Description: `Foreign key to the project`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"tags": {
			Description: `Tags are name/value pairs that enable you to categorize resources.`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"title": {
			Description: `User defined title of the zks instance. Title can be changed at any time`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetZksInstanceSummaryPropertyFields() (t []string) {
	return []string{
		"description",
		"id",
		"n_of_nodes",
		"name",
		"project_id",
		"tags",
		"title",
	}
}
