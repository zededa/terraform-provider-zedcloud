resource "zedcloud_datastore"  "test_datastore" {
	# required
	ds_fqdn = "my-datastore.my-company.com"
	ds_path = "download/AMD64"
	ds_type = "DATASTORE_TYPE_HTTPS"
	ds_status = "DATASTORE_STATUS_ACTIVE"
	name = "test"
	title = "title"

	# optional
	# api_key = "1234abcd"
	description = "description"
	need_clear_text = false
	region = "eu"
}
