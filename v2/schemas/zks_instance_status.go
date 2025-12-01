package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/go-openapi/strfmt"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZksInstanceStatusModel(d *schema.ResourceData) *models.ZksInstanceStatus {
	architecture, _ := d.Get("architecture").(string)
	var capacity *models.ClusterCapacity // ClusterCapacity
	capacityInterface, capacityIsSet := d.GetOk("capacity")
	if capacityIsSet && capacityInterface != nil {
		capacityMap := capacityInterface.([]interface{})
		if len(capacityMap) > 0 {
			capacity = ClusterCapacityModelFromMap(capacityMap[0].(map[string]interface{}))
		}
	}
	clusterPrefix, _ := d.Get("cluster_prefix").(string)
	created, _ := d.Get("created").(strfmt.DateTime)
	deploymentsInt, _ := d.Get("deployments").(int)
	deployments := int32(deploymentsInt)
	description, _ := d.Get("description").(string)
	fleetWorkspaceName, _ := d.Get("fleet_workspace_name").(string)
	id, _ := d.Get("id").(string)
	isImported, _ := d.Get("is_imported").(bool)
	kubeVersion, _ := d.Get("kube_version").(string)
	nOfNodesInt, _ := d.Get("n_of_nodes").(int)
	nOfNodes := int32(nOfNodesInt)
	name, _ := d.Get("name").(string)
	projectID, _ := d.Get("project_id").(string)
	projectName, _ := d.Get("project_name").(string)
	provider, _ := d.Get("provider").(string)
	var runState *models.RunState // RunState
	runStateInterface, runStateIsSet := d.GetOk("run_state")
	if runStateIsSet {
		runStateModel := runStateInterface.(string)
		runState = models.NewRunState(models.RunState(runStateModel))
	}
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
	totalNumberOfNodesInt, _ := d.Get("total_number_of_nodes").(int)
	totalNumberOfNodes := int32(totalNumberOfNodesInt)
	totalResourcesInt, _ := d.Get("total_resources").(int)
	totalResources := int32(totalResourcesInt)
	zksBaseURL, _ := d.Get("zks_base_url").(string)
	zksClusterID, _ := d.Get("zks_cluster_id").(string)
	return &models.ZksInstanceStatus{
		Architecture:       architecture,
		Capacity:           capacity,
		ClusterPrefix:      clusterPrefix,
		Created:            created,
		Deployments:        deployments,
		Description:        description,
		FleetWorkspaceName: fleetWorkspaceName,
		ID:                 id,
		IsImported:         isImported,
		KubeVersion:        kubeVersion,
		NOfNodes:           nOfNodes,
		Name:               name,
		ProjectID:          projectID,
		ProjectName:        projectName,
		Provider:           provider,
		RunState:           runState,
		Tags:               tags,
		Title:              title,
		TotalNumberOfNodes: totalNumberOfNodes,
		TotalResources:     totalResources,
		ZksBaseURL:         zksBaseURL,
		ZksClusterID:       zksClusterID,
	}
}

func ZksInstanceStatusModelFromMap(m map[string]interface{}) *models.ZksInstanceStatus {
	architecture := m["architecture"].(string)
	var capacity *models.ClusterCapacity // ClusterCapacity
	capacityInterface, capacityIsSet := m["capacity"]
	if capacityIsSet && capacityInterface != nil {
		capacityMap := capacityInterface.([]interface{})
		if len(capacityMap) > 0 {
			capacity = ClusterCapacityModelFromMap(capacityMap[0].(map[string]interface{}))
		}
	}
	//
	clusterPrefix := m["cluster_prefix"].(string)
	created := m["created"].(strfmt.DateTime)
	deployments := int32(m["deployments"].(int)) // int32
	description := m["description"].(string)
	fleetWorkspaceName := m["fleet_workspace_name"].(string)
	id := m["id"].(string)
	isImported := m["is_imported"].(bool)
	kubeVersion := m["kube_version"].(string)
	nOfNodes := int32(m["n_of_nodes"].(int)) // int32
	name := m["name"].(string)
	projectID := m["project_id"].(string)
	projectName := m["project_name"].(string)
	provider := m["provider"].(string)
	var runState *models.RunState // RunState
	runStateInterface, runStateIsSet := m["run_state"]
	if runStateIsSet {
		runStateModel := runStateInterface.(string)
		runState = models.NewRunState(models.RunState(runStateModel))
	}
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
	totalNumberOfNodes := int32(m["total_number_of_nodes"].(int)) // int32
	totalResources := int32(m["total_resources"].(int))           // int32
	zksBaseURL := m["zks_base_url"].(string)
	zksClusterID := m["zks_cluster_id"].(string)
	return &models.ZksInstanceStatus{
		Architecture:       architecture,
		Capacity:           capacity,
		ClusterPrefix:      clusterPrefix,
		Created:            created,
		Deployments:        deployments,
		Description:        description,
		FleetWorkspaceName: fleetWorkspaceName,
		ID:                 id,
		IsImported:         isImported,
		KubeVersion:        kubeVersion,
		NOfNodes:           nOfNodes,
		Name:               name,
		ProjectID:          projectID,
		ProjectName:        projectName,
		Provider:           provider,
		RunState:           runState,
		Tags:               tags,
		Title:              title,
		TotalNumberOfNodes: totalNumberOfNodes,
		TotalResources:     totalResources,
		ZksBaseURL:         zksBaseURL,
		ZksClusterID:       zksClusterID,
	}
}

