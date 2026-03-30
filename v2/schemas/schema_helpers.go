package schemas

import (
	"fmt"
	"log"
	"reflect"
	"slices"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	"github.com/zededa/terraform-provider-zedcloud/v2/pkg/compare"
)

// diffSuppressStringListOrder suppresses diffs in schema.List element order. If the lists elements differ, it still
// reports a diff.
// FIXME: Note, this is slow since it is called for every element of the schema.List. Thus, it's an expensive workaround
// that should be removed if possible.
func diffSuppressStringListOrder(key, oldValue, newValue string, d *schema.ResourceData) bool {
	// For a list, the key is path to the element, rather than the list.
	// E.g. "node_groups.2.ips.0"
	lastDotIndex := strings.LastIndex(key, ".")
	if lastDotIndex != -1 {
		key = string(key[:lastDotIndex])
	}

	oldData, newData := d.GetChange(key)
	if oldData == nil || newData == nil {
		return false
	}

	var newList []string
	for _, s := range newData.([]any) {
		newList = append(newList, s.(string))
	}
	var oldList []string
	for _, s := range oldData.([]any) {
		oldList = append(oldList, s.(string))
	}

	slices.Sort(oldList)
	slices.Sort(newList)

	return slices.Equal(oldList, newList)
}

// Sort and compare DNSList.
// FIXME: The API does not guarantee the order of the list items and terraform reports a diff if the order in state and
// config differs. Note, this is slow since it is called for every element of the schema.List. Thus, it's an expensive
// workaround that should be removed if possible.
func diffSuppressDNSListOrder(mapKey string) schema.SchemaDiffSuppressFunc {
	return func(key, oldValue, newValue string, d *schema.ResourceData) bool {
		oldData, newData := d.GetChange(mapKey)
		if newData == nil && oldData == nil {
			return true
		}

		if oldData == nil {
			return false
		}
		if newData == nil {
			return false
		}

		o := oldData.([]interface{})
		n := newData.([]interface{})
		if len(o) == 0 && len(n) == 0 {
			return true
		}
		if len(o) != len(n) {
			return false
		}

		var oldMapList []*models.StaticDNSList
		for _, m := range o {
			if m == nil {
				continue
			}
			oldMapList = append(oldMapList, StaticDNSListModelFromMap(m.(map[string]interface{})))
		}
		var newMapList []*models.StaticDNSList
		for _, m := range n {
			if m == nil {
				continue
			}
			newMapList = append(newMapList, StaticDNSListModelFromMap(m.(map[string]interface{})))
		}

		return CompareDNSLists(oldMapList, newMapList)
	}
}

func diffSuppressProxyListOrder(mapKey string) schema.SchemaDiffSuppressFunc {
	return func(key, oldValue, newValue string, d *schema.ResourceData) bool {
		oldData, newData := d.GetChange(mapKey)
		if newData == nil && oldData == nil {
			return true
		}

		if oldData == nil {
			return false
		}
		if newData == nil {
			return false
		}

		old := oldData.([]interface{})
		new := newData.([]interface{})
		if len(old) != len(new) {
			return false
		}

		var oldMapList []*models.Proxy
		for _, m := range old {
			if m == nil {
				continue
			}
			oldMapList = append(oldMapList, NetworkProxyModelFromMap(m.(map[string]interface{})))
		}
		var newMapList []*models.Proxy
		for _, m := range new {
			if m == nil {
				continue
			}
			newMapList = append(newMapList, NetworkProxyModelFromMap(m.(map[string]interface{})))
		}

		return CompareProxyLists(oldMapList, newMapList)
	}
}

func diffSuppressSystemInterfaceListOrder(mapKey string) schema.SchemaDiffSuppressFunc {
	return func(key, oldValue, newValue string, d *schema.ResourceData) bool {
		oldData, newData := d.GetChange(mapKey)
		if newData == nil && oldData == nil {
			return true
		}

		if oldData == nil {
			return false
		}
		if newData == nil {
			return false
		}

		old := oldData.([]interface{})
		new := newData.([]interface{})
		if len(old) != len(new) {
			return false
		}

		var oldMapList []*models.SysInterface
		for _, o := range old {
			if o == nil {
				continue
			}
			if oldElem, ok := o.(map[string]interface{}); ok {
				oldMapList = append(oldMapList, SysInterfaceModelFromMap(oldElem))
			}
		}
		var newMapList []*models.SysInterface
		for _, n := range new {
			if n == nil {
				continue
			}
			if newElem, ok := n.(map[string]interface{}); ok {
				newMapList = append(newMapList, SysInterfaceModelFromMap(newElem))
			}
		}

		interfaceMissmatch, reason := CompareSystemInterfaceList(oldMapList, newMapList)
		if interfaceMissmatch {
			log.Printf("SystemInterface list mismatch: %s\n", reason)
		}
		return interfaceMissmatch
	}
}

