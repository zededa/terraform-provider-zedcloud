# A custom role can be imported by its system-assigned ID.
terraform import zedcloud_role.readonly CCGFABAEqnH4je5PHZTXSmHOs-ZE

# ...or by role name. Anything that is not a 28-character system ID is
# looked up by name and resolved to its ID.
terraform import zedcloud_role.readonly readonly-operators
