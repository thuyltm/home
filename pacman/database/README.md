#### BoltDB
BoltDB is a pure Go key/value store that don't require a full database server such as Postgres or MySQL. Key-value pairs are stored in buckets, which are intended to group similar pairs. Thus, in order to get a value, you need to know a bucket and a key.

One important thing about BoltDB is that there are no data types: keys and values are byte arrays. Since we'll store Go structs in it, we'll need to serialize them,i.e. implement a mechanism of converting a Go struct into a byte array and restoring it back from a byte array

Bitcoin Core uses two "buckets" to store data:
1. blocks stores metadata describing all the blocks in a chain
2. chainstate stores the state of a chain

#### Persistence
1. Open a DB file
2. Check if there's a blockchain stored in it
3. If there's a blockchain:
    3.1. Set the tip to the last block hash stored in the DB
4. If there's no existing blockchain:
    4.1. Create the genesis block
    4.2 Store in the DB
    4.3 Save the genesis block's hash as the last block hash
    4.4 Set the top to the genesis block hash
5. Create a new Blockchain instance with the tip specific to its case