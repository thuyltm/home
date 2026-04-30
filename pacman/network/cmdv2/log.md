**Btcd Server Console Log**
```sh
[INF] CMGR: Server listening on 127.0.0.1:9333
[DBG] PEER: Received version (agent /Satoshi:5.64/tinybit:0.0.1, pver 70015, block -1) from 127.0.0.1:39818 (inbound)
[DBG] PEER: Negotiated protocol version 70015 for peer 127.0.0.1:39818 (inbound)
[DBG] CHAN: Added time sample of 0s (total: 1)
[DBG] PEER: Sending version (agent /btcwire:0.5.0/btcd:0.25.0/, pver 70016, block 0) to 127.0.0.1:39818 (inbound)
[DBG] PEER: Sending verack to 127.0.0.1:39818 (inbound)
[DBG] PEER: Received verack from 127.0.0.1:39818 (inbound)
[DBG] PEER: Connected to 127.0.0.1:39818
[DBG] SRVR: New peer 127.0.0.1:39818 (inbound)
[INF] SYNC: New valid peer 127.0.0.1:39818 (inbound) (/Satoshi:5.64/tinybit:0.0.1)
[WRN] SYNC: No sync peer candidates available
[INF] BTCD: Received signal (interrupt).  Shutting down...
[INF] BTCD: Gracefully shutting down the server...
[WRN] SRVR: Server shutting down
[INF] SYNC: Sync manager shutting down
[INF] SYNC: Lost peer 127.0.0.1:39818 (inbound)
[DBG] SYNC: Block handler shutting down: flushing blockchain caches...
[INF] CHAN: Flushing UTXO cache of 97 MiB with 0 entries to disk. For large sizes, this can take up to several minutes...
[INF] AMGR: Address manager shutting down
[INF] SRVR: Server shutdown complete
[INF] BTCD: Gracefully shutting down the database...
[INF] BTCD: Shutdown complete
```
**Tinybit Console Log**
```sh
DEBU[0000] received header: 161c141276657273696f6e00000000007100000033f6e8d9 
DEBU[0000] received message: version               
INFO[0000] VERSION: /btcwire:0.5.0/btcd:0.25.0/         
DEBU[0000] received header: 161c141276657261636b000000000000000000005df6e0e2 
DEBU[0000] received message: verack 
```