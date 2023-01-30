package schemas

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/exp/slices"
)

// diffSuppressStringListOrder suppresses diffs in schema.List element order.
// If the lists elements differ, it still reports a diff.
// Note, this is slow since it is called for every element of the schema.List.
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

func diffSuppressMapListOrder(mapKey string) schema.SchemaDiffSuppressFunc {
	// fmt.Println("=========================================================")
	return func(key, oldValue, newValue string, d *schema.ResourceData) bool {
		fmt.Println("start =========================================================")
		fmt.Println(mapKey)
		fmt.Println(key)
		oldData, newData := d.GetChange(mapKey)
		if oldData == nil || newData == nil {
			return false
		}

		oldMap := oldData.([]interface{})
		newMap := newData.([]interface{})

		fmt.Println("oldMap")
		fmt.Println(oldMap)
		fmt.Println("newMap")
		fmt.Println(newMap)

		fmt.Println("end =========================================================")
		return true
	}
}
