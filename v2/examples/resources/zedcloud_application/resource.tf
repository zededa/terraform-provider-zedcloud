# Tag an Edge app in the marketplace (requires zedcloud 20.0.7+ or 21.0.0+).
#
# At most 32 tags; keys 3-512 and values 3-256 characters. Keys must not
# contain < > % & ? / \ ' " or a backtick; values must not contain ' " \ or a
# backtick.
#
# Terraform manages the tags once this argument is set. While it is absent,
# tags set in the UI or through the API are left as they are; removing it after
# it was set clears the tags.
resource "zedcloud_application" "router" {
  name  = "edge-router"
  title = "Edge router"
  # ...

  tags = {
    customer = "acme"
    tier     = "gold"
  }

  manifest {
    # ...
  }
}
