#! /bin/bash
forge init priceoracle
rm -rf .git
git rm -r --cached .
forge install foundry-rs/forge-std
forge install rari-capital/solmate
forge install paulrberg/prb-math
cd lib/prd-math
git checkout e33a042e4d1673fe9b333830b75c4765ccf3f5f2
forge install abdk-consulting/abdk-libraries-solidity
forge install GNSPS/solidity-bytes-utils
npm install
# Error: Compiler error (/solidity/libsolidity/codegen/LValue.cpp:50): Stack too deep. 
# Try compiling with `--via-ir` (cli) or the equivalent `viaIR: true` (standard JSON) while enabling the optimizer
forge build --via-ir
forge test --via-ir
# running local blockchain
anvil --code-size-limit 1000000
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
forge script scripts/DeployDevelopment.s.sol --broadcast --fork-url http://localhost:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80  --code-size-limit 1000000
#Script ran successfully.

#== Logs ==
#  WETH address 0x5FbDB2315678afecb367f032d93F642f64180aa3
#  UNI address 0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0
#  USDC address 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512
#  USDT address 0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9
#  WBTC address 0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9
#  Factory address 0x5FC8d32690cc91D4c39d9d3abcBD16989F875707
#  Manager address 0x0165878A594ca255338adfa4d48449f69242Eb8F
#  Quoter address 0xa513E6E4b8f2a923D98304ec87F64353C4D5C853
#  USDT/USDC address 0xe2833013999eDc77f123714a709B5E2D78480D8c
#  WBTC/USDT address 0x7Fd85cc2DD4419d927200d2216317AFA2b22f07D
#  WETH/UNI address 0x8d9Ae3969dD32a0e2eC55e791a89e868F585B17b
#  WETH/USDC address 0xC68bE465082A8E0B93d70a1F408B472884bb628F