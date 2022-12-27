package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DeviceLisp resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceLispModel(d *schema.ResourceData) *models.DeviceLisp {
	eID := d.Get("e_id").(string)
	eIDHashLen := int64(d.Get("e_id_hash_len").(int))
	clientAddr := d.Get("client_addr").(string)
	eidAllocationPrefix := d.Get("eid_allocation_prefix").(strfmt.Base64)
	eidAllocationPrefixLen := int64(d.Get("eid_allocation_prefix_len").(int))
	lispInstance := int64(d.Get("lisp_instance").(int))
	lispMapServers := d.Get("lisp_map_servers").([]*models.LispServer) // []*LispServer
	mode := d.Get("mode").(string)
	zedServers := d.Get("zed_servers").([]*models.DevZedServer) // []*DevZedServer
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
	eidAllocationPrefix := m["eid_allocation_prefix"].(strfmt.Base64)
	eidAllocationPrefixLen := int64(m["eid_allocation_prefix_len"].(int)) // int64 true false false
	lispInstance := int64(m["lisp_instance"].(int))                       // int64 true false false
	lispMapServers := m["lisp_map_servers"].([]*models.LispServer)        // []*LispServer
	mode := m["mode"].(string)
	zedServers := m["zed_servers"].([]*models.DevZedServer) // []*DevZedServer
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

// Iterate throught and update the DeviceLisp resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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
func DeviceLispSchema() map[string]*schema.Schema {
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
			Description:  `EID allocation prefix`,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Required:     true,
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
				Schema: LispServerSchema(),
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
