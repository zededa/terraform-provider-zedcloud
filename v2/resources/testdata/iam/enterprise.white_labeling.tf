resource "zedcloud_enterprise" "test_tf_provider_white_labeling" {
  name  = "test_tf_provider_white_labeling"
  title = "test_tf_provider_white_labeling"

  white_labeling {
    primary_color   = "#0A2540"
    secondary_color = "#00B3A4"
    logo_url        = "https://acme.example.com/logo.svg"
    product_name    = "Acme Edge"
  }
}
