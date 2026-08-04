# White-labeling the console for an enterprise.
#
# Terraform owns these values: the API replaces the whole attribute map on every
# update, so an apply overwrites white-labeling configured in the UI, and removing
# the white_labeling block clears it.
resource "zedcloud_enterprise" "acme" {
  name  = "acme"
  title = "Acme"

  white_labeling {
    primary_color   = "#0A2540"
    secondary_color = "#00B3A4"

    # Must be a URL of at most 256 characters. The API stores this as a tag value,
    # so an inlined (base64) image does not fit.
    logo_url = "https://acme.example.com/logo.svg"

    product_name = "Acme Edge"
  }
}

# How the console picks these up
#
# The browser calls the unauthenticated GET /api/v1/cloud/environment, which maps the
# request host to an enterprise by matching it against that enterprise's
# controllerHostURL, and returns that enterprise's white-label attributes.
#
# Note that the API does not persist controller_host_url on a child enterprise: it
# accepts the request and silently discards the value. Every enterprise created
# through Terraform is a child of the default parent, so it cannot be host-resolved
# today and the console falls back to the default parent enterprise. Setting
# controller_host_url here would also leave a permanent diff, since the value never
# comes back on read - which is why it is deliberately left out of this example.
