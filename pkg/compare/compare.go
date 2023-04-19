package compare

import (
	"encoding/json"
	"fmt"
)

func Slices(a, b interface{}) (bool, error) {
	if a == nil && b == nil {
		return true, nil
	}

	if b == nil {
		return false, nil
	}
	if a == nil {
		return false, nil
	}
	as, bs, err := serializeSlice(a, b)
	if err != nil {
		return false, fmt.Errorf("could not serialize interfaces: %+v", err)
	}
	return equalSlices(as, bs), nil
}

func serializeSlice(a, b interface{}) ([]interface{}, []interface{}, error) {
	jsonBytes, err := json.Marshal(a)
	if err != nil {
		return nil, nil, err
	}
	var as []interface{}
	err = json.Unmarshal(jsonBytes, &as)
	if err != nil {
		return nil, nil, err
	}

	jsonBytes, err = json.Marshal(b)
	if err != nil {
		return nil, nil, err
	}
	var bs []interface{}
	err = json.Unmarshal(jsonBytes, &bs)
	if err != nil {
		return nil, nil, err
	}

	return as, bs, nil
}

// Recursively test the elements in the map for value equality.
func equalMaps(map1, map2 map[string]interface{}) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, value1 := range map1 {
		value2, ok := map2[key]
		if !ok {
			return false
		}

		switch v1 := value1.(type) {
		case map[string]interface{}:
			v2, ok := value2.(map[string]interface{})
			if !ok || !equalMaps(v1, v2) {
				return false
			}
		case []interface{}:
			v2, ok := value2.([]interface{})
			if !ok || !equalSlices(v1, v2) {
				return false
			}
		default:
			if value1 != value2 {
				return false
			}
		}
	}

	return true
}

// Recursively test the elements in the slice for value equality, ignoring the order of elements in the slices.
func equalSlices(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}

	// is each element of list a in list b?
	for _, value1 := range a {
		found := false
		for _, value2 := range b {
			switch v1 := value1.(type) {
			case map[string]interface{}:
				v2, ok := value2.(map[string]interface{})
				if ok && equalMaps(v1, v2) {
					found = true
				}
			case []interface{}:
				v2, ok := value2.([]interface{})
				if ok && equalSlices(v1, v2) {
					found = true
				}
			default:
				if value1 == value2 {
					found = true
				}
			}
			if found == true {
				break
			}
		}
		if !found {
			return false
		}
	}

	// is each element of list b in list a?
	for _, value1 := range b {
		found := false
		for _, value2 := range a {
			switch v1 := value1.(type) {
			case map[string]interface{}:
				v2, ok := value2.(map[string]interface{})
				if ok && equalMaps(v1, v2) {
					found = true
				}
			case []interface{}:
				v2, ok := value2.([]interface{})
				if ok && equalSlices(v1, v2) {
					found = true
				}
			default:
				if value1 == value2 {
					found = true
				}
			}
			if found == true {
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}
