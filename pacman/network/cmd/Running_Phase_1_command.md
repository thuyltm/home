```sh
% btcd --configfile ./pacman/network/btcd.conf
# 2026-04-22 10:43:07.829 [INF] CMGR: Server listening on 127.0.0.1:9333
bazel run //pacman/network/cmd
# INFO[0000] received: 161c141276657273696f6e000000000071000000646ceed4801101004d04000000000000f143e869000000004d0400000000000000000000000000000000ffff7f000001e8324d0400000000000000000000000000000000000000000000000074c020ae69d3667b1b2f627463776972653a302e352e302f627463643a302e32352e302f0000000001 
# INFO[0000] received: 161c141276657261636b000000000000000000005df6e0e2 
```
We received two messages from the btcd node
1. version: contains information about the other node's version
2. verack: the Acknowledged message

The Btcd console log is at this time
```sh
# 2026-04-22 10:43:45.663 [DBG] PEER: Received version (agent /Satoshi:5.64/tinybit:0.0.1/, pver 70015, block -1) from 127.0.0.1:59442 (inbound)
# 2026-4-22 10:43:45.663 [DBG] PEER: Negotiated protocol version 70015 for peer 127.0.0.1:59442 (inbound)
# 2026-04-22 10:43:45.663 [DBG] CHAN: Added time sample of 0s (total: 1)
# 2026-04-22 10:43:45.663 [DBG] PEER: Sending version (agent /btcwire:0.5.0/btcd:0.25.0/, pver 70016, block 0) to 127.0.0.1:59442 (inbound)
# 2026-04-22 10:43:45.663 [DBG] PEER: Sending verack to 127.0.0.1:59442 (inbound)
# 2026-04-22 10:44:15.671 [DBG] PEER: Cannot start peer 127.0.0.1:59442 (inbound): protocol negotiation timeout
```
The btcd node correctly decoded our message and replied to it. Version and verack messages as a part of version handshake procedure