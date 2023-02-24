resource "zedcloud_datastore"  "test_datastore" {
	# required
	ds_fqdn = "my-datastore.my-company.com"
	ds_path = "download/AMD64"
  ds_type = "DATASTORE_TYPE_AZUREBLOB"
	# ds_status = "DATASTORE_STATUS_ACTIVE"
  # need_clear_text = true
	name = "test"
	title = "title"

	# optional
	description = "description"
	region = "eu"

  # we cannot test for this, the API does not return any values
  # here so the plan command reports a diff after successfull test always
  # secret {
  #     api_key = "api_key"
  #     api_passwd =  "api_password"
  # }
}

# resource "zedcloud_datastore"  "test_datastore" {
# 	# required
# 	ds_fqdn = "my-datastore.my-company.com"
# 	ds_path = "download/AMD64"
#   ds_type = "DATASTORE_TYPE_AZUREBLOB"
# 	# ds_status = "DATASTORE_STATUS_ACTIVE"
# 	name = "test"
# 	title = "title"

# 	# optional
# 	description = "description"
# 	region = "eu"
#   secret {
#       api_key = var.api_key
#       api_passwd =  var.api_password
#   }
# }
