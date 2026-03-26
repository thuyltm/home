### Blockchain Network
Blockchain network is decentralized. There are nodes, and each node is a full-fledged member of the network. Nodes are connected directly to each other, which means that Blockchain network is a P2P network and there are no hierarchy in node roles
[!P2P-Network](https://jeiwan.net/images/p2p-network.png)

Nodes in such network are more difficult to implement, because they have to perform a lot of operations
1. Request other node's state
2. Compare it with it's own state
3. Update its state when it's outdated

### Node Roles
Despite being full-fledged, blockchain nodes can play different roles in the network
1. **Miner** Such nodes which run on powerful or specialized hardware have only goal is to mine new blocks as fast as possible because mining actually means solving PoW puzzles.
2. **Full node** These node must have the whole copy of blockchain to make a crucial decisions: they decide if a block or transaction is valie. It's very mandatory for network to have many full nodes
3. **Simplified Payment Verification** These nodes don't store a full copy of blockchain, but they still able to verify transactions. As a result, an SPV node depends on a full node to get data from

### Network implementation simplification
We don't have many computers to simulate a network with multiple nodes. But, we implemented a simplified approach where nodes are identified by ports rather than IP address. E.g., there will ne nodes with addresses: 127.0.0.1:3000. 127.0.0.1:3001, 127.0.0.1:3002, etc. We will call the port NODE_ID and you can open multiple terminal windows, set different NODE_IDs and have different nodes running.

This approach also requires having different blockchains and wallet files, for instance: blockchain_3000.db, blockchain_3001.db and wallet_3000.db, wallet_3001.db, etc.

We will have three nodes:
1. The central node: This is the node all other nodes will connect to, and also sends data between other nodes
2. A miner node: This node will store new transactions in mempool and when there're enough of transactions, it'll mine a new block
3. A wallet node: This node is used to send coins between wallets. Unlike SPV nodes, it'll store a full copy of blockchain


### The Scenario
1. The __central node__ creates a blockchain
2. Other (wallet) node connects to it and __downloads the blockchain__
3. One more (miner) node connects to the central node and downloads the blockchain
4. The wallet node creates a transaction
5. The __miner nodes__ receives the transaction and keeps it in its memory pool
6. When there are enough transactions in the memory pool, the miner starts mining a new block
7. __When a new block is mined, it's send to the central node__
8. The __wallet node synchronizes with the central node__
9. User of the wallet node checks that their payment was successful

### Specific Items
#### Version
```golang
type version struct {
    Version int
    BestHeight int
    AddrFrom string
}
```
The notable field BestHeight stores the length of the node's blockchain. When a node receives a version message, it checks if the node's blockchain is longer than the value of BestHeight. If it's not, the node will request and download missing blocks
```golang
var nodeAddress string
var knownNodes = []string{"localhost:3000"} //the address of the central node

func StartServer(nodeID, minerAddress string) {
    ...
    ln, err := net.Listen(protocol, nodeAddress)
    defer ln.Close()

    /*
        if current node is not the central one, it must send version message
        to the central node to find out if its blockchain is outdated
    */
    if nodeAddress != knownNodes[0] {
        sendVersion(knownNodes[0], bc)
    }

    for {
        conn, err := ln.Accept()
        go handleConnection(conn, bc)
    }
}
func handleVersion(request []byte, bc *Blockchain) {
	myBestHeight := bc.GetBestHeight()
	foreignerBestHeight := payload.BestHeight
	/*
		A node compares its BestHeight with the one from the message.
		If the node's blockchain is longer, it'll reply with version message;
		otherwise, it'll send getBlocks message
	*/
	if myBestHeight < foreignerBestHeight {
		sendGetBlocks(payload.AddrFrom)
	} else if myBestHeight > foreignerBestHeight {
		sendVersion(payload.AddrFrom, bc)
	}
}
```
Pay attention, it doesn't say "give me all your blocks", instead it requests a list of block hashes within an Inv struct. This reduces network load and, more importantly, allows blocks to be downloaded from different nodes simultaneously, advoiding the bottleneck of downloading dozens of gigabytes from a single source
```golang
func handleGetBlocks(request []byte, bc *Blockchain) {
    ...
    blocks := bc.GetBlockHashes()
    sendInv(payload.AddrFrom, "block", blocks)
}
```
#### inv
```golang
type inv struct {
    AddrFrom string
    Type string
    Items [][]byte
}
```
Inv just contains just their hashes which the Type field says whether these are blocks or transactions. The Inv struct lists blocks and transaction that node must fetch to align its data with peers
```golang
func handleInv(request []byte) {
	...
	if payload.Type == "block" {
		blocksInTransit = payload.Items

		blockHash := payload.Items[0]
		sendGetData(payload.AddrFrom, "block", blockHash)

		...
	}

	if payload.Type == "tx" {
		txID := payload.Items[0]

		if mempool[hex.EncodeToString(txID)].ID == nil {
			sendGetData(payload.AddrFrom, "tx", txID)
		}
	}
}
```
#### getdata
```golang
type getdata struct {
    AddrFrom string
    Type string
    ID []byte
}
```
The Inv struct presents a request to download a block or transaction identified by a specific hash
```golang
func handleGetData(request []byte, bc *Blockchain) {
	var buff bytes.Buffer
	var payload getdata

	...
    if payload.Type == "block" {
		block, err := bc.GetBlock([]byte(payload.ID))
		if err != nil {
			return
		}

		sendBlock(payload.AddrFrom, &block)
	}

	if payload.Type == "tx" {
		txID := hex.EncodeToString(payload.ID)
		tx := mempool[txID]

		sendTx(payload.AddrFrom, &tx)
		// delete(mempool, txID)
	}
}
```
#### block and tx
```golang
type block struct {
    AddrFrom string
    Block []byte
}

type tx struct {
    AddrFrom string
    Transaction []byte
}
```
Upon receiving a new block represented via block struct, we put it into node blockchain. Once all necessary blocks are downloaded, the UTXO set is reindexed
```golang
func handleBlock(request []byte, bc *Blockchain) {
	var buff bytes.Buffer
	var payload block
    ...
	blockData := payload.Block
	block := DeserializeBlock(blockData)

	fmt.Println("Received a new block!")
	bc.AddBlock(block)
    ...
    if len(blocksInTransit) > 0 {
		
	} else {
		UTXOSet := UTXOSet{Blockchain: bc}
		UTXOSet.Reindex()
	}
}
```
Handling tx messages is more difficult part
1. If currentNode is a central node, it'll forward the new transaction to other node specifically to miner node in the network. The central node must identify the miner node as a prerequisite, which occurs when two node synchronize
```golang
if nodeAddress == knownNodes[0] {
    for _, node := range knownNodes {
        if node != nodeAddress && node != payload.AddFrom {
            fmt.Printf("\nSend to miner node %s intentionally\n", node)
			sendTx(node, &tx)
        }
    }
}
```
2. If currentNode is a miner node, the mining jobs just begin when there are 2 or more transactions in the mempool.
    2.1. Transactions must be verified before being placed into the mempool
    2.2. A coinbase transaction with the reward is also put into a transaction
    2.3. Mining takes place
    2.4. After mining the block, the UTXO set is reindexed
    2.5. The mempool has been flushed
3. The miner node broadcasts the new block hash to all nodes in its peer list
### Extra Information
1. Node discover each other
```golang
knownNodes = append(knownNodes, payload.AddrList...)
```
2. Directly comparing node heights is not a precise method to determine if two node databases are identical; instead, use node content hashes or deep comparisons of the underlying data to identify discrepancies between two database nodes.
3. In the absence of a discovery service, nodes requires a minimum one sync to recognize each other
4. Once a process opens a connection to the Bolt memory database, ant other process attempting to access same database will be blocked
5. Serializing the same data can generate different values, particulary in JSON encoding, due to non-deterministic map iterator order, dynamic pointer addresses. which change across program execution

To ensure deterministic output: 
- Sort Map Keys: before marshaling, iterator through map keys in a sorted order
- Use Structs: Prioritize using struct over map[string]interface{}
- Use json.Marshal on structured data
6. Comparing database height is not a precise method, so we are assume they share same genesis block to initialize synchornization