func SetZksInstanceStatusResourceData(d *schema.ResourceData, m *models.ZksInstanceStatus) {
	d.Set("architecture", m.Architecture)
	d.Set("capacity", SetClusterCapacitySubResourceData([]*models.ClusterCapacity{m.Capacity}))
	d.Set("cluster_prefix", m.ClusterPrefix)
	d.Set("created", m.Created)
	d.Set("deployments", m.Deployments)
	d.Set("description", m.Description)
	d.Set("fleet_workspace_name", m.FleetWorkspaceName)
	d.Set("id", m.ID)
	d.Set("is_imported", m.IsImported)
	d.Set("kube_version", m.KubeVersion)
	d.Set("n_of_nodes", m.NOfNodes)
	d.Set("name", m.Name)
	d.Set("project_id", m.ProjectID)
	d.Set("project_name", m.ProjectName)
	d.Set("provider", m.Provider)
	d.Set("run_state", m.RunState)
	d.Set("tags", m.Tags)
	d.Set("title", m.Title)
	d.Set("total_number_of_nodes", m.TotalNumberOfNodes)
	d.Set("total_resources", m.TotalResources)
	d.Set("zks_base_url", m.ZksBaseURL)
	d.Set("zks_cluster_id", m.ZksClusterID)
}

func SetZksInstanceStatusSubResourceData(m []*models.ZksInstanceStatus) (d []*map[string]interface{}) {
	for _, ZksInstanceStatusModel := range m {
		if ZksInstanceStatusModel != nil {
			properties := make(map[string]interface{})
			properties["architecture"] = ZksInstanceStatusModel.Architecture
			properties["capacity"] = SetClusterCapacitySubResourceData([]*models.ClusterCapacity{ZksInstanceStatusModel.Capacity})
			properties["cluster_prefix"] = ZksInstanceStatusModel.ClusterPrefix
			properties["created"] = ZksInstanceStatusModel.Created.String()
			properties["deployments"] = ZksInstanceStatusModel.Deployments
			properties["description"] = ZksInstanceStatusModel.Description
			properties["fleet_workspace_name"] = ZksInstanceStatusModel.FleetWorkspaceName
			properties["id"] = ZksInstanceStatusModel.ID
			properties["is_imported"] = ZksInstanceStatusModel.IsImported
			properties["kube_version"] = ZksInstanceStatusModel.KubeVersion
			properties["n_of_nodes"] = ZksInstanceStatusModel.NOfNodes
			properties["name"] = ZksInstanceStatusModel.Name
			properties["project_id"] = ZksInstanceStatusModel.ProjectID
			properties["project_name"] = ZksInstanceStatusModel.ProjectName
			properties["provider"] = ZksInstanceStatusModel.Provider
			properties["run_state"] = ZksInstanceStatusModel.RunState
			properties["tags"] = ZksInstanceStatusModel.Tags
			properties["title"] = ZksInstanceStatusModel.Title
			properties["total_number_of_nodes"] = ZksInstanceStatusModel.TotalNumberOfNodes
			properties["total_resources"] = ZksInstanceStatusModel.TotalResources
			properties["zks_base_url"] = ZksInstanceStatusModel.ZksBaseURL
			properties["zks_cluster_id"] = ZksInstanceStatusModel.ZksClusterID
			d = append(d, &properties)
		}
	}
	return
}

func ZksInstanceStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"architecture": {
			Description: `Architecture of the zks instance`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"capacity": {
			Description: `Capacity information of the zks instance`,
			Type:        schema.TypeList, //GoType: ClusterCapacity
			Elem: &schema.Resource{
				Schema: ClusterCapacitySchema(),
			},
			Optional: true,
		},

		"cluster_prefix": {
			Description: `Cluster prefix for the ZKS instance`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"created": {
			Description:  `Creation time of the zks instance`,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"deployments": {
			Description: `Number of deployments in the zks instance`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"description": {
			Description: `Detailed description of the zks instance`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"fleet_workspace_name": {
			Description: `Name of the Fleet workspace associated with the ZKS instance`,
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

		"kube_version": {
			Description: `Kubernetes version of the zks instance`,
			Type:        schema.TypeString,
			Optional:    true,
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

		"project_name": {
			Description: `Name of the project to which the zks instance belongs`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"provider": {
			Description: `Provider of the zks instance`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"run_state": {
			Description: `Current state of the zks instance`,
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

		"total_number_of_nodes": {
			Description: `Total number of nodes in the ZKS instance`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"total_resources": {
			Description: `Total resources in the zks instance`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"zks_base_url": {
			Description: `Base URL of the ZKS instance`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"zks_cluster_id": {
			Description: `ZKS cluster ID`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetZksInstanceStatusPropertyFields() (t []string) {
	return []string{
		"architecture",
		"capacity",
		"cluster_prefix",
		"created",
		"deployments",
		"description",
		"fleet_workspace_name",
		"id",
		"is_imported",
		"kube_version",
		"n_of_nodes",
		"name",
		"project_id",
		"project_name",
		"provider",
		"run_state",
		"tags",
		"title",
		"total_number_of_nodes",
		"total_resources",
		"zks_base_url",
		"zks_cluster_id",
	}
}
