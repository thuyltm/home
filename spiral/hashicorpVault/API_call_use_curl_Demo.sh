#! /bin/bash
vault server -dev -dev-root-token-id root -dev-tls
export VAULT_ADDR='https://127.0.0.1:8200'
export VAULT_CACERT='/tmp/vault-tls2939726587/vault-ca.pem'
export VAULT_TOKEN=root
export CURL_CA_BUNDLE=$VAULT_CACERT
# this is how to enable the userpass auth method by invoking the Vault HTTP API
curl -H "X-Vault-Token: $VAULT_TOKEN" -X POST -d '{"type":"userpass"}' $VAULT_ADDR/v1/sys/auth/userpass
# To determine if it was successful
curl -H "X-Vault-Token: $VAULT_TOKEN" $VAULT_ADDR/v1/sys/auth | jq ".data"
# Create an user in userpass
curl -H "X-Vault-Token: $VAULT_TOKEN" -X POST \
 -d '{"password":"Imprint Bacteria Marathon Aflutter","token_policies":"developer-vault-policy"}' \
  $VAULT_ADDR/v1/auth/userpass/users/danielle-vault-user
# Create ACL policies
curl -H "X-Vault-Token: $VAULT_TOKEN" \
    -X PUT \
    -d '{"policy":"path \"dev-secrets/data/creds\" {\n  capabilities = [\"create\", \"update\"]\n}\n\npath \"dev-secrets/data/creds\" {\n  capabilities = [\"read\"]\n}\n"}' \
    $VAULT_ADDR/v1/sys/policies/acl/developer-vault-policy
# List the Vault policies
curl -s -H "X-Vault-Token: $VAULT_TOKEN" $VAULT_ADDR/v1/sys/policy | jq ".data.policies"
#[
#  "default",
#  "default-ceiling",
#  "developer-vault-policy",
#  "root"
#]
# check which secrets engines are already mounted
curl -s -H "X-Vault-Token: $VAULT_TOKEN" $VAULT_ADDR/v1/sys/mounts | jq ".data"
# create a secrets engine at dev-secrets
curl -H "X-Vault-Token: $VAULT_TOKEN" -X POST -d '{ "type":"kv-v2" }' $VAULT_ADDR/v1/sys/mounts/dev-secrets
# create a login username danielle-vault-user
curl -s -X POST -d '{ "password": "Imprint Bacteria Marathon Aflutter" }' $VAULT_ADDR/v1/auth/userpass/login/danielle-vault-user | jq
#{
#   "request_id":"ffc19c9e-7e58-07a7-3090-2eea42473509",
#   "lease_id":"",
#   "renewable":false,
#   "lease_duration":0,
#   "data":null,
#   "wrap_info":null,
#   "warnings":null,
#   "auth":{
#      "client_token":"hvs.CAESINo7OvEtsCIiBSN5fGj5FlSmTN8Bvm7Xtk6xrrm8X7ncGh4KHGh2cy4zYUdKQWU4dHZVUjhKTHBpZWNaOHBiN0M",
#      "accessor":"m4tSAhiUkWW4HzsVKmSlO4cH",
#      "policies":[
#         "default",
#         "developer-vault-policy"
#      ],
#      "token_policies":[
#         "default",
#         "developer-vault-policy"
#      ],
#      "metadata":{
#         "username":"danielle-vault-user"
#      },
#      "lease_duration":2764800,
#      "renewable":true,
#      "entity_id":"99a63f97-fbe3-48cf-7e9b-6978712e26e4",
#      "token_type":"service",
#      "orphan":true,
#      "mfa_requirement":null,
#      "num_uses":0
#   },
#   "mount_type":""
#}
export AUTH_CLIENT_TOKEN='hvs.CAESINo7OvEtsCIiBSN5fGj5FlSmTN8Bvm7Xtk6xrrm8X7ncGh4KHGh2cy4zYUdKQWU4dHZVUjhKTHBpZWNaOHBiN0M'
curl -s -H "X-Vault-Token: $AUTH_CLIENT_TOKEN" -X PUT -d '{ "data": {"password": "Driven Siberian Pantyhose Equinox"} }' $VAULT_ADDR/v1/dev-secrets/data/creds | jq
#{
#  "request_id": "3ef906ce-4f12-77da-ba79-06cc1bfa0d54",
#  "lease_id": "",
#  "renewable": false,
#  "lease_duration": 0,
#  "data": {
#    "created_time": "2026-06-30T04:23:50.153952867Z",
#    "custom_metadata": null,
#    "deletion_time": "",
#    "destroyed": false,
#    "version": 2
#  },
#  "wrap_info": null,
#  "warnings": null,
#  "auth": null,
#  "mount_type": "kv"
#}
curl -s -H "X-Vault-Token: $AUTH_CLIENT_TOKEN" $VAULT_ADDR/v1/dev-secrets/data/creds | jq
#{
#  "request_id": "c993ff7f-bf23-280c-5981-419e492f9f71",
#  "lease_id": "",
#  "renewable": false,
#  "lease_duration": 0,
#  "data": {
#    "data": {
#      "password": "Driven Siberian Pantyhose Equinox"
#    },
#    "metadata": {
#      "created_time": "2026-06-30T04:23:50.153952867Z",
#      "custom_metadata": null,
#      "deletion_time": "",
#      "destroyed": false,
#      "version": 2
#    }
#  },
#  "wrap_info": null,
#  "warnings": null,
#  "auth": null,
#  "mount_type": "kv"
#}