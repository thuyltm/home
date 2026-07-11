1. Ensure proper vault permission

Because every single workload service account that logs into Vault relies on delegated authentication to check its validity, its corresponding service account requires access the Kubernetes TokenReview API. Bind the standard _system:auth-delegator_ ClusterRole to Vault's Service Account
```sh
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: vault-tokenreview-binding
  namespace: default
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:auth-delegator
subjects:
- kind: ServiceAccount
  name: vault
  namespace: default
```

2. If your Vault server is hosted outside the Kubernetes cluster, you must disable the local CA/JWT behavior to ensure Vault knows it is running externally. This allow Vault to submit client's JWT as its own token to query the K8s TokenReview API
 and 
```sh
vault write auth/kubernetes/config \
    kubernetes_host="https://<YOUR_K8S_API_SERVER_URL>" \
    kubernetes_ca_cert="<PEM_ENCODED_CLUSTER_CA_CERT>" \
    disable_local_ca_jwt=true
```

3. If Vault is deployed inside your cluster, you do not need to provide a reviewer token a a CA certificate. Vault automatically default mount the short-lived token inside its own pod since kubernetes 1.24+, long-lived service account tokens are no longer auto-generated.
```sh
vault write auth/kubernetes/config \
    kubernetes_host="https://default.svc"
```