resource "zedcloud_artifact" "test_tf_provider" {
	name        = "test_tf_provider-logo__SUFFIX__.png"
	source      = "__TESTDATA__/artifact/logo.png"
	source_hash = filesha256("__TESTDATA__/artifact/logo.png")
}
