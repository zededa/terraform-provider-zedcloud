resource "zedcloud_edgenode" "test_tf_provider" {
		name = "test_tf_provider"
		model_id = "2f716b55-2639-486c-9a2f-55a2e94146a6"
		title = "complete-title"
    project_id = "4754cd0f-82d7-4e06-a68f-ff9e23e75ccf"
}

data "zedcloud_edgenode" "test_tf_provider" {
		name = "test_tf_provider"
    depends_on = [
        zedcloud_edgenode.test_tf_provider
    ]
}

resource "zedcloud_network_instance" "test_tf_provider" {
    depends_on = [
        zedcloud_edgenode.test_tf_provider,
    ]
	  # required
    device_id = data.zedcloud_edgenode.test_tf_provider.id
    name = "test_tf_provider"
    title = "title"
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
    opaque {
        oconfig = "Test OConfig"
        type = "OPAQUE_CONFIG_TYPE_UNSPECIFIED"
    }
    tags = {
        "ni-tag1" = "ni-tag-value-1"
        "ni-tag2" = "ni-tag-value-2"
    }
}

data "zedcloud_network_instance" "test_tf_provider" {
		name = "test_tf_provider"
    depends_on = [
        zedcloud_network_instance.test_tf_provider
    ]
}

resource "zedcloud_application" "test_tf_provider" {
    name = "ubuntu-all-ip"
    title = "ubuntu-all-ip"
    description = "ubuntu-all-ip"
    user_defined_version = "1.1"
    origin_type = "ORIGIN_LOCAL"
    manifest {
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
    	  	category = "Infrastructure"
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

data "zedcloud_application" "test_tf_provider" {
		name = "test_tf_provider"
    depends_on = [
        zedcloud_application.test_tf_provider
    ]
}

resource "zedcloud_application_instance"  "test_tf_provider" {
    depends_on = [
        zedcloud_application.test_tf_provider,
        zedcloud_edgenode.test_tf_provider,
        zedcloud_network_instance.test_tf_provider,
    ]

		# required
    name = "test_tf_provider"
    title = "tf test"
    description = "tf test"
    activate = "true"
    app_id = zedcloud_application.test_tf_provider.id
    device_id = zedcloud_edgenode.test_tf_provider.id

    interfaces {
	    	# required
        acls {
            matches {
                type = "protocol"
                value = "tcp"
            }
            actions {
                drop = true
                limit = true
                limitrate = 0
                limitunit = "unit"
                limitburst = 0
                portmap = true
                mapparams {
                    port = 0
                }
            }
            name = "tcp"
        }
	    	directattach = false
	    	eidregister {
	        	# required
	          app_cert = "123"
	          app_private_key = "123"
	          app_public_key = "123"
	        	display_name = "display name"
	        	e_id = "eID"
	        	e_id_hash_len = 2
	        	lisp_instance = 2
	        	lisp_map_servers {
	            	# required
	            	credential = "123"
	            	name_or_ip = "name"
            }
	        	lisp_signature = "sig"
	        	uuid = "1111"
        }
        intfname = "indirect"
	    	intforder = 1
	    	io {
	        	name = "adapter"
	        	tags = {
                "key" = "value"
            }
	        	type = "IO_TYPE_UNSPECIFIED"
        }
    		ipaddr = "127.0.0.1"
    		macaddr = "00:00:00:00:00:00"
        netinstname = zedcloud_network_instance.test_tf_provider.name
        privateip = false

	    	# optional
	    	access_vlan_id = 0
	    	default_net_instance = false
	    	netinstid = "id"
	    	netinsttag = {
            "key" = "value"
        }
	    	netname = "netname"
    }

		drives {
	    	# required
	    	drvtype = "drvtype"
	    	imagename = "imagename"
	    	maxsize = 0
	    	preserve = false
	    	readonly = false
	    	target = "target"

	    	# optional
	    	cleartext = false
	    	ignorepurge = false
	    	imvolname = "imvolname"
	    	mountpath = "mountpath"
	    	mvolname = "mvolname"
	    	volumelabel = "volumelabel"
    }

		# optional
    remote_console = true
    is_secret_updated = false
    collect_stats_ip_addr = "true"
		app_policy_id = ""
      app_type = "APP_TYPE_VM"
		bundleversion = ""
		cluster_id = ""
    custom_config {
	    	add = false
	    	allow_storage_resize = true
	    	field_delimiter = "###"
	    	name = "custom_config_name"
	    	override = false
        template = "I2Nsb3VkLWNvbmZpZwoKc3NoX3B3YXV0aDogWWVzCmhvc3RuYW1lOiBwb2N1c2VyCnVzZXJzOgogIC0gbmFtZTogcG9jdXNlcgogICAgc2hlbGw6IC9iaW4vYmFzaAogICAgc3VkbzogQUxMPShBTEwpIE5PUEFTU1dEOkFMTApjaHBhc3N3ZDoKICBsaXN0OiB8CiAgICAgIHBvY3VzZXI6cG9jdXNlcgogIGV4cGlyZTogZmFsc2UKCndyaXRlX2ZpbGVzOgogIC0gcGF0aDogL2hvbWUvcG9jdXNlci90ZXN0MQogICAgcGVybWlzc2lvbnM6ICcwNjQ0JwogICAgb3duZXI6IHBvY3VzZXI6cG9jdXNlcgogICAgZW5jb2Rpbmc6IGI2NAogICAgY29udGVudDogIyMjRklMRV9DT05URU5UIyMjCgpmaW5hbF9tZXNzYWdlOiAiVGhlIHN5c3RlbSBpcyBmaW5hbGx5IHVwLCBhZnRlciAkVVBUSU1FIHNlY29uZHMi"
	    	variable_groups {
	        	condition {
	            	# optional
	            	name = "condition"
	            	operator = "CONDITION_OPERATOR_EQUALTO"
	            	value = "val"
            }
	        	name = "var_group"
	        	required = false
            variables {
                    name = "FILE_CONTENT"
                    label = "Content to be written into the file"
                    max_length = ""
                    required = true
                    default = "Default content"
                    value = "Custom content"
                    format = "VARIABLE_FORMAT_TEXT"
                    encode = "FILE_ENCODING_UNSPECIFIED"
            }
        }
    }
		deployment_type = ""
		logs {
		    access = false
		}
		manifest_info {
	    	bundle_version = "1.0.0"
	    	next_bundle_version = "1.0.1"
	    	params = {
            "key" = "value"
        }
	    	transition_action = "INSTANCE_TA_NONE"
		}
		start_delay_in_seconds = 120
		user_data = ""
		user_defined_version = "1.0.0"
		vminfo {
	    	# required
	    	cpus = 2
	    	memory = 1024
        mode = "HV_PV"
	    	vnc = false
	    	# optional
	    	cpu_pinning_enabled = true
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
		# tags = {
        # "tag-key-1" = "tag-value-1"
    # }
}
