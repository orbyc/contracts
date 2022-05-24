require("@nomiclabs/hardhat-waffle");
require("dotenv").config();

// This is a sample Hardhat task. To learn how to create your own go to
// https://hardhat.org/guides/create-task.html
task("accounts", "Prints the list of accounts", async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners();

  for (const account of accounts) {
    console.log(account.address);
  }
});

// You need to export an object to set up your config
// Go to https://hardhat.org/config/ to learn more
ROPSTEN_PRIVATE_KEY = "";
/**
 * @type import('hardhat/config').HardhatUserConfig
 */
module.exports = {
  solidity: {
    version: "0.8.4",
    settings: {
      evmVersion: "constantinople",
      optimizer: {
        enabled: true,
        runs: 200,
      },
    },
  },

  networks: {
    hardhat: {
      allowUnlimitedContractSize: true,
    },
    ropsten: {
      url: `https://speedy-nodes-nyc.moralis.io/9dcfc2a35a472aedc480f310/eth/ropsten`,
      accounts: [process.env.ACCOUNT_PK],
    },
  },
};
