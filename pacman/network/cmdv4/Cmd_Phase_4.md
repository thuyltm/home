### Prepare data
```sh
btcd --configfile ./btcd.conf
btcwallet -C ./btcwallet.conf
btcctl -C ./btcctl-wallet.conf walletpassphrase "luxury hurry what trick slim easy congress ceiling analyst trick palace help" 3600
btcctl -C ./btcctl-wallet.conf getnewaddress
#MINER_ADDRESS
btcd --configfile ./btcd.conf --miningaddr=MINER_ADDRESS
#deploying 100 mining units to earn some BTC coins for MINER_ADDRESS
btcctl -C ./btcctl.conf generate 100
btcctl -C ./btcctl-wallet.conf getbalance
btcctl -C ./btcctl-wallet.conf getnewaddress alice
#ALICE_ADDRESS
```

### Start Server
1. Start Btcd Server
```sh
btcd --configfile ./btcd.conf --miningaddr=MINER_ADDRESS
```
2. Start Wallet Server
```sh
btcwallet -C ./btcwallet.conf
```
3. Start Tinybit Server
```sh
bazel run //pacman/network/cmdv4
```
4. Start a transaction
```sh
btcctl -C ./btcctl-wallet.conf sendtoaddress ALICE_ADDRESS 0.00001
```
Check mempool
```sh
[DBG] TXMP: Accepted transaction 66f3c40cb9a6cc7e31a228eee8752ea4e9fbd9cd3e411279ba61151fc45cc73d (pool size: 1)
```