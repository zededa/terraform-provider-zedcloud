# Same resource, different content supplied inline: every input is ForceNew,
# so this replaces the artifact (new id) and deletes the old one.
resource "zedcloud_artifact" "test_tf_provider" {
	name           = "test_tf_provider-logo__SUFFIX__.png"
	content_base64 = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAADElEQVR4nGPgUnUAAACsAHCSJZIBAAAAAElFTkSuQmCC"
}
