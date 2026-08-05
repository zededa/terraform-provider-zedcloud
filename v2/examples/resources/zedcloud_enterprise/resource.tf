# White-labeling the console for an enterprise.
#
# Terraform owns these values: the API replaces the whole attribute map on every
# update, so an apply overwrites white-labeling configured in the UI, and removing
# the white_labeling block clears it.
resource "zedcloud_enterprise" "acme" {
  name  = "acme"
  title = "Acme"

  # The host users browse to. The console matches this to pick whose branding to
  # serve, and it works on a child enterprise like this one.
  #
  # Write-only: the API stores it but never returns it on read, so Terraform keeps
  # the configured value in state rather than reading it back.
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

# How the console picks these up
#
# The browser calls the unauthenticated GET /api/v1/cloud/environment before anyone
# logs in. The API resolves the enterprise from the X-HOST header, which the ingress
# sets from the browser's Host header, by matching it against controller_host_url. It
# returns that enterprise's white-label attributes, and falls back to the default
# parent enterprise when no enterprise matches the host.
#
# So beyond the Terraform config, the host has to actually reach the controller:
# it must resolve to the ingress, terminate TLS, and be served by a matching ingress
# server block - otherwise the request never arrives with that X-HOST value and the
# console keeps serving the default parent's branding.
