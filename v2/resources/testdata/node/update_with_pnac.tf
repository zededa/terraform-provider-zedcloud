resource "zedcloud_cep_profile" "test_tf_provider_node_pnac_cep" {
  name        = "test_tf_provider_node_pnac_cep"
  title       = "Test CEP for Node PNAC"
  description = "CEP profile used by Node PNAC onboarding test"
  scep_url    = "https://scep.example.com/scep"

  use_controller_proxy = false

  ca_cert_pem = [
    "-----BEGIN CERTIFICATE-----\nMIIC3DCCAcQCCQDBjgBA7A9NrDANBgkqhkiG9w0BAQsFADAwMRAwDgYDVQQDDAdU\nZXN0IENBMQ8wDQYDVQQKDAZaZWRlZGExCzAJBgNVBAYTAlVTMB4XDTI2MDMxODA5\nMjkxMFoXDTM2MDMxNTA5MjkxMFowMDEQMA4GA1UEAwwHVGVzdCBDQTEPMA0GA1UE\nCgwGWmVkZWRhMQswCQYDVQQGEwJVUzCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC\nAQoCggEBALWqOnCHZqun7HSvjd4hSISKJdgRnlH91TnuJI6ABgA6u0szHPyxK4w7\nLnseSqBpnbXujqlLVXa6jI4Akm0Ycfo3+IBVbhKCxXf0Mxj/pfq7QR3IVjOt4pjE\nV+9pb489uAFRj7D0/0I/kUSpalO0AlXiCCJYEASBPQlzkEoWer5trAXKdY89+nVp\nj/i0rG9bAukwOa1t4PNZWQx4WdsC/bFN2Bts51/GKZ3ahjSfb1h2ZBJONs6IBcEY\neGkrNzfRPhMD5LWGwIvgHrheIhW+qpolv5v/QYEuwqmFct1XFdjrl6HDl0X8EekU\nNgxCN8IbiiD089ggZ1Cc3FsOkSY8rW0CAwEAATANBgkqhkiG9w0BAQsFAAOCAQEA\nApV+iIvnZPaxNhwFPHAdXbCwtlCr4Tc7HzZnTCdrRKN/pTRWbqCUMHu/DBA7TgNo\nMIsOOjJqw0nco2+9AAZ91bfL2GbM13o0sQ4cdXfiyU9xuD6RycS6tK4h4rRN496I\nMqAiVEK8Wg84di+/qpdh1xg87iKjlChuJcN2i05wIJG5sAS9W2w6ug7ihasxloPk\nyjLbsYXdwghR6FiTGYfnXyF+TD7M20EuAgujF281MvQ3KN3swlIZCbWKE/3roL24\nub6+O2KIvbyWLQHcAjLBv/4+pg3UrxiefPP0yeEEMzOuh9mLu49NXpE2JIXhQejF\naHEKNVshNtW5q4qXu4bVzQ==\n-----END CERTIFICATE-----",
  ]

  csr_profile {
    common_name          = "device-{{.DeviceID}}"
    organization         = "Zededa"
    organizational_unit  = "Engineering"
    country              = "US"
    renew_period_percent = 80
    key_type             = "KEY_TYPE_RSA_2048"
    hash_algorithm       = "HASH_ALGORITHM_SHA256"
  }

  secret {
    challenge_password = "test-challenge-password"
  }

  lifecycle {
    ignore_changes = [ca_cert_pem]
  }
}

resource "zedcloud_project" "test_tf_provider_node_pnac" {
  name  = "test_tf_provider_node_pnac"
  title = "Test Node PNAC Project"
  type  = "TAG_TYPE_PROJECT"

  attestation_policy {
    title = "test_tf_provider_node_pnac_attest"
    type  = "POLICY_TYPE_ATTESTATION"

    attestation_policy {
      type = "ATTEST_POLICY_TYPE_ACCEPT"
    }
  }

  port_based_network_access_control_policy {
    title = "Test Node PNAC Policy"
    type  = "POLICY_TYPE_PORT_BASED_NETWORK_ACCESS_CONTROL"

    port_based_network_access_control_policy {
      enable_port_based_network_access_control = true
      eap_method                               = "EAP_METHOD_TLS"
      eap_identity                             = "device-identity-node"
      certificate_enrollment_id                = zedcloud_cep_profile.test_tf_provider_node_pnac_cep.id
    }
  }
}

resource "zedcloud_brand" "test_tf_provider_node_pnac" {
  name        = "test_tf_provider_node_pnac"
  title       = "test_tf_provider_node_pnac"
  description = "description"
  origin_type = "ORIGIN_LOCAL"
}

resource "zedcloud_model" "test_tf_provider_node_pnac" {
  brand_id    = zedcloud_brand.test_tf_provider_node_pnac.id
  name        = "test_tf_provider_node_pnac"
  title       = "test_tf_provider_node_pnac"
  type        = "AMD64"
  origin_type = "ORIGIN_LOCAL"
  state       = "SYS_MODEL_STATE_ACTIVE"
  attr = {
    memory  = "8G"
    storage = "100G"
    Cpus    = "4"
  }
  io_member_list {
    ztype        = "IO_TYPE_ETH"
    phylabel     = "firstEth"
    usage        = "ADAPTER_USAGE_MANAGEMENT"
    assigngrp    = "eth0"
    logicallabel = "ethernet0"
    phyaddrs = {
      Ifname  = "eth0"
      PciLong = "0000:02:00.0"
    }
    usage_policy = {
      FreeUplink = true
    }
    cost = 0
  }
  depends_on = [zedcloud_brand.test_tf_provider_node_pnac]
}

resource "zedcloud_edgenode" "test_tf_provider_node_pnac" {
  depends_on = [
    zedcloud_project.test_tf_provider_node_pnac,
    zedcloud_model.test_tf_provider_node_pnac,
  ]
  name       = "test_tf_provider_node_pnac"
  title      = "test_tf_provider_node_pnac"
  model_id   = zedcloud_model.test_tf_provider_node_pnac.id
  project_id = zedcloud_project.test_tf_provider_node_pnac.id
  serialno   = "node-pnac-test-serial-001"
  admin_state = "ADMIN_STATE_ACTIVE"

  interfaces {
    intfname   = "eth0"
    intf_usage = "ADAPTER_USAGE_MANAGEMENT"
    enable_port_based_network_access_control = false
  }
}
