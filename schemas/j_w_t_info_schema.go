package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate JWTInfo resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func JWTInfoModel(d *schema.ResourceData) *models.JWTInfo {
	allowSec := int64(d.Get("allow_sec").(int))
	dispURL := d.Get("disp_url").(string)
	encrypt := d.Get("encrypt").(bool)
	expireSec := d.Get("expire_sec").(string)
	numInst := int64(d.Get("num_inst").(int))
	return &models.JWTInfo{
		AllowSec:  allowSec,
		DispURL:   dispURL,
		Encrypt:   encrypt,
		ExpireSec: expireSec,
		NumInst:   numInst,
	}
}

func JWTInfoModelFromMap(m map[string]interface{}) *models.JWTInfo {
	allowSec := int64(m["allow_sec"].(int)) // int64 false false false
	dispURL := m["disp_url"].(string)
	encrypt := m["encrypt"].(bool)
	expireSec := m["expire_sec"].(string)
	numInst := int64(m["num_inst"].(int)) // int64 false false false
	return &models.JWTInfo{
		AllowSec:  allowSec,
		DispURL:   dispURL,
		Encrypt:   encrypt,
		ExpireSec: expireSec,
		NumInst:   numInst,
	}
}

// Update the underlying JWTInfo resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetJWTInfoResourceData(d *schema.ResourceData, m *models.JWTInfo) {
	d.Set("allow_sec", m.AllowSec)
	d.Set("disp_url", m.DispURL)
	d.Set("encrypt", m.Encrypt)
	d.Set("expire_sec", m.ExpireSec)
	d.Set("num_inst", m.NumInst)
}

// Iterate throught and update the JWTInfo resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetJWTInfoSubResourceData(m []*models.JWTInfo) (d []*map[string]interface{}) {
	for _, JWTInfoModel := range m {
		if JWTInfoModel != nil {
			properties := make(map[string]interface{})
			properties["allow_sec"] = JWTInfoModel.AllowSec
			properties["disp_url"] = JWTInfoModel.DispURL
			properties["encrypt"] = JWTInfoModel.Encrypt
			properties["expire_sec"] = JWTInfoModel.ExpireSec
			properties["num_inst"] = JWTInfoModel.NumInst
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the JWTInfo resource defined in the Terraform configuration
func JWTInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allow_sec": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"disp_url": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"encrypt": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"expire_sec": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"num_inst": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the JWTInfo resource
func GetJWTInfoPropertyFields() (t []string) {
	return []string{
		"allow_sec",
		"disp_url",
		"encrypt",
		"expire_sec",
		"num_inst",
	}
}
