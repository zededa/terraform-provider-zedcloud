package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ProfileDeploymentResourceObjectStatusModel(d *schema.ResourceData) *models.ProfileDeploymentResourceObjectStatus {
	id, _ := d.Get("id").(string)
	var objectRunState *models.RunState // RunState
	objectRunStateInterface, objectRunStateIsSet := d.GetOk("object_run_state")
	if objectRunStateIsSet {
		objectRunStateModel := objectRunStateInterface.(string)
		objectRunState = models.NewRunState(models.RunState(objectRunStateModel))
	}
	var objectType *models.ObjectType // ObjectType
	objectTypeInterface, objectTypeIsSet := d.GetOk("object_type")
	if objectTypeIsSet {
		objectTypeModel := objectTypeInterface.(string)
		objectType = models.NewObjectType(models.ObjectType(objectTypeModel))
	}
	return &models.ProfileDeploymentResourceObjectStatus{
		ID:             id,
		ObjectRunState: objectRunState,
		ObjectType:     objectType,
	}
}

func ProfileDeploymentResourceObjectStatusModelFromMap(m map[string]interface{}) *models.ProfileDeploymentResourceObjectStatus {
	id := m["id"].(string)
	var objectRunState *models.RunState // RunState
	objectRunStateInterface, objectRunStateIsSet := m["object_run_state"]
	if objectRunStateIsSet {
		objectRunStateModel := objectRunStateInterface.(string)
		objectRunState = models.NewRunState(models.RunState(objectRunStateModel))
	}
	var objectType *models.ObjectType // ObjectType
	objectTypeInterface, objectTypeIsSet := m["object_type"]
	if objectTypeIsSet {
		objectTypeModel := objectTypeInterface.(string)
		objectType = models.NewObjectType(models.ObjectType(objectTypeModel))
	}
	return &models.ProfileDeploymentResourceObjectStatus{
		ID:             id,
		ObjectRunState: objectRunState,
		ObjectType:     objectType,
	}
}

func SetProfileDeploymentResourceObjectStatusResourceData(d *schema.ResourceData, m *models.ProfileDeploymentResourceObjectStatus) {
	d.Set("id", m.ID)
	d.Set("object_run_state", m.ObjectRunState)
	d.Set("object_type", m.ObjectType)
}

func SetProfileDeploymentResourceObjectStatusSubResourceData(m []*models.ProfileDeploymentResourceObjectStatus) (d []*map[string]interface{}) {
	for _, ProfileDeploymentResourceObjectStatusModel := range m {
		if ProfileDeploymentResourceObjectStatusModel != nil {
			properties := make(map[string]interface{})
			properties["id"] = ProfileDeploymentResourceObjectStatusModel.ID
			properties["object_run_state"] = ProfileDeploymentResourceObjectStatusModel.ObjectRunState
			properties["object_type"] = ProfileDeploymentResourceObjectStatusModel.ObjectType
			d = append(d, &properties)
		}
	}
	return
}

func ProfileDeploymentResourceObjectStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Description: `Unique ID of the object.`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"object_run_state": {
			Description: `object run state`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"object_type": {
			Description: `object type`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetProfileDeploymentResourceObjectStatusPropertyFields() (t []string) {
	return []string{
		"id",
		"object_run_state",
		"object_type",
	}
}
