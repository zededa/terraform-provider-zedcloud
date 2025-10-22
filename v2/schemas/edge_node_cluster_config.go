package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func EdgeNodeClusterConfigModel(d *schema.ResourceData) *models.EdgeNodeClusterConfig {
	clusterPrefix, _ := d.Get("cluster_prefix").(string)
	id, _ := d.Get("id").(string)
	isMaster, _ := d.Get("is_master").(bool)
	manifest, _ := d.Get("manifest").(strfmt.Base64)
	name, _ := d.Get("name").(string)
	projectID, _ := d.Get("project_id").(string)
	seedNodeID, _ := d.Get("seed_node_id").(string)
	seedNodeIP, _ := d.Get("seed_node_ip").(string)
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

	tieBreakerNodeID, _ := d.Get("tie_breaker_node_id").(string)
	token, _ := d.Get("token").(string)
	return &models.EdgeNodeClusterConfig{
		ClusterPrefix:    clusterPrefix,
		ID:               id,
		IsMaster:         isMaster,
		Manifest:         manifest,
		Name:             name,
		ProjectID:        projectID,
		SeedNodeID:       seedNodeID,
		SeedNodeIP:       seedNodeIP,
		Tags:             tags,
		TieBreakerNodeID: tieBreakerNodeID,
		Token:            token,
	}
}

func EdgeNodeClusterConfigModelFromMap(m map[string]interface{}) *models.EdgeNodeClusterConfig {
	clusterPrefix := m["cluster_prefix"].(string)
	id := m["id"].(string)
	isMaster := m["is_master"].(bool)
	manifest := m["manifest"].(strfmt.Base64)
	name := m["name"].(string)
	projectID := m["project_id"].(string)
	seedNodeID := m["seed_node_id"].(string)
	seedNodeIP := m["seed_node_ip"].(string)
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

	tieBreakerNodeID := m["tie_breaker_node_id"].(string)
	token := m["token"].(string)
	return &models.EdgeNodeClusterConfig{
		ClusterPrefix:    clusterPrefix,
		ID:               id,
		IsMaster:         isMaster,
		Manifest:         manifest,
		Name:             name,
		ProjectID:        projectID,
		SeedNodeID:       seedNodeID,
		SeedNodeIP:       seedNodeIP,
		Tags:             tags,
		TieBreakerNodeID: tieBreakerNodeID,
		Token:            token,
	}
}

func SetEdgeNodeClusterConfigResourceData(d *schema.ResourceData, m *models.EdgeNodeClusterConfig) {
	d.Set("cluster_prefix", m.ClusterPrefix)
	d.Set("id", m.ID)
	d.Set("is_master", m.IsMaster)
	d.Set("manifest", m.Manifest.String())
	d.Set("name", m.Name)
	d.Set("project_id", m.ProjectID)
	d.Set("seed_node_id", m.SeedNodeID)
	d.Set("seed_node_ip", m.SeedNodeIP)
	d.Set("tags", m.Tags)
	d.Set("tie_breaker_node_id", m.TieBreakerNodeID)
	d.Set("token", m.Token)
}

func SetEdgeNodeClusterConfigSubResourceData(m []*models.EdgeNodeClusterConfig) (d []*map[string]interface{}) {
	for _, EdgeNodeClusterConfigModel := range m {
		if EdgeNodeClusterConfigModel != nil {
			properties := make(map[string]interface{})
			properties["cluster_prefix"] = EdgeNodeClusterConfigModel.ClusterPrefix
			properties["id"] = EdgeNodeClusterConfigModel.ID
			properties["is_master"] = EdgeNodeClusterConfigModel.IsMaster
			properties["manifest"] = EdgeNodeClusterConfigModel.Manifest.String()
			properties["name"] = EdgeNodeClusterConfigModel.Name
			properties["project_id"] = EdgeNodeClusterConfigModel.ProjectID
			properties["seed_node_id"] = EdgeNodeClusterConfigModel.SeedNodeID
			properties["seed_node_ip"] = EdgeNodeClusterConfigModel.SeedNodeIP
			properties["tags"] = EdgeNodeClusterConfigModel.Tags
			properties["tie_breaker_node_id"] = EdgeNodeClusterConfigModel.TieBreakerNodeID
			properties["token"] = EdgeNodeClusterConfigModel.Token
			d = append(d, &properties)
		}
	}
	return
}

func EdgeNodeClusterConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cluster_prefix": {
			Description: `A cluster prefix`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: `System defined universally unique Id of the cluster`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"is_master": {
			Description: `Is it a master node?`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"manifest": {
			Description: `ZKS instance manifest`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"name": {
			Description: `User defined name of the cluster, unique across the enterprise. Once cluster is created, name can’t be changed`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"project_id": {
			Description: `Foreign key to the project`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"seed_node_id": {
			Description: `Seed node id`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"seed_node_ip": {
			Description: `Seed node ip`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"tags": {
			Description: `Tags are name/value pairs that enable you to categorize resources. Tag names are case insensitive with max_length 512 and min_length 3. Tag values are case sensitive with max_length 256 and min_length 3.`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"tie_breaker_node_id": {
			Description: `Tiebreaker node id`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"token": {
			Description: `Cluster token`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetEdgeNodeClusterConfigPropertyFields() (t []string) {
	return []string{
		"cluster_prefix",
		"id",
		"is_master",
		"manifest",
		"name",
		"project_id",
		"seed_node_id",
		"seed_node_ip",
		"tags",
		"tie_breaker_node_id",
		"token",
	}
}
