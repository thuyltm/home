#!/bin/bash
npx hardhat clean
npx hardhat compile
npx hardhat test test/Token.ts
# You are running Hardhat Ignition against an in-process instance of Hardhat Network.
# This will execute the deployment, but the results will be lost.
npx hardhat ignition deploy ./ignition/modules/Token.ts
# The password for keystore is Hardhat@123
npx hardhat keystore set SEPOLIA_RPC_URL -force
npx hardhat keystore set SEPOLIA_PRIVATE_KEY

npx hardhat ignition deploy ./ignition/modules/Token.ts --network sepolia
# [hardhat-keystore] Enter the password: ***********
npx hardhat ignition deploy ./ignition/modules/Token.ts --network sepolia
#[hardhat-keystore] Enter the password: ***********
#✔ Confirm deploy to network sepolia (11155111)? … yes
#Hardhat Ignition 🚀

#Deploying [ TokenModule ]

#Batch #1
#  Executed TokenModule#Token

#[ TokenModule ] successfully deployed 🚀

#Deployed Addresses

#TokenModule#Token - 0x32a4705621c428E2ae6Bef323C1AC7B97DC325B2

npx hardhat verify --network sepolia 0x222813e15F983ED5dc1897249043BF0D7bDB6311
#=== Etherscan ===
#[hardhat-keystore] Enter the password: ***********

#📤 Submitted source code for verification on Etherscan:

#  contracts/Token.sol:Token
#  Address: 0x32a4705621c428E2ae6Bef323C1AC7B97DC325B2

#⏳ Waiting for verification result...


#✅ Contract verified successfully on Etherscan!

#  contracts/Token.sol:Token
#  Explorer: https://sepolia.etherscan.io/address/0x32a4705621c428E2ae6Bef323C1AC7B97DC325B2#code

#=== Blockscout ===

#📤 Submitted source code for verification on Blockscout:

#  contracts/Token.sol:Token
#  Address: 0x32a4705621c428E2ae6Bef323C1AC7B97DC325B2

#⏳ Waiting for verification result...


#✅ Contract verified successfully on Blockscout!

#  contracts/Token.sol:Token
#  Explorer: https://eth-sepolia.blockscout.com/address/0x32a4705621c428E2ae6Bef323C1AC7B97DC325B2#code

#=== Sourcify ===

#The contract at 0x32a4705621c428E2ae6Bef323C1AC7B97DC325B2 has already been verified on Sourcify.

#If you need to verify a partially verified contract, please use the --force flag.

#Explorer: https://sourcify.dev/server/repo-ui/11155111/0x32a4705621c428E2ae6Bef323C1AC7B97DC325B2