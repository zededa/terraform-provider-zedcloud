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
	config "github.com/zededa/terraform-provider/client/projects"
	"github.com/zededa/terraform-provider/models"
)

func TestProject_Create_RequiredOnly(t *testing.T) {
	var gotCreated, expectCreated models.Tag
	// var gotUpdated, expectUpdated models.Tag

	// input configs
	createPath := "project/create_required_only.tf"
	inputCreate := mustGetTestInput(t, createPath)
	// updatePath := "project/update.tf"
	// inputUpdate := mustGetTestInput(t, updatePath)

	// expected output
	createPath = "project/create_required_only.yaml"
	mustGetExpectedOutput(t, createPath, &expectCreated)
	// updatePath = "project/update.yaml"
	// mustGetExpectedOutput(t, updatePath, &expectUpdated)

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { checkEnv(t) },
		CheckDestroy: testProjectDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: inputCreate,
				Check: resource.ComposeTestCheckFunc(
					testProjectExists("zedcloud_project.test_tf_provider", &gotCreated),
					resource.TestCheckResourceAttr("zedcloud_project.test_tf_provider", "name", "test_tf_provider"),
					resource.TestCheckResourceAttr("zedcloud_project.test_tf_provider", "title", "title"),
					resource.TestMatchResourceAttr(
						"zedcloud_project.test_tf_provider",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testProjectAttributes(t, &gotCreated, &expectCreated),
				),
			},
			// {
			// 	Config: inputUpdate,
			// 	Check: resource.ComposeTestCheckFunc(
			// 		testProjectExists("zedcloud_project.test_tf_provider", &gotUpdated),
			// 		resource.TestCheckResourceAttr("zedcloud_project.test_tf_provider", "name", "test_tf_provider"),
			// 		resource.TestCheckResourceAttr("zedcloud_project.test_tf_provider", "model_id", "2f716b55-2639-486c-9a2f-55a2e94146a6"),
			// 		resource.TestCheckResourceAttr("zedcloud_project.test_tf_provider", "title", "test_tf_provider-title"),
			// 		resource.TestMatchResourceAttr(
			// 			"zedcloud_project.test_tf_provider",
			// 			"id",
			// 			regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
			// 		),
			// 		testProjectAttributes(t, &gotUpdated, &expectUpdated),
			// 	),
			// },
		},
	})
}

// testProjectExists retrieves the Project and stores it in the provided *models.DeviceConfig.
func testProjectExists(resourceName string, projectModel *models.Tag) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Project not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Project ID is not set")
		}

		// retrieve the client established in Provider configuration
		client := testProvider.Meta().(*api_client.ZedcloudAPI)

		// retrieve the Project by referencing its state ID for API lookup
		params := config.GetByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.Project.GetByID(params, nil)
		if err != nil {
			return fmt.Errorf("could not fetch Project (%s): %w", rs.Primary.ID, err)
		}

		project := response.GetPayload()
		if project == nil {
			return errors.New("could not get response payload in Project existence test: project is nil")
		}

		*projectModel = *project
		return nil
	}
}

// testProjectAttributes verifies attributes are set correctly by Terraform
func testProjectAttributes(t *testing.T, got, expected *models.Tag) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		opts := cmpopts.IgnoreFields(
			models.Tag{},
			"ID",
			"AttestationPolicy",
			"Revision",
		)
		// API and YAML unmarshal might change order of list elements so we need to ignore order when comparing
		// if !schemas.CompareSystemInterfaceList(got.Interfaces, expected.Interfaces) {
		// 	return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), cmp.Diff(got.Interfaces, expected.Interfaces))
		// }
		if diff := cmp.Diff(*got, *expected, opts); len(diff) != 0 {
			return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
		}
		return nil
	}
}

// testProjectDestroy verifies the Project has been destroyed.
func testProjectDestroy(s *terraform.State) error {
	// retrieve the client established in Provider configuration
	client := testProvider.Meta().(*api_client.ZedcloudAPI)

	// loop through the resources in state, verifying each Project is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zedcloud_project" {
			continue
		}

		// retrieve the Project by referencing it's state ID for API lookup
		params := config.GetByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.Project.GetByID(params, nil)
		if err == nil {
			if deviceConfig := response.GetPayload(); deviceConfig != nil && deviceConfig.ID == rs.Primary.ID {
				return fmt.Errorf("destroy failed, Project (%s) still exists", deviceConfig.ID)
			}
			return nil
		}

		// if the error is equivalent to 404 not found, the Project is destroyed
		_, ok := err.(*config.ProjectsGetByIDNotFound)
		if !ok {
			return fmt.Errorf("destroy failed, expect status code 404 for Project (%s)", params.ID)
		}
	}
	return nil
}
