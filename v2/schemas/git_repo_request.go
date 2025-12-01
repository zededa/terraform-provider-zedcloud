package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func GitRepoRequestModel(d *schema.ResourceData) *models.GitRepoRequest {
	var data *models.GitRepoData // GitRepoData
	dataInterface, dataIsSet := d.GetOk("data")
	if dataIsSet && dataInterface != nil {
		dataMap := dataInterface.([]interface{})
		if len(dataMap) > 0 {
			data = GitRepoDataModelFromMap(dataMap[0].(map[string]interface{}))
		}
	}
	var metadata *models.GitRepoMetadata // GitRepoMetadata
	metadataInterface, metadataIsSet := d.GetOk("metadata")
	if metadataIsSet && metadataInterface != nil {
		metadataMap := metadataInterface.([]interface{})
		if len(metadataMap) > 0 {
			metadata = GitRepoMetadataModelFromMap(metadataMap[0].(map[string]interface{}))
		}
	}
	return &models.GitRepoRequest{
		Data:     data,
		Metadata: metadata,
	}
}

func GitRepoRequestModelFromMap(m map[string]interface{}) *models.GitRepoRequest {
	var data *models.GitRepoData // GitRepoData
	dataInterface, dataIsSet := m["data"]
	if dataIsSet && dataInterface != nil {
		dataMap := dataInterface.([]interface{})
		if len(dataMap) > 0 {
			data = GitRepoDataModelFromMap(dataMap[0].(map[string]interface{}))
		}
	}
	//
	var metadata *models.GitRepoMetadata // GitRepoMetadata
	metadataInterface, metadataIsSet := m["metadata"]
	if metadataIsSet && metadataInterface != nil {
		metadataMap := metadataInterface.([]interface{})
		if len(metadataMap) > 0 {
			metadata = GitRepoMetadataModelFromMap(metadataMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.GitRepoRequest{
		Data:     data,
		Metadata: metadata,
	}
}

func SetGitRepoRequestResourceData(d *schema.ResourceData, m *models.GitRepoRequest) {
	d.Set("data", SetGitRepoDataSubResourceData([]*models.GitRepoData{m.Data}))
	d.Set("metadata", SetGitRepoMetadataSubResourceData([]*models.GitRepoMetadata{m.Metadata}))
}

func SetGitRepoRequestSubResourceData(m []*models.GitRepoRequest) (d []*map[string]interface{}) {
	for _, GitRepoRequestModel := range m {
		if GitRepoRequestModel != nil {
			properties := make(map[string]interface{})
			properties["data"] = SetGitRepoDataSubResourceData([]*models.GitRepoData{GitRepoRequestModel.Data})
			properties["metadata"] = SetGitRepoMetadataSubResourceData([]*models.GitRepoMetadata{GitRepoRequestModel.Metadata})
			d = append(d, &properties)
		}
	}
	return
}

func GitRepoRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"data": {
			Description: `Data specification for the GitRepo (required)`,
			Type:        schema.TypeList, //GoType: GitRepoData
			Elem: &schema.Resource{
				Schema: GitRepoDataSchema(),
			},
			Required: true,
		},

		"metadata": {
			Description: `Metadata for the GitRepo (required)`,
			Type:        schema.TypeList, //GoType: GitRepoMetadata
			Elem: &schema.Resource{
				Schema: GitRepoMetadataSchema(),
			},
			Required: true,
		},
	}
}

func GetGitRepoRequestPropertyFields() (t []string) {
	return []string{
		"data",
		"metadata",
	}
}

func ConvertGitRepoResponseModelToRequestModel(m *models.GitRepoGetResponse) *models.GitRepoRequest {
	if m == nil {
		return nil
	}
	return &models.GitRepoRequest{
		Data: m.Data,
		Metadata: &models.GitRepoMetadata{
			Name:        &m.Name,
			Title:       m.Title,
			Description: m.Description,
		},
	}
}
