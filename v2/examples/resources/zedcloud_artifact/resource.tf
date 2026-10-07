# Set the logo of an Edge app.
#
# The image is uploaded as an artifact; the app references the artifact's id
# under the key "logo", which is the only key the console reads. The console
# accepts PNG or JPEG images up to 5 MB, and does not display URLs.
resource "zedcloud_artifact" "myapp_logo" {
  name = "myapp-logo.png"

  source = "${path.module}/assets/myapp-logo.png"
  # Re-upload when the file changes. Without it, Terraform only notices a change
  # of path.
  source_hash = filesha256("${path.module}/assets/myapp-logo.png")

  # Artifacts cannot be modified, so any change creates a new one. Create the
  # new artifact, repoint the app, then delete the old one.
  lifecycle {
    create_before_destroy = true
  }
}

resource "zedcloud_application" "myapp" {
  name  = "myapp"
  title = "My App"
  # ...

  manifest {
    # ...
    desc {
      app_category = "APP_CATEGORY_OTHERS"
      logo = {
        logo = zedcloud_artifact.myapp_logo.id
      }
    }
  }
}

# Content can also be given inline, for example from another module's output.
# Only a SHA-256 of it is kept in state.
resource "zedcloud_artifact" "license" {
  name           = "license.txt"
  content_base64 = base64encode(file("${path.module}/LICENSE"))
}

# Destroying or replacing an artifact deletes it from the controller. An app
# duplicated in the console shares the original's logo artifact; set
# retain_on_destroy to leave the file in place when Terraform lets go of it.
resource "zedcloud_artifact" "shared_logo" {
  name              = "shared-logo.png"
  source            = "${path.module}/assets/shared-logo.png"
  retain_on_destroy = true
}
