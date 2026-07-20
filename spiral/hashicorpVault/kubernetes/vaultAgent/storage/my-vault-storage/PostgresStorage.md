### postgres-pvc vs postgres-storage-postgres
- postgres-pvc: A single, traditional claim provisioned via your StorageClass to store the entire PGDATA directory. It is simpler to set up but less ideal for demanding production databases
- postgres-storage-postgres (PVC Groups): Instead of lumping everything together, it automatically provisions and orchestrates multiple separate PVCs per pos
1. Isolated WAL volumes: You can place your postres WAL files on a separte, dedicated PVC. This prevents heavy database reads/writes on the main data volume from starving the log files, which significantly improves write-heavy OLTP performance and crash recovery
2. Separate tablespaces: It enables highly accessed tables or indexes to be moved to isolated, high-spedd disks
3. Tiered storage