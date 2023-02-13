package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func TopTalkersResponseItemModel(d *schema.ResourceData) *models.TopTalkersResponseItem {
	total, _ := d.Get("total").(string)
	localAddr, _ := d.Get("local_addr").(string)
	localPortInt, _ := d.Get("local_port").(int)
	localPort := int64(localPortInt)
	protoInt, _ := d.Get("proto").(int)
	proto := int64(protoInt)
	remoteAddr, _ := d.Get("remote_addr").(string)
	remotePortInt, _ := d.Get("remote_port").(int)
	remotePort := int64(remotePortInt)
	rxTotal, _ := d.Get("rx_total").(string)
	txTotal, _ := d.Get("tx_total").(string)
	return &models.TopTalkersResponseItem{
		Total:      total,
		LocalAddr:  localAddr,
		LocalPort:  localPort,
		Proto:      proto,
		RemoteAddr: remoteAddr,
		RemotePort: remotePort,
		RxTotal:    rxTotal,
		TxTotal:    txTotal,
	}
}

func TopTalkersResponseItemModelFromMap(m map[string]interface{}) *models.TopTalkersResponseItem {
	total := m["total"].(string)
	localAddr := m["local_addr"].(string)
	localPort := int64(m["local_port"].(int)) // int64
	proto := int64(m["proto"].(int))          // int64
	remoteAddr := m["remote_addr"].(string)
	remotePort := int64(m["remote_port"].(int)) // int64
	rxTotal := m["rx_total"].(string)
	txTotal := m["tx_total"].(string)
	return &models.TopTalkersResponseItem{
		Total:      total,
		LocalAddr:  localAddr,
		LocalPort:  localPort,
		Proto:      proto,
		RemoteAddr: remoteAddr,
		RemotePort: remotePort,
		RxTotal:    rxTotal,
		TxTotal:    txTotal,
	}
}

// Update the underlying TopTalkersResponseItem resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetTopTalkersResponseItemResourceData(d *schema.ResourceData, m *models.TopTalkersResponseItem) {
	d.Set("total", m.Total)
	d.Set("local_addr", m.LocalAddr)
	d.Set("local_port", m.LocalPort)
	d.Set("proto", m.Proto)
	d.Set("remote_addr", m.RemoteAddr)
	d.Set("remote_port", m.RemotePort)
	d.Set("rx_total", m.RxTotal)
	d.Set("tx_total", m.TxTotal)
}

// Iterate through and update the TopTalkersResponseItem resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetTopTalkersResponseItemSubResourceData(m []*models.TopTalkersResponseItem) (d []*map[string]interface{}) {
	for _, TopTalkersResponseItemModel := range m {
		if TopTalkersResponseItemModel != nil {
			properties := make(map[string]interface{})
			properties["total"] = TopTalkersResponseItemModel.Total
			properties["local_addr"] = TopTalkersResponseItemModel.LocalAddr
			properties["local_port"] = TopTalkersResponseItemModel.LocalPort
			properties["proto"] = TopTalkersResponseItemModel.Proto
			properties["remote_addr"] = TopTalkersResponseItemModel.RemoteAddr
			properties["remote_port"] = TopTalkersResponseItemModel.RemotePort
			properties["rx_total"] = TopTalkersResponseItemModel.RxTotal
			properties["tx_total"] = TopTalkersResponseItemModel.TxTotal
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the TopTalkersResponseItem resource defined in the Terraform configuration
func TopTalkersResponseItemSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"total": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"local_addr": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"local_port": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"proto": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"remote_addr": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"remote_port": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"rx_total": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"tx_total": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the TopTalkersResponseItem resource
func GetTopTalkersResponseItemPropertyFields() (t []string) {
	return []string{
		"total",
		"local_addr",
		"local_port",
		"proto",
		"remote_addr",
		"remote_port",
		"rx_total",
		"tx_total",
	}
}
