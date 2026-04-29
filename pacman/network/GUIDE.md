https://jeiwan.net/posts/programming-bitcoin-network/

**A tiny Bitcoin network CLIENT is able to**
1. Connect to a Bitcoin network (whether that's mainnet, testnet, or a local network)
2. Introduce itself to the netwrok (what's called "version handshake")
3. Get information about current blockchain state from a node in the network
4. Download full blockchain history
5. Verify and relay new transactions
6. Verify and relay new blocks

**Full Bitcoin nodes serve multiple purpose**
1. Downloading all blocks and transactions
2. Verifying all downloaded blocks and transactions
3. Let other node get blocks and transations from this node's database
4. Verify and relay unconfirmed transactions

### What is btcd
Btcd is an alternative **full node bitcoin implementation** written in Go

Btcd downloads, validates, and servers the block chain using the exact rules (consensus algorithm) for block acceptance.

It also properly relays newly mined blocks, maintains a transaction pool, and relays individual transactions that have not yet made it into a block

One key difference between btcd and Bitcoin Core is that btcd does NOT include wallet functionality


#### Install btcd
https://github.com/btcsuite/btcd