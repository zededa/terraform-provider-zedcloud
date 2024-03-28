package resources

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	config "github.com/zededa/terraform-provider-zedcloud/v2/client/datastore"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	testhelper "github.com/zededa/terraform-provider-zedcloud/v2/testing"
)

func TestDatastore_Create(t *testing.T) {
	var got models.Datastore
	var expected models.Datastore

	// input config
	inputPath := "datastore/create.tf"
	input := testhelper.MustGetTestInput(t, inputPath)

	// expected output
	expectedPath := "datastore/create.yaml"
	testhelper.MustGetExpectedOutput(t, expectedPath, &expected)

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { checkDatastoreEnv(t) },
		CheckDestroy: testDatastoreDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: input,
				Check: resource.ComposeTestCheckFunc(
					testDatastoreExists("zedcloud_datastore.test_datastore", &got),
					resource.TestMatchResourceAttr(
						"zedcloud_datastore.test_datastore",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testDatastoreAttributes(t, &got, &expected),
				),
			},
		},
	})
}

func checkDatastoreEnv(t *testing.T) {
	if v := os.Getenv("TF_CLI_CONFIG_FILE"); v == "" {
		t.Fatal("TF_CLI_CONFIG_FILE must be set for acceptance tests, it should contain the dev_overrides config that points to local instance of the provider")
	}
	if v := os.Getenv("TF_VAR_zedcloud_token"); v == "" {
		t.Fatal("TF_VAR_zedcloud_token must be set for acceptance tests to access the zedcloud API")
	}
	// if v := os.Getenv("TF_VAR_api_key"); v == "" {
	// 	t.Fatal("TF_VAR_api_key must be set for acceptance tests to access the zedcloud API")
	// }
	// if v := os.Getenv("TF_VAR_api_password"); v == "" {
	// 	t.Fatal("TF_VAR_api_password must be set for acceptance tests to access the zedcloud API")
	// }
}

func testDatastoreExists(resourceName string, datastoreModel *models.Datastore) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Datastore not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Datastore ID is not set")
		}

		// retrieve the client established in Provider configuration
		client := testProvider.Meta().(*api_client.ZedcloudAPI)

		// retrieve the Datastore by referencing its state ID for API lookup
		params := config.GetByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.Datastore.GetByID(params, nil)
		if err != nil {
			return fmt.Errorf("could not fetch Datastore (%s): %w", rs.Primary.ID, err)
		}

		datastore := response.GetPayload()
		if datastore == nil {
			return errors.New("could not get response payload in Datastore existence test: datastore is nil")
		}

		*datastoreModel = *datastore
		return nil
	}
}

// testDatastoreAttributes verifies attributes are set correctly by Terraform
func testDatastoreAttributes(t *testing.T, got, expected *models.Datastore) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ignoredFields := []string{
			"ID",
			"Revision",
			"EncryptedSecrets",
		}
		// if expected.Proxy != nil && expected.Proxy.Proxies != nil {
		// 	ignoredFields = append(ignoredFields, "Proxy.Proxies")
		// }
		opts := cmpopts.IgnoreFields(models.Datastore{}, ignoredFields...)
		if diff := cmp.Diff(*got, *expected, opts); len(diff) != 0 {
			return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
		}
		return nil
	}
}

// testDatastoreInstanceDestroy verifies the Datastore has been destroyed.
func testDatastoreDestroy(s *terraform.State) error {
	// retrieve the client established in Provider configuration
	client := testProvider.Meta().(*api_client.ZedcloudAPI)

	// loop through the resources in state, verifying each Datastore is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zedcloud_datastore" {
			continue
		}

		// retrieve the Datastore by referencing it's state ID for API lookup
		params := config.GetByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.Datastore.GetByID(params, nil)
		if err == nil {
			if datastore := response.GetPayload(); datastore != nil && datastore.ID == rs.Primary.ID {
				return fmt.Errorf("destroy failed, Datastore (%s) still exists", datastore.ID)
			}
			return nil
		}

		// if the error is equivalent to 404 not found, the Datastore is destroyed
		_, ok := err.(*config.GetByIDNotFound)
		if !ok {
			return fmt.Errorf("destroy failed, expect status code 404 for Datastore (%s)", params.ID)
		}
	}
	return nil
}
