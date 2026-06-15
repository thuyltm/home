resource "kubernetes_namespace_v1" "caddi" {
  metadata {
    name = "caddi"
  }
}