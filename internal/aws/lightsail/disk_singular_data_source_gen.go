// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_lightsail_disk", diskDataSource)
}

// diskDataSource returns the Terraform awscc_lightsail_disk data source.
// This Terraform data source corresponds to the CloudFormation AWS::Lightsail::Disk resource.
func diskDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AddOns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of objects representing the add-ons to enable for the new instance.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A addon associate with a resource.",
		//	    "properties": {
		//	      "AddOnType": {
		//	        "description": "The add-on type",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "AutoSnapshotAddOnRequest": {
		//	        "additionalProperties": false,
		//	        "description": "An object that represents additional parameters when enabling or modifying the automatic snapshot add-on",
		//	        "properties": {
		//	          "SnapshotTimeOfDay": {
		//	            "description": "The daily time when an automatic snapshot will be created.",
		//	            "pattern": "^[0-9]{2}:00$",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "Status": {
		//	        "description": "Status of the Addon",
		//	        "enum": [
		//	          "Enabling",
		//	          "Disabling",
		//	          "Enabled",
		//	          "Terminating",
		//	          "Terminated",
		//	          "Disabled",
		//	          "Failed"
		//	        ],
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "AddOnType"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"add_ons": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: AddOnType
					"add_on_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The add-on type",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: AutoSnapshotAddOnRequest
					"auto_snapshot_add_on_request": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: SnapshotTimeOfDay
							"snapshot_time_of_day": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The daily time when an automatic snapshot will be created.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "An object that represents additional parameters when enabling or modifying the automatic snapshot add-on",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Status
					"status": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Status of the Addon",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of objects representing the add-ons to enable for the new instance.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AttachedTo
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of the attached Lightsail Instance",
		//	  "type": "string"
		//	}
		"attached_to": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of the attached Lightsail Instance",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AttachmentState
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Attachment State of the Lightsail disk",
		//	  "type": "string"
		//	}
		"attachment_state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Attachment State of the Lightsail disk",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AvailabilityZone
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Availability Zone in which to create your instance. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"availability_zone": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Availability Zone in which to create your instance. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DiskArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"disk_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: DiskName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The names to use for your new Lightsail disk.",
		//	  "maxLength": 254,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9][\\w\\-.]*[a-zA-Z0-9]$",
		//	  "type": "string"
		//	}
		"disk_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The names to use for your new Lightsail disk.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Iops
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Iops of the Lightsail disk",
		//	  "type": "integer"
		//	}
		"iops": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Iops of the Lightsail disk",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IsAttached
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Check is Disk is attached state",
		//	  "type": "boolean"
		//	}
		"is_attached": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Check is Disk is attached state",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Location
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Location of a resource.",
		//	  "properties": {
		//	    "AvailabilityZone": {
		//	      "description": "The Availability Zone in which to create your disk. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.",
		//	      "type": "string"
		//	    },
		//	    "RegionName": {
		//	      "description": "The Region Name in which to create your disk.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AvailabilityZone
				"availability_zone": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The Availability Zone in which to create your disk. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: RegionName
				"region_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The Region Name in which to create your disk.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Location of a resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Path
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Path of the  attached Disk",
		//	  "type": "string"
		//	}
		"path": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Path of the  attached Disk",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Resource type of Lightsail instance.",
		//	  "type": "string"
		//	}
		"resource_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Resource type of Lightsail instance.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SizeInGb
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Size of the Lightsail disk",
		//	  "type": "integer"
		//	}
		"size_in_gb": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Size of the Lightsail disk",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: State
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "State of the Lightsail disk",
		//	  "type": "string"
		//	}
		"state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "State of the Lightsail disk",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SupportCode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Support code to help identify any issues",
		//	  "type": "string"
		//	}
		"support_code": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Support code to help identify any issues",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Lightsail::Disk",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lightsail::Disk").WithTerraformTypeName("awscc_lightsail_disk")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"add_on_type":                  "AddOnType",
		"add_ons":                      "AddOns",
		"attached_to":                  "AttachedTo",
		"attachment_state":             "AttachmentState",
		"auto_snapshot_add_on_request": "AutoSnapshotAddOnRequest",
		"availability_zone":            "AvailabilityZone",
		"disk_arn":                     "DiskArn",
		"disk_name":                    "DiskName",
		"iops":                         "Iops",
		"is_attached":                  "IsAttached",
		"key":                          "Key",
		"location":                     "Location",
		"path":                         "Path",
		"region_name":                  "RegionName",
		"resource_type":                "ResourceType",
		"size_in_gb":                   "SizeInGb",
		"snapshot_time_of_day":         "SnapshotTimeOfDay",
		"state":                        "State",
		"status":                       "Status",
		"support_code":                 "SupportCode",
		"tags":                         "Tags",
		"value":                        "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
