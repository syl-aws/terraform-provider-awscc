// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package iam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_iam_role", roleResource)
}

// roleResource returns the Terraform awscc_iam_role resource.
// This Terraform resource corresponds to the CloudFormation AWS::IAM::Role resource.
func roleResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AssumeRolePolicyDocument
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The trust policy that is associated with this role. Trust policies define which entities can assume the role. You can associate only one trust policy with a role. For an example of a policy that can be used to assume a role, see [Template Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html#aws-resource-iam-role--examples). For more information about the elements that you can use in an IAM policy, see [Policy Elements Reference](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements.html) in the *User Guide*.",
		//	  "type": "string"
		//	}
		"assume_role_policy_document": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The trust policy that is associated with this role. Trust policies define which entities can assume the role. You can associate only one trust policy with a role. For an example of a policy that can be used to assume a role, see [Template Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html#aws-resource-iam-role--examples). For more information about the elements that you can use in an IAM policy, see [Policy Elements Reference](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements.html) in the *User Guide*.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description of the role that you provide.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description of the role that you provide.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ManagedPolicyArns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of Amazon Resource Names (ARNs) of the IAM managed policies that you want to attach to the role.\n For more information about ARNs, see [Amazon Resource Names (ARNs) and Service Namespaces](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the *General Reference*.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"managed_policy_arns": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of Amazon Resource Names (ARNs) of the IAM managed policies that you want to attach to the role.\n For more information about ARNs, see [Amazon Resource Names (ARNs) and Service Namespaces](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the *General Reference*.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MaxSessionDuration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default value of one hour is applied. This setting can have a value from 1 hour to 12 hours.\n Anyone who assumes the role from the CLI or API can use the ``DurationSeconds`` API parameter or the ``duration-seconds`` CLI parameter to request a longer session. The ``MaxSessionDuration`` setting determines the maximum duration that can be requested using the ``DurationSeconds`` parameter. If users don't specify a value for the ``DurationSeconds`` parameter, their security credentials are valid for one hour by default. This applies when you use the ``AssumeRole*`` API operations or the ``assume-role*`` CLI operations but does not apply when you use those operations to create a console URL. For more information, see [Using IAM roles](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html) in the *IAM User Guide*.",
		//	  "type": "integer"
		//	}
		"max_session_duration": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default value of one hour is applied. This setting can have a value from 1 hour to 12 hours.\n Anyone who assumes the role from the CLI or API can use the ``DurationSeconds`` API parameter or the ``duration-seconds`` CLI parameter to request a longer session. The ``MaxSessionDuration`` setting determines the maximum duration that can be requested using the ``DurationSeconds`` parameter. If users don't specify a value for the ``DurationSeconds`` parameter, their security credentials are valid for one hour by default. This applies when you use the ``AssumeRole*`` API operations or the ``assume-role*`` CLI operations but does not apply when you use those operations to create a console URL. For more information, see [Using IAM roles](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html) in the *IAM User Guide*.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Path
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "/",
		//	  "description": "The path to the role. For more information about paths, see [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the *IAM User Guide*.\n This parameter is optional. If it is not included, it defaults to a slash (/).\n This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of either a forward slash (/) by itself or a string that must begin and end with forward slashes. In addition, it can contain any ASCII character from the ! (``\\u0021``) through the DEL character (``\\u007F``), including most punctuation characters, digits, and upper and lowercased letters.",
		//	  "type": "string"
		//	}
		"path": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The path to the role. For more information about paths, see [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the *IAM User Guide*.\n This parameter is optional. If it is not included, it defaults to a slash (/).\n This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of either a forward slash (/) by itself or a string that must begin and end with forward slashes. In addition, it can contain any ASCII character from the ! (``\\u0021``) through the DEL character (``\\u007F``), including most punctuation characters, digits, and upper and lowercased letters.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				generic.StringDefaultValue("/"),
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PermissionsBoundary
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the policy used to set the permissions boundary for the role.\n For more information about permissions boundaries, see [Permissions boundaries for IAM identities](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html) in the *IAM User Guide*.",
		//	  "type": "string"
		//	}
		"permissions_boundary": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the policy used to set the permissions boundary for the role.\n For more information about permissions boundaries, see [Permissions boundaries for IAM identities](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html) in the *IAM User Guide*.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Policies
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Adds or updates an inline policy document that is embedded in the specified IAM role.\n When you embed an inline policy in a role, the inline policy is used as part of the role's access (permissions) policy. The role's trust policy is created at the same time as the role. You can update a role's trust policy later. For more information about IAM roles, go to [Using Roles to Delegate Permissions and Federate Identities](https://docs.aws.amazon.com/IAM/latest/UserGuide/roles-toplevel.html).\n A role can also have an attached managed policy. For information about policies, see [Managed Policies and Inline Policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html) in the *User Guide*.\n For information about limits on the number of inline policies that you can embed with a role, see [Limitations on Entities](https://docs.aws.amazon.com/IAM/latest/UserGuide/LimitationsOnEntities.html) in the *User Guide*.\n  If an external policy (such as ``AWS::IAM::Policy`` or",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Contains information about an attached policy.\n An attached policy is a managed policy that has been attached to a user, group, or role.\n For more information about managed policies, refer to [Managed Policies and Inline Policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html) in the *User Guide*.",
		//	    "properties": {
		//	      "PolicyDocument": {
		//	        "description": "The entire contents of the policy that defines permissions. For more information, see [Overview of JSON policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#access_policies-json).",
		//	        "type": "string"
		//	      },
		//	      "PolicyName": {
		//	        "description": "The friendly name (not ARN) identifying the policy.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "PolicyName",
		//	      "PolicyDocument"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"policies": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: PolicyDocument
					"policy_document": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The entire contents of the policy that defines permissions. For more information, see [Overview of JSON policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#access_policies-json).",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: PolicyName
					"policy_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The friendly name (not ARN) identifying the policy.",
						Required:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Adds or updates an inline policy document that is embedded in the specified IAM role.\n When you embed an inline policy in a role, the inline policy is used as part of the role's access (permissions) policy. The role's trust policy is created at the same time as the role. You can update a role's trust policy later. For more information about IAM roles, go to [Using Roles to Delegate Permissions and Federate Identities](https://docs.aws.amazon.com/IAM/latest/UserGuide/roles-toplevel.html).\n A role can also have an attached managed policy. For information about policies, see [Managed Policies and Inline Policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html) in the *User Guide*.\n For information about limits on the number of inline policies that you can embed with a role, see [Limitations on Entities](https://docs.aws.amazon.com/IAM/latest/UserGuide/LimitationsOnEntities.html) in the *User Guide*.\n  If an external policy (such as ``AWS::IAM::Policy`` or",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RoleId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"role_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RoleName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the IAM role, up to 64 characters in length. For valid values, see the ``RoleName`` parameter for the [CreateRole](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html) action in the *User Guide*.\n This parameter allows (per its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-. The role name must be unique within the account. Role names are not distinguished by case. For example, you cannot create roles named both \"Role1\" and \"role1\".\n If you don't specify a name, CFN generates a unique physical ID and uses that ID for the role name.\n If you specify a name, you must specify the ``CAPABILITY_NAMED_IAM`` value to acknowledge your template's capabilities. For more information, see [Acknowledging Resources in Templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/Use",
		//	  "type": "string"
		//	}
		"role_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the IAM role, up to 64 characters in length. For valid values, see the ``RoleName`` parameter for the [CreateRole](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html) action in the *User Guide*.\n This parameter allows (per its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-. The role name must be unique within the account. Role names are not distinguished by case. For example, you cannot create roles named both \"Role1\" and \"role1\".\n If you don't specify a name, CFN generates a unique physical ID and uses that ID for the role name.\n If you specify a name, you must specify the ``CAPABILITY_NAMED_IAM`` value to acknowledge your template's capabilities. For more information, see [Acknowledging Resources in Templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/Use",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of tags that are attached to the role. For more information about tagging, see [Tagging IAM resources](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the *IAM User Guide*.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A structure that represents user-provided metadata that can be associated with an IAM resource. For more information about tagging, see [Tagging IAM resources](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the *IAM User Guide*.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name that can be used to look up or retrieve the associated value. For example, ``Department`` or ``Cost Center`` are common choices.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value associated with this tag. For example, tags with a key name of ``Department`` could have values such as ``Human Resources``, ``Accounting``, and ``Support``. Tags with a key name of ``Cost Center`` might have values that consist of the number associated with the different cost centers in your company. Typically, many resources have tags with the same key name but with different values.\n   AWS always interprets the tag ``Value`` as a single string. If you need to store an array, you can store comma-separated values in the string. However, you must interpret the value in your code.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name that can be used to look up or retrieve the associated value. For example, ``Department`` or ``Cost Center`` are common choices.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value associated with this tag. For example, tags with a key name of ``Department`` could have values such as ``Human Resources``, ``Accounting``, and ``Support``. Tags with a key name of ``Cost Center`` might have values that consist of the number associated with the different cost centers in your company. Typically, many resources have tags with the same key name but with different values.\n   AWS always interprets the tag ``Value`` as a single string. If you need to store an array, you can store comma-separated values in the string. However, you must interpret the value in your code.",
						Required:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of tags that are attached to the role. For more information about tagging, see [Tagging IAM resources](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the *IAM User Guide*.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
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
		Description: "Creates a new role for your AWS-account.\n  For more information about roles, see [IAM roles](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html) in the *IAM User Guide*. For information about quotas for role names and the number of roles you can create, see [IAM and quotas](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html) in the *IAM User Guide*.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IAM::Role").WithTerraformTypeName("awscc_iam_role")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                         "Arn",
		"assume_role_policy_document": "AssumeRolePolicyDocument",
		"description":                 "Description",
		"key":                         "Key",
		"managed_policy_arns":         "ManagedPolicyArns",
		"max_session_duration":        "MaxSessionDuration",
		"path":                        "Path",
		"permissions_boundary":        "PermissionsBoundary",
		"policies":                    "Policies",
		"policy_document":             "PolicyDocument",
		"policy_name":                 "PolicyName",
		"role_id":                     "RoleId",
		"role_name":                   "RoleName",
		"tags":                        "Tags",
		"value":                       "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
