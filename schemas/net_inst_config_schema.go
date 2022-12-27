package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate NetInstConfig resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func NetInstConfigModel(d *schema.ResourceData) *models.NetInstConfig {
	clusterID := d.Get("cluster_id").(string)
	description := d.Get("description").(string)
	deviceDefault := d.Get("device_default").(string)
	deviceID := d.Get("device_id").(string)
	dhcp := d.Get("dhcp").(bool)
	dNSList := d.Get("dns_list").([]*models.StaticDNSList) // []*StaticDNSList
	id := d.Get("id").(string)
	var ip *models.DhcpServerConfig // DhcpServerConfig
	ipInterface, ipIsSet := d.GetOk("ip")
	if ipIsSet {
		ipMap := ipInterface.([]interface{})[0].(map[string]interface{})
		ip = DhcpServerConfigModelFromMap(ipMap)
	}
	kind := d.Get("kind").(*models.NetworkInstanceKind) // NetworkInstanceKind
	var lisp *models.LispConfig                         // LispConfig
	lispInterface, lispIsSet := d.GetOk("lisp")
	if lispIsSet {
		lispMap := lispInterface.([]interface{})[0].(map[string]interface{})
		lisp = LispConfigModelFromMap(lispMap)
	}
	name := d.Get("name").(string)
	networkPolicyID := d.Get("network_policy_id").(string)
	oconfig := d.Get("oconfig").(string)
	var opaque *models.NetInstOpaqueConfig // NetInstOpaqueConfig
	opaqueInterface, opaqueIsSet := d.GetOk("opaque")
	if opaqueIsSet {
		opaqueMap := opaqueInterface.([]interface{})[0].(map[string]interface{})
		opaque = NetInstOpaqueConfigModelFromMap(opaqueMap)
	}
	port := d.Get("port").(string)
	portTags := d.Get("port_tags").(map[string]string) // map[string]string
	projectID := d.Get("project_id").(string)
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := d.GetOk("revision")
	if revisionIsSet {
		revisionMap := revisionInterface.([]interface{})[0].(map[string]interface{})
		revision = ObjectRevisionModelFromMap(revisionMap)
	}
	tags := d.Get("tags").(map[string]string) // map[string]string
	title := d.Get("title").(string)
	typeVar := d.Get("type").(*models.NetworkInstanceDhcpType) // NetworkInstanceDhcpType
	return &models.NetInstConfig{
		ClusterID:       clusterID,
		Description:     description,
		DeviceDefault:   deviceDefault,
		DeviceID:        &deviceID, // string true false false
		Dhcp:            dhcp,
		DNSList:         dNSList,
		ID:              id,
		IP:              ip,
		Kind:            kind,
		Lisp:            lisp,
		Name:            &name, // string true false false
		NetworkPolicyID: networkPolicyID,
		Oconfig:         oconfig,
		Opaque:          opaque,
		Port:            &port, // string true false false
		PortTags:        portTags,
		ProjectID:       projectID,
		Revision:        revision,
		Tags:            tags,
		Title:           &title, // string true false false
		Type:            typeVar,
	}
}

func NetInstConfigModelFromMap(m map[string]interface{}) *models.NetInstConfig {
	clusterID := m["cluster_id"].(string)
	description := m["description"].(string)
	deviceDefault := m["device_default"].(string)
	deviceID := m["device_id"].(string)
	dhcp := m["dhcp"].(bool)
	dNSList := m["dns_list"].([]*models.StaticDNSList) // []*StaticDNSList
	id := m["id"].(string)
	var ip *models.DhcpServerConfig // DhcpServerConfig
	ipInterface, ipIsSet := m["ip"]
	if ipIsSet {
		ipMap := ipInterface.([]interface{})[0].(map[string]interface{})
		ip = DhcpServerConfigModelFromMap(ipMap)
	}
	//
	kind := m["kind"].(*models.NetworkInstanceKind) // NetworkInstanceKind
	var lisp *models.LispConfig                     // LispConfig
	lispInterface, lispIsSet := m["lisp"]
	if lispIsSet {
		lispMap := lispInterface.([]interface{})[0].(map[string]interface{})
		lisp = LispConfigModelFromMap(lispMap)
	}
	//
	name := m["name"].(string)
	networkPolicyID := m["network_policy_id"].(string)
	oconfig := m["oconfig"].(string)
	var opaque *models.NetInstOpaqueConfig // NetInstOpaqueConfig
	opaqueInterface, opaqueIsSet := m["opaque"]
	if opaqueIsSet {
		opaqueMap := opaqueInterface.([]interface{})[0].(map[string]interface{})
		opaque = NetInstOpaqueConfigModelFromMap(opaqueMap)
	}
	//
	port := m["port"].(string)
	portTags := m["port_tags"].(map[string]string)
	projectID := m["project_id"].(string)
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := m["revision"]
	if revisionIsSet {
		revisionMap := revisionInterface.([]interface{})[0].(map[string]interface{})
		revision = ObjectRevisionModelFromMap(revisionMap)
	}
	//
	tags := m["tags"].(map[string]string)
	title := m["title"].(string)
	typeVar := m["type"].(*models.NetworkInstanceDhcpType) // NetworkInstanceDhcpType
	return &models.NetInstConfig{
		ClusterID:       clusterID,
		Description:     description,
		DeviceDefault:   deviceDefault,
		DeviceID:        &deviceID,
		Dhcp:            dhcp,
		DNSList:         dNSList,
		ID:              id,
		IP:              ip,
		Kind:            kind,
		Lisp:            lisp,
		Name:            &name,
		NetworkPolicyID: networkPolicyID,
		Oconfig:         oconfig,
		Opaque:          opaque,
		Port:            &port,
		PortTags:        portTags,
		ProjectID:       projectID,
		Revision:        revision,
		Tags:            tags,
		Title:           &title,
		Type:            typeVar,
	}
}

