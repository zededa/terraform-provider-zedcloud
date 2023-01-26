package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate EIDRegister resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func EIDRegisterModel(d *schema.ResourceData) *models.EIDRegister {
	appCert, _ := d.Get("app_cert").(strfmt.Base64)
	appPrivateKey, _ := d.Get("app_private_key").(strfmt.Base64)
	appPublicKey, _ := d.Get("app_public_key").(strfmt.Base64)
	displayName, _ := d.Get("display_name").(string)
	eID, _ := d.Get("e_id").(string)
	eIDHashLenInt, _ := d.Get("e_id_hash_len").(int)
	eIDHashLen := int64(eIDHashLenInt)
	lispInstanceInt, _ := d.Get("lisp_instance").(int)
	lispInstance := int64(lispInstanceInt)
	lispMapServers, _ := d.Get("lisp_map_servers").([]*models.LispServer) // []*LispServer
	lispSignature, _ := d.Get("lisp_signature").(string)
	uuid, _ := d.Get("uuid").(string)
	return &models.EIDRegister{
		AppCert:        &appCert,       // strfmt.Base64 true false false
		AppPrivateKey:  &appPrivateKey, // strfmt.Base64 true false false
		AppPublicKey:   &appPublicKey,  // strfmt.Base64 true false false
		DisplayName:    &displayName,   // string true false false
		EID:            &eID,           // string true false false
		EIDHashLen:     &eIDHashLen,    // int64 true false false
		LispInstance:   &lispInstance,  // int64 true false false
		LispMapServers: lispMapServers,
		LispSignature:  &lispSignature, // string true false false
		UUID:           &uuid,          // string true false false
	}
}

func EIDRegisterModelFromMap(m map[string]interface{}) *models.EIDRegister {
	appCert := m["app_cert"].(strfmt.Base64)
	appPrivateKey := m["app_private_key"].(strfmt.Base64)
	appPublicKey := m["app_public_key"].(strfmt.Base64)
	displayName := m["display_name"].(string)
	eID := m["e_id"].(string)
	eIDHashLen := int64(m["e_id_hash_len"].(int))                  // int64 true false false
	lispInstance := int64(m["lisp_instance"].(int))                // int64 true false false
	lispMapServers := m["lisp_map_servers"].([]*models.LispServer) // []*LispServer
	lispSignature := m["lisp_signature"].(string)
	uuid := m["uuid"].(string)
	return &models.EIDRegister{
		AppCert:        &appCert,
		AppPrivateKey:  &appPrivateKey,
		AppPublicKey:   &appPublicKey,
		DisplayName:    &displayName,
		EID:            &eID,
		EIDHashLen:     &eIDHashLen,
		LispInstance:   &lispInstance,
		LispMapServers: lispMapServers,
		LispSignature:  &lispSignature,
		UUID:           &uuid,
	}
}

// Update the underlying EIDRegister resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetEIDRegisterResourceData(d *schema.ResourceData, m *models.EIDRegister) {
	d.Set("app_cert", m.AppCert.String())
	d.Set("app_private_key", m.AppPrivateKey.String())
	d.Set("app_public_key", m.AppPublicKey.String())
	d.Set("display_name", m.DisplayName)
	d.Set("e_id", m.EID)
	d.Set("e_id_hash_len", m.EIDHashLen)
	d.Set("lisp_instance", m.LispInstance)
	d.Set("lisp_map_servers", SetLispServerSubResourceData(m.LispMapServers))
	d.Set("lisp_signature", m.LispSignature)
	d.Set("uuid", m.UUID)
}

// Iterate through and update the EIDRegister resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetEIDRegisterSubResourceData(m []*models.EIDRegister) (d []*map[string]interface{}) {
	for _, EIDRegisterModel := range m {
		if EIDRegisterModel != nil {
			properties := make(map[string]interface{})
			properties["app_cert"] = EIDRegisterModel.AppCert.String()
			properties["app_private_key"] = EIDRegisterModel.AppPrivateKey.String()
			properties["app_public_key"] = EIDRegisterModel.AppPublicKey.String()
			properties["display_name"] = EIDRegisterModel.DisplayName
			properties["e_id"] = EIDRegisterModel.EID
			properties["e_id_hash_len"] = EIDRegisterModel.EIDHashLen
			properties["lisp_instance"] = EIDRegisterModel.LispInstance
			properties["lisp_map_servers"] = SetLispServerSubResourceData(EIDRegisterModel.LispMapServers)
			properties["lisp_signature"] = EIDRegisterModel.LispSignature
			properties["uuid"] = EIDRegisterModel.UUID
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the EIDRegister resource defined in the Terraform configuration
func EIDRegisterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_cert": {
			Description: `app certificate`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"app_private_key": {
			Description: `App private key`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"app_public_key": {
			Description: `App public key`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"display_name": {
			Description: `Display name`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"e_id": {
			Description: `EID`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"e_id_hash_len": {
			Description: `EID hash length`,
			Type:        schema.TypeInt,
			Required:    true,
		},

		"lisp_instance": {
			Description: `Lisp Instance`,
			Type:        schema.TypeInt,
			Required:    true,
		},

		"lisp_map_servers": {
			Description: `Lisp Map Server`,
			Type:        schema.TypeList, //GoType: []*LispServer
			Elem: &schema.Resource{
				Schema: LispServer(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Required: true,
		},

		"lisp_signature": {
			Description: `Lisp Signature`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"uuid": {
			Description: `UUID`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

// Retrieve property field names for updating the EIDRegister resource
func GetEIDRegisterPropertyFields() (t []string) {
	return []string{
		"app_cert",
		"app_private_key",
		"app_public_key",
		"display_name",
		"e_id",
		"e_id_hash_len",
		"lisp_instance",
		"lisp_map_servers",
		"lisp_signature",
		"uuid",
	}
}
