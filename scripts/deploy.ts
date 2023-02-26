import { ethers } from "hardhat";

async function main() {
  const Accounts = await ethers.getContractFactory("AccountControlMock");
  const accounts = await Accounts.deploy();
  await accounts.deployed();

  const SupplyChain = await ethers.getContractFactory("SupplyChain");
  const supplyChain = await SupplyChain.deploy(
    accounts.address,
    "orbyc Supply Chain",
    "https://wallet.orbyc.com/metadata/{id}"
  );
  await supplyChain.deployed();

  console.log("address:", supplyChain.address);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
