package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZksInstanceModel(d *schema.ResourceData) *models.ZksInstance {
	clusterPrefix, _ := d.Get("cluster_prefix").(string)
	description, _ := d.Get("description").(string)
	id, _ := d.Get("id").(string)
	isImported, _ := d.Get("is_imported").(bool)
	name, _ := d.Get("name").(string)
	var nodes []*models.ZksInstanceNode // []*ZksInstanceNode
	nodesInterface, nodesIsSet := d.GetOk("nodes")
	if nodesIsSet {
		var items []interface{}
		if listItems, isList := nodesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = nodesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ZksInstanceNodeModelFromMap(v.(map[string]interface{}))
			nodes = append(nodes, m)
		}
	}
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
	return &models.ZksInstance{
		ClusterPrefix: clusterPrefix,
		Description:   description,
		ID:            id,
		IsImported:    isImported,
		Name:          &name, // string
		Nodes:         nodes,
		ProjectID:     &projectID, // string
		Tags:          tags,
		Title:         title,
	}
}

func ZksInstanceModelFromMap(m map[string]interface{}) *models.ZksInstance {
	clusterPrefix := m["cluster_prefix"].(string)
	description := m["description"].(string)
	id := m["id"].(string)
	isImported := m["is_imported"].(bool)
	name := m["name"].(string)
	var nodes []*models.ZksInstanceNode // []*ZksInstanceNode
	nodesInterface, nodesIsSet := m["nodes"]
	if nodesIsSet {
		var items []interface{}
		if listItems, isList := nodesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = nodesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ZksInstanceNodeModelFromMap(v.(map[string]interface{}))
			nodes = append(nodes, m)
		}
	}
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
	return &models.ZksInstance{
		ClusterPrefix: clusterPrefix,
		Description:   description,
		ID:            id,
		IsImported:    isImported,
		Name:          &name,
		Nodes:         nodes,
		ProjectID:     &projectID,
		Tags:          tags,
		Title:         title,
	}
}

func SetZksInstanceResourceData(d *schema.ResourceData, m *models.ZksInstance) {
	d.Set("cluster_prefix", m.ClusterPrefix)
	d.Set("description", m.Description)
	d.Set("id", m.ID)
	d.Set("is_imported", m.IsImported)
	d.Set("name", m.Name)
	d.Set("nodes", SetZksInstanceNodeSubResourceData(m.Nodes))
	d.Set("project_id", m.ProjectID)
	d.Set("tags", m.Tags)
	d.Set("title", m.Title)
}

func SetZksInstanceSubResourceData(m []*models.ZksInstance) (d []*map[string]interface{}) {
	for _, ZksInstanceModel := range m {
		if ZksInstanceModel != nil {
			properties := make(map[string]interface{})
			properties["cluster_prefix"] = ZksInstanceModel.ClusterPrefix
			properties["description"] = ZksInstanceModel.Description
			properties["id"] = ZksInstanceModel.ID
			properties["is_imported"] = ZksInstanceModel.IsImported
			properties["name"] = ZksInstanceModel.Name
			properties["nodes"] = SetZksInstanceNodeSubResourceData(ZksInstanceModel.Nodes)
			properties["project_id"] = ZksInstanceModel.ProjectID
			properties["tags"] = ZksInstanceModel.Tags
			properties["title"] = ZksInstanceModel.Title
			d = append(d, &properties)
		}
	}
	return
}

func ZksInstanceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cluster_prefix": {
			Description: `Cluster prefix for the ZKS instance`,
			Type:        schema.TypeString,
			Optional:    true,
		},

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

		"is_imported": {
			Description: `Indicates if the ZKS instance is imported from an external source`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"name": {
			Description: `User defined name of the zks instance, unique across the enterprise. Once zks instance is created, name can’t be changed`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"nodes": {
			Description: `A list of nodes in the zks instance`,
			Type:        schema.TypeSet, //GoType: []*ZksInstanceNode
			Elem: &schema.Resource{
				Schema: ZksInstanceNodeSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
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
			Description: `User defined title of the zks instance. Title can be changed at any time`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetZksInstancePropertyFields() (t []string) {
	return []string{
		"cluster_prefix",
		"description",
		"id",
		"is_imported",
		"name",
		"nodes",
		"project_id",
		"tags",
		"title",
	}
}
