package test

// In JSON all numbers are floats, this casts ints to int recursively.
func MapItemsFloatToInt(m map[string]interface{}) map[string]interface{} {
	for k, v := range m {
		switch v.(type) {
		case float64:
			// check the v is integer or float
			if isIntegral(v.(float64)) {
				m[k] = int(v.(float64))
			}
		case map[string]interface{}:
			m[k] = MapItemsFloatToInt(m[k].(map[string]interface{}))
		case []interface{}:
			m[k] = sliceItemsFloatToInt(m[k].([]interface{}))
		}
	}
	return m
}

func isIntegral(val float64) bool {
	return val == float64(int(val))
}

func sliceItemsFloatToInt(s []interface{}) []interface{} {
	for i, v := range s {
		switch v.(type) {
		case float64:
			// check the v is integer or float
			if isIntegral(v.(float64)) {
				s[i] = int(v.(float64))
			}
		case map[string]interface{}:
			s[i] = MapItemsFloatToInt(v.(map[string]interface{}))
		case []interface{}:
			s[i] = sliceItemsFloatToInt(v.([]interface{}))
		}
	}
	return s
}
