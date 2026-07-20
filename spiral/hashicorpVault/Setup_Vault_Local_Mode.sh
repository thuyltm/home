#! /bin/bash
# Install binary library
wget -O - https://apt.releases.hashicorp.com/gpg | sudo gpg --dearmor -o /usr/share/keyrings/hashicorp-archive-keyring.gpg
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(grep -oP '(?<=UBUNTU_CODENAME=).*' /etc/os-release || lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/hashicorp.list
sudo apt update && sudo apt install vault
# dev mode in-memory storage backend
vault server -dev -dev-root-token-id root -dev-tls
# export VAULT_ADDR='https://127.0.0.1:8200'
# export VAULT_CACERT='/tmp/vault-tls3035908606/vault-ca.pem'
# The unseal key and root token are displayed below in case you want to
# seal/unseal the Vault or re-authenticate.

# Unseal Key: +j8wpZgZ802Gnp89s9PUS/BAa6NJ7otKH167F/AQP8c=
# Root Token: root
vault status
#Key             Value
#---             -----
#Seal Type       shamir
#Initialized     true
#Sealed          false
#Total Shares    1
#Threshold       1
#Version         2.0.3
#Build Date      2026-06-17T12:39:45Z
#Storage Type    inmem
#Cluster Name    vault-cluster-b1312512
#Cluster ID      c4b17da2-1ea7-caa7-48f2-51ef40459577
#HA Enabled      false
vault login
```sh
Token (will be hidden):
#Type in the literal string root, and press RETURN.
```
vault token lookup -accessor nyBHfiYFMhMranowfQy6dbf5
vault secrets list
#Path               Type              Accessor                   Description
#----               ----              --------                   -----------
#agent-registry/    agent_registry    agent-registry_87b39561    agent registry
#cubbyhole/         cubbyhole         cubbyhole_94a82b74         per-token private secret storage
#identity/          identity          identity_1d5d0bcf          identity store
#secret/            kv                kv_05a3c8b5                key/value secret storage
#sys/               system            system_4733e53c            system endpoints used for control, policy and debugging