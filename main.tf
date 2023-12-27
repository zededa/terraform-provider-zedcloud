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
  name  = "test_tf_provider"
  title = "title"

  # optional
  type = "TAG_TYPE_PROJECT"
  attestation_policy {
    # computed
    # id =
    # revision =
    # status =

    # required
    title = "title"
    type  = "POLICY_TYPE_ATTESTATION"

    # optional
    # app_policy = {
    # # required
    # apps = []
    # }
    attestation_policy {
      # computed
      # id =

      # required
      type = "ATTEST_POLICY_TYPE_ACCEPT"
    }
    # attr {
    # "attribue1" = "att-value1"
    # }
    # azure_policy = {
    # required
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
  }
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
  edgeview_policy {
    # computed
    # id =
    # revision =
    # status =

    #     # required
    #     name = ""
    #     title = ""
    #
    edgeview_policy {
      # required
      edgeview_allow = false
      # optional
      access_allow_change = true
      edgeviewcfg {
        # app_policy {
        # allow_app = true
        # }
        dev_policy {
          allow_dev = true
        }
        ext_policy {
          allow_ext = false
        }
        generation_id = 0
        jwt_info {
          allow_sec = 18000
          disp_url  = "zedcloud.local.zededa.net/api/v1/edge-view"
          encrypt   = false
          # expire_sec = ""
          num_inst = 1
        }
        # token = ""
      }
      max_expire_sec = 2592000
      max_inst       = 3
    }
    type = "POLICY_TYPE_EDGEVIEW"

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
  }
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

data "zedcloud_project" "test_tf_provider" {
  name  = "test_tf_provider"
  type  = "TAG_TYPE_PROJECT"
  title = "test_tf_provider"
  depends_on = [
    zedcloud_project.test_tf_provider
  ]
}

resource "zedcloud_edgenode" "test_tf_provider" {
  onboarding_key = "5d0767ee-0547-4569-b530-387e526f8cb9"
  serialno       = "2293dbe8-29ce-420c-8264-962857efc46b"
  # identity = "identity"

  # required
  name       = "test_tf_provider"
  model_id   = "ab66ab26-cbc2-4d66-829e-7bedb59aabba"
  project_id = data.zedcloud_project.test_tf_provider.id
  title      = "test_tf_provider-create_edgenode-title"

  admin_state = "ADMIN_STATE_ACTIVE"
  asset_id    = "asset_id"
  config_item {
    bool_value   = true
    float_value  = 1.0
    key          = "key"
    string_value = "string"
    uint32_value = 32
    uint64_value = 64
    value_type   = "value type"
  }
  # cpu = 2
  deployment_tag = "depl_tag"
  # deprecated = false
  description = "description"
  dev_location {
    city        = "berlin"
    country     = "germany"
    freeloc     = "freeloc"
    hostname    = "hostname"
    loc         = "52.520008, 13.404954"
    org         = "zededa"
    postal      = "10115"
    region      = "europe/west"
    underlay_ip = ""
  }
  generate_soft_serial = false
  tags = {
    "tag-key-1" = "tag-value-1"
  }
  token = "token"
}

data "zedcloud_edgenode" "test_tf_provider" {
  name     = "test_tf_provider"
  title    = "test_tf_provider"
  model_id = "ab66ab26-cbc2-4d66-829e-7bedb59aabba"
  depends_on = [
    zedcloud_edgenode.test_tf_provider
  ]
}

