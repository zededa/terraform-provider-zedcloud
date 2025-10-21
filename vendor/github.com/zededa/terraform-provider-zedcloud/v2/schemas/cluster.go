package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ClusterModel(d *schema.ResourceData) *models.Cluster {
	clusterPrefix, _ := d.Get("cluster_prefix").(string)
	description, _ := d.Get("description").(string)
	id, _ := d.Get("id").(string)
	name, _ := d.Get("name").(string)
	var nodes []*models.ClusterNode // []*ClusterNode
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
			m := ClusterNodeModelFromMap(v.(map[string]interface{}))
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
	return &models.Cluster{
		ClusterPrefix: clusterPrefix,
		Description:   description,
		ID:            id,
		Name:          &name, // string
		Nodes:         nodes,
		ProjectID:     &projectID, // string
		Tags:          tags,
		Title:         title,
	}
}

func ClusterModelFromMap(m map[string]interface{}) *models.Cluster {
	clusterPrefix := m["cluster_prefix"].(string)
	description := m["description"].(string)
	id := m["id"].(string)
	name := m["name"].(string)
	var nodes []*models.ClusterNode // []*ClusterNode
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
			m := ClusterNodeModelFromMap(v.(map[string]interface{}))
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
	return &models.Cluster{
		ClusterPrefix: clusterPrefix,
		Description:   description,
		ID:            id,
		Name:          &name,
		Nodes:         nodes,
		ProjectID:     &projectID,
		Tags:          tags,
		Title:         title,
	}
}

func SetClusterResourceData(d *schema.ResourceData, m *models.ClusterConfigSummary) {
	d.Set("cluster_prefix", m.ClusterPrefix)
	d.Set("description", m.Description)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("nodes", SetClusterNodeSubResourceData(m.Nodes))
	d.Set("project_id", m.ProjectID)
	d.Set("tags", m.Tags)
	d.Set("title", m.Title)
}

func SetClusterSubResourceData(m []*models.ClusterConfigSummary) (d []*map[string]interface{}) {
	for _, ClusterModel := range m {
		if ClusterModel != nil {
			properties := make(map[string]interface{})
			properties["cluster_prefix"] = ClusterModel.ClusterPrefix
			properties["description"] = ClusterModel.Description
			properties["id"] = ClusterModel.ID
			properties["name"] = ClusterModel.Name
			properties["nodes"] = SetClusterNodeSubResourceData(ClusterModel.Nodes)
			properties["project_id"] = ClusterModel.ProjectID
			properties["tags"] = ClusterModel.Tags
			properties["title"] = ClusterModel.Title
			d = append(d, &properties)
		}
	}
	return
}

func ClusterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cluster_prefix": {
			Description: `A cluster prefix. The default is: 10.244.244.2/28`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"description": {
			Description: `Detailed description of the cluster`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: `System defined universally unique Id of the cluster`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: `User defined name of the cluster, unique across the enterprise. Once cluster is created, name can’t be changed`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"nodes": {
			Description: `A list of nodes in the cluster`,
			Type:        schema.TypeList, //GoType: []*ClusterNode
			Elem: &schema.Resource{
				Schema: ClusterNodeSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Required: true,
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
			Description: `User defined title of the cluster. Title can be changed at any time`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetClusterPropertyFields() (t []string) {
	return []string{
		"cluster_prefix",
		"description",
		"id",
		"name",
		"nodes",
		"project_id",
		"tags",
		"title",
	}
}
