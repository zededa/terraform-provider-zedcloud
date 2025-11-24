package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func PrivateRepoGetByIDResponseModel(d *schema.ResourceData) *models.PrivateRepoGetByIDResponse {
	id, _ := d.Get("id").(string)
	var metadata *models.PrivateRepoMetadata // PrivateRepoMetadata
	metadataInterface, metadataIsSet := d.GetOk("metadata")
	if metadataIsSet && metadataInterface != nil {
		metadataMap := metadataInterface.([]interface{})
		if len(metadataMap) > 0 {
			metadata = PrivateRepoMetadataModelFromMap(metadataMap[0].(map[string]interface{}))
		}
	}
	repoIdentifier, _ := d.Get("repo_identifier").(string)
	var spec *models.PrivateRepoSpecInfo // PrivateRepoSpecInfo
	specInterface, specIsSet := d.GetOk("spec")
	if specIsSet && specInterface != nil {
		specMap := specInterface.([]interface{})
		if len(specMap) > 0 {
			spec = PrivateRepoSpecInfoModelFromMap(specMap[0].(map[string]interface{}))
		}
	}
	var status *models.PrivateRepoStatus // PrivateRepoStatus
	statusInterface, statusIsSet := d.GetOk("status")
	if statusIsSet && statusInterface != nil {
		statusMap := statusInterface.([]interface{})
		if len(statusMap) > 0 {
			status = PrivateRepoStatusModelFromMap(statusMap[0].(map[string]interface{}))
		}
	}
	return &models.PrivateRepoGetByIDResponse{
		ID:             id,
		Metadata:       metadata,
		RepoIdentifier: repoIdentifier,
		Spec:           spec,
		Status:         status,
	}
}

func PrivateRepoGetByIDResponseModelFromMap(m map[string]interface{}) *models.PrivateRepoGetByIDResponse {
	id := m["id"].(string)
	var metadata *models.PrivateRepoMetadata // PrivateRepoMetadata
	metadataInterface, metadataIsSet := m["metadata"]
	if metadataIsSet && metadataInterface != nil {
		metadataMap := metadataInterface.([]interface{})
		if len(metadataMap) > 0 {
			metadata = PrivateRepoMetadataModelFromMap(metadataMap[0].(map[string]interface{}))
		}
	}
	//
	repoIdentifier := m["repo_identifier"].(string)
	var spec *models.PrivateRepoSpecInfo // PrivateRepoSpecInfo
	specInterface, specIsSet := m["spec"]
	if specIsSet && specInterface != nil {
		specMap := specInterface.([]interface{})
		if len(specMap) > 0 {
			spec = PrivateRepoSpecInfoModelFromMap(specMap[0].(map[string]interface{}))
		}
	}
	//
	var status *models.PrivateRepoStatus // PrivateRepoStatus
	statusInterface, statusIsSet := m["status"]
	if statusIsSet && statusInterface != nil {
		statusMap := statusInterface.([]interface{})
		if len(statusMap) > 0 {
			status = PrivateRepoStatusModelFromMap(statusMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.PrivateRepoGetByIDResponse{
		ID:             id,
		Metadata:       metadata,
		RepoIdentifier: repoIdentifier,
		Spec:           spec,
		Status:         status,
	}
}

func SetPrivateRepoGetByIDResponseResourceData(d *schema.ResourceData, m *models.PrivateRepoGetByIDResponse) {
	d.Set("id", m.ID)
	d.Set("metadata", SetPrivateRepoMetadataSubResourceData([]*models.PrivateRepoMetadata{m.Metadata}))
	d.Set("repo_identifier", m.RepoIdentifier)
	d.Set("spec", SetPrivateRepoSpecInfoSubResourceData([]*models.PrivateRepoSpecInfo{m.Spec}))
	d.Set("status", SetPrivateRepoStatusSubResourceData([]*models.PrivateRepoStatus{m.Status}))
}

func SetPrivateRepoGetByIDResponseSubResourceData(m []*models.PrivateRepoGetByIDResponse) (d []*map[string]interface{}) {
	for _, PrivateRepoGetByIDResponseModel := range m {
		if PrivateRepoGetByIDResponseModel != nil {
			properties := make(map[string]interface{})
			properties["id"] = PrivateRepoGetByIDResponseModel.ID
			properties["metadata"] = SetPrivateRepoMetadataSubResourceData([]*models.PrivateRepoMetadata{PrivateRepoGetByIDResponseModel.Metadata})
			properties["repo_identifier"] = PrivateRepoGetByIDResponseModel.RepoIdentifier
			properties["spec"] = SetPrivateRepoSpecInfoSubResourceData([]*models.PrivateRepoSpecInfo{PrivateRepoGetByIDResponseModel.Spec})
			properties["status"] = SetPrivateRepoStatusSubResourceData([]*models.PrivateRepoStatus{PrivateRepoGetByIDResponseModel.Status})
			d = append(d, &properties)
		}
	}
	return
}

func PrivateRepoGetByIDResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Description: `Unique identifier of the private repository`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"metadata": {
			Description: `Metadata information for the private repository`,
			Type:        schema.TypeList, //GoType: PrivateRepoMetadata
			Elem: &schema.Resource{
				Schema: PrivateRepoMetadataSchema(),
			},
			Optional: true,
		},

		"repo_identifier": {
			Description: `Unique identifier for the private repository`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"spec": {
			Description: `Enhanced specification of the private repository`,
			Type:        schema.TypeList, //GoType: PrivateRepoSpecInfo
			Elem: &schema.Resource{
				Schema: PrivateRepoSpecInfoSchema(),
			},
			Optional: true,
		},

		"status": {
			Description: `Current operational status of the private repository`,
			Type:        schema.TypeList, //GoType: PrivateRepoStatus
			Elem: &schema.Resource{
				Schema: PrivateRepoStatusSchema(),
			},
			Optional: true,
		},
	}
}

func GetPrivateRepoGetByIDResponsePropertyFields() (t []string) {
	return []string{
		"id",
		"metadata",
		"repo_identifier",
		"spec",
		"status",
	}
}
