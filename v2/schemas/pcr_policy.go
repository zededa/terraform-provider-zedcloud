package schemas

import (
	"fmt"
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// PCR policy (NFR-189 / NFR-249): which TPM PCRs EVE uses to seal the vault key.
//
// The same rules are enforced by zedcloud (seine validatePCRIndices, the gilas
// deployment pre-check and thames) and by zcli. They are checked here at plan
// time as well, so a bad set fails before anything is sent.
const (
	PCRMaxIndex      = 15
	PCRVolatileIndex = 5
	PCRRequiredIndex = 0
)

// PCRDefaultIndices is the recommended set (same as zcli's default).
var PCRDefaultIndices = []int{0, 1, 2, 3, 4, 6, 7, 8, 9, 13, 14}

const pcrIndicesDescription = `PCR indices (0-15) used to seal the TPM volume key. PCR 0 is required; PCR 5 is volatile and not allowed. ` +
	`Recommended set: [0, 1, 2, 3, 4, 6, 7, 8, 9, 13, 14]. ` +
	`Warning: a set that includes PCRs which change on firmware or bootloader updates can leave the node unable to unseal its vault after such an update.`

const pcrPolicyVersionDescription = `Policy version, sent to EVE as the policy id. Managed by the provider: ` +
	`set to 1 when the policy is created and incremented when pcr_indices change.`

// ValidatePCRIndices checks a PCR index list against the controller's rules:
// non-empty, 0-15, no PCR 5, no duplicates, PCR 0 present.
func ValidatePCRIndices(indices []int) error {
	if len(indices) == 0 {
		return fmt.Errorf("PCR indices list cannot be empty")
	}
	seen := make(map[int]bool, len(indices))
	for _, idx := range indices {
		if idx < 0 || idx > PCRMaxIndex {
			return fmt.Errorf("invalid PCR index %d: must be between 0 and %d", idx, PCRMaxIndex)
		}
		if idx == PCRVolatileIndex {
			return fmt.Errorf("PCR %d is volatile and must not be included in the policy", PCRVolatileIndex)
		}
		if seen[idx] {
			return fmt.Errorf("duplicate PCR index %d", idx)
		}
		seen[idx] = true
	}
	if !seen[PCRRequiredIndex] {
		return fmt.Errorf("PCR %d (Core Root of Trust Measurement) must be included in the policy", PCRRequiredIndex)
	}
	return nil
}

func validatePCRIndexElem(v interface{}, k string) (ws []string, es []error) {
	idx, ok := v.(int)
	if !ok {
		return nil, []error{fmt.Errorf("%s: expected an integer", k)}
	}
	if idx < 0 || idx > PCRMaxIndex {
		es = append(es, fmt.Errorf("%s: invalid PCR index %d, must be between 0 and %d", k, idx, PCRMaxIndex))
	}
	if idx == PCRVolatileIndex {
		es = append(es, fmt.Errorf("%s: PCR %d is volatile and must not be included in the policy", k, PCRVolatileIndex))
	}
	return ws, es
}

func pcrIndicesSchema() *schema.Schema {
	return &schema.Schema{
		Description: pcrIndicesDescription,
		Type:        schema.TypeSet,
		Elem: &schema.Schema{
			Type:         schema.TypeInt,
			ValidateFunc: validatePCRIndexElem,
		},
		Required: true,
		MinItems: 1,
	}
}

// PCRIndicesFromInterface converts a TypeSet/TypeList of ints into a sorted
// []int64, the order zcli sends and the order the API returns.
func PCRIndicesFromInterface(v interface{}) []int64 {
	var items []interface{}
	switch t := v.(type) {
	case *schema.Set:
		items = t.List()
	case []interface{}:
		items = t
	}
	out := make([]int64, 0, len(items))
	for _, it := range items {
		if i, ok := it.(int); ok {
			out = append(out, int64(i))
		}
	}
	sort.Slice(out, func(a, b int) bool { return out[a] < out[b] })
	return out
}

func pcrIndicesToInterface(indices []int64) []interface{} {
	sorted := append([]int64(nil), indices...)
	sort.Slice(sorted, func(a, b int) bool { return sorted[a] < sorted[b] })
	out := make([]interface{}, 0, len(sorted))
	for _, i := range sorted {
		out = append(out, int(i))
	}
	return out
}

// PCRIndicesEqual reports whether two index lists hold the same set.
func PCRIndicesEqual(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	sa := append([]int64(nil), a...)
	sb := append([]int64(nil), b...)
	sort.Slice(sa, func(i, j int) bool { return sa[i] < sa[j] })
	sort.Slice(sb, func(i, j int) bool { return sb[i] < sb[j] })
	for i := range sa {
		if sa[i] != sb[i] {
			return false
		}
	}
	return true
}

// NextPCRPolicyVersion returns the version to send: 1 for a new policy,
// current+1 when the indices changed, current otherwise. EVE treats a change of
// either the indices or the version as a new policy, so the version must not
// change when the indices don't.
func NextPCRPolicyVersion(currentIndices []int64, currentVersion int64, newIndices []int64) int64 {
	if len(currentIndices) == 0 {
		if currentVersion > 0 {
			return currentVersion + 1
		}
		return 1
	}
	if PCRIndicesEqual(currentIndices, newIndices) {
		if currentVersion == 0 {
			return 1
		}
		return currentVersion
	}
	return currentVersion + 1
}

// ---- PCRPolicy: the body of a project-level policy (Policy.pcr_policy) ----

func PCRPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Description: `unique policy id`,
			Type:        schema.TypeString,
			Computed:    true,
		},
		"pcr_indices": pcrIndicesSchema(),
		"policy_version": {
			Description: pcrPolicyVersionDescription,
			Type:        schema.TypeInt,
			Computed:    true,
		},
	}
}

