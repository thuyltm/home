
# Enable Ingress Controller in Minikube
[Ingress Minikube](https://v1-32.docs.kubernetes.io/docs/tasks/access-application-cluster/ingress-minikube/)
```sh
minikube addons enable ingress
```
# Manual install Ingress Controller in Minikube using Helm
```sh
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm install nginx-ingress ingress-nginx/ingress-nginx
```

# Test directly from NodePort services
```sh
k patch service serviceb -p '{"spec": {"type":"NodePort"}}'
minikube service serviceb --url
curl ...
```
# Test through Ingress
1. Start the minikube tunnel first
2. Verify that the service type is Cluster IP
3. Deploy a new Ingress Controller
```sh
k apply -f caddi-cmd-ngress.yaml
# ingress.networking.k8s.io/caddi-cmd created
```
4. Confirm that an IP address is assigned to this Ingress Controller
5. Test
```sh
curl --resolve "caddi-cmd.local:80:$( minikube ip )" -i http://caddi-cmd.local/caddi-cmd/ping
#HTTP/1.1 200 OK
#Date: Fri, 15 May 2026 05:21:59 GMT
#Content-Type: application/json; charset=utf-8
#Content-Length: 18
#Connection: keep-alive

#{"message":"pong"}
curl --resolve "hello-chi.example:80:$( minikube ip )" -i http://hello-chi.example/servicea/hi
curl --resolve "hello-gorilla.example:80:$( minikube ip )" -i http://hello-gorilla.example/serviceb/hi
```

[GUIDE](https://kubernetes.github.io/ingress-nginx/user-guide/basic-usage/)