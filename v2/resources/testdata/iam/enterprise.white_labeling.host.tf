resource "zedcloud_enterprise" "test_tf_provider_wl_host" {
  name  = "test_tf_provider_wl_host"
  title = "test_tf_provider_wl_host"

  # Maps a browser host to this enterprise so the console serves its branding.
  # The API stores this but never returns it on read.
  controller_host_url = "tf-wl-host.local.zededa.net"

  white_labeling {
    primary_color = "#112233"
    product_name  = "Host Brand"
  }
}
