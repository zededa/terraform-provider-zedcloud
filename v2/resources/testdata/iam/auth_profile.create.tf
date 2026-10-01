resource "zedcloud_role" "test_tf_provider_authprofile" {
  name  = "test_tf_provider-authprofile-role__SUFFIX__"
  title = "test_tf_provider-authprofile-role__SUFFIX__"
  type  = "USER_ROLE_USER_DEFINED"
  state = "ROLE_STATE_ACTIVE"

  scopes {
    enterprise_filter = []
    project_filter    = []
  }
}

resource "zedcloud_auth_profile" "test_tf_provider" {
  depends_on = [
    zedcloud_role.test_tf_provider_authprofile
  ]

  name  = "test_tf_provider-authprofile__SUFFIX__"
  title = "test_tf_provider-authprofile__SUFFIX__"
  type  = "AUTH_TYPE_OAUTH"

  # Deliberately NOT active: only one profile can be active per enterprise,
  # and flipping that would change how everyone signs in to the tenant the
  # suite runs against.
  active = false

  default_role_id = zedcloud_role.test_tf_provider_authprofile.id

  oauth_profile {
    o_id_c_end_point = "https://example.invalid/.well-known/openid-configuration"
    client_id        = "test-tf-provider-client-id"

    # Not a credential: a sentinel the test looks for in the read-back to
    # prove the controller does not echo it. See UE-168.
    client_secret = "test-tf-provider-not-a-real-secret"
  }
}
