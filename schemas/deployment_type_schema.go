package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DeploymentType resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeploymentTypeModel(d *schema.ResourceData) *models.DeploymentType {
	deploymentType := d.Get("deployment_type").(models.DeploymentType)
	return &deploymentType
}

func DeploymentTypeModelFromMap(m map[string]interface{}) *models.DeploymentType {
	deploymentType := m["deployment_type"].(models.DeploymentType)
	return &deploymentType
}

// Update the underlying DeploymentType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeploymentTypeResourceData(d *schema.ResourceData, m *models.DeploymentType) {
}

// Iterate throught and update the DeploymentType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeploymentTypeSubResourceData(m []*models.DeploymentType) (d []*map[string]interface{}) {
	for _, DeploymentTypeModel := range m {
		if DeploymentTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeploymentType resource defined in the Terraform configuration
func DeploymentTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the DeploymentType resource
func GetDeploymentTypePropertyFields() (t []string) {
	return []string{}
}
