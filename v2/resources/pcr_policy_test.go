package resources

import (
	"reflect"
	"testing"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func pcr(version int64, indices ...int64) *models.Policy {
	return &models.Policy{
		Type:      models.NewPolicyType(models.PolicyTypePOLICYTYPEPCR),
		PcrPolicy: &models.PCRPolicy{PcrIndices: indices, PolicyVersion: version},
	}
}

func TestNodePCRPolicyForUpdate(t *testing.T) {
	cases := []struct {
		name        string
		configured  *models.Policy
		wasManaged  bool
		existing    *models.Policy
		wantNil     bool
		wantVersion int64
		wantIndices []int64
	}{
		{
			name:        "new policy from config",
			configured:  pcr(0, 0, 1, 2),
			wantVersion: 1, wantIndices: []int64{0, 1, 2},
		},
		{
			name:        "indices changed: bump from the live version",
			configured:  pcr(0, 0, 1),
			wasManaged:  true,
			existing:    pcr(4, 0, 1, 2),
			wantVersion: 5, wantIndices: []int64{0, 1},
		},
		{
			name:        "indices unchanged: keep version",
			configured:  pcr(0, 2, 1, 0),
			wasManaged:  true,
			existing:    pcr(4, 0, 1, 2),
			wantVersion: 4, wantIndices: []int64{2, 1, 0},
		},
		{
			name:        "zcli changed it in between: continue from its version",
			configured:  pcr(0, 0, 7),
			wasManaged:  true,
			existing:    pcr(9, 0, 1, 7),
			wantVersion: 10, wantIndices: []int64{0, 7},
		},
		{
			name:       "removed from config: reset",
			wasManaged: true,
			existing:   pcr(4, 0, 1, 2),
			wantNil:    true,
		},
		{
			name:        "never managed: keep the zcli-set policy",
			existing:    pcr(3, 0, 7),
			wantVersion: 3, wantIndices: []int64{0, 7},
		},
		{
			name:    "never managed, none on the node",
			wantNil: true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := nodePCRPolicyForUpdate(tc.configured, tc.wasManaged, tc.existing)
			if tc.wantNil {
				if got != nil {
					t.Fatalf("expected no policy, got %+v", got.PcrPolicy)
				}
				return
			}
			if got == nil || got.PcrPolicy == nil {
				t.Fatal("expected a policy")
			}
			if got.PcrPolicy.PolicyVersion != tc.wantVersion {
				t.Errorf("version = %d, want %d", got.PcrPolicy.PolicyVersion, tc.wantVersion)
			}
			if !reflect.DeepEqual(got.PcrPolicy.PcrIndices, tc.wantIndices) {
				t.Errorf("indices = %v, want %v", got.PcrPolicy.PcrIndices, tc.wantIndices)
			}
		})
	}
}

func TestDeploymentHasPolicies(t *testing.T) {
	if deploymentHasPolicies(&models.Deployment{}) {
		t.Fatal("empty deployment has no policies")
	}
	if !deploymentHasPolicies(&models.Deployment{DevicePolicies: []*models.DevicePolicy{{}}}) {
		t.Fatal("device policy not counted")
	}
	if !deploymentHasPolicies(&models.Deployment{EdgeviewPolicy: &models.EdgeViewPolicy{}}) {
		t.Fatal("edgeview policy not counted")
	}
}
