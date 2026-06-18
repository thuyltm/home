### Implement traffic routing using the Gateway API method utilizing gateway.gateway.networking.k8s.io and HttpRoute
[Guide](https://istio.io/latest/docs/tasks/traffic-management/ingress/ingress-control/)

Every Gateway is backed by a service of type LoadBalancer. Kubernetes services of type LoadBalancer are supported in most of cloud platform but in some local environment, you need to do the following
- When working with minikube cluster, running `minikube tunnel` in a different terminal to start an external load balancer
- For Kind cluster, install Cloud Provider Kind which connects to your KIND cluster and provisions new Load Balancer container for your Services [Guide](https://kind.sigs.k8s.io/docs/user/loadbalancer/)
Access the caddi-cmd service using curl:
```sh
 curl -v  -H "Host: caddi.cmd.com" http://10.111.116.203:80/caddi-cmd/ping
```
Note that the -H flag set the Host HTTP header to "caddi.cmd.com". This is needed because your ingress Gateway is configured to handle "caddi.cmd.com", you send your request to the ingress IP
```sh
export INGRESS_HOST=$(kubectl get gtw caddi-cmd-gateway -o jsonpath='{.status.addresses[0].value}')
export INGRESS_PORT=$(kubectl get gtw caddi-cmd-gateway -o jsonpath='{.spec.listeners[?(@.name=="http")].port}')
```

### Configure traffic routing using Istio API method utilizing gateway.networking.istio.io and VirtualService
#### If use gateway service istio-ingress/istio-ingress
```sh
kubectl get services -n istio-ingress
#NAME            TYPE           CLUSTER-IP      EXTERNAL-IP     PORT(S)                                      AGE
#istio-ingress   LoadBalancer   10.100.30.107   10.100.30.107   15021:30918/TCP,80:31685/TCP,443:31190/TCP   4d
curl -v  -H "Host: book.bookinfo.com" http://10.100.30.107:80/productpage | grep -o "<title>.*</title>"
```
#### If use gateway service istio-system/istio-ingressgateway
```sh
kubectl get services -n istio-system
#NAME                          TYPE           CLUSTER-IP       EXTERNAL-IP      PORT(S)                                          AGE
#istio-ingressgateway          LoadBalancer   10.109.253.242   10.109.253.242   15021:32000/TCP,80:31305/TCP,443:31726/TCP       19h
curl -v  -H "Host: book.bookinfo.com" http://10.109.253.242:80/productpage | grep -o "<title>.*</title>"
```