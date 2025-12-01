package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ClusterGroupModel(d *schema.ResourceData) *models.ClusterGroup {
	description, _ := d.Get("description").(string)
	id, _ := d.Get("id").(string)
	var manifest *models.ClusterGroupManifest // ClusterGroupManifest
	manifestInterface, manifestIsSet := d.GetOk("manifest")
	if manifestIsSet && manifestInterface != nil {
		manifestMap := manifestInterface.([]interface{})
		if len(manifestMap) > 0 {
			manifest = ClusterGroupManifestModelFromMap(manifestMap[0].(map[string]interface{}))
		}
	}
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
	return &models.ClusterGroup{
		Description: description,
		ID:          id,
		Manifest:    manifest,
		Name:        &name,      // string
		ProjectID:   &projectID, // string
		Tags:        tags,
		Title:       title,
	}
}

func ClusterGroupModelFromMap(m map[string]interface{}) *models.ClusterGroup {
	description := m["description"].(string)
	id := m["id"].(string)
	var manifest *models.ClusterGroupManifest // ClusterGroupManifest
	manifestInterface, manifestIsSet := m["manifest"]
	if manifestIsSet && manifestInterface != nil {
		manifestMap := manifestInterface.([]interface{})
		if len(manifestMap) > 0 {
			manifest = ClusterGroupManifestModelFromMap(manifestMap[0].(map[string]interface{}))
		}
	}
	//
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
	return &models.ClusterGroup{
		Description: description,
		ID:          id,
		Manifest:    manifest,
		Name:        &name,
		ProjectID:   &projectID,
		Tags:        tags,
		Title:       title,
	}
}

func SetClusterGroupResourceData(d *schema.ResourceData, m *models.ClusterGroup) {
	d.Set("description", m.Description)
	d.Set("id", m.ID)
	d.Set("manifest", SetClusterGroupManifestSubResourceData([]*models.ClusterGroupManifest{m.Manifest}))
	d.Set("name", m.Name)
	d.Set("project_id", m.ProjectID)
	d.Set("tags", m.Tags)
	d.Set("title", m.Title)
}

func SetClusterGroupSubResourceData(m []*models.ClusterGroup) (d []*map[string]interface{}) {
	for _, ClusterGroupModel := range m {
		if ClusterGroupModel != nil {
			properties := make(map[string]interface{})
			properties["description"] = ClusterGroupModel.Description
			properties["id"] = ClusterGroupModel.ID
			properties["manifest"] = SetClusterGroupManifestSubResourceData([]*models.ClusterGroupManifest{ClusterGroupModel.Manifest})
			properties["name"] = ClusterGroupModel.Name
			properties["project_id"] = ClusterGroupModel.ProjectID
			properties["tags"] = ClusterGroupModel.Tags
			properties["title"] = ClusterGroupModel.Title
			d = append(d, &properties)
		}
	}
	return
}

func ClusterGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Description: `Detailed description of the cluster group`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: `System defined universally unique Id of the cluster group`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"manifest": {
			Description: `The cluster group manifest`,
			Type:        schema.TypeList, //GoType: ClusterGroupManifest
			Elem: &schema.Resource{
				Schema: ClusterGroupManifestSchema(),
			},
			Optional: true,
		},

		"name": {
			Description: `User defined name of the cluster group, unique across the enterprise. Once cluster group is created, name can’t be changed`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"project_id": {
			Description: `Foreign key to the project`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"tags": {
			Description: `Tags are name/value pairs that enable you to categorize resources. Tag names are case insensitive with max_length 512 and min_length 3. Tag values are case sensitive with max_length 256 and min_length 3.`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"title": {
			Description: `User defined title of the cluster group. Title can be changed at any time`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetClusterGroupPropertyFields() (t []string) {
	return []string{
		"description",
		"id",
		"manifest",
		"name",
		"project_id",
		"tags",
		"title",
	}
}
