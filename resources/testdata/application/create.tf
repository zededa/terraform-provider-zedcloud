resource "zedcloud_application" "test_tf_provider" {
    name = "ubuntu-all-ip"
    title = "ubuntu-all-ip"
    description = "ubuntu-all-ip"
    user_defined_version = "1.1"
    origin_type = "ORIGIN_LOCAL"
    manifest {
        # computed
    		# ac_kind = "VMManifest"

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
        }
    		# container_detail = {
      		# container_create_option = ""
        # }
    		cpu_pinning_enabled = false
    		deployment_type = "DEPLOYMENT_TYPE_K3S"
    		desc {
    	  	category = "Infrastructure"
    	  	os = "Zenix"
    	  	app_category = "APP_CATEGORY_OTHERS"
    	  	# agreement_list {"" = ""}
    	  	# license_list {"" = ""}
    	  	# logo {"" = ""}
    	  	# screenshot_list {"" = ""}
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
    		# module = {
    	  	# environment {"" = ""}
    	  	# module_type = ""
    	  	# routes {"" = ""}
    	  	# twin_detail = ""
        # }
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
