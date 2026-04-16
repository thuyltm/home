import hardhatToolboxViemPlugin from "@nomicfoundation/hardhat-toolbox-viem";
import { configVariable, defineConfig } from "hardhat/config";
import hardhatVerify from "@nomicfoundation/hardhat-verify";
import hardhatEthers from "@nomicfoundation/hardhat-ethers";
import hardhatTypechain from "@nomicfoundation/hardhat-typechain";
import hardhatMocha from "@nomicfoundation/hardhat-mocha";
import hardhatEthersChaiMatchers from "@nomicfoundation/hardhat-ethers-chai-matchers";
import hardhatNetworkHelpers from "@nomicfoundation/hardhat-network-helpers";




export default defineConfig({
  plugins: [
    hardhatToolboxViemPlugin,
    hardhatVerify,
    hardhatEthers,
    hardhatTypechain,
    hardhatMocha,
    hardhatEthersChaiMatchers,
    hardhatNetworkHelpers,
  ],
  verify: {
    etherscan: {
      apiKey: "CISIXT4FX7TBKVKK9AHASJNC79TH8M1D3W",
    },
  },
  solidity: {
    profiles: {
      default: {
        version: "0.8.28",
      },
      production: {
        version: "0.8.28",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200,
          },
        },
      },
    },
  },
  networks: {
    hardhatMainnet: {
      type: "edr-simulated",
      chainType: "l1",
    },
    hardhatOp: {
      type: "edr-simulated",
      chainType: "op",
    },
    hardhat: {
      type: "edr-simulated",
      chainId: 1337,
    },
    localhost: {
      type: "http", // Required even for local JSON-RPC nodes
      url: "http://127.0.0.1:8545"
    },
    sepolia: {
      type: "http",
      chainType: "l1",
      url: configVariable("SEPOLIA_RPC_URL"),
      accounts: [configVariable("SEPOLIA_PRIVATE_KEY")],
    },
  },
});
