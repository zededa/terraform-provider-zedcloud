package testing

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/ghodss/yaml"
)

func MustGetTestInput(t *testing.T, path string) string {
	testdataDir, err := filepath.Abs("./testdata")
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := os.ReadFile(filepath.Join(testdataDir, path))
	if err != nil {
		t.Fatal(err)
	}
	return string(bytes)
}

func MustGetExpectedOutput(t *testing.T, path string, i interface{}) {
	testdataDir, err := filepath.Abs("./testdata")
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := os.ReadFile(filepath.Join(testdataDir, path))
	if err != nil {
		t.Fatal(err)
	}
	jsonBytes, _ := yaml.YAMLToJSON(bytes)
	if err := json.Unmarshal(jsonBytes, i); err != nil {
		t.Fatal(err)
	}
}

func ToYAML(path string, i interface{}) {
	testdataDir, err := filepath.Abs("./testdata")
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(filepath.Join(testdataDir, path), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	yamlBytes, err := yaml.Marshal(i)
	if err != nil {
		panic(err)
	}

	_, err = file.Write(yamlBytes)
	if err != nil {
		panic(err)
	}
}

// CheckEnv validates the necessary test API keys exist in the testing environment.
func CheckEnv(t *testing.T) {
	if v := os.Getenv("TF_CLI_CONFIG_FILE"); v == "" {
		t.Fatal("TF_CLI_CONFIG_FILE must be set for acceptance tests, it should contain the dev_overrides config that points to local instance of the provider")
	}
	if v := os.Getenv("TF_VAR_zedcloud_token"); v == "" {
		t.Fatal("TF_VAR_zedcloud_token must be set for acceptance tests to access the zedcloud API")
	}
}
