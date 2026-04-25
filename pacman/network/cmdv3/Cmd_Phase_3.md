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
3. Create an account
```sh
btcctl -C ./btcctl-wallet.conf createnewaccount alice
btcctl -C ./btcctl-wallet.conf listaccounts
```

```sh
btcctl -C ./btcctl-wallet.conf walletpassphrase "luxury hurry what trick slim easy congress c
eiling analyst trick palace help" 3600
btcctl -C ./btcctl-wallet.conf createnewaccount alice
btcctl -C ./btcctl-wallet.conf listaccounts
```

**Extra**: Adding a custome command to your Linux command line is done by adding your directory to your Path
```sh
export GOPATH=/home/thuy/go
export PATH=$PATH:$GOPATH/bin
```