func diffSuppressVlanAdapterListOrder(mapKey string) schema.SchemaDiffSuppressFunc {
	return func(key, oldValue, newValue string, d *schema.ResourceData) bool {
		oldData, newData := d.GetChange(mapKey)
		if newData == nil && oldData == nil {
			return true
		}

		if oldData == nil {
			return false
		}
		if newData == nil {
			return false
		}

		old := oldData.([]interface{})
		new := newData.([]interface{})
		if len(old) != len(new) {
			return false
		}

		var oldMapList []*models.VlanAdapter
		for _, o := range old {
			if o == nil {
				continue
			}
			if oldElem, ok := o.(map[string]interface{}); ok {
				oldMapList = append(oldMapList, VlanAdapterModelFromMap(oldElem))
			}
		}
		var newMapList []*models.VlanAdapter
		for _, n := range new {
			if n == nil {
				continue
			}
			if newElem, ok := n.(map[string]interface{}); ok {
				newMapList = append(newMapList, VlanAdapterModelFromMap(newElem))
			}
		}

		adapterMismatch, reason := CompareVlanAdapterLists(oldMapList, newMapList)
		if !adapterMismatch {
			log.Printf("VlanAdapter list mismatch: %s\n", reason)
		}
		return adapterMismatch
	}
}

func diffSuppressInterfaceListOrder(mapKey string) schema.SchemaDiffSuppressFunc {
	return func(key, oldValue, newValue string, d *schema.ResourceData) bool {
		oldData, newData := d.GetChange(mapKey)
		if newData == nil && oldData == nil {
			return true
		}

		if oldData == nil {
			return false
		}
		if newData == nil {
			return false
		}

		o := oldData.([]interface{})
		n := newData.([]interface{})
		if len(o) != len(n) {
			return false
		}

		oldMapList := make([]*models.Interface, len(o))
		for i, m := range o {
			oldMapList[i] = InterfaceModelFromMap(m.(map[string]interface{}))
		}
		newMapList := make([]*models.Interface, len(o))
		for i, m := range o {
			newMapList[i] = InterfaceModelFromMap(m.(map[string]interface{}))
		}

		return CompareInterfaceLists(oldMapList, newMapList)
	}
}

// suppressAppInterfacesChangeAfterCreation checks whether we have the newly created AppInterface
// or not. It suppress any differences after an AppInstance is created.
// AppInterfaces are never updated after they have been created
func suppressAppInterfacesChangeAfterCreation(mapKey string) schema.SchemaDiffSuppressFunc {
	return func(key, oldValue, newValue string, d *schema.ResourceData) bool {
		oldData, newData := d.GetChange(mapKey)
		if newData == nil && oldData == nil {
			return true
		}

		if oldData == nil {
			return false
		}
		if newData == nil {
			return false
		}

		o := oldData.([]interface{})
		n := newData.([]interface{})
		if len(o) != len(n) && len(o) == 0 {
			return false
		}

		return true
	}
}

// In most cases it's not allowed to change the resources' name after creation.
// The function supresses any differences between names once the resource is created.
func supressNameAfterCreation(mapKey string) schema.SchemaDiffSuppressFunc {
	return func(key, oldValue, newValue string, d *schema.ResourceData) bool {
		oldData, newData := d.GetChange(mapKey)
		if newData == nil && oldData == nil {
			return true
		}

		if oldData == nil {
			return false
		}
		if newData == nil {
			return false
		}

		o := oldData.(string)
		n := newData.(string)
		if o != n && o == "" {
			return false
		}

		return true
	}
}

func supress() schema.SchemaDiffSuppressFunc {
	return func(key, oldValue, newValue string, d *schema.ResourceData) bool {
		return true
	}
}

