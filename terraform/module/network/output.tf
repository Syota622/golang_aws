# Output: VPCs
output "vpc_main_id" {
  description = "The ID of the Dev VPC"
  value       = aws_vpc.main_vpc.id
}

# Output: Dev Subnets
output "public_subnet" {
  description = "The IDs of the Dev Public Subnets"
  value       = [for subnet in aws_subnet.public_subnet : subnet.id]
}

# Output: Prod Subnets
output "private_subnet" {
  description = "The IDs of the Prod Public Subnets"
  value       = [for subnet in aws_subnet.private_subnet : subnet.id]
}
