```sh
truffle migrate --config truffle-config.ovm.js --network optimistic_sepolia


#Compiling your contracts...
#===========================
#> Everything is up to date, there is nothing to compile.


#Starting migrations...
#======================
#> Network name:    'optimistic_sepolia'
#> Network id:      11155420
#> Block gas limit: 40000000 (0x2625a00)


#1_deploy_contracts.js
#=====================

#   Deploying 'SimpleStorage'
#   -------------------------
#   > transaction hash:    0x7635e96a93b18a6538bbda6bb6234a5ae1e5f716cce7311dc41a74f0089d03ec
#   > Blocks: 6            Seconds: 18
#   > contract address:    0x3d7004E419A47039DB41C1B357B5CAF9DD0bc183
#   > block number:        42997441
#   > block timestamp:     1777797422
#   > account:             0x2C20387adbF65cd74010b2693fdA5E3B4039538E
#   > balance:             0.058939761073951494
#   > gas used:            90565 (0x161c5)
#   > gas price:           2.50000025 gwei
#   > value sent:          0 ETH
#   > total cost:          0.00022641252264125 ETH

#   > Saving artifacts
#   -------------------------------------
#   > Total cost:     0.00022641252264125 ETH


#2_deploy_marketplace.js
#=======================

#   Deploying 'Marketplace'
#   -----------------------
#   > transaction hash:    0xcd60d8ac9ca1097fcd03d2b4fe7e4bea474ae804f562d2f4f77dec12c7818ac7
#   > Blocks: 9            Seconds: 22
#   > contract address:    0x944e20193632B7Cc748DCece5Ee8ABB6bC539CF2
#   > block number:        42997460
#   > block timestamp:     1777797460
#   > account:             0x2C20387adbF65cd74010b2693fdA5E3B4039538E
#   > balance:             0.056495490829515919
#   > gas used:            977708 (0xeeb2c)
#   > gas price:           2.50000025 gwei
#   > value sent:          0 ETH
#   > total cost:          0.002444270244427 ETH


#   Deploying 'BoredPetsNFT'
#   ------------------------
#   > transaction hash:    0x54b3ecfb63714f6f86755fcc61dcff4b71175bdb76e8aed85b8ad157ea303f20
#   > Blocks: 4            Seconds: 10
#   > contract address:    0x278c100CB40E14384032A3536616642EaFEDbcb2
#   > block number:        42997477
#   > block timestamp:     1777797494
#   > account:             0x2C20387adbF65cd74010b2693fdA5E3B4039538E
#   > balance:             0.052872322967184939
#   > gas used:            1449267 (0x161d33)
#   > gas price:           2.50000025 gwei
#   > value sent:          0 ETH
#   > total cost:          0.00362316786231675 ETH

#   > Saving artifacts
#   -------------------------------------
#   > Total cost:     0.00606743810674375 ETH

#Summary
#=======
#> Total deployments:   3
#> Final cost:          0.006293850629385 ETH
```