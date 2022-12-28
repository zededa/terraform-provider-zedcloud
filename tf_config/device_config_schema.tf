resource "device_config"  "device_config_test" {
		# computed
    # id =

		# required
		model_id = ""
		name = ""
		project_id = ""
		title = ""

		# optional
		admin_state = ""
		asset_id = ""
		base_image = []
		base_os_retry_counter = 0
		base_os_retry_time = ""
		client_ip = ""
		cluster_id = ""
		config_item = []
		cpu = 0
		debug_knob = # DebugKnobDetail
		default_net_inst = # NetInstConfig
		deployment_tag = ""
		deprecated = ""
		description = ""
		dev_location = # GeoLocation
		dlisp = # DeviceLisp
		edgeviewconfig = # EdgeviewCfg
		generate_soft_serial = false
		identity = ""
		interfaces = []
		location = ""
		memory = 0
		obkey = ""
		onboarding = # DeviceCerts
		prepare_power_off_counter = 0
		prepare_power_off_time = ""
		reset_counter = 0
		reset_time = ""
		revision = # ObjectRevision
		serialno = ""
		site_pictures = [""]
		storage = 0
		tags {"" = ""}
		thread = 0
		token = ""
		utype = ""
}
