package resources

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	api_client "github.com/zededa/terraform-provider/client"
	config "github.com/zededa/terraform-provider/client/application_instance"
	"github.com/zededa/terraform-provider/models"
	testhelper "github.com/zededa/terraform-provider/testing"
)

func TestApplicationInstance_Create(t *testing.T) {
	var got models.AppInstance
	var expected models.AppInstance

	// input config
	inputPath := "application_instance/create.tf"
	input := testhelper.MustGetTestInput(t, inputPath)

	// expected output
	expectedPath := "application_instance/create.yaml"
	testhelper.MustGetExpectedOutput(t, expectedPath, &expected)

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testApplicationInstanceDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: input,
				Check: resource.ComposeTestCheckFunc(
					testApplicationInstanceExists("zedcloud_application_instance.test_tf_provider", &got),
					resource.TestMatchResourceAttr(
						"zedcloud_application_instance.test_tf_provider",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					resource.TestMatchResourceAttr(
						"zedcloud_application_instance.test_tf_provider",
						"app_id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					resource.TestMatchResourceAttr(
						"zedcloud_application_instance.test_tf_provider",
						"device_id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testApplicationInstanceAttributes(t, &got, &expected),
				),
			},
		},
	})
}

// testApplicationInstanceExists retrieves the ApplicationInstance and stores it in the provided *models.DeviceConfig.
func testApplicationInstanceExists(resourceName string, applicationModel *models.AppInstance) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("ApplicationInstance not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ApplicationInstance ID is not set")
		}

		// retrieve the client established in Provider configuration
		client := testProvider.Meta().(*api_client.ZedcloudAPI)

		// retrieve the ApplicationInstance by referencing its state ID for API lookup
		params := config.GetByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.ApplicationInstance.GetByID(params, nil)
		if err != nil {
			return fmt.Errorf("could not fetch ApplicationInstance (%s): %w", rs.Primary.ID, err)
		}

		app := response.GetPayload()
		if app == nil {
			return errors.New("could not get response payload in ApplicationInstance existence test: application_instance is nil")
		}

		*applicationModel = *app
		return nil
	}
}

// testApplicationInstanceAttributes verifies attributes are set correctly by Terraform
func testApplicationInstanceAttributes(t *testing.T, got, expected *models.AppInstance) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ignoredFields := []string{
			"Imvolname",
			"Mvolname",
		}
		opts := cmpopts.IgnoreFields(models.Drive{}, ignoredFields...)
		if diff := cmp.Diff(*got.Drives[0], *expected.Drives[0], opts); len(diff) != 0 {
			return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
		}

		ignoredFields = []string{
			"ID",
			"Revision",
			"DeviceID",
			"AppID",
			"Interfaces",
			"Drives",
		}
		opts = cmpopts.IgnoreFields(models.AppInstance{}, ignoredFields...)
		if diff := cmp.Diff(*got, *expected, opts); len(diff) != 0 {
			return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
		}
		return nil
	}
}

// testApplicationInstanceInstanceDestroy verifies the ApplicationInstance has been destroyed.
func testApplicationInstanceDestroy(s *terraform.State) error {
	// retrieve the client established in Provider configuration
	client := testProvider.Meta().(*api_client.ZedcloudAPI)

	// loop through the resources in state, verifying each ApplicationInstance is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zedcloud_application_instance" {
			continue
		}

		// retrieve the ApplicationInstance by referencing it's state ID for API lookup
		params := config.GetByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.ApplicationInstance.GetByID(params, nil)
		if err == nil {
			if application := response.GetPayload(); application != nil && application.ID == rs.Primary.ID {
				return fmt.Errorf("destroy failed, ApplicationInstance (%s) still exists", application.ID)
			}
			return nil
		}

		// if the error is equivalent to 404 not found, the ApplicationInstance is destroyed
		_, ok := err.(*config.GetApplicationInstanceNotFound)
		if !ok {
			return fmt.Errorf("destroy failed, expect status code 404 for ApplicationInstance (%s)", params.ID)
		}
	}
	return nil
}
