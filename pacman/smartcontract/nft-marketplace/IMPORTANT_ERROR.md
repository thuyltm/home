The belowed error indicates that Metamask is not yet connected to your Local Dapp (for example Truffle Dashboard)
```sh
Error: Expected parameter 'from' not passed to function.
    at has (/home/thuy/.nvm/versions/node/v24.14.1/lib/node_modules/truffle/build/webpack:/packages/expect/dist/src/index.js:10:1)
    at Object.options (/home/thuy/.nvm/versions/node/v24.14.1/lib/node_modules/truffle/build/webpack:/packages/expect/dist/src/index.js:19:1)
    at Object.<anonymous> (/home/thuy/.nvm/versions/node/v24.14.1/lib/node_modules/truffle/build/webpack:/packages/migrate/dist/src/index.js:104:1)
    at Generator.next (<anonymous>)
    at /home/thuy/.nvm/versions/node/v24.14.1/lib/node_modules/truffle/build/webpack:/packages/migrate/dist/src/index.js:31:1
    at new Promise (<anonymous>)
    at exports.modules.22478.__awaiter (/home/thuy/.nvm/versions/node/v24.14.1/lib/node_modules/truffle/build/webpack:/packages/migrate/dist/src/index.js:27:1)
    at Object.run (/home/thuy/.nvm/versions/node/v24.14.1/lib/node_modules/truffle/build/webpack:/packages/migrate/dist/src/index.js:103:1)
    at __webpack_modules__.96146.module.exports (/home/thuy/.nvm/versions/node/v24.14.1/lib/node_modules/truffle/build/webpack:/packages/core/lib/commands/migrate/runMigrations.js:114:1)
    at processTicksAndRejections (node:internal/process/task_queues:104:5)
    at Object.__webpack_modules__.52423.module.exports [as run] (/home/thuy/.nvm/versions/node/v24.14.1/lib/node_modules/truffle/build/webpack:/packages/core/lib/commands/migrate/run.js:41:1)
    at runCommand (/home/thuy/.nvm/versions/node/v24.14.1/lib/node_modules/truffle/build/webpack:/packages/core/lib/command-utils.js:297:1)
```