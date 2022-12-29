package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate ClusterConfig resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ClusterConfigModel(d *schema.ResourceData) *models.ClusterConfig {
	minNodesRequiredInt, _ := d.Get("min_nodes_required").(int)
	minNodesRequired := int64(minNodesRequiredInt)
	return &models.ClusterConfig{
		MinNodesRequired: minNodesRequired,
	}
}

func ClusterConfigModelFromMap(m map[string]interface{}) *models.ClusterConfig {
	minNodesRequired := int64(m["min_nodes_required"].(int)) // int64 false false false
	return &models.ClusterConfig{
		MinNodesRequired: minNodesRequired,
	}
}

// Update the underlying ClusterConfig resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetClusterConfigResourceData(d *schema.ResourceData, m *models.ClusterConfig) {
	d.Set("min_nodes_required", m.MinNodesRequired)
}

// Iterate through and update the ClusterConfig resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetClusterConfigSubResourceData(m []*models.ClusterConfig) (d []*map[string]interface{}) {
	for _, ClusterConfigModel := range m {
		if ClusterConfigModel != nil {
			properties := make(map[string]interface{})
			properties["min_nodes_required"] = ClusterConfigModel.MinNodesRequired
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ClusterConfig resource defined in the Terraform configuration
func ClusterConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"min_nodes_required": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the ClusterConfig resource
func GetClusterConfigPropertyFields() (t []string) {
	return []string{
		"min_nodes_required",
	}
}
