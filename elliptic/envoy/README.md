# Install
```sh
helm install eg oci://docker.io/envoyproxy/gateway-helm --version v1.6.0 -n envoy-gateway-system --create-namespace
```

**Check**
```sh
Thank you for installing Envoy Gateway
Your release is named: eg
Your release is in namespace: envoy-gateway-system
To learn more about the release, try:
$ helm status eg -n envoy-gateway-system
$ helm get all eg -n envoy-gateway-system
$ helm list -n envoy-gateway-system
```

**Or Waiting for Envoy Gateway to become available**

```sh
kubectl wait --timeout=5m -n envoy-gateway-system deployment/envoy-gateway --for=condition=Available
```

# Information
The file _ingress.yaml_, _httproute.yaml_ will create 3 classes: GatewayClass, Gateway, HTTPRoute
- **GatewayClass** is cluster-scoped resource defined by the infrastructure provider. This resource presents a class of Gateways that can be instantiated. The **_GatewayClass.spec.controller_** field determines the controller implementation responsible for managing the _GatewayClass_.
- When a user creates a **Gateway**, some load balancing infrastructure is provisioned or configured by the GatewayClass controller. The Gateway spec defines the following:
   - GatewayClassName: defines the name of a GatewayClass object used by this Gateway
   - Listeners: - Define the hostnames, ports, protocol, termination, TLS settings and which routes can be attached to a listener
   - Addresses: - Define the network addresses requested for this gateway
- **HTTPRoute** is a Gateway API type for specifying routing behavior of HTTP requests from a Gateway listener to an API object, i.e Service
    The specification of an HTTPRoute consists of:
    - ParentRef: Define which gateways this route wants to be attached to
    - Rules: Define a list of rules to perform actions against matching HTTP request. Each rule consists of matches, filters (optional), backendRefs (optional), timeouts (optional), and name (optional) fields.



```sh
$ k get gatewayclass
NAME   CONTROLLER                                      ACCEPTED   AGE
eg     gateway.envoyproxy.io/gatewayclass-controller   True       3m3s
$ k get gateway
NAME   CLASS   ADDRESS          PROGRAMMED   AGE
eg     eg      10.106.165.113   True         3m21s
$ k get httproute
NAME      HOSTNAMES   AGE
backend               6m34s
```

# Testing
Open a new termination to support External IP in minikube
```sh
minikube tunnel
```
To get the external IP of the Envoy service, run
```sh
export GATEWAY_HOST=$(kubectl get gateway/eg -o jsonpath='{.status.addresses[0].value}')
```
Test
```sh
curl --verbose http://$GATEWAY_HOST
```