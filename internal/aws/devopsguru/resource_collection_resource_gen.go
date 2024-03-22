// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package devopsguru

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
	registry.AddResourceFactory("awscc_devopsguru_resource_collection", resourceCollectionResource)
}

// resourceCollectionResource returns the Terraform awscc_devopsguru_resource_collection resource.
// This Terraform resource corresponds to the CloudFormation AWS::DevOpsGuru::ResourceCollection resource.
func resourceCollectionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ResourceCollectionFilter
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Information about a filter used to specify which AWS resources are analyzed for anomalous behavior by DevOps Guru.",
		//	  "properties": {
		//	    "CloudFormation": {
		//	      "additionalProperties": false,
		//	      "description": "CloudFormation resource for DevOps Guru to monitor",
		//	      "properties": {
		//	        "StackNames": {
		//	          "description": "An array of CloudFormation stack names.",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "pattern": "^[a-zA-Z*]+[a-zA-Z0-9-]*$",
		//	            "type": "string"
		//	          },
		//	          "maxItems": 1000,
		//	          "minItems": 1,
		//	          "type": "array"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "Tags": {
		//	      "description": "Tagged resources for DevOps Guru to monitor",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Tagged resource for DevOps Guru to monitor",
		//	        "properties": {
		//	          "AppBoundaryKey": {
		//	            "description": "A Tag key for DevOps Guru app boundary.",
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "TagValues": {
		//	            "description": "Tag values of DevOps Guru app boundary.",
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "maxLength": 256,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            },
		//	            "maxItems": 1000,
		//	            "minItems": 1,
		//	            "type": "array"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"resource_collection_filter": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CloudFormation
				"cloudformation": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: StackNames
						"stack_names": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "An array of CloudFormation stack names.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.List{ /*START VALIDATORS*/
								listvalidator.SizeBetween(1, 1000),
								listvalidator.ValueStringsAre(
									stringvalidator.LengthBetween(1, 128),
									stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z*]+[a-zA-Z0-9-]*$"), ""),
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								generic.Multiset(),
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "CloudFormation resource for DevOps Guru to monitor",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Tags
				"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: AppBoundaryKey
							"app_boundary_key": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "A Tag key for DevOps Guru app boundary.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthBetween(1, 128),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: TagValues
							"tag_values": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "Tag values of DevOps Guru app boundary.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.List{ /*START VALIDATORS*/
									listvalidator.SizeBetween(1, 1000),
									listvalidator.ValueStringsAre(
										stringvalidator.LengthBetween(1, 256),
									),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
									generic.Multiset(),
									listplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "Tagged resources for DevOps Guru to monitor",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Information about a filter used to specify which AWS resources are analyzed for anomalous behavior by DevOps Guru.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceCollectionType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of ResourceCollection",
		//	  "enum": [
		//	    "AWS_CLOUD_FORMATION",
		//	    "AWS_TAGS"
		//	  ],
		//	  "type": "string"
		//	}
		"resource_collection_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of ResourceCollection",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
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
		Description: "This resource schema represents the ResourceCollection resource in the Amazon DevOps Guru.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DevOpsGuru::ResourceCollection").WithTerraformTypeName("awscc_devopsguru_resource_collection")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"app_boundary_key":           "AppBoundaryKey",
		"cloudformation":             "CloudFormation",
		"resource_collection_filter": "ResourceCollectionFilter",
		"resource_collection_type":   "ResourceCollectionType",
		"stack_names":                "StackNames",
		"tag_values":                 "TagValues",
		"tags":                       "Tags",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
