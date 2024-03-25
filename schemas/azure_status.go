package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func AzureStatusModel(d *schema.ResourceData) *models.AzureStatus {
	var azureDevStatus *models.AzureDevStatusDetail // AzureDevStatusDetail
	azureDevStatusInterface, azureDevStatusIsSet := d.GetOk("azure_dev_status")
	if azureDevStatusIsSet && azureDevStatusInterface != nil {
		azureDevStatusMap := azureDevStatusInterface.([]interface{})
		if len(azureDevStatusMap) > 0 {
			azureDevStatus = AzureDevStatusDetailModelFromMap(azureDevStatusMap[0].(map[string]interface{}))
		}
	}
	return &models.AzureStatus{
		AzureDevStatus: azureDevStatus,
	}
}

func AzureStatusModelFromMap(m map[string]interface{}) *models.AzureStatus {
	var azureDevStatus *models.AzureDevStatusDetail // AzureDevStatusDetail
	azureDevStatusInterface, azureDevStatusIsSet := m["azure_dev_status"]
	if azureDevStatusIsSet && azureDevStatusInterface != nil {
		azureDevStatusMap := azureDevStatusInterface.([]interface{})
		if len(azureDevStatusMap) > 0 {
			azureDevStatus = AzureDevStatusDetailModelFromMap(azureDevStatusMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.AzureStatus{
		AzureDevStatus: azureDevStatus,
	}
}

// Update the underlying AzureStatus resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAzureStatusResourceData(d *schema.ResourceData, m *models.AzureStatus) {
	d.Set("azure_dev_status", SetAzureDevStatusDetailSubResourceData([]*models.AzureDevStatusDetail{m.AzureDevStatus}))
}

// Iterate through and update the AzureStatus resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAzureStatusSubResourceData(m []*models.AzureStatus) (d []*map[string]interface{}) {
	for _, AzureStatusModel := range m {
		if AzureStatusModel != nil {
			properties := make(map[string]interface{})
			properties["azure_dev_status"] = SetAzureDevStatusDetailSubResourceData([]*models.AzureDevStatusDetail{AzureStatusModel.AzureDevStatus})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AzureStatus resource defined in the Terraform configuration
func AzureStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"azure_dev_status": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AzureDevStatusDetail
			Elem: &schema.Resource{
				Schema: AzureDevStatusDetailSchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the AzureStatus resource
func GetAzureStatusPropertyFields() (t []string) {
	return []string{
		"azure_dev_status",
	}
}
