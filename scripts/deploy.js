const hre = require("hardhat");

async function main() {
  const Suite = await hre.ethers.getContractFactory("Suite");
  const suite = await Suite.deploy();

  await suite.deployed();

  console.log("Agents suite deployed to:", suite.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
