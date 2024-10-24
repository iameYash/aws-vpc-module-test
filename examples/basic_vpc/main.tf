provider "aws" {
  region = "us-west-2"
}

module "vpc" {
  source = "../../modules/vpc"

  name = "test-vpc"
  cidr = "10.0.0.0/16"

  tags = {
    Environment = "test"
    Project     = "VPC Module Test"
  }
}