func diffSupressSliceOrder(mapKey string) schema.SchemaDiffSuppressFunc {
	return func(key, oldValue, newValue string, d *schema.ResourceData) bool {
		equal, err := compare.Slices(d.GetChange(mapKey))
		if err != nil {
			fmt.Printf("error comparing %s: %+v\n", mapKey, err)
			return false
		}
		return equal
	}
}

func diffSuppressIoMemberListOrder(mapKey string) schema.SchemaDiffSuppressFunc {
	return func(key, oldValue, newValue string, d *schema.ResourceData) bool {
		oldData, newData := d.GetChange(mapKey)
		if newData == nil && oldData == nil {
			return true
		}

		if oldData == nil {
			return false
		}
		if newData == nil {
			return false
		}

		o := oldData.([]interface{})
		n := newData.([]interface{})
		if len(o) != len(n) {
			return false
		}

		oldMapList := make([]*models.IoMember, len(o))
		for i, m := range o {
			oldMapList[i] = IoMemberModelFromMap(m.(map[string]interface{}))
		}
		newMapList := make([]*models.IoMember, len(n))
		for i, m := range n {
			newMapList[i] = IoMemberModelFromMap(m.(map[string]interface{}))
		}

		return CompareIoMemberLists(oldMapList, newMapList)
	}
}

func diffSuppressChangedList(mapKey string) schema.SchemaDiffSuppressFunc {
	return func(key, oldValue, newValue string, d *schema.ResourceData) bool {
		oldData, newData := d.GetChange(mapKey)
		if newData == nil && oldData == nil {
			return true
		}

		o := oldData.([]interface{})
		n := newData.([]interface{})
		if len(o) == 0 || len(n) == 0 {
			return true
		}
		if len(o) != len(n) {
			return false
		}

		return true
	}
}

func diffSupressMapInterfaceNonConfigChanges(field string, mapKey string) schema.SchemaDiffSuppressFunc {
	return func(key, oldValue, newValue string, d *schema.ResourceData) bool {
		oldData, newData := d.GetChange(field)

		idOld := ""
		oldDataList, ok := oldData.([]interface{})
		if ok && len(oldDataList) > 0 {
			gotOld, okOld := oldDataList[0].(map[string]interface{})
			if okOld {
				idOld, _ = gotOld[mapKey].(string)
			}
		}

		newDataList, ok := newData.([]interface{})
		if !ok || len(newDataList) == 0 {
			return true
		}

		gotNew, okNew := newDataList[0].(map[string]interface{})
		if !okNew {
			return true
		}

		id, _ := gotNew[mapKey].(string)
		if id == "" {
			return true
		} else if id == idOld {
			return true
		}

		return false
	}
}

func diffSupressMapInterfaceNonConfigChangesV2(field string, mapKeys []string) schema.SchemaDiffSuppressFunc {
	return func(key, oldValue, newValue string, d *schema.ResourceData) bool {
		oldData, newData := d.GetChange(field)

		oldValues := make(map[string]string, len(mapKeys))
		oldDataList, ok := oldData.([]interface{})
		if ok {
			for _, oldData := range oldDataList {
				gotOld, okOld := oldData.(map[string]interface{})
				if okOld {
					for _, mapKey := range mapKeys {
						idOld, _ := gotOld[mapKey].(string)
						if idOld != "" {
							oldValues[mapKey] = idOld
						}
					}
				}
			}
		}

		newDataList, ok := newData.([]interface{})
		if !ok || len(newDataList) == 0 {
			return true
		}

		for _, newData := range newDataList {
			gotNew, okNew := newData.(map[string]interface{})
			if !okNew {
				return true
			}

			allEqual := true

			for _, mapKey := range mapKeys {
				val, _ := gotNew[mapKey].(string)
				if val == "" {
					// Maybe the mapKey is not a string, but a bool
					if _, ok := gotNew[mapKey].(bool); !ok {
						return true
					}
				} else if val != oldValues[mapKey] {
					allEqual = false
					break
				}
			}

			if allEqual {
				return true
			}
		}

		return false
	}
}

func diffSupressStringNonConfigChanges(field string) schema.SchemaDiffSuppressFunc {
	return func(key, oldValue, newValue string, d *schema.ResourceData) bool {
		oldData, newData := d.GetChange(field)
		if oldData == nil || newData == nil {
			return true
		}

		old, _ := oldData.(string)
		new, okNew := newData.(string)
		if !okNew || new == "" || old == new {
			return true
		}

		return old == new
	}
}

