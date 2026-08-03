# White-labeling the console for an enterprise.
#
# The console resolves a browser request to an enterprise by matching the request
# host against controller_host_url, then applies that enterprise's white-labeling.
# Note that controller_host_url is currently only honored on a parent enterprise.
#
# Terraform owns these values: the API replaces the whole attribute map on every
# update, so an apply overwrites white-labeling configured in the UI, and removing
# the white_labeling block clears it.
resource "zedcloud_enterprise" "acme" {
  name                = "acme"
  title               = "Acme"
  controller_host_url = "acme.zededa.net"

  white_labeling {
    primary_color   = "#0A2540"
    secondary_color = "#00B3A4"

    # Must be a URL of at most 256 characters. The API stores this as a tag value,
    # so an inlined (base64) image does not fit.
    logo_url = "https://acme.example.com/logo.svg"

    product_name = "Acme Edge"
  }
}
