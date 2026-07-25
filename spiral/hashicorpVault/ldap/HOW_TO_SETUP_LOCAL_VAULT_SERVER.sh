#! /bin/bash
#########################################################################################
### 
### Install binary library
###########################################################################################
wget -O - https://apt.releases.hashicorp.com/gpg | sudo gpg --dearmor -o /usr/share/keyrings/hashicorp-archive-keyring.gpg
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(grep -oP '(?<=UBUNTU_CODENAME=).*' /etc/os-release || lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/hashicorp.list
sudo apt update && sudo apt install vault
##########################################################################################
###
### Setup a Local Vault Server
### A local Vault server use in-memory database
##########################################################################################
vault server -dev -dev-root-token-id root
export VAULT_ADDR=http://127.0.0.1:8200
export VAULT_TOKEN=root
export OPENLDAP_URL=127.0.0.1:389
###########################################################################################
###
### Login with administration priviledge
###########################################################################################
vault login
# Token (will be hidden): root
vault policy write admin admin-policy.hcl
vault token create -policy="admin"
#Key                  Value
#---                  -----
#token                hvs.CAESIAw7Hi4l5lXm8lJmm1FLwx5gaHGCk7TQTCnX_Sn399inGh4KHGh2cy4zWVhWTXhzRDFCMWFwTWhaY1ZLV0piZ1E
vault login
# Token (will be hidden): hvs.CAESIAw7Hi4l5lXm8lJmm1FLwx5gaHGCk7TQTCnX_Sn399inGh4KHGh2cy4zWVhWTXhzRDFCMWFwTWhaY1ZLV0piZ1E
###########################################################################################
###
### Enable the LDAP secrets engine
###########################################################################################
vault secrets enable ldap
###########################################################################################
###
### Configure LDAP secrets engine
###########################################################################################
vault write ldap/config \
    binddn=cn=admin,dc=learn,dc=example \
    bindpass=2LearnedVault \
    url=ldap://$OPENLDAP_URL
###########################################################################################
###
### Create a role
###########################################################################################
vault write ldap/static-role/learn \
    dn='cn=alice,ou=users,dc=learn,dc=example' \
    username='alice' \
    rotation_period="24h"
###########################################################################################
###
### Request OpenLDAP credentials
###########################################################################################
LDAP_PASSWORD=$(vault read --format=json ldap/static-cred/learn | jq -r ".data.password")
ldapsearch -b "cn=alice,ou=users,dc=learn,dc=example" \
    -D 'cn=alice,ou=users,dc=learn,dc=example' \
    -w $LDAP_PASSWORD