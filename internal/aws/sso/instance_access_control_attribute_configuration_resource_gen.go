// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package sso

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_sso_instance_access_control_attribute_configuration", instanceAccessControlAttributeConfigurationResource)
}

// instanceAccessControlAttributeConfigurationResource returns the Terraform awscc_sso_instance_access_control_attribute_configuration resource.
// This Terraform resource corresponds to the CloudFormation AWS::SSO::InstanceAccessControlAttributeConfiguration resource.
func instanceAccessControlAttributeConfigurationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccessControlAttributes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "[\\p{L}\\p{Z}\\p{N}_.:\\/=+\\-@]+",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Source": {
		//	            "insertionOrder": true,
		//	            "items": {
		//	              "maxLength": 256,
		//	              "minLength": 0,
		//	              "pattern": "[\\p{L}\\p{Z}\\p{N}_.:\\/=+\\-@\\[\\]\\{\\}\\$\\\\\"]*",
		//	              "type": "string"
		//	            },
		//	            "maxItems": 1,
		//	            "type": "array"
		//	          }
		//	        },
		//	        "required": [
		//	          "Source"
		//	        ],
		//	        "type": "object"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array"
		//	}
		"access_control_attributes": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							stringvalidator.RegexMatches(regexp.MustCompile("[\\p{L}\\p{Z}\\p{N}_.:\\/=+\\-@]+"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Source
							"source": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Required:    true,
								Validators: []validator.List{ /*START VALIDATORS*/
									listvalidator.SizeAtMost(1),
									listvalidator.ValueStringsAre(
										stringvalidator.LengthBetween(0, 256),
										stringvalidator.RegexMatches(regexp.MustCompile("[\\p{L}\\p{Z}\\p{N}_.:\\/=+\\-@\\[\\]\\{\\}\\$\\\\\"]*"), ""),
									),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Required: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtMost(50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InstanceAccessControlAttributeConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The InstanceAccessControlAttributeConfiguration property has been deprecated but is still supported for backwards compatibility purposes. We recomend that you use  AccessControlAttributes property instead.",
		//	  "properties": {
		//	    "AccessControlAttributes": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Key": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "pattern": "[\\p{L}\\p{Z}\\p{N}_.:\\/=+\\-@]+",
		//	            "type": "string"
		//	          },
		//	          "Value": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "Source": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "maxLength": 256,
		//	                  "minLength": 0,
		//	                  "pattern": "[\\p{L}\\p{Z}\\p{N}_.:\\/=+\\-@\\[\\]\\{\\}\\$\\\\\"]*",
		//	                  "type": "string"
		//	                },
		//	                "maxItems": 1,
		//	                "type": "array"
		//	              }
		//	            },
		//	            "required": [
		//	              "Source"
		//	            ],
		//	            "type": "object"
		//	          }
		//	        },
		//	        "required": [
		//	          "Key",
		//	          "Value"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 50,
		//	      "type": "array"
		//	    }
		//	  },
		//	  "required": [
		//	    "AccessControlAttributes"
		//	  ],
		//	  "type": "object"
		//	}
		"instance_access_control_attribute_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AccessControlAttributes
				"access_control_attributes": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Key
							"key": schema.StringAttribute{ /*START ATTRIBUTE*/
								Required: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 128),
									stringvalidator.RegexMatches(regexp.MustCompile("[\\p{L}\\p{Z}\\p{N}_.:\\/=+\\-@]+"), ""),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: Value
							"value": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Source
									"source": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Required:    true,
										Validators: []validator.List{ /*START VALIDATORS*/
											listvalidator.SizeAtMost(1),
											listvalidator.ValueStringsAre(
												stringvalidator.LengthBetween(0, 256),
												stringvalidator.RegexMatches(regexp.MustCompile("[\\p{L}\\p{Z}\\p{N}_.:\\/=+\\-@\\[\\]\\{\\}\\$\\\\\"]*"), ""),
											),
										}, /*END VALIDATORS*/
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Required: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Required: true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.SizeAtMost(50),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The InstanceAccessControlAttributeConfiguration property has been deprecated but is still supported for backwards compatibility purposes. We recomend that you use  AccessControlAttributes property instead.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InstanceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the AWS SSO instance under which the operation will be executed.",
		//	  "maxLength": 1224,
		//	  "minLength": 10,
		//	  "pattern": "arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):sso:::instance/(sso)?ins-[a-zA-Z0-9-.]{16}",
		//	  "type": "string"
		//	}
		"instance_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the AWS SSO instance under which the operation will be executed.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(10, 1224),
				stringvalidator.RegexMatches(regexp.MustCompile("arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):sso:::instance/(sso)?ins-[a-zA-Z0-9-.]{16}"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource Type definition for SSO InstanceAccessControlAttributeConfiguration",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSO::InstanceAccessControlAttributeConfiguration").WithTerraformTypeName("awscc_sso_instance_access_control_attribute_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_control_attributes":                       "AccessControlAttributes",
		"instance_access_control_attribute_configuration": "InstanceAccessControlAttributeConfiguration",
		"instance_arn": "InstanceArn",
		"key":          "Key",
		"source":       "Source",
		"value":        "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
