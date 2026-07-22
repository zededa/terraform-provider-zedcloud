package resources

import (
	"strings"
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// Covers the plan-time validation for ENG-2782: explicitly configuring
// netname/netid on a bond member interface must fail with an actionable
// error instead of the API's generic 400.
func TestCheckBondMemberInterfaces(t *testing.T) {
	iface := func(attrs map[string]cty.Value) cty.Value {
		base := map[string]cty.Value{
			"intfname":   cty.StringVal("eth0"),
			"intf_usage": cty.NullVal(cty.String),
			"netname":    cty.NullVal(cty.String),
			"netid":      cty.NullVal(cty.String),
		}
		for k, v := range attrs {
			base[k] = v
		}
		return cty.ObjectVal(base)
	}
	config := func(interfaces ...cty.Value) cty.Value {
		attrs := map[string]cty.Value{
			"name": cty.StringVal("test-node"),
		}
		if len(interfaces) == 0 {
			attrs["interfaces"] = cty.NullVal(cty.List(cty.EmptyObject))
		} else {
			attrs["interfaces"] = cty.TupleVal(interfaces)
		}
		return cty.ObjectVal(attrs)
	}
	bondMember := cty.StringVal(string(models.AdapterUsageADAPTERUSAGEBONDMEMBER))

	testCases := []struct {
		name        string
		rawConfig   cty.Value
		expectError string
	}{
		{
			name:      "null config",
			rawConfig: cty.NullVal(cty.EmptyObject),
		},
		{
			name:      "unknown config",
			rawConfig: cty.UnknownVal(cty.EmptyObject),
		},
		{
			name:      "config without interfaces attribute",
			rawConfig: cty.ObjectVal(map[string]cty.Value{"name": cty.StringVal("test-node")}),
		},
		{
			name:      "null interfaces list",
			rawConfig: config(),
		},
		{
			name: "bond member without network identity is valid",
			rawConfig: config(
				iface(map[string]cty.Value{
					"intf_usage": bondMember,
				}),
			),
		},
		{
			name: "management interface with network identity is valid",
			rawConfig: config(
				iface(map[string]cty.Value{
					"intf_usage": cty.StringVal(string(models.AdapterUsageADAPTERUSAGEMANAGEMENT)),
					"netname":    cty.StringVal("defaultIPv4-net"),
					"netid":      cty.StringVal("49972e18-a905-4740-938e-fd283b084257"),
				}),
			),
		},
		{
			name: "bond member with explicit netname is rejected",
			rawConfig: config(
				iface(map[string]cty.Value{
					"intf_usage": bondMember,
					"netname":    cty.StringVal("defaultIPv4-net"),
				}),
			),
			expectError: `interface "eth0"`,
		},
		{
			name: "bond member with explicit netid is rejected",
			rawConfig: config(
				iface(map[string]cty.Value{
					"intf_usage": bondMember,
					"netid":      cty.StringVal("49972e18-a905-4740-938e-fd283b084257"),
				}),
			),
			expectError: `interface "eth0"`,
		},
		{
			name: "mixed list flags only the offending bond member",
			rawConfig: config(
				iface(map[string]cty.Value{
					"intfname":   cty.StringVal("eth1"),
					"intf_usage": cty.StringVal(string(models.AdapterUsageADAPTERUSAGEMANAGEMENT)),
					"netname":    cty.StringVal("defaultIPv4-net"),
				}),
				iface(map[string]cty.Value{
					"intfname":   cty.StringVal("eth2"),
					"intf_usage": bondMember,
					"netname":    cty.StringVal("defaultIPv4-net"),
				}),
			),
			expectError: `interface "eth2"`,
		},
		{
			name: "bond member with unknown netid is not rejected at plan time",
			rawConfig: config(
				iface(map[string]cty.Value{
					"intf_usage": bondMember,
					"netid":      cty.UnknownVal(cty.String),
				}),
			),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := checkBondMemberInterfaces(tc.rawConfig)
			if tc.expectError == "" {
				if err != nil {
					t.Errorf("checkBondMemberInterfaces() = %v, expected no error", err)
				}
				return
			}
			if err == nil {
				t.Errorf("checkBondMemberInterfaces() = nil, expected error containing %q", tc.expectError)
				return
			}
			if !strings.Contains(err.Error(), tc.expectError) {
				t.Errorf("checkBondMemberInterfaces() = %v, expected error containing %q", err, tc.expectError)
			}
		})
	}
}
