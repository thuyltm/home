#! /bin/sh
kind create cluster --config kind-config.yaml
kubectl cluster-info --context kind-kind
#Kubernetes control plane is running at https://127.0.0.1:33963
#CoreDNS is running at https://127.0.0.1:33963/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy
kubectl get nodes
#NAME                 STATUS   ROLES           AGE     VERSION
#kind-control-plane   Ready    control-plane   2m20s   v1.36.1
#kind-worker          Ready    <none>          2m10s   v1.36.1
kubectl apply -f https://projectcontour.io/quickstart/contour.yaml
kubectl api-resources --verbs=list --namespaced -o name | xargs -n 1 kubectl get --show-kind --ignore-not-found -n projectcontour
#NAME              TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)                      AGE
#service/contour   ClusterIP      10.96.58.193    <none>        8001/TCP                     43m
#service/envoy     LoadBalancer   10.96.215.239   172.25.0.5    80:31137/TCP,443:30292/TCP   43m

kubectl apply -f https://projectcontour.io/examples/httpbin.yaml
#deployment.apps/httpbin created
#service/httpbin created
#ingress.networking.k8s.io/httpbin created
#NAME      CLASS     HOSTS   ADDRESS      PORTS   AGE
#httpbin   contour   *       172.25.0.5   80      42m
kubectl port-forward service/envoy 8888:80 -n projectcontour