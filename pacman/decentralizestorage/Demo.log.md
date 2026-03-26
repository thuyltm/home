We created two servers, CENTRAL_NODE and WALLET_NODE. 
### Synchronize 2 node without MINER NODE
To simplify two-node synchoronization, we will designate a CENTRAL_NODE with block-addition functionality
1. Initialize CENTRAL_NODE 
```sh
% export NODE_ID=3000
% bazel run //pacman/decentralizestorage -- createwallet 
Your new address: 1PGy2fpZxbEmNY5qyhJLGhfZNm9mmtPm9X
wallet_3000.dat
% bazel run //pacman/decentralizestorage:main -- startnode
Starting node 3000
No existing blockchain found. Create one first.
% bazel run //pacman/decentralizestorage:main -- createblockchain -address 1PGy2fpZxbEmNY5qyhJLGhfZNm9mmtPm9X
blockchain_3000.db
% bazel run //pacman/decentralizestorage:main -- startnode
Starting node 3000
```
2. Initialize WALLET_NODE

We clone blockchain_3000.db to blockchain_3001.db for the WALLET_NODE to simplify database content comparasion via height
```sh
% export NODE_ID=3001
% cp blockchain_3000.db blockchain_3001
% bazel run //pacman/decentralizestorage:main -- createwallet
Your new address: 1BRk4G1BuB1muz6RHjoDNqZ6QnWcEL68gZ
% bazel run //pacman/decentralizestorage:main -- startnode # with blockchain_3001.db
Starting node 3001
localhost:3001 sends a version request to localhost:3000 in order to compare block or transaction heighs
```
**the CENTRAL_NODE console output is at this time**
```sh
Received version command
myBestHeight 1, foreignBestHeight 1
```
3. Stop CENTRAL_NODE and WALLET_NODE server
4. Create new block in CENTRAL_NODE
```sh
% bazel run //pacman/decentralizestorage:main -- send -from 1PGy2fpZxbEmNY5qyhJLGhfZNm9mmtPm9X -to 1BRk4G1BuB1muz6RHjoDNqZ6QnWcEL68gZ -amount 2 -mine
000000b0f09ee9b883442c1c92d218d95ef5061da8eee2a220c38ae40c7dd798
Success!
% bazel run //pacman/decentralizestorage:main -- startnode
```
5. Start CENTRAL_NODE server firstly, followed by the WALLET_NODE server which will connect to the CENTRAL_NODE to synchronize data

**WALLET_NODE server console output**
```sh
% bazel run //pacman/decentralizestorage:main -- startnode
Starting node 3001
localhost:3001 sends a version request to localhost:3000 in order to compare block or transaction heighs
Received version command
myBestHeight 1, foreignBestHeight 2
receiving a getblocks request, localhost:3000 insert a real block
Received inv command
Recevied inventory with 2 block
localhost:3000 send getdata request to get a real block or transaction via a hash
Received block command
Received a new block!
Added block 000000b0f09ee9b883442c1c92d218d95ef5061da8eee2a220c38ae40c7dd798
localhost:3000 send getdata request to get a real block or transaction via a hash
Received block command
Received a new block!
Added block 000000d9ebff813fe87794a9ca7427fbaa910c36eee5b9e787c6fde339ce64e7
```

**Terminal log from CENTRAL_NODE server**
```sh
Starting node 3000
Received version command
myBestHeight 2, foreignBestHeight 1localhost:3000 sends a version request to localhost:3001 in order to compare block or transaction heighs
Received getblocks command
localhost:3001 receive a version request containing a list of missing block or transaction hashes
Received getdata command
Received getdata command
```
