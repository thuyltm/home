### The UTXO Set
Bitcoin Cores stores blocks in a blocks database and transaction outputs in achainstate database.

The structure of chainstate is:
1. 'c' + transaction hash -> unspent transaction output record for that transaction
2. 'B' + block hash -> up to which the database represents the unspent transaction outputs

#### Why do we want to have the UTXO set?
Since transactions are stored in blocks, the function __finds unspent transaction outputs that iterators over each block in the blockchain__ and checks every transaction in it. If there're 485860 blocks in Bitcoin and the whole database takes 140+ Gb of disk space. This means that one iterate over many blocks.

The solution to the problem is to have __an index that stores only unspent outputs__.Iterate over blocks once to construct a cache UTXO set and is later used to calculate balnace and validate new transaction

We have some methods used to find transactions:
1. Blockchain.FindUTXO - query all unspent outputs by iterating over blocks
2. UTXOSet.Reindex - identifies FindUTO to find unspent outputs and then cache the results
3. UTXOSet.FindSpendableOutputs - When creating a new transaction, this function __queries the UTXO set to retrieves the enough number of outputs__ holding required amount
4. UTXOSet.FindUTXO - __look up unspent outputs__ for a public key hash using __the indexed UTXO set__
5. Blockchain.FindTransaction - iterates over all blocks to locate a transaction in the blockchain by ID

To summary, actual transactions are logged in the "blockchain.db" bucket, and unspent outputs are recorded in the "utxoBucket" bucket

### P2PKH
In Bitcoin there is the Script programming language, which is used to lock transaction outputs; a trasnaction inputs provide data to unlock outputs.

Look at the script that is used in Bitcoin to perform payments:
```sh
<signature> <pubKey> OP_DUP OP_HASH160 <pubKeyHash> OP_EQUALVERIFY OP_CHECKSIG
```
This script is called _Pay to Public Key Hash_(P2PKH). This is the heart of Bitcoin payments: there's just a script that checks that provided signature and publick key are correct
