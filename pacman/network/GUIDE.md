https://jeiwan.net/posts/programming-bitcoin-network/

A tiny Bitcoin network client is able to
1. Connect to a Bitcoin network (whether that's mainnet, testnet, or a local network)
2. Introduce itself to the netwrok (what's called "version handshake")
3. Get information about current blockchain state from a node in the network
4. Download full blockchain history
5. Verify and relay new transactions
6. Verify and relay new blocks

A Project Layout
|-- btcd
    |-- blockchain_data
|-- btcd.conf
|-- cmd
    |-- tinybit.go
|-- go.mod
|-- go.sum
|-- main.go
|-- protocol
    |-- command.go
    |-- message.go
    |-- message_version.go
    |-- netaddr.go
    |-- protocol.go

# Install btcd
https://github.com/btcsuite/btcd