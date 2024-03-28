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
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	config "github.com/zededa/terraform-provider-zedcloud/v2/client/application"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	testhelper "github.com/zededa/terraform-provider-zedcloud/v2/testing"
)

func TestApplication_Create(t *testing.T) {
	var got models.Application
	var expected models.Application

	// input config
	inputPath := "application/create.tf"
	input := testhelper.MustGetTestInput(t, inputPath)

	// expected output
	expectedPath := "application/create.yaml"
	testhelper.MustGetExpectedOutput(t, expectedPath, &expected)

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testApplicationDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: input,
				Check: resource.ComposeTestCheckFunc(
					testApplicationExists("zedcloud_application.test_tf_provider", &got),
					resource.TestMatchResourceAttr(
						"zedcloud_application.test_tf_provider",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testApplicationAttributes(t, &got, &expected),
				),
			},
		},
	})
}

func TestApplication_Create_FromFile(t *testing.T) {
	var got models.Application
	var expected models.Application

	// input config
	inputPath := "application/create_from_file.tf"
	input := testhelper.MustGetTestInput(t, inputPath)

	// expected output
	expectedPath := "application/create_from_file.yaml"
	testhelper.MustGetExpectedOutput(t, expectedPath, &expected)

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testApplicationDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: input,
				Check: resource.ComposeTestCheckFunc(
					testApplicationExists("zedcloud_application.test_tf_provider", &got),
					resource.TestMatchResourceAttr(
						"zedcloud_application.test_tf_provider",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testApplicationAttributes(t, &got, &expected),
				),
			},
		},
	})
}

func testApplicationExists(resourceName string, applicationModel *models.Application) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Application not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Application ID is not set")
		}

		// retrieve the client established in Provider configuration
		client := testProvider.Meta().(*api_client.ZedcloudAPI)

		// retrieve the Application by referencing its state ID for API lookup
		params := config.GetByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.Application.GetByID(params, nil)
		if err != nil {
			return fmt.Errorf("could not fetch Application (%s): %w", rs.Primary.ID, err)
		}

		application := response.GetPayload()
		if application == nil {
			return errors.New("could not get response payload in Application existence test: application is nil")
		}

		*applicationModel = *application
		return nil
	}
}

// testApplicationAttributes verifies attributes are set correctly by Terraform
func testApplicationAttributes(t *testing.T, got, expected *models.Application) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ignoredFields := []string{
			"ID",
			"Revision",
		}
		opts := cmpopts.IgnoreFields(models.Application{}, ignoredFields...)
		if diff := cmp.Diff(*got, *expected, opts); len(diff) != 0 {
			return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
		}
		return nil
	}
}

// testApplicationInstanceDestroy verifies the Application has been destroyed.
func testApplicationDestroy(s *terraform.State) error {
	// retrieve the client established in Provider configuration
	client := testProvider.Meta().(*api_client.ZedcloudAPI)

	// loop through the resources in state, verifying each Application is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zedcloud_application" {
			continue
		}

		// retrieve the Application by referencing it's state ID for API lookup
		params := config.GetByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.Application.GetByID(params, nil)
		if err == nil {
			if application := response.GetPayload(); application != nil && application.ID == rs.Primary.ID {
				return fmt.Errorf("destroy failed, Application (%s) still exists", application.ID)
			}
			return nil
		}

		// if the error is equivalent to 404 not found, the Application is destroyed
		_, ok := err.(*config.GetApplicationNotFound)
		if !ok {
			return fmt.Errorf("destroy failed, expect status code 404 for Application (%s)", params.ID)
		}
	}
	return nil
}
