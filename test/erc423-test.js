const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("ERC423", function () {
  let suite;
  let owner, addr1, addr2;

  const suiteName = "Test ERC423 Suite"

  beforeEach(async function () {
    const ERC423 = await ethers.getContractFactory("ERC423");
    suite = await ERC423.deploy(suiteName);
    await suite.deployed();

    [owner, addr1, addr2, _] = await ethers.getSigners();
  });

  it("initializes the smart contract", async () => {
    const name = await suite.name();
    expect(name).to.equal(suiteName);
  })

  describe("define agent", function () {
    const agentId = 5;
    const agentInfo = JSON.stringify({ name: "Agent 1", value: 50 });

    beforeEach(async () => {
      await suite.defineAgent(addr1.address, agentId, agentInfo);
    });

    it("should have associate id with agent's address", async () => {
      id = await suite.idOf(addr1.address);

      expect(id).to.equal(agentId);
    });

    it("should return info for the given id", async () => {
      info = await suite.agentInfo(agentId);

      expect(info).to.equal(agentInfo);
    });
  });

  describe("define role", function () {
    const roleId = 1;
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
    const role1 = 1;
    const role2 = 2;

    const agent = 1;

    beforeEach(async () => {
      await suite.defineRole(role1, "");
      await suite.defineRole(role2, "");

      await suite.defineAgent(addr1.address, agent, "");
    });

    describe("when only one role is endorsed", async () => {
      beforeEach(async () => {
        await suite.grantRole(agent, role1);
      });

      it("should return true when asks for agent to have the role 1", async () => {
        const value = await suite.hasRole(agent, role1);

        expect(value).to.equal(true);
      });

      it("should return false when asks for agent to have the role 2", async () => {
        const value = await suite.hasRole(agent, role2);

        expect(value).to.equal(false);
      });
    });

    describe("when more than one role is endorsed", async () => {
      beforeEach(async () => {
        await suite.grantRole(agent, role1);
        await suite.grantRole(agent, role2);
      });

      it("should return true when asks for agent to have the role 1", async () => {
        const value = await suite.hasRole(agent, role1);

        expect(value).to.equal(true);
      });

      it("should return true when asks for agent to have the role 2", async () => {
        const value = await suite.hasRole(agent, role2);

        expect(value).to.equal(true);
      });
    });
  });

  describe("role revocation", function () {
    const role1 = 1;
    const role2 = 2;

    const agent = 1;

    beforeEach(async () => {
      await suite.defineRole(role1, "");
      await suite.defineRole(role2, "");

      await suite.defineAgent(addr1.address, agent, "");

      await suite.grantRole(agent, role1);
      await suite.grantRole(agent, role2);
    });

    it("should return true when ask for endorsed roles", async () => {
      const hasRole1 = await suite.hasRole(agent, role1);
      const hasRole2 = await suite.hasRole(agent, role2);

      expect(hasRole1 && hasRole2).to.equal(true);
    });

    describe("when user role 1 is revoked", async () => {
      beforeEach(async () => {
        await suite.revokeRole(agent, role1);
      });

      it("should return false when asks for agent to have the role 1", async () => {
        const value = await suite.hasRole(agent, role1);

        expect(value).to.equal(false);
      });

      it("should return true when asks for agent to have the role 2", async () => {
        const value = await suite.hasRole(agent, role2);

        expect(value).to.equal(true);
      });
    });
  });
});
