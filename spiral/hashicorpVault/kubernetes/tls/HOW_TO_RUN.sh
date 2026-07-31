#! /bin/sh
###########################################################################################################
# Store the certificates and key in the kubernetes secrets store
############################################################################################################
kubectl create secret generic vault-tls \
    --from-file=tls.key=tls.key \
    --from-file=tls.crt=tls.crt \
    --from-file=k8s.ca=k8s.ca
##########################################################
# Deploy the vault cluster via Helm with certificate
###########################################################
helm repo add hashicorp https://helm.releases.hashicorp.com
helm search repo hashicorp/vault
helm install vault hashicorp/vault -f values.yaml
###########################################################
## Initialize and Unseal vault-0
###########################################################
#kubectl scale statefulset vault --replicas=2
kubectl exec vault-0 -- vault status
kubectl exec vault-0 -- vault operator init -key-shares=1 -key-threshold=1 -format=json > cluster-keys.json
export VAULT_UNSEAL_KEY=$(jq -r ".unseal_keys_b64[]" cluster-keys.json)
kubectl exec vault-0 -- vault operator unseal $VAULT_UNSEAL_KEY
############################################################
## Join the pod vault-1 to the Raft cluster
#############################################################
kubectl exec -it vault-1 -- /bin/sh
/ $ vault operator raft join -address=https://vault-1.vault-internal:8200 \
  -leader-ca-cert="$(cat /vault/vault-tls/k8s.ca)" \
  -leader-client-cert="$(cat /vault/vault-tls/tls.crt)" \
  -leader-client-key="$(cat /vault/vault-tls/tls.key)" \
  https://vault-0.vault-internal:8200
#Key       Value
#---       -----
#Joined    true
kubectl exec -it vault-1 -- vault operator unseal -ca-cert=/vault/vault-tls/k8s.ca $VAULT_UNSEAL_KEY
# Vault-0 Log
#storage.raft.autopilot: Promoting server: id=vault-1 address=vault-1.vault-internal:8201 name=vault-1
#storage.raft: updating configuration: command=AddVoter server-id=vault-1 server-addr=vault-1.vault-internal:8201 servers="[{Suffrage:Voter ID:vault-0 Address:vault-0.vault-internal:8201} {Suffrage:Voter ID:vault-1 Address:vault-1.vault-internal:8201}]"
# Vault-1 Log
# core: vault is unsealed
# core: entering standby mode
# core: parsing information for new active node: active_cluster_addr=https://vault-0.vault-internal:8201 active_redirect_addr=https://10.244.0.6:8200
# core: refreshing forwarding connection: clusterAddr=https://vault-0.vault-internal:8201
# core: clearing forwarding clients
# core: done clearing forwarding clients
# core: done refreshing forwarding connection: clusterAddr=https://vault-0.vault-internal:8201
# core.cluster-listener: creating rpc dialer: address=vault-0.vault-internal:8201 alpn=req_fw_sb-act_v1 host=fw-b7ade47f-1f83-b7a6-575c-96650f68a0a9
kubectl exec -it vault-1 -- /bin/sh
/ $ vault login
# Token (will be hidden): <Root Token>
/ $ vault operator raft list-peers
#Node       Address                        State       Voter
#----       -------                        -----       -----
#vault-0    vault-0.vault-internal:8201    leader      true
#vault-1    vault-1.vault-internal:8201    follower    true
###########################################################
## Create a secret
###########################################################
vault secrets enable -path=secret kv-v2
vault kv put secret/tls/apitest username="apiuser" password="apipwd"
vault kv get secret/tls/apitest
###########################################################
## Expose the vault service
##########################################################
kubectl port-forward service/vault 8200:8200
curl -vvI --cacert tls.crt --header "X-Vault-Token:hvs.d9KrsOlQOKwDTj68OMYfW4gQ" \
    https://127.0.0.1:8200/v1/secret/data/tls/apitest