resource "zedcloud_project" "test_tf_provider" {
	# required
	name = "test_tf_provider-create_application"
	title = "title"

	# optional
	type = "TAG_TYPE_PROJECT"
	attestation_policy {
		# required
		title = "title"
		type = "POLICY_TYPE_ATTESTATION"

		attestation_policy {
			# required
			type = "ATTEST_POLICY_TYPE_ACCEPT"
		}
	}
	tag_level_settings {
		flow_log_transmission = "NETWORK_INSTANCE_FLOW_LOG_TRANSMISSION_UNSPECIFIED"
	}
}

resource "zedcloud_application" "test_tf_provider" {
	depends_on = [
		zedcloud_project.test_tf_provider
	]
	name = "test_tf_provider-ubuntu-all-ip"
	title = "test_tf_provider-ubuntu-all-ip"
	description = "test_tf_provider-ubuntu-all-ip"
	user_defined_version = "1.1"
	origin_type = "ORIGIN_LOCAL"
	project_access_list = [zedcloud_project.test_tf_provider.id]
	manifest {
		# computed
		ac_kind = "VMManifest"

		# optional
		name = "xenial-amd64-docker-20180725"
		ac_version = "1.2.0"
		app_type = "APP_TYPE_VM"
		configuration {
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
				}
			}
		}
		cpu_pinning_enabled = false
		deployment_type = "DEPLOYMENT_TYPE_K3S"
		desc {
			category = "APP_CATEGORY_OTHERS"
			os = "Zenix"
			app_category = "APP_CATEGORY_OTHERS"
			support = "support"
		}
		description = "description"
		display_name = "display_name"
		enablevnc = false
		images {
			# optional
			cleartext = false
			drvtype = "HDD"
			ignorepurge = false
			imageformat = "ISO"
			imagename = "test-xenial-amd64"
			imageid = ""
			maxsize = "1200000"
			mountpath = ""
			params {
				# optional
				name = "bootparam"
				value = "+k/sre"
			}
			preserve = true
			readonly = false
			target = "Disk"
			volumelabel = "vol_1"
		}
		interfaces {
			name = "indirect"
			directattach = false
			privateip = false
			acls {
				matches {
					type = "protocol"
					value = "tcp"
				}
				matches {
					type = "lport"
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
					type = "host"
					value = ""
				}
			}
		}
		owner {
			company = "Zededa Inc."
			email = "test@zededa.com"
			group = "testgroup"
			user = "testuser"
			website = "http://www.zededa.com"
		}
		resources {
			name = "cpus"
			value = "2"
		}
		resources {
			name = "memory"
			value = "1024000"
		}
		vmmode = "HV_HVM"
	}
}