func diffSuppressIfFieldValueEqual(field, value string) schema.SchemaDiffSuppressFunc {
	return func(key, oldValue, newValue string, d *schema.ResourceData) bool {
		_, newData := d.GetChange(field)
		if newData == nil {
			return false
		}
		newVal, okNew := newData.(string)
		if !okNew || newVal == "" {
			return false
		}
		if newVal == value {
			return true
		}

		return false
	}
}

func isFieldDefinedInTerraformConfig(d *schema.ResourceData, field string) bool {
	rawConfig := d.GetRawConfig()
	if !rawConfig.IsKnown() || rawConfig.IsNull() {
		return false
	}

	rawConfigMap := rawConfig.AsValueMap()
	rawField, ok := rawConfigMap[field]
	if !ok {
		return false
	}

	return rawField.IsKnown() && !rawField.IsNull()
}

func diffSuppressBaseImage(field string) schema.SchemaDiffSuppressFunc {
	return func(key, oldValue, newValue string, d *schema.ResourceData) bool {
		if !isFieldDefinedInTerraformConfig(d, field) {
			return true
		}

		oldData, newData := d.GetChange(field)
		if oldData == nil && newData == nil {
			return true
		}

		if oldData == nil || newData == nil {
			return false
		}

		return reflect.DeepEqual(oldData, newData)
	}
}

func diffSuppressEdgeviewConfig(field string) schema.SchemaDiffSuppressFunc {
	return func(key, oldValue, newValue string, d *schema.ResourceData) bool {
		if !isFieldDefinedInTerraformConfig(d, field) {
			return true
		}

		oldData, newData := d.GetChange(field)
		if oldData == nil && newData == nil {
			return true
		}

		if oldData == nil || newData == nil {
			return false
		}

		return reflect.DeepEqual(oldData, newData)
	}
}

// compareSysInterfaceIgnoringAutoFields compares two SysInterface structs for
// equality, deliberately skipping Intfname and Ztype because the API
// auto-populates them (Intfname is set to the adapter's logical label;
// Ztype is set to "IO_TYPE_ETH") and we must not treat those API-assigned
// values as user-driven changes.
func compareSysInterfaceIgnoringAutoFields(x, y *models.SysInterface) (bool, string) {
	if x == nil && y == nil {
		return true, ""
	}
	if x == nil || y == nil {
		return false, "one interface is nil, other is not"
	}
	if x.Cost != y.Cost {
		return false, fmt.Sprintf("Cost mismatch: %d vs %d", x.Cost, y.Cost)
	}
	if x.IntfUsage != nil && y.IntfUsage != nil {
		if *x.IntfUsage != *y.IntfUsage {
			return false, fmt.Sprintf("IntfUsage mismatch: %s vs %s", *x.IntfUsage, *y.IntfUsage)
		}
	} else if x.IntfUsage != y.IntfUsage {
		return false, "IntfUsage mismatch: one is nil, other is not"
	}
	if x.Ipaddr != y.Ipaddr {
		return false, fmt.Sprintf("Ipaddr mismatch: %s vs %s", x.Ipaddr, y.Ipaddr)
	}
	if x.Macaddr != y.Macaddr {
		return false, fmt.Sprintf("Macaddr mismatch: %s vs %s", x.Macaddr, y.Macaddr)
	}
	if x.Netname != y.Netname {
		return false, fmt.Sprintf("Netname mismatch: %s vs %s", x.Netname, y.Netname)
	}
	if x.Netid != y.Netid {
		return false, fmt.Sprintf("Netid mismatch: %s vs %s", x.Netid, y.Netid)
	}
	// Note: Intfname and Ztype are intentionally excluded — the API auto-assigns
	// them and they should not trigger plan diffs.
	return true, ""
}

