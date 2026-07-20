#! /bin/bash
# Debug generated template
# helm template vault hashicorp/vault -f helm-vault-db-values.yml --debug > rendered-output.yam
helm install vault hashicorp/vault -f helm-vault-db-values.yml
k get logs vault-0 -f
# Storage: postgresql (HA disabled)
k exec -it vault-0 -- vault status
#Key                Value
#---                -----
#Seal Type          shamir
#Initialized        false
#Sealed             true
#Total Shares       0
#Threshold          0
#Unseal Progress    0/0
#Unseal Nonce       n/a
#Version            2.0.2
#Build Date         2026-06-04T13:18:11Z
#Storage Type       postgresql
#HA Enabled         false
kubectl exec -it vault-0 -- vault operator init  -key-shares=1 -key-threshold=1 -format=json > cluster-keys.json
kubectl exec -ti vault-0 -- vault operator unseal $(jq -r ".unseal_keys_b64[]" cluster-keys.json)
#Key             Value
#---             -----
#Seal Type       shamir
#Initialized     true
#Sealed          false
#Total Shares    1
#Threshold       1
#Version         2.0.2
#Build Date      2026-06-04T13:18:11Z
#Storage Type    postgresql
#Cluster Name    vault-cluster-a5ef9232
#Cluster ID      01c42138-fce5-7afd-7018-03d82562b9f3
#HA Enabled      false
################################################################
##
##  SET A SECRET IN VAULT
##
#################################################################
kubectl exec -it vault-0 -- /bin/sh
/ $ vault login
# Token: hvs.PW8wyWXW579Mwd2PRQOnGSUA
/ $ vault secrets enable -path=internal kv-v2
/ $ vault kv put internal/database/config username="myuser" password="MyHome123"
/ $ vault kv get internal/database/config
#======= Metadata =======
#Key                Value
#---                -----
#created_time       2026-07-20T08:07:10.339989099Z
#custom_metadata    <nil>
#deletion_time      n/a
#destroyed          false
#version            1

#====== Data ======
#Key         Value
#---         -----
#password    MyHome123
#username    myuser
#################################################################
##
## CONFIGURE KUBERNETES AUTHENTICATION
##
#################################################################
kubectl cluster-info
#Kubernetes control plane is running at https://192.168.49.2:8443
#CoreDNS is running at https://192.168.49.2:8443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy
kubectl config view --minify
kubectl exec -it vault-0 -- /bin/sh
/ $ vault auth enable kubernetes
/ $ vault write auth/kubernetes/config \
    kubernetes_host=https://192.168.49.2:8443 \
    kubernetes_ca_cert=@/var/run/secrets/kubernetes.io/serviceaccount/ca.crt

/ $ vault policy write internal-app - <<EOF
path "internal/data/database/config" {
  capabilities = ["read"]
}
EOF
/ $ vault policy list
/ $ vault write auth/kubernetes/role/internal-app \
    bound_service_account_names=internal-app \
    bound_service_account_namespaces=default \
    policies=internal-app \
    ttl=24h \
    audience="https://kubernetes.default.svc.cluster.local"
/ $ vault list auth/kubernetes/role
##################################################################
##
##  Inject secrets into the pod
##
##################################################################
k apply -f sa.yaml
kubectl auth can-i --list --as=system:serviceaccount:default:internal-app
kubectl patch deployment orgchart --patch "$(cat patch-inject-secrets.yaml)"
# k get pods -l app=orgchart | tail -n+2 | awk '{print $1}' | xargs -I {} kubectl logs -f {}
k logs -f orgchart-7c699864bd-cc6kq -c vault-agent-init
# DEBUG k logs -f vault-0
k logs -f orgchart-7c699864bd-cc6kq -c vault-agent
k exec -it orgchart-7c699864bd-cc6kq -c orgchart -- cat /vault/secrets/database-config.txt
#data: map[password:MyHome123 username:myuser]
#metadata: map[created_time:2026-07-20T08:07:10.339989099Z custom_metadata:<nil> deletion_time: destroyed:false version:1]
kubectl patch deployment orgchart --patch "$(cat patch-inject-secrets-update.yaml)"
k exec -it orgchart-7db795b9c4-5j4qj -c orgchart -- cat /vault/secrets/database-config.txt
#postgresql://myuser:MyHome123@postgres-service.default.svc.cluster.local:5432/mydb?sslmode=disable(base)