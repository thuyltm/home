Demo2 demonstrates a 3-node blockchain, where a dedicated MINER_NODE mine new block into its local database and subsequently synchronizes update across for the CENTRAL_NODE and WALLET_NODE
1. Start CENTRAL_NODE server
```sh
% export NODE_ID=3000
% bazel run //pacman/decentralizestorage:main -- createwallet
Your new address: 18VC76Yo4XpHAMDgCN5CPBi6Hk5tRYqFKi
% bazel run //pacman/decentralizestorage:main -- createblockchain -address 18VC76Yo4XpHAMDgCN5CPBi6Hk5tRYqFKi
000000f24a14bd4301bcf4d90662cbbc54c1ff4051bb961be7a9d23502b4bec2
Done!

Success!
% bazel run //pacman/decentralizestorage:main -- startnode
Starting node 3000
```
2. Start MINER_NODE server
```sh
% export NODE_ID=3002
% cp blockchain_genesis.db blockchain_3002.db
% bazel run //pacman/decentralizestorage:main -- createwallet
Your new address: 14K4SdksR7RXzaHnt9RExLSxsRppcwyy4S
% bazel run //pacman/decentralizestorage:main -- startnode -miner 14K4SdksR7RXzaHnt9RExLSxsRppcwyy4S
Starting node 3002
Mining is on. Address to receive rewards:  14K4SdksR7RXzaHnt9RExLSxsRppcwyy4S
I am miner at 14K4SdksR7RXzaHnt9RExLSxsRppcwyy4S
localhost:3002 sends a version request to localhost:3000 in order to compare block or transaction heighs
```
The CENTRAL_NODE console outputs:
```sh
Received version command
myBestHeight 1, foreignBestHeight 1
```
3. Start WALLET_NODE server

```sh
% export NODE_ID=3001
% bazel run //pacman/decentralizestorage:main -- createwallet
Your new address: 1GuwaLK8TjbnmgVc13JcaXkMz7DSwr77BS

% cp blockchain_genesis.db blockchain_3001.db
% bazel run //pacman/decentralizestorage:main -- startnode
Starting node 3001

localhost:3001 sends a version request to localhost:3000 in order to compare block or transaction heighs
```
Stop WALLET_NODE server and begin to demo How to create a new block in a 3-node distributed blockchain system
```sh
% bazel run //pacman/decentralizestorage:main -- getbalance -address 18VC76Yo4XpHAMDgCN5CPBi6Hk5tRYqFKi
Balance of '18VC76Yo4XpHAMDgCN5CPBi6Hk5tRYqFKi': 10
% bazel run //pacman/decentralizestorage:main -- send -from 18VC76Yo4XpHAMDgCN5CPBi6Hk5tRYqFKi -to 1GuwaLK8TjbnmgVc13JcaXkMz7DSwr77BS -amount 1
send to localhost:3000 insert a new transaction
Success!
```
The CENTRAL_NODE console outputs:
```sh
Received tx command

Send to miner node localhost:3002 intentionally

Submit the transaction to localhost:3002 to insert
Received inv command
Received inventory with 1 block

SendGetData: Query localhost:3002 to retrieve a detail block or transaction via a hash
Received block command
Received a new block!
Added block 000045909c78b7369d3e9b7cf3bb9338c1d33431fb2312f82c64599c3c3f0fd5
```
The MINER_NODE console outputs:
```sh
Received tx command
MineBlock is called
000045909c78b7369d3e9b7cf3bb9338c1d33431fb2312f82c64599c3c3f0fd5

New block is mined!

SendInv: synchronize localhost:3000 with a given blocks or transaction hash list
Received getdata command

SendBlock: forward a given block to localhost:3000
```
Finally start WALLET_NODE to sync
```sh
% bazel run //pacman/decentralizestorage:main -- startnode
```
