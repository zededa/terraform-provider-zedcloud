package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func NetworkInstanceModel(d *schema.ResourceData) *models.NetworkInstance {
	clusterID, _ := d.Get("cluster_id").(string)
	description, _ := d.Get("description").(string)
	deviceDefault, _ := d.Get("device_default").(bool)
	deviceID, _ := d.Get("device_id").(string)
	dhcp, _ := d.Get("dhcp").(bool)
	var dNSList []*models.StaticDNSList // []*StaticDNSList
	dnsListInterface, dnsListIsSet := d.GetOk("dns_list")
	if dnsListIsSet {
		var items []interface{}
		if listItems, isList := dnsListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = dnsListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := StaticDNSListModelFromMap(v.(map[string]interface{}))
			dNSList = append(dNSList, m)
		}
	}
	var edgeNodeCluster *models.NetInstEdgeNodeCluster // NetInstEdgeNodeCluster
	edgeNodeClusterInterface, edgeNodeClusterIsSet := d.GetOk("edge_node_cluster")
	if edgeNodeClusterIsSet && edgeNodeClusterInterface != nil {
		edgeNodeClusterMap := edgeNodeClusterInterface.([]interface{})
		if len(edgeNodeClusterMap) > 0 {
			edgeNodeCluster = NetInstEdgeNodeClusterModelFromMap(edgeNodeClusterMap[0].(map[string]interface{}))
		}
	}
	id, _ := d.Get("id").(string)
	var ip *models.DhcpServerConfig // DhcpServerConfig
	ipInterface, ipIsSet := d.GetOk("ip")
	if ipIsSet && ipInterface != nil {
		ipMap := ipInterface.([]interface{})
		if len(ipMap) > 0 {
			ip = DHCPServerModelFromMap(ipMap[0].(map[string]interface{}))
		}
	}
	var kind *models.NetworkInstanceKind // NetworkInstanceKind
	kindInterface, kindIsSet := d.GetOk("kind")
	if kindIsSet {
		kindModel := kindInterface.(string)
		kind = models.NewNetworkInstanceKind(models.NetworkInstanceKind(kindModel))
	}
	var lisp *models.LispConfig // LispConfig
	lispInterface, lispIsSet := d.GetOk("lisp")
	if lispIsSet && lispInterface != nil {
		lispMap := lispInterface.([]interface{})
		if len(lispMap) > 0 {
			lisp = LispModelFromMap(lispMap[0].(map[string]interface{}))
		}
	}
	mtuInt, _ := d.Get("mtu").(int)
	mtu := int64(mtuInt)
	name, _ := d.Get("name").(string)
	networkPolicyID, _ := d.Get("network_policy_id").(string)
	oconfig, _ := d.Get("oconfig").(string)
	var opaque *models.NetInstOpaqueConfig // NetInstOpaqueConfig
	opaqueInterface, opaqueIsSet := d.GetOk("opaque")
	if opaqueIsSet && opaqueInterface != nil {
		opaqueMap := opaqueInterface.([]interface{})
		if len(opaqueMap) > 0 {
			opaque = NetworkInstanceOpaqueModelFromMap(opaqueMap[0].(map[string]interface{}))
		}
	}
	port, _ := d.Get("port").(string)
	portTags := map[string]string{}
	portTagsInterface, portTagsIsSet := d.GetOk("port_tags")
	if portTagsIsSet {
		portTagsMap := portTagsInterface.(map[string]interface{})
		for k, v := range portTagsMap {
			if v == nil {
				continue
			}
			portTags[k] = v.(string)
		}
	}

	propagateConnectedRoutes, _ := d.Get("propagate_connected_routes").(bool)
	var staticRoutes []*models.StaticIPRoute // []*StaticIPRoute
	staticRoutesInterface, staticRoutesIsSet := d.GetOk("static_routes")
	if staticRoutesIsSet {
		var items []interface{}
		if listItems, isList := staticRoutesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = staticRoutesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := StaticIPRouteModelFromMap(v.(map[string]interface{}))
			staticRoutes = append(staticRoutes, m)
		}
	}

	projectID, _ := d.Get("project_id").(string)
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := d.GetOk("revision")
	if revisionIsSet && revisionInterface != nil {
		revisionMap := revisionInterface.([]interface{})
		if len(revisionMap) > 0 {
			revision = ObjectRevisionModelFromMap(revisionMap[0].(map[string]interface{}))
		}
	}
	tags := map[string]string{}
	tagsInterface, tagsIsSet := d.GetOk("tags")
	if tagsIsSet {
		tagsMap := tagsInterface.(map[string]interface{})
		for k, v := range tagsMap {
			if v == nil {
				continue
			}
			tags[k] = v.(string)
		}
	}

	title, _ := d.Get("title").(string)
	var typeVar *models.NetworkInstanceDhcpType // NetworkInstanceDhcpType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewNetworkInstanceDhcpType(models.NetworkInstanceDhcpType(typeModel))
	}
	return &models.NetworkInstance{
		ClusterID:       			clusterID,
		Description:    			description,
		DeviceDefault:   			deviceDefault,
		DeviceID:        			&deviceID, // string true false false
		Dhcp:            			dhcp,
		DNSList:         			dNSList,
		EdgeNodeCluster:          edgeNodeCluster,
		ID:              			id,
		IP:              			ip,
		Kind:            			kind,
		Lisp:            			lisp,
		Mtu:            			mtu,
		Name:            			&name, // string true false false
		NetworkPolicyID: 			networkPolicyID,
		Oconfig:         			oconfig,
		Opaque:          			opaque,
		Port:            			&port, // string true false false
		PortTags:        			portTags,
		ProjectID:       			projectID,
		PropagateConnectedRoutes: 	propagateConnectedRoutes,
		Revision:        			revision,
		StaticRoutes:	 			staticRoutes,
		Tags:            			tags,
		Title:           			&title, // string true false false
		Type:            			typeVar,
	}
}