func diffSuppressBondAdapterListOrder(mapKey string) schema.SchemaDiffSuppressFunc {
	return func(key, oldValue, newValue string, d *schema.ResourceData) bool {
		oldData, newData := d.GetChange(mapKey)
		if newData == nil && oldData == nil {
			return true
		}
		if oldData == nil || newData == nil {
			return false
		}

		old := oldData.([]interface{})
		new := newData.([]interface{})
		if len(old) != len(new) {
			return false
		}

		var oldList []*models.BondAdapter
		for _, o := range old {
			if o == nil {
				continue
			}
			if elem, ok := o.(map[string]interface{}); ok {
				oldList = append(oldList, BondAdapterModelFromMap(elem))
			}
		}
		var newList []*models.BondAdapter
		for _, n := range new {
			if n == nil {
				continue
			}
			if elem, ok := n.(map[string]interface{}); ok {
				newList = append(newList, BondAdapterModelFromMap(elem))
			}
		}

		equal, reason := compareBondAdaptersForDiffSuppress(oldList, newList)
		if !equal {
			log.Printf("BondAdapter list mismatch: %s\n", reason)
		}
		return equal
	}
}

// compareBondAdaptersForDiffSuppress compares two BondAdapter lists for diff
// suppression. It is more thorough than CompareBondAdapterLists (which is used
// in acceptance tests against YAML snapshots that omit the Interface sub-block)
// because it also checks the nested SysInterface — but using the lenient
// compareSysInterfaceIgnoringAutoFields helper so that API-assigned Intfname /
// Ztype values are never treated as user-driven changes.
func compareBondAdaptersForDiffSuppress(a, b []*models.BondAdapter) (bool, string) {
	if len(a) != len(b) {
		return false, fmt.Sprintf("bond_adapters length mismatch: %d vs %d", len(a), len(b))
	}

	slices.SortFunc(a, func(i, j *models.BondAdapter) int {
		return strings.Compare(strings.ToLower(i.LogicalLabel), strings.ToLower(j.LogicalLabel))
	})
	slices.SortFunc(b, func(i, j *models.BondAdapter) int {
		return strings.Compare(strings.ToLower(i.LogicalLabel), strings.ToLower(j.LogicalLabel))
	})

	reason := ""
	equal := slices.EqualFunc(a, b, func(x, y *models.BondAdapter) bool {
		if x.LogicalLabel != y.LogicalLabel {
			reason = fmt.Sprintf("LogicalLabel mismatch: %s vs %s", x.LogicalLabel, y.LogicalLabel)
			return false
		}
		if x.BondMode != nil && y.BondMode != nil {
			if *x.BondMode != *y.BondMode {
				reason = fmt.Sprintf("BondMode mismatch: %s vs %s", *x.BondMode, *y.BondMode)
				return false
			}
		} else if x.BondMode != y.BondMode {
			reason = "BondMode mismatch: one is nil, other is not"
			return false
		}
		xLacpRate := models.LacpRateLACPRATEUNSPECIFIED
		if x.LacpRate != nil {
			xLacpRate = *x.LacpRate
		}
		yLacpRate := models.LacpRateLACPRATEUNSPECIFIED
		if y.LacpRate != nil {
			yLacpRate = *y.LacpRate
		}
		if xLacpRate != yLacpRate {
			reason = fmt.Sprintf("LacpRate mismatch: %s vs %s", xLacpRate, yLacpRate)
			return false
		}
		xNames := slices.Clone(x.LowerLayerNames)
		yNames := slices.Clone(y.LowerLayerNames)
		slices.Sort(xNames)
		slices.Sort(yNames)
		if !slices.Equal(xNames, yNames) {
			reason = fmt.Sprintf("LowerLayerNames mismatch: %v vs %v", x.LowerLayerNames, y.LowerLayerNames)
			return false
		}
		if x.Arp != nil && y.Arp != nil {
			if x.Arp.Interval != y.Arp.Interval {
				reason = fmt.Sprintf("Arp.Interval mismatch: %d vs %d", x.Arp.Interval, y.Arp.Interval)
				return false
			}
		} else if x.Arp != y.Arp {
			reason = "Arp mismatch: one is nil, other is not"
			return false
		}
		ok, ifaceReason := compareSysInterfaceIgnoringAutoFields(x.Interface, y.Interface)
		if !ok {
			reason = fmt.Sprintf("Interface mismatch: %s", ifaceReason)
			return false
		}
		return true
	})
	return equal, reason
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// compareStringPointers compares two string pointers for equality.
// if both are nil, they are equal.
// if one is nil and the other is empty string, they are equal.
// otherwise, compare the string values.
func compareStringPointers(x, y *string) bool {
	// normalize converts nil pointer to empty string
	normalize := func(s *string) string {
		if s == nil {
			return ""
		}
		return *s
	}
	return normalize(x) == normalize(y)
}
