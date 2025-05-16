package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZserviceDeploymentUpdateResponseModel(d *schema.ResourceData) *models.ZserviceDeploymentUpdateResponse {
	deploymentID, _ := d.Get("deployment_id").(string)
	error, _ := d.Get("error").(string)
	var workflowInfo []*models.ZserviceWorkflowInfo // []*ZserviceWorkflowInfo
	workflowInfoInterface, workflowInfoIsSet := d.GetOk("workflow_info")
	if workflowInfoIsSet {
		var items []interface{}
		if listItems, isList := workflowInfoInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = workflowInfoInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ZserviceWorkflowInfoModelFromMap(v.(map[string]interface{}))
			workflowInfo = append(workflowInfo, m)
		}
	}
	return &models.ZserviceDeploymentUpdateResponse{
		DeploymentID: deploymentID,
		Error:        error,
		WorkflowInfo: workflowInfo,
	}
}

func ZserviceDeploymentUpdateResponseModelFromMap(m map[string]interface{}) *models.ZserviceDeploymentUpdateResponse {
	deploymentID := m["deployment_id"].(string)
	error := m["error"].(string)
	var workflowInfo []*models.ZserviceWorkflowInfo // []*ZserviceWorkflowInfo
	workflowInfoInterface, workflowInfoIsSet := m["workflow_info"]
	if workflowInfoIsSet {
		var items []interface{}
		if listItems, isList := workflowInfoInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = workflowInfoInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ZserviceWorkflowInfoModelFromMap(v.(map[string]interface{}))
			workflowInfo = append(workflowInfo, m)
		}
	}
	return &models.ZserviceDeploymentUpdateResponse{
		DeploymentID: deploymentID,
		Error:        error,
		WorkflowInfo: workflowInfo,
	}
}

func SetZserviceDeploymentUpdateResponseResourceData(d *schema.ResourceData, m *models.ZserviceDeploymentUpdateResponse) {
	d.Set("deployment_id", m.DeploymentID)
	d.Set("error", m.Error)
	d.Set("workflow_info", SetZserviceWorkflowInfoSubResourceData(m.WorkflowInfo))
}

func SetZserviceDeploymentUpdateResponseSubResourceData(m []*models.ZserviceDeploymentUpdateResponse) (d []*map[string]interface{}) {
	for _, ZserviceDeploymentUpdateResponseModel := range m {
		if ZserviceDeploymentUpdateResponseModel != nil {
			properties := make(map[string]interface{})
			properties["deployment_id"] = ZserviceDeploymentUpdateResponseModel.DeploymentID
			properties["error"] = ZserviceDeploymentUpdateResponseModel.Error
			properties["workflow_info"] = SetZserviceWorkflowInfoSubResourceData(ZserviceDeploymentUpdateResponseModel.WorkflowInfo)
			d = append(d, &properties)
		}
	}
	return
}

func ZserviceDeploymentUpdateResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"deployment_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"error": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"workflow_info": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*ZserviceWorkflowInfo
			Elem: &schema.Resource{
				Schema: ZserviceWorkflowInfoSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
	}
}

func GetZserviceDeploymentUpdateResponsePropertyFields() (t []string) {
	return []string{
		"deployment_id",
		"error",
		"workflow_info",
	}
}
