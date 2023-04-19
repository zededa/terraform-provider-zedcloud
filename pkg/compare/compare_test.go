package compare

import (
	"testing"

	"github.com/zededa/terraform-provider/models"
	testhelper "github.com/zededa/terraform-provider/testing"
)

// This test code uses a table-driven approach, where each test case is represented by a struct with the necessary inputs
// and expected output. The test function iterates through the test cases and runs a sub-test for each case. If the actual
// result doesn't match the expected result, the test fails with an error message.
func TestCompareMaps(t *testing.T) {
	testCases := []struct {
		name     string
		map1     map[string]interface{}
		map2     map[string]interface{}
		expected bool
	}{
		{
			name: "equal maps",
			map1: map[string]interface{}{
				"a": "a",
				"b": 30,
				"c": true,
				"d": []interface{}{"a", "b"},
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			map2: map[string]interface{}{
				"a": "a",
				"b": 30,
				"c": true,
				"d": []interface{}{"a", "b"},
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			expected: true,
		},
		{
			name: "unequal maps",
			map1: map[string]interface{}{
				"a": "a",
				"b": 30,
				"c": true,
				"d": []interface{}{"a", "b"},
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			map2: map[string]interface{}{
				"a": "Jane",
				"b": 35,
				"c": false,
				"d": []interface{}{"c", "d"},
				"e": map[string]interface{}{
					"f": "Foo",
					"g": "Bar",
					"h": "Baz",
					"i": "67890",
				},
			},
			expected: false,
		},
		{
			name: "maps with different values",
			map1: map[string]interface{}{
				"a": "a",
				"b": 30,
				"c": true,
				"d": []interface{}{"a", "b"},
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			map2: map[string]interface{}{
				"a": "a",
				"b": 30,
				"c": true,
				"d": []interface{}{"a", "b"},
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
				},
			},
			expected: false,
		},
		{
			name: "maps with different keys",
			map1: map[string]interface{}{
				"a": "a",
				"b": 30,
				"c": true,
				"d": []interface{}{"a", "b"},
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			map2: map[string]interface{}{
				"a": "a",
				"b": 30,
				"c": true,
				"d": []interface{}{"a", "b"},
				"location": map[string]interface{}{
					"latitude":  37.7749,
					"longitude": -122.4194,
				},
			},
			expected: false,
		},
		{
			name:     "maps with empty values",
			map1:     map[string]interface{}{},
			map2:     map[string]interface{}{},
			expected: true,
		},
		{
			name: "equal maps, slices with different order",
			map1: map[string]interface{}{
				"a": "a",
				"b": 30,
				"c": true,
				"d": []interface{}{"b", "a"},
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			map2: map[string]interface{}{
				"a": "a",
				"b": 30,
				"c": true,
				"d": []interface{}{"a", "b"},
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			expected: true,
		},
		{
			name: "equal maps, slices with map elements in same order",
			map1: map[string]interface{}{
				"a": "a",
				"d": []interface{}{
					map[string]interface{}{
						"f": "foo",
						"g": "bar",
						"h": "baz",
						"i": 12345,
					},
					map[string]interface{}{
						"f": "foobar",
						"g": "barfoo",
						"h": "bazbar",
						"i": "1234",
					},
				},
				"b": 30,
				"c": true,
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			map2: map[string]interface{}{
				"a": "a",
				"b": 30,
				"c": true,
				"d": []interface{}{
					map[string]interface{}{
						"f": "foo",
						"g": "bar",
						"h": "baz",
						"i": 12345,
					},
					map[string]interface{}{
						"f": "foobar",
						"g": "barfoo",
						"h": "bazbar",
						"i": "1234",
					},
				},
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			expected: true,
		},
		{
			name: "equal maps, unequal slices with map elements",
			map1: map[string]interface{}{
				"a": "a",
				"d": []interface{}{
					map[string]interface{}{
						"g": "bar",
						"h": "baz",
						"i": 12345,
					},
					map[string]interface{}{
						"f": "foobar",
						"g": "barfoo",
						"h": "bazbar",
						"i": "1234",
					},
				},
				"b": 30,
				"c": true,
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			map2: map[string]interface{}{
				"a": "a",
				"b": 30,
				"c": true,
				"d": []interface{}{
					map[string]interface{}{
						"f": "foo",
						"g": "bar",
						"h": "baz",
						"i": 12345,
					},
					map[string]interface{}{
						"f": "foobar",
						"g": "barfoo",
						"h": "bazbar",
						"i": "1234",
					},
				},
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			expected: false,
		},
		{
			name: "equal maps, slices with map elements in different order",
			map1: map[string]interface{}{
				"a": "a",
				"d": []interface{}{
					map[string]interface{}{
						"f": "foobar",
						"g": "barfoo",
						"h": "bazbar",
						"i": "1234",
					},
					map[string]interface{}{
						"f": "foo",
						"g": "bar",
						"h": "baz",
						"i": 12345,
					},
				},
				"b": 30,
				"c": true,
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			map2: map[string]interface{}{
				"a": "a",
				"b": 30,
				"c": true,
				"d": []interface{}{
					map[string]interface{}{
						"f": "foo",
						"g": "bar",
						"h": "baz",
						"i": 12345,
					},
					map[string]interface{}{
						"f": "foobar",
						"g": "barfoo",
						"h": "bazbar",
						"i": "1234",
					},
				},
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			expected: true,
		},
		{
			name: "equal maps, slices with unequal map elements in different order",
			map1: map[string]interface{}{
				"a": "a",
				"d": []interface{}{
					map[string]interface{}{
						"f": "foobar",
						"g": "barfoo",
						"h": "bazbar",
						"i": "1234",
					},
					map[string]interface{}{
						"f": "foo",
						"g": "bar",
						"h": "baz",
					},
				},
				"b": 30,
				"c": true,
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			map2: map[string]interface{}{
				"a": "a",
				"b": 30,
				"c": true,
				"d": []interface{}{
					map[string]interface{}{
						"f": "foo",
						"g": "bar",
						"h": "baz",
						"i": 12345,
					},
					map[string]interface{}{
						"f": "foobar",
						"g": "barfoo",
						"h": "bazbar",
						"i": "1234",
					},
				},
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			expected: false,
		},
		{
			name: "equal maps, unequal slices in same order with slice elements",
			map1: map[string]interface{}{
				"a": "a",
				"d": []interface{}{
					[]interface{}{
						"c",
						"d",
					},
				},
				"b": 30,
				"c": true,
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			map2: map[string]interface{}{
				"a": "a",
				"b": 30,
				"c": true,
				"d": []interface{}{
					[]interface{}{
						"a",
						"b",
					},
					[]interface{}{
						"c",
						"d",
					},
				},
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			expected: false,
		},
		{
			name: "equal maps, slices in same order with slice elements",
			map1: map[string]interface{}{
				"a": "a",
				"d": []interface{}{
					[]interface{}{
						"a",
						"b",
					},
					[]interface{}{
						"c",
						"d",
					},
				},
				"b": 30,
				"c": true,
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			map2: map[string]interface{}{
				"a": "a",
				"b": 30,
				"c": true,
				"d": []interface{}{
					[]interface{}{
						"a",
						"b",
					},
					[]interface{}{
						"c",
						"d",
					},
				},
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			expected: true,
		},
		{
			name: "equal maps, equal slices in same order with unequal slice elements",
			map1: map[string]interface{}{
				"a": "a",
				"d": []interface{}{
					[]interface{}{
						"a",
						"b",
					},
					[]interface{}{
						"d",
					},
				},
				"b": 30,
				"c": true,
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			map2: map[string]interface{}{
				"a": "a",
				"b": 30,
				"c": true,
				"d": []interface{}{
					[]interface{}{
						"a",
						"b",
					},
					[]interface{}{
						"c",
						"d",
					},
				},
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			expected: false,
		},
		{
			name: "equal maps, slices in different order with slice elements in same order",
			map1: map[string]interface{}{
				"a": "a",
				"d": []interface{}{
					[]interface{}{
						"c",
						"d",
					},
					[]interface{}{
						"a",
						"b",
					},
				},
				"b": 30,
				"c": true,
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			map2: map[string]interface{}{
				"a": "a",
				"b": 30,
				"c": true,
				"d": []interface{}{
					[]interface{}{
						"a",
						"b",
					},
					[]interface{}{
						"c",
						"d",
					},
				},
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			expected: true,
		},
		{
			name: "equal maps, equal slices in different order with unequal slice elements in same order",
			map1: map[string]interface{}{
				"a": "a",
				"d": []interface{}{
					[]interface{}{
						"c",
						"d",
					},
					[]interface{}{
						"b",
					},
				},
				"b": 30,
				"c": true,
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			map2: map[string]interface{}{
				"a": "a",
				"b": 30,
				"c": true,
				"d": []interface{}{
					[]interface{}{
						"a",
						"b",
					},
					[]interface{}{
						"c",
						"d",
					},
				},
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			expected: false,
		},
		{
			name: "equal maps, slices in different order with slice elements in different order",
			map1: map[string]interface{}{
				"a": "a",
				"d": []interface{}{
					[]interface{}{
						"d",
						"c",
					},
					[]interface{}{
						"b",
						"a",
					},
				},
				"b": 30,
				"c": true,
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			map2: map[string]interface{}{
				"a": "a",
				"b": 30,
				"c": true,
				"d": []interface{}{
					[]interface{}{
						"a",
						"b",
					},
					[]interface{}{
						"c",
						"d",
					},
				},
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			expected: true,
		},
		{
			name: "equal maps, slices in different order with slice elements in different order with map elements with slices in different order with slice elements in different order",
			map1: map[string]interface{}{
				"a": "a",
				"d": []interface{}{
					[]interface{}{
						map[string]interface{}{
							"f": "foo",
							"g": "bar",
							"h": "baz",
							"i": 12345,
						},
						map[string]interface{}{
							"g": "bar",
							"h": "bazbar",
							"i": 12345,
							"d": []interface{}{
								[]interface{}{
									"a",
									"b",
								},
								[]interface{}{
									"c",
									"d",
								},
							},
						},
					},
					[]interface{}{
						map[string]interface{}{
							"g": "bar",
							"h": "bazbar",
							"i": 12345,
							"d": []interface{}{
								[]interface{}{
									"a",
									"b",
								},
								[]interface{}{
									"c",
									"d",
								},
							},
						},
						map[string]interface{}{
							"f": "foo",
							"g": "bar",
							"h": "baz",
							"i": 12345,
						},
					},
				},
				"b": 30,
				"c": true,
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			map2: map[string]interface{}{
				"a": "a",
				"d": []interface{}{
					[]interface{}{
						map[string]interface{}{
							"g": "bar",
							"h": "bazbar",
							"i": 12345,
							"d": []interface{}{
								[]interface{}{
									"d",
									"c",
								},
								[]interface{}{
									"b",
									"a",
								},
							},
						},
						map[string]interface{}{
							"f": "foo",
							"g": "bar",
							"h": "baz",
							"i": 12345,
						},
					},
					[]interface{}{
						map[string]interface{}{
							"f": "foo",
							"g": "bar",
							"h": "baz",
							"i": 12345,
						},
						map[string]interface{}{
							"g": "bar",
							"h": "bazbar",
							"i": 12345,
							"d": []interface{}{
								[]interface{}{
									"d",
									"c",
								},
								[]interface{}{
									"b",
									"a",
								},
							},
						},
					},
				},
				"b": 30,
				"c": true,
				"e": map[string]interface{}{
					"f": "foo",
					"g": "bar",
					"h": "baz",
					"i": 12345,
				},
			},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := equalMaps(tc.map1, tc.map2)
			if result != tc.expected {
				t.Errorf("Expected %t but got %t", tc.expected, result)
			}
		})
	}
}

// change acl order and mapparams port
func TestCompareSlices(t *testing.T) {
	var a, b models.AppInstance
	testhelper.MustGetExpectedOutput(t, "application_instance_a.yaml", &a)
	testhelper.MustGetExpectedOutput(t, "application_instance_b.yaml", &b)

	euqal, err := Slices(a.Interfaces, b.Interfaces)
	if err != nil {
		t.Error(err)
	}

	if euqal {
		t.Fatal("expect not euqal")
	}
}
