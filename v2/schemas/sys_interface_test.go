package schemas

import (
	"testing"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func TestCompareSysInterfaceList(t *testing.T) {
	// Helper function to create pointer to AdapterUsage
	adapterUsagePtr := func(usage models.AdapterUsage) *models.AdapterUsage {
		return models.NewAdapterUsage(usage)
	}

	// Helper function to create pointer to NetworkDHCPType
	netDhcpPtr := func(dhcpType models.NetworkDHCPType) *models.NetworkDHCPType {
		return models.NewNetworkDHCPType(dhcpType)
	}

	testCases := []struct {
		name     string
		list1    []*models.SysInterface
		list2    []*models.SysInterface
		expected bool
	}{
		{
			name:     "both lists empty",
			list1:    []*models.SysInterface{},
			list2:    []*models.SysInterface{},
			expected: true,
		},
		{
			name:     "both lists nil",
			list1:    nil,
			list2:    nil,
			expected: true,
		},
		{
			name:     "one empty, one nil (should be equal)",
			list1:    []*models.SysInterface{},
			list2:    nil,
			expected: true,
		},
		{
			name: "different lengths",
			list1: []*models.SysInterface{
				{
					Intfname: "eth0",
					Cost:     0,
				},
			},
			list2:    []*models.SysInterface{},
			expected: false,
		},
		{
			name: "identical single item lists",
			list1: []*models.SysInterface{
				{
					Cost:         10,
					IntfUsage:    adapterUsagePtr(models.AdapterUsageADAPTERUSAGEMANAGEMENT),
					Intfname:     "eth0",
					Ipaddr:       "192.168.1.10",
					Macaddr:      "00:11:22:33:44:55",
					NetDhcp:      netDhcpPtr(models.NetworkDHCPTypeNETWORKDHCPTYPESTATIC),
					Netid:        "net-123",
					Netname:      "management",
					SharedLabels: []string{"label1", "label2"},
					Tags:         map[string]string{"env": "prod", "team": "infra"},
					Ztype:        "virtual",
				},
			},
			list2: []*models.SysInterface{
				{
					Cost:         10,
					IntfUsage:    adapterUsagePtr(models.AdapterUsageADAPTERUSAGEMANAGEMENT),
					Intfname:     "eth0",
					Ipaddr:       "192.168.1.10",
					Macaddr:      "00:11:22:33:44:55",
					NetDhcp:      netDhcpPtr(models.NetworkDHCPTypeNETWORKDHCPTYPESTATIC),
					Netid:        "net-123",
					Netname:      "management",
					SharedLabels: []string{"label1", "label2"},
					Tags:         map[string]string{"env": "prod", "team": "infra"},
					Ztype:        "virtual",
				},
			},
			expected: true,
		},
		{
			name: "identical multiple items in same order",
			list1: []*models.SysInterface{
				{
					Intfname:  "eth0",
					IntfUsage: adapterUsagePtr(models.AdapterUsageADAPTERUSAGEMANAGEMENT),
					Cost:      10,
				},
				{
					Intfname:  "eth1",
					IntfUsage: adapterUsagePtr(models.AdapterUsageADAPTERUSAGEAPPDIRECT),
					Cost:      20,
				},
			},
			list2: []*models.SysInterface{
				{
					Intfname:  "eth0",
					IntfUsage: adapterUsagePtr(models.AdapterUsageADAPTERUSAGEMANAGEMENT),
					Cost:      10,
				},
				{
					Intfname:  "eth1",
					IntfUsage: adapterUsagePtr(models.AdapterUsageADAPTERUSAGEAPPDIRECT),
					Cost:      20,
				},
			},
			expected: true,
		},
		{
			name: "identical multiple items in different order (should be sorted and equal)",
			list1: []*models.SysInterface{
				{
					Intfname:  "eth1",
					IntfUsage: adapterUsagePtr(models.AdapterUsageADAPTERUSAGEAPPDIRECT),
					Cost:      20,
				},
				{
					Intfname:  "eth0",
					IntfUsage: adapterUsagePtr(models.AdapterUsageADAPTERUSAGEMANAGEMENT),
					Cost:      10,
				},
			},
			list2: []*models.SysInterface{
				{
					Intfname:  "eth0",
					IntfUsage: adapterUsagePtr(models.AdapterUsageADAPTERUSAGEMANAGEMENT),
					Cost:      10,
				},
				{
					Intfname:  "eth1",
					IntfUsage: adapterUsagePtr(models.AdapterUsageADAPTERUSAGEAPPDIRECT),
					Cost:      20,
				},
			},
			expected: true,
		},
		{
			name: "different Cost values",
			list1: []*models.SysInterface{
				{
					Intfname: "eth0",
					Cost:     10,
				},
			},
			list2: []*models.SysInterface{
				{
					Intfname: "eth0",
					Cost:     20,
				},
			},
			expected: false,
		},
		{
			name: "different IntfUsage - both set",
			list1: []*models.SysInterface{
				{
					Intfname:  "eth0",
					IntfUsage: adapterUsagePtr(models.AdapterUsageADAPTERUSAGEMANAGEMENT),
				},
			},
			list2: []*models.SysInterface{
				{
					Intfname:  "eth0",
					IntfUsage: adapterUsagePtr(models.AdapterUsageADAPTERUSAGEAPPDIRECT),
				},
			},
			expected: false,
		},
		{
			name: "IntfUsage - one nil, one set",
			list1: []*models.SysInterface{
				{
					Intfname:  "eth0",
					IntfUsage: adapterUsagePtr(models.AdapterUsageADAPTERUSAGEMANAGEMENT),
				},
			},
			list2: []*models.SysInterface{
				{
					Intfname:  "eth0",
					IntfUsage: nil,
				},
			},
			expected: false,
		},
		{
			name: "IntfUsage - both nil",
			list1: []*models.SysInterface{
				{
					Intfname:  "eth0",
					IntfUsage: nil,
				},
			},
			list2: []*models.SysInterface{
				{
					Intfname:  "eth0",
					IntfUsage: nil,
				},
			},
			expected: true,
		},
		{
			name: "different Intfname",
			list1: []*models.SysInterface{
				{
					Intfname: "eth0",
				},
			},
			list2: []*models.SysInterface{
				{
					Intfname: "eth1",
				},
			},
			expected: false,
		},
		{
			name: "different Ipaddr",
			list1: []*models.SysInterface{
				{
					Intfname: "eth0",
					Ipaddr:   "192.168.1.10",
				},
			},
			list2: []*models.SysInterface{
				{
					Intfname: "eth0",
					Ipaddr:   "192.168.1.20",
				},
			},
			expected: false,
		},
		{
			name: "different Macaddr",
			list1: []*models.SysInterface{
				{
					Intfname: "eth0",
					Macaddr:  "00:11:22:33:44:55",
				},
			},
			list2: []*models.SysInterface{
				{
					Intfname: "eth0",
					Macaddr:  "00:11:22:33:44:66",
				},
			},
			expected: false,
		},
		{
			name: "different Netname",
			list1: []*models.SysInterface{
				{
					Intfname: "eth0",
					Netname:  "management",
				},
			},
			list2: []*models.SysInterface{
				{
					Intfname: "eth0",
					Netname:  "data",
				},
			},
			expected: false,
		},
		{
			name: "different NetDhcp - both set",
			list1: []*models.SysInterface{
				{
					Intfname: "eth0",
					NetDhcp:  netDhcpPtr(models.NetworkDHCPTypeNETWORKDHCPTYPESTATIC),
				},
			},
			list2: []*models.SysInterface{
				{
					Intfname: "eth0",
					NetDhcp:  netDhcpPtr(models.NetworkDHCPTypeNETWORKDHCPTYPECLIENT),
				},
			},
			expected: false,
		},
		{
			name: "NetDhcp - one nil, one set",
			list1: []*models.SysInterface{
				{
					Intfname: "eth0",
					NetDhcp:  netDhcpPtr(models.NetworkDHCPTypeNETWORKDHCPTYPESTATIC),
				},
			},
			list2: []*models.SysInterface{
				{
					Intfname: "eth0",
					NetDhcp:  nil,
				},
			},
			expected: false,
		},
		{
			name: "NetDhcp - both nil",
			list1: []*models.SysInterface{
				{
					Intfname: "eth0",
					NetDhcp:  nil,
				},
			},
			list2: []*models.SysInterface{
				{
					Intfname: "eth0",
					NetDhcp:  nil,
				},
			},
			expected: true,
		},
		{
			name: "different Ztype",
			list1: []*models.SysInterface{
				{
					Intfname: "eth0",
					Ztype:    "physical",
				},
			},
			list2: []*models.SysInterface{
				{
					Intfname: "eth0",
					Ztype:    "virtual",
				},
			},
			expected: false,
		},
		{
			name: "different Tags - different values",
			list1: []*models.SysInterface{
				{
					Intfname: "eth0",
					Tags:     map[string]string{"env": "prod"},
				},
			},
			list2: []*models.SysInterface{
				{
					Intfname: "eth0",
					Tags:     map[string]string{"env": "dev"},
				},
			},
			expected: false,
		},
		{
			name: "different Tags - different keys",
			list1: []*models.SysInterface{
				{
					Intfname: "eth0",
					Tags:     map[string]string{"env": "prod"},
				},
			},
			list2: []*models.SysInterface{
				{
					Intfname: "eth0",
					Tags:     map[string]string{"environment": "prod"},
				},
			},
			expected: false,
		},
		{
			name: "different Tags - one nil, one empty",
			list1: []*models.SysInterface{
				{
					Intfname: "eth0",
					Tags:     nil,
				},
			},
			list2: []*models.SysInterface{
				{
					Intfname: "eth0",
					Tags:     map[string]string{},
				},
			},
			expected: true, // reflect.DeepEqual treats nil and empty map as equal
		},
		{
			name: "same Tags - multiple entries",
			list1: []*models.SysInterface{
				{
					Intfname: "eth0",
					Tags:     map[string]string{"env": "prod", "team": "infra", "region": "us-west"},
				},
			},
			list2: []*models.SysInterface{
				{
					Intfname: "eth0",
					Tags:     map[string]string{"env": "prod", "team": "infra", "region": "us-west"},
				},
			},
			expected: true,
		},
		{
			name: "sorting by Intfname - case insensitive",
			list1: []*models.SysInterface{
				{
					Intfname:  "Eth1",
					IntfUsage: adapterUsagePtr(models.AdapterUsageADAPTERUSAGEMANAGEMENT),
				},
				{
					Intfname:  "eth0",
					IntfUsage: adapterUsagePtr(models.AdapterUsageADAPTERUSAGEMANAGEMENT),
				},
			},
			list2: []*models.SysInterface{
				{
					Intfname:  "eth0",
					IntfUsage: adapterUsagePtr(models.AdapterUsageADAPTERUSAGEMANAGEMENT),
				},
				{
					Intfname:  "Eth1",
					IntfUsage: adapterUsagePtr(models.AdapterUsageADAPTERUSAGEMANAGEMENT),
				},
			},
			expected: true,
		},
		{
			name: "sorting by IntfUsage when Intfname is same",
			list1: []*models.SysInterface{
				{
					Intfname:  "eth0",
					IntfUsage: adapterUsagePtr(models.AdapterUsageADAPTERUSAGEMANAGEMENT),
				},
				{
					Intfname:  "eth0",
					IntfUsage: adapterUsagePtr(models.AdapterUsageADAPTERUSAGEAPPDIRECT),
				},
			},
			list2: []*models.SysInterface{
				{
					Intfname:  "eth0",
					IntfUsage: adapterUsagePtr(models.AdapterUsageADAPTERUSAGEAPPDIRECT),
				},
				{
					Intfname:  "eth0",
					IntfUsage: adapterUsagePtr(models.AdapterUsageADAPTERUSAGEMANAGEMENT),
				},
			},
			expected: true,
		},
		{
			name: "complex scenario with multiple interfaces",
			list1: []*models.SysInterface{
				{
					Cost:      5,
					IntfUsage: adapterUsagePtr(models.AdapterUsageADAPTERUSAGEDISABLED),
					Intfname:  "eth2",
					Netname:   "storage",
					Ztype:     "virtual",
					Tags:      map[string]string{"role": "storage"},
				},
				{
					Cost:      10,
					IntfUsage: adapterUsagePtr(models.AdapterUsageADAPTERUSAGEMANAGEMENT),
					Intfname:  "eth0",
					Ipaddr:    "192.168.1.10",
					Macaddr:   "00:11:22:33:44:55",
					NetDhcp:   netDhcpPtr(models.NetworkDHCPTypeNETWORKDHCPTYPESTATIC),
					Netid:     "net-123",
					Netname:   "management",
					Ztype:     "physical",
					Tags:      map[string]string{"env": "prod", "team": "infra"},
				},
				{
					Cost:      20,
					IntfUsage: adapterUsagePtr(models.AdapterUsageADAPTERUSAGEAPPDIRECT),
					Intfname:  "eth1",
					Ipaddr:    "10.0.0.5",
					NetDhcp:   netDhcpPtr(models.NetworkDHCPTypeNETWORKDHCPTYPECLIENT),
					Netid:     "net-456",
					Netname:   "data",
					Ztype:     "virtual",
					Tags:      map[string]string{"env": "prod"},
				},
			},
			list2: []*models.SysInterface{
				{
					Cost:      10,
					IntfUsage: adapterUsagePtr(models.AdapterUsageADAPTERUSAGEMANAGEMENT),
					Intfname:  "eth0",
					Ipaddr:    "192.168.1.10",
					Macaddr:   "00:11:22:33:44:55",
					NetDhcp:   netDhcpPtr(models.NetworkDHCPTypeNETWORKDHCPTYPESTATIC),
					Netid:     "net-123",
					Netname:   "management",
					Ztype:     "physical",
					Tags:      map[string]string{"env": "prod", "team": "infra"},
				},
				{
					Cost:      20,
					IntfUsage: adapterUsagePtr(models.AdapterUsageADAPTERUSAGEAPPDIRECT),
					Intfname:  "eth1",
					Ipaddr:    "10.0.0.5",
					NetDhcp:   netDhcpPtr(models.NetworkDHCPTypeNETWORKDHCPTYPECLIENT),
					Netid:     "net-456",
					Netname:   "data",
					Ztype:     "virtual",
					Tags:      map[string]string{"env": "prod"},
				},
				{
					Cost:      5,
					IntfUsage: adapterUsagePtr(models.AdapterUsageADAPTERUSAGEDISABLED),
					Intfname:  "eth2",
					Netname:   "storage",
					Ztype:     "virtual",
					Tags:      map[string]string{"role": "storage"},
				},
			},
			expected: true,
		},
		{
			name: "SharedLabels field is equal",
			list1: []*models.SysInterface{
				{
					Intfname:     "eth0",
					SharedLabels: []string{"label1", "label2"},
				},
			},
			list2: []*models.SysInterface{
				{
					Intfname:     "eth0",
					SharedLabels: []string{"label1", "label2"},
				},
			},
			expected: true,
		},
		{
			name: "SharedLabels field is not equal",
			list1: []*models.SysInterface{
				{
					Intfname:     "eth0",
					SharedLabels: []string{"label1", "label2"},
				},
			},
			list2: []*models.SysInterface{
				{
					Intfname:     "eth0",
					SharedLabels: []string{"label3", "label4"},
				},
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, _ := CompareSystemInterfaceList(tc.list1, tc.list2)
			if result != tc.expected {
				t.Errorf("CompareSysInterfaceList() = %v, expected %v", result, tc.expected)
			}
		})
	}
}
