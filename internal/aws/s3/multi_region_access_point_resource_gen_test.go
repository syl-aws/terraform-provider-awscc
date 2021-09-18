// Code generated by generators/resource/main.go; DO NOT EDIT.

package s3_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSS3MultiRegionAccessPoint_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::S3::MultiRegionAccessPoint", "awscc_s3_multi_region_access_point", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}