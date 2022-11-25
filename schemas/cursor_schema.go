package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate Cursor resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func CursorModel(d *schema.ResourceData) *models.Cursor {
	orderBy := d.Get("order_by").(string)
	pageNum := int64(d.Get("page_num").(int))
	pageSize := int64(d.Get("page_size").(int))
	pageToken := d.Get("page_token").(string)
	totalPages := int64(d.Get("total_pages").(int))
	return &models.Cursor{
		OrderBy:    orderBy,
		PageNum:    pageNum,
		PageSize:   pageSize,
		PageToken:  pageToken,
		TotalPages: totalPages,
	}
}

func CursorModelFromMap(m map[string]interface{}) *models.Cursor {
	orderBy := m["order_by"].(string)
	pageNum := int64(m["page_num"].(int))   // int64 false false false
	pageSize := int64(m["page_size"].(int)) // int64 false false false
	pageToken := m["page_token"].(string)
	totalPages := int64(m["total_pages"].(int)) // int64 false false false
	return &models.Cursor{
		OrderBy:    orderBy,
		PageNum:    pageNum,
		PageSize:   pageSize,
		PageToken:  pageToken,
		TotalPages: totalPages,
	}
}

// Update the underlying Cursor resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetCursorResourceData(d *schema.ResourceData, m *models.Cursor) {
	d.Set("order_by", m.OrderBy)
	d.Set("page_num", m.PageNum)
	d.Set("page_size", m.PageSize)
	d.Set("page_token", m.PageToken)
	d.Set("total_pages", m.TotalPages)
}

// Iterate throught and update the Cursor resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetCursorSubResourceData(m []*models.Cursor) (d []*map[string]interface{}) {
	for _, CursorModel := range m {
		if CursorModel != nil {
			properties := make(map[string]interface{})
			properties["order_by"] = CursorModel.OrderBy
			properties["page_num"] = CursorModel.PageNum
			properties["page_size"] = CursorModel.PageSize
			properties["page_token"] = CursorModel.PageToken
			properties["total_pages"] = CursorModel.TotalPages
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the Cursor resource defined in the Terraform configuration
func CursorSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"order_by": {
			Description: `OrderBy helps in sorting the list response`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"page_num": {
			Description: `Page Number`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"page_size": {
			Description: `Defines the page size`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"page_token": {
			Description: `Page Token`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"total_pages": {
			Description: `Total number of pages to be fetched.`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the Cursor resource
func GetCursorPropertyFields() (t []string) {
	return []string{
		"order_by",
		"page_num",
		"page_size",
		"page_token",
		"total_pages",
	}
}
