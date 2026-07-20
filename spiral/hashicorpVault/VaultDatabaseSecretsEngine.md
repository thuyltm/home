HashiCorp Vault dynamically generate and manage credentials for your PostgresSQL database. Vault connects to Postgres as an admin and automatically create temporary database users or rotates existing ones.
- Dynamic Secrets: Vault creates a unique, ephemeral PostgreSQL user every time an app requests access. The user automatically expires when ist Time-To-Live finishes
- Static Secrets: Vault manages a set database user and automatically rotates its password at pre-defined intervals
A PostgreSQL database act as the encrypted data storage layer for a self-hosted HashiCorp Vault server