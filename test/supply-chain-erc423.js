const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("Orbyc SupplyChain ERC423", function () {
  let supplyChain;
  let owner, addr1, addr2, addr3;

  let adminRole, providerRole, orbycAccount;

  const account1 = "0x122733e2e7c995732f444f7fD77E0715D6448E09";
  const account2 = "0x12fbeFEa020E81B0ECb8ed92eEBC5276B1D25a53";
  const account3 = "0x12345982421011440aAC0c42E9A64490a378a316";

  beforeEach(async () => {
    const Array = await ethers.getContractFactory("Array");
    const array = await Array.deploy();
    await array.deployed();

    const SupplyChain = await ethers.getContractFactory("SupplyChain", {
      libraries: {
        Array: array.address,
      },
    });
    supplyChain = await SupplyChain.deploy("Orbyc Supply Chain");
    await supplyChain.deployed();

    [owner, addr1, addr2, addr3] = await ethers.getSigners();

    adminRole = await supplyChain.ADMIN_ROLE();
    providerRole = await supplyChain.PROVIDER_ROLE();
    orbycAccount = await supplyChain.ORBYC_ACCOUNT();
  });

  it("initializes the smart contract", async () => {
    const agent = await supplyChain.accountOf(owner.address);
    expect(agent).to.equal(orbycAccount);

    const isAdmin = await supplyChain.hasRole(agent, adminRole);
    expect(isAdmin).to.equal(true);
  });

  describe("role permissions", () => {
    const role1 = 1 << 10;

    beforeEach(async () => {
      await supplyChain.defineAgent(addr1.address, account1, "");
      await supplyChain.defineAgent(addr2.address, account2, "");

      await supplyChain.grantRole(account1, adminRole);
    });

    describe("agent and role registration", function () {
      it("should succeed if account 1 with role admin tries to register a role", async () => {
        await supplyChain.connect(addr1).defineRole(role1, "defined");
        expect(await supplyChain.roleInfo(role1)).to.equal("defined");
      });

      it("should succeed if account 1 with role admin tries to register an agent", async () => {
        await supplyChain.connect(addr1).defineAgent(addr3.address, account3, "defined");
        expect(await supplyChain.accountInfo(account3)).to.equal("defined");
      });

      it("should fail if account 2 with no role tries to register an agent", async () => {
        await expect(
          supplyChain.connect(addr2).defineAgent(addr3.address, account3, "")
        ).to.be.revertedWith("Error: agent has not the required role");
      });

      it("should fails if account 2 with no role tries to register a role", async () => {
        await expect(supplyChain.connect(addr2).defineRole(role1, "")).to.be.revertedWith(
          "Error: agent has not the required role"
        );
      });
    });

    describe("role grant and revocation", function () {
      beforeEach(async () => {
        await supplyChain.defineAgent(addr3.address, account3, "");
        await supplyChain.grantRole(account2, providerRole);
      });

      it("should succeed if agent1 with role admin tries to grant or revoke role", async () => {
        await supplyChain.connect(addr1).grantRole(account3, adminRole);
        await supplyChain.connect(addr1).revokeRole(account2, providerRole);

        expect(await supplyChain.hasRole(account3, adminRole)).to.equal(true);
        expect(await supplyChain.hasRole(account2, providerRole)).to.equal(false);
      });

      it("should fail if agent2 with no role tries to grant or revoke role", async () => {
        await expect(supplyChain.connect(addr2).grantRole(account3, adminRole)).to.be.revertedWith(
          "Error: agent has not the required role"
        );

        await expect(supplyChain.connect(addr2).revokeRole(account3, adminRole)).to.be.revertedWith(
          "Error: agent has not the required role"
        );
      });
    });
  });

  describe("remove agents", () => {
    beforeEach(async () => {
      await supplyChain.defineAgent(addr1.address, account1, "");
      await supplyChain.defineAgent(addr2.address, account2, "");

      await supplyChain.grantRole(account1, adminRole);
    });

    describe("when account has not admin role", async () => {
      it("should fail if try to remove agent", async () => {
        await expect(supplyChain.connect(addr2).removeAgent(addr1.address)).to.be.revertedWith(
          "Error: agent has not the required role"
        );
      });
    });

    describe("when account has admin role", async () => {
      beforeEach(async () => {
        await supplyChain.removeAgent(addr1.address);
      });

      it("should fail when trying to get account of agent", async () => {
        await expect(supplyChain.accountOf(addr1.address)).to.be.revertedWith(
          "ERC423: agent has been removed"
        );
      });

      it("should fail when defining another agent with address1", async () => {
        await expect(supplyChain.defineAgent(addr1.address, account1, "")).to.be.revertedWith(
          "ERC423: agent has been removed"
        );
      });
    });
  });
});
