
1. Run Truffle Dashboard in a separate terminal window
```sh
truffle dashboard
#Truffle Dashboard running at http://localhost:24012
#DashboardProvider RPC endpoint running at http://localhost:24012/rpc
```
2. Deploying contracts on Truffle Dashboard requires you to approve the signature request from MetaMask account
```sh
truffle migrate --config truffle-config.js --network dashboard
#This version of µWS is not compatible with your Node.js build:

#Error: Cannot find module '../binaries/uws_linux_x64_137.node'
#Require stack:
#- /home/thuy/.nvm/versions/node/v24.14.1/lib/node_modules/truffle/node_modules/ganache/node_modules/@trufflesuite/uws-js-unofficial/src/uws.js
#- /home/thuy/.nvm/versions/node/v24.14.1/lib/node_modules/truffle/node_modules/ganache/dist/node/core.js
#- /home/thuy/.nvm/versions/node/v24.14.1/lib/node_modules/truffle/build/migrate.bundled.js
#- /home/thuy/.nvm/versions/node/v24.14.1/lib/node_modules/truffle/node_modules/original-require/index.js
#- /home/thuy/.nvm/versions/node/v24.14.1/lib/node_modules/truffle/build/cli.bundled.js
#Falling back to a NodeJS implementation; performance may be degraded.



#Compiling your contracts...
#===========================
#> Everything is up to date, there is nothing to compile.


#Starting migrations...
#======================
#> Network name:    'dashboard'
#> Network id:      11155420
#> Block gas limit: 40000000 (0x2625a00)


#1_deploy_contracts.js
#=====================

#   Replacing 'SimpleStorage'
#   -------------------------
#   > transaction hash:    0xb69d78de45ab8c2997672c2893d60d20505df1fa6c157359c8fc922bf25e20fasage.
#   > Blocks: 0            Seconds: 0
#   > contract address:    0x2619A20EC4F6ea4A4AD5e323363C91dA3D1eB11f
#   > block number:        42988279
#   > block timestamp:     1777779098
#   > account:             0x2C20387adbF65cd74010b2693fdA5E3B4039538E
#   > balance:             0.095207942020779921
#   > gas used:            125677 (0x1eaed)
#   > gas price:           2.5000005 gwei
#   > value sent:          0 ETH
#   > total cost:          0.0003141925628385 ETH

#   > Saving artifacts
#   -------------------------------------
#   > Total cost:     0.0003141925628385 ETH


#2_deploy_marketplace.js
#=======================

#   Deploying 'Marketplace'
#   -----------------------
#   > transaction hash:    0x0ff2464bedcfc4dddaece2588666f4f74559aee6df37b12449b4fd4aa4f54744sage.
#   > Blocks: 11           Seconds: 25
#   > contract address:    0x30eBe4F9a578F52CC5DFA41B585d9c7b78Eb0F63
#   > block number:        42988293
#   > block timestamp:     1777779126
#   > account:             0x2C20387adbF65cd74010b2693fdA5E3B4039538E
#   > balance:             0.091044269104401527
#   > gas used:            1665469 (0x1969bd)
#   > gas price:           2.5000005 gwei
#   > value sent:          0 ETH
#   > total cost:          0.0041636733327345 ETH


#   Deploying 'BoredPetsNFT'
#   ------------------------
#   > transaction hash:    0x1d9845b77745182f8096ad4f2cdbfa669468a352de98baf2a582c569a0bc2887sage.
#   > Blocks: 6            Seconds: 12
#   > contract address:    0x27eB387e3646b95645D3E7488B9Af9343aB691Cc
#   > block number:        42988304
#   > block timestamp:     1777779148
#   > account:             0x2C20387adbF65cd74010b2693fdA5E3B4039538E
#   > balance:             0.091044269104401527
#   > gas used:            2429201 (0x251111)
#   > gas price:           2.5000005 gwei
#   > value sent:          0 ETH
#   > total cost:          0.0060730037146005 ETH

#   > Saving artifacts
#   -------------------------------------
#   > Total cost:     0.010236677047335 ETH

#Summary
#=======
#> Total deployments:   3
#> Final cost:          0.0105508696101735 ETH#
```
