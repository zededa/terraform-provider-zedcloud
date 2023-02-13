package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate ClusterType resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ClusterTypeModel(d *schema.ResourceData) *models.ClusterType {
	clusterType, _ := d.Get("cluster_type").(models.ClusterType)
	return &clusterType
}

func ClusterTypeModelFromMap(m map[string]interface{}) *models.ClusterType {
	clusterType := m["cluster_type"].(models.ClusterType)
	return &clusterType
}

// Update the underlying ClusterType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetClusterTypeResourceData(d *schema.ResourceData, m *models.ClusterType) {
}

// Iterate through and update the ClusterType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetClusterTypeSubResourceData(m []*models.ClusterType) (d []*map[string]interface{}) {
	for _, ClusterTypeModel := range m {
		if ClusterTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ClusterType resource defined in the Terraform configuration
func ClusterTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the ClusterType resource
func GetClusterTypePropertyFields() (t []string) {
	return []string{}
}