resource "zedcloud_application" "test_tf_provider" {
  name                 = "test_tf_provider"
  title                = "ubuntu-all-ip"
  description          = "ubuntu-all-ip"
  user_defined_version = "1.1"
  origin_type          = "ORIGIN_LOCAL"
  manifest {
    # computed
    ac_kind = "VMManifest"

    # optional
    name       = "xenial-amd64-docker-20180725"
    ac_version = "1.2.0"
    app_type   = "APP_TYPE_VM"
    configuration {
      custom_config {
        add                  = false
        allow_storage_resize = false
        field_delimiter      = ","
        name                 = "custom_config_name"
        override             = false
        template             = "dGVzdA=="
        variable_groups {
          condition {
            # optional
            name     = "condition"
            operator = "CONDITION_OPERATOR_EQUALTO"
            value    = "val"
          }
          name     = "var_group"
          required = false
        }
      }
    }
    cpu_pinning_enabled = false
    deployment_type     = "DEPLOYMENT_TYPE_K3S"
    desc {
      category     = "Infrastructure"
      os           = "Zenix"
      app_category = "APP_CATEGORY_OTHERS"
      support      = "support"
    }
    description  = "description"
    display_name = "display_name"
    enablevnc    = false
    # images {
    # # optional
    # cleartext = false
    # drvtype = "HDD"
    # ignorepurge = false
    # imageformat = "ISO"
    # imagename = "xenial-amd64-docker-20180725_1"
    # imageid = ""
    # maxsize = "1200000"
    # mountpath = ""
    # params {
    # # optional
    # name = "bootparam"
    # value = "+k/sre"
    # }
    # preserve = true
    # readonly = false
    # target = "Disk"
    # # volumelabel = "vol_1"
    # }
    interfaces {
      name         = "indirect"
      directattach = false
      privateip    = false
      acls {
        matches {
          type  = "protocol"
          value = "tcp"
        }
        matches {
          type  = "lport"
          value = "8022"
        }
        actions {
          portmap = true
          portmapto {
            app_port = 22
          }
        }
      }
      acls {
        matches {
          type  = "host"
          value = ""
        }
      }
    }
    owner {
      company = "Zededa Inc."
      email   = "test@zededa.com"
      group   = "testgroup"
      user    = "testuser"
      website = "http://www.zededa.com"
    }
    resources {
      name  = "cpus"
      value = "2"
    }
    resources {
      name  = "memory"
      value = "1024000"
    }
    vmmode = "HV_HVM"
  }
}

data "zedcloud_application" "test_tf_provider" {
  name  = "test_tf_provider"
  title = "test_tf_provider"
  depends_on = [
    zedcloud_application.test_tf_provider
  ]
}

resource "zedcloud_volume_instance" "test_tf_provider" {
  device_id   = data.zedcloud_edgenode.test_tf_provider.id
  name        = "test_tf_provider"
  title       = "test_title"
  description = "test_description"
  type        = "VOLUME_INSTANCE_TYPE_BLOCKSTORAGE"

  multiattach = true
  cleartext   = true
  accessmode  = "VOLUME_INSTANCE_ACCESS_MODE_READWRITE"
  size_bytes  = "1024"
  implicit    = false
  label       = "label"
  tags = {
    "tag-key-1" = "tag-value-1"
    "tag-key-2" = "tag-value-2"
  }
  depends_on = [
    data.zedcloud_edgenode.test_tf_provider
  ]
}

data "zedcloud_volume_instance" "test_tf_provider" {
  name      = "test_tf_provider"
  title     = "test_tf_provider"
  type      = "VOLUME_INSTANCE_TYPE_CONTENT_TREE"
  device_id = data.zedcloud_edgenode.test_tf_provider.id
  depends_on = [
    zedcloud_volume_instance.test_tf_provider
  ]
}

resource "zedcloud_network_instance" "test_tf_provider" {
  depends_on = [
    zedcloud_edgenode.test_tf_provider,
  ]
  # required
  device_id = data.zedcloud_edgenode.test_tf_provider.id
  name      = "test_tf_provider"
  title     = "title"
  kind      = "NETWORK_INSTANCE_KIND_LOCAL"
  port      = "eth1"

  # optional
  description    = "zedcloud_network_instance-complete-description"
  type           = "NETWORK_INSTANCE_DHCP_TYPE_V4"
  device_default = false
  dhcp           = false
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
  opaque {
    oconfig = "Test OConfig"
    type    = "OPAQUE_CONFIG_TYPE_UNSPECIFIED"
  }
  tags = {
    "ni-tag1" = "ni-tag-value-1"
    "ni-tag2" = "ni-tag-value-2"
  }
}

data "zedcloud_network_instance" "test_tf_provider" {
  name      = "test_tf_provider"
  title     = "test_tf_provider"
  kind      = "NETWORK_INSTANCE_KIND_LOCAL"
  device_id = data.zedcloud_edgenode.test_tf_provider.id
  depends_on = [
    zedcloud_network_instance.test_tf_provider
  ]
}

