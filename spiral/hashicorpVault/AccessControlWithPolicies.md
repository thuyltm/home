1. Vault create a root policy during initialization. The root policy, attached to the root token when server initialization completes, provides an initial superuser to enable secrets engines, define policies, and configure authentication methods
2. Vault create a default policy.

Vault uses policies to govern the behavior of clients and instrument Role-Based Access Control. A policy defines a list of paths. Each path expresses the allowed capabilities. You must define capabilities for a path