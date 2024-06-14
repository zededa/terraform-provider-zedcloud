package resources

import (
	"errors"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	config "github.com/zededa/terraform-provider-zedcloud/v2/client/identity_access_management"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	testhelper "github.com/zededa/terraform-provider-zedcloud/v2/testing"
)

func TestUser_Create(t *testing.T) {
	var got models.DetailedUser
	var expected models.DetailedUser

	// input config
	inputPath := "iam/user.create.tf"
	input := testhelper.MustGetTestInput(t, inputPath)

	// expected output
	expectedPath := "iam/user.create.yaml"
	testhelper.MustGetExpectedOutput(t, expectedPath, &expected)

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testUserDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: input,
				Check: resource.ComposeTestCheckFunc(
					testUserExists("zedcloud_user.test_tf_provider", &got),
					resource.TestMatchResourceAttr(
						"zedcloud_user.test_tf_provider",
						"id",
						regexp.MustCompile("^[0-9A-Za-z_=-]{28}$"),
					),
					testUserAttributes(t, &got, &expected),
				),
			},
		},
	})
}

// testUserExists retrieves the User and stores it in the provided *models.DeviceConfig.
func testUserExists(resourceName string, userModel *models.DetailedUser) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("User not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("User ID is not set")
		}

		// retrieve the client established in Provider configuration
		client := testProvider.Meta().(*api_client.ZedcloudAPI)

		// retrieve the ApplicationInstance by referencing its state ID for API lookup
		params := config.NewIdentityAccessManagementGetUserParams()
		params.ID = rs.Primary.ID
		response, err := client.IdentityAccessManagement.IdentityAccessManagementGetUser(params, nil)
		if err != nil {
			return fmt.Errorf("could not fetch User (%s): %w", rs.Primary.ID, err)
		}

		user := response.GetPayload()
		if user == nil {
			return errors.New("could not get response payload in User existence test: user is nil")
		}

		*userModel = *user
		return nil
	}
}

// testUserAttributes verifies attributes are set correctly by Terraform
func testUserAttributes(t *testing.T, got, expected *models.DetailedUser) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ignoredFields := []string{
			"ID",
			"AllowedEnterprises",
			"Revision",
			"RoleID",
			"SfdcID",
			"EnterpriseID",
			"CustomUserInput",
		}
		opts := cmpopts.IgnoreFields(models.DetailedUser{}, ignoredFields...)
		if diff := cmp.Diff(*got, *expected, opts); len(diff) != 0 {
			return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
		}
		return nil
	}
}

// testUserDestroy verifies the User has been destroyed.
func testUserDestroy(s *terraform.State) error {
	// retrieve the client established in Provider configuration
	client := testProvider.Meta().(*api_client.ZedcloudAPI)

	// loop through the resources in state, verifying each User is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zedcloud_user" {
			continue
		}

		// retrieve the User by referencing its state ID for API lookup
		params := config.NewIdentityAccessManagementGetUserParams()
		params.ID = rs.Primary.ID
		response, err := client.IdentityAccessManagement.IdentityAccessManagementGetUser(params, nil)
		if err == nil {
			if user := response.GetPayload(); user != nil && user.ID == rs.Primary.ID {
				return fmt.Errorf("destroy failed, User (%s) still exists", user.ID)
			}
			return nil
		}

		// if the error is equivalent to 404 not found, the ApplicationInstance is destroyed
		_, ok := err.(*config.IdentityAccessManagementGetUserNotFound)
		if !ok {
			return fmt.Errorf("destroy failed, expect status code 404 for User (%s)", params.ID)
		}
	}
	return nil
}
