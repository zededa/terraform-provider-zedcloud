// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0
//

terraform {
  required_providers {
    zedcloud = {
      source = "zededa/zedcloud"
      //version = "0.0.7-alpha"
    }
  }
}

provider "zedcloud" {
  zedcloud_url = "zedcloud.local.zededa.net"
  zedcloud_token = "c8xBwliOAXyrnTRh46dWDREM0lXZHqsrWWTyDMAM4UQ4gnoLBhXXUS6Pq4NaOuPCJsDBOLJAhhdooeLOipufmRMWeYc1c5oC-iU2B8cRnJRlBYHvxP5-Rf7XjwzdhsczJ6CXvzQU4lVwRvgEEWo_W06kVM5zmLHAExPWXUJ3ncE="
  zservices_url = "zservices.local.zededa.net"
}


resource "zedcloud_zasset_groups" "example" {
  name        = "example-asset-group"
  description = "My example asset group"
  # Add other required fields as per your schema, e.g.:
  asset_ids {
    ids = [ "asset-id-1", "asset-id-2" , "fced5689-5b23-4e6e-8a37-62030dcaa4ac" ]
  }
  # tags        = ["tag1", "tag2"]
}

resource "zedcloud_zasset_groups" "example2" {
  name        = "example-asset-group2"
  description = "My example asset group"
  # Add other required fields as per your schema, e.g.:
  asset_ids {
    ids = [ "asset-id-3", "asset-id-4" ]
  }
  # tags        = ["tag1", "tag2"]
}

data "zedcloud_zasset_groups" "example2" {
  name = zedcloud_zasset_groups.example2.name
  depends_on = [
    zedcloud_zasset_groups.example2
  ]
}

resource "zedcloud_zdeployments" "example_dep" {
  deployment_name = "test1"
  # deployment_id   = "761a9560-bc23-44f6-ba1d-f8d53ff36431"

  chart {
    chart_id        = "07309ceb-256c-410a-8e0e-e9ab2cb979b8"
    chart_name      = "zededa-helm-chart"
    chart_version   = "0.1.1"
    repository_url  = "http://chartmuseum.zservices.svc.cluster.local:8090/"
    values          = <<EOT
{"activate":true,"devId":"fced5689-5b23-4e6e-8a37-62030dcaa4ac","devTags":{"demo-devices":"blr"},"edgeNodeCluster":{"designatedNodeID":"","id":""},"logs":{"access":true},"mainAppInstDescription":"main app instance created using helm charts.","mainAppInstDrives":[],"mainAppInstInterfaces":[{"accessVlanId":0,"acls":[],"defaultNetInstance":false,"directattach":false,"intfname":"eth0","intforder":1,"io":null,"ipaddr":"","macaddr":"","netinstname":"","netinsttag":{"key":"demo-net-inst"}},{"accessVlanId":0,"acls":[],"defaultNetInstance":false,"directattach":false,"intfname":"eth1","intforder":2,"io":null,"ipaddr":"","macaddr":"","netinstname":"","netinsttag":{"key":"demo-net-inst"}}],"mainAppInstNamePrefix":"main-app-inst","mainAppInstTags":{},"mainAppInstTitle":"main-app-inst","mainAppName":"zservice-main-app","netInstDescription":"demo network instance created using helm charts.","netInstKind":"NETWORK_INSTANCE_KIND_LOCAL","netInstNamePrefix":"demo-local-net-inst","netInstPort":"uplink","netInstTags":{"key":"demo-net-inst"},"netInstTitle":"demo-local-net-inst","netInstType":"NETWORK_INSTANCE_DHCP_TYPE_V4","netInstip":{"dhcpRange":{"end":"10.3.0.160","start":"10.3.0.150"},"dns":["10.3.0.1"],"gateway":"10.3.0.1","subnet":"10.3.0.0/16"},"patchEnvelopeAction":"PATCH_ENVELOPE_ACTION_ACTIVATE","patchEnvelopeArtifacts":[{"base64Artifact":{"base64Data":"aGVsbG8gZnJvbSB6aGVsbQ==","base64MetaData":"aGVsbG8gZnJvbSB6aGVsbQ==","fileNameToUse":"summit2024.html"},"format":"OpaqueObjectCategoryInline"}],"patchEnvelopeDescription":"demo patch envelope created using helm charts.","patchEnvelopeName":"demo-side-car-patch-envelope","patchEnvelopeProjectId":"b67f1311-9e18-4473-8725-4495778faf88","patchEnvelopeProjectName":"default-project","patchEnvelopeTitle":"demo-side-car-patch-envelope","patchartifactvolume":false,"projectId":"b67f1311-9e18-4473-8725-4495778faf88","sideCarAppInstDescription":"sidecar app instance created using helm charts.","sideCarAppName":"zservice-sidecar-app","sideCarInstDrives":[],"sideCarInstInterfaces":[{"accessVlanId":0,"acls":[],"defaultNetInstance":false,"directattach":false,"intfname":"eth0","intforder":1,"io":null,"ipaddr":"","macaddr":"","netinstname":"","netinsttag":{"key":"demo-net-inst"}}],"sideCarInstNamePrefix":"side-car-inst","sideCarInstTags":{"app":"sidecar"},"sideCarInstTitle":"side-car-isnt","type":"VOLUME_INSTANCE_TYPE_CONTENT_TREE","volInstAccessmode":"VOLUME_INSTANCE_ACCESS_MODE_READONLY","volInstCleartext":true,"volInstContentTreeId":"","volInstImage":"ubuntu-sshd","volInstImplicit":false,"volInstNamePrefix":"ubuntu-sshd-volume-inst","volInstProjectId":"b67f1311-9e18-4473-8725-4495778faf88","volInstSizeBytes":"0","volInstTags":{},"volInstTitle":"ubuntu-sshd-volume-inst","volInstdescription":"ubuntu-sshd-volume-inst","volInstlabel":"tubuntu-sshd-volume-inst","volInstvolInstMultiattach":false}
EOT
  }

  target_asset {
    group_id = zedcloud_zasset_groups.example.id
    target_selector {
      asset_ids {
        values = ["fced5689-5b23-4e6e-8a37-62030dcaa4ac"]
      }
    }
  }

  depends_on = [
    zedcloud_zasset_groups.example
  ]
}