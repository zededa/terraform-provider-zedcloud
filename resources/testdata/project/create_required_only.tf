resource "zedcloud_project" "test_tf_provider" {
		# computed
    # id =
    # app_policy =
    # attr =
    # cloud_policy =
    # module_policy =
    # numdevices =
    # revision =

		# required
		name = "test_tf_provider"
    title = "title"

		# optional
		type = "TAG_TYPE_PROJECT"
		# attestation_policy = {
		#     # computed
        # # id =
        # # revision =
        # # status =

		#     # required
		#     name = "test_tf_provider"
		#     title = "title"
		#     type = "POLICY_TYPE_UNSPECIFIED"

		#     # optional
		#     app_policy = {
	        	# # required
	        	# apps = []
        # }
		#     attestation_policy = {
	        	# # computed
            # # id =
	        	# # required
	        	# type = "ATTEST_POLICY_TYPE_UNSPECIFIED"
        # }
		#     attr {
		# 			"attribue1" = "att-value1"
		# 		}
		#     azure_policy = {
	        	# # required
	        	# app_id = ""
	        	# app_password = ""
	        	# tenant_id = ""

	        	# # optional
	        	# azure_resource_and_services = {
	            	# # optional
	            	# s_k_u = {
	                	# # optional
	                	# capacity = ""
	                	# name = ""
	                	# tier = ""
                # }
	            	# create_by_default = false
	            	# name = ""
	            	# region = ""
	            	# resource_group_name = ""
	            	# subscription_id = ""
            # }
	        	# certificate = {
	            	# # optional
	            	# basic_contraints_valid = false
	            	# cert = ""
	            	# crypto_key = ""
	            	# ecdsa_encryption = {
	                	# # optional
	                	# curve = ""
                # }
	            	# encrypted_secrets {"" = ""}
	            	# exportable = false
	            	# extended_key_usage = [""]
	            	# issuer = {
	                	# # optional
	                	# common_name = ""
	                	# country = [""]
	                	# locality = [""]
	                	# organization = [""]
	                	# organizational_unit = [""]
	                	# postal_code = [""]
	                	# province = [""]
	                	# serial_number = ""
                # }
	            	# key_usage = 0
	            	# pass_phrase = ""
	            	# public_key = ""
	            	# public_key_algorithm = ""
	            	# pvt_key = ""
	            	# reuse_key = false
	            	# rsa_ecryption = {
	                	# # optional
	                	# rsa_bits = ""
                # }
	            	# san_values = {
	                	# # optional
	                	# dns = [""]
	                	# emaild_ids = [""]
	                	# hosts = [""]
	                	# ips = [""]
	                	# upns = [""]
	                	# uris = [""]
                # }
	            	# serial_number = ""
	            	# signature_algorithm = ""
	            	# subject = {
	                	# # optional
	                	# dns = [""]
	                	# emaild_ids = [""]
	                	# hosts = [""]
	                	# ips = [""]
	                	# upns = [""]
	                	# uris = [""]
                # }
	            	# valid_from = 0.0
	            	# valid_till = 0.0
            # }
	        	# crypto_key = ""
	        	# custom_deployment_managed = false
	        	# encrypted_secrets {"" = ""}
        # }
		#     cluster_policy ={
        		# # required
        		# app_policy_id = ""
        		# network_policy_id = ""
        		# type = ""
        		# # optional
        		# cluster_config = {
	            	# # optional
	            	# min_nodes_required = 0
            # }
        # }
		#     description = ""
		#     edgeview_policy = {
	        	# # required
	        	# edgeview_allow = false
	        	# # optional
	        	# access_allow_change = false
	        	# edgeviewcfg = # EdgeviewCfg
	        	# max_expire_sec = 0
	        	# max_inst = 0
        # }
		#     local_operator_console_policy = {
        		# # required
        		# loc_url = ""
        # }
		#     module_policy =  {
	        	# # required
	        	# apps = []
	        	# priority = 0
	        	# # optional
	        	# etag = ""
	        	# azure_edge_agent = {
	            	# # computed
                # # drives =
                # # id =
	            	# # required
	            	# cpus = 0
	            	# drives = 0
	            	# manifest_json = {
	                	# # optional
	                	# ac_kind = ""
	                	# ac_version = ""
	                	# app_type = ""
	                	# configuration = {
	                    	# # optional
	                    	# custom_config = {
	                        	# # optional
	                        	# add = false
	                        	# allow_storage_resize = false
	                        	# field_delimiter = ""
	                        	# name = ""
	                        	# override = false
	                        	# template = ""
	                        	# variable_groups = []
                        # }
                    # }
	                	# container_detail = {
	                    	# # optional
	                    	# container_create_option = ""
                    # }
	                	# cpu_pinning_enabled = false
	                	# deployment_type = ""
	                	# desc = {
	                    	# # required
	                    	# app_category = ""
	                    	# # optional
	                    	# agreement_list {"" = ""}
	                    	# category"All"
	                    	# license_list {"" = ""}
	                    	# logo {"" = ""}
	                    	# os = ""
	                    	# screenshot_list {"" = ""}
	                    	# support = ""
                    # }
	                	# description = ""
	                	# display_name = ""
	                	# enablevnc = false
	                	# images = []
	                	# interfaces = []
	                	# module = {
	                    	# # optional
	                    	# environment {"" = ""}
	                    	# module_type = ""
	                    	# routes {"" = ""}
	                    	# twin_detail = ""
                    # }
	                	# name = ""
	                	# owner = {
	                    	# # optional
	                    	# company = ""
	                    	# email = ""
	                    	# group = ""
	                    	# user = ""
	                    	# website = ""
                    # }
	                	# permissions = []
	                	# resources = []
	                	# vmmode = ""
                # }
	            	# memory = 0
	            	# name = ""
	            	# networks = 0
	            	# origin_type = ""
	            	# title = ""
	            	# # optional
	            	# app_id = ""
	            	# app_version = ""
	            	# description = ""
	            	# interfaces = []
	            	# name_app_part = ""
	            	# name_project_part = ""
	            	# naming_scheme = ""
	            	# parent_detail = {
	                	# # computed
                    # # id_of_parent_object =
                    # # version_of_parent_object =
	                	# # required
	                	# id_of_parent_object = ""
	                	# # optional
	                	# reference_exists = false
	                	# update_available = false
                # }
	            	# start_delay_in_seconds = 0
	            	# storage = 0
	            	# tags {"" = ""}
            # }
	        	# azure_edge_hub = {
	            	# # computed
                # # drives =
                # # id =
	            	# # required
	            	# cpus = 0
	            	# drives = 0
	            	# manifest_json = {
	                	# # optional
	                	# ac_kind = ""
	                	# ac_version = ""
	                	# app_type = ""
	                	# configuration = {
	                    	# # optional
	                    	# custom_config = {
	                        	# # optional
	                        	# add = false
	                        	# allow_storage_resize = false
	                        	# field_delimiter = ""
	                        	# name = ""
	                        	# override = false
	                        	# template = ""
	                        	# variable_groups = []
                        # }
                    # }
	                	# container_detail = {
	                    	# # optional
	                    	# container_create_option = ""
                    # }
	                	# cpu_pinning_enabled = false
	                	# deployment_type = ""
	                	# desc = {
	                    	# # required
	                    	# app_category = ""
	                    	# # optional
	                    	# agreement_list {"" = ""}
	                    	# category"All"
	                    	# license_list {"" = ""}
	                    	# logo {"" = ""}
	                    	# os = ""
	                    	# screenshot_list {"" = ""}
	                    	# support = ""
                    # }
	                	# description = ""
	                	# display_name = ""
	                	# enablevnc = false
	                	# images = []
	                	# interfaces = []
	                	# module = {
	                    	# # optional
	                    	# environment {"" = ""}
	                    	# module_type = ""
	                    	# routes {"" = ""}
	                    	# twin_detail = ""
                    # }
	                	# name = ""
	                	# owner = {
	                    	# # optional
	                    	# company = ""
	                    	# email = ""
	                    	# group = ""
	                    	# user = ""
	                    	# website = ""
                    # }
	                	# permissions = []
	                	# resources = []
	                	# vmmode = ""
                # }
	            	# memory = 0
	            	# name = ""
	            	# networks = 0
	            	# origin_type = ""
	            	# title = ""
	            	# # optional
	            	# app_id = ""
	            	# app_version = ""
	            	# description = ""
	            	# interfaces = []
	            	# name_app_part = ""
	            	# name_project_part = ""
	            	# naming_scheme = ""
	            	# parent_detail = {
	                	# # computed
                    # # id_of_parent_object =
                    # # version_of_parent_object =
	                	# # required
	                	# id_of_parent_object = ""
	                	# # optional
	                	# reference_exists = false
	                	# update_available = false
                # }
	            	# start_delay_in_seconds = 0
	            	# storage = 0
	            	# tags {"" = ""}
            # }
	        	# labels {"" = ""}
	        	# metrics = {
	            	# # optional
	            	# queries {"" = ""}
	            	# results {"" = ""}
            # }
	        	# routes {"" = ""}
	        	# target_condition = ""
	        	# target_condition_new {"" = ""}
        # }
		#     network_policy = {
	        	# # required
	        	# net_instance_config = []
        # }
		#     status_message = ""
    # }
		# deployment = {
	    	# # optional
	    	# app_inst_policies = []
	    	# cluster_policy = {
	        	# # optional
	        	# id = ""
	        	# name = ""
	        	# revision = {
	            	# # required
	            	# created_at =
	            	# created_by = ""
	            	# curr = ""
	            	# updated_at =
	            	# updated_by = ""
	            	# # optional
	            	# prev = ""
		# 				}
	        	# title = ""
		# 		}
	    	# deployment_tag = ""
	    	# device_policies = []
	    	# id = ""
	    	# integration_policy = {
	        	# # optional
	        	# id = ""
	        	# name = ""
	        	# revision = {
	            	# # required
	            	# created_at =
	            	# created_by = ""
	            	# curr = ""
	            	# updated_at =
	            	# updated_by = ""
	            	# # optional
	            	# prev = ""
		# 				}
	        	# title = ""
		# 		}
	    	# name = ""
	    	# network_inst_policies = []
	    	# revision = {
	            	# # required
	            	# created_at =
	            	# created_by = ""
	            	# curr = ""
	            	# updated_at =
	            	# updated_by = ""
	            	# # optional
	            	# prev = ""
		# 		}
	    	# title = ""
	    	# volume_inst_policies = []
    # }
		# description = ""
		# edgeview_policy = {
		#     # computed
        # # id =
        # # revision =
        # # status =

		#     # required
		#     name = ""
		#     title = ""
		#     type = ""

		#     # optional
		#     app_policy = {
	        	# # required
	        	# apps = []
        # }
		#     attestation_policy = {
	        	# # computed
            # # id =
	        	# # required
	        	# type = ""
        # }
		#     attr {"" = ""}
		#     azure_policy = {
	        	# # required
	        	# app_id = ""
	        	# app_password = ""
	        	# tenant_id = ""

	        	# # optional
	        	# azure_resource_and_services = {
	            	# # optional
	            	# s_k_u = {
	                	# # optional
	                	# capacity = ""
	                	# name = ""
	                	# tier = ""
                # }
	            	# create_by_default = false
	            	# name = ""
	            	# region = ""
	            	# resource_group_name = ""
	            	# subscription_id = ""
            # }
	        	# certificate = {
	            	# # optional
	            	# basic_contraints_valid = false
	            	# cert = ""
	            	# crypto_key = ""
	            	# ecdsa_encryption = {
	                	# # optional
	                	# curve = ""
                # }
	            	# encrypted_secrets {"" = ""}
	            	# exportable = false
	            	# extended_key_usage = [""]
	            	# issuer = {
	                	# # optional
	                	# common_name = ""
	                	# country = [""]
	                	# locality = [""]
	                	# organization = [""]
	                	# organizational_unit = [""]
	                	# postal_code = [""]
	                	# province = [""]
	                	# serial_number = ""
                # }
	            	# key_usage = 0
	            	# pass_phrase = ""
	            	# public_key = ""
	            	# public_key_algorithm = ""
	            	# pvt_key = ""
	            	# reuse_key = false
	            	# rsa_ecryption = {
	                	# # optional
	                	# rsa_bits = ""
                # }
	            	# san_values = {
	                	# # optional
	                	# dns = [""]
	                	# emaild_ids = [""]
	                	# hosts = [""]
	                	# ips = [""]
	                	# upns = [""]
	                	# uris = [""]
                # }
	            	# serial_number = ""
	            	# signature_algorithm = ""
	            	# subject = {
	                	# # optional
	                	# dns = [""]
	                	# emaild_ids = [""]
	                	# hosts = [""]
	                	# ips = [""]
	                	# upns = [""]
	                	# uris = [""]
                # }
	            	# valid_from = 0.0
	            	# valid_till = 0.0
            # }
	        	# crypto_key = ""
	        	# custom_deployment_managed = false
	        	# encrypted_secrets {"" = ""}
        # }
		#     cluster_policy ={
        		# # required
        		# app_policy_id = ""
        		# network_policy_id = ""
        		# type = ""
        		# # optional
        		# cluster_config = {
	            	# # optional
	            	# min_nodes_required = 0
            # }
        # }
		#     description = ""
		#     edgeview_policy = {
	        	# # required
	        	# edgeview_allow = false
	        	# # optional
	        	# access_allow_change = false
	        	# edgeviewcfg = # EdgeviewCfg
	        	# max_expire_sec = 0
	        	# max_inst = 0
        # }
		#     local_operator_console_policy = {
        		# # required
        		# loc_url = ""
        # }
		#     module_policy =  {
	        	# # required
	        	# apps = []
	        	# priority = 0
	        	# # optional
	        	# etag = ""
	        	# azure_edge_agent = {
	            	# # computed
                # # drives =
                # # id =
	            	# # required
	            	# cpus = 0
	            	# drives = 0
	            	# manifest_json = {
	                	# # optional
	                	# ac_kind = ""
	                	# ac_version = ""
	                	# app_type = ""
	                	# configuration = {
	                    	# # optional
	                    	# custom_config = {
	                        	# # optional
	                        	# add = false
	                        	# allow_storage_resize = false
	                        	# field_delimiter = ""
	                        	# name = ""
	                        	# override = false
	                        	# template = ""
	                        	# variable_groups = []
                        # }
                    # }
	                	# container_detail = {
	                    	# # optional
	                    	# container_create_option = ""
                    # }
	                	# cpu_pinning_enabled = false
	                	# deployment_type = ""
	                	# desc = {
	                    	# # required
	                    	# app_category = ""
	                    	# # optional
	                    	# agreement_list {"" = ""}
	                    	# category"All"
	                    	# license_list {"" = ""}
	                    	# logo {"" = ""}
	                    	# os = ""
	                    	# screenshot_list {"" = ""}
	                    	# support = ""
                    # }
	                	# description = ""
	                	# display_name = ""
	                	# enablevnc = false
	                	# images = []
	                	# interfaces = []
	                	# module = {
	                    	# # optional
	                    	# environment {"" = ""}
	                    	# module_type = ""
	                    	# routes {"" = ""}
	                    	# twin_detail = ""
                    # }
	                	# name = ""
	                	# owner = {
	                    	# # optional
	                    	# company = ""
	                    	# email = ""
	                    	# group = ""
	                    	# user = ""
	                    	# website = ""
                    # }
	                	# permissions = []
	                	# resources = []
	                	# vmmode = ""
                # }
	            	# memory = 0
	            	# name = ""
	            	# networks = 0
	            	# origin_type = ""
	            	# title = ""
	            	# # optional
	            	# app_id = ""
	            	# app_version = ""
	            	# description = ""
	            	# interfaces = []
	            	# name_app_part = ""
	            	# name_project_part = ""
	            	# naming_scheme = ""
	            	# parent_detail = {
	                	# # computed
                    # # id_of_parent_object =
                    # # version_of_parent_object =
	                	# # required
	                	# id_of_parent_object = ""
	                	# # optional
	                	# reference_exists = false
	                	# update_available = false
                # }
	            	# start_delay_in_seconds = 0
	            	# storage = 0
	            	# tags {"" = ""}
            # }
	        	# azure_edge_hub = {
	            	# # computed
                # # drives =
                # # id =
	            	# # required
	            	# cpus = 0
	            	# drives = 0
	            	# manifest_json = {
	                	# # optional
	                	# ac_kind = ""
	                	# ac_version = ""
	                	# app_type = ""
	                	# configuration = {
	                    	# # optional
	                    	# custom_config = {
	                        	# # optional
	                        	# add = false
	                        	# allow_storage_resize = false
	                        	# field_delimiter = ""
	                        	# name = ""
	                        	# override = false
	                        	# template = ""
	                        	# variable_groups = []
                        # }
                    # }
	                	# container_detail = {
	                    	# # optional
	                    	# container_create_option = ""
                    # }
	                	# cpu_pinning_enabled = false
	                	# deployment_type = ""
	                	# desc = {
	                    	# # required
	                    	# app_category = ""
	                    	# # optional
	                    	# agreement_list {"" = ""}
	                    	# category"All"
	                    	# license_list {"" = ""}
	                    	# logo {"" = ""}
	                    	# os = ""
	                    	# screenshot_list {"" = ""}
	                    	# support = ""
                    # }
	                	# description = ""
	                	# display_name = ""
	                	# enablevnc = false
	                	# images = []
	                	# interfaces = []
	                	# module = {
	                    	# # optional
	                    	# environment {"" = ""}
	                    	# module_type = ""
	                    	# routes {"" = ""}
	                    	# twin_detail = ""
                    # }
	                	# name = ""
	                	# owner = {
	                    	# # optional
	                    	# company = ""
	                    	# email = ""
	                    	# group = ""
	                    	# user = ""
	                    	# website = ""
                    # }
	                	# permissions = []
	                	# resources = []
	                	# vmmode = ""
                # }
	            	# memory = 0
	            	# name = ""
	            	# networks = 0
	            	# origin_type = ""
	            	# title = ""
	            	# # optional
	            	# app_id = ""
	            	# app_version = ""
	            	# description = ""
	            	# interfaces = []
	            	# name_app_part = ""
	            	# name_project_part = ""
	            	# naming_scheme = ""
	            	# parent_detail = {
	                	# # computed
                    # # id_of_parent_object =
                    # # version_of_parent_object =
	                	# # required
	                	# id_of_parent_object = ""
	                	# # optional
	                	# reference_exists = false
	                	# update_available = false
                # }
	            	# start_delay_in_seconds = 0
	            	# storage = 0
	            	# tags {"" = ""}
            # }
	        	# labels {"" = ""}
	        	# metrics = {
	            	# # optional
	            	# queries {"" = ""}
	            	# results {"" = ""}
            # }
	        	# routes {"" = ""}
	        	# target_condition = ""
	        	# target_condition_new {"" = ""}
        # }
		#     network_policy = {
	        	# # required
	        	# net_instance_config = []
        # }
		#     status_message = ""
    # }
		# local_operator_console_policy = {
		#     # computed
        # # id =
        # # revision =
        # # status =

		#     # required
		#     name = ""
		#     title = ""
		#     type = ""

		#     # optional
		#     app_policy = {
	        	# # required
	        	# apps = []
        # }
		#     attestation_policy = {
	        	# # computed
            # # id =
	        	# # required
	        	# type = ""
        # }
		#     attr {"" = ""}
		#     azure_policy = {
	        	# # required
	        	# app_id = ""
	        	# app_password = ""
	        	# tenant_id = ""

	        	# # optional
	        	# azure_resource_and_services = {
	            	# # optional
	            	# s_k_u = {
	                	# # optional
	                	# capacity = ""
	                	# name = ""
	                	# tier = ""
                # }
	            	# create_by_default = false
	            	# name = ""
	            	# region = ""
	            	# resource_group_name = ""
	            	# subscription_id = ""
            # }
	        	# certificate = {
	            	# # optional
	            	# basic_contraints_valid = false
	            	# cert = ""
	            	# crypto_key = ""
	            	# ecdsa_encryption = {
	                	# # optional
	                	# curve = ""
                # }
	            	# encrypted_secrets {"" = ""}
	            	# exportable = false
	            	# extended_key_usage = [""]
	            	# issuer = {
	                	# # optional
	                	# common_name = ""
	                	# country = [""]
	                	# locality = [""]
	                	# organization = [""]
	                	# organizational_unit = [""]
	                	# postal_code = [""]
	                	# province = [""]
	                	# serial_number = ""
                # }
	            	# key_usage = 0
	            	# pass_phrase = ""
	            	# public_key = ""
	            	# public_key_algorithm = ""
	            	# pvt_key = ""
	            	# reuse_key = false
	            	# rsa_ecryption = {
	                	# # optional
	                	# rsa_bits = ""
                # }
	            	# san_values = {
	                	# # optional
	                	# dns = [""]
	                	# emaild_ids = [""]
	                	# hosts = [""]
	                	# ips = [""]
	                	# upns = [""]
	                	# uris = [""]
                # }
	            	# serial_number = ""
	            	# signature_algorithm = ""
	            	# subject = {
	                	# # optional
	                	# dns = [""]
	                	# emaild_ids = [""]
	                	# hosts = [""]
	                	# ips = [""]
	                	# upns = [""]
	                	# uris = [""]
                # }
	            	# valid_from = 0.0
	            	# valid_till = 0.0
            # }
	        	# crypto_key = ""
	        	# custom_deployment_managed = false
	        	# encrypted_secrets {"" = ""}
        # }
		#     cluster_policy ={
        		# # required
        		# app_policy_id = ""
        		# network_policy_id = ""
        		# type = ""
        		# # optional
        		# cluster_config = {
	            	# # optional
	            	# min_nodes_required = 0
            # }
        # }
		#     description = ""
		#     edgeview_policy = {
	        	# # required
	        	# edgeview_allow = false
	        	# # optional
	        	# access_allow_change = false
	        	# edgeviewcfg = # EdgeviewCfg
	        	# max_expire_sec = 0
	        	# max_inst = 0
        # }
		#     local_operator_console_policy = {
        		# # required
        		# loc_url = ""
        # }
		#     module_policy =  {
	        	# # required
	        	# apps = []
	        	# priority = 0
	        	# # optional
	        	# etag = ""
	        	# azure_edge_agent = {
	            	# # computed
                # # drives =
                # # id =
	            	# # required
	            	# cpus = 0
	            	# drives = 0
	            	# manifest_json = {
	                	# # optional
	                	# ac_kind = ""
	                	# ac_version = ""
	                	# app_type = ""
	                	# configuration = {
	                    	# # optional
	                    	# custom_config = {
	                        	# # optional
	                        	# add = false
	                        	# allow_storage_resize = false
	                        	# field_delimiter = ""
	                        	# name = ""
	                        	# override = false
	                        	# template = ""
	                        	# variable_groups = []
                        # }
                    # }
	                	# container_detail = {
	                    	# # optional
	                    	# container_create_option = ""
                    # }
	                	# cpu_pinning_enabled = false
	                	# deployment_type = ""
	                	# desc = {
	                    	# # required
	                    	# app_category = ""
	                    	# # optional
	                    	# agreement_list {"" = ""}
	                    	# category"All"
	                    	# license_list {"" = ""}
	                    	# logo {"" = ""}
	                    	# os = ""
	                    	# screenshot_list {"" = ""}
	                    	# support = ""
                    # }
	                	# description = ""
	                	# display_name = ""
	                	# enablevnc = false
	                	# images = []
	                	# interfaces = []
	                	# module = {
	                    	# # optional
	                    	# environment {"" = ""}
	                    	# module_type = ""
	                    	# routes {"" = ""}
	                    	# twin_detail = ""
                    # }
	                	# name = ""
	                	# owner = {
	                    	# # optional
	                    	# company = ""
	                    	# email = ""
	                    	# group = ""
	                    	# user = ""
	                    	# website = ""
                    # }
	                	# permissions = []
	                	# resources = []
	                	# vmmode = ""
                # }
	            	# memory = 0
	            	# name = ""
	            	# networks = 0
	            	# origin_type = ""
	            	# title = ""
	            	# # optional
	            	# app_id = ""
	            	# app_version = ""
	            	# description = ""
	            	# interfaces = []
	            	# name_app_part = ""
	            	# name_project_part = ""
	            	# naming_scheme = ""
	            	# parent_detail = {
	                	# # computed
                    # # id_of_parent_object =
                    # # version_of_parent_object =
	                	# # required
	                	# id_of_parent_object = ""
	                	# # optional
	                	# reference_exists = false
	                	# update_available = false
                # }
	            	# start_delay_in_seconds = 0
	            	# storage = 0
	            	# tags {"" = ""}
            # }
	        	# azure_edge_hub = {
	            	# # computed
                # # drives =
                # # id =
	            	# # required
	            	# cpus = 0
	            	# drives = 0
	            	# manifest_json = {
	                	# # optional
	                	# ac_kind = ""
	                	# ac_version = ""
	                	# app_type = ""
	                	# configuration = {
	                    	# # optional
	                    	# custom_config = {
	                        	# # optional
	                        	# add = false
	                        	# allow_storage_resize = false
	                        	# field_delimiter = ""
	                        	# name = ""
	                        	# override = false
	                        	# template = ""
	                        	# variable_groups = []
                        # }
                    # }
	                	# container_detail = {
	                    	# # optional
	                    	# container_create_option = ""
                    # }
	                	# cpu_pinning_enabled = false
	                	# deployment_type = ""
	                	# desc = {
	                    	# # required
	                    	# app_category = ""
	                    	# # optional
	                    	# agreement_list {"" = ""}
	                    	# category"All"
	                    	# license_list {"" = ""}
	                    	# logo {"" = ""}
	                    	# os = ""
	                    	# screenshot_list {"" = ""}
	                    	# support = ""
                    # }
	                	# description = ""
	                	# display_name = ""
	                	# enablevnc = false
	                	# images = []
	                	# interfaces = []
	                	# module = {
	                    	# # optional
	                    	# environment {"" = ""}
	                    	# module_type = ""
	                    	# routes {"" = ""}
	                    	# twin_detail = ""
                    # }
	                	# name = ""
	                	# owner = {
	                    	# # optional
	                    	# company = ""
	                    	# email = ""
	                    	# group = ""
	                    	# user = ""
	                    	# website = ""
                    # }
	                	# permissions = []
	                	# resources = []
	                	# vmmode = ""
                # }
	            	# memory = 0
	            	# name = ""
	            	# networks = 0
	            	# origin_type = ""
	            	# title = ""
	            	# # optional
	            	# app_id = ""
	            	# app_version = ""
	            	# description = ""
	            	# interfaces = []
	            	# name_app_part = ""
	            	# name_project_part = ""
	            	# naming_scheme = ""
	            	# parent_detail = {
	                	# # computed
                    # # id_of_parent_object =
                    # # version_of_parent_object =
	                	# # required
	                	# id_of_parent_object = ""
	                	# # optional
	                	# reference_exists = false
	                	# update_available = false
                # }
	            	# start_delay_in_seconds = 0
	            	# storage = 0
	            	# tags {"" = ""}
            # }
	        	# labels {"" = ""}
	        	# metrics = {
	            	# # optional
	            	# queries {"" = ""}
	            	# results {"" = ""}
            # }
	        	# routes {"" = ""}
	        	# target_condition = ""
	        	# target_condition_new {"" = ""}
        # }
		#     network_policy = {
	        	# # required
	        	# net_instance_config = []
        # }
		#     status_message = ""
    # }
		# network_policy = {
		#     # computed
        # # id =
        # # revision =
        # # status =

		#     # required
		#     name = ""
		#     title = ""
		#     type = ""

		#     # optional
		#     app_policy = {
	        	# # required
	        	# apps = []
        # }
		#     attestation_policy = {
	        	# # computed
            # # id =
	        	# # required
	        	# type = ""
        # }
		#     attr {"" = ""}
		#     azure_policy = {
	        	# # required
	        	# app_id = ""
	        	# app_password = ""
	        	# tenant_id = ""

	        	# # optional
	        	# azure_resource_and_services = {
	            	# # optional
	            	# s_k_u = {
	                	# # optional
	                	# capacity = ""
	                	# name = ""
	                	# tier = ""
                # }
	            	# create_by_default = false
	            	# name = ""
	            	# region = ""
	            	# resource_group_name = ""
	            	# subscription_id = ""
            # }
	        	# certificate = {
	            	# # optional
	            	# basic_contraints_valid = false
	            	# cert = ""
	            	# crypto_key = ""
	            	# ecdsa_encryption = {
	                	# # optional
	                	# curve = ""
                # }
	            	# encrypted_secrets {"" = ""}
	            	# exportable = false
	            	# extended_key_usage = [""]
	            	# issuer = {
	                	# # optional
	                	# common_name = ""
	                	# country = [""]
	                	# locality = [""]
	                	# organization = [""]
	                	# organizational_unit = [""]
	                	# postal_code = [""]
	                	# province = [""]
	                	# serial_number = ""
                # }
	            	# key_usage = 0
	            	# pass_phrase = ""
	            	# public_key = ""
	            	# public_key_algorithm = ""
	            	# pvt_key = ""
	            	# reuse_key = false
	            	# rsa_ecryption = {
	                	# # optional
	                	# rsa_bits = ""
                # }
	            	# san_values = {
	                	# # optional
	                	# dns = [""]
	                	# emaild_ids = [""]
	                	# hosts = [""]
	                	# ips = [""]
	                	# upns = [""]
	                	# uris = [""]
                # }
	            	# serial_number = ""
	            	# signature_algorithm = ""
	            	# subject = {
	                	# # optional
	                	# dns = [""]
	                	# emaild_ids = [""]
	                	# hosts = [""]
	                	# ips = [""]
	                	# upns = [""]
	                	# uris = [""]
                # }
	            	# valid_from = 0.0
	            	# valid_till = 0.0
            # }
	        	# crypto_key = ""
	        	# custom_deployment_managed = false
	        	# encrypted_secrets {"" = ""}
        # }
		#     cluster_policy ={
        		# # required
        		# app_policy_id = ""
        		# network_policy_id = ""
        		# type = ""
        		# # optional
        		# cluster_config = {
	            	# # optional
	            	# min_nodes_required = 0
            # }
        # }
		#     description = ""
		#     edgeview_policy = {
	        	# # required
	        	# edgeview_allow = false
	        	# # optional
	        	# access_allow_change = false
	        	# edgeviewcfg = # EdgeviewCfg
	        	# max_expire_sec = 0
	        	# max_inst = 0
        # }
		#     local_operator_console_policy = {
        		# # required
        		# loc_url = ""
        # }
		#     module_policy =  {
	        	# # required
	        	# apps = []
	        	# priority = 0
	        	# # optional
	        	# etag = ""
	        	# azure_edge_agent = {
	            	# # computed
                # # drives =
                # # id =
	            	# # required
	            	# cpus = 0
	            	# drives = 0
	            	# manifest_json = {
	                	# # optional
	                	# ac_kind = ""
	                	# ac_version = ""
	                	# app_type = ""
	                	# configuration = {
	                    	# # optional
	                    	# custom_config = {
	                        	# # optional
	                        	# add = false
	                        	# allow_storage_resize = false
	                        	# field_delimiter = ""
	                        	# name = ""
	                        	# override = false
	                        	# template = ""
	                        	# variable_groups = []
                        # }
                    # }
	                	# container_detail = {
	                    	# # optional
	                    	# container_create_option = ""
                    # }
	                	# cpu_pinning_enabled = false
	                	# deployment_type = ""
	                	# desc = {
	                    	# # required
	                    	# app_category = ""
	                    	# # optional
	                    	# agreement_list {"" = ""}
	                    	# category"All"
	                    	# license_list {"" = ""}
	                    	# logo {"" = ""}
	                    	# os = ""
	                    	# screenshot_list {"" = ""}
	                    	# support = ""
                    # }
	                	# description = ""
	                	# display_name = ""
	                	# enablevnc = false
	                	# images = []
	                	# interfaces = []
	                	# module = {
	                    	# # optional
	                    	# environment {"" = ""}
	                    	# module_type = ""
	                    	# routes {"" = ""}
	                    	# twin_detail = ""
                    # }
	                	# name = ""
	                	# owner = {
	                    	# # optional
	                    	# company = ""
	                    	# email = ""
	                    	# group = ""
	                    	# user = ""
	                    	# website = ""
                    # }
	                	# permissions = []
	                	# resources = []
	                	# vmmode = ""
                # }
	            	# memory = 0
	            	# name = ""
	            	# networks = 0
	            	# origin_type = ""
	            	# title = ""
	            	# # optional
	            	# app_id = ""
	            	# app_version = ""
	            	# description = ""
	            	# interfaces = []
	            	# name_app_part = ""
	            	# name_project_part = ""
	            	# naming_scheme = ""
	            	# parent_detail = {
	                	# # computed
                    # # id_of_parent_object =
                    # # version_of_parent_object =
	                	# # required
	                	# id_of_parent_object = ""
	                	# # optional
	                	# reference_exists = false
	                	# update_available = false
                # }
	            	# start_delay_in_seconds = 0
	            	# storage = 0
	            	# tags {"" = ""}
            # }
	        	# azure_edge_hub = {
	            	# # computed
                # # drives =
                # # id =
	            	# # required
	            	# cpus = 0
	            	# drives = 0
	            	# manifest_json = {
	                	# # optional
	                	# ac_kind = ""
	                	# ac_version = ""
	                	# app_type = ""
	                	# configuration = {
	                    	# # optional
	                    	# custom_config = {
	                        	# # optional
	                        	# add = false
	                        	# allow_storage_resize = false
	                        	# field_delimiter = ""
	                        	# name = ""
	                        	# override = false
	                        	# template = ""
	                        	# variable_groups = []
                        # }
                    # }
	                	# container_detail = {
	                    	# # optional
	                    	# container_create_option = ""
                    # }
	                	# cpu_pinning_enabled = false
	                	# deployment_type = ""
	                	# desc = {
	                    	# # required
	                    	# app_category = ""
	                    	# # optional
	                    	# agreement_list {"" = ""}
	                    	# category"All"
	                    	# license_list {"" = ""}
	                    	# logo {"" = ""}
	                    	# os = ""
	                    	# screenshot_list {"" = ""}
	                    	# support = ""
                    # }
	                	# description = ""
	                	# display_name = ""
	                	# enablevnc = false
	                	# images = []
	                	# interfaces = []
	                	# module = {
	                    	# # optional
	                    	# environment {"" = ""}
	                    	# module_type = ""
	                    	# routes {"" = ""}
	                    	# twin_detail = ""
                    # }
	                	# name = ""
	                	# owner = {
	                    	# # optional
	                    	# company = ""
	                    	# email = ""
	                    	# group = ""
	                    	# user = ""
	                    	# website = ""
                    # }
	                	# permissions = []
	                	# resources = []
	                	# vmmode = ""
                # }
	            	# memory = 0
	            	# name = ""
	            	# networks = 0
	            	# origin_type = ""
	            	# title = ""
	            	# # optional
	            	# app_id = ""
	            	# app_version = ""
	            	# description = ""
	            	# interfaces = []
	            	# name_app_part = ""
	            	# name_project_part = ""
	            	# naming_scheme = ""
	            	# parent_detail = {
	                	# # computed
                    # # id_of_parent_object =
                    # # version_of_parent_object =
	                	# # required
	                	# id_of_parent_object = ""
	                	# # optional
	                	# reference_exists = false
	                	# update_available = false
                # }
	            	# start_delay_in_seconds = 0
	            	# storage = 0
	            	# tags {"" = ""}
            # }
	        	# labels {"" = ""}
	        	# metrics = {
	            	# # optional
	            	# queries {"" = ""}
	            	# results {"" = ""}
            # }
	        	# routes {"" = ""}
	        	# target_condition = ""
	        	# target_condition_new {"" = ""}
        # }
		#     network_policy = {
	        	# # required
	        	# net_instance_config = []
        # }
		#     status_message = ""
    # }
}
