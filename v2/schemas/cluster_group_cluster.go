package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/go-openapi/strfmt"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ClusterGroupClusterModel(d *schema.ResourceData) *models.ClusterGroupCluster {
	age, _ := d.Get("age").(strfmt.DateTime)
	bundleReady, _ := d.Get("bundle_ready").(string)
	lastSeen, _ := d.Get("last_seen").(strfmt.DateTime)
	name, _ := d.Get("name").(string)
	reposReady, _ := d.Get("repos_ready").(string)
	resourcesInt, _ := d.Get("resources").(int)
	resources := int64(resourcesInt)
	var state *models.RunState // RunState
	stateInterface, stateIsSet := d.GetOk("state")
	if stateIsSet {
		stateModel := stateInterface.(string)
		state = models.NewRunState(models.RunState(stateModel))
	}
	return &models.ClusterGroupCluster{
		Age:         age,
		BundleReady: bundleReady,
		LastSeen:    lastSeen,
		Name:        name,
		ReposReady:  reposReady,
		Resources:   resources,
		State:       state,
	}
}

func ClusterGroupClusterModelFromMap(m map[string]interface{}) *models.ClusterGroupCluster {
	age := m["age"].(strfmt.DateTime)
	bundleReady := m["bundle_ready"].(string)
	lastSeen := m["last_seen"].(strfmt.DateTime)
	name := m["name"].(string)
	reposReady := m["repos_ready"].(string)
	resources := int64(m["resources"].(int)) // int64
	var state *models.RunState               // RunState
	stateInterface, stateIsSet := m["state"]
	if stateIsSet {
		stateModel := stateInterface.(string)
		state = models.NewRunState(models.RunState(stateModel))
	}
	return &models.ClusterGroupCluster{
		Age:         age,
		BundleReady: bundleReady,
		LastSeen:    lastSeen,
		Name:        name,
		ReposReady:  reposReady,
		Resources:   resources,
		State:       state,
	}
}

func SetClusterGroupClusterResourceData(d *schema.ResourceData, m *models.ClusterGroupCluster) {
	d.Set("age", m.Age)
	d.Set("bundle_ready", m.BundleReady)
	d.Set("last_seen", m.LastSeen)
	d.Set("name", m.Name)
	d.Set("repos_ready", m.ReposReady)
	d.Set("resources", m.Resources)
	d.Set("state", m.State)
}

func SetClusterGroupClusterSubResourceData(m []*models.ClusterGroupCluster) (d []*map[string]interface{}) {
	for _, ClusterGroupClusterModel := range m {
		if ClusterGroupClusterModel != nil {
			properties := make(map[string]interface{})
			properties["age"] = ClusterGroupClusterModel.Age.String()
			properties["bundle_ready"] = ClusterGroupClusterModel.BundleReady
			properties["last_seen"] = ClusterGroupClusterModel.LastSeen.String()
			properties["name"] = ClusterGroupClusterModel.Name
			properties["repos_ready"] = ClusterGroupClusterModel.ReposReady
			properties["resources"] = ClusterGroupClusterModel.Resources
			properties["state"] = ClusterGroupClusterModel.State
			d = append(d, &properties)
		}
	}
	return
}

func ClusterGroupClusterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"age": {
			Description:  `Age of the cluster in the cluster group`,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"bundle_ready": {
			Description: `Bundle readiness of the cluster in the cluster group`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"last_seen": {
			Description:  `Last seen timestamp of the cluster in the cluster group`,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"name": {
			Description: `Name of the cluster in the cluster group`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"repos_ready": {
			Description: `Repository readiness of the cluster in the cluster group`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"resources": {
			Description: `Number of resources in the cluster in the cluster group`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"state": {
			Description: `State of the cluster in the cluster group`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetClusterGroupClusterPropertyFields() (t []string) {
	return []string{
		"age",
		"bundle_ready",
		"last_seen",
		"name",
		"repos_ready",
		"resources",
		"state",
	}
}
