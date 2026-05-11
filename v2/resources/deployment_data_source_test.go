package resources

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	testhelper "github.com/zededa/terraform-provider-zedcloud/v2/testing"
)

func TestDeployment_DataSource(t *testing.T) {
	createCfg := testhelper.MustGetTestInput(t, "deployment/create.tf")
	dsCfg := testhelper.MustGetTestInput(t, "deployment/datasource.tf")
	combined := createCfg + "\n" + dsCfg

	uuidPattern := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testDeploymentDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:             combined,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.zedcloud_deployment.by_name",
						"name",
						"test_tf_provider-deployment",
					),
					resource.TestCheckResourceAttr(
						"data.zedcloud_deployment.by_name",
						"deployment_tag",
						"depl:1234",
					),
					resource.TestMatchResourceAttr(
						"data.zedcloud_deployment.by_name",
						"id",
						uuidPattern,
					),
					resource.TestCheckResourceAttrPair(
						"data.zedcloud_deployment.by_name",
						"id",
						"zedcloud_deployment.tf_deployment",
						"id",
					),
					resource.TestCheckResourceAttrPair(
						"data.zedcloud_deployment.by_id",
						"name",
						"zedcloud_deployment.tf_deployment",
						"name",
					),
					resource.TestCheckResourceAttrPair(
						"data.zedcloud_deployment.by_id",
						"deployment_tag",
						"zedcloud_deployment.tf_deployment",
						"deployment_tag",
					),
				),
			},
		},
	})
}
