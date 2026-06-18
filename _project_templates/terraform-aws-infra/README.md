# terraform-aws-infra

AWS infrastructure starter using reusable Terraform modules for VPC and EKS.

## Prerequisites
- Terraform >= 1.7
- AWS credentials configured

## Usage
```bash
terraform init
terraform plan -var="env=dev" -var="cluster_name=devops-yasan-eks"
terraform apply
```

## Modules
- `modules/vpc`: VPC and public/private subnets
- `terraform-aws-modules/eks/aws`: managed EKS cluster
