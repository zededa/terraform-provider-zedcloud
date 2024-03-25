package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func VCEStatusModel(d *schema.ResourceData) *models.VCEStatus {
	var vceStatusDetail *models.VCEStatusDetail // VCEStatusDetail
	vceStatusDetailInterface, vceStatusDetailIsSet := d.GetOk("vce_status_detail")
	if vceStatusDetailIsSet && vceStatusDetailInterface != nil {
		vceStatusDetailMap := vceStatusDetailInterface.([]interface{})
		if len(vceStatusDetailMap) > 0 {
			vceStatusDetail = VCEStatusDetailModelFromMap(vceStatusDetailMap[0].(map[string]interface{}))
		}
	}
	return &models.VCEStatus{
		VceStatusDetail: vceStatusDetail,
	}
}

func VCEStatusModelFromMap(m map[string]interface{}) *models.VCEStatus {
	var vceStatusDetail *models.VCEStatusDetail // VCEStatusDetail
	vceStatusDetailInterface, vceStatusDetailIsSet := m["vce_status_detail"]
	if vceStatusDetailIsSet && vceStatusDetailInterface != nil {
		vceStatusDetailMap := vceStatusDetailInterface.([]interface{})
		if len(vceStatusDetailMap) > 0 {
			vceStatusDetail = VCEStatusDetailModelFromMap(vceStatusDetailMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.VCEStatus{
		VceStatusDetail: vceStatusDetail,
	}
}

// Update the underlying VCEStatus resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetVCEStatusResourceData(d *schema.ResourceData, m *models.VCEStatus) {
	d.Set("vce_status_detail", SetVCEStatusDetailSubResourceData([]*models.VCEStatusDetail{m.VceStatusDetail}))
}

// Iterate through and update the VCEStatus resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetVCEStatusSubResourceData(m []*models.VCEStatus) (d []*map[string]interface{}) {
	for _, VCEStatusModel := range m {
		if VCEStatusModel != nil {
			properties := make(map[string]interface{})
			properties["vce_status_detail"] = SetVCEStatusDetailSubResourceData([]*models.VCEStatusDetail{VCEStatusModel.VceStatusDetail})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the VCEStatus resource defined in the Terraform configuration
func VCEStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"vce_status_detail": {
			Description: ``,
			Type:        schema.TypeList, //GoType: VCEStatusDetail
			Elem: &schema.Resource{
				Schema: VCEStatusDetailSchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the VCEStatus resource
func GetVCEStatusPropertyFields() (t []string) {
	return []string{
		"vce_status_detail",
	}
}
