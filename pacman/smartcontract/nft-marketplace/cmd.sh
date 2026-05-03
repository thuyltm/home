#! /bin/bash
truffle unbox optimism nft-marketplace
npm install
truffle compile
truffle dashboard
truffle migrate --config truffle-config.ovm.js --network dashboard
truffle test test/marketplace.js
#Contract: Marketplace
#marketplace 0xF12b5dd4EAD5F743C6BaA640B0216200e89B60Da
#token_owner 0xf17f52151EbEF6C7334FAD080c5704D77216b732
#buyer 0xC5fdf4076b8F3A5357c5E395ab970B5B54098Fef
#    ✔ should validate before listing (245ms)
#    ✔ should list nft (377ms)
#    ✔ should validate before buying
#    ✔ should modify listings when nft is bought (544ms)
#    ✔ should validate reselling
#    ✔ should resell nft (801ms)


#  6 passing (2s)
truffle test test/simplestorage.js --config truffle-config.ovm.js --network ganache
#Contract: SimpleStorage
#   ✔ ...should store the value 89.

# 1 passing (85ms)
truffle exec scripts/run.js --config truffle-config.ovm.js --network ganache
#Using network 'ganache'.

#MINT AND LIST 3 NFTs
#Minted and listed 36
#Minted and listed 37
#Minted and listed 38
#listedNfts: 12
#myNfts: 0
#myListedNfts 0

#BUY 2 NFTs
#listedNfts: 10
#myNfts: 0
#myListedNfts 0

#RESELL 1 NFT
#listedNfts: 11
#myNfts: 0
#myListedNfts 0