func NetworkInstanceModelFromMap(m map[string]interface{}) *models.NetworkInstance {
	clusterID := m["cluster_id"].(string)
	description := m["description"].(string)
	deviceDefault := m["device_default"].(bool)
	deviceID := m["device_id"].(string)
	dhcp := m["dhcp"].(bool)
	var dNSList []*models.StaticDNSList // []*StaticDNSList
	dnsListInterface, dnsListIsSet := m["dns_list"]
	if dnsListIsSet {
		var items []interface{}
		if listItems, isList := dnsListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = dnsListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := StaticDNSListModelFromMap(v.(map[string]interface{}))
			dNSList = append(dNSList, m)
		}
	}
	var edgeNodeCluster *models.NetInstEdgeNodeCluster // NetInstEdgeNodeCluster
	edgeNodeClusterInterface, edgeNodeClusterIsSet := m["edge_node_cluster"]
	if edgeNodeClusterIsSet && edgeNodeClusterInterface != nil {
		edgeNodeClusterMap := edgeNodeClusterInterface.([]interface{})
		if len(edgeNodeClusterMap) > 0 {
			edgeNodeCluster = NetInstEdgeNodeClusterModelFromMap(edgeNodeClusterMap[0].(map[string]interface{}))
		}
	}
	id := m["id"].(string)
	var ip *models.DhcpServerConfig // DhcpServerConfig
	ipInterface, ipIsSet := m["ip"]
	if ipIsSet && ipInterface != nil {
		ipMap := ipInterface.([]interface{})
		if len(ipMap) > 0 {
			ip = DHCPServerModelFromMap(ipMap[0].(map[string]interface{}))
		}
	}
	//
	var kind *models.NetworkInstanceKind // NetworkInstanceKind
	kindInterface, kindIsSet := m["kind"]
	if kindIsSet {
		kindModel := kindInterface.(string)
		kind = models.NewNetworkInstanceKind(models.NetworkInstanceKind(kindModel))
	}
	var lisp *models.LispConfig // LispConfig
	lispInterface, lispIsSet := m["lisp"]
	if lispIsSet && lispInterface != nil {
		lispMap := lispInterface.([]interface{})
		if len(lispMap) > 0 {
			lisp = LispModelFromMap(lispMap[0].(map[string]interface{}))
		}
	}
	//
	mtu := int64(m["mtu"].(int)) // int64
	name := m["name"].(string)
	networkPolicyID := m["network_policy_id"].(string)
	oconfig := m["oconfig"].(string)
	var opaque *models.NetInstOpaqueConfig // NetInstOpaqueConfig
	opaqueInterface, opaqueIsSet := m["opaque"]
	if opaqueIsSet && opaqueInterface != nil {
		opaqueMap := opaqueInterface.([]interface{})
		if len(opaqueMap) > 0 {
			opaque = NetworkInstanceOpaqueModelFromMap(opaqueMap[0].(map[string]interface{}))
		}
	}
	//
	port := m["port"].(string)
	portTags := map[string]string{}
	portTagsInterface, portTagsIsSet := m["port_tags"]
	if portTagsIsSet {
		portTagsMap := portTagsInterface.(map[string]interface{})
		for k, v := range portTagsMap {
			if v == nil {
				continue
			}
			portTags[k] = v.(string)
		}
	}

	propagateConnectedRoutes := m["propagate_connected_routes"].(bool)
	var staticRoutes []*models.StaticIPRoute // []*StaticIPRoute
	staticRoutesInterface, staticRoutesIsSet := m["static_routes"]
	if staticRoutesIsSet {
		var items []interface{}
		if listItems, isList := staticRoutesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = staticRoutesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := StaticIPRouteModelFromMap(v.(map[string]interface{}))
			staticRoutes = append(staticRoutes, m)
		}
	}

	projectID := m["project_id"].(string)
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := m["revision"]
	if revisionIsSet && revisionInterface != nil {
		revisionMap := revisionInterface.([]interface{})
		if len(revisionMap) > 0 {
			revision = ObjectRevisionModelFromMap(revisionMap[0].(map[string]interface{}))
		}
	}
	//
	tags := map[string]string{}
	tagsInterface, tagsIsSet := m["tags"]
	if tagsIsSet {
		tagsMap := tagsInterface.(map[string]interface{})
		for k, v := range tagsMap {
			if v == nil {
				continue
			}
			tags[k] = v.(string)
		}
	}

	title := m["title"].(string)
	var typeVar *models.NetworkInstanceDhcpType // NetworkInstanceDhcpType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewNetworkInstanceDhcpType(models.NetworkInstanceDhcpType(typeModel))
	}
	return &models.NetworkInstance{
		ClusterID:       			clusterID,
		Description:     			description,
		DeviceDefault:   			deviceDefault,
		DeviceID:        			&deviceID,
		Dhcp:            			dhcp,
		DNSList:         			dNSList,
		EdgeNodeCluster:          edgeNodeCluster,
		ID:              			id,
		IP:              			ip,
		Kind:            			kind,
		Lisp:            			lisp,
		Mtu: 		   	 			mtu,
		Name:            			&name,
		NetworkPolicyID: 			networkPolicyID,
		Oconfig:         			oconfig,
		Opaque:          			opaque,
		Port:            			&port,
		PortTags:        			portTags,
		ProjectID:       			projectID,
		PropagateConnectedRoutes: 	propagateConnectedRoutes,
		Revision:       		 	revision,
		StaticRoutes:            	staticRoutes,
		Tags:            			tags,
		Title:        			   	&title,
		Type:          				typeVar,
	}
}

