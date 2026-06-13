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

# 1. Create an isolated Namespace
resource "kubernetes_namespace" "web_server" {
  metadata {
    name = "web-server-namespace"
  }
}

# 2. Deploy the NGINX Server
resource "kubernetes_deployment" "nginx" {
  metadata {
    name      = "nginx-server"
    namespace = kubernetes_namespace.web_server.metadata[0].name
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = "nginx-web"
      }
    }

    template {
      metadata {
        labels = {
          app = "nginx-web"
        }
      }

      spec {
        container {
          image = "nginx:latest"
          name  = "nginx-container"

          port {
            container_port = 80
          }
        }
      }
    }
  }
}

# 3. Expose the Server via NodePort Service
resource "kubernetes_service" "nginx_service" {
  metadata {
    name      = "nginx-service"
    namespace = kubernetes_namespace.web_server.metadata[0].name
  }

  spec {
    selector = {
      app = kubernetes_deployment.nginx.spec[0].template[0].metadata[0].labels.app
    }

    port {
      port        = 80
      target_port = 80
      node_port   = 30080
    }

    type = "NodePort"
  }
}
