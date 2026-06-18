variable "region" {
  description = "AWS region"
  type        = string
  default     = "us-east-1"
}

variable "env" {
  description = "Deployment environment"
  type        = string
  default     = "dev"
}

variable "cluster_name" {
  description = "EKS cluster name"
  type        = string
}

variable "node_count" {
  description = "Desired number of EKS worker nodes"
  type        = number
  default     = 2
}
