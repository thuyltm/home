![MerkleTree](https://github.com/cbergoon/merkletree/blob/master/merkle_tree.png?raw=true)
(https://github.com/cbergoon/merkletree/blob/master/merkle_tree.png?raw=true)

A Merkle Tree is a hash tree that provides an efficient way to verify the contents of a set data

Each of these items is inserted into a leaf node and a tree of hashes is __constructed bottom up using a hash of left and right children__. The root node will be a hash of all other nodes in the tree. This property allows the tree verified by on the hash of the root node of the tree. 

__The benefit of the tree structure is verifying any single content entry in the tree will require only $nlog2(n)$ steps in the worst case__
 
![MerkleTree in Blockchain](https://jeiwan.net/images/merkle-tree-diagram.png)

Because of the decentralized nature of Bitcoin, every node in the network must be independent and self-sufficient, i.e. every node must store a full copy of the blockchain. Also, there's certain internet traffic required to interact with other nodes and download new blocks. Nodes are full-fledged participants of the network, they have responsibilities: they myst verify transactions and blocks.

This rule becomes more difficult to follow: it's not likely that everyone will run a full node. There was a solution for this problem: Simplified Payment Verification. SPV is a light Bitcoin node that doesn't download the whole blockchain and doesn't verify blocks and transactions. Instead, it finds transactions in blocks (to verify payments) and is linked to a full node to retrieve just necessary data. This mechanism allows having multiple light wallet nodes with running just one full node.

For SPV to be possible, Merkle tree comes into play to check if a block contains certain transaction without downloading the whole block

A Merkle tree is built for each block, and it starts with leaves (the bottom of the tree), where a leaf is a transaction hash. 

Moving from the bottom up, leaves are grouped in pairs, their hashes are concatenated, and __a new hash is obtained from the concatenated hashes__. The new hashes form a new tree nodes. This process is repeated until there's just one node, which is called the root of the tree. The root hash is then used as the unique representation of the transactions, is saved in block headers, and is used in the proof-of-work.


### Run
```sh
bazel test //pacman/merkletree:merkletree_test --test_output=all
```