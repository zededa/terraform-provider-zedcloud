package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func SecretInfoResponseModel(d *schema.ResourceData) *models.SecretInfoResponse {
	typeVar, _ := d.Get("type").(string)
	var data *models.SSHCredentials // SSHCredentials
	dataInterface, dataIsSet := d.GetOk("data")
	if dataIsSet && dataInterface != nil {
		dataMap := dataInterface.([]interface{})
		if len(dataMap) > 0 {
			data = SSHCredentialsModelFromMap(dataMap[0].(map[string]interface{}))
		}
	}
	id, _ := d.Get("id").(string)
	kind, _ := d.Get("kind").(string)
	var metadata *models.SecretMetadata // SecretMetadata
	metadataInterface, metadataIsSet := d.GetOk("metadata")
	if metadataIsSet && metadataInterface != nil {
		metadataMap := metadataInterface.([]interface{})
		if len(metadataMap) > 0 {
			metadata = SecretMetadataModelFromMap(metadataMap[0].(map[string]interface{}))
		}
	}
	return &models.SecretInfoResponse{
		Type:     typeVar,
		Data:     data,
		ID:       id,
		Kind:     kind,
		Metadata: metadata,
	}
}

func SecretInfoResponseModelFromMap(m map[string]interface{}) *models.SecretInfoResponse {
	typeVar := m["type"].(string)
	var data *models.SSHCredentials // SSHCredentials
	dataInterface, dataIsSet := m["data"]
	if dataIsSet && dataInterface != nil {
		dataMap := dataInterface.([]interface{})
		if len(dataMap) > 0 {
			data = SSHCredentialsModelFromMap(dataMap[0].(map[string]interface{}))
		}
	}
	//
	id := m["id"].(string)
	kind := m["kind"].(string)
	var metadata *models.SecretMetadata // SecretMetadata
	metadataInterface, metadataIsSet := m["metadata"]
	if metadataIsSet && metadataInterface != nil {
		metadataMap := metadataInterface.([]interface{})
		if len(metadataMap) > 0 {
			metadata = SecretMetadataModelFromMap(metadataMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.SecretInfoResponse{
		Type:     typeVar,
		Data:     data,
		ID:       id,
		Kind:     kind,
		Metadata: metadata,
	}
}

func SetSecretInfoResponseResourceData(d *schema.ResourceData, m *models.SecretInfoResponse) {
	d.Set("type", m.Type)
	d.Set("data", SetSSHCredentialsSubResourceData([]*models.SSHCredentials{m.Data}))
	d.Set("id", m.ID)
	d.Set("kind", m.Kind)
	d.Set("metadata", SetSecretMetadataSubResourceData([]*models.SecretMetadata{m.Metadata}))
}

func SetSecretInfoResponseSubResourceData(m []*models.SecretInfoResponse) (d []*map[string]interface{}) {
	for _, SecretInfoResponseModel := range m {
		if SecretInfoResponseModel != nil {
			properties := make(map[string]interface{})
			properties["type"] = SecretInfoResponseModel.Type
			properties["data"] = SetSSHCredentialsSubResourceData([]*models.SSHCredentials{SecretInfoResponseModel.Data})
			properties["id"] = SecretInfoResponseModel.ID
			properties["kind"] = SecretInfoResponseModel.Kind
			properties["metadata"] = SetSecretMetadataSubResourceData([]*models.SecretMetadata{SecretInfoResponseModel.Metadata})
			d = append(d, &properties)
		}
	}
	return
}

func SecretInfoResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"type": {
			Description: `Type of the secret`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"data": {
			Description: `SSH key credentials`,
			Type:        schema.TypeList, //GoType: SSHCredentials
			Elem: &schema.Resource{
				Schema: SSHCredentialsSchema(),
			},
			Optional: true,
		},

		"id": {
			Description: `Unique identifier of the secret`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"kind": {
			Description: `Kubernetes resource kind`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"metadata": {
			Description: `Metadata information for the secret`,
			Type:        schema.TypeList, //GoType: SecretMetadata
			Elem: &schema.Resource{
				Schema: SecretMetadataSchema(),
			},
			Optional: true,
		},
	}
}

func GetSecretInfoResponsePropertyFields() (t []string) {
	return []string{
		"type",
		"data",
		"id",
		"kind",
		"metadata",
	}
}
