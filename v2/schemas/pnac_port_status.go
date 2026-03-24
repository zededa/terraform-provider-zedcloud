package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func PNACPortStatusModelFromMap(m map[string]interface{}) *models.PNACPortStatus {
	enabled := m["enabled"].(bool)

	var err *models.DeviceError
	errInterface, errIsSet := m["err"]
	if errIsSet {
		errList := errInterface.([]interface{})
		if len(errList) > 0 && errList[0] != nil {
			err = DeviceErrorModelFromMap(errList[0].(map[string]interface{}))
		}
	}

	lastAuthTimestamp := m["last_auth_timestamp"].(string)

	var supplicantState *models.PNACSupplicantState
	supplicantStateInterface, supplicantStateIsSet := m["supplicant_state"]
	if supplicantStateIsSet {
		supplicantStateStr := supplicantStateInterface.(string)
		if supplicantStateStr != "" {
			ss := models.PNACSupplicantState(supplicantStateStr)
			supplicantState = &ss
		}
	}

	_ = lastAuthTimestamp // lastAuthTimestamp is a string representation, used for display only

	return &models.PNACPortStatus{
		Enabled:         enabled,
		Err:             err,
		SupplicantState: supplicantState,
	}
}

// SetPNACPortStatusSubResourceData iterates through and updates the PNACPortStatus resource data within a pagination response
func SetPNACPortStatusSubResourceData(m []*models.PNACPortStatus) (d []*map[string]interface{}) {
	for _, PNACPortStatusModel := range m {
		if PNACPortStatusModel != nil {
			properties := make(map[string]interface{})
			properties["enabled"] = PNACPortStatusModel.Enabled
			properties["err"] = SetDeviceErrorSubResourceData([]*models.DeviceError{PNACPortStatusModel.Err})
			properties["last_auth_timestamp"] = PNACPortStatusModel.LastAuthTimestamp.String()
			if PNACPortStatusModel.SupplicantState != nil {
				properties["supplicant_state"] = string(*PNACPortStatusModel.SupplicantState)
			} else {
				properties["supplicant_state"] = ""
			}
			d = append(d, &properties)
		}
	}
	return
}

// PNACPortStatusSchema returns the schema mapping for the PNACPortStatus resource
func PNACPortStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enabled": {
			Description: `Whether 802.1X authentication is enabled on this port`,
			Type:        schema.TypeBool,
			Computed:    true,
		},

		"supplicant_state": {
			Description: `802.1X supplicant state`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"last_auth_timestamp": {
			Description: `Timestamp of the most recent successful 802.1X authentication`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"err": {
			Description: `Supplicant authentication error details`,
			Type:        schema.TypeList,
			Elem: &schema.Resource{
				Schema: DeviceErrorSchema(),
			},
			Computed: true,
		},
	}
}
