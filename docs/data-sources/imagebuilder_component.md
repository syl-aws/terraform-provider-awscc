---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_imagebuilder_component Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::ImageBuilder::Component
---

# awscc_imagebuilder_component (Data Source)

Data Source schema for AWS::ImageBuilder::Component



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **arn** (String) The Amazon Resource Name (ARN) of the component.
- **change_description** (String) The change description of the component.
- **data** (String) The data of the component.
- **description** (String) The description of the component.
- **encrypted** (Boolean) The encryption status of the component.
- **kms_key_id** (String) The KMS key identifier used to encrypt the component.
- **name** (String) The name of the component.
- **platform** (String) The platform of the component.
- **supported_os_versions** (List of String) The operating system (OS) version supported by the component.
- **tags** (Map of String) The tags associated with the component.
- **type** (String) The type of the component denotes whether the component is used to build the image or only to test it.
- **uri** (String) The uri of the component.
- **version** (String) The version of the component.

