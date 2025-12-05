
A GatewayClass indicates which controller should manage Gateway. The _GatewayClass.spec.controller_ field determines the controller implementation responsible for managing the Gateway

A Gateway resource is created that reference to a GatewayClass.

The Gateway controller takes over to provision and manage Proxy deployment

- Gateway:
   - GatewayClassName: defines the name of a GatewayClass object used by this Gateway
   - Listeners: - Define the hostnames, ports, protocol, termination, TLS settings and which routes can be attached to a listener
   - Addresses: - Define the network addresses requested for this gateway
- HTTPRoute is a routing behavior of HTTP requests from a Gateway listener to an API object, i.e Service
    - ParentRef: Define which gateways this route wants to be attached to
    - Rules: Define a list of rules to perform actions against matching HTTP request. Each rule consists of matches, filters (optional), backendRefs (optional), timeouts (optional), and name (optional) fields.

Note: Check envoy-default-{GATEWAYNAME}-.. created automatically by typing `k get services -n envoy-gateway-system`
# Testing
Open a new termination to support External IP in minikube
```sh
minikube tunnel
```
To get the external IP of the Envoy service, run
```sh
export GATEWAY_HOST=$(kubectl get gateway/gateway-servicea -o jsonpath='{.status.addresses[0].value}')
```
Test
```sh
curl --resolve "www.example.com:80:$( echo $GATEWAY_HOST )" -i www.example.com/servicea
curl --resolve "www.example.com:80:$( echo $GATEWAY_HOST )" -i www.example.com/serviceb
```