func SetNetworkInstanceResourceData(d *schema.ResourceData, m *models.NetworkInstance) {
	d.Set("cluster_id", m.ClusterID)
	d.Set("description", m.Description)
	d.Set("device_default", m.DeviceDefault)
	d.Set("device_id", m.DeviceID)
	d.Set("dhcp", m.Dhcp)
	d.Set("dns_list", SetStaticDNSListSubResourceData(m.DNSList))
	d.Set("edge_node_cluster", SetNetInstEdgeNodeClusterSubResourceData([]*models.NetInstEdgeNodeCluster{m.EdgeNodeCluster}))
	d.Set("id", m.ID)
	d.Set("ip", SetDHCPServerSubResourceData([]*models.DhcpServerConfig{m.IP}))
	d.Set("kind", m.Kind)
	d.Set("lisp", SetLispSubResourceData([]*models.LispConfig{m.Lisp}))
	d.Set("mtu", m.Mtu)
	d.Set("name", m.Name)
	d.Set("network_policy_id", m.NetworkPolicyID)
	d.Set("oconfig", m.Oconfig)
	d.Set("opaque", SetNetworkInstanceOpaqueSubResourceData([]*models.NetInstOpaqueConfig{m.Opaque}))
	d.Set("port", m.Port)
	d.Set("port_tags", m.PortTags)
	d.Set("project_id", m.ProjectID)
	d.Set("propagate_connected_routes", m.PropagateConnectedRoutes)
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("static_routes", SetStaticIPRouteSubResourceData(m.StaticRoutes))
	d.Set("tags", m.Tags)
	d.Set("title", m.Title)
	d.Set("type", m.Type)
}

