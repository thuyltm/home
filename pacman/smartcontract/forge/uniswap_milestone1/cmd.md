1. Create a new project
```sh
forge init hello_foundry
```
2. Build contract
```sh
forge build
```
3. Run test
```sh
forge test
```
4. Add library
```sh
forge install transmissions11/solmate
```
5. Start a local Ethereum node
```sh
anvil --code-size-limit 50000
#                          (_) | |
#      __ _   _ __   __   __  _  | |
#     / _` | | '_ \  \ \ / / | | | |
#    | (_| | | | | |  \ V /  | | | |
#     \__,_| |_| |_|   \_/   |_| |_|

#    0.1.0 (d89f6af 2022-06-24T00:15:17.897682Z)
#    https://github.com/foundry-rs/foundry
#...
# Listening on 127.0.0.1:8545
```
Due to a big contracts that don't fit into the Ethereum contract size limit, we need to configure Anvil to allow bigger smart contracts
6. Deploy smart contracts to Anvil
```sh
forge script script/DeployDevelopment.s.sol --broadcast --rpc-url http://localhost:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80  --code-size-limit 50000
```
