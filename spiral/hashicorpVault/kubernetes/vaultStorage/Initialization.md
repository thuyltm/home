In HashiCorp Vault, initialization is a mandatory one-time setup process for a new cluster. Upon initialization,the system performs these actions
- Storage Backend Prep: It prepares the configured storage (database or cloud storage) to receive encrypted data
- __Master Key__ Generation: It generates an internal master key that encrypts the data stored within Vault
- Sharding (__Unseal Key__): It splits the master key into multiple shards and defines a threshold needed to reconstruct the key
- Initial __Root Token__: It generates a master "root token" with full administractive privileges to the Vault

1. If you lose your unseal keys or the inital root token, Vault provides no backdoor

You can only view and obtain HashiCorp Vault unseal keys immediately during the initial execution of the _vault operator init_ command. Vault prints these keys to your terminal output exactly once. Vault never stores the unseal keys plaintext internally, making it impossible to query or recover them from a running server after initialization has finished.

2. If you delete your cluster/pod but kept the old Persistent Volumnes, Vault admit the data is already intialized but sealed, blocking new initialization
```sh
k delete pvc -l app.kubernetes.io/instance=vault
```
Follow the steps below to reinitialize Vault

2.1. Stop the Vault service or container
```sh
kubectl scale statefulset vault --replicas=0
kubectl scale deployment vault --replicas=0
```
2.2. Wipe the physical storage path
```sh
k delete pvc -l app.kubernetes.io/instance=vault
```
2.3. Start the Vault service again
```sh
kubectl scale statefulset vault --replicas=1
kubectl scale deployment vault-agent-injector --replicas=1
```
2.4. Check Vault status
```sh
kubectl exec vault-0 -- vault status
#Key                     Value
#---                     -----
#Seal Type               shamir
#Initialized             false
#Sealed                  true
#Total Shares            0
#Threshold               0
#Unseal Progress         0/0
#Unseal Nonce            n/a
#Version                 2.0.2
#Build Date              2026-06-04T13:18:11Z
#Storage Type            raft
#Removed From Cluster    false
#HA Enabled              true
```
3. a Vault pod stuck in a 0/1 READY state which means the status already initialized but sealed, you simply need to unseal the Vault instance. Readiness probes fail when a pod is sealed.
```sh
k exec -it vault-0 -- vault status
---                     -----
Seal Type               shamir
Initialized             true
------------------------------
|Sealed                  true|
------------------------------
Total Shares            1
Threshold               1
-----------------------------
|Unseal Progress         0/1|
-----------------------------
Unseal Nonce            n/a
Version                 2.0.2
Build Date              2026-06-04T13:18:11Z
Storage Type            raft
Removed From Cluster    false
HA Enabled              true
```