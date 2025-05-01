package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func DockerRuntimeProtocolVersionModel(d *schema.ResourceData) *models.DockerRuntimeProtocolVersion {
	dockerRuntimeProtocolVersion, _ := d.Get("docker_runtime_protocol_version").(models.DockerRuntimeProtocolVersion)
	return &dockerRuntimeProtocolVersion
}

func DockerRuntimeProtocolVersionModelFromMap(m map[string]interface{}) *models.DockerRuntimeProtocolVersion {
	dockerRuntimeProtocolVersion := m["docker_runtime_protocol_version"].(models.DockerRuntimeProtocolVersion)
	return &dockerRuntimeProtocolVersion
}

func SetDockerRuntimeProtocolVersionResourceData(d *schema.ResourceData, m *models.DockerRuntimeProtocolVersion) {
}

func SetDockerRuntimeProtocolVersionSubResourceData(m []*models.DockerRuntimeProtocolVersion) (d []*map[string]interface{}) {
	for _, DockerRuntimeProtocolVersionModel := range m {
		if DockerRuntimeProtocolVersionModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func DockerRuntimeProtocolVersionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetDockerRuntimeProtocolVersionPropertyFields() (t []string) {
	return []string{}
}
