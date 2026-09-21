// Copyright (c) Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package testing

import (
	"fmt"
)

// Environment contract for the CI-709 real-node test. Unlike EnvRealNodeName,
// these name OBJECTS ON a cluster rather than the cluster's nodes, because
// CI-709 is about instances, not devices.
//
// test/e2e-cluster/ci709_repro.tf creates all three when
// TF_VAR_repro_ci709=true; its outputs give the names.
const (
	// EnvClusterNIName names a cluster-scoped zedcloud_network_instance.
	EnvClusterNIName = "ZEDCLOUD_TEST_CLUSTER_NI_NAME"
	// EnvClusterVolName names a cluster-scoped zedcloud_volume_instance.
	EnvClusterVolName = "ZEDCLOUD_TEST_CLUSTER_VOL_NAME"
	// EnvClusterAppInstName names a cluster-scoped zedcloud_application_instance.
	EnvClusterAppInstName = "ZEDCLOUD_TEST_CLUSTER_APPINST_NAME"
)

// ClusterScopedInstance holds the fields the controller assigns to an instance
// that was created with only an `edge_node_cluster` and no `device_id`.
//
// These six fields across three resources are exactly what PR #234 changed to
// `Optional + Computed`, and `app_type` additionally lost its
// `Default: "APP_TYPE_UNSPECIFIED"` there.
type ClusterScopedInstance struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	DeviceID  string `json:"deviceId"`
	ClusterID string `json:"clusterId"`
	// AppType is populated for application instances only; volume and network
	// instances have no such field.
	AppType         string `json:"appType"`
	EdgeNodeCluster *struct {
		ID string `json:"id"`
	} `json:"edgeNodeCluster"`
}

// ClusterScoped reports whether the controller resolved a designated node for
// this instance, which is the precondition for the CI-709 assertion to mean
// anything. An instance with an empty DeviceID was never assigned one, so a
// clean plan would prove nothing.
func (c ClusterScopedInstance) ClusterScoped() bool {
	return c.DeviceID != "" && c.EdgeNodeCluster != nil && c.EdgeNodeCluster.ID != ""
}

// InstanceKind selects which API collection to look a name up in.
type InstanceKind string

const (
	KindNetworkInstance     InstanceKind = "netinsts"
	KindVolumeInstance      InstanceKind = "volumes/instances"
	KindApplicationInstance InstanceKind = "apps/instances"
)

// GetClusterScopedInstance reads an instance by name and returns the
// server-assigned fields.
//
// By NAME rather than by id so the caller can pass a
// test/e2e-cluster output straight through without plumbing UUIDs, and so the
// test stays readable in CI logs. All three collections expose
// /v1/<kind>/name/{name}.
func GetClusterScopedInstance(kind InstanceKind, name string) (ClusterScopedInstance, error) {
	var out ClusterScopedInstance
	if name == "" {
		return out, fmt.Errorf("empty instance name for kind %s", kind)
	}
	if err := apiGet(fmt.Sprintf("%s/name/%s", kind, name), &out); err != nil {
		return out, fmt.Errorf("reading %s %q: %w", kind, name, err)
	}
	return out, nil
}
