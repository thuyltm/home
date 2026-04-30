### Liveliness check
The btcd node sends 'ping' message to TinyBit node and expects an answer. If there's no answer, it removes our node from the list of its peers

Pinging other nodes in a network is a common way of checking their liveliness, i.e if the other node is still running and responding. If a node fails to reply in time and with proper message, it gets removed from the list of peers.

### List of peers
Each nodes in the Bitcoin network should connected to serveral nodes. Nodes continuously exchange ping and pong message to monitor peer health
The node can have a list of peers
```go
type Node struct {
    ...
    Peers map[string]*Peer
    ...
}
```