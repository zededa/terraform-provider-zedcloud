package resources

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Regression tests for ENG-2805: removing the last bond_adapter (or
// vlan_adapters) block from the config must produce a plan diff. GetChange's
// "new" side merges config over state, and the hcl2 shim omits empty block
// lists from the legacy config entirely, so the merge falls back to state and
// the list diff-suppress functions used to see old==new and suppress the
// removal. The suppress functions now consult the raw config (which always
// carries the attribute, mirroring what PlanResourceChange provides via
// InstanceState.RawConfig) to detect the removed blocks.

// nodeStateWithBondAdapter returns an InstanceState for a node with one bond
// adapter (lag-test-01 over eth0/eth2) and its two bond member interfaces.
func nodeStateWithBondAdapter() *terraform.InstanceState {
	return &terraform.InstanceState{
		ID: "609fe5d4-d6a6-4063-a8a5-fbd5cfdcb8d6",
		Attributes: map[string]string{
			"id":    "609fe5d4-d6a6-4063-a8a5-fbd5cfdcb8d6",
			"name":  "lag-device",
			"title": "lag-device",

			"interfaces.#":            "2",
			"interfaces.0.intfname":   "eth0",
			"interfaces.0.intf_usage": "ADAPTER_USAGE_BOND_MEMBER",
			"interfaces.0.net_dhcp":   "NETWORK_DHCP_TYPE_UNSPECIFIED",
			"interfaces.0.cost":       "0",
			"interfaces.0.tags.%":     "0",
			"interfaces.1.intfname":   "eth2",
			"interfaces.1.intf_usage": "ADAPTER_USAGE_BOND_MEMBER",
			"interfaces.1.net_dhcp":   "NETWORK_DHCP_TYPE_UNSPECIFIED",
			"interfaces.1.cost":       "0",
			"interfaces.1.tags.%":     "0",

			"bond_adapter.#":                        "1",
			"bond_adapter.0.logical_label":          "lag-test-01",
			"bond_adapter.0.bond_mode":              "BOND_MODE_802_3AD",
			"bond_adapter.0.lacp_rate":              "LACP_RATE_FAST",
			"bond_adapter.0.lower_layer_names.#":    "2",
			"bond_adapter.0.lower_layer_names.0":    "eth0",
			"bond_adapter.0.lower_layer_names.1":    "eth2",
			"bond_adapter.0.interface.#":            "1",
			"bond_adapter.0.interface.0.intfname":   "lag-test-01",
			"bond_adapter.0.interface.0.intf_usage": "ADAPTER_USAGE_MANAGEMENT",
			"bond_adapter.0.interface.0.netname":    "defaultIPv4-net",
			"bond_adapter.0.interface.0.netid":      "49972e18-a905-4740-938e-fd283b084257",
			"bond_adapter.0.interface.0.net_dhcp":   "NETWORK_DHCP_TYPE_CLIENT",
			"bond_adapter.0.interface.0.cost":       "0",
			"bond_adapter.0.interface.0.tags.%":     "0",
		},
	}
}

// diffKeysWithPrefix returns the diff attribute keys starting with prefix.
func diffKeysWithPrefix(diff *terraform.InstanceDiff, prefix string) []string {
	var keys []string
	if diff == nil {
		return keys
	}
	for k := range diff.Attributes {
		if strings.HasPrefix(k, prefix) {
			keys = append(keys, k)
		}
	}
	return keys
}

func TestNodeResource_BondAdapterRemoval_PlansDiff(t *testing.T) {
	r := NodeResource()

	state := nodeStateWithBondAdapter()
	// PlanResourceChange always attaches the raw config to the prior state
	// before diffing; a removed block list is present there as an empty list.
	state.RawConfig = cty.ObjectVal(map[string]cty.Value{
		"bond_adapter": cty.ListValEmpty(cty.EmptyObject),
	})

	// config: bond_adapter removed entirely, members flipped to MANAGEMENT
	config := terraform.NewResourceConfigRaw(map[string]interface{}{
		"name":  "lag-device",
		"title": "lag-device",
		"interfaces": []interface{}{
			map[string]interface{}{
				"intfname":   "eth0",
				"intf_usage": "ADAPTER_USAGE_MANAGEMENT",
				"netname":    "defaultIPv4-net",
				"netid":      "49972e18-a905-4740-938e-fd283b084257",
				"net_dhcp":   "NETWORK_DHCP_TYPE_CLIENT",
			},
			map[string]interface{}{
				"intfname":   "eth2",
				"intf_usage": "ADAPTER_USAGE_MANAGEMENT",
				"netname":    "defaultIPv4-net",
				"netid":      "49972e18-a905-4740-938e-fd283b084257",
				"net_dhcp":   "NETWORK_DHCP_TYPE_CLIENT",
			},
		},
	})

	diff, err := r.Diff(context.Background(), state, config, nil)
	if err != nil {
		t.Fatalf("diff error: %s", err)
	}
	if diff == nil {
		t.Fatal("diff is nil, expected the bond_adapter removal to be planned")
	}

	countDiff, ok := diff.Attributes["bond_adapter.#"]
	if !ok {
		t.Fatalf(
			"bond_adapter removal missing from diff, bond_adapter keys: %v",
			diffKeysWithPrefix(diff, "bond_adapter"),
		)
	}
	if countDiff.Old != "1" || countDiff.New != "0" {
		t.Errorf("bond_adapter.# diff = %q => %q, expected \"1\" => \"0\"", countDiff.Old, countDiff.New)
	}
}

func TestNodeResource_BondAdapterUnchanged_SuppressesDiff(t *testing.T) {
	r := NodeResource()

	bondConfig := map[string]interface{}{
		"logical_label":     "lag-test-01",
		"bond_mode":         "BOND_MODE_802_3AD",
		"lacp_rate":         "LACP_RATE_FAST",
		"lower_layer_names": []interface{}{"eth0", "eth2"},
		"interface": []interface{}{
			map[string]interface{}{
				"intfname":   "lag-test-01",
				"intf_usage": "ADAPTER_USAGE_MANAGEMENT",
				"netname":    "defaultIPv4-net",
				"netid":      "49972e18-a905-4740-938e-fd283b084257",
				"net_dhcp":   "NETWORK_DHCP_TYPE_CLIENT",
			},
		},
	}

	state := nodeStateWithBondAdapter()
	state.RawConfig = cty.ObjectVal(map[string]cty.Value{
		"bond_adapter": cty.ListVal([]cty.Value{
			cty.ObjectVal(map[string]cty.Value{
				"logical_label": cty.StringVal("lag-test-01"),
			}),
		}),
	})

	config := terraform.NewResourceConfigRaw(map[string]interface{}{
		"name":  "lag-device",
		"title": "lag-device",
		"interfaces": []interface{}{
			map[string]interface{}{
				"intfname":   "eth0",
				"intf_usage": "ADAPTER_USAGE_BOND_MEMBER",
			},
			map[string]interface{}{
				"intfname":   "eth2",
				"intf_usage": "ADAPTER_USAGE_BOND_MEMBER",
			},
		},
		"bond_adapter": []interface{}{bondConfig},
	})

	diff, err := r.Diff(context.Background(), state, config, nil)
	if err != nil {
		t.Fatalf("diff error: %s", err)
	}

	if keys := diffKeysWithPrefix(diff, "bond_adapter"); len(keys) > 0 {
		t.Errorf("unchanged bond_adapter produced diff keys: %v", keys)
	}
}

// TestNodeResource_BondAdapterAddition_PlansDiff pins the ENG-2782 forward
// scenario (convert interfaces to BOND_MEMBER and add a bond_adapter block)
// at the plan level, proving the ENG-2805 raw-config guard does not interfere
// with it: the guard is inert on addition because the config still carries the
// bond_adapter key, so GetChange never falls back to state.
func TestNodeResource_BondAdapterAddition_PlansDiff(t *testing.T) {
	r := NodeResource()

	// state: no bond adapter, members still MANAGEMENT with netid/netname
	state := &terraform.InstanceState{
		ID: "609fe5d4-d6a6-4063-a8a5-fbd5cfdcb8d6",
		Attributes: map[string]string{
			"id":    "609fe5d4-d6a6-4063-a8a5-fbd5cfdcb8d6",
			"name":  "lag-device",
			"title": "lag-device",

			"interfaces.#":            "2",
			"interfaces.0.intfname":   "eth0",
			"interfaces.0.intf_usage": "ADAPTER_USAGE_MANAGEMENT",
			"interfaces.0.netname":    "defaultIPv4-net",
			"interfaces.0.netid":      "49972e18-a905-4740-938e-fd283b084257",
			"interfaces.0.net_dhcp":   "NETWORK_DHCP_TYPE_CLIENT",
			"interfaces.0.cost":       "0",
			"interfaces.0.tags.%":     "0",
			"interfaces.1.intfname":   "eth2",
			"interfaces.1.intf_usage": "ADAPTER_USAGE_MANAGEMENT",
			"interfaces.1.netname":    "defaultIPv4-net",
			"interfaces.1.netid":      "49972e18-a905-4740-938e-fd283b084257",
			"interfaces.1.net_dhcp":   "NETWORK_DHCP_TYPE_CLIENT",
			"interfaces.1.cost":       "0",
			"interfaces.1.tags.%":     "0",

			"bond_adapter.#": "0",
		},
	}
	// raw config declares one bond_adapter block
	state.RawConfig = cty.ObjectVal(map[string]cty.Value{
		"bond_adapter": cty.ListVal([]cty.Value{
			cty.ObjectVal(map[string]cty.Value{
				"logical_label": cty.StringVal("lag-test-01"),
			}),
		}),
	})

	config := terraform.NewResourceConfigRaw(map[string]interface{}{
		"name":  "lag-device",
		"title": "lag-device",
		"interfaces": []interface{}{
			map[string]interface{}{
				"intfname":   "eth0",
				"intf_usage": "ADAPTER_USAGE_BOND_MEMBER",
			},
			map[string]interface{}{
				"intfname":   "eth2",
				"intf_usage": "ADAPTER_USAGE_BOND_MEMBER",
			},
		},
		"bond_adapter": []interface{}{
			map[string]interface{}{
				"logical_label":     "lag-test-01",
				"bond_mode":         "BOND_MODE_802_3AD",
				"lacp_rate":         "LACP_RATE_FAST",
				"lower_layer_names": []interface{}{"eth0", "eth2"},
				"interface": []interface{}{
					map[string]interface{}{
						"intfname":   "lag-test-01",
						"intf_usage": "ADAPTER_USAGE_MANAGEMENT",
						"netname":    "defaultIPv4-net",
						"netid":      "49972e18-a905-4740-938e-fd283b084257",
						"net_dhcp":   "NETWORK_DHCP_TYPE_CLIENT",
					},
				},
			},
		},
	})

	diff, err := r.Diff(context.Background(), state, config, nil)
	if err != nil {
		t.Fatalf("diff error: %s", err)
	}
	if diff == nil {
		t.Fatal("diff is nil, expected bond_adapter addition and interface conversion to be planned")
	}

	countDiff, ok := diff.Attributes["bond_adapter.#"]
	if !ok {
		t.Fatal("bond_adapter.# missing from diff, addition not planned")
	}
	if countDiff.Old != "0" || countDiff.New != "1" {
		t.Errorf("bond_adapter.# diff = %q => %q, expected \"0\" => \"1\"", countDiff.Old, countDiff.New)
	}

	usageDiff, ok := diff.Attributes["interfaces.0.intf_usage"]
	if !ok {
		t.Fatal("interfaces.0.intf_usage missing from diff, conversion not planned")
	}
	if usageDiff.New != "ADAPTER_USAGE_BOND_MEMBER" {
		t.Errorf("interfaces.0.intf_usage new = %q, expected ADAPTER_USAGE_BOND_MEMBER", usageDiff.New)
	}
}

func TestNodeResource_VlanAdapterRemoval_PlansDiff(t *testing.T) {
	r := NodeResource()

	state := &terraform.InstanceState{
		ID: "609fe5d4-d6a6-4063-a8a5-fbd5cfdcb8d6",
		Attributes: map[string]string{
			"id":    "609fe5d4-d6a6-4063-a8a5-fbd5cfdcb8d6",
			"name":  "vlan-device",
			"title": "vlan-device",

			"interfaces.#":            "1",
			"interfaces.0.intfname":   "eth0",
			"interfaces.0.intf_usage": "ADAPTER_USAGE_VLANS_ONLY",
			"interfaces.0.net_dhcp":   "NETWORK_DHCP_TYPE_UNSPECIFIED",
			"interfaces.0.cost":       "0",
			"interfaces.0.tags.%":     "0",

			"vlan_adapters.#":                  "1",
			"vlan_adapters.0.logical_label":    "vlan-100",
			"vlan_adapters.0.interface_name":   "vlan-100",
			"vlan_adapters.0.lower_layer_name": "eth0",
			"vlan_adapters.0.vlan_id":          "100",
		},
	}
	state.RawConfig = cty.ObjectVal(map[string]cty.Value{
		"vlan_adapters": cty.ListValEmpty(cty.EmptyObject),
	})

	config := terraform.NewResourceConfigRaw(map[string]interface{}{
		"name":  "vlan-device",
		"title": "vlan-device",
		"interfaces": []interface{}{
			map[string]interface{}{
				"intfname":   "eth0",
				"intf_usage": "ADAPTER_USAGE_MANAGEMENT",
				"netname":    "defaultIPv4-net",
				"net_dhcp":   "NETWORK_DHCP_TYPE_CLIENT",
			},
		},
	})

	diff, err := r.Diff(context.Background(), state, config, nil)
	if err != nil {
		t.Fatalf("diff error: %s", err)
	}
	if diff == nil {
		t.Fatal("diff is nil, expected the vlan_adapters removal to be planned")
	}

	countDiff, ok := diff.Attributes["vlan_adapters.#"]
	if !ok {
		t.Fatalf(
			"vlan_adapters removal missing from diff, vlan_adapters keys: %v",
			diffKeysWithPrefix(diff, "vlan_adapters"),
		)
	}
	if countDiff.Old != "1" || countDiff.New != "0" {
		t.Errorf("vlan_adapters.# diff = %q => %q, expected \"1\" => \"0\"", countDiff.Old, countDiff.New)
	}
}
