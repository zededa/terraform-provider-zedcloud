package schemas

import (
	"time"

	"github.com/go-openapi/strfmt"
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
			Default:     "10.244.244.2/28",
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
			Type:        schema.TypeSet, //GoType: []*ClusterNode
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

		// Once nodes are members of an edge-node cluster, EVE-OS must be
		// upgraded through the cluster rather than per node: the controller
		// rolls the image out one node at a time and migrates workloads as it
		// goes. Setting base_image here is what drives
		// PUT /v1/cluster/id/{id}/upgrade.
		"base_image": {
			Description: `The EVE-OS image the cluster should run. Changing image_name starts a ` +
				`cluster-scoped, rolling upgrade: the controller upgrades the member nodes one at a ` +
				`time, migrating workloads between them. The apply returns as soon as the rollout ` +
				`has been accepted, not when it has finished -- watch upgrade_status for progress. ` +
				`Do NOT set base_image on the zedcloud_edgenode resources of clustered nodes; the ` +
				`controller rejects a per-node base image for a cluster member.`,
			Type:     schema.TypeList, //GoType: []*BaseOSImage
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: ClusterBaseImageSchema(),
			},
			Optional: true,
			// Same reasoning as the node resource: the controller can set a
			// base image out of band (the UI, another rollout), and a config
			// that never mentions base_image must not be dragged into a diff
			// by that. See CI-736 and CI-836.
			DiffSuppressFunc: diffSuppressBaseImage("base_image"),
		},

		"upgrade_status": {
			Description: `Per-node progress of the most recent cluster EVE-OS upgrade, as reported ` +
				`by the cluster reporter.`,
			Type:     schema.TypeList, //GoType: []*EdgeNodeClusterUpgradeStatusRespNode
			Computed: true,
			Elem: &schema.Resource{
				Schema: ClusterUpgradeStatusSchema(),
			},
		},
	}
}

// ClusterBaseImageSchema is the cluster-scoped subset of BaseOSImage. The
// upgrade API only reads imageName and activate; uuid, version and imvolId are
// per-device bookkeeping and have no meaning for a cluster.
func ClusterBaseImageSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"image_name": {
			Description: `Name of the EVE-OS image to roll out across the cluster, e.g. "16.5.0-k-amd64". Must be an IMAGE_TYPE_EVE image in IMAGE_STATUS_READY.`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"activate": {
			Description: `Activate the image as each node receives it. Defaults to true; false stages the image without switching nodes onto it.`,
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
		},
	}
}

func ClusterUpgradeStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"node_id": {
			Description: `Id of the cluster member node`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"upgradeable_eve_os": {
			Description: `The EVE-OS image this node has been asked to move to. Empty until the rollout reaches the node.`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"status": {
			Description: `One of STATUS_UNSPECIFIED, STATUS_IN_PROGRESS, STATUS_COMPLETED, STATUS_FAILED`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"updated_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

// ClusterBaseImageModel builds the PUT /v1/cluster/id/{id}/upgrade body from
// the base_image block. The API takes a plain BaseOSImage; uuid and version are
// required by the swagger definition but ignored by the controller for a
// cluster upgrade, so they go out empty rather than invented.
func ClusterBaseImageModel(d *schema.ResourceData) *models.BaseOSImage {
	raw, isSet := d.GetOk("base_image")
	if !isSet {
		return nil
	}
	items, isList := raw.([]interface{})
	if !isList || len(items) == 0 || items[0] == nil {
		return nil
	}
	m, isMap := items[0].(map[string]interface{})
	if !isMap {
		return nil
	}

	imageName, _ := m["image_name"].(string)
	if imageName == "" {
		return nil
	}
	activate, activateIsSet := m["activate"].(bool)
	if !activateIsSet {
		activate = true
	}
	empty := ""

	return &models.BaseOSImage{
		ImageName: &imageName,
		Activate:  &activate,
		UUID:      &empty,
		Version:   &empty,
	}
}

// SetClusterUpgradeResourceData writes the cluster upgrade status into state,
// and derives base_image from it.
//
// The cluster object itself does not carry the image: the only record of what a
// cluster was asked to run is the per-node upgrade bookkeeping. So the image is
// taken from the most recently updated node row that names one, which is what
// the rollout most recently requested.
func SetClusterUpgradeResourceData(d *schema.ResourceData, m *models.EdgeNodeClusterUpgradeStatusResp) {
	if m == nil {
		return
	}

	d.Set("upgrade_status", SetClusterUpgradeStatusSubResourceData(m.Nodes))

	var imageName string
	var newest strfmt.DateTime
	for _, n := range m.Nodes {
		if n == nil || n.UpgradeableEveOs == "" {
			continue
		}
		if imageName == "" || time.Time(n.UpdatedAt).After(time.Time(newest)) {
			imageName = n.UpgradeableEveOs
			newest = n.UpdatedAt
		}
	}
	if imageName == "" {
		// No rollout has ever been requested for this cluster. Leave whatever
		// the config says alone rather than writing an empty block into state.
		return
	}

	// activate is not reported back by the API, so keep the configured value
	// and fall back to the schema default.
	activate := true
	if raw, isSet := d.GetOk("base_image"); isSet {
		if items, isList := raw.([]interface{}); isList && len(items) > 0 && items[0] != nil {
			if cur, isMap := items[0].(map[string]interface{}); isMap {
				if a, ok := cur["activate"].(bool); ok {
					activate = a
				}
			}
		}
	}

	d.Set("base_image", []interface{}{
		map[string]interface{}{
			"image_name": imageName,
			"activate":   activate,
		},
	})
}

func SetClusterUpgradeStatusSubResourceData(m []*models.EdgeNodeClusterUpgradeStatusRespNode) (d []*map[string]interface{}) {
	for _, node := range m {
		if node == nil {
			continue
		}
		properties := make(map[string]interface{})
		properties["node_id"] = node.NodeID
		properties["upgradeable_eve_os"] = node.UpgradeableEveOs
		if node.Status != nil {
			properties["status"] = string(*node.Status)
		} else {
			properties["status"] = ""
		}
		properties["created_at"] = node.CreatedAt.String()
		properties["updated_at"] = node.UpdatedAt.String()
		d = append(d, &properties)
	}
	return
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
		"base_image",
		"upgrade_status",
	}
}
