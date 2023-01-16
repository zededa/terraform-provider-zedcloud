// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0
//

terraform {
  required_providers {
    zedcloud = {
      source = "zededa/zedcloud"
      //version = "0.0.7-alpha"
    }
  }
}

# data "zedcloud_edgenode" "data-test_tf_provider" {
# 		model_id = "2f716b55-2639-486c-9a2f-55a2e94146a6"
# 		name = "test_tf_provider-create_edgenode"
# 		project_id = "4754cd0f-82d7-4e06-a68f-ff9e23e75ccf"
# 		title = "test_tf_provider-create_edgenode-title"
# }

# data "zedcloud_datastore" "docker-registry" {
#    name = "Zededa-Dockerhub"
# }

resource "zedcloud_edgenode" "test_tf_provider" {
		# computed
    # id = data-test_tf_provider.id

    obkey = "5d0767ee-0547-4569-b530-387e526f8cb9"
		serialno = "2293dbe8-29ce-420c-8264-962857efc46b"
    identity = "identity"

		# required
		name = "test_tf_provider-create_edgenode"
		model_id = "2f716b55-2639-486c-9a2f-55a2e94146a6"
		project_id = "4754cd0f-82d7-4e06-a68f-ff9e23e75ccf"
		title = "test_tf_provider-create_edgenode-title"

    admin_state = "ADMIN_STATE_ACTIVE"
    asset_id = "asset_id"
    client_ip = "1.1.1.1"
    location = "berlin"
    # cluster_id = "1" conflict
    # base_image {}
    config_item {
	      bool_value = true
	      float_value = 1.0
	    	key = "key"
	    	string_value = "string"
	    	uint32_value = 32
	    	uint64_value = 64
	    	value_type = "value type"
    }
    # cpu = 2
    deployment_tag = "depl_tag"
    deprecated = false
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
		interfaces {}
    memory = 32
    prepare_power_off_counter = 0
    prepare_power_off_time = "undocumented format"
		site_pictures = []
    storage = 64
		tags = {
        "tag-key-1" = "tag-value-1"
    }
    thread = 1
    token = "token"

		# default_net_inst = {
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
		# 		}
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

# resource "zedcloud_edgenode" "test_tf_provider" {
#     admin_state = "ADMIN_STATE_ACTIVE"
#     client_ip = ""
#     cluster_id = ""
#     config_item = [
#         {
# 	          bool_value = false
# 	          float_value = 0.0
# 	        	key = ""
# 	        	string_value = ""
# 	        	uint32_value = 0
# 	        	uint64_value = 0
# 	        	value_type = ""
#         }
#     ]
#     cpu = 0
#     debug_knob_detail = {
#         debug_knob = false
#         expired = false
#         expiry = 1000000
#     }

# 		default_net_inst = {
# 		cluster_id = ""
# 		description = ""
# 		device_default = ""
# 		device_id = ""
# 		dhcp = false
# 		dns_list = {
# 		  addrs = ""
# 		  hostname = ""
# 		}
# 		ip = {
# 	  	dhcp_range = [
# 		    end = ""
# 		    start = ""
# 		  ]
# 		dns = ""
# 		domain = ""
# 		gateway = ""
# 		mask = ""
# }

# data "zedcloud_edgeapp" "ds_vmware_velocloud_sd_wan_medium" {
#    name = "vmware-velocloud-sd-wan-medium"
# }

# resource "zedcloud_network_instance" "TF-nwinst-ABC-test-kalyan-1" {
#     name = "TF-nwinst-ABC-test-kalyan-1"
#     description = "TF-nwinst-ABC-test-kalyan-1"
#     title = "TF-nwinst-ABC-test-kalyan-1"
#     device_default = false
#     device_id = data.zedcloud_edgenode.data-Sample-Device.id
#     kind = "NETWORK_INSTANCE_KIND_SWITCH"
#     type = "NETWORK_INSTANCE_DHCP_TYPE_V4"
# }

# resource "zedcloud_edgeapp" "TF-Sample-ubuntu-all-ip" {
#     name = "TF-Sample-ubuntu-all-ip"
#     title = "TF-Sample-ubuntu-all-ip"
#     description = "TF-Sample-ubuntu-all-ip"
#     manifest {
#         name = "TF-Sample-ubuntu-all-ip"
#         /*
#         desc_detail {
#             //app_category = "APP_CATEGORY_OPERATING_SYSTEM"
#             category = "infrastructure"
#             os = "Zenix"
#         }
#         */
#         image {
#             imagename = "Ztest-xenial-amd64"
#             maxsize = "1200000"
#             param {
#                 name = "bootparam"
#                 value = "+k/sre"
#             }
#             readonly = false
#             preserve = true
#             target = "Disk"
#             drvtype = "HDD"
#         }
#         interface {
#             name = "indirect"
#             directattach = false
#             privateip = false
#             acl {
#                 match {
#                     type = "protocol"
#                     value = "tcp"
#                 }
#                 match {
#                     type = "lport"
#                     value = "8022"
#                 }
#                 action {
#                     portmap = true
#                     mapparam {
#                         port = 22
#                     }
#                 }
#             }
#             acl {
#                 match {
#                     type = "host"
#                     value = ""
#                 }
#             }
#         }
#         enablevnc = true
#         owner {
#             user = "testuser"
#             email = "sampleemail@example.com"
#             company = "Zededa Inc."
#         }
#         resource {
#             name = "cpus"
#             value = "2"
#         }
#         resource {
#             name = "memory"
#             value = "1024000"
#         }
#     }
#     user_defined_version = "1.1"
#     cpus = 2
#     drives = 1
#     memory = 1024000
#     networks = 1
# }


# /*
# resource "zedcloud_edgeapp" "TF-AppManifestFile-Sample-ubuntu-all-ip" {
#     name = "TF-Sample-ubuntu-all-ip"
#     title = "TF-Sample-ubuntu-all-ip"
#     description = "TF-Sample-ubuntu-all-ip"
#     manifest_file = "./TFTest-ubuntu-xenial-16.04.json"
#     user_defined_version = "1.1"
#     cpus = 2
#     drives = 1
#     memory = 1024000
#     networks = 1
# }
# */
# data "zedcloud_network_instance" "TF-Edgenode-Default" {
#     name = "defaultLocal-gcp-alpha-acme-kalyan-1"
# }

# /*
# resource "zedcloud_network" "TF-test-sample-network-kalyan-1" {
#     name = "TF-Test-Network-kalyan-1"
#     enterprise_default = false
#     description = "Test Network for TF testing"
#     kind = "NETWORK_KIND_V4"
#     title = "TF-Test-Network"
#     //project_id = data.zedcloud_edgenode.data-Sample-Device.project_id
#     project_id = "b8552e75-54ee-4607-a715-1469dd132004"
#     ip {
#         dhcp = "NETWORK_DHCP_TYPE_CLIENT"
#     }
#     dns_list {
#         addrs = [ "10.10.10.1", "10.2.2.2"]
#         hostname = "www.example.com"
#     }
# }

# resource "zedcloud_network_instance" "TF-nwinst-gcp-alpha-acme-kalyan-1" {
#     name = "TF-nwinst-gcp-alpha-acme-kalyan-1"
#     description = "TF1-nwinst-gcp-alpha-acme-kalyan-1"
#     title = "TF-nwinst-gcp-alpha-acme-kalyan-1"
#     device_default = false
#     //project_id = "b8552e75-54ee-4607-a715-1469dd132004"
#     port = "uplink"
#     device_id = data.zedcloud_edgenode.data-Sample-Device.id
#     kind = "NETWORK_INSTANCE_KIND_LOCAL"
#     type = "NETWORK_INSTANCE_DHCP_TYPE_V4"
#     ip {
#         subnet = "192.168.10.0/24"
#         gateway = "192.168.10.1"
#         dns = [
#             "192.168.10.1"
#         ]
#         dhcp_range {
#             start = "192.168.10.32"
#             end = "192.168.10.40"
#         }
#     }
# }
# */
# resource "zedcloud_image" "TFTEST-image-container" {
#     name             = "tftest1-alpine-3.11-cont"
#     title            = "tftest1-alpine-3.11-cont-1"
#     datastore_id     = "9a322166-ec36-4e7f-b2c1-3e0199b076e8"
#     image_arch       = "AMD64"
#     image_format     = "CONTAINER"
#     image_rel_url    = "alpine:3.11"
#     image_type       = "IMAGE_TYPE_APPLICATION"
# }
# /*
# resource "zedcloud_edgeapp_instance" "Sample-EdgeAppInstance" {
#     name = "TF-Sample-Resource-ubuntu-bionic-EdgeAppInstance"
#     // name = "TF-Adarsh-Sample-Resource-ubuntu-bionic-EdgeAppInstance"
#     activate = true
#     app_type = "APP_TYPE_VM"
#     app_id = zedcloud_edgeapp.TF-Sample-ubuntu-all-ip.id
#     description = "TerraForm Sample App Instance"
#     device_id = data.zedcloud_edgenode.data-Sample-Device.id
#     interface {
#         intfname = "indirect"
#         netinstname = zedcloud_network_instance.TF-nwinst-ABC-test-kalyan-1.name
#     }
#     logs {
#        access = false
#     }
#     title = "TerraForm Sample App Instance"
#     remote_console = true
#     custom_config{
#         name = "Custom cloud-init config"
#         add = true
#         override = false
#         allow_storage_resize = true
#         field_delimiter = "###"
#         template = "I2Nsb3VkLWNvbmZpZwoKc3NoX3B3YXV0aDogWWVzCmhvc3RuYW1lOiBwb2N1c2VyCnVzZXJzOgogIC0gbmFtZTogcG9jdXNlcgogICAgc2hlbGw6IC9iaW4vYmFzaAogICAgc3VkbzogQUxMPShBTEwpIE5PUEFTU1dEOkFMTApjaHBhc3N3ZDoKICBsaXN0OiB8CiAgICAgIHBvY3VzZXI6cG9jdXNlcgogIGV4cGlyZTogZmFsc2UKCndyaXRlX2ZpbGVzOgogIC0gcGF0aDogL2hvbWUvcG9jdXNlci90ZXN0MQogICAgcGVybWlzc2lvbnM6ICcwNjQ0JwogICAgb3duZXI6IHBvY3VzZXI6cG9jdXNlcgogICAgZW5jb2Rpbmc6IGI2NAogICAgY29udGVudDogIyMjRklMRV9DT05URU5UIyMjCgpmaW5hbF9tZXNzYWdlOiAiVGhlIHN5c3RlbSBpcyBmaW5hbGx5IHVwLCBhZnRlciAkVVBUSU1FIHNlY29uZHMi"
#         variable_group {
#                 name = "Default Group 1"
#                 variable {
#                         name = "FILE_CONTENT"
#                         label = "Content to be written into the file"
#                         max_length = ""
#                         required = true
#                         default = "Default content"
#                         value = "Custom content"
#                         format = "VARIABLE_FORMAT_TEXT"
#                         encode = "FILE_ENCODING_UNSPECIFIED"
#                 }
#                 required = true
#         }
#     }
# }

# resource "zedcloud_edgeapp_instance" "Sample-EdgeAppInstance-2" {
#     name = "TF-Sample-Resource-ubuntu-bionic-EdgeAppInstance-2"
#     activate = true
#     app_type = "APP_TYPE_VM"
#     app_id = zedcloud_edgeapp.TF-Sample-ubuntu-all-ip.id
#     description = "TerraForm Sample App Instance"
#     device_id = data.zedcloud_edgenode.data-Sample-Device.id
#     interface {
#         intfname = "indirect"
#         netinstname = zedcloud_network_instance.TF-nwinst-ABC-test-kalyan-1.name
#     }
#     title = "TerraForm Sample App Instance"
# }
# */
