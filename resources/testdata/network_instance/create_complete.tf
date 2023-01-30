resource "zedcloud_network_instance" "complete" {
	  # required
    device_id = data.zedcloud_edgenode.complete.id
    name = "complete-name"
    title = "complete-title"
    kind = "NETWORK_INSTANCE_KIND_LOCAL"
    port = "eth1"

	  # optional
    description = "zedcloud_network_instance-complete-description"
    type = "NETWORK_INSTANCE_DHCP_TYPE_V4"
    device_default = false
	  dhcp = false
    dns_list {
        addrs = [
            "10.1.1.1",
            "10.1.1.2"
        ]
        hostname = "wwww.ns1.example.com"
    }
    dns_list {
        addrs = [
            "10.1.2.1",
            "10.1.2.2"
        ]
        hostname = "wwww.ns2.example.com"
    }
    ip {
        # not supported by API
        # domain = "htttp://example.com"
        gateway = "10.1.0.1"
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
