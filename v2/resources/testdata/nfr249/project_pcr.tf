resource "zedcloud_project" "test_tf_provider_pcr" {
	name  = "test_tf_provider-nfr249-project-pcr__SUFFIX__"
	title = "test_tf_provider-nfr249-project-pcr__SUFFIX__"
	type  = "TAG_TYPE_PROJECT"

	attestation_policy {
		title = "title"
		type  = "POLICY_TYPE_ATTESTATION"
		attestation_policy {
			type = "ATTEST_POLICY_TYPE_ACCEPT"
		}
	}
	tag_level_settings {
		flow_log_transmission = "NETWORK_INSTANCE_FLOW_LOG_TRANSMISSION_UNSPECIFIED"
		interface_ordering    = "INTERFACE_ORDERING_ENABLED"
	}

	#@PROJECT_EXTRA@
}
