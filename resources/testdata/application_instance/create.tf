resource "zedcloud_application" "test_tf_provider" {
		name = "test_tf_provider"
    title = "ubuntu-all-ip"
    description = "ubuntu-all-ip"
    user_defined_version = "1.1"
    origin_type = "ORIGIN_LOCAL"
}

data "zedcloud_application" "test_tf_provider" {
		name = "test_tf_provider"
    depends_on = [
        zedcloud_application.test_tf_provider
    ]
}

resource "zedcloud_edgenode" "test_tf_provider" {
		name = "test_tf_provider"
		model_id = "2f716b55-2639-486c-9a2f-55a2e94146a6"
		title = "complete-title"
}

data "zedcloud_edgenode" "test_tf_provider" {
		name = "test_tf_provider"
    depends_on = [
        zedcloud_edgenode.test_tf_provider
    ]
}

resource "app_instance"  "test_tf_provider" {
    depends_on = [
        zedcloud_application.test_tf_provider,
        zedcloud_edgenode.test_tf_provider,
        zedcloud_network_instance.test_tf_provider
    ]

		# required
    name = "test_tf_provider"
    activate = true
    description = "tf test"

    app_id = zedcloud_application.test_tf_provider.id
    device_id = zedcloud_edgenode.test_tf_provider.id

    interface {
        intfname = "indirect"
        netinstname = zedcloud_network_instance.test_tf_provider.name
    }
    title = "tf test"
    remote_console = true

		drives = []

		# optional
		app_policy_id = ""
    app_type = "APP_TYPE_VM"
		bundleversion = ""
		cluster_id = ""
		collect_stats_ip_addr = ""
    custom_config {
	    	add = false
	    	allow_storage_resize = false
	    	field_delimiter = ","
	    	name = "custom_config_name"
	    	override = false
	    	template = "dGVzdA=="
	    	variable_groups {
	        	condition {
	            	# optional
	            	name = "condition"
	            	operator = "CONDITION_OPERATOR_EQUALTO"
	            	value = "val"
            }
	        	name = "var_group"
	        	required = false
	        	# variables {
	            	# # required
	            	# format = ""
	            	# label = ""
	            	# name = ""
	            	# required = false
	            	# # optional
	            	# default = ""
	            	# encode = ""
	            	# max_length = ""
	            	# options {
	                	# # optional
	                	# label = ""
	                	# value = ""
                # }
	            	# process_input = ""
	            	# type = ""
	            	# value = ""
            # }
        }
    }
		deployment_type = ""
		description = ""
		is_secret_updated = false
		logs = {
		    access = ""
		}
		manifest_info = {
	    	bundle_version = ""
	    	details = {
	        	authentication_type = ""
	        	cloud_to_device_message_count = 0
	        	description = ""
	        	desired = ""
	        	last_desired_status = ""
	        	module_count = 0
	        	reported = ""
	        	status_code = 0
	        	tags = ""
				}
	    	next_bundle_version = ""
	    	params {"" = ""}
	    	transition_action = ""
		}
		purge = {
	    	counter = 0
		    ops_time = ""
		}
		refresh = {
	    	counter = 0
		    ops_time = ""
		}
		remote_console = false
		restart = {
	    	counter = 0
		    ops_time = ""
		}
		start_delay_in_seconds = 0
		tags {"" = ""}
		user_data = ""
		user_defined_version = ""
		vminfo = {
	    	# required
	    	cpus = 0
	    	memory = 0
	    	mode = ""
	    	vnc = false
	    	# optional
	    	cpu_pinning_enabled = false
		}
}
