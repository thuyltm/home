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
bazel run //pacman/network/cmdv4
btcctl -C ./btcctl-wallet.conf getnewaddress alice
#ALICE_ADDRESS
btcctl -C ./btcctl-wallet.conf sendtoaddress Sd7MgBzz4B4BofGkrf5zL631YLw1KXygaN   0.00001
```