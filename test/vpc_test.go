package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformVpcExample(t *testing.T) {
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/vpc",
		Vars: map[string]interface{}{
			"private_subnets":"10.0.0.192/26",
            "public_subnets":"10.0.0.128/26",
		},
	})

	terraform.InitAndApply(t, terraformOptions)

	defer terraform.Destroy(t, terraformOptions)

	// VPC test
	vpc_cidr_block := terraform.Output(t, terraformOptions, "vpc_cidr_block")
	vpc_id := terraform.Output(t, terraformOptions, "vpc_id")
	assert.Equal(t, "10.0.0.0/16", vpc_cidr_block)

	// Internet gateway test
	igw_vpc_id := terraform.Output(t, terraformOptions, "igw_vpc_id")
	assert.Equal(t, vpc_id, igw_vpc_id)

	// public subnet test
	public_subnet_vpc := terraform.Output(t, terraformOptions, "public_subnet_vpc")
	public_subnet_id := terraform.Output(t, terraformOptions, "public_subnet_id")
	assert.Equal(t, vpc_id, public_subnet_vpc)

	// private subnet test
	private_subnet_vpc := terraform.Output(t, terraformOptions, "private_subnet_vpc")
	assert.Equal(t, vpc_id, private_subnet_vpc)

	// nat gateway test
	nat_subnet_id := terraform.Output(t, terraformOptions, "nat_subnet_id")
	assert.Equal(t, public_subnet_id, nat_subnet_id)

	// public route table test
	public_route_table_vpc := terraform.Output(t, terraformOptions, "public_route_table_vpc")
	public_route_table_id := terraform.Output(t, terraformOptions, "public_route_table_id")
	assert.Equal(t, vpc_id, public_route_table_vpc)

	// public route table test
	private_route_table_vpc := terraform.Output(t, terraformOptions, "private_route_table_vpc")
	private_route_table_id := terraform.Output(t, terraformOptions, "private_route_table_id")
	assert.Equal(t, vpc_id, private_route_table_vpc)

	// public route table association test
	public_route_table_association_tableId := terraform.Output(t, terraformOptions, "public_route_table_association_tableId")
	assert.Equal(t, public_route_table_id, public_route_table_association_tableId)

	// private route table association test
	private_route_table_association_tableId := terraform.Output(t, terraformOptions, "private_route_table_association_tableId")
	assert.Equal(t, private_route_table_id, private_route_table_association_tableId)
}