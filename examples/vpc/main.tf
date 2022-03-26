terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

# Configure the AWS Provider
provider "aws" {
  region     = "us-east-1"
  access_key = "AKIA2EFS3QBOE7DFB5MG"
  secret_key = "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"
}

module "terraform_aws_vpc" {
  source="../../"
  public_subnets= var.public_subnets
  private_subnets= var.private_subnets
}