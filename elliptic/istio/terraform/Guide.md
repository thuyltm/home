### Virtual Service Istio vs HttpRoute
Both VirtualService and HTTPRoute handle L7 traffic routing in Istio, but they belong to different API families. 
- VirtualService is Istio's custom, legacy configuration

    Scope: Hight flexible, handles HTTP, TCP, TLS and gRPC all in one object
- HTTPRoute is the standardized, vendor-agnostic Kubernetes Gateway API resource

    Scope: Specifically built for HTTP/HTTPS traffic

### Istio-Ingess vs Istio-System
- Istio-Ingress (or istio-ingressgateway): The actual networking proxy that exposes your applications to external traffic
- Istio-system: The standard Kubernetes namespace where Istio control plane components (like istiod) are installed