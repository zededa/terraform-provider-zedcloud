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
	config "github.com/zededa/terraform-provider-zedcloud/v2/client/volume_instance"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	testhelper "github.com/zededa/terraform-provider-zedcloud/v2/testing"
)

func TestVolumeInstance_Create(t *testing.T) {
	var got models.VolumeInstance
	var expected models.VolumeInstance

	// input config
	inputPath := "volume_instance/create.tf"
	input := testhelper.MustGetTestInput(t, inputPath)

	// expected output
	expectedPath := "volume_instance/create.yaml"
	testhelper.MustGetExpectedOutput(t, expectedPath, &expected)

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testVolumeInstanceDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: input,
				Check: resource.ComposeTestCheckFunc(
					testVolumeInstanceExists("zedcloud_volume_instance.test_volume_instance", &got),
					resource.TestMatchResourceAttr(
						"zedcloud_volume_instance.test_volume_instance",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					resource.TestMatchResourceAttr(
						"zedcloud_volume_instance.test_volume_instance",
						"device_id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testVolumeInstanceAttributes(t, &got, &expected),
				),
			},
		},
	})
}

func testVolumeInstanceExists(resourceName string, volume_instanceModel *models.VolumeInstance) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("VolumeInstance not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("VolumeInstance ID is not set")
		}

		// retrieve the client established in Provider configuration
		client := testProvider.Meta().(*api_client.ZedcloudAPI)

		// retrieve the VolumeInstance by referencing its state ID for API lookup
		params := config.GetByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.VolumeInstance.GetByID(params, nil)
		if err != nil {
			return fmt.Errorf("could not fetch VolumeInstance (%s): %w", rs.Primary.ID, err)
		}

		volumeInst := response.GetPayload()
		if volumeInst == nil {
			return errors.New("could not get response payload in VolumeInstance existence test: volumeInst is nil")
		}

		*volume_instanceModel = *volumeInst
		return nil
	}
}

// testVolumeInstanceAttributes verifies attributes are set correctly by Terraform
func testVolumeInstanceAttributes(t *testing.T, got, expected *models.VolumeInstance) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ignoredFields := []string{
			"ID",
			"Revision",
			"ProjectID",
			"DeviceID",
			"Purge",
			"SizeBytes",
		}
		opts := cmpopts.IgnoreFields(models.VolumeInstance{}, ignoredFields...)
		if diff := cmp.Diff(*got, *expected, opts); len(diff) != 0 {
			return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
		}
		return nil
	}
}

// testVolumeInstanceInstanceDestroy verifies the VolumeInstance has been destroyed.
func testVolumeInstanceDestroy(s *terraform.State) error {
	// retrieve the client established in Provider configuration
	client := testProvider.Meta().(*api_client.ZedcloudAPI)

	// loop through the resources in state, verifying each VolumeInstance is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zedcloud_volume_instance" {
			continue
		}

		// retrieve the VolumeInstance by referencing it's state ID for API lookup
		params := config.GetByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.VolumeInstance.GetByID(params, nil)
		if err == nil {
			if volumeInst := response.GetPayload(); volumeInst != nil && volumeInst.ID == rs.Primary.ID {
				return fmt.Errorf("destroy failed, VolumeInstance (%s) still exists", volumeInst.ID)
			}
			return nil
		}

		// if the error is equivalent to 404 not found, the VolumeInstance is destroyed
		_, ok := err.(*config.GetByIDNotFound)
		if !ok {
			return fmt.Errorf("destroy failed, expect status code 404 for VolumeInstance (%s)", params.ID)
		}
	}
	return nil
}
