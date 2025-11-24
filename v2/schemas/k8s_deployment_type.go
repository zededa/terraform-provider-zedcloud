package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func K8sDeploymentTypeModel(d *schema.ResourceData) *models.K8sDeploymentType {
	k8sDeploymentType, _ := d.Get("k8s_deployment_type").(models.K8sDeploymentType)
	return &k8sDeploymentType
}

func K8sDeploymentTypeModelFromMap(m map[string]interface{}) *models.K8sDeploymentType {
	k8sDeploymentType := m["k8s_deployment_type"].(models.K8sDeploymentType)
	return &k8sDeploymentType
}

func SetK8sDeploymentTypeResourceData(d *schema.ResourceData, m *models.K8sDeploymentType) {
}

func SetK8sDeploymentTypeSubResourceData(m []*models.K8sDeploymentType) (d []*map[string]interface{}) {
	for _, K8sDeploymentTypeModel := range m {
		if K8sDeploymentTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func K8sDeploymentTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetK8sDeploymentTypePropertyFields() (t []string) {
	return []string{}
}
