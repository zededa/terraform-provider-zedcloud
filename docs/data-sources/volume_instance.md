---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "zedcloud_volume_instance Data Source - terraform-provider-zedcloud"
subcategory: ""
description: |-
  
---

# zedcloud_volume_instance (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `accessmode` (String) Access mode
- `cleartext` (Boolean) flag to keep the contents of the volume unencrypted (in clear text)
- `content_tree_id` (String) content tree ID
- `description` (String) Detailed description of the volume instance.
- `device_id` (String) id of the device on which volume instance is created
- `image` (String) name of the image
- `implicit` (Boolean) flag to create implicit volumes
- `label` (String) label
- `multiattach` (Boolean) flag to enable the volume to be attached to multiple app instances
- `name` (String) User defined name of the volume instance, unique across the enterprise. Once object is created, name can’t be changed.
- `project_id` (String) id of the project in which the volume instance is created
- `size_bytes` (String) size of volume
- `tags` (Map of String) Tags are name/value pairs that enable you to categorize resources. Tag names are case insensitive with max_length 512 and min_length 3. Tag values are case sensitive with max_length 256 and min_length 3.
- `title` (String) User defined title of the volume instance. Title can be changed at any time.
- `type` (String) type of Volume Instance

### Read-Only

- `id` (String) System defined universally unique Id of the volume instance.
- `revision` (List of Object) system defined Revision info of the object (see [below for nested schema](#nestedatt--revision))

<a id="nestedatt--revision"></a>
### Nested Schema for `revision`

Read-Only:

- `created_at` (String)
- `created_by` (String)
- `curr` (String)
- `prev` (String)
- `updated_at` (String)
- `updated_by` (String)

