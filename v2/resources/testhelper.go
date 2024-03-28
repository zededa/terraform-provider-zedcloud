package resources

import (
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v3"
)

func mustGetTestInput(t *testing.T, path string) string {
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

func mustGetExpectedOutput(t *testing.T, path string, i interface{}) {
	testdataDir, err := filepath.Abs("./testdata")
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := os.ReadFile(filepath.Join(testdataDir, path))
	if err != nil {
		t.Fatal(err)
	}
	if err := yaml.Unmarshal(bytes, i); err != nil {
		t.Fatal(err)
	}
}

func toYAML(path string, i interface{}) {
	testdataDir, err := filepath.Abs("./testdata")
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(filepath.Join(testdataDir, path), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = yaml.NewEncoder(file).Encode(i)
	if err != nil {
		panic(err)
	}
}

// checkEnv validates the necessary test API keys exist in the testing environment.
func checkEnv(t *testing.T) {
	if v := os.Getenv("TF_CLI_CONFIG_FILE"); v == "" {
		t.Fatal("TF_CLI_CONFIG_FILE must be set for acceptance tests, it should contain the dev_overrides config that points to local instance of the provider")
	}
	if v := os.Getenv("TF_VAR_zedcloud_token"); v == "" {
		t.Fatal("TF_VAR_zedcloud_token must be set for acceptance tests to access the zedcloud API")
	}
}
