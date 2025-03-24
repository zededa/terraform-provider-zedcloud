package resources

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	"github.com/zededa/terraform-provider-zedcloud/v2/client/profile_deployment_service"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	testhelper "github.com/zededa/terraform-provider-zedcloud/v2/testing"
)

func TestProfileDeployment_Create(t *testing.T) {
	var got, expected models.ProfileDeployment

	// input config
	createPath := "profile_deployments/create.tf"
	input := testhelper.MustGetTestInput(t, createPath)

	// expected output
	createPath = "profile_deployments/create.yaml"
	testhelper.MustGetExpectedOutput(t, createPath, &expected)

	// terraform acceptance testcase
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testhelper.CheckEnv(t)
		},
		CheckDestroy: testProfileDeploymentDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:             input,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testProfileDeploymentExists("zedcloud_profile_deployment.test_tf_provider", &got),
					// resource.TestMatchResourceAttr(
					// 	"zedcloud_profile_deployment.test_tf_provider",
					// 	"id",
					// 	regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					// ),
					// resource.TestMatchResourceAttr(
					// 	"zedcloud_profile_deployment.test_tf_provider",
					// 	"project_id",
					// 	regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					// ),
					testProfileDeploymentAttributes(t, &got, &expected),
				),
			},
		},
	})
}

// testProfileDeploymentExists retrieves the Profile Deployment and stores it in the provided *models.ProfileDeployment.
func testProfileDeploymentExists(name string, p *models.ProfileDeployment) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return nil
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ProfileDeployment ID is not set")
		}

		projectID := rs.Primary.Attributes["project_id"]
		if projectID == "" {
			return fmt.Errorf("ProfileDeployment project_id is not set")
		}

		// retrieve the client established in Provider configuration
		client := testProvider.Meta().(*api_client.ZedcloudAPI)

		// retreive the ProfileDeployment by ID
		params := profile_deployment_service.NewProfileDeploymentServiceGetProfileDeploymentParams()
		params.ID = rs.Primary.ID
		resp, err := client.ProfileDeployment.ProfileDeploymentServiceGetProfileDeployment(params, nil)
		if err != nil {
			return fmt.Errorf("could not retrieve ProfileDeployment(%s): %w", rs.Primary.ID, err)
		}

		profileDeployment := resp.GetPayload()
		if profileDeployment == nil {
			return fmt.Errorf("could not get response payload for ProfileDeployment(%s)", rs.Primary.ID)
		}

		*p = *profileDeployment
		return nil
	}
}

// testProfileDeploymentAttributes checks the ProfileDeployment attributes match the expected attributes.
func testProfileDeploymentAttributes(t *testing.T, got, expected *models.ProfileDeployment) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ignoredFields := []string{
			"ID",
			"ProjectID",
			"TargetAssetGroup",
			"AppProfileInfo",
		}
		opts := cmpopts.IgnoreFields(models.ProfileDeployment{}, ignoredFields...)
		if diff := cmp.Diff(*got, *expected, opts); len(diff) != 0 {
			return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
		}
		return nil
	}
}

// testProfileDeploymentDestroy verifies the ProfileDeployment is destroyed.
func testProfileDeploymentDestroy(s *terraform.State) error {
	// retrieve the client established in Provider configuration
	client := testProvider.Meta().(*api_client.ZedcloudAPI)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zedcloud_profile_deployment" {
			continue
		}

		// retreive the ProfileDeployment by ID
		params := profile_deployment_service.NewProfileDeploymentServiceGetProfileDeploymentParams()
		params.ID = rs.Primary.ID
		response, err := client.ProfileDeployment.ProfileDeploymentServiceGetProfileDeployment(params, nil)
		if err == nil {
			if profileDeploymentResp := response.GetPayload(); profileDeploymentResp != nil && profileDeploymentResp.ID == rs.Primary.ID {
				return fmt.Errorf("destroy failed, ProfileDeployment(%s) still exists", rs.Primary.ID)
			}
			return nil
		}

		// if the error is equivalent to 404, the ProfileDeployment has been destroyed
		_, ok := err.(*profile_deployment_service.ProfileDeploymentServiceGetProfileDeploymentNotFound)
		if !ok {
			return fmt.Errorf("destroy failed, expect status code 404 for ProfileDeployment(%s)", params.ID)
		}
	}
	return nil
}
