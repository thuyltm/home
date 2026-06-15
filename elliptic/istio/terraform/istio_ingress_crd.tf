# create a namespace for your Ingress Gateway
resource "kubernetes_namespace_v1" "istio_ingress" {
  metadata {
    name = "istio-ingress"
  }
}

#  Istio Ingress Gateway Deployment
resource "helm_release" "istio_ingress" {
  name       = "istio-ingress"
  repository = "https://istio-release.storage.googleapis.com/charts"
  chart      = "gateway"
  namespace  = kubernetes_namespace_v1.istio_ingress.id
  version    = "1.30.1"
  # Configuration overrides (adjust according to your cloud provider)
  values = [
    yamlencode({
      service = {
        type = "LoadBalancer"
      }
    })
  ]
  # Ensure the Istio Base chart and control plane are deployed first
  depends_on = [
    helm_release.istio_base,
    helm_release.istiod
  ]
}

# 1. Install Istio Base (CRDs)
resource "helm_release" "istio_base" {
  name             = "istio-base"
  repository       = "https://istio-release.storage.googleapis.com/charts"
  chart            = "base"
  namespace        = "istio-system"
  create_namespace = true
  version          = "1.30.1" # match this with your desired version
}

# 2. Install Istio Discovery / Control Plane (istiod)
resource "helm_release" "istiod" {
  name       = "istiod"
  repository = "https://istio-release.storage.googleapis.com/charts"
  chart      = "istiod"
  namespace  = "istio-system"
  version    = "1.30.1"

  # Explicitly wait for CRDs to establish before provisioning istiod
  depends_on = [helm_release.istio_base]
}