provider "aws" {
  region = "ap-northeast-1"
  default_tags {
    tags = {
      env = "prod"
    }
  }
}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.12.0"
    }
  }

  required_version = ">= 1.5"

  backend "s3" {
    bucket         = "learning-terraform-prod" # create terraform s3 bucket
    region         = "ap-northeast-1"
    key            = "terraform.tfstate"
    encrypt        = true
    dynamodb_table = "learning_terraform_state_lock_prod" # dynamodb lock
  }
}
