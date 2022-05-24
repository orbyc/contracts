const hre = require("hardhat");

async function main() {
  const Array = await ethers.getContractFactory("Array");
  const array = await Array.deploy();
  await array.deployed();

  const SupplyChain = await ethers.getContractFactory("SupplyChain", {
    libraries: {
      Array: array.address,
    },
  });
  const supply = await SupplyChain.deploy("Orbyc Supply Chain", "OSC");
  await supply.deployed();

  console.log("Supply Chain deployed to:", supply.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
