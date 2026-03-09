#### What is blockchain
A blockchain is __a peer-to-peer network__, meaning it is a system of nodes, or computers, that all talk to one another

__Each node on the network maintains a copy of the data on the blockchain__. For example, they all know how much money Alice has in her account. Instead of a bank storing all this data inside a central database, the blockchain stores it redundantly on each node in the network. This makes it virtually impossible to tamper with the data.

Alice needs a Bitcoin wallet in order to send the transaction. She signs the transaction  with her private key in order to make it official.

A select group of nodes called "miners" process her transaction, and record it on the blockchain. __Bitcoin miners must solve a mathematical puzzle called a "proof-of-work" problem in order to record Alice's transaction on the blockchain__. Each miner competes with one another to guess __a random encrypted number called a "nonce"__. The miner that guesses this number first submits their answer, or "proof of work". This authorizes them to record the transaction to the blockchain and receive a Bitcoin reward for winning the mining competition. Whenever they record it, they create a new transaction record on the blockchain. Groupings of transactions are called __"blocks" which get "chained together" to make up the blockchain__.

#### Blockchain Block
A blockchain is a series of blocks linked by cryptographic hashes. Each block contains:
- Data: Information stored in the block
- Hash: A SHA-256 hash of the block's content
- Previous Hash: The hash of the previous block, linking blocks together
- Nonce: A counter used in mining

#### Consensus Algorithm

During the mining process, the network reaches "consensus", meaning that each node verifies that they have the same valid data as everyone else. If they all agree, then the transaction is complete. 
1. Proof of Work (PoW)
2. Proof of Stake (PoS)
3. Proof of Authority (PoA)
4. Raft and IBFT/QBFT

#### Smart contracts
All of the backend code for the applicaton will be made up of smart contracts. These are immutable building blocks of blockchain applications. Once the codes is put on the blockchain, no one can change it, and we know it will work the same way every time. 

Building a blockchain from scratch in Go involves fundamental concepts like block structure, hashing, proof-of-work, and chaining blocks.

#### Step by step guide to building a Blockchain
1. Define block structure, calculate SHA-256 hashes, create the genenis block
2. Proof-Of-Work and Mining
3. Distributed storage for blocks (Wallet and Transactions)
4. Peer-to-peer network

#### Local Environments (Fastest & Free)
- Hardhat: Built-in network for rapid development and testing in JS/TS
- Ganache (Truffle Suite): Personal, local blockchain for quick simulation
- Geth/Parity: Used to create private chains for specialized, isolated tests

#### Tutorial
https://github.com/hlongvu/blockchain-go-vietnamese?tab=readme-ov-file

https://jeiwan.net/posts/building-blockchain-in-go-part-1/

https://github.com/volodymyrprokopyuk/go-blockchain

https://blog.logrocket.com/build-blockchain-with-go/

https://goethereumbook.org/en/