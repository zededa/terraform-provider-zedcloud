package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func ContainerDetailModel(d *schema.ResourceData) *models.ContainerDetail {
	containerCreateOption, _ := d.Get("container_create_option").(string)
	return &models.ContainerDetail{
		ContainerCreateOption: containerCreateOption,
	}
}

func ContainerDetailModelFromMap(m map[string]interface{}) *models.ContainerDetail {
	containerCreateOption := m["container_create_option"].(string)
	return &models.ContainerDetail{
		ContainerCreateOption: containerCreateOption,
	}
}

func SetContainerDetailResourceData(d *schema.ResourceData, m *models.ContainerDetail) {
	d.Set("container_create_option", m.ContainerCreateOption)
}

func SetContainerDetailSubResourceData(m []*models.ContainerDetail) (d []*map[string]interface{}) {
	for _, ContainerDetailModel := range m {
		if ContainerDetailModel != nil {
			properties := make(map[string]interface{})
			properties["container_create_option"] = ContainerDetailModel.ContainerCreateOption
			d = append(d, &properties)
		}
	}
	return
}

func ContainerDetail() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"container_create_option": {
			Description: `Create options direct the creation of the Docker container`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetContainerDetailPropertyFields() (t []string) {
	return []string{
		"container_create_option",
	}
}
