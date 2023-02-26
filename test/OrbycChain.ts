import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { ethers } from "hardhat";

describe("Lock", function () {
  async function deployOrbycChainFixture() {
    const [owner, otherAccount] = await ethers.getSigners();

    const Accounts = await ethers.getContractFactory("AccountControlMock");
    const accounts = await Accounts.deploy();

    const SupplyChain = await ethers.getContractFactory("SupplyChain");
    const supplyChain = await SupplyChain.deploy(
      accounts.address,
      "orbyc Supply Chain",
      "https://wallet.orbyc.com/metadata/{id}"
    );

    return { supplyChain, owner, otherAccount };
  }

  describe("Deployment", function () {
    it("Should set the right role", async function () {
      const { supplyChain, owner } = await loadFixture(deployOrbycChainFixture);

      expect(await supplyChain.hasRole(1, owner.address)).to.be.true;
    });
  });
});
