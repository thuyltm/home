When you create a standard Kubernetes Ingress resource while using Istio as your controller, Istio does not physically create a visible Gate in your cluster. Instead, the Istio control plan converts the Ingress object into internal Envoy routing configurations dynamically inside memory

Because Istio translates standard Ingress objects silently behind the scenes, you cannot see them using _kubectl get gateway_. You have to check Istio's internal debug runtime to see if it processed your Ingress properly
```sh
istioctl analyze
istioctl x internal-debug configz
```

#### Install Istioctl
https://istio.io/latest/docs/setup/getting-started/