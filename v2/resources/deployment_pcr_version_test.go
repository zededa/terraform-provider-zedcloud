package resources

import (
	"testing"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// ENG-3116: the controller stores the policyVersion it is sent for a
// deployment PCR device policy (it does not manage it), and a policy with
// version 0 never reaches the nodes.

func pcrDeployment(indices ...int64) *models.Deployment {
	return &models.Deployment{DevicePolicies: []*models.DevicePolicy{{
		PolicySubType: models.NewDevicePolicyType(models.DevicePolicyTypeDEVICEPOLICYTYPEPCR),
		PcrPolicy:     &models.DevicePCRPolicy{PcrIndices: indices},
	}}}
}

func oldPCRState(version int, indices ...interface{}) []interface{} {
	return []interface{}{map[string]interface{}{
		"policy_sub_type": string(models.DevicePolicyTypeDEVICEPOLICYTYPEPCR),
		"pcr_policy": []interface{}{map[string]interface{}{
			"pcr_indices":    indices,
			"policy_version": version,
		}},
	}}
}

func TestDeploymentPCRVersion(t *testing.T) {
	cases := []struct {
		name string
		new  *models.Deployment
		old  []interface{}
		want int64
	}{
		{"create", pcrDeployment(0, 1, 2, 3, 4, 7), nil, 1},
		{"unchanged indices", pcrDeployment(0, 1, 2, 3, 4, 7), oldPCRState(3, 0, 1, 2, 3, 4, 7), 3},
		{"unchanged indices, other order", pcrDeployment(7, 4, 3, 2, 1, 0), oldPCRState(3, 0, 1, 2, 3, 4, 7), 3},
		{"changed indices", pcrDeployment(0, 1, 2, 3), oldPCRState(3, 0, 1, 2, 3, 4, 7), 4},
		{"state has version 0 (created before the fix)", pcrDeployment(0, 1, 2, 3, 4, 7), oldPCRState(0, 0, 1, 2, 3, 4, 7), 1},
		{"policy added to a deployment without one", pcrDeployment(0, 1), []interface{}{}, 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			setDeploymentPCRPolicyVersions(c.new, c.old)
			if got := c.new.DevicePolicies[0].PcrPolicy.PolicyVersion; got != c.want {
				t.Errorf("policyVersion = %d, want %d", got, c.want)
			}
		})
	}
}

func TestDeploymentPCRVersion_IgnoresOtherPolicies(t *testing.T) {
	m := &models.Deployment{DevicePolicies: []*models.DevicePolicy{
		{PolicySubType: models.NewDevicePolicyType(models.DevicePolicyTypeDEVICEPOLICYTYPEPCR)}, // no pcr_policy block
		nil,
	}}
	setDeploymentPCRPolicyVersions(m, nil) // must not panic
}
