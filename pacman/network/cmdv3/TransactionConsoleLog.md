### Command
```sh
% btcctl -C ./btcctl-wallet.conf createnewaccount alice
% btcctl -C ./btcctl-wallet.conf listaccounts
% btcctl -C ./btcctl-wallet.conf getnewaddress
# MINING_ADDRESS
% btcctl -C ./btcctl-wallet.conf getnewaddress alice
# ALICE_ADDRESS
% btcctl -C ./btcctl-wallet.conf getbalance
% btcctl -C ./btcctl-wallet.conf walletpassphrase "luxury hurry what trick slim easy congress ceiling analyst trick palace help" 3600
% btcctl -C ./btcctl-wallet.conf sendtoaddress SN4B1ysAhLopyCdkpqipwysysfWnmwXyKV  0.00001
```

### Btcd Server Console Log
```sh
btcd --configfile ./btcd.conf --miningaddr=ScbrmUSkzbkAELjcD1MQod3mbGqYhb1cBF
# PEER: Received version (agent /Satoshi:5.64/tinybit:0.0.1/, pver 70015, block -1) from 127.0.0.1:56632 (inbound)
# PEER: Negotiated protocol version 70015 for peer 127.0.0.1:56632 (inbound)
# CHAN: Added time sample of 0s (total: 1)
# PEER: Sending version (agent /btcwire:0.5.0/btcd:0.25.0/, pver 70016, block 100) to 127.0.0.1:56632 (inbound)
# PEER: Sending verack to 127.0.0.1:56632 (inbound)
# PEER: Received verack from 127.0.0.1:56632 (inbound)
# PEER: Connected to 127.0.0.1:56632
# SRVR: New peer 127.0.0.1:56632 (inbound)
# SYNC: New valid peer 127.0.0.1:56632 (inbound) (/Satoshi:5.64/tinybit:0.0.1/)
# SYNC: Caught up to block 12568e7708b09ea02eb4e3062e7ea6e36f077627cff0cfaf15d892a5d029cb07(100)
# PEER: Sending ping to 127.0.0.1:56632 (inbound)
# RPCS: Received command <notifyreceived> from 127.0.0.1:56324
# RPCS: Received command <notifyreceived> from 127.0.0.1:56324
# RPCS: Received command <getinfo> from 127.0.0.1:56324
# RPCS: Received command <sendrawtransaction> from 127.0.0.1:56324
# TXMP: Accepted transaction 2b8a650e3e69b3a334e4842b5f2dc71a8757b669917216c4aaa6673df2d5f03b (pool size: 1)
# PEER: Sending inv (tx 2b8a650e3e69b3a334e4842b5f2dc71a8757b669917216c4aaa6673df2d5f03b) to 127.0.0.1:56632 (inbound)
# PEER: Sending ping to 127.0.0.1:56632 (inbound)
# PEER: Peer 127.0.0.1:56632 (inbound) no answer for 5m0s -- disconnecting
# SYNC: Lost peer 127.0.0.1:56632 (inbound)
# SRVR: Removed peer 127.0.0.1:56632 (inbound)
```

### Btcd Wallet Server Console Log
```sh
btcwallet -C ./btcwallet.conf
# BTCW: Version 0.15.1-alpha
# RPCS: Listening on 127.0.0.1:18554
# BTCW: Attempting RPC client connection to localhost:18556
# CHIO: Established connection to RPC server localhost:18556
# BTWL: Opened wallet
# BTWL: RECOVERY MODE ENABLED -- rescanning for used addresses with recovery_window=250
# BTWL: Started rescan from block 683e86bd5c6d110d91b94b97137ba6bfe02dbbdb8e3dff722a669b5d69d77af6 (height 0) for 4 addrs, 0 outpoints
# BTWL: Catching up block hashes to height 0, this might take a while
# BTWL: Done catching up block hashes
# BTWL: Finished rescan for 4 addresses (synced to block 683e86bd5c6d110d91b94b97137ba6bfe02dbbdb8e3dff722a669b5d69d77af6, height 0)
# BTWL: The wallet has been temporarily unlocked
# TMGR: Inserting unconfirmed transaction 2b8a650e3e69b3a334e4842b5f2dc71a8757b669917216c4aaa6673df2d5f03b
# RPCS: Successfully sent transaction 2b8a650e3e69b3a334e4842b5f2dc71a8757b669917216c4aaa6673df2d5f03b
```

### Tinybit Server Console Log
```sh
bazel run //pacman/network/cmdv3
# DEBU[0000] received header: 161c141276657273696f6e0000000000710000006ad35218 
# DEBU[0000] received message: version               
# INFO[0000] VERSION: /btcwire:0.5.0/btcd:0.25.0/         
# DEBU[0000] received header: 161c141276657261636b000000000000000000005df6e0e2 
# DEBU[0000] received message: verack               
# DEBU[0120] received header: 161c141270696e670000000000000000080000001f868a49 
# DEBU[0120] received message: ping               
# DEBU[0120] received header: 20f3183729a41f95            
# ERRO[0120] EOF                                          
# DEBU[0170] received header: 161c1412696e7600000000000000000025000000cd9a2ba3 
# DEBU[0170] received message: inv               
# DEBU[0170] received header: 01010000003bf0d5f23d67a6aac416729169b657871ac72d 
# DEBU[0170] received message: ;���=g���r               
# DEBU[0170] received header: 5f2b84e434a3b3693e0e658a2b  
# ERRO[0170] EOF                                          
# DEBU[0240] received header: 161c141270696e6700000000000000000800000035380bf5 
# DEBU[0240] received message: ping               
# DEBU[0240] received header: 49df6da07e183680            
# ERRO[0240] EOF 
```