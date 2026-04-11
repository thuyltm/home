New Solidity version (v0.8.20+) use default settings for the Shanghai or Cancun hardforks, which include opcodes (like PUSH0 or MCOPY) that your version of Ganache does not support

#### Common Fixes
- Option 1: Downgrade Solidity (Recommended)

Change the pragma in your contract **pragma solidity 0.8.19**; or update your compiler version in truffle-config.js to 0.8.19 to avoid generate the incompatible PUSH0 opcode by default

- Option 2: Manually set the EVM Version

If you must use a newer Solidity version, you can tell the compiler to target an older EVM version that Ganache supports (like Paris, London, or Berlin)

Update truffle-config.js
```sh
compilers: {
  solc: {
    version: "0.8.20", // or your version
    settings: {
      evmVersion: "paris" 
    }
  }
}
```

- Option 3: Use a Modern Alternative

Since Ganache is deprecated and no longer frequently updated to support the latest Ethereum hardforks, 
many developers have move to Hardhat or Foundry for better support for modern Solidity opcodes