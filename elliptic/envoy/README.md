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
**Note**: 

Namespace: envoy-gateway-system

ConfigMap/envoy-gateway-config, Deployment/envoy-gateway, Service/envoy-gateway

# Create 3 classes: GatewayClass, Gateway, HTTPRoute
```sh
$ k get gatewayclass
NAME   CONTROLLER                                      ACCEPTED   AGE
eg     gateway.envoyproxy.io/gatewayclass-controller   True       3m3s
$ k get gateway
NAME   CLASS   ADDRESS          PROGRAMMED   AGE
eg     eg      10.106.165.113   True         3m21s
$ k get httproute
NAME              HOSTNAMES                AGE
route-servicea    www.example.com          6m34s
```

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

# Testing Local RateLimit
```sh
for i in {1..4}; do curl --resolve "www.example.com:80:$( echo $GATEWAY_HOST )" -i www.example.com/servicea/hi; sleep 1; done
#Output
HTTP/1.1 200 OK
date: Fri, 28 Nov 2025 04:55:57 GMT
content-length: 9
content-type: text/plain; charset=utf-8
x-ratelimit-limit: 1
x-ratelimit-remaining: 0
x-ratelimit-reset: 3599

hi Go ChiHTTP/1.1 429 Too Many Requests
content-length: 18
content-type: text/plain
x-ratelimit-limit: 1
x-ratelimit-remaining: 0
x-ratelimit-reset: 3598
date: Fri, 28 Nov 2025 04:55:58 GMT

local_rate_limitedHTTP/1.1 429 Too Many Requests
```