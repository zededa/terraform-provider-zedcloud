// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0
//

// export TF_VAR_zedcloud_token=
variable "zedcloud_token" {
    description = "ZEDCloud token"
    sensitive = true
    type        = string
}

resource "zedcloud_edgenode" "required_only" {
		name = "required_only"
		model_id = "2f716b55-2639-486c-9a2f-55a2e94146a6"
		project_id = "4754cd0f-82d7-4e06-a68f-ff9e23e75ccf"
		title = "required_only-title"
}
