package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func TwinDetailModel(d *schema.ResourceData) *models.TwinDetail {
	authenticationType, _ := d.Get("authentication_type").(string)
	cloudToDeviceMessageCountInt, _ := d.Get("cloud_to_device_message_count").(int)
	cloudToDeviceMessageCount := int64(cloudToDeviceMessageCountInt)
	description, _ := d.Get("description").(string)
	desired, _ := d.Get("desired").(strfmt.Base64)
	lastDesiredStatus, _ := d.Get("last_desired_status").(strfmt.Base64)
	moduleCountInt, _ := d.Get("module_count").(int)
	moduleCount := int64(moduleCountInt)
	reported, _ := d.Get("reported").(strfmt.Base64)
	statusCodeInt, _ := d.Get("status_code").(int)
	statusCode := int32(statusCodeInt)
	tags, _ := d.Get("tags").(strfmt.Base64)
	return &models.TwinDetail{
		AuthenticationType:        authenticationType,
		CloudToDeviceMessageCount: cloudToDeviceMessageCount,
		Description:               description,
		Desired:                   desired,
		LastDesiredStatus:         lastDesiredStatus,
		ModuleCount:               moduleCount,
		Reported:                  reported,
		StatusCode:                statusCode,
		Tags:                      tags,
	}
}

func TwinDetailModelFromMap(m map[string]interface{}) *models.TwinDetail {
	authenticationType := m["authentication_type"].(string)
	cloudToDeviceMessageCount := int64(m["cloud_to_device_message_count"].(int)) // int64
	description := m["description"].(string)
	desired := m["desired"].(strfmt.Base64)
	lastDesiredStatus := m["last_desired_status"].(strfmt.Base64)
	moduleCount := int64(m["module_count"].(int)) // int64
	reported := m["reported"].(strfmt.Base64)
	statusCode := int32(m["status_code"].(int)) // int32
	tags := m["tags"].(strfmt.Base64)
	return &models.TwinDetail{
		AuthenticationType:        authenticationType,
		CloudToDeviceMessageCount: cloudToDeviceMessageCount,
		Description:               description,
		Desired:                   desired,
		LastDesiredStatus:         lastDesiredStatus,
		ModuleCount:               moduleCount,
		Reported:                  reported,
		StatusCode:                statusCode,
		Tags:                      tags,
	}
}

// Update the underlying TwinDetail resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetTwinDetailResourceData(d *schema.ResourceData, m *models.TwinDetail) {
	d.Set("authentication_type", m.AuthenticationType)
	d.Set("cloud_to_device_message_count", m.CloudToDeviceMessageCount)
	d.Set("description", m.Description)
	d.Set("desired", m.Desired.String())
	d.Set("last_desired_status", m.LastDesiredStatus.String())
	d.Set("module_count", m.ModuleCount)
	d.Set("reported", m.Reported.String())
	d.Set("status_code", m.StatusCode)
	d.Set("tags", m.Tags.String())
}

// Iterate through and update the TwinDetail resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetTwinDetailSubResourceData(m []*models.TwinDetail) (d []*map[string]interface{}) {
	for _, TwinDetailModel := range m {
		if TwinDetailModel != nil {
			properties := make(map[string]interface{})
			properties["authentication_type"] = TwinDetailModel.AuthenticationType
			properties["cloud_to_device_message_count"] = TwinDetailModel.CloudToDeviceMessageCount
			properties["description"] = TwinDetailModel.Description
			properties["desired"] = TwinDetailModel.Desired.String()
			properties["last_desired_status"] = TwinDetailModel.LastDesiredStatus.String()
			properties["module_count"] = TwinDetailModel.ModuleCount
			properties["reported"] = TwinDetailModel.Reported.String()
			properties["status_code"] = TwinDetailModel.StatusCode
			properties["tags"] = TwinDetailModel.Tags.String()
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the TwinDetail resource defined in the Terraform configuration
func TwinDetailSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"authentication_type": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"cloud_to_device_message_count": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"description": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"desired": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"last_desired_status": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"module_count": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"reported": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"status_code": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"tags": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the TwinDetail resource
func GetTwinDetailPropertyFields() (t []string) {
	return []string{
		"authentication_type",
		"cloud_to_device_message_count",
		"description",
		"desired",
		"last_desired_status",
		"module_count",
		"reported",
		"status_code",
		"tags",
	}
}
