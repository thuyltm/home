Install from the source code
```sh
cd /tmp
git clone https://github.com/btcsuite/btcwallet
cd btcwallet
go install -v . ./cmd/...
```
If you run __go install__ on a package, the resulting binary is placed in:
- The $GOBIN directory, if it is set
- The $GOPATH/bin directory if $GOBIN is not set
- The ~/go/bin directory by default on most systems if neither is explicitly configured