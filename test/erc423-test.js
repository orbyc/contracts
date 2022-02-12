const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("ERC423", function () {
  let suite;
  let owner, addr1, addr2;

  const suiteName = "Test ERC423 Suite";

  beforeEach(async function () {
    const ERC423 = await ethers.getContractFactory("ERC423");
    suite = await ERC423.deploy(suiteName);
    await suite.deployed();

    [owner, addr1, addr2, _] = await ethers.getSigners();
  });

  it("initializes the smart contract", async () => {
    const name = await suite.name();
    expect(name).to.equal(suiteName);
  });

  describe("define agent", function () {
    const account = "0xA001da7DcdF4BD8463823669FB9c43Fcb807EA61";
    const accountInfo = JSON.stringify({ name: "Agent 1", value: 50 });

    beforeEach(async () => {
      await suite.defineAgent(addr1.address, account, accountInfo);
    });

    it("should have associate account with agent's address", async () => {
      const acc = await suite.accountOf(addr1.address);

      expect(acc).to.equal(account);
    });

    it("should return info for the given account", async () => {
      const info = await suite.accountInfo(account);

      expect(info).to.equal(accountInfo);
    });
  });

  describe("remove agent", () => {
    const account = "0xA001da7DcdF4BD8463823669FB9c43Fcb807EA61";

    beforeEach(async () => {
      await suite.defineAgent(addr1.address, account, "");
    });

    it("should be reverted after get account of removed agent", async () => {
      await suite.removeAgent(addr1.address);

      await expect(suite.accountOf(addr1.address)).to.be.revertedWith(
        "ERC423: agent has been removed"
      );
    });
  });

  describe("define role", function () {
    const roleId = 1 << 0;
    const roleInfo = JSON.stringify({ name: "Admin", values: [1, 2, 3] });

    beforeEach(async () => {
      await suite.defineRole(roleId, roleInfo);
    });

    it("should return info for the given role", async () => {
      info = await suite.roleInfo(roleId);

      expect(info).to.equal(roleInfo);
    });
  });

  describe("role endorsement", function () {
    const role1 = 1 << 1;
    const role2 = 1 << 2;

    const account = "0xA001da7DcdF4BD8463823669FB9c43Fcb807EA61";

    beforeEach(async () => {
      await suite.defineRole(role1, "");
      await suite.defineRole(role2, "");

      await suite.defineAgent(addr1.address, account, "");
    });

    describe("when only one role is endorsed", async () => {
      beforeEach(async () => {
        await suite.grantRole(account, role1);
      });

      it("should return true when asks for account to have the role 1", async () => {
        const value = await suite.hasRole(account, role1);

        expect(value).to.equal(true);
      });

      it("should return false when asks for account to have the role 2", async () => {
        const value = await suite.hasRole(account, role2);

        expect(value).to.equal(false);
      });
    });

    describe("when more than one role is endorsed", async () => {
      beforeEach(async () => {
        await suite.grantRole(account, role1);
        await suite.grantRole(account, role2);
      });

      it("should return true when asks for account to have the role 1", async () => {
        const value = await suite.hasRole(account, role1);

        expect(value).to.equal(true);
      });

      it("should return true when asks for account to have the role 2", async () => {
        const value = await suite.hasRole(account, role2);

        expect(value).to.equal(true);
      });
    });
  });

  describe("role revocation", function () {
    const role1 = 1 << 1;
    const role2 = 1 << 2;

    const account = "0xA001da7DcdF4BD8463823669FB9c43Fcb807EA61";

    beforeEach(async () => {
      await suite.defineRole(role1, "");
      await suite.defineRole(role2, "");

      await suite.defineAgent(addr1.address, account, "");

      await suite.grantRole(account, role1);
      await suite.grantRole(account, role2);
    });

    it("should return true when ask for endorsed roles", async () => {
      const hasRole1 = await suite.hasRole(account, role1);
      const hasRole2 = await suite.hasRole(account, role2);

      expect(hasRole1 && hasRole2).to.equal(true);
    });

    describe("when user role 1 is revoked", async () => {
      beforeEach(async () => {
        await suite.revokeRole(account, role1);
      });

      it("should return false when asks for agent to have the role 1", async () => {
        const value = await suite.hasRole(account, role1);

        expect(value).to.equal(false);
      });

      it("should return true when asks for agent to have the role 2", async () => {
        const value = await suite.hasRole(account, role2);

        expect(value).to.equal(true);
      });
    });
  });
});
