const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("Orbyc Suite", function () {
  let suite;
  let owner, addr1, addr2, addr3;

  let adminRole, editorRole, orbycAgent, nullAgent;

  const agent1 = 1001;
  const agent2 = 1002;
  const agent3 = 1003;

  beforeEach(async () => {
    const Suite = await ethers.getContractFactory("Suite");
    suite = await Suite.deploy();
    await suite.deployed();

    [owner, addr1, addr2, addr3] = await ethers.getSigners();

    adminRole = await suite.ADMIN_ROLE();
    editorRole = await suite.EDITOR_ROLE();
    orbycAgent = await suite.ORBYC_AGENT();
    nullAgent = await suite.NULL_AGENT();
  });

  it("initializes the smart contract", async () => {
    const agent = await suite.idOf(owner.address);
    expect(agent).to.equal(orbycAgent);

    const isAdmin = await suite.hasRole(agent, adminRole);
    const isEditor = await suite.hasRole(agent, editorRole);
    expect(isAdmin && isEditor).to.equal(true);
  });

  describe("role permissions", () => {
    const role1 = 100;

    beforeEach(async () => {
      await suite.defineAgent(addr1.address, agent1, "");
      await suite.defineAgent(addr2.address, agent2, "");

      await suite.grantRole(agent1, adminRole);
      await suite.grantRole(agent2, editorRole);
    });

    describe("agent and role registration", function () {
      it("should succeed if agent1 with role admin try to register a role", async () => {
        await suite.connect(addr1).defineRole(role1, "defined");
        expect(await suite.roleInfo(role1)).to.equal("defined");
      });

      it("should fail if agent1 with role admin try to register an agent", async () => {
        await expect(
          suite.connect(addr1).defineAgent(addr3.address, agent3, "")
        ).to.be.revertedWith("Error: agent has not the required role");
      });

      it("should succeed if agent2 with role editor try to register an agent", async () => {
        await suite.connect(addr2).defineAgent(addr3.address, agent3, "defined");
        expect(await suite.agentInfo(agent3)).to.equal("defined");
      });

      it("should fails if agent2 with role editor try to register a role", async () => {
        await expect(suite.connect(addr2).defineRole(role1, "")).to.be.revertedWith(
          "Error: agent has not the required role"
        );
      });
    });

    describe("role grant and revocation", function () {
      beforeEach(async () => {
        await suite.defineAgent(addr3.address, agent3, "");
      });

      it("should fail if agent2 with role editor try to grant or revoke role", async () => {
        await expect(suite.connect(addr2).grantRole(agent3, adminRole)).to.be.revertedWith(
          "Error: agent has not the required role"
        );

        await expect(suite.connect(addr2).revokeRole(agent3, adminRole)).to.be.revertedWith(
          "Error: agent has not the required role"
        );
      });

      it("should succeed if agent1 with role admin try to grant or revoke role", async () => {
        await suite.connect(addr1).grantRole(agent3, adminRole);
        await suite.connect(addr1).revokeRole(agent2, editorRole);

        expect(await suite.hasRole(agent3, adminRole)).to.equal(true);
        expect(await suite.hasRole(agent2, editorRole)).to.equal(false);
      });
    });
  });

  describe("banning addresses", () => {
    beforeEach(async () => {
      await suite.defineAgent(addr1.address, agent1, "");
      await suite.defineAgent(addr2.address, agent2, "");

      await suite.grantRole(agent1, adminRole);
    });

    describe("when agent has not admin role", async () => {
      it("should fail if try to bann an address", async () => {
        await expect(suite.connect(addr2).bannAddress(addr1.address)).to.be.revertedWith(
          "Error: agent has not the required role"
        );
      });
    });

    describe("when agent has admin and editor role", async () => {
      beforeEach(async () => {
        await suite.bannAddress(addr1.address);
      });

      it("should associate address1 to nullAgent", async () => {
        expect(await suite.idOf(addr1.address)).to.equal(nullAgent);
      });

      it("should fail when defining another agent with address1", async () => {
        await expect(suite.defineAgent(addr1.address, agent1, "")).to.be.revertedWith(
          "Error: agent is null address, can not operate"
        );
      });

      it("should fail when grant roles to nullAddress", async () => {
        await expect(suite.grantRole(nullAgent, adminRole)).to.be.revertedWith(
          "Error: can not operate over null agent"
        );
      });

      it("should fail when revoke roles from nullAddress", async () => {
        await expect(suite.revokeRole(nullAgent, adminRole)).to.be.revertedWith(
          "Error: can not operate over null agent"
        );
      });
    });
  });
});
