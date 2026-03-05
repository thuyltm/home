### Proof of Work
A key idea of blockchain is that one has to __perform some hard work to put data in it__. It is this hard work that makes blockchain secure and consistent.

In blockchain, some participants of the network work to sustain the network, to add new blocks of it, and __get a reward for their work__.

This whole "__do hard work and prove__" mechanism is called proof-of-work. The difficulty of this work increases from time to time to keep new blocks rate at about 6 blocks per hour. In Bitcoin, the goal of such work is to find a hash for a block, that meets some requirements.

Proof-of-Work algorithms must meet a requirement: __doing the work is hard, but verifying the proof is easy__.

### Hashing
Hashing is a process of obtaining a hash for specified data. A hash is a unique representation of the data it was calcualted on. A hash function is a function that takes data of arbitrary size and produces a fixed size hash.

Here are some features of hashing:
1. __Original data cannot be restored from a hash__. Thus, __hashing is not encryption__
2. Certain data can have only one hash and __the hash is unique__
3. Changing even one byte in the input data will result in a completely different hash

In blockchain, hashing is used to guarantee the consistency of a block.

Name convention of algorithm-hash: HMAC-SHA1, HMAC-MD5, HMAC-SHA256, RSA-SHA1, etc

### Hashcash
The idea "...to require a use to compute a moderately hard, but not intractable function..."

For email users, a textual encoding of a hashcash stamp is added to the header of an email to prove the sender has expended a modest amount of CPU time calculating the stamp prior to sending the email. The receiver can, at negligible computational cost, verify that the stamp is valid. However, the only know way to find a header with the necessary properties is brute force, trying random values until the answer is found.

The core differences between Hashcash and other proof-of-work systems lie in purpose. Hashcash was designed to combat email spam by making it computationally expensive to send mass emails. On the other hand, PoW systems in cryptocurrencies serve as a consesus mechanism, an agreement on the state of the network and validitity of transactions.

Hashcash served as a foundatonal concept for the PoW system implemented in Bitcoin mining. In the Bitcoin network, miners compete to solve a mathematical puzzle based on the network's hashing algorithm, SHA-256. It transforms transaction data into a unique string of characters, typically 64 characters long, called a Bitcoin hash. This process requires significant computational effort, effectively serving as proof-of-work. Once a miner finds a valid solution to the puzzle, they broadcast the new block to the network. Other participants then verify the validity of the block and its transactions, accepting it into the blockchain if it meets the consensus rules.

Bitcoin uses Hashcash, a Proof-of-Work algorithm, as the mining core.

Hashcash can be split into the following steps:
1. The sender prepares a header, which contains the recipient's email address, and information proving that the required computation has been performed, and appends a counter value initialized to a random number
2. It then computes a hash of the data + counter combination
3. If the first 20 bits of the hash are all zeros that meets certain requirements, then this is an acceptable header. 
4. If not, then the sender increments the counter and tries the hash again

The number of times that the sender needs to try to get a valid hash value is modeled by geometric distribution. Hence the sender will on average have to try $2^{20}$ values to find a valid header

