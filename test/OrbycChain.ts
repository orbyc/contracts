import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { expect } from "chai";
import { ethers } from "hardhat";

describe("Lock", function () {
  async function deployOrbycChainFixture() {
    const [owner, otherAccount] = await ethers.getSigners();

    const Accounts = await ethers.getContractFactory("AccountControlMock");
    const accounts = await Accounts.deploy();

    const OrbycChain = await ethers.getContractFactory("OrbycChain");
    const orbyc = await OrbycChain.deploy(
      "orbyc Chain",
      "OCH",
      accounts.address
    );

    return { orbyc, owner, otherAccount };
  }

  describe("Deployment", function () {
    it("Should set the right role", async function () {
      const { orbyc, owner } = await loadFixture(deployOrbycChainFixture);

      expect(await orbyc.hasRole(1, owner.address)).to.be.true;
    });
  });
});
