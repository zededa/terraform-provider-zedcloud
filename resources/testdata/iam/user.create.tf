resource "zedcloud_user" "test_tf_provider" {
  email = "user1@example.com"
  first_name = "user"
  role_id = "AAGFABCnEbWH56VAfZUHmv6ZfXRc"
  username = "user2@example.com"
}

data "zedcloud_user" "test_tf_provider" {
  depends_on = [
    zedcloud_user.test_tf_provider
  ]
  username = "user2@example.com"
  email = zedcloud_user.test_tf_provider.email
  role_id = zedcloud_user.test_tf_provider.role_id
}