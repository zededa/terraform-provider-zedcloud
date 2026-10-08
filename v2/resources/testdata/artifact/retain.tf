resource "zedcloud_artifact" "test_tf_provider" {
	name              = "test_tf_provider-logo-retained__SUFFIX__.png"
	source            = "__TESTDATA__/artifact/logo.png"
	retain_on_destroy = true
}
