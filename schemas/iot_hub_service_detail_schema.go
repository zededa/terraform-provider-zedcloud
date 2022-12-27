package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate IotHubServiceDetail resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func IotHubServiceDetailModel(d *schema.ResourceData) *models.IotHubServiceDetail {
	var serviceDetail *models.AzureResourceAndServiceDetail // AzureResourceAndServiceDetail
	serviceDetailInterface, serviceDetailIsSet := d.GetOk("service_detail")
	if serviceDetailIsSet {
		serviceDetailMap := serviceDetailInterface.([]interface{})[0].(map[string]interface{})
		serviceDetail = AzureResourceAndServiceDetailModelFromMap(serviceDetailMap)
	}
	return &models.IotHubServiceDetail{
		ServiceDetail: serviceDetail,
	}
}

func IotHubServiceDetailModelFromMap(m map[string]interface{}) *models.IotHubServiceDetail {
	var serviceDetail *models.AzureResourceAndServiceDetail // AzureResourceAndServiceDetail
	serviceDetailInterface, serviceDetailIsSet := m["service_detail"]
	if serviceDetailIsSet {
		serviceDetailMap := serviceDetailInterface.([]interface{})[0].(map[string]interface{})
		serviceDetail = AzureResourceAndServiceDetailModelFromMap(serviceDetailMap)
	}
	//
	return &models.IotHubServiceDetail{
		ServiceDetail: serviceDetail,
	}
}

// Update the underlying IotHubServiceDetail resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetIotHubServiceDetailResourceData(d *schema.ResourceData, m *models.IotHubServiceDetail) {
	d.Set("service_detail", SetAzureResourceAndServiceDetailSubResourceData([]*models.AzureResourceAndServiceDetail{m.ServiceDetail}))
}

// Iterate throught and update the IotHubServiceDetail resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetIotHubServiceDetailSubResourceData(m []*models.IotHubServiceDetail) (d []*map[string]interface{}) {
	for _, IotHubServiceDetailModel := range m {
		if IotHubServiceDetailModel != nil {
			properties := make(map[string]interface{})
			properties["service_detail"] = SetAzureResourceAndServiceDetailSubResourceData([]*models.AzureResourceAndServiceDetail{IotHubServiceDetailModel.ServiceDetail})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the IotHubServiceDetail resource defined in the Terraform configuration
func IotHubServiceDetailSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"service_detail": {
			Description: ``,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the IotHubServiceDetail resource
func GetIotHubServiceDetailPropertyFields() (t []string) {
	return []string{
		"service_detail",
	}
}
