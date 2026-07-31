### Certificate Signing Request Issue fail
1. Missing or Invalid signerName

Set _apiVersion: certificates.k8s.io/v1_ and specify a valid Kubernetes Signer Name such as _kubernetes.io/kube-apiserver-client_ or _kubernetes.io/kubelet-serving_
2. Incorrect Subject or Organization for Nodes

Update the Common Name (CN) to start with _system:node:_ and ensure the Organization (O) is set to _system:nodes_
3. Missing Controller Manager Flags

Ensure _cluster-signing-cert-file_ and _cluster-signing-key-file_ flags are correctly configured on your control plane's _kube-controller-manager_
```sh
ps aux | grep kube-controller-manager
# OR
kubectl describe pod -l component=kube-controller-manager -n kube-system
```
Minikube stores its cluster CA certicate and key files locally in ~/.minikube
- CA certificate: ~/.minikube/ca.crt
- CA private key: ~/.minikube/ca.key
- Client Certificate & Key: ~/minikube/profiles/minikube