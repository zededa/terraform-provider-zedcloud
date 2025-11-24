package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/go-openapi/strfmt"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ClusterGroupStatusModel(d *schema.ResourceData) *models.ClusterGroupStatus {
	age, _ := d.Get("age").(strfmt.DateTime)
	var clusterReadiness *models.ClusterReadiness // ClusterReadiness
	clusterReadinessInterface, clusterReadinessIsSet := d.GetOk("cluster_readiness")
	if clusterReadinessIsSet && clusterReadinessInterface != nil {
		clusterReadinessMap := clusterReadinessInterface.([]interface{})
		if len(clusterReadinessMap) > 0 {
			clusterReadiness = ClusterReadinessModelFromMap(clusterReadinessMap[0].(map[string]interface{}))
		}
	}
	clusterTags := map[string]string{}
	clusterTagsInterface, clusterTagsIsSet := d.GetOk("clusterTags")
	if clusterTagsIsSet {
		clusterTagsMap := clusterTagsInterface.(map[string]interface{})
		for k, v := range clusterTagsMap {
			if v == nil {
				continue
			}
			clusterTags[k] = v.(string)
		}
	}

	var clusters []*models.ClusterGroupCluster // []*ClusterGroupCluster
	clustersInterface, clustersIsSet := d.GetOk("clusters")
	if clustersIsSet {
		var items []interface{}
		if listItems, isList := clustersInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = clustersInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ClusterGroupClusterModelFromMap(v.(map[string]interface{}))
			clusters = append(clusters, m)
		}
	}
	description, _ := d.Get("description").(string)
	id, _ := d.Get("id").(string)
	name, _ := d.Get("name").(string)
	namespace, _ := d.Get("namespace").(string)
	projectID, _ := d.Get("project_id").(string)
	projectName, _ := d.Get("project_name").(string)
	var resources *models.ClusterGroupResources // ClusterGroupResources
	resourcesInterface, resourcesIsSet := d.GetOk("resources")
	if resourcesIsSet && resourcesInterface != nil {
		resourcesMap := resourcesInterface.([]interface{})
		if len(resourcesMap) > 0 {
			resources = ClusterGroupResourcesModelFromMap(resourcesMap[0].(map[string]interface{}))
		}
	}
	var state *models.RunState // RunState
	stateInterface, stateIsSet := d.GetOk("state")
	if stateIsSet {
		stateModel := stateInterface.(string)
		state = models.NewRunState(models.RunState(stateModel))
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
	return &models.ClusterGroupStatus{
		Age:              age,
		ClusterReadiness: clusterReadiness,
		ClusterTags:      clusterTags,
		Clusters:         clusters,
		Description:      description,
		ID:               id,
		Name:             name,
		Namespace:        namespace,
		ProjectID:        projectID,
		ProjectName:      projectName,
		Resources:        resources,
		State:            state,
		Tags:             tags,
		Title:            title,
	}
}

func ClusterGroupStatusModelFromMap(m map[string]interface{}) *models.ClusterGroupStatus {
	age := m["age"].(strfmt.DateTime)
	var clusterReadiness *models.ClusterReadiness // ClusterReadiness
	clusterReadinessInterface, clusterReadinessIsSet := m["cluster_readiness"]
	if clusterReadinessIsSet && clusterReadinessInterface != nil {
		clusterReadinessMap := clusterReadinessInterface.([]interface{})
		if len(clusterReadinessMap) > 0 {
			clusterReadiness = ClusterReadinessModelFromMap(clusterReadinessMap[0].(map[string]interface{}))
		}
	}
	//
	clusterTags := map[string]string{}
	clusterTagsInterface, clusterTagsIsSet := m["cluster_tags"]
	if clusterTagsIsSet {
		clusterTagsMap := clusterTagsInterface.(map[string]interface{})
		for k, v := range clusterTagsMap {
			if v == nil {
				continue
			}
			clusterTags[k] = v.(string)
		}
	}

	var clusters []*models.ClusterGroupCluster // []*ClusterGroupCluster
	clustersInterface, clustersIsSet := m["clusters"]
	if clustersIsSet {
		var items []interface{}
		if listItems, isList := clustersInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = clustersInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ClusterGroupClusterModelFromMap(v.(map[string]interface{}))
			clusters = append(clusters, m)
		}
	}
	description := m["description"].(string)
	id := m["id"].(string)
	name := m["name"].(string)
	namespace := m["namespace"].(string)
	projectID := m["project_id"].(string)
	projectName := m["project_name"].(string)
	var resources *models.ClusterGroupResources // ClusterGroupResources
	resourcesInterface, resourcesIsSet := m["resources"]
	if resourcesIsSet && resourcesInterface != nil {
		resourcesMap := resourcesInterface.([]interface{})
		if len(resourcesMap) > 0 {
			resources = ClusterGroupResourcesModelFromMap(resourcesMap[0].(map[string]interface{}))
		}
	}
	//
	var state *models.RunState // RunState
	stateInterface, stateIsSet := m["state"]
	if stateIsSet {
		stateModel := stateInterface.(string)
		state = models.NewRunState(models.RunState(stateModel))
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
	return &models.ClusterGroupStatus{
		Age:              age,
		ClusterReadiness: clusterReadiness,
		ClusterTags:      clusterTags,
		Clusters:         clusters,
		Description:      description,
		ID:               id,
		Name:             name,
		Namespace:        namespace,
		ProjectID:        projectID,
		ProjectName:      projectName,
		Resources:        resources,
		State:            state,
		Tags:             tags,
		Title:            title,
	}
}