resource "zedcloud_application_instance" "test_tf_provider" {
  depends_on = [
    zedcloud_application.test_tf_provider,
    zedcloud_edgenode.test_tf_provider,
    zedcloud_network_instance.test_tf_provider
  ]

  # required
  name        = "test_tf_provider"
  title       = "tf test"
  description = "tf test"
  activate    = "true"
  app_id      = zedcloud_application.test_tf_provider.id
  device_id   = zedcloud_edgenode.test_tf_provider.id
  drives {
    # required
    drvtype   = "UNSPECIFIED"
    imagename = "ubuntu-tiny"
    maxsize   = "0"
    preserve  = false
    readonly  = false
    target    = "Disk"

    # optional
    cleartext   = false
    ignorepurge = false
    # imvolname = ""
    # mountpath = ""
    # mvolname = ""
    # volumelabel = ""
  }
  interfaces {
    directattach = false
    privateip    = false
    netinstname  = zedcloud_network_instance.test_tf_provider.name
    intfname     = "indirect"
    # io {
    # name = "adapter"
    # tags = {
    # "key" = "value"
    # }
    # type = "IO_TYPE_UNSPECIFIED"
    # }
    # acls {
    # id = 1
    #     matches {
    #         type = "protocol"
    #         value = "tcp"
    #     }
    #     matches {
    #         type = "lport"
    #         value = "8022"
    #     }
    #     actions {
    #         portmap = true
    #         mapparams {
    #             port = 22
    #         }
    #     }
    # }
    # acls {
    # id = 2
    #     matches {
    #         type = "host"
    #         value = ""
    #     }
    # }
    acls {
      #   id = 1
      actions {
        drop       = false
        limit      = false
        limitburst = 0
        limitrate  = 0
        portmap    = true

        mapparams {
          port = 22
        }
      }

      matches {
        type  = "protocol"
        value = "tcp"
      }
      matches {
        type  = "lport"
        value = "8022"
      }
    }
    acls {
      #   id = 2
      matches {
        type = "host"
      }
    }
  }

  # interfaces {
  # # required
  #     acls {
  #         matches {
  #             type = "protocol"
  #             value = "tcp"
  #         }
  #         actions {
  #             drop = true
  #             limit = true
  #             limitrate = 0
  #             limitunit = "unit"
  #             limitburst = 0
  #             portmap = true
  #             mapparams {
  #                 port = 0
  #             }
  #         }
  #         name = "tcp"
  #     }
  # directattach = false
  # eidregister {
  # # required
  # display_name = "display name"
  # e_id = "eID"
  # e_id_hash_len = 2
  # lisp_instance = 2
  # lisp_map_servers {
  # # required
  # credential = "123"
  # name_or_ip = "name"
  #         }
  # lisp_signature = "sig"
  # uuid = "1111"
  #     }
  #     intfname = "indirect"
  # intforder = 1
  # io {
  # name = "adapter"
  # tags = {
  #             "key" = "value"
  #         }
  # type = "IO_TYPE_UNSPECIFIED"
  #     }
  # 		ipaddr = "127.0.0.1"
  # 		macaddr = "00:00:00:00:00:00"
  #     netinstname = zedcloud_network_instance.test_tf_provider.name
  #     privateip = false

  # # optional
  # access_vlan_id = 0
  # default_net_instance = false
  # netinsttag = {
  #         "key" = "value"
  #     }
  # netname = "netname"
  # }

  # optional
  remote_console        = false
  is_secret_updated     = false
  collect_stats_ip_addr = "true"
  app_policy_id         = ""
  app_type              = "APP_TYPE_VM"
  cluster_id            = ""
  custom_config {
    add                  = false
    allow_storage_resize = false
    override             = false
    # field_delimiter = "###"
    # name = "custom_config_name"
    #     template = "I2Nsb3VkLWNvbmZpZwoKc3NoX3B3YXV0aDogWWVzCmhvc3RuYW1lOiBwb2N1c2VyCnVzZXJzOgogIC0gbmFtZTogcG9jdXNlcgogICAgc2hlbGw6IC9iaW4vYmFzaAogICAgc3VkbzogQUxMPShBTEwpIE5PUEFTU1dEOkFMTApjaHBhc3N3ZDoKICBsaXN0OiB8CiAgICAgIHBvY3VzZXI6cG9jdXNlcgogIGV4cGlyZTogZmFsc2UKCndyaXRlX2ZpbGVzOgogIC0gcGF0aDogL2hvbWUvcG9jdXNlci90ZXN0MQogICAgcGVybWlzc2lvbnM6ICcwNjQ0JwogICAgb3duZXI6IHBvY3VzZXI6cG9jdXNlcgogICAgZW5jb2Rpbmc6IGI2NAogICAgY29udGVudDogIyMjRklMRV9DT05URU5UIyMjCgpmaW5hbF9tZXNzYWdlOiAiVGhlIHN5c3RlbSBpcyBmaW5hbGx5IHVwLCBhZnRlciAkVVBUSU1FIHNlY29uZHMi"
    # variable_groups {
    # condition {
    # # optional
    # name = "condition"
    # operator = "CONDITION_OPERATOR_EQUALTO"
    # value = "val"
    #         }
    # name = "var_group"
    # required = false
    #         variables {
    #                 name = "FILE_CONTENT"
    #                 label = "Content to be written into the file"
    #                 max_length = ""
    #                 required = true
    #                 default = "Default content"
    #                 value = "Custom content"
    #                 format = "VARIABLE_FORMAT_TEXT"
    #                 encode = "FILE_ENCODING_UNSPECIFIED"
    #         }
    #     }
  }
  logs {
    access = false
  }
  manifest_info {
    # params = {
    # "key" = "value"
    # }
    transition_action = "INSTANCE_TA_NONE"
  }
  start_delay_in_seconds = 120
  user_data              = ""
  vminfo {
    # required
    cpus = 2
    mode = "HV_HVM"
    vnc  = false
    # optional
    # cpu_pinning_enabled = true
  }
  # not supported
  # purge {
  # counter = 2
  # }
  # refresh {
  # counter = 2
  # }
  # restart {
  # counter = 2
  # }
  tags = {
    "tag-key-1" = "tag-value-1"
  }
}

