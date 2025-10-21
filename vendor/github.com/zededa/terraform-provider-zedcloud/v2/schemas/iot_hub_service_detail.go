package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func IotHubServiceDetailModel(d *schema.ResourceData) *models.IotHubServiceDetail {
	var serviceDetail *models.AzureResourceAndServiceDetail // AzureResourceAndServiceDetail
	serviceDetailInterface, serviceDetailIsSet := d.GetOk("service_detail")
	if serviceDetailIsSet && serviceDetailInterface != nil {
		serviceDetailMap := serviceDetailInterface.([]interface{})
		if len(serviceDetailMap) > 0 {
			serviceDetail = AzureResourceAndServiceDetailModelFromMap(serviceDetailMap[0].(map[string]interface{}))
		}
	}
	return &models.IotHubServiceDetail{
		ServiceDetail: serviceDetail,
	}
}

func IotHubServiceDetailModelFromMap(m map[string]interface{}) *models.IotHubServiceDetail {
	var serviceDetail *models.AzureResourceAndServiceDetail // AzureResourceAndServiceDetail
	serviceDetailInterface, serviceDetailIsSet := m["service_detail"]
	if serviceDetailIsSet && serviceDetailInterface != nil {
		serviceDetailMap := serviceDetailInterface.([]interface{})
		if len(serviceDetailMap) > 0 {
			serviceDetail = AzureResourceAndServiceDetailModelFromMap(serviceDetailMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.IotHubServiceDetail{
		ServiceDetail: serviceDetail,
	}
}

func SetIotHubServiceDetailResourceData(d *schema.ResourceData, m *models.IotHubServiceDetail) {
	d.Set("service_detail", SetAzureResourceAndServiceDetailSubResourceData([]*models.AzureResourceAndServiceDetail{m.ServiceDetail}))
}

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

func IotHubServiceDetail() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"service_detail": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AzureResourceAndServiceDetail
			Elem: &schema.Resource{
				Schema: AzureResourceAndServiceDetail(),
			},
			Optional: true,
		},
	}
}

func GetIotHubServiceDetailPropertyFields() (t []string) {
	return []string{
		"service_detail",
	}
}
