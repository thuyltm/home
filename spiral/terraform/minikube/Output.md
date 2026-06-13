Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # kubernetes_deployment.nginx will be created
  + resource "kubernetes_deployment" "nginx" {
      + id               = (known after apply)
      + wait_for_rollout = true

      + metadata {
          + generation       = (known after apply)
          **+ name             = "nginx-server"**
          **+ namespace        = "web-server-namespace"**
          + resource_version = (known after apply)
          + uid              = (known after apply)
        }

      + spec {
          + min_ready_seconds         = 0
          + paused                    = false
          + progress_deadline_seconds = 600
          + replicas                  = "1"
          + revision_history_limit    = 10

          + selector {
              + match_labels = {
                  **+ "app" = "nginx-web"**
                }
            }

          + strategy (known after apply)

          + template {
              + metadata {
                  + generation       = (known after apply)
                  + labels           = {
                      + "app" = "nginx-web"
                    }
                  + name             = (known after apply)
                  + resource_version = (known after apply)
                  + uid              = (known after apply)
                }
              + spec {
                  + automount_service_account_token  = true
                  + dns_policy                       = "ClusterFirst"
                  + enable_service_links             = true
                  + host_ipc                         = false
                  + host_network                     = false
                  + host_pid                         = false
                  + hostname                         = (known after apply)
                  + node_name                        = (known after apply)
                  + restart_policy                   = "Always"
                  + scheduler_name                   = (known after apply)
                  + service_account_name             = (known after apply)
                  + share_process_namespace          = false
                  + termination_grace_period_seconds = 30

                  + container {
                      + image                      = "nginx:latest"
                      + image_pull_policy          = (known after apply)
                      + name                       = "nginx-container"
                      + restart_policy             = (known after apply)
                      + stdin                      = false
                      + stdin_once                 = false
                      + termination_message_path   = "/dev/termination-log"
                      + termination_message_policy = (known after apply)
                      + tty                        = false

                      + port {
                          + container_port = 80
                          + protocol       = "TCP"
                        }

                      + resources (known after apply)
                    }

                  + image_pull_secrets (known after apply)

                  + readiness_gate (known after apply)
                }
            }
        }
    }

  # kubernetes_namespace.web_server will be created
  + resource "kubernetes_namespace" "web_server" {
      + id                               = (known after apply)
      + wait_for_default_service_account = false

      + metadata {
          + generation       = (known after apply)
          + name             = "web-server-namespace"
          + resource_version = (known after apply)
          + uid              = (known after apply)
        }
    }

  # kubernetes_service.nginx_service will be created
  + resource "kubernetes_service" "nginx_service" {
      + id                     = (known after apply)
      + status                 = (known after apply)
      + wait_for_load_balancer = true

      + metadata {
          + generation       = (known after apply)
          + name             = "nginx-service"
          + namespace        = "web-server-namespace"
          + resource_version = (known after apply)
          + uid              = (known after apply)
        }

      + spec {
          + allocate_load_balancer_node_ports = true
          + cluster_ip                        = (known after apply)
          + cluster_ips                       = (known after apply)
          + external_traffic_policy           = (known after apply)
          + health_check_node_port            = (known after apply)
          + internal_traffic_policy           = (known after apply)
          + ip_families                       = (known after apply)
          + ip_family_policy                  = (known after apply)
          + publish_not_ready_addresses       = false
          + selector                          = {
              + "app" = "nginx-web"
            }
          + session_affinity                  = "None"
          + type                              = "NodePort"

          + port {
              + node_port   = 30080
              + port        = 80
              + protocol    = "TCP"
              + target_port = "80"
            }

          + session_affinity_config (known after apply)
        }
    }

Plan: 3 to add, 0 to change, 0 to destroy.