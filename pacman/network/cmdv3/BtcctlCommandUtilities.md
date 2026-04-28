1. Generates a set number of blocks
```sh
btcctl -C ./btcctl-wallet.conf help generate
```
2. Creates a new account. The wallet must be unlocked for this request to succeed
```sh
btcctl -C ./btcctl-wallet.conf help createnewaccount
```
3. Returns a JSON object of all accounts and their balances
```sh
btcctl -C ./btcctl-wallet.conf help listaccounts
```
4. Unlock wallet passphase
```sh
btcctl -C ./btcctl-wallet.conf help walletpassphrase
```
5. Generates and returns a new payment address
```sh
btcctl -C ./btcctl-wallet.conf help getnewaddress
```