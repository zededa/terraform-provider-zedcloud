---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "zedcloud_user Data Source - terraform-provider-zedcloud"
subcategory: ""
description: |-
  
---

# zedcloud_user (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `email` (String) Email of the user
- `role_id` (String) Role associated with the user
- `username` (String) User defined name

### Optional

- `allowed_enterprises` (Block List) Permitted list of enterprises with their associated roles (see [below for nested schema](#nestedblock--allowed_enterprises))
- `custom_user_input` (Map of String) Custom user parameters
- `first_name` (String) First name of the user
- `full_name` (String) Full name of the user
- `hubspot_id` (String)
- `last_login_time` (String) Last login time of the user
- `last_logout_time` (String) Last logout time of the user
- `locale` (String) Locale of the user
- `notify_pref` (String) Notification preference of the user
- `phone` (String) Phone number of the user
- `sfdc_id` (String)
- `time_zone` (String) Preferred time zone of the user
- `type` (String) Type of the user

### Read-Only

- `email_state` (String) Email state
- `enterprise_id` (String) Origin enterprise of the user
- `id` (String) Unique system defined user ID
- `phone_state` (String) Phone state
- `revision` (List of Object) system defined info (see [below for nested schema](#nestedatt--revision))
- `state` (String) User state
- `totp_enabled` (Boolean) Is TOTP enrolment enabled

<a id="nestedblock--allowed_enterprises"></a>
### Nested Schema for `allowed_enterprises`

Optional:

- `name` (String)
- `role_id` (String)

Read-Only:

- `id` (String)


<a id="nestedatt--revision"></a>
### Nested Schema for `revision`

Read-Only:

- `created_at` (String)
- `created_by` (String)
- `curr` (String)
- `prev` (String)
- `updated_at` (String)
- `updated_by` (String)
