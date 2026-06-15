resource "kubernetes_deployment_v1" "caddi-cmd" {
  metadata {
    name      = "caddi-cmd-deployment"
    namespace = kubernetes_namespace_v1.caddi.metadata[0].name
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = "caddi-cmd"
      }
    }

    template {
      metadata {
        labels = {
          app = "caddi-cmd"
        }
      }

      spec {
        # Bind the pull secret to the deployment pod template
        image_pull_secrets {
          name = kubernetes_secret_v1.docker_pull_secret.metadata[0].name
        }
        container {
          image = "thuyltm2201/caddi_cmd:notls_latest"
          name  = "caddi-cmd-container"

          port {
            container_port = 8080
          }

        }
      }
    }
  }
}