// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package frauddetector

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_frauddetector_detector", detectorDataSource)
}

// detectorDataSource returns the Terraform awscc_frauddetector_detector data source.
// This Terraform data source corresponds to the CloudFormation AWS::FraudDetector::Detector resource.
func detectorDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the detector.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the detector.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AssociatedModels
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The models to associate with this detector.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A model to associate with a detector.",
		//	    "properties": {
		//	      "Arn": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "maxItems": 10,
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"associated_models": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Arn
					"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The models to associate with this detector.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreatedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time when the detector was created.",
		//	  "type": "string"
		//	}
		"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time when the detector was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the detector.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the detector.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DetectorId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the detector",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[0-9a-z_-]+$",
		//	  "type": "string"
		//	}
		"detector_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the detector",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DetectorVersionId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The active version ID of the detector",
		//	  "type": "string"
		//	}
		"detector_version_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The active version ID of the detector",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DetectorVersionStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The desired detector version status for the detector",
		//	  "enum": [
		//	    "DRAFT",
		//	    "ACTIVE"
		//	  ],
		//	  "type": "string"
		//	}
		"detector_version_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The desired detector version status for the detector",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EventType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The event type to associate this detector with.",
		//	  "properties": {
		//	    "Arn": {
		//	      "description": "The ARN of the event type.",
		//	      "type": "string"
		//	    },
		//	    "CreatedTime": {
		//	      "description": "The time when the event type was created.",
		//	      "type": "string"
		//	    },
		//	    "Description": {
		//	      "description": "The description of the event type.",
		//	      "maxLength": 128,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "EntityTypes": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Arn": {
		//	            "type": "string"
		//	          },
		//	          "CreatedTime": {
		//	            "description": "The time when the entity type was created.",
		//	            "type": "string"
		//	          },
		//	          "Description": {
		//	            "description": "The description.",
		//	            "maxLength": 256,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Inline": {
		//	            "type": "boolean"
		//	          },
		//	          "LastUpdatedTime": {
		//	            "description": "The time when the entity type was last updated.",
		//	            "type": "string"
		//	          },
		//	          "Name": {
		//	            "type": "string"
		//	          },
		//	          "Tags": {
		//	            "description": "Tags associated with this entity type.",
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "Key": {
		//	                  "maxLength": 128,
		//	                  "minLength": 1,
		//	                  "type": "string"
		//	                },
		//	                "Value": {
		//	                  "maxLength": 256,
		//	                  "minLength": 0,
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "required": [
		//	                "Key",
		//	                "Value"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "maxItems": 200,
		//	            "type": "array",
		//	            "uniqueItems": false
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "minItems": 1,
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "EventVariables": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Arn": {
		//	            "type": "string"
		//	          },
		//	          "CreatedTime": {
		//	            "description": "The time when the event variable was created.",
		//	            "type": "string"
		//	          },
		//	          "DataSource": {
		//	            "enum": [
		//	              "EVENT"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "DataType": {
		//	            "enum": [
		//	              "STRING",
		//	              "INTEGER",
		//	              "FLOAT",
		//	              "BOOLEAN"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "DefaultValue": {
		//	            "type": "string"
		//	          },
		//	          "Description": {
		//	            "description": "The description.",
		//	            "maxLength": 256,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Inline": {
		//	            "type": "boolean"
		//	          },
		//	          "LastUpdatedTime": {
		//	            "description": "The time when the event variable was last updated.",
		//	            "type": "string"
		//	          },
		//	          "Name": {
		//	            "type": "string"
		//	          },
		//	          "Tags": {
		//	            "description": "Tags associated with this event variable.",
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "Key": {
		//	                  "maxLength": 128,
		//	                  "minLength": 1,
		//	                  "type": "string"
		//	                },
		//	                "Value": {
		//	                  "maxLength": 256,
		//	                  "minLength": 0,
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "required": [
		//	                "Key",
		//	                "Value"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "maxItems": 200,
		//	            "type": "array",
		//	            "uniqueItems": false
		//	          },
		//	          "VariableType": {
		//	            "enum": [
		//	              "AUTH_CODE",
		//	              "AVS",
		//	              "BILLING_ADDRESS_L1",
		//	              "BILLING_ADDRESS_L2",
		//	              "BILLING_CITY",
		//	              "BILLING_COUNTRY",
		//	              "BILLING_NAME",
		//	              "BILLING_PHONE",
		//	              "BILLING_STATE",
		//	              "BILLING_ZIP",
		//	              "CARD_BIN",
		//	              "CATEGORICAL",
		//	              "CURRENCY_CODE",
		//	              "EMAIL_ADDRESS",
		//	              "FINGERPRINT",
		//	              "FRAUD_LABEL",
		//	              "FREE_FORM_TEXT",
		//	              "IP_ADDRESS",
		//	              "NUMERIC",
		//	              "ORDER_ID",
		//	              "PAYMENT_TYPE",
		//	              "PHONE_NUMBER",
		//	              "PRICE",
		//	              "PRODUCT_CATEGORY",
		//	              "SHIPPING_ADDRESS_L1",
		//	              "SHIPPING_ADDRESS_L2",
		//	              "SHIPPING_CITY",
		//	              "SHIPPING_COUNTRY",
		//	              "SHIPPING_NAME",
		//	              "SHIPPING_PHONE",
		//	              "SHIPPING_STATE",
		//	              "SHIPPING_ZIP",
		//	              "USERAGENT"
		//	            ],
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "minItems": 1,
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "Inline": {
		//	      "type": "boolean"
		//	    },
		//	    "Labels": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Arn": {
		//	            "type": "string"
		//	          },
		//	          "CreatedTime": {
		//	            "description": "The time when the label was created.",
		//	            "type": "string"
		//	          },
		//	          "Description": {
		//	            "description": "The description.",
		//	            "maxLength": 256,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Inline": {
		//	            "type": "boolean"
		//	          },
		//	          "LastUpdatedTime": {
		//	            "description": "The time when the label was last updated.",
		//	            "type": "string"
		//	          },
		//	          "Name": {
		//	            "type": "string"
		//	          },
		//	          "Tags": {
		//	            "description": "Tags associated with this label.",
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "Key": {
		//	                  "maxLength": 128,
		//	                  "minLength": 1,
		//	                  "type": "string"
		//	                },
		//	                "Value": {
		//	                  "maxLength": 256,
		//	                  "minLength": 0,
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "required": [
		//	                "Key",
		//	                "Value"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "maxItems": 200,
		//	            "type": "array",
		//	            "uniqueItems": false
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "minItems": 2,
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "LastUpdatedTime": {
		//	      "description": "The time when the event type was last updated.",
		//	      "type": "string"
		//	    },
		//	    "Name": {
		//	      "description": "The name for the event type",
		//	      "maxLength": 64,
		//	      "minLength": 1,
		//	      "pattern": "^[0-9a-z_-]+$",
		//	      "type": "string"
		//	    },
		//	    "Tags": {
		//	      "description": "Tags associated with this event type.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Key": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Value": {
		//	            "maxLength": 256,
		//	            "minLength": 0,
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Key",
		//	          "Value"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 200,
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"event_type": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Arn
				"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the event type.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: CreatedTime
				"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The time when the event type was created.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Description
				"description": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The description of the event type.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: EntityTypes
				"entity_types": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Arn
							"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: CreatedTime
							"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The time when the entity type was created.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Description
							"description": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The description.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Inline
							"inline": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: LastUpdatedTime
							"last_updated_time": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The time when the entity type was last updated.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Tags
							"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
								NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Key
										"key": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: Value
										"value": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
								}, /*END NESTED OBJECT*/
								Description: "Tags associated with this entity type.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: EventVariables
				"event_variables": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Arn
							"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: CreatedTime
							"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The time when the event variable was created.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: DataSource
							"data_source": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: DataType
							"data_type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: DefaultValue
							"default_value": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Description
							"description": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The description.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Inline
							"inline": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: LastUpdatedTime
							"last_updated_time": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The time when the event variable was last updated.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Tags
							"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
								NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Key
										"key": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: Value
										"value": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
								}, /*END NESTED OBJECT*/
								Description: "Tags associated with this event variable.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: VariableType
							"variable_type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Inline
				"inline": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Labels
				"labels": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Arn
							"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: CreatedTime
							"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The time when the label was created.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Description
							"description": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The description.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Inline
							"inline": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: LastUpdatedTime
							"last_updated_time": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The time when the label was last updated.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Tags
							"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
								NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Key
										"key": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: Value
										"value": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
								}, /*END NESTED OBJECT*/
								Description: "Tags associated with this label.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: LastUpdatedTime
				"last_updated_time": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The time when the event type was last updated.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Name
				"name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name for the event type",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Tags
				"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Key
							"key": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Value
							"value": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "Tags associated with this event type.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The event type to associate this detector with.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LastUpdatedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time when the detector was last updated.",
		//	  "type": "string"
		//	}
		"last_updated_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time when the detector was last updated.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RuleExecutionMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "FIRST_MATCHED",
		//	    "ALL_MATCHED"
		//	  ],
		//	  "type": "string"
		//	}
		"rule_execution_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Rules
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Arn": {
		//	        "type": "string"
		//	      },
		//	      "CreatedTime": {
		//	        "description": "The time when the event type was created.",
		//	        "type": "string"
		//	      },
		//	      "Description": {
		//	        "description": "The description.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "DetectorId": {
		//	        "type": "string"
		//	      },
		//	      "Expression": {
		//	        "type": "string"
		//	      },
		//	      "Language": {
		//	        "enum": [
		//	          "DETECTORPL"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "LastUpdatedTime": {
		//	        "description": "The time when the event type was last updated.",
		//	        "type": "string"
		//	      },
		//	      "Outcomes": {
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Arn": {
		//	              "type": "string"
		//	            },
		//	            "CreatedTime": {
		//	              "description": "The time when the outcome was created.",
		//	              "type": "string"
		//	            },
		//	            "Description": {
		//	              "description": "The description.",
		//	              "maxLength": 256,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            },
		//	            "Inline": {
		//	              "type": "boolean"
		//	            },
		//	            "LastUpdatedTime": {
		//	              "description": "The time when the outcome was last updated.",
		//	              "type": "string"
		//	            },
		//	            "Name": {
		//	              "type": "string"
		//	            },
		//	            "Tags": {
		//	              "description": "Tags associated with this outcome.",
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "additionalProperties": false,
		//	                "properties": {
		//	                  "Key": {
		//	                    "maxLength": 128,
		//	                    "minLength": 1,
		//	                    "type": "string"
		//	                  },
		//	                  "Value": {
		//	                    "maxLength": 256,
		//	                    "minLength": 0,
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "required": [
		//	                  "Key",
		//	                  "Value"
		//	                ],
		//	                "type": "object"
		//	              },
		//	              "maxItems": 200,
		//	              "type": "array",
		//	              "uniqueItems": false
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "minItems": 1,
		//	        "type": "array",
		//	        "uniqueItems": false
		//	      },
		//	      "RuleId": {
		//	        "type": "string"
		//	      },
		//	      "RuleVersion": {
		//	        "type": "string"
		//	      },
		//	      "Tags": {
		//	        "description": "Tags associated with this event type.",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Key": {
		//	              "maxLength": 128,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            },
		//	            "Value": {
		//	              "maxLength": 256,
		//	              "minLength": 0,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "Key",
		//	            "Value"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "maxItems": 200,
		//	        "type": "array",
		//	        "uniqueItems": false
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"rules": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Arn
					"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: CreatedTime
					"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The time when the event type was created.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Description
					"description": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The description.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: DetectorId
					"detector_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Expression
					"expression": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Language
					"language": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: LastUpdatedTime
					"last_updated_time": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The time when the event type was last updated.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Outcomes
					"outcomes": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Arn
								"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: CreatedTime
								"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The time when the outcome was created.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: Description
								"description": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The description.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: Inline
								"inline": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: LastUpdatedTime
								"last_updated_time": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The time when the outcome was last updated.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: Name
								"name": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: Tags
								"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
									NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: Key
											"key": schema.StringAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
											// Property: Value
											"value": schema.StringAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
									}, /*END NESTED OBJECT*/
									Description: "Tags associated with this outcome.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: RuleId
					"rule_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: RuleVersion
					"rule_version": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Tags
					"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Key
								"key": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: Value
								"value": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Description: "Tags associated with this event type.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Tags associated with this detector.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
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
		//	  "maxItems": 200,
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Tags associated with this detector.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::FraudDetector::Detector",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::FraudDetector::Detector").WithTerraformTypeName("awscc_frauddetector_detector")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                     "Arn",
		"associated_models":       "AssociatedModels",
		"created_time":            "CreatedTime",
		"data_source":             "DataSource",
		"data_type":               "DataType",
		"default_value":           "DefaultValue",
		"description":             "Description",
		"detector_id":             "DetectorId",
		"detector_version_id":     "DetectorVersionId",
		"detector_version_status": "DetectorVersionStatus",
		"entity_types":            "EntityTypes",
		"event_type":              "EventType",
		"event_variables":         "EventVariables",
		"expression":              "Expression",
		"inline":                  "Inline",
		"key":                     "Key",
		"labels":                  "Labels",
		"language":                "Language",
		"last_updated_time":       "LastUpdatedTime",
		"name":                    "Name",
		"outcomes":                "Outcomes",
		"rule_execution_mode":     "RuleExecutionMode",
		"rule_id":                 "RuleId",
		"rule_version":            "RuleVersion",
		"rules":                   "Rules",
		"tags":                    "Tags",
		"value":                   "Value",
		"variable_type":           "VariableType",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
