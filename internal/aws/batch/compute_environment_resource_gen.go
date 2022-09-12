// Code generated by generators/resource/main.go; DO NOT EDIT.

package batch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_batch_compute_environment", computeEnvironmentResourceType)
}

// computeEnvironmentResourceType returns the Terraform awscc_batch_compute_environment resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Batch::ComputeEnvironment resource type.
func computeEnvironmentResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"compute_environment_arn": {
			// Property: ComputeEnvironmentArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"compute_environment_name": {
			// Property: ComputeEnvironmentName
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"compute_resources": {
			// Property: ComputeResources
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "AllocationStrategy": {
			//       "type": "string"
			//     },
			//     "BidPercentage": {
			//       "type": "integer"
			//     },
			//     "DesiredvCpus": {
			//       "type": "integer"
			//     },
			//     "Ec2Configuration": {
			//       "insertionOrder": false,
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "ImageIdOverride": {
			//             "type": "string"
			//           },
			//           "ImageType": {
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "ImageType"
			//         ],
			//         "type": "object"
			//       },
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "Ec2KeyPair": {
			//       "type": "string"
			//     },
			//     "ImageId": {
			//       "type": "string"
			//     },
			//     "InstanceRole": {
			//       "type": "string"
			//     },
			//     "InstanceTypes": {
			//       "insertionOrder": false,
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "LaunchTemplate": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "LaunchTemplateId": {
			//           "type": "string"
			//         },
			//         "LaunchTemplateName": {
			//           "type": "string"
			//         },
			//         "Version": {
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "MaxvCpus": {
			//       "type": "integer"
			//     },
			//     "MinvCpus": {
			//       "type": "integer"
			//     },
			//     "PlacementGroup": {
			//       "type": "string"
			//     },
			//     "SecurityGroupIds": {
			//       "insertionOrder": false,
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "SpotIamFleetRole": {
			//       "type": "string"
			//     },
			//     "Subnets": {
			//       "insertionOrder": false,
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "Tags": {
			//       "additionalProperties": false,
			//       "description": "A key-value pair to associate with a resource.",
			//       "patternProperties": {
			//         "": {
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "Type": {
			//       "type": "string"
			//     },
			//     "UpdateToLatestImageVersion": {
			//       "default": false,
			//       "type": "boolean"
			//     }
			//   },
			//   "required": [
			//     "Subnets",
			//     "Type",
			//     "MaxvCpus"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"allocation_strategy": {
						// Property: AllocationStrategy
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"bid_percentage": {
						// Property: BidPercentage
						Type:     types.Int64Type,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"desiredv_cpus": {
						// Property: DesiredvCpus
						Type:     types.Int64Type,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"ec_2_configuration": {
						// Property: Ec2Configuration
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"image_id_override": {
									// Property: ImageIdOverride
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"image_type": {
									// Property: ImageType
									Type:     types.StringType,
									Required: true,
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
							resource.UseStateForUnknown(),
						},
					},
					"ec_2_key_pair": {
						// Property: Ec2KeyPair
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"image_id": {
						// Property: ImageId
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"instance_role": {
						// Property: InstanceRole
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"instance_types": {
						// Property: InstanceTypes
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
							resource.UseStateForUnknown(),
						},
					},
					"launch_template": {
						// Property: LaunchTemplate
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"launch_template_id": {
									// Property: LaunchTemplateId
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"launch_template_name": {
									// Property: LaunchTemplateName
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
								"version": {
									// Property: Version
									Type:     types.StringType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
									},
								},
							},
						),
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"maxv_cpus": {
						// Property: MaxvCpus
						Type:     types.Int64Type,
						Required: true,
					},
					"minv_cpus": {
						// Property: MinvCpus
						Type:     types.Int64Type,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"placement_group": {
						// Property: PlacementGroup
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"security_group_ids": {
						// Property: SecurityGroupIds
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
							resource.UseStateForUnknown(),
						},
					},
					"spot_iam_fleet_role": {
						// Property: SpotIamFleetRole
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
							resource.RequiresReplace(),
						},
					},
					"subnets": {
						// Property: Subnets
						Type:     types.ListType{ElemType: types.StringType},
						Required: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
						},
					},
					"tags": {
						// Property: Tags
						Description: "A key-value pair to associate with a resource.",
						// Pattern: ""
						Type:     types.MapType{ElemType: types.StringType},
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"type": {
						// Property: Type
						Type:     types.StringType,
						Required: true,
					},
					"update_to_latest_image_version": {
						// Property: UpdateToLatestImageVersion
						Type:     types.BoolType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							DefaultValue(types.Bool{Value: false}),
							resource.UseStateForUnknown(),
						},
						// UpdateToLatestImageVersion is a write-only property.
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"replace_compute_environment": {
			// Property: ReplaceComputeEnvironment
			// CloudFormation resource type schema:
			// {
			//   "default": true,
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				DefaultValue(types.Bool{Value: true}),
				resource.UseStateForUnknown(),
			},
			// ReplaceComputeEnvironment is a write-only property.
		},
		"service_role": {
			// Property: ServiceRole
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A key-value pair to associate with a resource.",
			//   "patternProperties": {
			//     "": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "A key-value pair to associate with a resource.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"unmanagedv_cpus": {
			// Property: UnmanagedvCpus
			// CloudFormation resource type schema:
			// {
			//   "type": "integer"
			// }
			Type:     types.Int64Type,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"update_policy": {
			// Property: UpdatePolicy
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "JobExecutionTimeoutMinutes": {
			//       "default": 30,
			//       "type": "integer"
			//     },
			//     "TerminateJobsOnUpdate": {
			//       "default": false,
			//       "type": "boolean"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"job_execution_timeout_minutes": {
						// Property: JobExecutionTimeoutMinutes
						Type:     types.Int64Type,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							DefaultValue(types.Int64{Value: 30}),
							resource.UseStateForUnknown(),
						},
					},
					"terminate_jobs_on_update": {
						// Property: TerminateJobsOnUpdate
						Type:     types.BoolType,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							DefaultValue(types.Bool{Value: false}),
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::Batch::ComputeEnvironment",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Batch::ComputeEnvironment").WithTerraformTypeName("awscc_batch_compute_environment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allocation_strategy":            "AllocationStrategy",
		"bid_percentage":                 "BidPercentage",
		"compute_environment_arn":        "ComputeEnvironmentArn",
		"compute_environment_name":       "ComputeEnvironmentName",
		"compute_resources":              "ComputeResources",
		"desiredv_cpus":                  "DesiredvCpus",
		"ec_2_configuration":             "Ec2Configuration",
		"ec_2_key_pair":                  "Ec2KeyPair",
		"image_id":                       "ImageId",
		"image_id_override":              "ImageIdOverride",
		"image_type":                     "ImageType",
		"instance_role":                  "InstanceRole",
		"instance_types":                 "InstanceTypes",
		"job_execution_timeout_minutes":  "JobExecutionTimeoutMinutes",
		"launch_template":                "LaunchTemplate",
		"launch_template_id":             "LaunchTemplateId",
		"launch_template_name":           "LaunchTemplateName",
		"maxv_cpus":                      "MaxvCpus",
		"minv_cpus":                      "MinvCpus",
		"placement_group":                "PlacementGroup",
		"replace_compute_environment":    "ReplaceComputeEnvironment",
		"security_group_ids":             "SecurityGroupIds",
		"service_role":                   "ServiceRole",
		"spot_iam_fleet_role":            "SpotIamFleetRole",
		"state":                          "State",
		"subnets":                        "Subnets",
		"tags":                           "Tags",
		"terminate_jobs_on_update":       "TerminateJobsOnUpdate",
		"type":                           "Type",
		"unmanagedv_cpus":                "UnmanagedvCpus",
		"update_policy":                  "UpdatePolicy",
		"update_to_latest_image_version": "UpdateToLatestImageVersion",
		"version":                        "Version",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/ComputeResources/UpdateToLatestImageVersion",
		"/properties/ReplaceComputeEnvironment",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
