resource "zedcloud_edgenode" "test_tf_provider" {
		name = "test_tf_provider-create_edgenode"
		model_id = "2f716b55-2639-486c-9a2f-55a2e94146a6"
		project_id = "4754cd0f-82d7-4e06-a68f-ff9e23e75ccf"
		title = "title"
}

data "zedcloud_edgenode" "test_tf_provider" {
		name = "test_tf_provider-create_edgenode"
		model_id = zedcloud_edgenode.test_tf_provider.model_id
		title = zedcloud_edgenode.test_tf_provider.title
    depends_on = [
        zedcloud_edgenode.test_tf_provider
    ]
}

resource "zedcloud_network_instance" "complete" {
    depends_on = [
        data.zedcloud_edgenode.test_tf_provider
    ]

	  # required
    device_id = data.zedcloud_edgenode.test_tf_provider.id
    name = "complete"
    title = "complete"
    kind = "NETWORK_INSTANCE_KIND_LOCAL"
    port = "eth1"

	  # optional
    description = "zedcloud_network_instance-complete-description"
    type = "NETWORK_INSTANCE_DHCP_TYPE_V4"
    device_default = false
	  dhcp = false
    dns_list {
        addrs = [
            "10.1.2.2",
            "10.1.2.1"
        ]
        hostname = "wwww.ns2.example.com"
    }
    dns_list {
        addrs = [
            "10.1.1.1",
            "10.1.1.2"
        ]
        hostname = "wwww.ns1.example.com"
    }
    ip {
        domain = "htttp://example.com"

        gateway = "10.0.20.1"
        subnet = "10.0.20.0/24"

        dhcp_range {
            end = "10.0.20.100"
            start = "10.0.20.50"
        }
        # dns = [
        #     "10.0.20.1"
        # ]
        mask = "255.255.255.0"
        ntp = "10.1.0.2"
    }
    opaque {
        oconfig = "Test OConfig"
        type = "OPAQUE_CONFIG_TYPE_UNSPECIFIED"
    }
    tags = {
        "ni-tag1" = "ni-tag-value-1"
        "ni-tag2" = "ni-tag-value-2"
    }
    # not supported yet / by API?
	  # lisp = {
    		# # optional
    		# allocate = false
    		# allocationprefix = ""
    		# allocationprefixlen = 0
    		# exportprivate = false
    		# lispiid = 0
    		# sp = []
    # }
	  # network_policy_id = ""
	  # oconfig = ""
	  # opaque = {
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
    		# dns_list = []
    		# ip = {
        		# # optional
        		# dhcp_range = {
            		# # optional
            		# end = ""
            		# start = ""
            # }
        		# dns = []
        		# domain = ""
        		# gateway = ""
        		# mask = ""
        		# ntp = ""
        		# subnet = ""
        # }
    		# lisp = # LispConfig
    		# network_policy_id = ""
    		# oconfig = ""
    		# opaque =
    		# port_tags {}
    		# project_id = ""
	     	# revision = {
         		# # required
         		# created_at =
         		# created_by = ""
         		# curr = ""
         		# updated_at =
         		# updated_by = ""
         		# # optional
         		# prev = ""
        # }
    		# tags {}
    		# type = ""
    # }
	  # port_tags {}
	  # project_id = ""
	  # revision = {
    		# # required
    		# created_at =
    		# created_by = ""
    		# curr = ""
    		# updated_at =
    		# updated_by = ""
    		# # optional
    		# prev = ""
    # }
	  # tags {}
	  # type = ""
     # }
	 	# dlisp = {
     		# e_id = "e id"
     		# e_id_hash_len = 1
     		# client_addr = ""
     		# eid_allocation_prefix = "prefix"
     		# eid_allocation_prefix_len = 0
     		# lisp_instance = 0
     		# lisp_map_servers = []
     		# mode = ""
     		# zed_servers = []
     # }
	 	# edgeviewconfig = {
     		# # optional
     		# app_policy = {
         		# # optional
         		# allow_app = false
         # }
     		# dev_policy = {
         		# # optional
         		# allow_dev = false
         # }
     		# ext_policy = {
         		# # optional
         		# allow_ext = false
         # }
     		# generation_id = 0
     		# jwt_info = {
         		# # optional
         		# allow_sec = 0
         		# disp_url = ""
         		# encrypt = false
         		# expire_sec = ""
         		# num_inst = 0
         # }
     		# token = ""
     # }
}
