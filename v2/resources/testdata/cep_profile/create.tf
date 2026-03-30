resource "zedcloud_cep_profile" "test_tf_provider" {
  name        = "test_tf_provider_cep"
  title       = "Test CEP Profile"
  description = "Test SCEP certificate enrollment profile"
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
