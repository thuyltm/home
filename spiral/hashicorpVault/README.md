HashiCorp Vault is an enterprise-grade secrets management tool used to securely store, access, and control sensitive data like passwords, API keys, digital certificates, and encryption keys.

[Install](https://developer.hashicorp.com/vault/install)
[Setup](https://developer.hashicorp.com/vault/tutorials/get-started/setup)

### Dev server
the dev server uses an in-memory storage backend. This means that when you stop a dev server, you lose access to any data you wrote to the server

### Token
A token validates a Vault client access to Vault and what actions the client can perform. Most actions in Vault require a token

### Policy
Policies provide a declarative way to grant or forbid access to operations in Vault. Everything in Vault is path based. Policies determines the operations that you need to configure on a specific path

### Dynamic secrets
Dynamic secrets do not exist until read, so the risk of being stolen is greatly reduced. Because Vault has built-in revocation mechanisms, Vault revokes dynamic secrets after use