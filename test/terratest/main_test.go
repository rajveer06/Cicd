package test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformDeployment(t *testing.T) {
	options := &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../",
	}

	// Defer the destroy until after the test has run.
	defer terraform.Destroy(t, options)

	// Run terraform init and apply. Fail the test if there are any errors.
	initAndApply := terraform.InitAndApply(t, options)

	// Use Terraform output to get the value of an output variable.
	vpcID := terraform.Output(t, options, "vpc_id")

	// Check if the VPC ID is not empty.
	assert.NotEmpty(t, vpcID)

	// Additional assertions can be added based on your deployment.
	// For example, check if specific resources are created, tags are applied, etc.
	// assert.True(t, strings.HasPrefix(vpcID, "vpc-"))

	// Validate that the VPC ID is present in the Terraform output.
	assert.True(t, strings.Contains(initAndApply, vpcID))
}
