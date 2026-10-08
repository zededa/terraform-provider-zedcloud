# Plan-only: a URL in desc.logo must fail validation (the UI cannot render it).
resource "zedcloud_application" "test_tf_provider" {
	name        = "test_tf_provider-logo-url__SUFFIX__"
	title       = "test_tf_provider-logo-url__SUFFIX__"
	origin_type = "ORIGIN_LOCAL"
	manifest {
		ac_kind    = "VMManifest"
		ac_version = "1.2.0"
		name       = "logo-url"
		desc {
			app_category = "APP_CATEGORY_OTHERS"
			logo = {
				logo = "https://example.com/logo.png"
			}
		}
	}
}
