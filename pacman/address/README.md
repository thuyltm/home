### Bitcoin address
There are no user accounts, your personal data is not required and not stored anywhere in Bitcoin. Bitcoin address is something that identifies you as the owner of transaction outputs. If you want to send coins to someone, you need to know their address. But addresses are not something that identifies you as the owner of a wallet. In fact, such addresses are a representation of public keys. 

__When you use a Bitcoin client to genrate a new address, a pair of keys is generated for you__. Public keys are not sensitive and can be disclosed to anyone. No one but the owner should have access to private keys because private key servve as the identifier of the owner. The one who controls the private key controls all the coins sent to this key in Bitcoin. 

### Digitatl Signatures
There's a concept of digital signature algorithms that guarantee:
1. data wasn't modified while being transferred from a sender to a recipient
2. the sender cannot deny sending the data
Digital signing happens with the usage of a private key, and verification requires a public key.

Every transaction input is signed by the one who created the transaction. Every transaction in Bitcoin must be verified before being put in a block.

Let's now review the full lifecycle of a transaction:
1. In the beginning, there's the genesis block that contains a coinbase transaction. There are no real inputs in coinbase transactions, so signing is not necessary.
2. When one sends coins, a transaction is created. This transaction will be linked to the previous one. Every input will store a public key (not hashed) and a signature of the whole transaction.
3. Node receiving the transaction will verify it. It will check public key and the signature to ensure that the transaction is created by the real owner of the coins
4. Upon validation, it'll put the transaction in a block and start mining it
5. When the block is mined, every other node in the network receives a message saying the block is mined and adds the block to the blockchain
6. Completed transaction outputs are referenced in future transactions

### Generate Private key
Bitcoin uses elliptic curves to generate private keys. Elliptic curves can be used to generate really big and random numbers. The curve used by Bitcoin can randomly pick a number between 0 and $2^{256}$, which is approximately $10^{77}$. Such a huge upper limit means that it's almost impossible to generate the same private key twice

### Base58
Bitcoin uses the Base58 algorithm to convert public keys into human readable format. The algorithm is very similar to famous Base64, but it uses shorter alphabet: 

Since hasing is one way, which cannot be reversed, it's not possible to extract the public key from the hash. But we can check it by running it though the save hash functions and comparing the hashes

### Wallet
A wallet manages the private key needed to sign transaction and public key derived from it




