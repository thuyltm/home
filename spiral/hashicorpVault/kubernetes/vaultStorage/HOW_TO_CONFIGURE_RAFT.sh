#! /bin/sh
helm install vault hashicorp/vault --values helm-vault-raft-values.yml
kubectl exec -ti vault-0 -- vault operator init -key-shares=1 -key-threshold=1 -format=json > cluster-keys.json
export VAULT_UNSEAL_KEY=$(jq -r ".unseal_keys_b64[]" cluster-keys.json)
kubectl exec -ti vault-0 -- vault operator unseal $VAULT_UNSEAL_KEY
kubectl exec -ti vault-1 -- vault operator unseal $VAULT_UNSEAL_KEY
#Error unsealing: Error making API request.
#URL: PUT http://127.0.0.1:8200/v1/sys/unseal
#Code: 400. Errors:
#* Vault is not initialized
#command terminated with exit code 2
#######################################################################
##
## Join the pod vault-1 to the Raft cluster
########################################################################
kubectl exec -it vault-1 -- /bin/sh
> vault operator raft join http://vault-0.vault-internal:8200
#Key       Value
#---       -----
#Joined    true
#Vault-1: [DEBUG] core: parsing information for new active node: active_cluster_addr=https://vault-0.vault-internal:8201 active_redirect_addr=http://10.244.0.6:8200
#[DEBUG] core: refreshing forwarding connection: clusterAddr=https://vault-0.vault-internal:8201
#[DEBUG] core: clearing forwarding clients
#]DEBUG] core: done clearing forwarding clients
#DEBUG] core: done refreshing forwarding connection: clusterAddr=https://vault-0.vault-internal:8201
#[DEBUG] core.cluster-listener: creating rpc dialer: address=vault-0.vault-internal:8201 alpn=req_fw_sb-act_v1 host=fw-ea71090b-54ce-bc6d-4860-eeb41d01afd8
#[DEBUG] core.cluster-listener: performing client cert lookup
#[TRACE] storage.raft: triggering raft config reload due to initial timeout
#[TRACE] storage.raft: reloaded raft config to set lower timeouts: config="raft.ReloadableConfig{TrailingLogs:0x2800, SnapshotInterval:120000000000, SnapshotThreshold:0x2000, HeartbeatTimeout:5000000000, ElectionTimeout:5000000000}"
# Vault-0: [TRACE] core: forwarding RPC: echo received: node_id=vault-1 applied_index=67 term=3 desired_suffrage=voter sdk_version=2.0.2 upgrade_version=2.0.2 redundancy_zone=""
> vault operator raft list-peers
#Error reading the raft cluster configuration: Error making API request.
#URL: GET http://127.0.0.1:8200/v1/sys/storage/raft/configuration
#Code: 503. Errors:
#* Vault is sealed
kubectl exec -ti vault-1 -- vault operator unseal $VAULT_UNSEAL_KEY
kubectl exec -it vault-1 -- /bin/sh
/ $ vault operator raft list-peers
#Error reading the raft cluster configuration: Error making API request.
#URL: GET http://127.0.0.1:8200/v1/sys/storage/raft/configuration
#Code: 403. Errors:
#* permission denied
/ $ vault login
# Token: hvs.3bFkaZw8vbSw7w4pyPdYpppC
/ $ vault operator raft list-peers
#Node       Address                        State       Voter
#----       -------                        -----       -----
#vault-0    vault-0.vault-internal:8201    leader      true
#vault-1    vault-1.vault-internal:8201    follower    true
#vault-2    vault-2.vault-internal:8201    follower    true
> $ vault operator raft autopilot state
#Healthy:                         true
#Failure Tolerance:               1
#Leader:                          vault-0
#Voters:
#   vault-0
#   vault-1
#   vault-2
#Servers:
#   vault-0
#      Name:              vault-0
#      Address:           vault-0.vault-internal:8201
#      Status:            leader
#      Node Status:       alive
#      Healthy:           true
#      Last Contact:      0s
#      Last Term:         3
#      Last Index:        83
#      Version:           2.0.2
#      Node Type:         voter
#   vault-1
#      Name:              vault-1
#      Address:           vault-1.vault-internal:8201
#      Status:            voter
#      Node Status:       alive
#      Healthy:           true
#      Last Contact:      1.654361137s
#      Last Term:         3
#      Last Index:        83
#      Version:           2.0.2
#     Node Type:         voter
#   vault-2
#      Name:              vault-2
#      Address:           vault-2.vault-internal:8201
#      Status:            voter
#      Node Status:       alive
#      Healthy:           true
#      Last Contact:      2.56602958s
#      Last Term:         3
#      Last Index:        83
#      Version:           2.0.2
#      Node Type:         voter