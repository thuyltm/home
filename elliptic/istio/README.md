# Install Istio Ingress
You can install Istio using the minimal profile
```sh
helm repo add istio https://istio-release.storage.googleapis.com/charts
helm repo update
helm install istio-base istio/base -n istio-system --set defaultRevision=default --create-namespace
helm status istio-base -n istio-system
#NAME: istio-base
#LAST DEPLOYED: Fri May 15 15:56:54 2026
#NAMESPACE: istio-system
#STATUS: deployed
#REVISION: 1
#DESCRIPTION: Install complete
#RESOURCES:
#==> v1/ServiceAccount
#NAME                           AGE
#istio-reader-service-account   43s

#==> v1/CustomResourceDefinition
#NAME                              CREATED AT
#wasmplugins.extensions.istio.io   2026-05-15T08:56:55Z
#destinationrules.networking.istio.io   2026-05-15T08:56:56Z
#envoyfilters.networking.istio.io   2026-05-15T08:56:55Z
#gateways.networking.istio.io   2026-05-15T08:56:55Z
#proxyconfigs.networking.istio.io   2026-05-15T08:56:55Z
#serviceentries.networking.istio.io   2026-05-15T08:56:55Z
#sidecars.networking.istio.io   2026-05-15T08:56:55Z
#virtualservices.networking.istio.io   2026-05-15T08:56:55Z
#workloadentries.networking.istio.io   2026-05-15T08:56:55Z
#workloadgroups.networking.istio.io   2026-05-15T08:56:55Z
#authorizationpolicies.security.istio.io   2026-05-15T08:56:55Z
#peerauthentications.security.istio.io   2026-05-15T08:56:55Z
#requestauthentications.security.istio.io   2026-05-15T08:56:55Z
#telemetries.telemetry.istio.io   2026-05-15T08:56:55Z

#==> v1/ValidatingWebhookConfiguration
#NAME                       WEBHOOKS   AGE
#istiod-default-validator   1          43s


#TEST SUITE: None
#NOTES:
#Istio base successfully installed!

#Install the Istio discovery chart which deploys the istiod service
helm install istiod istio/istiod -n istio-system --wait
#NAME: istiod
#LAST DEPLOYED: Sat May 16 15:14:13 2026
#NAMESPACE: istio-system
#STATUS: deployed
#REVISION: 1
#DESCRIPTION: Install complete
#RESOURCES:
#==> v1/Service
#NAME     TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                                 AGE
#istiod   ClusterIP   10.110.133.115   <none>        15010/TCP,15012/TCP,443/TCP,15014/TCP   45s

#==> v1/MutatingWebhookConfiguration
#NAME                     WEBHOOKS   AGE
#istio-sidecar-injector   4          45s

#==> v1/ValidatingWebhookConfiguration
#istio-validator-istio-system   1     45s

#==> v1/ConfigMap
#NAME     DATA   AGE
#values   2      45s
#istio   2     45s
#istio-sidecar-injector   2     45s

#==> v1/ClusterRole
#NAME                              CREATED AT
#istiod-clusterrole-istio-system   2026-05-16T08:14:14Z
#istiod-gateway-controller-istio-system   2026-05-16T08:14:14Z
#istio-reader-clusterrole-istio-system   2026-05-16T08:14:14Z

#==> v1/ClusterRoleBinding
#NAME                              ROLE                                          AGE
#istiod-clusterrole-istio-system   ClusterRole/istiod-clusterrole-istio-system   45s
#istiod-gateway-controller-istio-system   ClusterRole/istiod-gateway-controller-istio-system   45s
#istio-reader-clusterrole-istio-system   ClusterRole/istio-reader-clusterrole-istio-system   45s

#==> v1/RoleBinding
#NAME     ROLE          AGE
#istiod   Role/istiod   45s

#==> v1/Deployment
#NAME     READY   UP-TO-DATE   AVAILABLE   AGE
#istiod   1/1     1            1           45s

#==> v1/Pod(related)
#NAME                     READY   STATUS    RESTARTS   AGE
#istiod-54c8f97d7-2qf9h   1/1     Running   0          45s

#==> v2/HorizontalPodAutoscaler
#NAME     REFERENCE           TARGETS              MINPODS   MAXPODS   REPLICAS   AGE
#istiod   Deployment/istiod   cpu: <unknown>/80%   1         5         1          45s

#==> v1/ServiceAccount
#NAME     AGE
#istiod   45s

#==> v1/Role
#NAME     CREATED AT
#istiod   2026-05-16T08:14:14Z


#TEST SUITE: None
#NOTES:
#"istiod" successfully installed!

#Next steps:
#  * Deploy a Gateway: https://istio.io/latest/docs/setup/additional-setup/gateway/

# Install an ingress gateway
helm install istio-ingress istio/gateway -n istio-ingress --wait --create-namespace
#NAME: istio-ingress
#LAST DEPLOYED: Sat May 16 15:22:48 2026
#NAMESPACE: istio-ingress
#STATUS: deployed
#REVISION: 1
#DESCRIPTION: Install complete
#RESOURCES:
#==> v1/Deployment
#NAME            READY   UP-TO-DATE   AVAILABLE   AGE
#istio-ingress   1/1     1            1           36s

#==> v1/Pod(related)
#NAME                             READY   STATUS    RESTARTS   AGE
#istio-ingress-7fc5959fc7-b7rvn   1/1     Running   0          36s

#==> v2/HorizontalPodAutoscaler
#NAME            REFERENCE                  TARGETS              MINPODS   MAXPODS   REPLICAS   AGE
#istio-ingress   Deployment/istio-ingress   cpu: <unknown>/80%   1         5         1          36s

#==> v1/ServiceAccount
#NAME            AGE
#istio-ingress   36s

#==> v1/Role
#NAME            CREATED AT
#istio-ingress   2026-05-16T08:22:49Z

#==> v1/RoleBinding
#NAME            ROLE                 AGE
#istio-ingress   Role/istio-ingress   36s

#==> v1/Service
#NAME            TYPE           CLUSTER-IP      EXTERNAL-IP     PORT(S)                                      AGE
#istio-ingress   LoadBalancer   10.110.192.16   10.110.192.16   15021:31871/TCP,80:32543/TCP,443:31225/TCP   36s


#TEST SUITE: None
#NOTES:
#"istio-ingress" successfully installed!
```
# Perform a gateway call to verify
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