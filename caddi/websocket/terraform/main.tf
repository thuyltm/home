terraform {
  required_providers {
    helm = {
      source = "hashicorp/helm"
      version = "4.2.0"
    }
  }
}

provider "helm" {
  kubernetes {
    config_path = "~/.kube/config"  # Path to your Kubernetes config file
    # localhost registry with password protection
    registry {
        url = "oci://localhost:5000"
        username = "username"
        password = "password"
    }

    # private registry
    registry {
        url = "oci://private.registry"
        username = "username"
        password = "password"
    }
  }
}

resource "helm_release" "nginx" {
  name       = "my-nginx"
  repository = "https://charts.bitnami.com/bitnami"
  chart      = "nginx-ingress-controller"
  version    = "8.0.1"

  set {
    name  = "service.type"
    value = "ClusterIP"
  }
}

resource "helm_release" "redis" {
  name       = "my-redis"
  repository = "https://charts.bitnami.com/bitnami"
  chart      = "redis"
  version    = "15.0.10"
}

resource "helm_release" "prometheus" {
  name       = "my-prometheus"
  repository = "https://prometheus-community.github.io/helm-charts"
  chart      = "prometheus"
  version    = "15.0.0"
}