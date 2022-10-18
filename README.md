# Table Of Contents
1. [Terraform Customer Documentation](#Terraform-Customer-Documentation)
2. [Schema documentation](#Schema-documentation)
3. [Directory Structure](#Directory-Structure)
4. [Getting Started for Developers](#Getting-Started-for-Developers)
5. [Setting up Terraform to use terraform-provider-zededa for Developers](#Setting-up-Terraform-to-use-terraform-provider-zededa-for-Developers)
6. [Sample Terraform configuration files](#Sample-Terraform-configuration-files)

# Terraform Customer Documentation
For Customer documentation, see [Terraform Provider ZEDCloud](https://app.gitbook.com/@information-experience-zededa/s/zedcontrol/terraform-provider-zedcloud/terraform-provider-zedcloud)

# Schema documentation
1. [Provider schema](https://github.com/zededa/terraform-provider-zedcloud/blob/main/docs/index.md)
2. [Resources](https://github.com/zededa/terraform-provider-zedcloud/tree/main/docs/resources)
3. [Data Sources](https://github.com/zededa/terraform-provider-zedcloud/tree/main/docs/data-sources)

## Generating Schema documentation
Use the follong command from the terraform-provider-zedcloud directory to generate
schema documentation:

```
go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
```

# Directory Structure:
## provider
This directory has the actual code for the provider. It has the the following files:
1. data_source_\<object\>.go
    - Provides the read method for the object ( reads from zedcontrol )

2. resource_\<object\>.go
    - Provides create / update / delete methods for the object

3. provider.go
    - This file has the definition for the Provider, defining the data sources
      and the resources.

4. flattenutils.go
    - Utility routines to flatten API structures into TF Schema format.

5. resourcedatautils.go
    - Provides utility routines to convert TF Resource data into API structures.

## schemas
This directory has schemas for the objects. The same schema is used for data_sources
as well as the resources.

## tftest
This is a test directory. This currently has the Sample schema files and any other
files needed to test the provider.

# Getting Started for Developers
1. Check out github.com/zededa/zedcloud-api repo.
   1. Do a make in zedcloud-api (_ cd zededa/zedcloud-api; make _)
        1. This downloads the swagger files and generates client code.
        2. This needs to be done only once in your workspace, until there is a new
            version of the API
2. Do a make in the top level directory to generate the terraform-provider-zedcloud binary

# Setting up Terraform to use terraform-provider-zededa for Developers

If not already done so, install the terraform module on your machine. Instructions
can be found at https://terraform.io

To make terraform find and use terraform-provider-zededa, the following must be configured.

1. Set the following environment variables:
    1. TF_CLI_CONFIG_FILE  
        This must be set to a .tfrc file. A sample is provided in the toplevel
        directory (terraform-provider-zedcloud/terraform.tfrc).  
        Ex: export TF_CLI_CONFIG_FILE=<path-to-terraform-provider-zedcloud>/terraform.tfrc
    2. TF_LOG_PATH  
         Configures the path to the log file. All logging from the terraform module
         would be redirected to this file.  
         Ex: export TF_LOG_PATH=<path-to-terraform-provider-zedcloud>/tftest/terraform.log
    3. TF_LOG  
         Set this to the desired logging level. Ex: (export TF_LOG=TRACE)

# Sample Terraform configuration files
Sample terraform configuration files can be found in tftest directory.
1. onedevice_oneappinst.tf.sample
2. multinode_multiappinst.tf.sample

To try these out, copy them into a .tf file (for ex: main.tf). terraform automatically
find and process all .tf files in the module.
