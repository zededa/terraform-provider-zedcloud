resource "zedcloud_cep_profile" "test_tf_provider_pnac_v2_cep" {
  name        = "test_tf_provider_pnac_v2_cep"
  title       = "Test CEP for PNAC v2"
  description = "CEP profile used by PNAC v2 deployment policy test"
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

resource "zedcloud_project" "test_tf_provider_pnac_v2" {
  name  = "test_tf_provider_pnac_v2"
  title = "Test PNAC v2 Project"
  type  = "TAG_TYPE_DEPLOYMENT"

  tag_level_settings {
    flow_log_transmission = "NETWORK_INSTANCE_FLOW_LOG_TRANSMISSION_DISABLED"
    interface_ordering    = "INTERFACE_ORDERING_DISABLED"
  }
}

resource "zedcloud_deployment" "test_tf_provider_pnac_v2_depl" {
  depends_on = [
    zedcloud_cep_profile.test_tf_provider_pnac_v2_cep,
    zedcloud_project.test_tf_provider_pnac_v2,
  ]

  name           = "test_tf_provider_pnac_v2_depl"
  title          = "Test PNAC v2 Deployment"
  deployment_tag = "pnac:v2"

  project_id = zedcloud_project.test_tf_provider_pnac_v2.id

  device_policies {
    meta_data {
      name  = "test_tf_provider_pnac_v2_device_policy"
      title = "Test PNAC v2 Device Policy"
      policy_target_condition = {
        "pnac" : "v2"
      }
    }
    policy_sub_type = "DEVICE_POLICY_TYPE_PORT_BASED_NETWORK_ACCESS_CONTROL"

    port_based_network_access_control_policy {
      enable_port_based_network_access_control = true
      eap_method                               = "EAP_METHOD_TLS"
      eap_identity                             = "device-identity-v2"
      certificate_enrollment_id                = zedcloud_cep_profile.test_tf_provider_pnac_v2_cep.id
    }
  }
}
