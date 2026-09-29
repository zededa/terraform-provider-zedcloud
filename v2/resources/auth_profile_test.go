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

// authProfileSecretSentinel is the client_secret written by the fixture. It is
// not a credential -- it is a marker the audit below looks for in what the API
// hands back.
const authProfileSecretSentinel = "test-tf-provider-not-a-real-secret"

// TestAuthProfile_CreateAndImport covers UE-168.
//
// The blocking question on that ticket was whether
// GET /v1/authorization/profiles/id/{id} returns clientSecret, cryptoKey and
// encryptedSecrets in clear, since import writes whatever comes back into
// plaintext Terraform state. It does not: fillAuthProfile in zedcloud
// srvs/indusv2/authprofileproc.go blanks all three on every read, for both the
// by-id and the query paths. testAuthProfileSecretsRedacted pins that, so a
// controller change that starts echoing secrets fails here rather than
// silently leaking them into every user's state file.
func TestAuthProfile_CreateAndImport(t *testing.T) {
	var got models.AuthorizationProfile

	input := testhelper.MustGetTestInput(t, "iam/auth_profile.create.tf")
	name := suffixed("test_tf_provider-authprofile")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testAuthProfileDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			// Create
			{
				Config: input,
				Check: resource.ComposeTestCheckFunc(
					testAuthProfileExists("zedcloud_auth_profile.test_tf_provider", &got),
					testAuthProfileSecretsRedacted(&got),
					resource.TestCheckResourceAttr("zedcloud_auth_profile.test_tf_provider", "name", name),
					resource.TestCheckResourceAttr("zedcloud_auth_profile.test_tf_provider", "type", "AUTH_TYPE_OAUTH"),
					resource.TestCheckResourceAttr("zedcloud_auth_profile.test_tf_provider", "active", "false"),
					resource.TestCheckResourceAttr("zedcloud_auth_profile.test_tf_provider", "oauth_profile.0.client_id", "test-tf-provider-client-id"),
					resource.TestMatchResourceAttr(
						"zedcloud_auth_profile.test_tf_provider",
						"id",
						regexp.MustCompile("^[0-9A-Za-z_=-]{28}$"),
					),
				),
			},
			// Import by system ID.
			//
			// client_secret is ignored because it is write-only: the operator
			// supplies it, the controller encrypts it and never gives it back,
			// so the imported state legitimately cannot reproduce the value the
			// create step held. That asymmetry is the point of UE-168, not a
			// round-trip defect -- every other attribute must still match.
			{
				ResourceName:            "zedcloud_auth_profile.test_tf_provider",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"oauth_profile.0.client_secret"},
			},
			// Import by profile name (UE-166 §4.1).
			{
				ResourceName:            "zedcloud_auth_profile.test_tf_provider",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"oauth_profile.0.client_secret"},
				ImportStateIdFunc:       importIDFromAttribute("zedcloud_auth_profile.test_tf_provider", "name"),
			},
		},
	})
}

// testAuthProfileSecretsRedacted is the UE-168 audit, run as an assertion.
//
// It never logs a secret value, only whether one came back.
func testAuthProfileSecretsRedacted(profile *models.AuthorizationProfile) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		oauth := profile.OauthProfile
		if oauth == nil {
			return errors.New("auth profile read returned no oauth_profile; cannot audit secret exposure")
		}

		if oauth.ClientSecret != "" {
			if oauth.ClientSecret == authProfileSecretSentinel {
				return errors.New("SECURITY: the API echoed client_secret back verbatim; " +
					"importing an auth profile would write it into plaintext state")
			}
			return errors.New("SECURITY: the API returned a non-empty client_secret")
		}
		if oauth.CryptoKey != "" {
			return errors.New("SECURITY: the API returned a non-empty crypto_key")
		}
		if len(oauth.EncryptedSecrets) != 0 {
			return fmt.Errorf("SECURITY: the API returned %d encrypted_secrets entries",
				len(oauth.EncryptedSecrets))
		}

		return nil
	}
}

// testAuthProfileExists retrieves the profile by its state ID and stores it.
func testAuthProfileExists(resourceName string, profileModel *models.AuthorizationProfile) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("auth profile not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("auth profile ID is not set")
		}

		client := testProvider.Meta().(*api_client.ZedcloudAPI)

		params := config.NewIdentityAccessManagementGetAuthProfileParams()
		params.ID = rs.Primary.ID
		response, err := client.IdentityAccessManagement.IdentityAccessManagementGetAuthProfile(params, nil)
		if err != nil {
			return fmt.Errorf("could not fetch auth profile (%s): %w", rs.Primary.ID, err)
		}

		profile := response.GetPayload()
		if profile == nil {
			return errors.New("could not get response payload in auth profile existence test: profile is nil")
		}

		*profileModel = *profile
		return nil
	}
}

// testAuthProfileDestroy verifies every auth profile in state has been destroyed.
func testAuthProfileDestroy(s *terraform.State) error {
	client := testProvider.Meta().(*api_client.ZedcloudAPI)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zedcloud_auth_profile" {
			continue
		}

		params := config.NewIdentityAccessManagementGetAuthProfileParams()
		params.ID = rs.Primary.ID

		if _, err := client.IdentityAccessManagement.IdentityAccessManagementGetAuthProfile(params, nil); err == nil {
			return fmt.Errorf("auth profile %s still exists", rs.Primary.ID)
		}
	}

	return nil
}
