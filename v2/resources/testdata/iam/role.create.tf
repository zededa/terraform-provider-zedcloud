resource "zedcloud_role" "test_tf_provider" {
  name  = "test_tf_provider-import-role__SUFFIX__"
  title = "test_tf_provider-import-role__SUFFIX__"
  type  = "USER_ROLE_USER_DEFINED"
  state = "ROLE_STATE_ACTIVE"

  scopes {
    access_app          = "PermissionAccessRead"
    access_app_instance = "PermissionAccessRead"
    access_device       = "PermissionAccessRead"
    access_edge_app     = "PermissionAccessRead"
    access_enterprise   = "PermissionAccessRead"
    access_storage      = "PermissionAccessRead"
    access_user         = "PermissionAccessNone"
    enterprise_filter   = []
    project_filter      = []
  }
}
