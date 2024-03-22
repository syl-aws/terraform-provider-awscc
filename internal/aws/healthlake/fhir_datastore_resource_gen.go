// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package healthlake

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_healthlake_fhir_datastore", fHIRDatastoreResource)
}

// fHIRDatastoreResource returns the Terraform awscc_healthlake_fhir_datastore resource.
// This Terraform resource corresponds to the CloudFormation AWS::HealthLake::FHIRDatastore resource.
func fHIRDatastoreResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The time that a Data Store was created.",
		//	  "properties": {
		//	    "Nanos": {
		//	      "description": "Nanoseconds.",
		//	      "type": "integer"
		//	    },
		//	    "Seconds": {
		//	      "description": "Seconds since epoch.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Seconds",
		//	    "Nanos"
		//	  ],
		//	  "type": "object"
		//	}
		"created_at": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Nanos
				"nanos": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "Nanoseconds.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Seconds
				"seconds": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Seconds since epoch.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The time that a Data Store was created.",
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DatastoreArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name used in the creation of the Data Store.",
		//	  "pattern": "^arn:aws((-us-gov)|(-iso)|(-iso-b)|(-cn))?:healthlake:[a-zA-Z0-9-]+:[0-9]{12}:datastore/.+?",
		//	  "type": "string"
		//	}
		"datastore_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name used in the creation of the Data Store.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DatastoreEndpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AWS endpoint for the Data Store. Each Data Store will have it's own endpoint with Data Store ID in the endpoint URL.",
		//	  "maxLength": 10000,
		//	  "type": "string"
		//	}
		"datastore_endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AWS endpoint for the Data Store. Each Data Store will have it's own endpoint with Data Store ID in the endpoint URL.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DatastoreId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AWS-generated ID number for the Data Store.",
		//	  "maxLength": 32,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"datastore_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AWS-generated ID number for the Data Store.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DatastoreName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The user-generated name for the Data Store.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"datastore_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The user-generated name for the Data Store.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DatastoreStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of the Data Store. Possible statuses are 'CREATING', 'ACTIVE', 'DELETING', or 'DELETED'.",
		//	  "enum": [
		//	    "CREATING",
		//	    "ACTIVE",
		//	    "DELETING",
		//	    "DELETED"
		//	  ],
		//	  "type": "string"
		//	}
		"datastore_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of the Data Store. Possible statuses are 'CREATING', 'ACTIVE', 'DELETING', or 'DELETED'.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DatastoreTypeVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The FHIR version. Only R4 version data is supported.",
		//	  "enum": [
		//	    "R4"
		//	  ],
		//	  "type": "string"
		//	}
		"datastore_type_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The FHIR version. Only R4 version data is supported.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"R4",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IdentityProviderConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The identity provider configuration for the datastore",
		//	  "properties": {
		//	    "AuthorizationStrategy": {
		//	      "description": "Type of Authorization Strategy. The two types of supported Authorization strategies are SMART_ON_FHIR_V1 and AWS_AUTH.",
		//	      "enum": [
		//	        "SMART_ON_FHIR_V1",
		//	        "AWS_AUTH"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "FineGrainedAuthorizationEnabled": {
		//	      "description": "Flag to indicate if fine-grained authorization will be enabled for the datastore",
		//	      "type": "boolean"
		//	    },
		//	    "IdpLambdaArn": {
		//	      "description": "The Amazon Resource Name (ARN) of the Lambda function that will be used to decode the access token created by the authorization server.",
		//	      "maxLength": 256,
		//	      "minLength": 49,
		//	      "pattern": "arn:aws[-a-z]*:lambda:[a-z]{2}-[a-z]+-\\d{1}:\\d{12}:function:[a-zA-Z0-9\\-_\\.]+(:(\\$LATEST|[a-zA-Z0-9\\-_]+))?",
		//	      "type": "string"
		//	    },
		//	    "Metadata": {
		//	      "description": "The JSON metadata elements for identity provider configuration.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "AuthorizationStrategy"
		//	  ],
		//	  "type": "object"
		//	}
		"identity_provider_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AuthorizationStrategy
				"authorization_strategy": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Type of Authorization Strategy. The two types of supported Authorization strategies are SMART_ON_FHIR_V1 and AWS_AUTH.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"SMART_ON_FHIR_V1",
							"AWS_AUTH",
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: FineGrainedAuthorizationEnabled
				"fine_grained_authorization_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Flag to indicate if fine-grained authorization will be enabled for the datastore",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: IdpLambdaArn
				"idp_lambda_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The Amazon Resource Name (ARN) of the Lambda function that will be used to decode the access token created by the authorization server.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(49, 256),
						stringvalidator.RegexMatches(regexp.MustCompile("arn:aws[-a-z]*:lambda:[a-z]{2}-[a-z]+-\\d{1}:\\d{12}:function:[a-zA-Z0-9\\-_\\.]+(:(\\$LATEST|[a-zA-Z0-9\\-_]+))?"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Metadata
				"metadata": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The JSON metadata elements for identity provider configuration.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The identity provider configuration for the datastore",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PreloadDataConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The preloaded data configuration for the Data Store. Only data preloaded from Synthea is supported.",
		//	  "properties": {
		//	    "PreloadDataType": {
		//	      "description": "The type of preloaded data. Only Synthea preloaded data is supported.",
		//	      "enum": [
		//	        "SYNTHEA"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "PreloadDataType"
		//	  ],
		//	  "type": "object"
		//	}
		"preload_data_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: PreloadDataType
				"preload_data_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The type of preloaded data. Only Synthea preloaded data is supported.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"SYNTHEA",
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The preloaded data configuration for the Data Store. Only data preloaded from Synthea is supported.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SseConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The server-side encryption key configuration for a customer provided encryption key.",
		//	  "properties": {
		//	    "KmsEncryptionConfig": {
		//	      "additionalProperties": false,
		//	      "description": "The customer-managed-key (CMK) used when creating a Data Store. If a customer owned key is not specified, an AWS owned key will be used for encryption.",
		//	      "properties": {
		//	        "CmkType": {
		//	          "description": "The type of customer-managed-key (CMK) used for encryption. The two types of supported CMKs are customer owned CMKs and AWS owned CMKs.",
		//	          "enum": [
		//	            "CUSTOMER_MANAGED_KMS_KEY",
		//	            "AWS_OWNED_KMS_KEY"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "KmsKeyId": {
		//	          "description": "The KMS encryption key id/alias used to encrypt the Data Store contents at rest.",
		//	          "maxLength": 400,
		//	          "minLength": 1,
		//	          "pattern": "(arn:aws((-us-gov)|(-iso)|(-iso-b)|(-cn))?:kms:)?([a-z]{2}-[a-z]+(-[a-z]+)?-\\d:)?(\\d{12}:)?(((key/)?[a-zA-Z0-9-_]+)|(alias/[a-zA-Z0-9:/_-]+))",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "CmkType"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "KmsEncryptionConfig"
		//	  ],
		//	  "type": "object"
		//	}
		"sse_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: KmsEncryptionConfig
				"kms_encryption_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CmkType
						"cmk_type": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The type of customer-managed-key (CMK) used for encryption. The two types of supported CMKs are customer owned CMKs and AWS owned CMKs.",
							Required:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"CUSTOMER_MANAGED_KMS_KEY",
									"AWS_OWNED_KMS_KEY",
								),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: KmsKeyId
						"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The KMS encryption key id/alias used to encrypt the Data Store contents at rest.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 400),
								stringvalidator.RegexMatches(regexp.MustCompile("(arn:aws((-us-gov)|(-iso)|(-iso-b)|(-cn))?:kms:)?([a-z]{2}-[a-z]+(-[a-z]+)?-\\d:)?(\\d{12}:)?(((key/)?[a-zA-Z0-9-_]+)|(alias/[a-zA-Z0-9:/_-]+))"), ""),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The customer-managed-key (CMK) used when creating a Data Store. If a customer owned key is not specified, an AWS owned key will be used for encryption.",
					Required:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The server-side encryption key configuration for a customer provided encryption key.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair. A tag consists of a tag key and a tag value. Tag keys and tag values are both required, but tag values can be empty (null) strings.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key of the tag.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value of the tag.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key of the tag.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value of the tag.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
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
		Description: "HealthLake FHIR Datastore",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::HealthLake::FHIRDatastore").WithTerraformTypeName("awscc_healthlake_fhir_datastore")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"authorization_strategy":             "AuthorizationStrategy",
		"cmk_type":                           "CmkType",
		"created_at":                         "CreatedAt",
		"datastore_arn":                      "DatastoreArn",
		"datastore_endpoint":                 "DatastoreEndpoint",
		"datastore_id":                       "DatastoreId",
		"datastore_name":                     "DatastoreName",
		"datastore_status":                   "DatastoreStatus",
		"datastore_type_version":             "DatastoreTypeVersion",
		"fine_grained_authorization_enabled": "FineGrainedAuthorizationEnabled",
		"identity_provider_configuration":    "IdentityProviderConfiguration",
		"idp_lambda_arn":                     "IdpLambdaArn",
		"key":                                "Key",
		"kms_encryption_config":              "KmsEncryptionConfig",
		"kms_key_id":                         "KmsKeyId",
		"metadata":                           "Metadata",
		"nanos":                              "Nanos",
		"preload_data_config":                "PreloadDataConfig",
		"preload_data_type":                  "PreloadDataType",
		"seconds":                            "Seconds",
		"sse_configuration":                  "SseConfiguration",
		"tags":                               "Tags",
		"value":                              "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
