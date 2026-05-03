1. Transaction's maxFeePerGas (15000000) is less than the block's baseFeePerGas (696185213)
```sh
truffle test test/simplestorage.js
#Contract: SimpleStorage
#  1 failing

#  1) Contract: SimpleStorage
#       ...should store the value 89.:
#     Error: VM Exception while processing transaction: Transaction's maxFeePerGas (15000000) is less than the block's baseFeePerGas (696185213) (vm hf=shanghai -> block -> tx)
#      at Context.<anonymous> (test/simplestorage.js:13:44)
#      at processTicksAndRejections (node:internal/process/task_queues:104:5)

```
2. 
```sh
truffle test test/simplestorage.js --config truffle-config.ovm.js --network dashboard
#2 failing

#  1) Contract: SimpleStorage
#       ...should store the value 89.:
#     AssertionError: The value 89 was not stored.: expected <BN: 0> to equal 89
#      at Context.<anonymous> (test/simplestorage.js:18:12)
#      at processTicksAndRejections (node:internal/process/task_queues:104:5)

#  2) Contract: SimpleStorage
#       "after each" hook: after test for "...should store the value 89.":
#     Error: Internal JSON-RPC error.
#      at /home/thuy/.nvm/versions/node/v24.14.1/lib/node_modules/truffle/build/webpack:/packages/provider/wrapper.js:25:1
#      at /home/thuy/.nvm/versions/node/v24.14.1/lib/node_modules/truffle/build/webpack:/packages/provider/wrapper.js:166:1
#      at /home/thuy/.nvm/versions/node/v24.14.1/lib/node_modules/truffle/build/webpack:/node_modules/web3-providers-http/lib/index.js:127:1
#      at processTicksAndRejections (node:internal/process/task_queues:104:5)
```