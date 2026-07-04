An unseal key in HashiCorp Vault is a cryptographic key used to reconstruct the master key needed to decrypt the database encryption key.

Master/Root Key: The keyring itself is **encrypted by a master/root key**

Because **Vault always starts in a sealed state for security**, you must use unseal keys to unlock the server and access your stored data. Vault remains locked by default. Every time Vault restarts, the master key must be reconstructed using the unseal keys to decrypt the internal encryption layer. Only after unsealing, you use the inital root token to log in.

By default, Vault uses an algorithm called Shamir's Secret Sharing to divide the master key into multiple key shares. TO unseal the server, a specific threshold of these unique key shares must be entered. This prevents a single person from holding the power to unlock the entire vault
```sh
vault operator unseal
```
Vault will then prompt you to enter the unseal key securely. You will repeat this process with different key shares until the required threshold is met

Unseal keys should be distributed to different trusted individuals to ensure high security anc compliance