func PCRPolicyModelFromMap(m map[string]interface{}) *models.PCRPolicy {
	id, _ := m["id"].(string)
	version, _ := m["policy_version"].(int)
	return &models.PCRPolicy{
		ID:            id,
		PcrIndices:    PCRIndicesFromInterface(m["pcr_indices"]),
		PolicyVersion: int64(version),
	}
}

func SetPCRPolicySubResourceData(m []*models.PCRPolicy) (d []*map[string]interface{}) {
	for _, p := range m {
		if p == nil {
			continue
		}
		properties := map[string]interface{}{
			"id":             p.ID,
			"pcr_indices":    pcrIndicesToInterface(p.PcrIndices),
			"policy_version": int(p.PolicyVersion),
		}
		d = append(d, &properties)
	}
	return
}

// ---- DevicePCRPolicy: deployment device policy (zedcloud_deployment) ----

func DevicePCRPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"pcr_indices": pcrIndicesSchema(),
		"policy_version": {
			Description: `Policy version. Managed by the controller for deployment policies (1, then +1 on each update).`,
			Type:        schema.TypeInt,
			Computed:    true,
		},
	}
}

func DevicePCRPolicyModelFromMap(m map[string]interface{}) *models.DevicePCRPolicy {
	version, _ := m["policy_version"].(int)
	return &models.DevicePCRPolicy{
		PcrIndices:    PCRIndicesFromInterface(m["pcr_indices"]),
		PolicyVersion: int64(version),
	}
}

func SetDevicePCRPolicySubResourceData(m []*models.DevicePCRPolicy) (d []*map[string]interface{}) {
	for _, p := range m {
		if p == nil {
			continue
		}
		properties := map[string]interface{}{
			"pcr_indices":    pcrIndicesToInterface(p.PcrIndices),
			"policy_version": int(p.PolicyVersion),
		}
		d = append(d, &properties)
	}
	return
}

// ---- Edge node: a small block wrapped into Policy{type: POLICY_TYPE_PCR} ----

// NodePCRPolicySchema is zedcloud_edgenode.pcr_policy. The device API takes a
// full PolicyConfig, but only its type and pcrPolicy are used, so the
// provider exposes just the indices and wraps them on the wire.
func NodePCRPolicySchema() *schema.Schema {
	return &schema.Schema{
		Description: `Device-level PCR policy for TPM volume key sealing. Overrides the project's PCR policy for this node ` +
			`(for nodes in a deployment project, a matching deployment PCR device policy takes precedence over it). ` +
			`Managed only once set: a policy set outside Terraform (e.g. with zcli) is kept while this block is absent; ` +
			`removing the block after it was set resets the node to the project or EVE defaults.`,
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"pcr_indices": pcrIndicesSchema(),
				"policy_version": {
					Description: pcrPolicyVersionDescription,
					Type:        schema.TypeInt,
					Computed:    true,
				},
			},
		},
	}
}

// NodePCRPolicyModel builds the wire policy from the node's pcr_policy block,
// or returns nil when the block is absent. The version is left at the value
// in the block; the resource sets the final version from the live node.
func NodePCRPolicyModel(v interface{}) *models.Policy {
	l, _ := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}
	m := l[0].(map[string]interface{})
	version, _ := m["policy_version"].(int)
	return &models.Policy{
		Type: models.NewPolicyType(models.PolicyTypePOLICYTYPEPCR),
		PcrPolicy: &models.PCRPolicy{
			PcrIndices:    PCRIndicesFromInterface(m["pcr_indices"]),
			PolicyVersion: int64(version),
		},
	}
}

// FlattenNodePCRPolicy converts the node's wire policy into the block value.
func FlattenNodePCRPolicy(p *models.Policy) []interface{} {
	if p == nil || p.PcrPolicy == nil || len(p.PcrPolicy.PcrIndices) == 0 {
		return nil
	}
	return []interface{}{map[string]interface{}{
		"pcr_indices":    pcrIndicesToInterface(p.PcrPolicy.PcrIndices),
		"policy_version": int(p.PcrPolicy.PolicyVersion),
	}}
}
