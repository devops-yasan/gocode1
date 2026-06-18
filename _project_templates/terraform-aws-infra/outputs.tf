output "vpc_id" {
  value       = module.vpc.vpc_id
  description = "VPC ID"
}

output "cluster_endpoint" {
  value       = module.eks.cluster_endpoint
  description = "EKS API endpoint"
}

output "kubeconfig" {
  value       = "aws eks update-kubeconfig --name ${var.cluster_name} --region ${var.region}"
  description = "Command to update local kubeconfig"
}
