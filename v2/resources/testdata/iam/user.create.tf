resource "zedcloud_role" "test_tf_provider" {
  name = "test-role"
  title = "test-role"
  type = "USER_ROLE_SYSTEM_DEFINED"
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