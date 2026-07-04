Integrated Storage (Raft) uses a built-in consensus algorithm that stores data directly on Vault nodes because all the nodes in a Vault cluster will have a replicated copy of Vault's data

Operators can manually add a node by executing the join command and pointing it to the active leader node
```sh
vault operator raft join http://<leader-IP>:8200
```