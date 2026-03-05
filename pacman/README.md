A blockchain is a peer-to-peer network, meaning it is a system of nodes, or computers, that all talk to one another

Each node on the network maintains a copy of the data on the blockchain. For example, they all know how much money Alice has in her account. Instead of a bank storing all this data inside a central database, the blockchain stores it redundantly on each node in the network. This makes it virtually impossible to tamper with the data.

Alice needs a Bitcoin wallet in order to send the transaction. She signs the transaction  with her private key in order to make it official.

A select group of nodes called "miners" process her transaction, and record it on the blockchain. Bitcoin miners must solve a mathematical puzzle called a "proof-of-work" problem in order to record Alice's transaction on the blockchain. Each miner competes with one another to guess a random encrypted number called a "nonce". The miner that guesses this number first submits their answer, or "proof of work". This authorizes them to record the transaction to the blockchain and receive a Bitcoin reward for winning the mining competition. Whenever they record it, they create a new transaction record on the blockchain. Groupings of transactions are called "blocks" which get "chained together" to make up the blockchain.

**Consensus Algorithm**

During the mining process, the network reaches "consensus", meaning that each node verifies that they have the same valid data as everyone else. If they all agree, then the transaction is complete. 

**Smart contracts**
All of the backend code for the applicaton will be made up of smart contracts. These are immutable building blocks of blockchain applications. Once the codes is put on the blockchain, no one can change it, and we know it will work the same way every time. 

Building a blockchain from scratch in Go involves fundamental concepts like block structure, hashing, proof-of-work, and chaining blocks.


#### Tutorial
https://github.com/hlongvu/blockchain-go-vietnamese?tab=readme-ov-file

https://jeiwan.net/posts/building-blockchain-in-go-part-1/

https://github.com/volodymyrprokopyuk/go-blockchain

https://blog.logrocket.com/build-blockchain-with-go/