func SetClusterGroupStatusResourceData(d *schema.ResourceData, m *models.ClusterGroupStatus) {
	d.Set("age", m.Age)
	d.Set("cluster_readiness", SetClusterReadinessSubResourceData([]*models.ClusterReadiness{m.ClusterReadiness}))
	d.Set("cluster_tags", m.ClusterTags)
	d.Set("clusters", SetClusterGroupClusterSubResourceData(m.Clusters))
	d.Set("description", m.Description)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("namespace", m.Namespace)
	d.Set("project_id", m.ProjectID)
	d.Set("project_name", m.ProjectName)
	d.Set("resources", SetClusterGroupResourcesSubResourceData([]*models.ClusterGroupResources{m.Resources}))
	d.Set("state", m.State)
	d.Set("tags", m.Tags)
	d.Set("title", m.Title)
}

func SetClusterGroupStatusSubResourceData(m []*models.ClusterGroupStatus) (d []*map[string]interface{}) {
	for _, ClusterGroupStatusModel := range m {
		if ClusterGroupStatusModel != nil {
			properties := make(map[string]interface{})
			properties["age"] = ClusterGroupStatusModel.Age.String()
			properties["cluster_readiness"] = SetClusterReadinessSubResourceData([]*models.ClusterReadiness{ClusterGroupStatusModel.ClusterReadiness})
			properties["cluster_tags"] = ClusterGroupStatusModel.ClusterTags
			properties["clusters"] = SetClusterGroupClusterSubResourceData(ClusterGroupStatusModel.Clusters)
			properties["description"] = ClusterGroupStatusModel.Description
			properties["id"] = ClusterGroupStatusModel.ID
			properties["name"] = ClusterGroupStatusModel.Name
			properties["namespace"] = ClusterGroupStatusModel.Namespace
			properties["project_id"] = ClusterGroupStatusModel.ProjectID
			properties["project_name"] = ClusterGroupStatusModel.ProjectName
			properties["resources"] = SetClusterGroupResourcesSubResourceData([]*models.ClusterGroupResources{ClusterGroupStatusModel.Resources})
			properties["state"] = ClusterGroupStatusModel.State
			properties["tags"] = ClusterGroupStatusModel.Tags
			properties["title"] = ClusterGroupStatusModel.Title
			d = append(d, &properties)
		}
	}
	return
}

func ClusterGroupStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"age": {
			Description:  `Age of the cluster group`,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"cluster_readiness": {
			Description: `Readiness of the clusters in the cluster group`,
			Type:        schema.TypeList, //GoType: ClusterReadiness
			Elem: &schema.Resource{
				Schema: ClusterReadinessSchema(),
			},
			Optional: true,
		},

		"cluster_tags": {
			Description: `Cluster tags are name/value pairs that enable you to categorize clusters`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"clusters": {
			Description: `List of clusters in the cluster group`,
			Type:        schema.TypeList, //GoType: []*ClusterGroupCluster
			Elem: &schema.Resource{
				Schema: ClusterGroupClusterSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

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

		"name": {
			Description: `User defined name of the cluster group, unique across the enterprise. Once cluster group is created, name can’t be changed`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"namespace": {
			Description: `Namespace of the cluster group`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"project_id": {
			Description: `Foreign key to the project`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"project_name": {
			Description: `Name of the project to which this cluster group belongs`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"resources": {
			Description: `Resources in the cluster group`,
			Type:        schema.TypeList, //GoType: ClusterGroupResources
			Elem: &schema.Resource{
				Schema: ClusterGroupResourcesSchema(),
			},
			Optional: true,
		},

		"state": {
			Description: `Run state of the cluster group`,
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

		"title": {
			Description: `User defined title of the cluster group. Title can be changed at any time`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetClusterGroupStatusPropertyFields() (t []string) {
	return []string{
		"age",
		"cluster_readiness",
		"cluster_tags",
		"clusters",
		"description",
		"id",
		"name",
		"namespace",
		"project_id",
		"project_name",
		"resources",
		"state",
		"tags",
		"title",
	}
}
