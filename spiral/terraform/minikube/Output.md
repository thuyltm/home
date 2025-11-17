The console printout these below info when running "terraform apply"
```sh
Apply complete! Resources: 5 added, 0 changed, 0 destroyed.

Outputs:

container_image = "nginxinc/nginx-unprivileged:1.25-alpine"
deployment_generation = 1
deployment_labels = tomap({
  "app" = "MyExampleApp"
  "environment" = "dev"
})
deployment_name = "terraform-example"
deployment_replicas = "1"
kubernetes_connection_info = {
  "config_context" = "minikube"
  "config_path" = "~/.kube/config"
}
namespace_name = "k8s-ns-by-tf"
namespace_uid = "a3e3cf76-4ca0-4776-9ac3-3e599c75b29c"
pod_security_settings = {
  "read_only_root_filesystem" = true
  "run_as_non_root" = true
}
resource_quota_status = tomap({
  "limits.cpu" = "50m"
  "limits.memory" = "10Mi"
  "pods" = "1"
  "requests.cpu" = "25m"
  "requests.memory" = "5Mi"
})
service_cluster_ip = "10.96.58.95"
service_endpoint = "To access the service within the cluster; use: terraform-example-svc.k8s-ns-by-tf.svc.cluster.local"
service_name = "terraform-example-svc"
service_ports = tolist([
  80,
])
```
The later query is
```sh
terraform output -raw service_endpoint
```