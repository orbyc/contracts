const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("ERC245", function () {
  let supply;
  let owner, addr1;

  const supplyName = "Test ERC245 Supply Chain";

  beforeEach(async function () {
    const Array = await ethers.getContractFactory("Array");
    const array = await Array.deploy();
    await array.deployed();

    const ERC245 = await ethers.getContractFactory("ERC245", {
      libraries: {
        Array: array.address,
      },
    });
    supply = await ERC245.deploy(supplyName);
    await supply.deployed();

    [owner, addr1, _, _] = await ethers.getSigners();
  });

  it("initializes the smart contract", async () => {
    const name = await supply.name();
    expect(name).to.equal(supplyName);
  });

  describe("certificate issuing", () => {
    const cert = {
      certId: 100,
      metadata: JSON.stringify({ issuer: 100 }),
    };

    beforeEach(async () => {
      await supply.issueCertificate(cert.certId, cert.metadata);
    });

    it("returns the certificate info correctly", async () => {
      const [issuer, metadata] = await supply.certificateInfo(cert.certId);
      expect(issuer).to.equal(owner.address);
      expect(metadata).to.equal(cert.metadata);
    });

    it("should fails issuing the same certificate", async () => {
      await expect(supply.issueCertificate(cert.certId, cert.metadata)).to.be.revertedWith(
        "Error: certificate already exists"
      );
    });
  });

  describe("movement issuing", () => {
    const move = {
      moveId: 1010,
      lat: "80000",
      lng: "-4000",
      co2e: 500,
      certId: 100,
      metadata: JSON.stringify({ issuer: 100 }),
    };

    beforeEach(async () => {
      await supply.issueCertificate(move.certId, "");
      await supply.issueMovement(
        move.moveId,
        move.lat,
        move.lng,
        move.co2e,
        move.certId,
        move.metadata
      );
    });

    it("returns the movement info correctly", async () => {
      const [issuer, lat, lng, co2e, certId, metadata] = await supply.movementInfo(move.moveId);
      expect(issuer).to.equal(owner.address);
      expect(lat).to.equal(move.lat);
      expect(lng).to.equal(move.lng);
      expect(co2e).to.equal(move.co2e);
      expect(certId).to.equal(move.certId);
      expect(metadata).to.equal(move.metadata);
    });

    it("should fail issuing the same movement", async () => {
      await expect(
        supply.issueMovement(move.moveId, move.lat, move.lng, move.co2e, move.certId, move.metadata)
      ).to.be.revertedWith("Error: movement already exists");
    });

    it("should fail if the certificate has not been issued", async () => {
      await expect(
        supply.issueMovement(
          move.moveId + 1,
          move.lat,
          move.lng,
          move.co2e,
          move.certId + 1,
          move.metadata
        )
      ).to.be.revertedWith("Error: certificate does not exists");
    });
  });

  describe("asset issuing", () => {
    const asset = {
      assetId: 1010,
      owner: "0xA001da7DcdF4BD8463823669FB9c43Fcb807EA61",
      co2e: 500,
      certId: 100,
      metadata: JSON.stringify({ issuer: 100 }),
    };

    beforeEach(async () => {
      await supply.issueCertificate(asset.certId, "");
      await supply.issueAsset(asset.assetId, asset.owner, asset.co2e, asset.certId, asset.metadata);
    });

    it("returns the asset info correctly", async () => {
      const [_owner, _issuer, co2e, certId, metadata] = await supply.assetInfo(asset.assetId);
      expect(_owner).to.equal(asset.owner);
      expect(_issuer).to.equal(owner.address);
      expect(co2e).to.equal(asset.co2e);
      expect(certId).to.equal(asset.certId);
      expect(metadata).to.equal(asset.metadata);
    });

    it("should fails issuing the same asset", async () => {
      await expect(
        supply.issueAsset(asset.assetId, asset.owner, asset.co2e, asset.certId, asset.metadata)
      ).to.be.revertedWith("Error: asset already exists");
    });

    it("should fail if the certificate has not been issued", async () => {
      await expect(
        supply.issueAsset(
          asset.assetId + 1,
          asset.owner,
          asset.co2e,
          asset.certId + 1,
          asset.metadata
        )
      ).to.be.revertedWith("Error: certificate does not exists");
    });
  });

  describe("asset operations", () => {
    const asset = {
      assetId: 1000,
      owner: "0xA271251D16B1e5224164663ADCE5e60e028595dC",
      co2e: 500,
      certId: 100,
      metadata: JSON.stringify({ issuer: 100 }),
    };

    beforeEach(async () => {
      asset.owner = owner.address;
      await supply.issueCertificate(asset.certId, "");
      await supply.issueAsset(asset.assetId, asset.owner, asset.co2e, asset.certId, asset.metadata);
    });

    describe("transfers", () => {
      it("should change the owner when transfer is made", async () => {
        await supply.transferFrom(owner.address, addr1.address, asset.assetId);

        const [ownerAddr, issuerAddr] = await supply.assetInfo(asset.assetId);

        expect(ownerAddr).to.equal(addr1.address);
        expect(issuerAddr).to.equal(owner.address);
      });
    });

    describe("add movement to asset", () => {
      const move = {
        moveId: 2000,
        lat: "80000",
        lng: "-4000",
        co2e: 200,
        certId: 100,
        metadata: JSON.stringify({ issuer: 100 }),
      };

      beforeEach(async () => {
        await supply.issueMovement(
          move.moveId,
          move.lat,
          move.lng,
          move.co2e,
          move.certId,
          move.metadata
        );
        await supply.addMovements(asset.assetId, [move.moveId]);
      });

      it("should not affect asset co2 emission", async () => {
        const [_, __, co2e, ___, ____] = await supply.assetInfo(asset.assetId);
        expect(co2e).to.equal(asset.co2e);
      });

      it("should be present in asset traceability", async () => {
        const [[moveId], [lat], [lng]] = await supply.assetTraceability(asset.assetId);
        expect(moveId).to.equal(move.moveId);
        expect(lat).to.equal(move.lat);
        expect(lng).to.equal(move.lng);
      });

      it("should fails if same movement is added to the same asset", async () => {
        await expect(supply.addMovements(asset.assetId, [move.moveId])).to.be.revertedWith(
          "Error: movement already added"
        );
      });

      it("should fail when added inexistent movement", async () => {
        await expect(supply.addMovements(asset.assetId, [2002])).to.be.revertedWith(
          "Error: can not add inexistent movement"
        );
      });
    });

    describe("add composition to asset", () => {
      const parent = {
        assetId: 1012,
        owner: "0xA355538257b8716f0f9c0f30F94cE82f5cd44a8b",
        co2e: 1000,
        certId: 100,
        metadata: JSON.stringify({ issuer: 100 }),
      };

      beforeEach(async () => {
        await supply.issueAsset(
          parent.assetId,
          parent.owner,
          parent.co2e,
          parent.certId,
          parent.metadata
        );
        await supply.addParents(asset.assetId, [parent.assetId], [100]);
      });

      it("should not affect asset co2 emission", async () => {
        const [_, __, co2e, ___, ____] = await supply.assetInfo(asset.assetId);
        expect(co2e).to.equal(asset.co2e);
      });

      it("should be present in asset traceability", async () => {
        const [[assetId], [percent]] = await supply.assetComposition(asset.assetId);
        expect(assetId).to.equal(parent.assetId);
        expect(percent).to.equal(100);
      });

      it("should fail if same parent is added to same asset", async () => {
        await expect(supply.addParents(asset.assetId, [parent.assetId], [100])).to.be.revertedWith(
          "Error: parent already present"
        );
      });

      it("should fail when added inexistent parent", async () => {
        await expect(supply.addParents(asset.assetId, [2002], [100])).to.be.revertedWith(
          "Error: can not add inexistent asset"
        );
      });
    });

    describe("alteration to children", () => {
      const parent = {
        assetId: 1012,
        owner: "0xA495ac289C5Ff8e9F665595653F295c9c9c4cacc",
        co2e: 1000,
        certId: 100,
        metadata: JSON.stringify({ issuer: 100 }),
      };

      beforeEach(async () => {
        await supply.issueAsset(
          parent.assetId,
          parent.owner,
          parent.co2e,
          parent.certId,
          parent.metadata
        );
        await supply.addParents(asset.assetId, [parent.assetId], [100]);
      });

      describe("when movement added to parent", () => {
        const move = {
          moveId: 1010,
          lat: "80000",
          lng: "-4000",
          co2e: 200,
          certId: 100,
          metadata: JSON.stringify({ issuer: 100 }),
        };

        beforeEach(async () => {
          await supply.issueMovement(
            move.moveId,
            move.lat,
            move.lng,
            move.co2e,
            move.certId,
            move.metadata
          );
          await supply.addMovements(parent.assetId, [move.moveId]);
        });

        it("should not affect parent co2 emission", async () => {
          const [_, __, co2e, ___, ____] = await supply.assetInfo(parent.assetId);
          expect(co2e).to.equal(parent.co2e);
        });

        it("should not affect asset co2 emission", async () => {
          const [_, __, co2e, ___, ____] = await supply.assetInfo(asset.assetId);
          expect(co2e).to.equal(asset.co2e);
        });
      });

      describe("when composition added to parent", () => {
        const grand = {
          assetId: 1014,
          owner: "0xA495ac289C5Ff8e9F665595653F295c9c9c4cacc",
          co2e: 1500,
          certId: 100,
          metadata: JSON.stringify({ issuer: 100 }),
        };

        beforeEach(async () => {
          await supply.issueAsset(
            grand.assetId,
            grand.owner,
            grand.co2e,
            grand.certId,
            grand.metadata
          );
          await supply.addParents(parent.assetId, [grand.assetId], [100]);
        });

        it("should not affect parent co2 emission", async () => {
          const [_, __, co2e, ___, ____] = await supply.assetInfo(parent.assetId);
          expect(co2e).to.equal(parent.co2e);
        });

        it("should not affect asset co2 emission", async () => {
          const [_, __, co2e, ___, ____] = await supply.assetInfo(asset.assetId);
          expect(co2e).to.equal(asset.co2e);
        });
      });
    });
  });
});