data "zedcloud_application_instance" "test_tf_provider" {
  depends_on = [
    zedcloud_application_instance.test_tf_provider
  ]
  name      = "test_tf_provider"
  title     = "test_tf_provider"
  device_id = data.zedcloud_edgenode.test_tf_provider.id
  app_id    = data.zedcloud_application.test_tf_provider.id
}


resource "zedcloud_image"  "test_tf_provider" {
  name = "test_tf_provider"
  title = "test_tf_provider"
  image_arch = "UNSPECIFIED"
  image_format = "RAW"
  datastore_id = "55c029bd-5d0a-46dd-b0ab-c6c9895b0d9a"
  image_rel_url = "test_url"
  image_size_bytes = 5000
  image_type = "IMAGE_TYPE_ARTIFACT"
  image_sha256 = "815fa14f544f2e955af18848667320fefd90d82e42497f429b53cc57cc76fb5c"
  project_access_list = []
}


resource "zedcloud_patch_envelope" "test_tf_provider" {
  name   = "test_tf_provider"
  action = "PATCH_ENVELOPE_ACTION_ACTIVATE"
  title  = "test_tf_provider"
  description = "test tf provider"
  user_defined_version = "1.0"

  artifacts {
    format = "OpaqueObjectCategoryInline"
    base64_artifact {
      base64_data      = "YXJ0aWZhY3QgZGF0YQ=="
      file_name_to_use = "testfile"
    }
  }
  artifacts {
    format = "OpaqueObjectCategoryExternalBinary"
    binary_artifact {
      image_id = resource.zedcloud_image.test_tf_provider.id
      image_name = resource.zedcloud_image.test_tf_provider.name
      file_name_to_use = "new_filename"
    }
  }
  project_name = data.zedcloud_project.test_tf_provider.name
  project_id = data.zedcloud_project.test_tf_provider.id
  depends_on = [
    zedcloud_image.test_tf_provider
    ]
}

resource "zedcloud_patch_reference_update" "test_tf_provider" {
  app_inst_id_list = [resource.zedcloud_application_instance.test_tf_provider.id]
  patchenvelope_id = resource.zedcloud_patch_envelope.test_tf_provider.id
  project_id = resource.zedcloud_project.test_tf_provider.id
  depends_on = [
    zedcloud_application_instance.test_tf_provider,
    zedcloud_patch_envelope.test_tf_provider,
    zedcloud_project.test_tf_provider
    ]

}

