package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func EIDRegisterModel(d *schema.ResourceData) *models.EIDRegister {
	appCert, _ := d.Get("app_cert").(string)
	appPrivateKey, _ := d.Get("app_private_key").(string)
	appPublicKey, _ := d.Get("app_public_key").(string)
	displayName, _ := d.Get("display_name").(string)
	eID, _ := d.Get("e_id").(string)
	eIDHashLenInt, _ := d.Get("e_id_hash_len").(int)
	eIDHashLen := int64(eIDHashLenInt)
	lispInstanceInt, _ := d.Get("lisp_instance").(int)
	lispInstance := int64(lispInstanceInt)
	var lispMapServers []*models.LispServer // []*LispServer
	LispMapServersInterface, LispMapServersIsSet := d.GetOk("lisp_map_servers")
	if LispMapServersIsSet {
		var items []interface{}
		if listItems, isList := LispMapServersInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = LispMapServersInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := LispServerModelFromMap(v.(map[string]interface{}))
			lispMapServers = append(lispMapServers, m)
		}
	}
	lispSignature, _ := d.Get("lisp_signature").(string)
	uuid, _ := d.Get("uuid").(string)
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

func EIDRegisterModelFromMap(m map[string]interface{}) *models.EIDRegister {
	appCert := m["app_cert"].(string)
	appPrivateKey := m["app_private_key"].(string)
	appPublicKey := m["app_public_key"].(string)
	displayName := m["display_name"].(string)
	eID := m["e_id"].(string)
	eIDHashLen := int64(m["e_id_hash_len"].(int))
	lispInstance := int64(m["lisp_instance"].(int))
	var lispMapServers []*models.LispServer
	LispMapServersInterface, LispMapServersIsSet := m["lisp_map_servers"]
	if LispMapServersIsSet {
		var items []interface{}
		if listItems, isList := LispMapServersInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = LispMapServersInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := LispServerModelFromMap(v.(map[string]interface{}))
			lispMapServers = append(lispMapServers, m)
		}
	}
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

func SetEIDRegisterResourceData(d *schema.ResourceData, m *models.EIDRegister) {
	d.Set("app_cert", m.AppCert)
	d.Set("app_private_key", m.AppPrivateKey)
	d.Set("app_public_key", m.AppPublicKey)
	d.Set("display_name", m.DisplayName)
	d.Set("e_id", m.EID)
	d.Set("e_id_hash_len", m.EIDHashLen)
	d.Set("lisp_instance", m.LispInstance)
	d.Set("lisp_map_servers", SetLispServerSubResourceData(m.LispMapServers))
	d.Set("lisp_signature", m.LispSignature)
	d.Set("uuid", m.UUID)
}

func SetEIDRegisterSubResourceData(m []*models.EIDRegister) (d []*map[string]interface{}) {
	for _, EIDRegisterModel := range m {
		if EIDRegisterModel != nil {
			properties := make(map[string]interface{})
			properties["app_cert"] = EIDRegisterModel.AppCert
			properties["app_private_key"] = EIDRegisterModel.AppPrivateKey
			properties["app_public_key"] = EIDRegisterModel.AppPublicKey
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
func EIDRegister() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_cert": {
			Description: `app certificate`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"app_private_key": {
			Description: `App private key`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"app_public_key": {
			Description: `App public key`,
			Type:        schema.TypeString,
			Computed:    true,
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
