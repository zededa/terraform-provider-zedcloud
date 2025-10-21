package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// Function to perform the following actions:
// (1) Translate BlobStatus resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func BlobStatusModel(d *schema.ResourceData) *models.BlobStatus {
	hash, _ := d.Get("hash").(string)
	refCount, _ := d.Get("ref_count").(string)
	sizeMB, _ := d.Get("size_m_b").(string)
	var swState *models.SWState // SWState
	swStateInterface, swStateIsSet := d.GetOk("sw_state")
	if swStateIsSet {
		swStateModel := swStateInterface.(string)
		swState = models.NewSWState(models.SWState(swStateModel))
	}
	return &models.BlobStatus{
		Hash:     hash,
		RefCount: refCount,
		SizeMB:   sizeMB,
		SwState:  swState,
	}
}

func BlobStatusModelFromMap(m map[string]interface{}) *models.BlobStatus {
	hash := m["hash"].(string)
	refCount := m["ref_count"].(string)
	sizeMB := m["size_m_b"].(string)
	swState := m["sw_state"].(*models.SWState) // SWState
	return &models.BlobStatus{
		Hash:     hash,
		RefCount: refCount,
		SizeMB:   sizeMB,
		SwState:  swState,
	}
}

// Update the underlying BlobStatus resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetBlobStatusResourceData(d *schema.ResourceData, m *models.BlobStatus) {
	d.Set("hash", m.Hash)
	d.Set("ref_count", m.RefCount)
	d.Set("size_m_b", m.SizeMB)
	d.Set("sw_state", m.SwState)
}

// Iterate through and update the BlobStatus resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetBlobStatusSubResourceData(m []*models.BlobStatus) (d []*map[string]interface{}) {
	for _, BlobStatusModel := range m {
		if BlobStatusModel != nil {
			properties := make(map[string]interface{})
			properties["hash"] = BlobStatusModel.Hash
			properties["ref_count"] = BlobStatusModel.RefCount
			properties["size_m_b"] = BlobStatusModel.SizeMB
			properties["sw_state"] = BlobStatusModel.SwState
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the BlobStatus resource defined in the Terraform configuration
func BlobStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"hash": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"ref_count": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"size_m_b": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"sw_state": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the BlobStatus resource
func GetBlobStatusPropertyFields() (t []string) {
	return []string{
		"hash",
		"ref_count",
		"size_m_b",
		"sw_state",
	}
}
