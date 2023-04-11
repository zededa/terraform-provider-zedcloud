package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AzureResourceAndServicesModel(d *schema.ResourceData) *models.AzureResourceAndServices {
	var dpsService *models.DPSServiceDetail // DPSServiceDetail
	dpsServiceInterface, dpsServiceIsSet := d.GetOk("dps_service")
	if dpsServiceIsSet && dpsServiceInterface != nil {
		dpsServiceMap := dpsServiceInterface.([]interface{})
		if len(dpsServiceMap) > 0 {
			dpsService = DPSServiceDetailModelFromMap(dpsServiceMap[0].(map[string]interface{}))
		}
	}
	var iotHubService []*models.IotHubServiceDetail // []*IotHubServiceDetail
	iotHubServiceInterface, iotHubServiceIsSet := d.GetOk("iot_hub_service")
	if iotHubServiceIsSet {
		var items []interface{}
		if listItems, isList := iotHubServiceInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = iotHubServiceInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := IotHubServiceDetailModelFromMap(v.(map[string]interface{}))
			iotHubService = append(iotHubService, m)
		}
	}
	var resourceGroup []*models.ResourceGroupDetail // []*ResourceGroupDetail
	resourceGroupInterface, resourceGroupIsSet := d.GetOk("resource_group")
	if resourceGroupIsSet {
		var items []interface{}
		if listItems, isList := resourceGroupInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = resourceGroupInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ResourceGroupDetailModelFromMap(v.(map[string]interface{}))
			resourceGroup = append(resourceGroup, m)
		}
	}
	return &models.AzureResourceAndServices{
		DpsService:    dpsService,
		IotHubService: iotHubService,
		ResourceGroup: resourceGroup,
	}
}

func AzureResourceAndServicesModelFromMap(m map[string]interface{}) *models.AzureResourceAndServices {
	var dpsService *models.DPSServiceDetail // DPSServiceDetail
	dpsServiceInterface, dpsServiceIsSet := m["dps_service"]
	if dpsServiceIsSet && dpsServiceInterface != nil {
		dpsServiceMap := dpsServiceInterface.([]interface{})
		if len(dpsServiceMap) > 0 {
			dpsService = DPSServiceDetailModelFromMap(dpsServiceMap[0].(map[string]interface{}))
		}
	}
	//
	var iotHubService []*models.IotHubServiceDetail // []*IotHubServiceDetail
	iotHubServiceInterface, iotHubServiceIsSet := m["iot_hub_service"]
	if iotHubServiceIsSet {
		var items []interface{}
		if listItems, isList := iotHubServiceInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = iotHubServiceInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := IotHubServiceDetailModelFromMap(v.(map[string]interface{}))
			iotHubService = append(iotHubService, m)
		}
	}
	var resourceGroup []*models.ResourceGroupDetail // []*ResourceGroupDetail
	resourceGroupInterface, resourceGroupIsSet := m["resource_group"]
	if resourceGroupIsSet {
		var items []interface{}
		if listItems, isList := resourceGroupInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = resourceGroupInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ResourceGroupDetailModelFromMap(v.(map[string]interface{}))
			resourceGroup = append(resourceGroup, m)
		}
	}
	return &models.AzureResourceAndServices{
		DpsService:    dpsService,
		IotHubService: iotHubService,
		ResourceGroup: resourceGroup,
	}
}

func SetAzureResourceAndServicesResourceData(d *schema.ResourceData, m *models.AzureResourceAndServices) {
	d.Set("dps_service", SetDPSServiceDetailSubResourceData([]*models.DPSServiceDetail{m.DpsService}))
	d.Set("iot_hub_service", SetIotHubServiceDetailSubResourceData(m.IotHubService))
	d.Set("resource_group", SetResourceGroupDetailSubResourceData(m.ResourceGroup))
}

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

func AzureResourceAndServices() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"dps_service": {
			Description: `dps service attached to cloud policy`,
			Type:        schema.TypeList, //GoType: DPSServiceDetail
			Elem: &schema.Resource{
				Schema: DPSServiceDetail(),
			},
			Required: true,
		},

		"iot_hub_service": {
			Description: `list of iothubs attached to cloud policy`,
			Type:        schema.TypeList, //GoType: []*IotHubServiceDetail
			Elem: &schema.Resource{
				Schema: IotHubServiceDetail(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Required: true,
		},

		"resource_group": {
			Description: `list of resource groups attached to cloud policy`,
			Type:        schema.TypeList, //GoType: []*ResourceGroupDetail
			Elem: &schema.Resource{
				Schema: ResourceGroupDetail(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Required: true,
		},
	}
}

func GetAzureResourceAndServicesPropertyFields() (t []string) {
	return []string{
		"dps_service",
		"iot_hub_service",
		"resource_group",
	}
}
