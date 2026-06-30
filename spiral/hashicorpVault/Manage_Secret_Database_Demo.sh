#! /bin/bash
############Start Vault Serve###############################
#the dev server uses an in-memory storage backend.
vault server -dev -dev-root-token-id root -dev-tls
export VAULT_ADDR='https://127.0.0.1:8200'
export VAULT_CACERT='/tmp/vault-tls567626400/vault-ca.pem'
export VAULT_TOKEN=root
###########Create Secret###################################
vault secrets enable -path=kvv2 kv-v2
# Start Postgres database
docker run -d --name learn-postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -p 5432:5432 --rm postgres
# 1. Enable/Start a database secret engine in Vault
vault secrets enable database
# Success! Enabled the database secrets engine at: database/
# 2. Configure the database secrets engine with the connection credentials for the PostgreSQL database
vault WRITE database/config/postgresql \
    plugin_name="postgresql-database-plugin" \
    connection_url="postgresql://{{username}}:{{password}}@127.0.0.1:5432/postgres?sslmode=disable" \
    allowed_roles=readonly \
    username="root" \
    password="password"
# Success! Data written to: database/config/postgresql
vault READ database/config/postgresql
# 3. configure the PostgreSQL secrets engine with the role named readonly
# Create a role name ro
docker exec -i learn-postgres psql -U root -c "CREATE ROLE \"ro\" NOINHERIT;"
# Grant the ability to read all tabales to the role name ro
docker exec -i learn-postgres psql -U root -c "GRANT SELECT ON ALL TABLES IN SCHEMA public TO \"ro\";"
tee readonly.sql <<EOF
CREATE ROLE "{{name}}" WITH LOGIN PASSWORD '{{password}}' VALID UNTIL '{{expiration}}' INHERIT;
GRANT ro TO "{{name}}";
EOF
# The SQL statement contains the templatized fields {{name}}, {{password}}, and {{expiration}}
# Vault provides these values when creating the credentials
vault WRITE database/roles/readonly db_name=postgresql creation_statements=@readonly.sql default_ttl=1h max_ttl=24h
#Success! Data written to: database/roles/readonly
vault READ database/roles/readonly
vault READ database/creds/readonly
# Vault provides these values when creating the credentials
#Key                Value
#---                -----
#lease_id           database/creds/readonly/y0c09Z7re3BkZqQUBNfxV1eL
#lease_duration     1h
#lease_renewable    true
#password           F-mmqAg3R3hU1eF2NG8K
#username           v-token-readonly-UG6joTYzqduDI8ukjnDA-1782538945
# To verify that Vault create this user for PostgreSQL
docker exec -i learn-postgres psql -U root -c "SELECT usename, valuntil FROM pg_user;"
#                     usename                      |        valuntil        
#--------------------------------------------------+------------------------
# root                                             | 
# v-token-readonly-UG6joTYzqduDI8ukjnDA-1782538945 | 2026-06-27 06:42:30+00
