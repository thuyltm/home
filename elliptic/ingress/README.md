
# Enable Ingress Controller in Minikube
[Ingress Minikube](https://v1-32.docs.kubernetes.io/docs/tasks/access-application-cluster/ingress-minikube/)
```sh
minikube addons enable ingress
```
# Test directly from NodePort services
```sh
k patch service serviceb -p '{"spec": {"type":"ClusterIP"}}'
minikube service serviceb --url
curl ...
```
# Test through Ingress
```sh
curl --resolve "hello-chi.example:80:$( minikube ip )" -i http://hello-chi.example/servicea/hi
curl --resolve "hello-gorilla.example:80:$( minikube ip )" -i http://hello-gorilla.example/serviceb/hi
```