package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func DockerRuntimeVersionModel(d *schema.ResourceData) *models.DockerRuntimeVersion {
	dockerRuntimeVersion, _ := d.Get("docker_runtime_version").(models.DockerRuntimeVersion)
	return &dockerRuntimeVersion
}

func DockerRuntimeVersionModelFromMap(m map[string]interface{}) *models.DockerRuntimeVersion {
	dockerRuntimeVersion := m["docker_runtime_version"].(models.DockerRuntimeVersion)
	return &dockerRuntimeVersion
}

func SetDockerRuntimeVersionResourceData(d *schema.ResourceData, m *models.DockerRuntimeVersion) {
}

func SetDockerRuntimeVersionSubResourceData(m []*models.DockerRuntimeVersion) (d []*map[string]interface{}) {
	for _, DockerRuntimeVersionModel := range m {
		if DockerRuntimeVersionModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func DockerRuntimeVersionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetDockerRuntimeVersionPropertyFields() (t []string) {
	return []string{}
}
