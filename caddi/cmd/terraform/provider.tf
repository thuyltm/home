terraform {
  required_providers {
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = ">= 2.0.0"
    }
  }
}

# Configure the provider to point to your local Minikube context
provider "kubernetes" {
  config_path = "~/.kube/config"
}