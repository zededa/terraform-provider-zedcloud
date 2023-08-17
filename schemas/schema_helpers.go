package schemas

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
	"github.com/zededa/terraform-provider/pkg/compare"
	"golang.org/x/exp/slices"
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
		if len(o) != len(n) {
			return false
		}

		oldMapList := make([]*models.StaticDNSList, len(o))
		for i, m := range o {
			oldMapList[i] = StaticDNSListModelFromMap(m.(map[string]interface{}))
		}
		newMapList := make([]*models.StaticDNSList, len(o))
		for i, m := range o {
			newMapList[i] = StaticDNSListModelFromMap(m.(map[string]interface{}))
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

		o := oldData.([]interface{})
		n := newData.([]interface{})
		if len(o) != len(n) {
			return false
		}

		oldMapList := make([]*models.Proxy, len(o))
		for i, m := range o {
			oldMapList[i] = NetworkProxyModelFromMap(m.(map[string]interface{}))
		}
		newMapList := make([]*models.Proxy, len(o))
		for i, m := range o {
			newMapList[i] = NetworkProxyModelFromMap(m.(map[string]interface{}))
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

		o := oldData.([]interface{})
		n := newData.([]interface{})
		if len(o) != len(n) {
			return false
		}

		oldMapList := make([]*models.SystemInterface, len(o))
		for i, m := range o {
			oldMapList[i] = SystemInterfaceModelFromMap(m.(map[string]interface{}))
		}
		newMapList := make([]*models.SystemInterface, len(o))
		for i, m := range o {
			newMapList[i] = SystemInterfaceModelFromMap(m.(map[string]interface{}))
		}

		return CompareSystemInterfaceList(oldMapList, newMapList)
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
