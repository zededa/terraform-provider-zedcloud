# Import by snapshot id. List an app instance's snapshots with the
# zedcloud_app_instance_snapshots data source. triggers can't be recovered:
# if the configuration sets them, the next apply takes a new snapshot unless
# you add lifecycle { ignore_changes = [triggers] }.
terraform import zedcloud_app_instance_snapshot.before_upgrade 3f0d3b9e-6a3c-4b8e-9d7e-2c1f8e4a5b6c
