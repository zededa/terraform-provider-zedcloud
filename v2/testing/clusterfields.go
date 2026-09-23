// Copyright (c) Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package testing

import (
	"fmt"
)

// NodeClusterFields is the subset of an edge node's CONFIG object that the
// controller populates by itself once the node joins an edge-node cluster.
//
// These are the two attributes CI-834 is about. The provider declares neither
// of them in any configuration, so before c932b75d -- which flipped both from
// Optional to Computed in v2/schemas/node.go -- Terraform read the values back,
// saw no matching configuration, and planned to REMOVE them, which tears down
// the cluster membership.
//
// Deliberately hand-rolled rather than reusing models.Node: the JSON tags the
// generated model carries are the ones the generated client needs, and this
// exists precisely to observe the wire format independently of the code under
// test. If the provider stopped reading these fields correctly, a helper built
// on the same models would be just as wrong.
type NodeClusterFields struct {
	ClusterInterface string `json:"clusterInterface"`
	EdgeNodeCluster  *struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		ClusterPrefix string `json:"clusterPrefix"`
		SeedNodeID    string `json:"seedNodeId"`
		SeedNodeIP    string `json:"seedNodeIp"`
		IsMaster      bool   `json:"isMaster"`
		ProjectID     string `json:"projectId"`
		// Token is a shared cluster secret. Read so the struct matches the
		// wire format, never logged.
		Token string `json:"token"`
	} `json:"edgeNodeCluster"`
}

// Clustered reports whether the controller has populated both cluster fields.
func (n NodeClusterFields) Clustered() bool {
	return n.ClusterInterface != "" && n.EdgeNodeCluster != nil && n.EdgeNodeCluster.ID != ""
}

// GetNodeClusterFields reads the controller-populated cluster attributes off a
// live edge node.
//
// Note this is the device CONFIG endpoint, not /status: clusterInterface and
// edgeNodeCluster live on the config object, which is what the provider reads
// and therefore what produces a diff.
func GetNodeClusterFields(nodeID string) (NodeClusterFields, error) {
	var out NodeClusterFields
	if nodeID == "" {
		return out, fmt.Errorf("empty node id")
	}
	if err := apiGet(fmt.Sprintf("devices/id/%s", nodeID), &out); err != nil {
		return out, fmt.Errorf("reading cluster fields for node %s: %w", nodeID, err)
	}
	return out, nil
}
