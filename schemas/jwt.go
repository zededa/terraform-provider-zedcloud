package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func JWTInfoModel(d *schema.ResourceData) *models.JWTInfo {
	allowSecInt, _ := d.Get("allow_sec").(int)
	allowSec := int64(allowSecInt)
	dispURL, _ := d.Get("disp_url").(string)
	encrypt, _ := d.Get("encrypt").(bool)
	expireSec, _ := d.Get("expire_sec").(string)
	numInstInt, _ := d.Get("num_inst").(int)
	numInst := int64(numInstInt)
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

func SetJWTInfoResourceData(d *schema.ResourceData, m *models.JWTInfo) {
	d.Set("allow_sec", m.AllowSec)
	d.Set("disp_url", m.DispURL)
	d.Set("encrypt", m.Encrypt)
	d.Set("expire_sec", m.ExpireSec)
	d.Set("num_inst", m.NumInst)
}

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
func JWT() map[string]*schema.Schema {
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
			Computed:    true,
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
