package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/go-openapi/strfmt"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func DeploymentGetResponseModel(d *schema.ResourceData) *models.DeploymentGetResponse {
	createdAt, _ := d.Get("created_at").(strfmt.DateTime)
	createdBy, _ := d.Get("created_by").(string)
	deploymentStatus, _ := d.Get("deployment_status").(string)
	description, _ := d.Get("description").(string)
	id, _ := d.Get("id").(string)
	name, _ := d.Get("name").(string)
	projectID, _ := d.Get("project_id").(string)
	var spec *models.DeploymentData // DeploymentData
	specInterface, specIsSet := d.GetOk("spec")
	if specIsSet && specInterface != nil {
		specMap := specInterface.([]interface{})
		if len(specMap) > 0 {
			spec = DeploymentDataModelFromMap(specMap[0].(map[string]interface{}))
		}
	}
	title, _ := d.Get("title").(string)
	var typeVar *models.K8sDeploymentType // K8sDeploymentType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewK8sDeploymentType(models.K8sDeploymentType(typeModel))
	}
	updatedAt, _ := d.Get("updated_at").(strfmt.DateTime)
	updatedBy, _ := d.Get("updated_by").(string)
	return &models.DeploymentGetResponse{
		CreatedAt:        createdAt,
		CreatedBy:        createdBy,
		DeploymentStatus: deploymentStatus,
		Description:      description,
		ID:               id,
		Name:             name,
		ProjectID:        projectID,
		Spec:             spec,
		Title:            title,
		Type:             typeVar,
		UpdatedAt:        updatedAt,
		UpdatedBy:        updatedBy,
	}
}

func DeploymentGetResponseModelFromMap(m map[string]interface{}) *models.DeploymentGetResponse {
	createdAt := m["created_at"].(strfmt.DateTime)
	createdBy := m["created_by"].(string)
	deploymentStatus := m["deployment_status"].(string)
	description := m["description"].(string)
	id := m["id"].(string)
	name := m["name"].(string)
	projectID := m["project_id"].(string)
	var spec *models.DeploymentData // DeploymentData
	specInterface, specIsSet := m["spec"]
	if specIsSet && specInterface != nil {
		specMap := specInterface.([]interface{})
		if len(specMap) > 0 {
			spec = DeploymentDataModelFromMap(specMap[0].(map[string]interface{}))
		}
	}
	//
	title := m["title"].(string)
	var typeVar *models.K8sDeploymentType // K8sDeploymentType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewK8sDeploymentType(models.K8sDeploymentType(typeModel))
	}
	updatedAt := m["updated_at"].(strfmt.DateTime)
	updatedBy := m["updated_by"].(string)
	return &models.DeploymentGetResponse{
		CreatedAt:        createdAt,
		CreatedBy:        createdBy,
		DeploymentStatus: deploymentStatus,
		Description:      description,
		ID:               id,
		Name:             name,
		ProjectID:        projectID,
		Spec:             spec,
		Title:            title,
		Type:             typeVar,
		UpdatedAt:        updatedAt,
		UpdatedBy:        updatedBy,
	}
}

func SetDeploymentGetResponseResourceData(d *schema.ResourceData, m *models.DeploymentGetResponse) {
	d.Set("created_at", m.CreatedAt)
	d.Set("created_by", m.CreatedBy)
	d.Set("deployment_status", m.DeploymentStatus)
	d.Set("description", m.Description)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("project_id", m.ProjectID)
	d.Set("resource_status_map", m.ResourceStatusMap)
	d.Set("spec", SetDeploymentDataSubResourceData([]*models.DeploymentData{m.Spec}))
	d.Set("title", m.Title)
	d.Set("type", m.Type)
	d.Set("updated_at", m.UpdatedAt)
	d.Set("updated_by", m.UpdatedBy)
}

func SetDeploymentGetResponseSubResourceData(m []*models.DeploymentGetResponse) (d []*map[string]interface{}) {
	for _, DeploymentGetResponseModel := range m {
		if DeploymentGetResponseModel != nil {
			properties := make(map[string]interface{})
			properties["created_at"] = DeploymentGetResponseModel.CreatedAt.String()
			properties["created_by"] = DeploymentGetResponseModel.CreatedBy
			properties["deployment_status"] = DeploymentGetResponseModel.DeploymentStatus
			properties["description"] = DeploymentGetResponseModel.Description
			properties["id"] = DeploymentGetResponseModel.ID
			properties["name"] = DeploymentGetResponseModel.Name
			properties["project_id"] = DeploymentGetResponseModel.ProjectID
			properties["resource_status_map"] = DeploymentGetResponseModel.ResourceStatusMap
			properties["spec"] = SetDeploymentDataSubResourceData([]*models.DeploymentData{DeploymentGetResponseModel.Spec})
			properties["title"] = DeploymentGetResponseModel.Title
			properties["type"] = DeploymentGetResponseModel.Type
			properties["updated_at"] = DeploymentGetResponseModel.UpdatedAt.String()
			properties["updated_by"] = DeploymentGetResponseModel.UpdatedBy
			d = append(d, &properties)
		}
	}
	return
}

func DeploymentGetResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"created_at": {
			Description:  `Creation timestamp`,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"created_by": {
			Description: `User who created the deployment`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"deployment_status": {
			Description: `Deployment status`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"description": {
			Description: `Description of the deployment`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: `Deployment ID`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: `Release name`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"project_id": {
			Description: `Project ID associated with the deployment`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"spec": {
			Description: `Deployment spec`,
			Type:        schema.TypeList, //GoType: DeploymentData
			Elem: &schema.Resource{
				Schema: DeploymentDataSchema(),
			},
			Optional: true,
		},

		"title": {
			Description: `Title of the deployment`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"type": {
			Description: `Type of deployment`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"updated_at": {
			Description:  `Last update timestamp`,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"updated_by": {
			Description: `User who last updated the deployment`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetDeploymentGetResponsePropertyFields() (t []string) {
	return []string{
		"created_at",
		"created_by",
		"deployment_status",
		"description",
		"id",
		"name",
		"project_id",
		"resource_status_map",
		"spec",
		"title",
		"type",
		"updated_at",
		"updated_by",
	}
}
