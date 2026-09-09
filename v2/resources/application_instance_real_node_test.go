// Copyright (c) Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package resources

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	testhelper "github.com/zededa/terraform-provider-zedcloud/v2/testing"
)

var uuidRE = regexp.MustCompile(
	`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)

// TestApplicationInstance_RealNode deploys a real container onto a real,
// onboarded edge node and waits until the device reports it running.
//
// This is the test the rest of the suite cannot be: every other acceptance test
// fabricates a zedcloud_edgenode against a synthetic brand/model with an invented
// serial, so nothing ever checks in and an app instance is "created" without ever
// running. Here the controller must actually push the config, the device must pull
// the image from Docker Hub and start the workload -- so the assertion is
// SW_STATE_RUNNING, not merely "the POST returned 200".
//
// It is the intended replacement for the reason TestApplicationInstance_CreateCompose
// carries the comment "skipped until we decide how to provide a real device for
// testing".
//
// Requires (otherwise skipped):
//
//	ZEDCLOUD_ACC_REAL_NODE=1
//	ZEDCLOUD_TEST_NODE_NAME=<name of an onboarded edge node>
//	ZEDCLOUD_TEST_SUFFIX=<per-run suffix>      (generated if absent)
//	TF_VAR_zedcloud_url / TF_VAR_zedcloud_token
//
// See test/e2e/ for the Terraform config that creates a suitable node, and
// docs/design/e2e-virtual-eve-node-testing.md for how this fits into CI.
func TestApplicationInstance_RealNode(t *testing.T) {
	node := testhelper.RealNode(t)

	const resourceName = "zedcloud_application_instance.real"

	input := testhelper.MustGetTestInputWithVars(t,
		"application_instance/create_real_node.tf",
		map[string]string{
			"NODE_ID": node.ID,
			"SUFFIX":  node.Suffix,
		},
	)

	var got models.AppInstance

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testApplicationInstanceDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: input,
				Check: resource.ComposeTestCheckFunc(
					testApplicationInstanceExists(resourceName, &got),
					resource.TestMatchResourceAttr(resourceName, "id", uuidRE),
					resource.TestMatchResourceAttr(resourceName, "app_id", uuidRE),

					// The instance must be bound to the live node, not to some
					// node a fixture invented for itself.
					resource.TestCheckResourceAttr(resourceName, "device_id", node.ID),
					resource.TestCheckResourceAttr(
						"zedcloud_network_instance.real", "device_id", node.ID),

					// The payload: the workload actually started on the device.
					// Budget covers config push, image pull from Docker Hub and
					// container start.
					waitAppInstanceRunning(t, resourceName, 15*time.Minute),
				),
			},
		},
	})
}

// waitAppInstanceRunning blocks until the device reports the workload running.
func waitAppInstanceRunning(t *testing.T, appInstName string, timeout time.Duration) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[appInstName]
		if !ok {
			return fmt.Errorf("%s not found in state", appInstName)
		}
		id := rs.Primary.ID
		if id == "" {
			return fmt.Errorf("%s has an empty id", appInstName)
		}

		return testhelper.WaitForAppInstanceSwState(
			t, id, models.SWStateSWSTATERUNNING, timeout)
	}
}
