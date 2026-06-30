#! /bin/bash
vault server -dev -dev-root-token-id root -dev-tls
export VAULT_ADDR='https://127.0.0.1:8200'
export VAULT_CACERT='/tmp/vault-tls2295685568/vault-ca.pem'
export VAULT_TOKEN=root
cd oliver
terraform apply -auto-approve
#vault_policy.developer-vault-policy: Modifying... [id=developer-vault-policy]
#vault_mount.dev-secrets: Creating...
#vault_auth_backend.userpass: Creating...
#vault_policy.developer-vault-policy: Modifications complete after 0s [id=developer-vault-policy]
#vault_auth_backend.userpass: Creation complete after 0s [id=userpass]
#vault_mount.dev-secrets: Creation complete after 0s [id=dev-secrets]
#vault_generic_endpoint.denielle-user: Creating...
#vault_generic_endpoint.denielle-user: Creation complete after 0s [id=auth/userpass/users/danielle-vault-user]
vault auth list
#Path         Type        Accessor                  Description                Version
#----         ----        --------                  -----------                -------
#token/       token       auth_token_1d69bc86       token based credentials    n/a
#userpass/    userpass    auth_userpass_8c0ec27e    n/a                        n/a
vault policy list
#default
#default-ceiling
#developer-vault-policy
#root
vault secrets list
#Path               Type              Accessor                   Description
#----               ----              --------                   -----------
#agent-registry/    agent_registry    agent-registry_3afb6bfa    agent registry
#cubbyhole/         cubbyhole         cubbyhole_676f0892         per-token private secret storage
#dev-secrets/       kv                kv_4fb6bfe7                n/a
#identity/          identity          identity_1189e747          identity store
#secret/            kv                kv_f867caed                key/value secret storage
#sys/               system            system_b1b61b8f            system endpoints used for control, policy and debugging
vault read /dev-secrets/data/creds
#Key         Value
#---         -----
#data        map[password:change-default-password]
#metadata    map[created_time:2026-06-28T11:38:26.128942274Z custom_metadata:<nil> deletion_time: destroyed:false version:1]