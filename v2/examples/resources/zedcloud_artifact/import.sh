# Import by artifact id (<uuid>_<name>), for example the id already stored in an
# app's manifest.desc.logo.logo. The API does not return file content, so the
# imported state has no source/content_base64. Add
#   lifecycle { ignore_changes = [source, source_hash, content_base64] }
# to the resource, or the next apply will upload a new artifact.
terraform import zedcloud_artifact.myapp_logo 07a928b1-c1ae-11f1-b06c-ca49049db012_logo