func SetNetworkInstanceSubResourceData(m []*models.NetworkInstance) (d []*map[string]interface{}) {
	for _, NetInstConfigModel := range m {
		if NetInstConfigModel != nil {
			properties := make(map[string]interface{})
			properties["cluster_id"] = NetInstConfigModel.ClusterID
			properties["description"] = NetInstConfigModel.Description
			properties["device_default"] = NetInstConfigModel.DeviceDefault
			properties["device_id"] = NetInstConfigModel.DeviceID
			properties["dhcp"] = NetInstConfigModel.Dhcp
			properties["dns_list"] = SetStaticDNSListSubResourceData(NetInstConfigModel.DNSList)
			properties["edge_node_cluster"] = SetNetInstEdgeNodeClusterSubResourceData([]*models.NetInstEdgeNodeCluster{NetInstConfigModel.EdgeNodeCluster})
			properties["id"] = NetInstConfigModel.ID
			properties["ip"] = SetDHCPServerSubResourceData([]*models.DhcpServerConfig{NetInstConfigModel.IP})
			properties["kind"] = NetInstConfigModel.Kind
			properties["lisp"] = SetLispSubResourceData([]*models.LispConfig{NetInstConfigModel.Lisp})
			properties["mtu"] = NetInstConfigModel.Mtu
			properties["name"] = NetInstConfigModel.Name
			properties["network_policy_id"] = NetInstConfigModel.NetworkPolicyID
			properties["oconfig"] = NetInstConfigModel.Oconfig
			properties["opaque"] = SetNetworkInstanceOpaqueSubResourceData([]*models.NetInstOpaqueConfig{NetInstConfigModel.Opaque})
			properties["port"] = NetInstConfigModel.Port
			properties["port_tags"] = NetInstConfigModel.PortTags
			properties["project_id"] = NetInstConfigModel.ProjectID
			properties["propagate_connected_routes"] = NetInstConfigModel.PropagateConnectedRoutes
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{NetInstConfigModel.Revision})
			properties["static_routes"] = SetStaticIPRouteSubResourceData(NetInstConfigModel.StaticRoutes)
			properties["tags"] = NetInstConfigModel.Tags
			properties["title"] = NetInstConfigModel.Title
			properties["type"] = NetInstConfigModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func NetworkInstance() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cluster_id": {
			Description: `ID of the Cluster in which the network instance is configured`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"description": {
			Description: `Detailed description of the network instance`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"device_default": {
			Description: `Flag to indicate if this is the default network instance for the device`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"device_id": {
			Description: 	`ID of the device on which network instance is created`,
			Type:        	schema.TypeString,
			Optional:    	true,
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
				Schema: StaticDNSList(),
			},
			DiffSuppressFunc: diffSuppressDNSListOrder("dns_list"),
			Optional:         true,
			Computed:         true,
		},

		"edge_node_cluster": {
			Description: `Edge Node Cluster`,
			Type:        schema.TypeList, //GoType: NetInstEdgeNodeCluster
			Elem: &schema.Resource{
				Schema: NetInstEdgeNodeClusterSchema(),
			},
			Optional: true,
			DiffSuppressFunc: diffSupressMapInterfaceNonConfigChanges("edge_node_cluster", "id"),
		},

		"id": {
			Description: `System defined universally unique ID of the network instance`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"ip": {
			Description: `DHCP Server Configuration`,
			Type:        schema.TypeList, //GoType: DhcpServerConfig
			Elem: &schema.Resource{
				Schema: DHCPServer(),
			},
			Optional: true,
			Computed: true,
		},

		"kind": {
			Description: `Kind of Network Instance:
NETWORK_INSTANCE_KIND_UNSPECIFIED
NETWORK_INSTANCE_KIND_TRANSPARENT
NETWORK_INSTANCE_KIND_SWITCH
NETWORK_INSTANCE_KIND_LOCAL
NETWORK_INSTANCE_KIND_CLOUD
NETWORK_INSTANCE_KIND_MESH
NETWORK_INSTANCE_KIND_HONEYPOT`,
			Type:     schema.TypeString,
			Required: true,
		},

		"lisp": {
			Description: `Lisp Config : read only for now. Deprecated.`,
			Type:        schema.TypeList, //GoType: LispConfig
			Elem: &schema.Resource{
				Schema: Lisp(),
			},
			Computed: true,
		},

		"mtu": {
			Description: `Maximum transmission unit (MTU) to set for the network instance and all application interfaces connected to it`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"name": {
			Description: `User defined name of the network instance, unique across the enterprise. Once object is created, name can’t be changed`,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
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
			Type:        schema.TypeList, //GoType: NetInstOpaqueConfig
			Elem: &schema.Resource{
				Schema: NetworkInstanceOpaque(),
			},
			Optional: true,
		},

		"port": {
			Description: `name of port mapping in the model`,
			Type:        schema.TypeString,
			Optional:    true,
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
			Computed:    true,
		},

		"propagate_connected_routes": {
			Description: `Automatically propagate connected routes`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"revision": {
			Description: `system defined info for the object`,
			Type:        schema.TypeList, //GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevision(),
			},
			Computed: true,
		},

		"static_routes": {
			Description: `List of Static IP routes`,
			Type:        schema.TypeList, //GoType: []*StaticIPRoute
			Elem: &schema.Resource{
				Schema: StaticIPRouteSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
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
			Description: `Type of DHCP for this Network Instance:
NETWORK_INSTANCE_DHCP_TYPE_V4
NETWORK_INSTANCE_DHCP_TYPE_V6
NETWORK_INSTANCE_DHCP_TYPE_CRYPTOEID
NETWORK_INSTANCE_DHCP_TYPE_CRYPTOV4
NETWORK_INSTANCE_DHCP_TYPE_CRYPTOV6`,
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

// Retrieve property field names for updating the NetInstConfig resource
func GetNetworkInstancePropertyFields() (t []string) {
	return []string{
		"cluster_id",
		"description",
		"device_default",
		"device_id",
		"dhcp",
		"dns_list",
		"edge_node_cluster",
		"id",
		"ip",
		"kind",
		"lisp",
		"mtu",
		"name",
		"network_policy_id",
		"oconfig",
		"opaque",
		"port",
		"port_tags",
		"project_id",
		"propagate_connected_routes",
		"revision",
		"static_routes",
		"tags",
		"title",
		"type",
	}
}
