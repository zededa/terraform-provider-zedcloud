# A custom role can be imported by its system-assigned ID.
terraform import zedcloud_role.readonly CCGFABAEqnH4je5PHZTXSmHOs-ZE

# ...or by role name. Anything that is not a 28-character system ID is
# looked up by name and resolved to its ID.
terraform import zedcloud_role.readonly readonly-operators

# Built-in roles (type USER_ROLE_SYSTEM_DEFINED or USER_ROLE_CLUSTER, for
# example SysAdmin and SysRoot) cannot be imported: only the platform root
# user can change or delete them, so Terraform could never apply a plan for
# them. The import fails with an error. Reference them with the zedcloud_role
# data source instead.
