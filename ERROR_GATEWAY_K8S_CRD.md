1.

```sh
Error: API did not recognize GroupVersionKind from manifest (CRD may not be installed)
│ 
│   with kubernetes_manifest.caddi-cmd-gateway,
│   on istio-gateway.tf line 1, in resource "kubernetes_manifest" "caddi-cmd-gateway":
│    1: resource "kubernetes_manifest" "caddi-cmd-gateway" {
│ 
│ no matches for kind "Gateway" in group "gateway.networking.k8s.io"
```
The aboved error is caused by the Gateway API resources (like GatewayClass, Gateway, and HTTPRoute) do not come built-in with standard Kubernetes clusters. 

Install K8s Gateway API CRUD
```sh
kubectl apply -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.5.1/standard-install.yaml
```
Check
```sh
kubectl api-resources | grep -i gateway
```
Next you install Istio or Envoy Gateway

2.
```sh
failed to get cpu utilization: unable to get metrics for resource cpu: unable to fetch metrics from resource metrics API: the server could not find the requested resource (get pods.metrics.k8s.io)
```
This error indicates that your Kubernetes Horizontal Pod Autoscaler (HPA) cannot find the Metrics API (metrics.k8s.io) because the metrics-server component is either missing, disabled, or misconfigured

Verify the Metrics API Status
```sh
kubectl get apiservice v1beta1.metrics.k8s.io
```
Enable metric server in minikube
```sh
minikube addons enable metrics-server
```
