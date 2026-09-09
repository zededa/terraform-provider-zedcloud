resource "zedcloud_enterprise" "test_tf_provider_white_labeling_legacy" {
  name  = "test_tf_provider_wl_legacy__SUFFIX__"
  title = "test_tf_provider_wl_legacy__SUFFIX__"

  # Pre-white_labeling style: the $ztag keys set directly in the raw attributes
  # map. Configs like this must keep working without a permanent diff.
  attributes = {
    "$ztag.entp.zui.ux.color.primary" = "#0A2540"
    "$ztag.entp.zui.ux.product.name"  = "Acme Edge"
  }
}
