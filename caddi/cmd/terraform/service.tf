resource "kubernetes_service_v1" "caddi_cmd_service" {
  metadata {
    name      = "caddi-cmd-service"
    namespace = kubernetes_namespace_v1.caddi.metadata[0].name
  }

  spec {
    selector = {
      app = kubernetes_deployment_v1.caddi-cmd.spec[0].template[0].metadata[0].labels.app
    }

    port {
      port        = 8080
      target_port = 8080
    }

    type = "NodePort"
  }
}