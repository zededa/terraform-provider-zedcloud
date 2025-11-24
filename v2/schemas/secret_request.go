package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func SecretRequestModel(d *schema.ResourceData) *models.SecretRequest {
	var data *models.SecretData // SecretData
	dataInterface, dataIsSet := d.GetOk("data")
	if dataIsSet && dataInterface != nil {
		dataMap := dataInterface.([]interface{})
		if len(dataMap) > 0 {
			data = SecretDataModelFromMap(dataMap[0].(map[string]interface{}))
		}
	}
	var metadata *models.SecretMetadata // SecretMetadata
	metadataInterface, metadataIsSet := d.GetOk("metadata")
	if metadataIsSet && metadataInterface != nil {
		metadataMap := metadataInterface.([]interface{})
		if len(metadataMap) > 0 {
			metadata = SecretMetadataModelFromMap(metadataMap[0].(map[string]interface{}))
		}
	}
	return &models.SecretRequest{
		Data:     data,
		Metadata: metadata,
	}
}

func SecretRequestModelFromMap(m map[string]interface{}) *models.SecretRequest {
	var data *models.SecretData // SecretData
	dataInterface, dataIsSet := m["data"]
	if dataIsSet && dataInterface != nil {
		dataMap := dataInterface.([]interface{})
		if len(dataMap) > 0 {
			data = SecretDataModelFromMap(dataMap[0].(map[string]interface{}))
		}
	}
	//
	var metadata *models.SecretMetadata // SecretMetadata
	metadataInterface, metadataIsSet := m["metadata"]
	if metadataIsSet && metadataInterface != nil {
		metadataMap := metadataInterface.([]interface{})
		if len(metadataMap) > 0 {
			metadata = SecretMetadataModelFromMap(metadataMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.SecretRequest{
		Data:     data,
		Metadata: metadata,
	}
}

func SetSecretRequestResourceData(d *schema.ResourceData, m *models.SecretRequest) {
	d.Set("data", SetSecretDataSubResourceData([]*models.SecretData{m.Data}))
	d.Set("metadata", SetSecretMetadataSubResourceData([]*models.SecretMetadata{m.Metadata}))
}

func SetSecretRequestSubResourceData(m []*models.SecretRequest) (d []*map[string]interface{}) {
	for _, SecretRequestModel := range m {
		if SecretRequestModel != nil {
			properties := make(map[string]interface{})
			properties["data"] = SetSecretDataSubResourceData([]*models.SecretData{SecretRequestModel.Data})
			properties["metadata"] = SetSecretMetadataSubResourceData([]*models.SecretMetadata{SecretRequestModel.Metadata})
			d = append(d, &properties)
		}
	}
	return
}

func SecretRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"data": {
			Description: `Secret data including credentials or keys. This field is required.`,
			Type:        schema.TypeList, //GoType: SecretData
			Elem: &schema.Resource{
				Schema: SecretDataSchema(),
			},
			Required: true,
		},

		"metadata": {
			Description: `Metadata for the secret to be created. This field is required.`,
			Type:        schema.TypeList, //GoType: SecretMetadata
			Elem: &schema.Resource{
				Schema: SecretMetadataSchema(),
			},
			Required: true,
		},
	}
}

func GetSecretRequestPropertyFields() (t []string) {
	return []string{
		"data",
		"metadata",
	}
}

func ConvertGetSecretResponseModelToSecretRequestModel(m *models.SecretInfoResponse) *models.SecretRequest {
	if m == nil {
		return nil
	}
	return &models.SecretRequest{
		Metadata: m.Metadata,
	}
}
