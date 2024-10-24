# AWS VPC Module Test

This repository contains a Terraform module for creating an AWS VPC and tests for that module.

## Structure

- `modules/vpc/`: Contains the main VPC module
- `examples/basic_vpc/`: Contains an example usage of the VPC module
- `tests/`: Contains Go tests for the VPC module

## Running Tests

1. Ensure you have Go and Terraform installed.
2. Set up your AWS credentials.
3. Run `go test -v ./tests`

