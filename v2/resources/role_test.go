package resources

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	config "github.com/zededa/terraform-provider-zedcloud/v2/client/identity_access_management"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	testhelper "github.com/zededa/terraform-provider-zedcloud/v2/testing"
)

// TestRole_CreateAndImport covers NFR-165 for custom roles: create one, then
// bring the same object back under management by ID and by name.
//
// ImportStateVerify compares the imported state against the state the create
// step produced, so it fails on any attribute the Read path cannot reproduce
// from the API alone -- which is exactly the failure mode that made roles
// non-importable in the first place.
func TestRole_CreateAndImport(t *testing.T) {
	var got models.Role

	input := testhelper.MustGetTestInput(t, "iam/role.create.tf")
	name := suffixed("test_tf_provider-import-role")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testRoleDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			// Create
			{
				Config: input,
				Check: resource.ComposeTestCheckFunc(
					testRoleExists("zedcloud_role.test_tf_provider", &got),
					resource.TestCheckResourceAttr("zedcloud_role.test_tf_provider", "name", name),
					resource.TestCheckResourceAttr("zedcloud_role.test_tf_provider", "title", name),
					resource.TestCheckResourceAttr("zedcloud_role.test_tf_provider", "type", "USER_ROLE_USER_DEFINED"),
					resource.TestCheckResourceAttr("zedcloud_role.test_tf_provider", "scopes.#", "1"),
					resource.TestCheckResourceAttr("zedcloud_role.test_tf_provider", "scopes.0.access_device", "PermissionAccessRead"),
					resource.TestCheckResourceAttr("zedcloud_role.test_tf_provider", "scopes.0.access_user", "PermissionAccessNone"),
					resource.TestMatchResourceAttr(
						"zedcloud_role.test_tf_provider",
						"id",
						regexp.MustCompile("^[0-9A-Za-z_=-]{28}$"),
					),
				),
			},
			// Import by system ID.
			{
				ResourceName:      "zedcloud_role.test_tf_provider",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Import by role name (NFR-165 §4.1).
			{
				ResourceName:      "zedcloud_role.test_tf_provider",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: importIDFromAttribute("zedcloud_role.test_tf_provider", "name"),
			},
		},
	})
}

// testRoleExists retrieves the Role by its state ID and stores it.
func testRoleExists(resourceName string, roleModel *models.Role) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Role not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Role ID is not set")
		}

		client := testProvider.Meta().(*api_client.ZedcloudAPI)

		params := config.NewIdentityAccessManagementGetRoleParams()
		params.ID = rs.Primary.ID
		response, err := client.IdentityAccessManagement.IdentityAccessManagementGetRole(params, nil)
		if err != nil {
			return fmt.Errorf("could not fetch Role (%s): %w", rs.Primary.ID, err)
		}

		role := response.GetPayload()
		if role == nil {
			return errors.New("could not get response payload in Role existence test: role is nil")
		}

		*roleModel = *role
		return nil
	}
}

// testRoleDestroy verifies every Role in state has been destroyed.
func testRoleDestroy(s *terraform.State) error {
	client := testProvider.Meta().(*api_client.ZedcloudAPI)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zedcloud_role" {
			continue
		}

		params := config.NewIdentityAccessManagementGetRoleParams()
		params.ID = rs.Primary.ID

		if _, err := client.IdentityAccessManagement.IdentityAccessManagementGetRole(params, nil); err == nil {
			return fmt.Errorf("role %s still exists", rs.Primary.ID)
		}
	}

	return nil
}
