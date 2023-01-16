package resources

import (
	"os"
	"path/filepath"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"zedcloud": testAccProvider,
	}
}

func getTestConfig(path string) (string, error) {
	testdataDir, err := filepath.Abs("./testdata")
	if err != nil {
		return "", err
	}
	bytes, err := os.ReadFile(filepath.Join(testdataDir, path))
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
