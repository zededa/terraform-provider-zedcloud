package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DevDataSecAtRest resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DevDataSecAtRestModel(d *schema.ResourceData) *models.DevDataSecAtRest {
	var errInfo *models.DeviceError // DeviceError
	errInfoInterface, errInfoIsSet := d.GetOk("err_info")
	if errInfoIsSet {
		errInfoMap := errInfoInterface.([]interface{})[0].(map[string]interface{})
		errInfo = DeviceErrorModelFromMap(errInfoMap)
	}
	name, _ := d.Get("name").(string)
	statusModel, _ := d.Get("status").(models.DeviceDataSecurityAtRestStatus) // DeviceDataSecurityAtRestStatus
	status := &statusModel
	if !ok {
		status = nil
	}
	return &models.DevDataSecAtRest{
		ErrInfo: errInfo,
		Name:    name,
		Status:  status,
	}
}

func DevDataSecAtRestModelFromMap(m map[string]interface{}) *models.DevDataSecAtRest {
	var errInfo *models.DeviceError // DeviceError
	errInfoInterface, errInfoIsSet := m["err_info"]
	if errInfoIsSet {
		errInfoMap := errInfoInterface.([]interface{})[0].(map[string]interface{})
		errInfo = DeviceErrorModelFromMap(errInfoMap)
	}
	//
	name := m["name"].(string)
	status := m["status"].(*models.DeviceDataSecurityAtRestStatus) // DeviceDataSecurityAtRestStatus
	return &models.DevDataSecAtRest{
		ErrInfo: errInfo,
		Name:    name,
		Status:  status,
	}
}

// Update the underlying DevDataSecAtRest resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDevDataSecAtRestResourceData(d *schema.ResourceData, m *models.DevDataSecAtRest) {
	d.Set("err_info", SetDeviceErrorSubResourceData([]*models.DeviceError{m.ErrInfo}))
	d.Set("name", m.Name)
	d.Set("status", m.Status)
}

// Iterate through and update the DevDataSecAtRest resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDevDataSecAtRestSubResourceData(m []*models.DevDataSecAtRest) (d []*map[string]interface{}) {
	for _, DevDataSecAtRestModel := range m {
		if DevDataSecAtRestModel != nil {
			properties := make(map[string]interface{})
			properties["err_info"] = SetDeviceErrorSubResourceData([]*models.DeviceError{DevDataSecAtRestModel.ErrInfo})
			properties["name"] = DevDataSecAtRestModel.Name
			properties["status"] = DevDataSecAtRestModel.Status
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DevDataSecAtRest resource defined in the Terraform configuration
func DevDataSecAtRestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"err_info": {
			Description: ``,
			Type:        schema.TypeList, //GoType: DeviceError
			Elem: &schema.Resource{
				Schema: DeviceErrorSchema(),
			},
			Optional: true,
		},

		"name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"status": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the DevDataSecAtRest resource
func GetDevDataSecAtRestPropertyFields() (t []string) {
	return []string{
		"err_info",
		"name",
		"status",
	}
}
