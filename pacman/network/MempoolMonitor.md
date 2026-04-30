There're no centralized nodes in the bitcoin network, all transactions and all blocks are delivered to every node in the network

### Transactions and blocks transfer in this way:
1. When a node gets a new transaction, it sends 'inv' message to its peers. 'inv' means inventory. 'Inv' doesn't contain full data, only hashes
2. Any peer receiving the message can decide whether to fetch this data or not
3. If a peer wants full data, it sends a 'getdata' reply specifying a list of hashes it want to get full data for
4. Upon receiving a 'getdata' message, a node checks the object types, transactions or blocks, and send them using the appropriate 'tx' for transaction and 'block' for block message

### Mempool 
Mempool is simply a list of transactions that haven't been mined yet. Mempool keeps all unconfirmed transactions in memory. As soon a transaction was added to a block and the block was mined, the transaction gets removed from mempool

Every node has its own mempool. While there might be discrepancies between mempools of different nodes, this doesn'r bring any danger or harm

When you create, sign and send a new transaction, it's received by all other nodes and is stored in their mempools. When a new block containing your transaction is mined, it gets delivered to all nodes and is removed from its mempool

Implement a JSON-RPC server for mempool monitoring