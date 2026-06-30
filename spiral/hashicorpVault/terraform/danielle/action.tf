# a vault_kv_secret_v2 resource creates a secret named creds in the dev-secrets secret engine
resource "vault_kv_secret_v2" "creds" {
    mount = "dev-secrets"
    name = "creds"
    data_json = jsonencode(
        {
            password = "change-default-password",
        }
    )
}
