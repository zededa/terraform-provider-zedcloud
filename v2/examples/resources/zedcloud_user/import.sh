# A user can be imported by its system-assigned ID.
terraform import zedcloud_user.alice AAGFABAEqnH4je5PHZTXSmHOs-XC

# ...or, more conveniently, by username. Anything that is not a
# 28-character system ID is looked up by name and resolved to its ID.
terraform import zedcloud_user.alice alice@corp.com
