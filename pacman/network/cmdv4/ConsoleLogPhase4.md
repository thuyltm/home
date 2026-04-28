1. TinyBit console log
```sh
INFO[0000] Running JSON-RPC server on port 9334         
DEBU[0000] received message: version               
INFO[0000] VERSION: /btcwire:0.5.0/btcd:0.25.0/         
DEBU[0000] received message: verack               
DEBU[0120] received message: ping               
DEBU[0230] received message: inv               
DEBU[0230] received message: tx               
DEBU[0230] transaction: {Version:1 Flag:0 TxInCount:1 TxIn:[{PreviousOutput:{Hash:[5 110 30 120 56 58 93 228 208 60 139 181 177 66 125 251 104 151 166 146 195 18 250 137 248 137 6 20 73 103 224 245] Index:0} ScriptLength:106 SignatureScript:[71 48 68 2 32 20 238 197 82 209 55 48 124 233 34 103 90 123 64 248 239 73 30 201 196 132 16 247 147 226 9 138 51 196 131 130 14 2 32 33 31 180 95 244 165 7 186 208 76 120 64 240 18 246 167 228 57 24 49 172 55 157 78 106 252 155 188 90 34 107 78 1 33 2 176 233 65 89 133 114 196 6 109 194 216 64 201 223 88 54 231 191 109 20 25 152 239 97 143 77 217 74 2 203 153 158] Sequence:4294967295}] TxOutCount:2 TxOut:[{Value:1000 PkScriptLength:25 PkScript:[118 169 20 173 126 171 110 144 8 119 45 252 91 41 248 62 253 198 31 42 176 158 240 136 172]} {Value:4999998773 PkScriptLength:25 PkScript:[118 169 20 162 14 182 246 52 250 25 241 42 190 170 59 33 200 139 31 184 100 155 213 136 172]}] TxWitness:{Count:0 Witness:[]} LockTime:0} 
DEBU[0240] received message: inv               
DEBU[0240] received message: ping               
DEBU[0240] received message: tx               
DEBU[0240] transaction: {Version:1 Flag:0 TxInCount:1 TxIn:[{PreviousOutput:{Hash:[5 110 30 120 56 58 93 228 208 60 139 181 177 66 125 251 104 151 166 146 195 18 250 137 248 137 6 20 73 103 224 245] Index:0} ScriptLength:106 SignatureScript:[71 48 68 2 32 20 238 197 82 209 55 48 124 233 34 103 90 123 64 248 239 73 30 201 196 132 16 247 147 226 9 138 51 196 131 130 14 2 32 33 31 180 95 244 165 7 186 208 76 120 64 240 18 246 167 228 57 24 49 172 55 157 78 106 252 155 188 90 34 107 78 1 33 2 176 233 65 89 133 114 196 6 109 194 216 64 201 223 88 54 231 191 109 20 25 152 239 97 143 77 217 74 2 203 153 158] Sequence:4294967295}] TxOutCount:2 TxOut:[{Value:1000 PkScriptLength:25 PkScript:[118 169 20 173 126 171 110 144 8 119 45 252 91 41 248 62 253 198 31 42 176 158 240 136 172]} {Value:4999998773 PkScriptLength:25 PkScript:[118 169 20 162 14 182 246 52 250 25 241 42 190 170 59 33 200 139 31 184 100 155 213 136 172]}] TxWitness:{Count:0 Witness:[]} LockTime:0}
```
2. Btcd Console Log
```sh
[DBG] CHAN: Added time sample of 0s (total: 1)
[DBG] PEER: Sending version (agent /btcwire:0.5.0/btcd:0.25.0/, pver 70016, block 170) to 127.0.0.1:47566 (inbound)
[DBG] PEER: Sending verack to 127.0.0.1:47566 (inbound)
[DBG] PEER: Received verack from 127.0.0.1:47566 (inbound)
[DBG] PEER: Connected to 127.0.0.1:47566
[DBG] SRVR: New peer 127.0.0.1:47566 (inbound)
[INF] SYNC: New valid peer 127.0.0.1:47566 (inbound) (/Satoshi:5.64/tinybit:0.0.1/)
[INF] SYNC: Caught up to block 4bddf2382734cf4b6ecbfc054d63e080436ee940093b1d2453b38a34c75eb639(170)
[DBG] PEER: Sending ping to 127.0.0.1:47566 (inbound)
[DBG] PEER: Received pong from 127.0.0.1:47566 (inbound)
[DBG] RPCS: Received command <notifyreceived> from 127.0.0.1:58370
[DBG] RPCS: Received command <notifyreceived> from 127.0.0.1:58370
[DBG] RPCS: Received command <notifyreceived> from 127.0.0.1:58370
[DBG] RPCS: Received command <getinfo> from 127.0.0.1:58370
[DBG] RPCS: Received command <sendrawtransaction> from 127.0.0.1:58370
[DBG] TXMP: Accepted transaction 66f3c40cb9a6cc7e31a228eee8752ea4e9fbd9cd3e411279ba61151fc45cc73d (pool size: 1)
[DBG] PEER: Sending inv (tx 66f3c40cb9a6cc7e31a228eee8752ea4e9fbd9cd3e411279ba61151fc45cc73d) to 127.0.0.1:47566 (inbound)
[DBG] PEER: Received getdata (tx 66f3c40cb9a6cc7e31a228eee8752ea4e9fbd9cd3e411279ba61151fc45cc73d) from 127.0.0.1:47566 (inbound)
[DBG] PEER: Sending tx (hash 66f3c40cb9a6cc7e31a228eee8752ea4e9fbd9cd3e411279ba61151fc45cc73d, 1 inputs, 2 outputs, lock height 0) to 127.0.0.1:47566 (inbound)
[DBG] PEER: Sending inv (tx 66f3c40cb9a6cc7e31a228eee8752ea4e9fbd9cd3e411279ba61151fc45cc73d) to 127.0.0.1:47566 (inbound)
[DBG] PEER: Sending ping to 127.0.0.1:47566 (inbound)
[DBG] PEER: Received getdata (tx 66f3c40cb9a6cc7e31a228eee8752ea4e9fbd9cd3e411279ba61151fc45cc73d) from 127.0.0.1:47566 (inbound)
[DBG] PEER: Sending tx (hash 66f3c40cb9a6cc7e31a228eee8752ea4e9fbd9cd3e411279ba61151fc45cc73d, 1 inputs, 2 outputs, lock height 0) to 127.0.0.1:47566 (inbound)
[DBG] PEER: Received pong from 127.0.0.1:47566 (inbound)
[DBG] PEER: Sending ping to 127.0.0.1:47566 (inbound)
[DBG] PEER: Received pong from 127.0.0.1:47566 (inbound)
```
3. BtcWallet console log
```sh
[INF] BTWL: The wallet has been temporarily unlocked
[INF] TMGR: Inserting unconfirmed transaction 66f3c40cb9a6cc7e31a228eee8752ea4e9fbd9cd3e411279ba61151fc45cc73d
[INF] RPCS: Successfully sent transaction 66f3c40cb9a6cc7e31a228eee8752ea4e9fbd9cd3e411279ba61151fc45cc73d
```