import { ethers } from "hardhat";

async function main() {
  const accounts = "0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d";

  const OrbycChain = await ethers.getContractFactory("OrbycChain");
  const orbyc = await OrbycChain.deploy("orbyc Chain", "OCH", accounts);
  await orbyc.deployed();

  console.log("address:", orbyc.address);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
