#! /bin/bash
forge init quotecontract
rm -rf .git
forge install foundry-rs/forge-std
forge install transmissions11/solmate
forge install paulrberg/prb-math
cd lib/prd-math
git checkout e33a042e4d1673fe9b333830b75c4765ccf3f5f2
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
forge script script/DeployDevelopment.s.sol --broadcast --fork-url http://localhost:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80  --code-size-limit 50000
#== Logs ==
#  WETH address 0x5FbDB2315678afecb367f032d93F642f64180aa3
#  USDC address 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512
#  Pool address 0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0
#  Manager address 0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9

## Setting up 1 EVM.

#==========================

#Chain 31337

#Estimated gas price: 2.000000001 gwei

#Estimated total gas used for script: 6016868

#Estimated amount required: 0.012033736006016868 ETH

#==========================

##### anvil-hardhat
#✅  [Success] Hash: 0x5291a18eb785c7d12a11faacdfd17955a9f4a26a1ea586f2e9e801ba6b7d6400
#Contract: ERC20Mintable
#Contract Address: 0x5FbDB2315678afecb367f032d93F642f64180aa3
#Block: 1
#Paid: 0.001308185001308185 ETH (1308185 gas * 1.000000001 gwei)


##### anvil-hardhat
#✅  [Success] Hash: 0x9f4a7205b0630e62c9c1c9a257ef530e2aded41041108e318b72314e950f3865
#Contract: UniswapV3Manager
#Contract Address: 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512
#Block: 2
#Paid: 0.001158869955936875 ETH (1308125 gas * 0.885901543 gwei)


##### anvil-hardhat
#✅  [Success] Hash: 0xd360844d9e3dcdaf7534aca3c7b01dc93e729fabb6292fbfc7fd35df95228965
#Contract: ERC20Mintable
#Function: mint(address,uint256)
#Block: 5
#Paid: 0.000042065205625212 ETH (68794 gas * 0.611466198 gwei)


##### anvil-hardhat
#✅  [Success] Hash: 0x1e8dafa14f543de0d048706dec705212e51063e50d9e71cead4cc3aab342d47a
#Contract: ERC20Mintable
#Contract Address: 0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0
#Block: 2
#Paid: 0.001051881398391851 ETH (1187357 gas * 0.885901543 gwei)


##### anvil-hardhat
#✅  [Success] Hash: 0x74e8b68a308644c84efb135357fa5be8970ad7a1096f059959f219c60402e3eb
#Contract: UniswapV3Pool
#Block: 5
#Paid: 0.000042079880813964 ETH (68818 gas * 0.611466198 gwei)


##### anvil-hardhat
#✅  [Success] Hash: 0x55bb91dfc3e1ed5dc00eebce4e35d8a48ae46e63e30154d0df4d8f6568dc9c54
#Contract: ERC20Mintable
#Function: mint(address,uint256)
#Contract Address: 0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9
#Block: 4
#Paid: 0.000465158311605024 ETH (669882 gas * 0.694388432 gwei)

#✅ Sequence #1 on anvil-hardhat | Total Paid: 0.004068239753681111 ETH (4611161 gas * avg 0.781520652 gwei)
                                                                                                                                                                         

#==========================

#ONCHAIN EXECUTION COMPLETE & SUCCESSFUL.

#Transactions saved to: /home/thuy/Documents/Learn/home/pacman/smartcontract/uniswap/quotecontract/broadcast/DeployDevelopment.s.sol/31337/run-latest.json

#Sensitive values saved to: /home/thuy/Documents/Learn/home/pacman/smartcontract/uniswap/quotecontract/cache/DeployDevelopment.s.sol/31337/run-latest.json