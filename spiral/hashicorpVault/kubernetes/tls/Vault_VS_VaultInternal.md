In HashiCorp Vault deployments on Kubernetes using Helm, vault and vault-internal are two distinct services
- vault: Exposes a stable cluster IP or load balancer for external client requests
- vault-internal: A headless internal service used for intra-cluster communication, pod-to-pod discovery, and Raft consensus join operations