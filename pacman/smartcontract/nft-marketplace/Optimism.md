### Optimism vs Ethereum
- **Optimism** is a Layer-2 (L2) scaling solution designed to make Ethereum transactions faster and roughly 10x-100x cheaper, while inheriting Ethereum's security. 
- **Ethereum** (L1) serves as the secure, decentralized settlement layer

Key Comparisions:
- Relationship: Optimism is built on top of Ethereum. It is not a competitor, but rather an extension aimed at scaling Ethereum's ecosystem
- Cost and Speed: Optimism offers significantly lower fees and faster transaction times compared to Ethereum mainnet which is expensive and congested
- Compatibility:  Optimism is EVM (Ethereum Virtual Machine) equivalent

When to use which:
- Use Ethereum (L1) for high-value transaction, security-critical actions and bridging large assets
- Use Optimism (L2) for daily DeFi usage, NFT minting, and small transactions where low fees and speed are prioritized

### GOERLI
Goerli was a Ethereum testnet, using Proof-of-Authority. As of 2026, it is no longer recommended or supported for development, with the ecosystem having full migrated to Sepolia for testing smart contracts and dApps

### Setup testing environment
**@openzeppelin/test-helpers/configure** is a module within the **OpenZeppelin Test Helpers** library that initializes and configure the environment--specifically the web3 provider--before tests run