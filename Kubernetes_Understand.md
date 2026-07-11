1> Deployment vs Stateful
- a deployment is designed for __stateless__ applications where pods are identical and interchangeable. Deployments usually use temporary storage that is wiped if a pod restarts.
- a statefulset is built for __stateful__ applications (like databases) that require persistent data. StatefulSets natively handle persistent volume claims (PVCs) for each individual pod, ensuring data is never lost.

2> To find the internal IP address of a Kubernetes hose (Node), run the below command
```sh
k get nodes -o wide
#NAME       STATUS   ROLES           AGE   VERSION   INTERNAL-IP    EXTERNAL-IP   OS-IMAGE                         KERNEL-VERSION      CONTAINER-RUNTIME
#minikube   Ready    control-plane   15d   v1.35.1   192.168.49.2   <none>        Debian GNU/Linux 12 (bookworm)   6.17.0-20-generic   docker://29.2.1
```

3> The __system:auth-delegator__ is a __built-in Kubernetes ClusterRole__ that grants a service account the permissions necessary to query the TokenReview API. It allows delegated authentication, letting services securely verify if a client's Kubernetes service account token is valid.

In Kubernetes, an account name itself does not have a default, fixed role. An account name is connected to a specific Role or ClusterRole using RoleBindings.

4> To get the Token and CA certificate (ca.crt) for a Kubernetes ServiceAccount

Method 1: This is a standard approach to fetch a short-lived token alongside the cluster's CA certificate
```sh
kubectl create token <YOUR_SERVICE_ACCOUNT_NAME> -n <YOUR_NAMESPACE>
kubectl config view --raw --flatten --minify -o jsonpath='{.clusters[0].cluster.certificate-authority-data}' | base64 -d > ca.crt
```
Method 2: If you need a long-live token, you must manually bind a Secret to the ServiceAccount
1. Create the token secret
```sh
apiVersion: v1
kind: Secret
metadata:
    name: sa-token-secret
    namespace: <YOUR_NAMESPACE>
    annotations:
        kubernetes.io/service-account.name: <YOUR_SERVICE_ACCOUNT_NAME>
type: 
    kubernetes.io/service-account-token
```
2. Retrieve the Token and CA Certificate
```sh
kubectl get secret sa-token-secret -n <YOUR_NAMESPACE> -o jsonpath='{.data.token}' | base64 -d
kubectl get secret sa-token-secret -n <YOUR_NAMESPACE> -o jsonpath='{.data.ca\.crt}' | base64 -d > ca.crt
kubectl create token <YOUR_SERVICE_ACCOUNT_NAME>  --audience="https://kubernetes.default.svc.cluster.local" -n <YOUR_NAMESPACE>
```
Method 3: From inside a running pod

You can view or use them instantly at the following paths:
- CA certificate: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
- Bearer Token: /var/run/secrets/kubernetes.io/serviceaccount/token

To copy a CA certificate from a Kubernetes Secret to a local file. The CA certificates from the sa-token-secret Secret and the Kubernetes cluter are identical
```sh
kubectl get secret sa-token-secret -o "jsonpath={.data['ca\.crt']}" | base64 -d > ca-sa.crt
kubectl config view --raw --flatten --minify -o jsonpath='{.clusters[0].cluster.certificate-authority-data}' | base64 -d > ca.crt
```
5> Invalid bearer token ServiceAccount UID ($uid) does not match claim ($uid-xxxxx-service-account)

The error "ServiceAccount UID does not match claim" occurs because the ServiceAccount was deleted and recreated which changes it a new unique identifier (UID). Meanwhile a running Pod is still trying to use an authentication token bound to the old UID

Because Kubernetes service account tokens are cryptographically bound to the specific UID of the object, matching names is not enough
```sh
# 1. Get the live ServiceAccount UID
kubectl get serviceaccount <sa-name> -n <namespace> -o jsonpath='{.metadata.uid}'
# 2. View the claim UID inside the pod's token
kubectl exec <pod-name> -n <namespace> -- cat /var/run/secrets/kubernetes.io/serviceaccount/token | cut -d. -f2 | base64 --decode
```