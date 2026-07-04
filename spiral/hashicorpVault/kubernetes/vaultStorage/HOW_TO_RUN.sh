#! /bin/bash
minikube start
helm repo add hashicorp https://helm.releases.hashicorp.com
helm repo update
helm search repo hashicorp/vault
helm install vault hashicorp/vault --values helm-vault-raft-values.yml
#RESOURCES:
#==> v1/Deployment
#NAME                   READY   UP-TO-DATE   AVAILABLE   AGE
#vault-agent-injector   0/1     1            0           14s

#==> v1/MutatingWebhookConfiguration
#NAME                       WEBHOOKS   AGE
#vault-agent-injector-cfg   1          14s

#==> v1/ServiceAccount
#NAME                   AGE
#vault-agent-injector   14s
#vault   14s

#==> v1/Role
#NAME                   CREATED AT
#vault-discovery-role   2026-07-01T02:53:04Z

#==> v1/Service
#NAME                       TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)   AGE
#vault-agent-injector-svc   ClusterIP   10.104.145.114   <none>        443/TCP   14s
#vault-active   ClusterIP   10.98.45.168   <none>   8200/TCP,8201/TCP   14s
#vault-standby   ClusterIP   10.104.2.79   <none>   8200/TCP,8201/TCP   14s
#vault-internal   ClusterIP   None   <none>   8200/TCP,8201/TCP   14s
#vault   ClusterIP   10.98.183.108   <none>   8200/TCP,8201/TCP   14s

#==> v1/Pod(related)
#NAME                                    READY   STATUS              RESTARTS   AGE
#vault-agent-injector-6d6f756749-d4gf4   0/1     ContainerCreating   0          14s
#vault-0   0/1   ContainerCreating   0     14s
#vault-1   0/1   ContainerCreating   0     14s
#vault-2   0/1   ContainerCreating   0     14s

#==> v1/StatefulSet
#NAME    READY   AGE
#vault   0/3     14s

#==> v1/PodDisruptionBudget
#NAME    MIN AVAILABLE   MAX UNAVAILABLE   ALLOWED DISRUPTIONS   AGE
#vault   N/A             1                 0                     14s

#==> v1/ConfigMap
#NAME           DATA   AGE
#vault-config   1      14s

#==> v1/ClusterRole
#NAME                               CREATED AT
#vault-agent-injector-clusterrole   2026-07-01T02:53:04Z

#==> v1/ClusterRoleBinding
#NAME                           ROLE                                           AGE
#vault-agent-injector-binding   ClusterRole/vault-agent-injector-clusterrole   14s
#vault-server-binding   ClusterRole/system:auth-delegator   14s

#==> v1/RoleBinding
#NAME                          ROLE                        AGE
#vault-discovery-rolebinding   Role/vault-discovery-role   14s

###############################################################################################################
# Initialize Vault
################################################################################################################
kubectl exec -ti vault-0 -- vault operator init -key-shares=1 -key-threshold=1 -format=json > cluster-keys.json
#NAME                                    READY   STATUS    RESTARTS   AGE
#vault-0                                 0/1     Running   0          41m
export VAULT_UNSEAL_KEY=$(jq -r ".unseal_keys_b64[]" cluster-keys.json)
export ROOT_TOKEN=$(jq -r ".root_token" cluster-keys.json)
# Use unseal key to unlock server and access your stored data
kubectl exec -ti vault-0 -- vault operator unseal $VAULT_UNSEAL_KEY
#Key                     Value
#---                     -----
#Seal Type               shamir
#Initialized             true
#Sealed                  false
#Total Shares            1
#Threshold               1
#Version                 2.0.2
#Build Date              2026-06-04T13:18:11Z
#Storage Type            raft
#Cluster Name            vault-integrated-storage
#Cluster ID              07a14e91-fb42-9dcb-b712-8e601b930376
#Removed From Cluster    false
#HA Enabled              true
#HA Cluster              https://vault-0.vault-internal:8201
#HA Mode                 active
#Active Since            2026-07-04T07:37:35.541803067Z
#Raft Committed Index    38
#Raft Applied Index      38
#NAME                                    READY   STATUS    RESTARTS   AGE
#vault-0                                 1/1     Running   0          41m
kubectl exec -ti vault-1 -- vault operator init -key-shares=1 -key-threshold=1 -format=json > cluster-keys-1.json
kubectl exec -ti vault-1 -- vault operator unseal $(jq -r ".unseal_keys_b64[]" cluster-keys-1.json)
# Create Raft cluster
kubectl exec -ti vault-1 -- vault operator raft join http://vault-0.vault-internal:8200
###################################################################################################################
# Set a secret in Vault
###################################################################################################################
kubectl exec --stdin=true --tty=true vault-0 -- /bin/sh
# Vault requires clients to authenticate first before it allows any further actions
# Login with the root token when prompted. This is ROOT_TOKEN
vault login
vault secrets list
#Path               Type              Accessor                   Description
#----               ----              --------                   -----------
#agent-registry/    agent_registry    agent-registry_664465ae    agent registry
#cubbyhole/         cubbyhole         cubbyhole_24046979         per-token private secret storage
#identity/          identity          identity_7e55371e          identity store
#sys/               system            system_f1a10237            system endpoints used for control, policy and debugging
# Enable an instance of the kv-v2 secrets engine at the path secret
vault secrets enable -path=secret kv-v2
# Success! Enabled the kv-v2 secrets engine at: secret/
# create a secret at path secret/webapp/config
vault kv put secret/webapp/config username="static-user" password="static-password"
# verify the secret at the path secret/webapp/config
vault kv get secret/webapp/config
#====== Secret Path ======
#secret/data/webapp/config

#======= Metadata =======
#Key                Value
#---                -----
#created_time       2026-07-04T08:11:14.449289305Z
#custom_metadata    <nil>
#deletion_time      n/a
#destroyed          false
#version            1

#====== Data ======
#Key         Value
#---         -----
#password    static-password
#username    static-user