// Update the underlying NetInstConfig resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetInstConfigResourceData(d *schema.ResourceData, m *models.NetInstConfig) {
	d.Set("cluster_id", m.ClusterID)
	d.Set("description", m.Description)
	d.Set("device_default", m.DeviceDefault)
	d.Set("device_id", m.DeviceID)
	d.Set("dhcp", m.Dhcp)
	d.Set("dns_list", SetStaticDNSListSubResourceData(m.DNSList))
	d.Set("id", m.ID)
	d.Set("ip", SetDhcpServerConfigSubResourceData([]*models.DhcpServerConfig{m.IP}))
	d.Set("kind", m.Kind)
	d.Set("lisp", SetLispConfigSubResourceData([]*models.LispConfig{m.Lisp}))
	d.Set("name", m.Name)
	d.Set("network_policy_id", m.NetworkPolicyID)
	d.Set("oconfig", m.Oconfig)
	d.Set("opaque", SetNetInstOpaqueConfigSubResourceData([]*models.NetInstOpaqueConfig{m.Opaque}))
	d.Set("port", m.Port)
	d.Set("port_tags", m.PortTags)
	d.Set("project_id", m.ProjectID)
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("tags", m.Tags)
	d.Set("title", m.Title)
	d.Set("type", m.Type)
}

// Iterate throught and update the NetInstConfig resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetInstConfigSubResourceData(m []*models.NetInstConfig) (d []*map[string]interface{}) {
	for _, NetInstConfigModel := range m {
		if NetInstConfigModel != nil {
			properties := make(map[string]interface{})
			properties["cluster_id"] = NetInstConfigModel.ClusterID
			properties["description"] = NetInstConfigModel.Description
			properties["device_default"] = NetInstConfigModel.DeviceDefault
			properties["device_id"] = NetInstConfigModel.DeviceID
			properties["dhcp"] = NetInstConfigModel.Dhcp
			properties["dns_list"] = SetStaticDNSListSubResourceData(NetInstConfigModel.DNSList)
			properties["id"] = NetInstConfigModel.ID
			properties["ip"] = SetDhcpServerConfigSubResourceData([]*models.DhcpServerConfig{NetInstConfigModel.IP})
			properties["kind"] = NetInstConfigModel.Kind
			properties["lisp"] = SetLispConfigSubResourceData([]*models.LispConfig{NetInstConfigModel.Lisp})
			properties["name"] = NetInstConfigModel.Name
			properties["network_policy_id"] = NetInstConfigModel.NetworkPolicyID
			properties["oconfig"] = NetInstConfigModel.Oconfig
			properties["opaque"] = SetNetInstOpaqueConfigSubResourceData([]*models.NetInstOpaqueConfig{NetInstConfigModel.Opaque})
			properties["port"] = NetInstConfigModel.Port
			properties["port_tags"] = NetInstConfigModel.PortTags
			properties["project_id"] = NetInstConfigModel.ProjectID
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{NetInstConfigModel.Revision})
			properties["tags"] = NetInstConfigModel.Tags
			properties["title"] = NetInstConfigModel.Title
			properties["type"] = NetInstConfigModel.Type
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetInstConfig resource defined in the Terraform configuration
func NetInstConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cluster_id": {
			Description: `id of the Cluster in which the network instance is configured`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"description": {
			Description: `Detailed description of the network instance`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"device_default": {
			Description: `flag to indicate if this is the default network instance for the device`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"device_id": {
			Description: `id of the device on which network instance is created`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"dhcp": {
			Description: `Deprecated`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"dns_list": {
			Description: `List of Static DNS entries`,
			Type:        schema.TypeList, //GoType: []*StaticDNSList
			Elem: &schema.Resource{
				Schema: StaticDNSListSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"id": {
			Description: `System defined universally unique Id of the network instance`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"ip": {
			Description: `Dhcp Server Configuration`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},

		"kind": {
			Description: `Kind of Network Instance ( Local, Switch etc )`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Required: true,
		},

		"lisp": {
			Description: `Lisp Config : read only for now. Deprecated.`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},

		"name": {
			Description: `User defined name of the network instance, unique across the enterprise. Once object is created, name can’t be changed`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"network_policy_id": {
			Description: `id of the network policy to be attached to this network instance`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"oconfig": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"opaque": {
			Description: `Service specific Config`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},

		"port": {
			Description: `name of port mapping in the model`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"port_tags": {
			Description: `Tags are name/value pairs that enable you to categorize resources. Tag names are case insensitive with max_length 512 and min_length 3. Tag values are case sensitive with max_length 256 and min_length 3.`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"project_id": {
			Description: `id of the project in which network instance is created`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"revision": {
			Description: `system defined info for the object`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},

		"tags": {
			Description: `Tags are name/value pairs that enable you to categorize resources. Tag names are case insensitive with max_length 512 and min_length 3. Tag values are case sensitive with max_length 256 and min_length 3.`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"title": {
			Description: `User defined title of the network instance. Title can be changed at any time`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"type": {
			Description: `Type of DHCP for this Network Instance`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the NetInstConfig resource
func GetNetInstConfigPropertyFields() (t []string) {
	return []string{
		"cluster_id",
		"description",
		"device_default",
		"device_id",
		"dhcp",
		"dns_list",
		"id",
		"ip",
		"kind",
		"lisp",
		"name",
		"network_policy_id",
		"oconfig",
		"opaque",
		"port",
		"port_tags",
		"project_id",
		"revision",
		"tags",
		"title",
		"type",
	}
}
