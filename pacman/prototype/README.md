#### Block
In blockchain it's __blocks that store valuable information__. For example, bitcoin blocks store transactions, the essence of any cryptocurrency. Besides this, a block contains some technical information, like its version, current timestamp and the hash of the previous block and current block

__They way hashes are calculates is very important feature of blockchain__, and it's this feature that makes blockchain secure. This is an intentional architectural design, which makes adding new blocks difficult, thus preventing their modification after they're added.

__Blocks are stored in the insertion order and that each block is linked to the previous one.__ This structure allows to quickly get the latest block in a chain

The actual blockchain is much more complex. Adding new blocks requires some work: one has to perform the machanism called Proof-of-Work before getting a permission to add block. Also, blockchain is distributed database that has no single decision maker. Thus, a new block must be confirmed and approved by other paricipants of the network (this mechanism is called consensus). ANd there's no transactions in our blockchain yet