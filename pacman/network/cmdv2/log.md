**Btcd Server Console Log**
```sh
2026-04-23 11:08:12.016 [INF] CMGR: Server listening on 127.0.0.1:9333
2026-04-23 11:23:03.161 [DBG] PEER: Received version (agent /Satoshi:5.64/tinybit:0.0.1, pver 70015, block -1) from 127.0.0.1:39818 (inbound)
2026-04-23 11:23:03.161 [DBG] PEER: Negotiated protocol version 70015 for peer 127.0.0.1:39818 (inbound)
2026-04-23 11:23:03.162 [DBG] CHAN: Added time sample of 0s (total: 1)
2026-04-23 11:23:03.162 [DBG] PEER: Sending version (agent /btcwire:0.5.0/btcd:0.25.0/, pver 70016, block 0) to 127.0.0.1:39818 (inbound)
2026-04-23 11:23:03.162 [DBG] PEER: Sending verack to 127.0.0.1:39818 (inbound)
2026-04-23 11:23:03.162 [DBG] PEER: Received verack from 127.0.0.1:39818 (inbound)
2026-04-23 11:23:03.162 [DBG] PEER: Connected to 127.0.0.1:39818
2026-04-23 11:23:03.162 [DBG] SRVR: New peer 127.0.0.1:39818 (inbound)
2026-04-23 11:23:03.162 [INF] SYNC: New valid peer 127.0.0.1:39818 (inbound) (/Satoshi:5.64/tinybit:0.0.1)
2026-04-23 11:23:03.162 [WRN] SYNC: No sync peer candidates available
^C2026-04-23 11:24:31.201 [INF] BTCD: Received signal (interrupt).  Shutting down...
2026-04-23 11:24:31.201 [INF] BTCD: Gracefully shutting down the server...
2026-04-23 11:24:31.201 [WRN] SRVR: Server shutting down
2026-04-23 11:24:31.216 [INF] SYNC: Sync manager shutting down
2026-04-23 11:24:31.216 [INF] SYNC: Lost peer 127.0.0.1:39818 (inbound)
2026-04-23 11:24:31.216 [DBG] SYNC: Block handler shutting down: flushing blockchain caches...
2026-04-23 11:24:31.216 [INF] CHAN: Flushing UTXO cache of 97 MiB with 0 entries to disk. For large sizes, this can take up to several minutes...
2026-04-23 11:24:31.216 [INF] AMGR: Address manager shutting down
2026-04-23 11:24:31.217 [INF] SRVR: Server shutdown complete
2026-04-23 11:24:31.217 [INF] BTCD: Gracefully shutting down the database...
2026-04-23 11:24:31.219 [INF] BTCD: Shutdown complete
```
**Tinybit Console Log**
```sh
DEBU[0000] received header: 161c141276657273696f6e00000000007100000033f6e8d9 
DEBU[0000] received message: version               
INFO[0000] VERSION: /btcwire:0.5.0/btcd:0.25.0/         
DEBU[0000] received header: 161c141276657261636b000000000000000000005df6e0e2 
DEBU[0000] received message: verack 
```