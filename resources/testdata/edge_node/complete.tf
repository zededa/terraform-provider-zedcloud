resource "zedcloud_edgenode" "complete" {
    # not supported by api
    # storage = 64
    # client_ip = "1.1.1.1"
    # deprecated = "false"
    # identity = "identity"
    # location = "berlin"
    # memory = 32
    # prepare_power_off_time = "undocumented format"
    # prepare_power_off_counter = 2
    # thread = 1
    # cluster_id = "1" conflict
    # cpu = 2
		# edgeviewconfig {
    		# # optional
    		# app_policy {
        		# # optional
        		# allow_app = false
        # }
    		# dev_policy {
        		# # optional
        		# allow_dev = false
        # }
    		# ext_policy {
        		# # optional
        		# allow_ext = false
        # }
    		# generation_id = 0
    		# jwt_info {
        		# # optional
        		# allow_sec = 1
            # disp_url = "http://localhost"
        		# encrypt = false
        		# expire_sec = "6000"
        		# num_inst = 1
        # }
    		# # token = "token"
    # }

    # we cannot setup interfaces via api calls it seems
		# interfaces {
	     	# # optional
    		# cost = 255
    		# intf_usage = "ADAPTER_USAGE_DISABLED"
    		# intfname = "system_interface_1_name"
    		# ipaddr = "127.0.0.1"
    		# macaddr = "00:00:00:00:00:00"
    		# netname = "system_interface_1_netname"
    		# tags = {
          # "system_interface_1_key" = "system_interface_1_value"
        # }
    # }

    config_item {
	      bool_value = true
	      float_value = 1.0
	    	key = "key"
	    	string_value = "string"
	    	uint32_value = 32
	    	uint64_value = 64
	    	value_type = "value type"
    }

		# required
		name = "complete"
		model_id = "2f716b55-2639-486c-9a2f-55a2e94146a6"
		project_id = "4754cd0f-82d7-4e06-a68f-ff9e23e75ccf"
		title = "complete-title"

    obkey = "5d0767ee-0547-4569-b530-387e526f8cb9"
		serialno = "2293dbe8-29ce-420c-8264-962857efc46b"

    admin_state = "ADMIN_STATE_ACTIVE"
    asset_id = "asset_id"
    # base_image {}
    deployment_tag = "depl_tag"
    description = "description"
		dev_location {
    		city = "berlin"
    		country = "germany"
    		freeloc = "freeloc"
    		hostname = "hostname"
    		loc = "52.520008, 13.404954"
    		org = "zededa"
    		postal = "10115"
    		region = "europe/west"
    		underlay_ip = ""
    }
		generate_soft_serial = false
		site_pictures = []
		tags = {
        "tag-key-1" = "tag-value-1"
    }
    token = "token"

    # not supported by API
		# dlisp {
    		# e_id = "e id"
    		# e_id_hash_len = 1
    		# client_addr = ""
    		# eid_allocation_prefix = "prefix"
    		# eid_allocation_prefix_len = 0
    		# lisp_instance = 0
    		# lisp_map_servers {
            # name_or_ip = "string"
            # credential = "string"
        # }
    		# mode = ""
    		# zed_servers {
            # host_name = "hostname"
            # e_id = [
                # "string"
            # ]
        # }
    # }

		default_net_inst {
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
	    	dns_list {}
	    	ip {
	        	# optional
	        	dhcp_range {
             		# optional
            		end = "end"
            		start = "start"
            }
	        	dns = ["1.1.1.1"]
	        	domain = ""
	        	gateway = ""
	        	mask = ""
	        	ntp = ""
	        	subnet = ""
				}
	    	lisp {
        		# optional
        		allocate = false
        		allocationprefix = ""
        		allocationprefixlen = 0
        		exportprivate = false
        		lispiid = 0
        		sp {}
        }
	    	network_policy_id = ""
	    	oconfig = ""
	    	port_tags = {
            "tag-key-1" = "tag-value-1"
        }
	    	project_id = ""
	    	tags = {}
	    	type = ""
	    	# opaque {
        		# # computed
            # # id =

        		# # required
        		# device_id = ""
        		# kind = ""
        		# name = ""
        		# port = ""
        		# title = ""

        		# # optional
        		# cluster_id = ""
        		# description = ""
        		# device_default = ""
        		# dhcp = false
        		# dns_list {}
        		# ip {
            		# # optional
            		# dhcp_range {
                		# # optional
                		# end = ""
                		# start = ""
                # }
            		# dns {}
            		# domain = ""
            		# gateway = ""
            		# mask = ""
            		# ntp = ""
            		# subnet = ""
            # }
        		# # lisp = # LispConfig
        		# network_policy_id = ""
        		# oconfig = ""
        		# # opaque =
        		# port_tags {}
        		# project_id = ""
        		# tags {}
        		# type = ""
        # }
    }
}
