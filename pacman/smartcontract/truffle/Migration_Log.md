truffle migrate
This version of µWS is not compatible with your Node.js build:

Error: Cannot find module '../binaries/uws_linux_x64_137.node'
Require stack:
- /home/thuy/.nvm/versions/node/v24.14.1/lib/node_modules/truffle/node_modules/ganache/node_modules/@trufflesuite/uws-js-unofficial/src/uws.js
- /home/thuy/.nvm/versions/node/v24.14.1/lib/node_modules/truffle/node_modules/ganache/dist/node/core.js
- /home/thuy/.nvm/versions/node/v24.14.1/lib/node_modules/truffle/build/migrate.bundled.js
- /home/thuy/.nvm/versions/node/v24.14.1/lib/node_modules/truffle/node_modules/original-require/index.js
- /home/thuy/.nvm/versions/node/v24.14.1/lib/node_modules/truffle/build/cli.bundled.js
Falling back to a NodeJS implementation; performance may be degraded.



Compiling your contracts...
===========================
> Everything is up to date, there is nothing to compile.


Starting migrations...
======================
> Network name:    'development'
> Network id:      5777
> Block gas limit: 6721975 (0x6691b7)


1_initial_migration.js
======================

   Deploying 'Migration'
   ---------------------
   > transaction hash:    0x48787d2dbe41ef5cd236e0414e95618205c09f7bf20ded46dbe66cbf4f19ca71
   > Blocks: 0            Seconds: 0
   > contract address:    0x32a4705621c428E2ae6Bef323C1AC7B97DC325B2
   > block number:        8
   > block timestamp:     1775720244
   > account:             0x2C20387adbF65cd74010b2693fdA5E3B4039538E
   > balance:             99.991401248000211499
   > gas used:            155222 (0x25e56)
   > gas price:           2.962400908 gwei
   > value sent:          0 ETH
   > total cost:          0.000459829793741576 ETH

   > Saving artifacts
   -------------------------------------
   > Total cost:     0.000459829793741576 ETH


2_deploy_contract.js
====================

   Replacing 'Adoption'
   --------------------
   > transaction hash:    0x7b769475846a7026886e46d72ec3a1c122ab89f43e77e2745d412e828f45cfda
   > Blocks: 0            Seconds: 0
   > contract address:    0x33cb6439C93fB0a9178a4B34a02A87C594e27374
   > block number:        9
   > block timestamp:     1775720244
   > account:             0x2C20387adbF65cd74010b2693fdA5E3B4039538E
   > balance:             99.990907555725059647
   > gas used:            169813 (0x29755)
   > gas price:           2.907270204 gwei
   > value sent:          0 ETH
   > total cost:          0.000493692275151852 ETH

   > Saving artifacts
   -------------------------------------
   > Total cost:     0.000493692275151852 ETH

Summary
=======
> Total deployments:   2
> Final cost:          0.000953522068893428 ETH