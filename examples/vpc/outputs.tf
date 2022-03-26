// VPC Outputs
output "vpc_arn" {
    value= module.terraform_aws_vpc.vpc_arn
}
output "vpc_id" {
    value= module.terraform_aws_vpc.vpc_id
}
output "vpc_cidr_block" {
    value= module.terraform_aws_vpc.vpc_cidr_block
}
// internet gateway outputs
output "igw_arn" {
    value= module.terraform_aws_vpc.igw_arn
}
output "igw_id" {
    value= module.terraform_aws_vpc.igw_id
}
output "igw_vpc_id" {
    value= module.terraform_aws_vpc.igw_vpc_id
}
// public subnet outputs
output "public_subnet_vpc" {
    value= module.terraform_aws_vpc.public_subnet_vpc
}
output "public_subnet_id" {
    value= module.terraform_aws_vpc.public_subnet_id
}
output "public_subnet_cidr" {
    value= module.terraform_aws_vpc.public_subnet_cidr
}
// private subnet outputs
output "private_subnet_vpc" {
    value= module.terraform_aws_vpc.private_subnet_vpc
}
output "private_subnet_cidr" {
    value= module.terraform_aws_vpc.private_subnet_cidr
}
// nat gateway outputs
output "nat_allocation_id" {
    value= module.terraform_aws_vpc.nat_allocation_id
}
output "nat_subnet_id" {
    value= module.terraform_aws_vpc.nat_subnet_id
}
// route table outputs
output "public_route_table_vpc" {
    value= module.terraform_aws_vpc.public_route_table_vpc
}
output "public_route_table_id" {
    value= module.terraform_aws_vpc.public_route_table_id
}
output "private_route_table_vpc" {
    value= module.terraform_aws_vpc.private_route_table_vpc
}
output "private_route_table_id" {
    value= module.terraform_aws_vpc.private_route_table_id
}
// route table association
output "public_route_table_association_tableId" {
    value= module.terraform_aws_vpc.public_route_table_association_tableId
}
output "private_route_table_association_tableId" {
    value= module.terraform_aws_vpc.private_route_table_association_tableId
}