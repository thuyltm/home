#! /bin/bash
forge init flashloans
rm -rf .git
forge install foundry-rs/forge-std
forge install rari-capital/solmate
forge install paulrberg/prb-math
cd lib/prd-math
git checkout e33a042e4d1673fe9b333830b75c4765ccf3f5f2
forge install abdk-consulting/abdk-libraries-solidity
forge build
forge test
# running local blockchain
anvil --code-size-limit 50000
#                             _   _
#                            (_) | |
#      __ _   _ __   __   __  _  | |
#     / _` | | '_ \  \ \ / / | | | |
#    | (_| | | | | |  \ V /  | | | |
#     \__,_| |_| |_|   \_/   |_| |_|

#    0.1.0 (d89f6af 2022-06-24T00:15:17.897682Z)
#    https://github.com/foundry-rs/foundry
#...
# Listening on 127.0.0.1:8545
# deploy to smartcontract to Anvil Ethereum node
forge script scripts/DeployDevelopment.s.sol --broadcast --fork-url http://localhost:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80  --code-size-limit 50000
#== Logs ==
#  WETH address 0x5FbDB2315678afecb367f032d93F642f64180aa3
#  USDC address 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512
#  Pool address 0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0
#  Manager address 0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9
#  Quoter address 0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9