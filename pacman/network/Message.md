In the Bitcoin network, nodes communicate with each other by exchanging messages
1. Implementing messages as they're defined in the protocol specification
2. Properly encoding and decoding them
3. Send them over the network

### Message structure
1. Magic is a 4 byte network identifier.
2. Commmand is a 12 byte command name
3. Length is the length of message payload
4. Checksum verifies the integrity of message payload. It's calculated as SHA256(SHA256(payload))
5. Payload is the actual message. It's serialized before being put here

Message serialization and deserialization is very important. Without proper (de)serialization it won't be possible to build communication between nodes.

### Version Message structure
1. Version sepecifies Bitcoin protocol version
2. Services sepecifies features supported by our node
3. Timestamp current timestamp in seconds
4. AddrRecv, AddrFrom contain destination and source network addresses
5. Nonce a random number that allows to distinguish similar messages
6. UserAgent is analogous to the User-Agent HTTP header
7. StartHeight holds the number of the last block our node stores
8. Relay tells whether the connected node shoud send us transactions or not. If setting it to true, we'll want to receive all transactions from other node