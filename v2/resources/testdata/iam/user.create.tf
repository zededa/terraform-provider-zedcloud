resource "zedcloud_role" "test_tf_provider" {
  name = "test_tf_provider-test-role__SUFFIX__"
  title = "test_tf_provider-test-role__SUFFIX__"
  type = "USER_ROLE_USER_DEFINED"
  state = "ROLE_STATE_ACTIVE"
  scopes {
    enterprise_filter = []
    project_filter = []
  }
}

resource "zedcloud_user" "test_tf_provider" {
  depends_on = [
    zedcloud_role.test_tf_provider
  ]
  email = "user1__SUFFIX__@example.com"
  first_name = "user"
  role_id = zedcloud_role.test_tf_provider.id
  username = "user2__SUFFIX__@example.com"
  type = "AUTH_TYPE_LOCAL"

  # NFR-165 §3.4 regression guard. The model builder read the camelCase
  # swagger key, so this map was dropped from every request body and an
  # imported user could never converge. Keep it asserted, not ignored.
  #
  # The key must be exactly three `_`-separated segments -- see
  # ValidateCustomParam in zedcloud libs/zutils/validate.go. A plain
  # "department" is rejected by the controller with HTTP 400.
  custom_user_input = {
    tf_provider_department = "networking"
  }
}