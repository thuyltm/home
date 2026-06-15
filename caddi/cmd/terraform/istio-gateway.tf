resource "kubernetes_secret_v1" "caddi_tls_secret" {
  metadata {
    name      = "caddi-tls-credential"
    namespace = "istio-ingress"
  }

  type = "kubernetes.io/tls"

  data = {
    "tls.crt" = file("../server.pem")
    "tls.key" = file("../server-key.pem")
  }
}

resource "kubernetes_manifest" "caddi_cmd_gateway" {
    manifest = {
        apiVersion = "networking.istio.io/v1"
        kind = "Gateway"

        metadata = {
            name = "caddi-cmd-gateway"
            namespace = "istio-ingress"
        }

        spec = {
            selector = {
                istio = "ingress"
            }

            servers = [{
                port = {
                    number = 80
                    name = "http"
                    protocol = "HTTP"
                }
                hosts = ["caddi.cmd.com"]
                tls = {
                    httpsRedirect = true
                }
            },
            {
                port = {
                    number = 443
                    name = "https-443"
                    protocol = "HTTPS"
                }
                hosts = ["caddi.cmd.com"]
                tls = {
                    mode = "SIMPLE"
                    credentialName = "caddi-tls-credential"
                }
            }]
        }
    }
}

resource "kubernetes_manifest" "caddi_cmd_virtual_service" {
    depends_on = [kubernetes_manifest.caddi_cmd_gateway]

    manifest = {
        apiVersion = "networking.istio.io/v1beta1"
        kind = "VirtualService"
        metadata = {
            name = "caddi-cmd-router"
            namespace = "caddi" # namespace where your actual workload runs
        }
        spec = {
            hosts = ["caddi.cmd.com"]
            gateways = ["istio-ingress/caddi-cmd-gateway"]
            http = [
                {
                    match = [
                        {
                            uri = {
                                prefix = "/caddi-cmd"
                            }
                        }
                    ]
                    route = [
                        {
                            destination = {
                                host = "caddi-cmd-service.caddi.svc.cluster.local"
                                port = {
                                    number = 8080
                                }
                            }
                        }
                    ]
                }
            ]
        }
    }
}