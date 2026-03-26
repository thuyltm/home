# pamac
Build the basic Blockchain using pure Golang
1. Transaction

A Transaction is a cryptographically signed, structured data entry representing a transfer of value or a smart contract operation

2. Block

A block is a digital container that stores a batch of validate transactions, a timestamp, and a cryptographic hash of the previous block.

3. Blockchain

A blockchain use cryptographic hashes to link data blocks in a secure, chronological chain

4. Pow

Proof of Work is a decentralized consensus mechanism in blockchian that solive complex mathematical puzzles to validate transactions and create new blocks.

5. Wallet

A wallet stores a private key to sign and authorize transactions and a public key for generating an address for receiving funds

6. Merkle Tree

A Merkle Tree is a hierarchical data strucure to efficiently verify large sets of data, such as transactions within a block. The root acts as a mathematical summary of every transaction in that block


**Key Characteristics:**
- Decentralized: The entire ledger - a specialized record book for financial transactions - is shared across a network of nodes. Every node has a copy of the blockchain
- Immutable: data cannot be deleted or changed once added
- Hashing: each block contains a unique cryptographic hash of its own data, as well as the hash of the previous hash
- Chain: these blocks are chronologically linked together, forming a chain
- Consensus: Before a new block is added, a consensus mechanism (e.g. Proof of Work or Proof of Stake) must be executed.
- Secure: Cryptographic techniques make it highly resistant to hacking and fraud