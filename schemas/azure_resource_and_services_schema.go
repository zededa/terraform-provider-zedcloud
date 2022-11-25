package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate AzureResourceAndServices resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AzureResourceAndServicesModel(d *schema.ResourceData) *models.AzureResourceAndServices {
	var dpsService *models.DPSServiceDetail // DPSServiceDetail
	dpsServiceInterface, dpsServiceIsSet := d.GetOk("dps_service")
	if dpsServiceIsSet {
		dpsServiceMap := dpsServiceInterface.([]interface{})[0].(map[string]interface{})
		dpsService = DPSServiceDetailModelFromMap(dpsServiceMap)
	}
	iotHubService := d.Get("iot_hub_service").([]*models.IotHubServiceDetail) // []*IotHubServiceDetail
	resourceGroup := d.Get("resource_group").([]*models.ResourceGroupDetail)  // []*ResourceGroupDetail
	return &models.AzureResourceAndServices{
		DpsService:    dpsService,
		IotHubService: iotHubService,
		ResourceGroup: resourceGroup,
	}
}

func AzureResourceAndServicesModelFromMap(m map[string]interface{}) *models.AzureResourceAndServices {
	var dpsService *models.DPSServiceDetail // DPSServiceDetail
	dpsServiceInterface, dpsServiceIsSet := m["dps_service"]
	if dpsServiceIsSet {
		dpsServiceMap := dpsServiceInterface.([]interface{})[0].(map[string]interface{})
		dpsService = DPSServiceDetailModelFromMap(dpsServiceMap)
	}
	//
	iotHubService := m["iot_hub_service"].([]*models.IotHubServiceDetail) // []*IotHubServiceDetail
	resourceGroup := m["resource_group"].([]*models.ResourceGroupDetail)  // []*ResourceGroupDetail
	return &models.AzureResourceAndServices{
		DpsService:    dpsService,
		IotHubService: iotHubService,
		ResourceGroup: resourceGroup,
	}
}

// Update the underlying AzureResourceAndServices resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAzureResourceAndServicesResourceData(d *schema.ResourceData, m *models.AzureResourceAndServices) {
	d.Set("dps_service", SetDPSServiceDetailSubResourceData([]*models.DPSServiceDetail{m.DpsService}))
	d.Set("iot_hub_service", SetIotHubServiceDetailSubResourceData(m.IotHubService))
	d.Set("resource_group", SetResourceGroupDetailSubResourceData(m.ResourceGroup))
}

// Iterate throught and update the AzureResourceAndServices resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAzureResourceAndServicesSubResourceData(m []*models.AzureResourceAndServices) (d []*map[string]interface{}) {
	for _, AzureResourceAndServicesModel := range m {
		if AzureResourceAndServicesModel != nil {
			properties := make(map[string]interface{})
			properties["dps_service"] = SetDPSServiceDetailSubResourceData([]*models.DPSServiceDetail{AzureResourceAndServicesModel.DpsService})
			properties["iot_hub_service"] = SetIotHubServiceDetailSubResourceData(AzureResourceAndServicesModel.IotHubService)
			properties["resource_group"] = SetResourceGroupDetailSubResourceData(AzureResourceAndServicesModel.ResourceGroup)
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AzureResourceAndServices resource defined in the Terraform configuration
func AzureResourceAndServicesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"dps_service": {
			Description: `dps service attached to cloud policy`,
			Type:        schema.TypeList, //GoType: DPSServiceDetail
			Elem: &schema.Resource{
				Schema: DPSServiceDetailSchema(),
			},
			Required: true,
		},

		"iot_hub_service": {
			Description: `list of iothubs attached to cloud policy`,
			Type:        schema.TypeList, //GoType: []*IotHubServiceDetail
			Elem: &schema.Resource{
				Schema: IotHubServiceDetailSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Required:   true,
		},

		"resource_group": {
			Description: `list of resource groups attached to cloud policy`,
			Type:        schema.TypeList, //GoType: []*ResourceGroupDetail
			Elem: &schema.Resource{
				Schema: ResourceGroupDetailSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Required:   true,
		},
	}
}

// Retrieve property field names for updating the AzureResourceAndServices resource
func GetAzureResourceAndServicesPropertyFields() (t []string) {
	return []string{
		"dps_service",
		"iot_hub_service",
		"resource_group",
	}
}
