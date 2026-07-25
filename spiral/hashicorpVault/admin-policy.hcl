# Mount secrets engines
path "sys/mounts/*" {
    capabilities = [ "create", "read", "update", "delete", "list"]
}
# Configure the ldap secrets engine and create roles
path "ldap/*" {
    capabilities = ["create", "read", "update", "delete", "list"]
}
# Write ACL policies
path "sys/policies/acl/*" {
    capabilities = [ "create", "read", "update", "delete", "list" ]
}
# Manage tokesn for verification
path "auth/token/create" {
    capabilities = [ "create", "read", "update", "delete", "list", "sudo" ]
}