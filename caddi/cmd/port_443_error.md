An Istio Gateway port 443 error usually points to a TLS mismatch, a Kubernetes Service port mismatch or missing certificates

1. A Kubernetes Service port mismatch

- Ensure your Ingress Service has port 443/TCP mapped to the correct target port
- Ensure your Gateway resource has protocol HTTPS set and is correctly referencing the Kubernetes Secret for TLS termination

2. Missing or Misconfigured TLS Secret
- Ensure your certificate and key are base64-encoded and placed as Secret in the istio-system namespace
- Configure the ReferenceGrant to allow many namespace to share a single TLS certificate/key secret

3. Port Conflict between multiple gateways
- Ensure different Ingress Gateway controllers are deployed if handling distict domains