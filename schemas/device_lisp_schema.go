package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func DeviceLispModel(d *schema.ResourceData) *models.DeviceLisp {
	eID, _ := d.Get("e_id").(string)
	eIDHashLenInt, _ := d.Get("e_id_hash_len").(int)
	eIDHashLen := int64(eIDHashLenInt)
	clientAddr, _ := d.Get("client_addr").(string)
	eidAllocationPrefix, _ := d.Get("eid_allocation_prefix").(string)
	eidAllocationPrefixLenInt, _ := d.Get("eid_allocation_prefix_len").(int)
	eidAllocationPrefixLen := int64(eidAllocationPrefixLenInt)
	lispInstanceInt, _ := d.Get("lisp_instance").(int)
	lispInstance := int64(lispInstanceInt)
	var lispMapServers []*models.LispServer // []*LispServer
	lispMapServersInterface, lispMapServersIsSet := d.GetOk("lisp_map_servers")
	if lispMapServersIsSet {
		for _, v := range lispMapServersInterface.([]interface{}) {
			if v == nil {
				continue
			}
			m := LispServerModelFromMap(v.(map[string]interface{}))
			lispMapServers = append(lispMapServers, m)
		}
	}
	mode, _ := d.Get("mode").(string)
	var zedServers []*models.DevZedServer // []*DevZedServer
	zedServersInterface, zedServersIsSet := d.GetOk("zed_servers")
	if zedServersIsSet {
		for _, v := range zedServersInterface.([]interface{}) {
			if v == nil {
				continue
			}
			m := DevZedServerModelFromMap(v.(map[string]interface{}))
			zedServers = append(zedServers, m)
		}
	}
	return &models.DeviceLisp{
		EID:                    &eID,                    // string true false false
		EIDHashLen:             &eIDHashLen,             // int64 true false false
		ClientAddr:             &clientAddr,             // string true false false
		EidAllocationPrefix:    &eidAllocationPrefix,    // strfmt.Base64 true false false
		EidAllocationPrefixLen: &eidAllocationPrefixLen, // int64 true false false
		LispInstance:           &lispInstance,           // int64 true false false
		LispMapServers:         lispMapServers,
		Mode:                   &mode, // string true false false
		ZedServers:             zedServers,
	}
}

func DeviceLispModelFromMap(m map[string]interface{}) *models.DeviceLisp {
	eID := m["e_id"].(string)
	eIDHashLen := int64(m["e_id_hash_len"].(int)) // int64 true false false
	clientAddr := m["client_addr"].(string)
	eidAllocationPrefix := m["eid_allocation_prefix"].(string)
	eidAllocationPrefixLen := int64(m["eid_allocation_prefix_len"].(int)) // int64 true false false
	lispInstance := int64(m["lisp_instance"].(int))                       // int64 true false false
	var lispMapServers []*models.LispServer                               // []*LispServer
	lispMapServersInterface, lispMapServersIsSet := m["lisp_map_servers"]
	if lispMapServersIsSet {
		for _, v := range lispMapServersInterface.([]interface{}) {
			if v == nil {
				continue
			}
			m := LispServerModelFromMap(v.(map[string]interface{}))
			lispMapServers = append(lispMapServers, m)
		}
	}
	mode := m["mode"].(string)
	var zedServers []*models.DevZedServer // []*DevZedServer
	zedServersInterface, zedServersIsSet := m["zed_servers"]
	if zedServersIsSet {
		for _, v := range zedServersInterface.([]interface{}) {
			if v == nil {
				continue
			}
			m := DevZedServerModelFromMap(v.(map[string]interface{}))
			zedServers = append(zedServers, m)
		}
	}
	return &models.DeviceLisp{
		EID:                    &eID,
		EIDHashLen:             &eIDHashLen,
		ClientAddr:             &clientAddr,
		EidAllocationPrefix:    &eidAllocationPrefix,
		EidAllocationPrefixLen: &eidAllocationPrefixLen,
		LispInstance:           &lispInstance,
		LispMapServers:         lispMapServers,
		Mode:                   &mode,
		ZedServers:             zedServers,
	}
}

// Update the underlying DeviceLisp resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceLispResourceData(d *schema.ResourceData, m *models.DeviceLisp) {
	d.Set("e_id", m.EID)
	d.Set("e_id_hash_len", m.EIDHashLen)
	d.Set("client_addr", m.ClientAddr)
	d.Set("eid_allocation_prefix", m.EidAllocationPrefix)
	d.Set("eid_allocation_prefix_len", m.EidAllocationPrefixLen)
	d.Set("lisp_instance", m.LispInstance)
	d.Set("lisp_map_servers", SetLispServerSubResourceData(m.LispMapServers))
	d.Set("mode", m.Mode)
	d.Set("zed_servers", SetDevZedServerSubResourceData(m.ZedServers))
}

func SetDeviceLispSubResourceData(m []*models.DeviceLisp) (d []*map[string]interface{}) {
	for _, DeviceLispModel := range m {
		if DeviceLispModel != nil {
			properties := make(map[string]interface{})
			properties["e_id"] = DeviceLispModel.EID
			properties["e_id_hash_len"] = DeviceLispModel.EIDHashLen
			properties["client_addr"] = DeviceLispModel.ClientAddr
			properties["eid_allocation_prefix"] = DeviceLispModel.EidAllocationPrefix
			properties["eid_allocation_prefix_len"] = DeviceLispModel.EidAllocationPrefixLen
			properties["lisp_instance"] = DeviceLispModel.LispInstance
			properties["lisp_map_servers"] = SetLispServerSubResourceData(DeviceLispModel.LispMapServers)
			properties["mode"] = DeviceLispModel.Mode
			properties["zed_servers"] = SetDevZedServerSubResourceData(DeviceLispModel.ZedServers)
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceLisp resource defined in the Terraform configuration
func DeviceLisp() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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

		"client_addr": {
			Description: `Client Address`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"eid_allocation_prefix": {
			Description: `EID allocation prefix`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"eid_allocation_prefix_len": {
			Description: `EID allocation prefix length`,
			Type:        schema.TypeInt,
			Required:    true,
		},

		"lisp_instance": {
			Description: `LISP instance`,
			Type:        schema.TypeInt,
			Required:    true,
		},

		"lisp_map_servers": {
			Description: `List of Lisp servers`,
			Type:        schema.TypeList, //GoType: []*LispServer
			Elem: &schema.Resource{
				Schema: LispServer(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Required: true,
		},

		"mode": {
			Description: `TEMP : flag to indicate which version of LISP data plane should be running on the device`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"zed_servers": {
			Description: `Zed development servers`,
			Type:        schema.TypeList, //GoType: []*DevZedServer
			Elem: &schema.Resource{
				Schema: DevZedServerSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Required: true,
		},
	}
}

// Retrieve property field names for updating the DeviceLisp resource
func GetDeviceLispPropertyFields() (t []string) {
	return []string{
		"e_id",
		"e_id_hash_len",
		"client_addr",
		"eid_allocation_prefix",
		"eid_allocation_prefix_len",
		"lisp_instance",
		"lisp_map_servers",
		"mode",
		"zed_servers",
	}
}
