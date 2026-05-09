Follow these commands in sequence to complete the process
1. Start ganache testnet
2. Start truffle dashboard
```sh
truffle dashboard
```
3. Migrate smart contracts to ganache testnet using truffle dashboard
```sh
truffle migrate --config truffle-config.ovm.js --network dashboard
# Marketplace.json
#{
#  "contractName": "Marketplace",
#  "abi": [
#    "networks": {
#        "5777": {
#            "events": {},
#            "links": {},
#            "address": "0x7C5249a34549c5164da68EF4FEAe95EA177b9875",
#            "transactionHash": "0x770f1737b2395be22365c5dee651c91b2f7498115fc94215e4cdbff6915b6813"
#        }
#    }
#  ]
#}
```
3. Import Ganache account to MetaMask
4. Start UI
```sh
npm install
npm run build
npm run dev
```