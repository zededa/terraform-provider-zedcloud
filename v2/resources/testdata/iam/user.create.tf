resource "zedcloud_role" "test_tf_provider" {
  name = "test_tf_provider-test-role"
  title = "test_tf_provider-test-role"
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
  email = "user1@example.com"
  first_name = "user"
  role_id = zedcloud_role.test_tf_provider.id
  username = "user2@example.com"
  type = "AUTH_TYPE_LOCAL"
}