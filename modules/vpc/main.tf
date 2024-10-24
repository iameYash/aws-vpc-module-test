resource "aws_vpc" "this" {
  cidr_block           = var.cidr
  enable_dns_hostnames = var.enable_dns_hostnames
  enable_dns_support   = var.enable_dns_support
  
  tags = merge(
    { "Name" = var.name },
    var.tags,
  )
}

# Add other resources like subnets, route tables, etc.
