resource "zedcloud_edgeapp" "ubuntu-all-ip" {
    name = "ubuntu-all-ip"
    title = "ubuntu-all-ip"
    description = "ubuntu-all-ip"
    manifest_file = "./TFTest-ubuntu-xenial-16.04.json"
    user_defined_version = "1.1"
    manifest_json = {
    		# optional
        name = "xenial-amd64-docker-20180725"
    		ac_kind = "VMManifest"
    		ac_version = "1.2.0"
    		app_type = "vm"
    		configuration = {
    		  custom_config = {
	          	add = false
	          	allow_storage_resize = false
	          	field_delimiter = ","
	          	name = "custom_config_name"
	          	override = false
	          	template = "dGVzdA=="
	          	variable_groups {
	              	condition {
	                  	# optional
	                  	name = "name"
	                  	operator = "operator"
	                  	value = "val"
                  }
	              	name = "var_group"
	              	required = false
	              	variables {
	                  	# required
	                  	format = ""
	                  	label = ""
	                  	name = ""
	                  	required = false
	                  	# optional
	                  	default = ""
	                  	encode = ""
	                  	max_length = ""
	                  	options {
	                      	# optional
	                      	label = ""
	                      	value = ""
                      }
	                  	process_input = ""
	                  	type = ""
	                  	value = ""
                  }
              }
          }
        }
    		# container_detail = {
      		# container_create_option = ""
        # }
    		cpu_pinning_enabled = false
    		deployment_type = "k3s"
    		desc = {
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
	        	imageformat = ""
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
                    portmapto = {
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
    		owner = {
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
