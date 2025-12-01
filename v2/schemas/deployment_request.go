package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func DeploymentRequestModel(d *schema.ResourceData) *models.DeploymentRequest {
	var data *models.DeploymentData // DeploymentData
	dataInterface, dataIsSet := d.GetOk("data")
	if dataIsSet && dataInterface != nil {
		dataMap := dataInterface.([]interface{})
		if len(dataMap) > 0 {
			data = DeploymentDataModelFromMap(dataMap[0].(map[string]interface{}))
		}
	}
	var metadata *models.Metadata // Metadata
	metadataInterface, metadataIsSet := d.GetOk("metadata")
	if metadataIsSet && metadataInterface != nil {
		metadataMap := metadataInterface.([]interface{})
		if len(metadataMap) > 0 {
			metadata = MetadataModelFromMap(metadataMap[0].(map[string]interface{}))
		}
	}
	var typeVar *models.K8sDeploymentType // K8sDeploymentType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewK8sDeploymentType(models.K8sDeploymentType(typeModel))
	}
	return &models.DeploymentRequest{
		Data:     data,
		Metadata: metadata,
		Type:     typeVar,
	}
}

func DeploymentRequestModelFromMap(m map[string]interface{}) *models.DeploymentRequest {
	var data *models.DeploymentData // DeploymentData
	dataInterface, dataIsSet := m["data"]
	if dataIsSet && dataInterface != nil {
		dataMap := dataInterface.([]interface{})
		if len(dataMap) > 0 {
			data = DeploymentDataModelFromMap(dataMap[0].(map[string]interface{}))
		}
	}
	//
	var metadata *models.Metadata // Metadata
	metadataInterface, metadataIsSet := m["metadata"]
	if metadataIsSet && metadataInterface != nil {
		metadataMap := metadataInterface.([]interface{})
		if len(metadataMap) > 0 {
			metadata = MetadataModelFromMap(metadataMap[0].(map[string]interface{}))
		}
	}
	//
	var typeVar *models.K8sDeploymentType // K8sDeploymentType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewK8sDeploymentType(models.K8sDeploymentType(typeModel))
	}
	return &models.DeploymentRequest{
		Data:     data,
		Metadata: metadata,
		Type:     typeVar,
	}
}

func SetDeploymentRequestResourceData(d *schema.ResourceData, m *models.DeploymentRequest) {
	d.Set("data", SetDeploymentDataSubResourceData([]*models.DeploymentData{m.Data}))
	d.Set("metadata", SetMetadataSubResourceData([]*models.Metadata{m.Metadata}))
	d.Set("type", m.Type)
}

func SetDeploymentRequestSubResourceData(m []*models.DeploymentRequest) (d []*map[string]interface{}) {
	for _, DeploymentRequestModel := range m {
		if DeploymentRequestModel != nil {
			properties := make(map[string]interface{})
			properties["data"] = SetDeploymentDataSubResourceData([]*models.DeploymentData{DeploymentRequestModel.Data})
			properties["metadata"] = SetMetadataSubResourceData([]*models.Metadata{DeploymentRequestModel.Metadata})
			properties["type"] = DeploymentRequestModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func DeploymentRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"data": {
			Description: `Data for the deployment (required)`,
			Type:        schema.TypeList, //GoType: DeploymentData
			Elem: &schema.Resource{
				Schema: DeploymentDataSchema(),
			},
			Required: true,
		},

		"metadata": {
			Description: `Metadata for the deployment (required)`,
			Type:        schema.TypeList, //GoType: Metadata
			Elem: &schema.Resource{
				Schema: MetadataSchema(),
			},
			Required: true,
		},

		"type": {
			Description: `Type of deployment (required). Example: 'HELMCHART'`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func GetDeploymentRequestPropertyFields() (t []string) {
	return []string{
		"data",
		"metadata",
		"type",
	}
}

func ConvertDeploymentGetResponseToDeploymentRequestModel(m *models.DeploymentGetResponse) *models.DeploymentRequest {
	if m == nil {
		return nil
	}
	return &models.DeploymentRequest{
		Metadata: &models.Metadata{
			Name:        &m.Name,
			Description: m.Description,
			Title:       m.Title,
			ProjectID:   &m.ProjectID,
		},
		Data: m.Spec,
		Type: m.Type,
	}
}
