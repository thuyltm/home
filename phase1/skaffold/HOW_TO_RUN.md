Start K8s cluster
```sh
minikube start
```
Create skaffold.yaml
```sh
skaffold init
```
Deploy k8s-service.yaml and k8s-deployment.yaml into minikube

Note: 
- NodePort for deploying in minikube
- LoadBalance for deploying in Cloud
```sh
skaffold run --tail
```
The easiest way to access this service is to let minikube launch a web browser for you
```sh
minikube service my-skaffold-app-service
```

Webrowser: http://192.168.49.2:30000/ping