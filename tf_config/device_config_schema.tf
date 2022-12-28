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
		debug_knob = {
	    	# computed
        # id =

	    	# optional
	    	debug_knob = false
	    	expired = false
	    	expiry = ""
		}
		default_net_inst = {
	    	# computed
        # id =

	    	# required
	    	device_id = ""
	    	kind = ""
	    	name = ""
	    	port = ""
	    	title = ""

	    	# optional
	    	cluster_id = ""
	    	description = ""
	    	device_default = ""
	    	dhcp = false
	    	dns_list = []
	    	ip = {
	        	# optional
	        	dhcp_range = {
             		# optional
            		end = ""
            		start = ""
            }
	        	dns = []
	        	domain = ""
	        	gateway = ""
	        	mask = ""
	        	ntp = ""
	        	subnet = ""
				}
	    	lisp = {
        		# optional
        		allocate = false
        		allocationprefix = ""
        		allocationprefixlen = 0
        		exportprivate = false
        		lispiid = 0
        		sp = []
        }
	    	network_policy_id = ""
	    	oconfig = ""
	    	opaque = {
        		# computed
            # id =

        		# required
        		device_id = ""
        		kind = ""
        		name = ""
        		port = ""
        		title = ""

        		# optional
        		cluster_id = ""
        		description = ""
        		device_default = ""
        		dhcp = false
        		dns_list = []
        		ip = {
            		# optional
            		dhcp_range = {
                		# optional
                		end = ""
                		start = ""
                }
            		dns = []
            		domain = ""
            		gateway = ""
            		mask = ""
            		ntp = ""
            		subnet = ""
            }
        		lisp = # LispConfig
        		network_policy_id = ""
        		oconfig = ""
        		opaque =
        		port_tags {}
        		project_id = ""
	         	revision = {
             		# required
             		created_at =
             		created_by = ""
             		curr = ""
             		updated_at =
             		updated_by = ""
             		# optional
             		prev = ""
            }
        		tags {}
        		type = ""
        }
	    	port_tags {}
	    	project_id = ""
	    	revision = {
        		# required
        		created_at =
        		created_by = ""
        		curr = ""
        		updated_at =
        		updated_by = ""
        		# optional
        		prev = ""
        }
	    	tags {}
	    	type = ""
    }
		deployment_tag = ""
		deprecated = ""
		description = ""
		dev_location = {
    		# optional
    		city = ""
    		country = ""
    		freeloc = ""
    		hostname = ""
    		latlong = ""
    		loc = ""
    		org = ""
    		postal = ""
    		region = ""
    		underlay_ip = ""
    }
		dlisp = {
    		# required
    		e_id = ""
    		e_id_hash_len = 0
    		client_addr = ""
    		eid_allocation_prefix = ""
    		eid_allocation_prefix_len = 0
    		lisp_instance = 0
    		lisp_map_servers = []
    		mode = ""
    		zed_servers = []
    }
		edgeviewconfig = {
    		# optional
    		app_policy = {
        		# optional
        		allow_app = false
        }
    		dev_policy = {
        		# optional
        		allow_dev = false
        }
    		ext_policy = {
        		# optional
        		allow_ext = false
        }
    		generation_id = 0
    		jwt_info = {
        		# optional
        		allow_sec = 0
        		disp_url = ""
        		encrypt = false
        		expire_sec = ""
        		num_inst = 0
        }
    		token = ""
    }
		generate_soft_serial = false
		identity = ""
		interfaces = []
		location = ""
		memory = 0
		obkey = ""
		onboarding = {
    		# optional
    		pem_cert = ""
    		pem_key = ""
    }
		prepare_power_off_counter = 0
		prepare_power_off_time = ""
		reset_counter = 0
		reset_time = ""
		revision = {
    		# required
    		created_at =
    		created_by = ""
    		curr = ""
    		updated_at =
    		updated_by = ""

    		# optional
    		prev = ""
    }
		serialno = ""
		site_pictures = []
		storage = 0
		tags {}
		thread = 0
		token = ""
		utype = ""
}
