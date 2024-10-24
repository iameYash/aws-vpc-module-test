package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestVPCModule(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic_vpc",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	vpcID := terraform.Output(t, terraformOptions, "vpc_id")
	vpcCIDR := terraform.Output(t, terraformOptions, "vpc_cidr_block")

	assert.NotEmpty(t, vpcID)
	assert.Equal(t, "10.0.0.0/16", vpcCIDR)
}
