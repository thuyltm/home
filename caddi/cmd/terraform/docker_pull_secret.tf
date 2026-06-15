resource "kubernetes_secret_v1" "docker_pull_secret" {
  metadata {
    name      = "docker-pull-secret"
    namespace = kubernetes_namespace_v1.caddi.metadata[0].name
  }

  # Terraform automatically base64 encodes values in the data block
  data = {
    ".dockerconfigjson" = jsonencode({
      auths = {
        "https://index.docker.io/v1/" = { # Replace with your private registry URL if not Docker Hub
          username = "thuyltm2201"
          #password = "your-docker-password-or-token"
          email = "thuyltm2201@gmail.com"
          auth  = base64encode("thuyltm2201:dckr_pat_Kj84tl0AIdar6v2RHAduGKR40dU")
        }
      }
    })
  }

  type = "kubernetes.io/dockerconfigjson"
}