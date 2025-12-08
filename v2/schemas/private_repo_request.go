package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func PrivateRepoRequestModel(d *schema.ResourceData) *models.PrivateRepoRequest {
	var metadata *models.PrivateRepoMetadata // PrivateRepoMetadata
	metadataInterface, metadataIsSet := d.GetOk("metadata")
	if metadataIsSet && metadataInterface != nil {
		metadataMap := metadataInterface.([]interface{})
		if len(metadataMap) > 0 {
			metadata = PrivateRepoMetadataModelFromMap(metadataMap[0].(map[string]interface{}))
		}
	}
	var spec *models.PrivateRepoSpec // PrivateRepoSpec
	specInterface, specIsSet := d.GetOk("spec")
	if specIsSet && specInterface != nil {
		specMap := specInterface.([]interface{})
		if len(specMap) > 0 {
			spec = PrivateRepoSpecModelFromMap(specMap[0].(map[string]interface{}))
		}
	}
	return &models.PrivateRepoRequest{
		Metadata: metadata,
		Spec:     spec,
	}
}

func PrivateRepoRequestModelFromMap(m map[string]interface{}) *models.PrivateRepoRequest {
	var metadata *models.PrivateRepoMetadata // PrivateRepoMetadata
	metadataInterface, metadataIsSet := m["metadata"]
	if metadataIsSet && metadataInterface != nil {
		metadataMap := metadataInterface.([]interface{})
		if len(metadataMap) > 0 {
			metadata = PrivateRepoMetadataModelFromMap(metadataMap[0].(map[string]interface{}))
		}
	}
	//
	var spec *models.PrivateRepoSpec // PrivateRepoSpec
	specInterface, specIsSet := m["spec"]
	if specIsSet && specInterface != nil {
		specMap := specInterface.([]interface{})
		if len(specMap) > 0 {
			spec = PrivateRepoSpecModelFromMap(specMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.PrivateRepoRequest{
		Metadata: metadata,
		Spec:     spec,
	}
}

func SetPrivateRepoRequestResourceData(d *schema.ResourceData, m *models.PrivateRepoRequest) {
	d.Set("metadata", SetPrivateRepoMetadataSubResourceData([]*models.PrivateRepoMetadata{m.Metadata}))
	d.Set("spec", SetPrivateRepoSpecSubResourceData([]*models.PrivateRepoSpec{m.Spec}))
}

func SetPrivateRepoRequestSubResourceData(m []*models.PrivateRepoRequest) (d []*map[string]interface{}) {
	for _, PrivateRepoRequestModel := range m {
		if PrivateRepoRequestModel != nil {
			properties := make(map[string]interface{})
			properties["metadata"] = SetPrivateRepoMetadataSubResourceData([]*models.PrivateRepoMetadata{PrivateRepoRequestModel.Metadata})
			properties["spec"] = SetPrivateRepoSpecSubResourceData([]*models.PrivateRepoSpec{PrivateRepoRequestModel.Spec})
			d = append(d, &properties)
		}
	}
	return
}

func PrivateRepoRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"metadata": {
			Description: `Metadata for the private repository (required)`,
			Type:        schema.TypeList, //GoType: PrivateRepoMetadata
			Elem: &schema.Resource{
				Schema: PrivateRepoMetadataSchema(),
			},
			Required: true,
		},

		"spec": {
			Description: `Specification for the private repository (required)`,
			Type:        schema.TypeList, //GoType: PrivateRepoSpec
			Elem: &schema.Resource{
				Schema: PrivateRepoSpecSchema(),
			},
			Required: true,
		},
	}
}

func GetPrivateRepoRequestPropertyFields() (t []string) {
	return []string{
		"metadata",
		"spec",
	}
}

func ConvertPrivateRepoResponseModelToPrivateRepoRequestModel(m *models.PrivateRepoGetByIDResponse,
) *models.PrivateRepoRequest {
	var spec *models.PrivateRepoSpec
	if m.Spec != nil {
		spec = &models.PrivateRepoSpec{
			RepoDetails: m.Spec.RepoDetails,
		}
		if m.Spec.Auth != nil {
			spec.Auth = &models.PrivateRepoAuthConfig{}
			if m.Spec.Auth.Secret != nil {
				spec.Auth.Secret = &models.PrivateRepoSecretConfig{
					Name: m.Spec.Auth.Secret.Name,
					Type: m.Spec.Auth.Secret.Type,
				}
			}
		}
	}

	return &models.PrivateRepoRequest{
		Metadata: m.Metadata,
		Spec:     spec,
	}
}
