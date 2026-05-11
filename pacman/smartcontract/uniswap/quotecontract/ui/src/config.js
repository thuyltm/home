const config = {
  token0Address: '0x0165878A594ca255338adfa4d48449f69242Eb8F',
  token1Address: '0xa513E6E4b8f2a923D98304ec87F64353C4D5C853',
  poolAddress: '0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6',
  managerAddress: '0x8A791620dd6260079BF849Dc5567aDC3F2FdC318',
  quoterAddress: '0x610178dA211FEF7D417bC0e6FeD39F05609AD788',
  ABIs: {
    'ERC20': require('./abi/ERC20.json'),
    'Pool': require('./abi/Pool.json'),
    'Manager': require('./abi/Manager.json'),
    'Quoter': require('./abi/Quoter.json')
  }
};

export default config;