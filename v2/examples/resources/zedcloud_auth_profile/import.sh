# An authorization profile can be imported by its system-assigned ID.
terraform import zedcloud_auth_profile.okta AAGFABAEqnH4je5PHZTXSmHOs-XC

# ...or by profile name. Anything that is not a 28-character system ID is
# looked up by name and resolved to its ID.
terraform import zedcloud_auth_profile.okta corp-okta

# Note: the controller never returns oauth_profile.client_secret, so an
# imported profile carries no secret material. Supply client_secret in your
# configuration after importing; the first plan will show it as the only
# change.
