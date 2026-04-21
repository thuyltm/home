1> Start a local Harhat network with 10 pre-funded accounnt, each having 10,000 ETH
```sh
npx hardhat node
# Account #0:  0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266 (10000 ETH)
# Private Key: 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
```
2> Setup a Hardhat network for MetaMask with these below information
Default RPC URL: http://127.0.0.1:8545
Chain ID: 31337
3> Import the first Hardhat development account into MetaMask via the "Import an account" option within 'Add Wallet' menu. This account comes with 10000 ETH
4> Deploy Token smartcontract on a Hardhat network
```sh
npx hardhat ignition deploy ./ignition/modules/Token.ts --network localhost
#Hardhat Ignition 🚀

#Deploying [ TokenModule ]

#Batch #1
  #Executed TokenModule#Token

#[ TokenModule ] successfully deployed 🚀

#Deployed Addresses

#TokenModule#Token - 0x5FbDB2315678afecb367f032d93F642f64180aa3
```
The Hardhat network console log is at this time
```sh
eth_call
  Contract deployment: Token
  Contract address:    0x5fbdb2315678afecb367f032d93f642f64180aa3
  From:                0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266

eth_sendTransaction
  Contract deployment: Token
  Contract address:    0x5fbdb2315678afecb367f032d93f642f64180aa3
  Transaction:         0x63622997d8806675928c1366026c167c353155b0034c9d43d9a375e8c8d74c3b
  From:                0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266
  Value:               0 ETH
  Gas used:            439462 of 439462
  Block #1:            0x4791d173815f2e6ae71d313455e0db074fe7fd713c97753348dd36c9ef12ab2e
```
5> Copy the deployed contract address from the file `ignition/deployements/chain-31337/deployed_addresses.json` to `frontend/src/App.js`
```javascript
const tokenAddress = "0x5fbdb2315678afecb367f032d93f642f64180aa3"
```
6> Import the deployed contract address to the MetaMask. This account was allocated 100,000 MHT and is visibile on MetaMask
7> Start npm server
```sh
npm start
```
8> The `window.ethereum.request({ method: 'eth_requestAccounts'})` code triggers a a popup, asking for persmission to coonect user Ethereum wallet (such as MetaMask) to a dApp
9> Send some MDT Token to another account


