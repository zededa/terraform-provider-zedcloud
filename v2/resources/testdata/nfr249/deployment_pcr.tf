resource "zedcloud_project" "test_tf_provider_deployment_pcr" {
	name  = "test_tf_provider-nfr249-deployment__SUFFIX__"
	title = "test_tf_provider-nfr249-deployment__SUFFIX__"
	type  = "TAG_TYPE_DEPLOYMENT"
	tag_level_settings {
		flow_log_transmission = "NETWORK_INSTANCE_FLOW_LOG_TRANSMISSION_DISABLED"
		interface_ordering    = "INTERFACE_ORDERING_DISABLED"
	}
}

resource "zedcloud_deployment" "test_tf_provider_deployment_pcr" {
	name           = "test_tf_provider-nfr249-deployment__SUFFIX__"
	project_id     = zedcloud_project.test_tf_provider_deployment_pcr.id

	#@DEPLOYMENT_EXTRA@
}
