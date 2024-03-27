package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// Function to perform the following actions:
// (1) Translate NetworkCounters resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func NetworkCountersModel(d *schema.ResourceData) *models.NetworkCounters {
	ifName, _ := d.Get("if_name").(string)
	rxACLDrops, _ := d.Get("rx_acl_drops").(uint64)
	rxACLRateLimitDrops, _ := d.Get("rx_acl_rate_limit_drops").(uint64)
	rxBytes, _ := d.Get("rx_bytes").(uint64)
	rxDrops, _ := d.Get("rx_drops").(uint64)
	rxErrors, _ := d.Get("rx_errors").(uint64)
	rxPkts, _ := d.Get("rx_pkts").(uint64)
	txACLDrops, _ := d.Get("tx_acl_drops").(uint64)
	txACLRateLimitDrops, _ := d.Get("tx_acl_rate_limit_drops").(uint64)
	txBytes, _ := d.Get("tx_bytes").(uint64)
	txDrops, _ := d.Get("tx_drops").(uint64)
	txErrors, _ := d.Get("tx_errors").(uint64)
	txPkts, _ := d.Get("tx_pkts").(uint64)
	return &models.NetworkCounters{
		IfName:              ifName,
		RxACLDrops:          rxACLDrops,
		RxACLRateLimitDrops: rxACLRateLimitDrops,
		RxBytes:             rxBytes,
		RxDrops:             rxDrops,
		RxErrors:            rxErrors,
		RxPkts:              rxPkts,
		TxACLDrops:          txACLDrops,
		TxACLRateLimitDrops: txACLRateLimitDrops,
		TxBytes:             txBytes,
		TxDrops:             txDrops,
		TxErrors:            txErrors,
		TxPkts:              txPkts,
	}
}

func NetworkCountersModelFromMap(m map[string]interface{}) *models.NetworkCounters {
	ifName := m["if_name"].(string)
	rxACLDrops := m["rx_acl_drops"].(uint64)
	rxACLRateLimitDrops := m["rx_acl_rate_limit_drops"].(uint64)
	rxBytes := m["rx_bytes"].(uint64)
	rxDrops := m["rx_drops"].(uint64)
	rxErrors := m["rx_errors"].(uint64)
	rxPkts := m["rx_pkts"].(uint64)
	txACLDrops := m["tx_acl_drops"].(uint64)
	txACLRateLimitDrops := m["tx_acl_rate_limit_drops"].(uint64)
	txBytes := m["tx_bytes"].(uint64)
	txDrops := m["tx_drops"].(uint64)
	txErrors := m["tx_errors"].(uint64)
	txPkts := m["tx_pkts"].(uint64)
	return &models.NetworkCounters{
		IfName:              ifName,
		RxACLDrops:          rxACLDrops,
		RxACLRateLimitDrops: rxACLRateLimitDrops,
		RxBytes:             rxBytes,
		RxDrops:             rxDrops,
		RxErrors:            rxErrors,
		RxPkts:              rxPkts,
		TxACLDrops:          txACLDrops,
		TxACLRateLimitDrops: txACLRateLimitDrops,
		TxBytes:             txBytes,
		TxDrops:             txDrops,
		TxErrors:            txErrors,
		TxPkts:              txPkts,
	}
}

// Update the underlying NetworkCounters resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetworkCountersResourceData(d *schema.ResourceData, m *models.NetworkCounters) {
	d.Set("if_name", m.IfName)
	d.Set("rx_acl_drops", m.RxACLDrops)
	d.Set("rx_acl_rate_limit_drops", m.RxACLRateLimitDrops)
	d.Set("rx_bytes", m.RxBytes)
	d.Set("rx_drops", m.RxDrops)
	d.Set("rx_errors", m.RxErrors)
	d.Set("rx_pkts", m.RxPkts)
	d.Set("tx_acl_drops", m.TxACLDrops)
	d.Set("tx_acl_rate_limit_drops", m.TxACLRateLimitDrops)
	d.Set("tx_bytes", m.TxBytes)
	d.Set("tx_drops", m.TxDrops)
	d.Set("tx_errors", m.TxErrors)
	d.Set("tx_pkts", m.TxPkts)
}

// Iterate through and update the NetworkCounters resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetworkCountersSubResourceData(m []*models.NetworkCounters) (d []*map[string]interface{}) {
	for _, NetworkCountersModel := range m {
		if NetworkCountersModel != nil {
			properties := make(map[string]interface{})
			properties["if_name"] = NetworkCountersModel.IfName
			properties["rx_acl_drops"] = NetworkCountersModel.RxACLDrops
			properties["rx_acl_rate_limit_drops"] = NetworkCountersModel.RxACLRateLimitDrops
			properties["rx_bytes"] = NetworkCountersModel.RxBytes
			properties["rx_drops"] = NetworkCountersModel.RxDrops
			properties["rx_errors"] = NetworkCountersModel.RxErrors
			properties["rx_pkts"] = NetworkCountersModel.RxPkts
			properties["tx_acl_drops"] = NetworkCountersModel.TxACLDrops
			properties["tx_acl_rate_limit_drops"] = NetworkCountersModel.TxACLRateLimitDrops
			properties["tx_bytes"] = NetworkCountersModel.TxBytes
			properties["tx_drops"] = NetworkCountersModel.TxDrops
			properties["tx_errors"] = NetworkCountersModel.TxErrors
			properties["tx_pkts"] = NetworkCountersModel.TxPkts
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetworkCounters resource defined in the Terraform configuration
func NetworkCountersSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"if_name": {
			Description: `ifName`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"rx_acl_drops": {
			Description: `Rx ACL Rate Drops`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"rx_acl_rate_limit_drops": {
			Description: `Rx ACL Rate Limit Drops`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"rx_bytes": {
			Description: `Rx Bytes`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"rx_drops": {
			Description: `Rx Drops`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"rx_errors": {
			Description: `Rx Errors`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"rx_pkts": {
			Description: `Rx packets`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"tx_acl_drops": {
			Description: `Tx ACL Rate Drops`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"tx_acl_rate_limit_drops": {
			Description: `Tx ACL Rate Limit Drops`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"tx_bytes": {
			Description: `Tx Bytes`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"tx_drops": {
			Description: `Tx Drops`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"tx_errors": {
			Description: `Tx Errors`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"tx_pkts": {
			Description: `Tx Packets`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the NetworkCounters resource
func GetNetworkCountersPropertyFields() (t []string) {
	return []string{
		"if_name",
		"rx_acl_drops",
		"rx_acl_rate_limit_drops",
		"rx_bytes",
		"rx_drops",
		"rx_errors",
		"rx_pkts",
		"tx_acl_drops",
		"tx_acl_rate_limit_drops",
		"tx_bytes",
		"tx_drops",
		"tx_errors",
		"tx_pkts",
	}
}
