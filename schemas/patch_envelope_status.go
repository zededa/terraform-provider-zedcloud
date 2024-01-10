package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func PatchEnvelopeStatusModel(d *schema.ResourceData) *models.PatchEnvelopeStatus {
	deviceID, _ := d.Get("device_id").(string)
	deviceName, _ := d.Get("device_name").(string)
	var errors []string
	errorsInterface, errorsIsSet := d.GetOk("errors")
	if errorsIsSet {
		var items []interface{}
		if listItems, isList := errorsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = errorsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			errors = append(errors, v.(string))
		}
	}
	id, _ := d.Get("id").(string)
	name, _ := d.Get("name").(string)
	projectID, _ := d.Get("project_id").(string)
	projectName, _ := d.Get("project_name").(string)
	var state *models.PatchEnvelopeState // PatchEnvelopeState
	stateInterface, stateIsSet := d.GetOk("state")
	if stateIsSet {
		stateModel := stateInterface.(string)
		state = models.NewPatchEnvelopeState(models.PatchEnvelopeState(stateModel))
	}
	userDefinedVersion, _ := d.Get("user_defined_version").(string)
	return &models.PatchEnvelopeStatus{
		DeviceID:           deviceID,
		DeviceName:         deviceName,
		Errors:             errors,
		ID:                 id,
		Name:               name,
		ProjectID:          projectID,
		ProjectName:        projectName,
		State:              state,
		UserDefinedVersion: userDefinedVersion,
	}
}

func PatchEnvelopeStatusModelFromMap(m map[string]interface{}) *models.PatchEnvelopeStatus {
	deviceID := m["device_id"].(string)
	deviceName := m["device_name"].(string)
	var errors []string
	errorsInterface, errorsIsSet := m["errors"]
	if errorsIsSet {
		var items []interface{}
		if listItems, isList := errorsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = errorsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			errors = append(errors, v.(string))
		}
	}
	id := m["id"].(string)
	name := m["name"].(string)
	projectID := m["project_id"].(string)
	projectName := m["project_name"].(string)
	var state *models.PatchEnvelopeState // PatchEnvelopeState
	stateInterface, stateIsSet := m["state"]
	if stateIsSet {
		stateModel := stateInterface.(string)
		state = models.NewPatchEnvelopeState(models.PatchEnvelopeState(stateModel))
	}
	userDefinedVersion := m["user_defined_version"].(string)
	return &models.PatchEnvelopeStatus{
		DeviceID:           deviceID,
		DeviceName:         deviceName,
		Errors:             errors,
		ID:                 id,
		Name:               name,
		ProjectID:          projectID,
		ProjectName:        projectName,
		State:              state,
		UserDefinedVersion: userDefinedVersion,
	}
}

func SetPatchEnvelopeStatusResourceData(d *schema.ResourceData, m *models.PatchEnvelopeStatus) {
	d.Set("device_id", m.DeviceID)
	d.Set("device_name", m.DeviceName)
	d.Set("errors", m.Errors)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("project_id", m.ProjectID)
	d.Set("project_name", m.ProjectName)
	d.Set("state", m.State)
	d.Set("user_defined_version", m.UserDefinedVersion)
}

func SetPatchEnvelopeStatusSubResourceData(m []*models.PatchEnvelopeStatus) (d []*map[string]interface{}) {
	for _, PatchEnvelopeStatusModel := range m {
		if PatchEnvelopeStatusModel != nil {
			properties := make(map[string]interface{})
			properties["device_id"] = PatchEnvelopeStatusModel.DeviceID
			properties["device_name"] = PatchEnvelopeStatusModel.DeviceName
			properties["errors"] = PatchEnvelopeStatusModel.Errors
			properties["id"] = PatchEnvelopeStatusModel.ID
			properties["name"] = PatchEnvelopeStatusModel.Name
			properties["project_id"] = PatchEnvelopeStatusModel.ProjectID
			properties["project_name"] = PatchEnvelopeStatusModel.ProjectName
			properties["state"] = PatchEnvelopeStatusModel.State
			properties["user_defined_version"] = PatchEnvelopeStatusModel.UserDefinedVersion
			d = append(d, &properties)
		}
	}
	return
}

func PatchEnvelopeStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"device_id": {
			Description: `device id`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"device_name": {
			Description: `device name`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"errors": {
			Description: `Patch envelope error`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"id": {
			Description: `patch envelope id`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: `User defined name of the patch envelope, unique across the enterprise. Once patch envelope is created, name can’t be changed`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"project_id": {
			Description: `project id`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"project_name": {
			Description: `project name`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"state": {
			Description: `Patch envelope state`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"user_defined_version": {
			Description: `User defined version for the given patch envelope`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetPatchEnvelopeStatusPropertyFields() (t []string) {
	return []string{
		"device_id",
		"device_name",
		"errors",
		"id",
		"name",
		"project_id",
		"project_name",
		"state",
		"user_defined_version",
	}
}
