package resources

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	asset_group_service "github.com/zededa/terraform-provider-zedcloud/v2/client/asset_group_service"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	testhelper "github.com/zededa/terraform-provider-zedcloud/v2/testing"
)

func TestAssetGroup_Create(t *testing.T) {
	var got, expected models.AssetGroupRead

	// input config
	createPath := "asset_groups/create.tf"
	input := testhelper.MustGetTestInput(t, createPath)

	// expected output
	createPath = "asset_groups/create.yaml"
	testhelper.MustGetExpectedOutput(t, createPath, &expected)

	// terraform acceptance testcase
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testhelper.CheckEnv(t)
		},
		CheckDestroy: testAssetGroupDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:             input,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testAssetGroupExists("zedcloud_asset_group.test_tf_provider", &got),
					resource.TestMatchResourceAttr(
						"zedcloud_asset_group.test_tf_provider",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					resource.TestMatchResourceAttr(
						"zedcloud_asset_group.test_tf_provider",
						"project_id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testAssetGroupAttributes(t, &got, &expected),
				),
			},
		},
	})
}

// testAssetGroupExists retrieves the Asset Group and stores it in the provided *models.AssetGroup.
func testAssetGroupExists(resourceName string, assetGroupModel *models.AssetGroupRead) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("AssetGroup not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("AssetGroup ID is not set")
		}

		projectID := rs.Primary.Attributes["project_id"]
		if projectID == "" {
			return fmt.Errorf("Project ID is not set")
		}

		// retrieve the client established in Provider configuration
		client := testProvider.Meta().(*api_client.ZedcloudAPI)

		// retrieve the AssetGroup by referencing its state ID for API lookup
		params := asset_group_service.NewAssetGroupServiceGetAssetGroupParams()
		params.ID = rs.Primary.ID
		response, err := client.AssetGroup.AssetGroupServiceGetAssetGroup(params, nil)
		if err != nil {
			return fmt.Errorf("could not fetch Assetgroup (%s): %w", rs.Primary.ID, err)
		}

		assetGroupResp := response.GetPayload()
		if assetGroupResp == nil {
			return errors.New("could not get response payload in AssetGroup existence test: assetgroup is nil")
		}

		*assetGroupModel = *assetGroupResp
		return nil
	}
}

// testAssetGroupAttributes verifies attributes are set correctly by Terraform
func testAssetGroupAttributes(t *testing.T, got, expected *models.AssetGroupRead) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		gotAssetIds := got.AssetIds
		expectedAssetIds := expected.AssetIds
		if gotAssetIds != nil && expectedAssetIds != nil {
			if len(gotAssetIds.Ids) != len(expectedAssetIds.Ids) {
				return fmt.Errorf("unexpected number of assets: got %d, expected %d", len(gotAssetIds.Ids), len(expectedAssetIds.Ids))
			}
		}

		gotAssetTags := got.AssetTags
		expectedAssetTags := expected.AssetTags
		if gotAssetTags != nil && expectedAssetTags != nil {
			if len(gotAssetTags.AssetTag) != len(expectedAssetTags.AssetTag) {
				return fmt.Errorf("unexpected number of asset tags: got %d, expected %d", len(gotAssetTags.AssetTag), len(expectedAssetTags.AssetTag))
			}
		}

		ignoredFields := []string{
			"ID",
			"ProjectID",
			"AssetIds",
		}
		opts := cmpopts.IgnoreFields(models.AssetGroupRead{}, ignoredFields...)
		if diff := cmp.Diff(*got, *expected, opts); len(diff) != 0 {
			return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
		}
		return nil
	}
}

// testAssetGroupDestroy verifies the AssetGroup has been destroyed.
func testAssetGroupDestroy(s *terraform.State) error {
	// retrieve the client established in Provider configuration
	client := testProvider.Meta().(*api_client.ZedcloudAPI)

	// loop through the resources in state, verifying each assetgroup is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zedcloud_asset_group" {
			continue
		}

		// retrieve the Asset group by referencing it's state ID for API lookup
		params := asset_group_service.NewAssetGroupServiceGetAssetGroupParams()
		params.ID = rs.Primary.ID
		response, err := client.AssetGroup.AssetGroupServiceGetAssetGroup(params, nil)
		if err == nil {
			if assetGroupResp := response.GetPayload(); assetGroupResp != nil && assetGroupResp.ID == rs.Primary.ID {
				return fmt.Errorf("destroy failed, Asset Group (%s) still exists", assetGroupResp.ID)
			}
			return nil
		}

		// if we use an http client with retries,
		// it overrrides AssetGroupServiceGetAssetGroupNotFound error
		if strings.Contains(err.Error(), "unexpected HTTP status 404 Not Found") {
			return nil
		}

		// if the error is equivalent to 404 not found, the Asset group is destroyed
		_, ok := err.(*asset_group_service.AssetGroupServiceGetAssetGroupNotFound)
		if !ok {
			return fmt.Errorf("destroy failed, expect status code 404 for Asset Group (%s)", params.ID)
		}
	}
	return nil
}
