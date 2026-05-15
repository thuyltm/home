To resolve the error 
```text
" no matches for kind "HTTPRoute" in version "gateway.networking.k8s.io/v1"
ensure CRDs are installed first"
```
you must manually install the Kubernetes Gateway API Custom Resource Defitnitions (CRDs), as they are not included by default in standard Kubernetes clusters
```sh
kubectl get crd gateways.gateway.networking.k8s.io &> /dev/null || \
  { kubectl kustomize "github.com/kubernetes-sigs/gateway-api/config/crd?ref=v1.4.0" | kubectl apply -f -; }
#customresourcedefinition.apiextensions.k8s.io/backendtlspolicies.gateway.networking.k8s.io created
#customresourcedefinition.apiextensions.k8s.io/gatewayclasses.gateway.networking.k8s.io created
#customresourcedefinition.apiextensions.k8s.io/gateways.gateway.networking.k8s.io created
#customresourcedefinition.apiextensions.k8s.io/grpcroutes.gateway.networking.k8s.io created
#customresourcedefinition.apiextensions.k8s.io/httproutes.gateway.networking.k8s.io created
#customresourcedefinition.apiextensions.k8s.io/referencegrants.gateway.networking.k8s.io created
```