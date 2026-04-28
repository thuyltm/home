# Start server Btcd
```sh
btcd --configfile ./btcd.conf
```
# Start server Btcwallet
1. Create a new wallet
```sh
btcwallet -C ./btcwallet.conf --create
# This create a new btcwallet folder containing rpc.cert and the simnet/wallet.db database
```
2. Start Btcd Wallet Server
```sh
btcwallet -C ./btcwallet.conf
```
3. Create an account and bind wallet address to an account
```sh
# unlock the wallet
btcctl -C ./btcctl-wallet.conf walletpassphrase "luxury hurry what trick slim easy congress ceiling analyst trick palace help" 3600
btcctl -C ./btcctl-wallet.conf createnewaccount alice
btcctl -C ./btcctl-wallet.conf listaccounts
btcctl -C ./btcctl-wallet.conf getnewaddress
btcctl -C ./btcctl-wallet.conf getnewaddress alice
```
4. Miners are rewarded BTC when they generate a new block
```sh
btcd --configgile ./btcd.conf --miningaddr=MINING_ADDRESS
btcctl -C ./btcctl.conf generate 100
btcctl -C ./btcctl-wallet.conf getbalance
```

**Extra**: Adding a custome command to your Linux command line is done by adding your directory to your Path
```sh
export GOPATH=/home/thuy/go
export PATH=$PATH:$GOPATH/bin
```