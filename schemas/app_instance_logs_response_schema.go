package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AppInstanceLogsResponseModel(d *schema.ResourceData) *models.AppInstanceLogsResponse {
	var content *models.AppInstanceLogsQueryResponse // AppInstanceLogsQueryResponse
	contentInterface, contentIsSet := d.GetOk("content")
	if contentIsSet && contentInterface != nil {
		contentMap := contentInterface.([]interface{})
		if len(contentMap) > 0 {
			content = AppInstanceLogsQueryResponseModelFromMap(contentMap[0].(map[string]interface{}))
		}
	}
	var result *models.ZsrvResponse // ZsrvResponse
	resultInterface, resultIsSet := d.GetOk("result")
	if resultIsSet && resultInterface != nil {
		resultMap := resultInterface.([]interface{})
		if len(resultMap) > 0 {
			result = ZsrvResponseModelFromMap(resultMap[0].(map[string]interface{}))
		}
	}
	return &models.AppInstanceLogsResponse{
		Content: content,
		Result:  result,
	}
}

func AppInstanceLogsResponseModelFromMap(m map[string]interface{}) *models.AppInstanceLogsResponse {
	var content *models.AppInstanceLogsQueryResponse // AppInstanceLogsQueryResponse
	contentInterface, contentIsSet := m["content"]
	if contentIsSet && contentInterface != nil {
		contentMap := contentInterface.([]interface{})
		if len(contentMap) > 0 {
			content = AppInstanceLogsQueryResponseModelFromMap(contentMap[0].(map[string]interface{}))
		}
	}
	//
	var result *models.ZsrvResponse // ZsrvResponse
	resultInterface, resultIsSet := m["result"]
	if resultIsSet && resultInterface != nil {
		resultMap := resultInterface.([]interface{})
		if len(resultMap) > 0 {
			result = ZsrvResponseModelFromMap(resultMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.AppInstanceLogsResponse{
		Content: content,
		Result:  result,
	}
}

// Update the underlying AppInstanceLogsResponse resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppInstanceLogsResponseResourceData(d *schema.ResourceData, m *models.AppInstanceLogsResponse) {
	d.Set("content", SetAppInstanceLogsQueryResponseSubResourceData([]*models.AppInstanceLogsQueryResponse{m.Content}))
	d.Set("result", SetZsrvResponseSubResourceData([]*models.ZsrvResponse{m.Result}))
}

// Iterate through and update the AppInstanceLogsResponse resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppInstanceLogsResponseSubResourceData(m []*models.AppInstanceLogsResponse) (d []*map[string]interface{}) {
	for _, AppInstanceLogsResponseModel := range m {
		if AppInstanceLogsResponseModel != nil {
			properties := make(map[string]interface{})
			properties["content"] = SetAppInstanceLogsQueryResponseSubResourceData([]*models.AppInstanceLogsQueryResponse{AppInstanceLogsResponseModel.Content})
			properties["result"] = SetZsrvResponseSubResourceData([]*models.ZsrvResponse{AppInstanceLogsResponseModel.Result})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AppInstanceLogsResponse resource defined in the Terraform configuration
func AppInstanceLogsResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"content": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AppInstanceLogsQueryResponse
			Elem: &schema.Resource{
				Schema: AppInstanceLogsQueryResponseSchema(),
			},
			Optional: true,
		},

		"result": {
			Description: ``,
			Type:        schema.TypeList, //GoType: ZsrvResponse
			Elem: &schema.Resource{
				Schema: ZsrvResponseSchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the AppInstanceLogsResponse resource
func GetAppInstanceLogsResponsePropertyFields() (t []string) {
	return []string{
		"content",
		"result",
	}
}
