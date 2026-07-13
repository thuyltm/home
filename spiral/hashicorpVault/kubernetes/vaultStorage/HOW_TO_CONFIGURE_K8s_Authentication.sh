#! /bin/bash
kubectl exec -ti vault-0 -- vault status
kubectl exec -ti vault-0 -- vault operator unseal $(jq -r ".unseal_keys_b64[]" cluster-keys.json)
kubectl get node -o wide
#NAME       STATUS   ROLES           AGE   VERSION   INTERNAL-IP    EXTERNAL-IP   OS-IMAGE                         KERNEL-VERSION      CONTAINER-RUNTIME
#minikube   Ready    control-plane   19d   v1.35.1   192.168.49.2   <none>        Debian GNU/Linux 12 (bookworm)   6.17.0-20-generic   docker://29.2.1
kubectl cluster-info
export K8S_SERVICE_ACCOUNT_TOKEN=$(k exec sample-deployment-57c9596999-bjdrw -- cat /var/run/secrets/kubernetes.io/serviceaccount/vault-token | cut -d. -f2)
# View kubectl exec sample-deployment-57c9596999-bjdrw -- cat /var/run/secrets/kubernetes.io/serviceaccount/vault-token | cut -d. -f2 | base64 --decode
kubectl exec --stdin=true --tty=true vault-0 -- /bin/sh
/ $ vault login
# Token (will be hidden): hvs.l9l8JHeGV50o8wSPc3y9LuSk
/ $ vault auth list
/ $ vault AUTH enable kubernetes
# Success! Enabled kubernetes auth method at: kubernetes/
# When Vault is running in a Kubernetes cluster, the Kubernetes auth method can be configured to use the service account token of the pod that Vault is running in. 
/ $ cat /var/run/secrets/kubernetes.io/serviceaccount/token | cut -d.  | base64 -d
# https://www.jwt.io/
/ $ vault write auth/kubernetes/config \
    kubernetes_host="https://192.168.49.2:8443"
# OR Vault is running outside of the Kubernetes cluster
/ $ vault write auth/kubernetes/config \
    kubernetes_host="https://192.168.49.2:8443" \
    kubernetes_ca_cert=@/var/run/secrets/kubernetes.io/serviceaccount/ca.crt
    disable_local_ca_jwt=true
    disable_iss_validation=true

/ $ vault write auth/kubernetes/role/hashicupsapp \
     bound_service_account_names=vault \
     bound_service_account_namespaces=default \
     policies=default,dev-secrets \
     ttl=1h \
     audience="https://kubernetes.default.svc.cluster.local"

/ $ vault list auth/kubernetes/role
/ $ vault read auth/kubernetes/role/hashicupsapp

kubectl exec -it sample-deployment-57c9596999-bjdrw -- /bin/bash
/ $ export K8S_SERVICE_ACCOUNT_TOKEN=$(cat /var/run/secrets/kubernetes.io/serviceaccount/vault-token)
/ $ curl -H "X-Vault-Token: hvs.l9l8JHeGV50o8wSPc3y9LuSk" \
        --header "X-Vault-Namespace: default" \
        --request POST \
        --data '{"jwt": "'"$K8S_SERVICE_ACCOUNT_TOKEN"'", "role": "hashicupsapp"}' \
        http://vault.default.svc.cluster.local:8200/v1/auth/kubernetes/login

#{
#   "request_id":"67784044-d0d4-9678-6793-1880a7f71a36",
#   "lease_id":"",
#   "renewable":false,
#   "lease_duration":0,
#   "data":null,
#   "wrap_info":null,
#   "warnings":null,
#   "auth":{
#      "client_token":"hvs.CAESIB3h5WZ_pY5W2SZI-qbLXAUcO9yoK6FfpRRU2fF41mkRGh4KHGh2cy5QS3BSME5KZWhFZUtOMDZiME1USTlMa1M",
#      "accessor":"u8POutAKUBjFfWQzEGXGNz8O",
#      "policies":[
#         "default",
#         "dev-secrets"
#      ],
#      "token_policies":[
#         "default",
#         "dev-secrets"
#      ],
#      "metadata":{
#         "role":"hashicupsapp",
#         "service_account_name":"vault",
#         "service_account_namespace":"default",
#         "service_account_secret_name":"",
#         "service_account_uid":"f69f07f4-711f-4c4a-98c5-41a924fcfc79"
#      },
#      "lease_duration":3600,
#      "renewable":true,
#      "entity_id":"858cb14f-6fc3-ff2a-ebce-2c372b24bc00",
#      "token_type":"service",
#      "orphan":true,
#      "mfa_requirement":null,
#      "num_uses":0
#   },
#   "mount_type":""
#}
# enable an instance of the KV secrets engine at the path secret/ with version 2.
vault secrets enable -path=secret kv-v2
# create a secret at path secret/data/dev-secrets/config with the key-value pair username=devuser and password=devpass.
vault kv put secret/dev-secrets/config username=devuser password=devpass
# verify the deifinition of the secret at the path secret/dev-secrets/config
vault kv get secret/dev-secrets/config
# Write out the policy named dev-secrets that enables the read capability for the secrets at path secret/data/dev/*.
vault policy write dev-secrets - <<EOF
path "secret/data/dev-secrets/config" {
  capabilities = ["read"]
}
EOF

