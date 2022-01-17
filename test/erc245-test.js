const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("ERC245", function () {
  let supply,
    issuer = 1;
  let owner;

  const supplyName = "Test ERC245 Supply Chain";

  beforeEach(async function () {
    const Array = await ethers.getContractFactory("Array");
    const array = await Array.deploy();
    await array.deployed();

    const ERC245 = await ethers.getContractFactory("ERC245Private", {
      libraries: {
        Array: array.address,
      },
    });
    supply = await ERC245.deploy(supplyName);
    await supply.deployed();

    [owner, _, _, _] = await ethers.getSigners();
  });

  it("initializes the smart contract", async () => {
    const name = await supply.name();
    expect(name).to.equal(supplyName);
  });

  it("returns the id of a sender", async () => {
    const id = await supply.idOf(owner.address);
    expect(id).to.equal(issuer);
  });

  describe("certificate issuing", () => {
    const cert = {
      id: 100,
      metadata: JSON.stringify({ issuer: 100 }),
    };

    beforeEach(async () => {
      await supply.issueCertificate(cert.id, cert.metadata);
    });

    it("returns the certificate info correctly", async () => {
      const [id, metadata] = await supply.certificateInfo(cert.id);
      expect(id).to.equal(issuer);
      expect(metadata).to.equal(cert.metadata);
    });

    it("should fails issuing the same certificate", async () => {
      await expect(supply.issueCertificate(cert.id, cert.metadata)).to.be.revertedWith(
        "Error: certificate already exists"
      );
    });
  });

  describe("movement issuing", () => {
    const move = {
      id: 1010,
      lat: "80000",
      lng: "-4000",
      co2: 500,
      cert: 100,
      metadata: JSON.stringify({ issuer: 100 }),
    };

    beforeEach(async () => {
      await supply.issueMovement(move.id, move.lat, move.lng, move.co2, move.cert, move.metadata);
    });

    it("returns the movement info correctly", async () => {
      const [id, lat, lng, co2, cert, metadata] = await supply.movementInfo(move.id);
      expect(id).to.equal(issuer);
      expect(lat).to.equal(move.lat);
      expect(lng).to.equal(move.lng);
      expect(co2).to.equal(move.co2);
      expect(cert).to.equal(move.cert);
      expect(metadata).to.equal(move.metadata);
    });

    it("should fail issuing the same movement", async () => {
      await expect(
        supply.issueMovement(move.id, move.lat, move.lng, move.co2, move.cert, move.metadata)
      ).to.be.revertedWith("Error: movement already exists");
    });
  });

  describe("asset issuing", () => {
    const asset = {
      id: 1010,
      owner: 200,
      co2: 500,
      cert: 100,
      metadata: JSON.stringify({ issuer: 100 }),
    };

    beforeEach(async () => {
      await supply.issueAsset(asset.id, asset.owner, asset.co2, asset.cert, asset.metadata);
    });

    it("returns the asset info correctly", async () => {
      const [owner, id, co2, cert, metadata] = await supply.assetInfo(asset.id);
      expect(owner).to.equal(asset.owner);
      expect(id).to.equal(issuer);
      expect(co2).to.equal(asset.co2);
      expect(cert).to.equal(asset.cert);
      expect(metadata).to.equal(asset.metadata);
    });

    it("should fails issuing the same asset", async () => {
      await expect(
        supply.issueAsset(asset.id, asset.owner, asset.co2, asset.cert, asset.metadata)
      ).to.be.revertedWith("Error: asset already exists");
    });
  });

  describe("asset operations", () => {
    const asset = {
      id: 1010,
      owner: 200,
      co2: 500,
      cert: 100,
      metadata: JSON.stringify({ issuer: 100 }),
    };

    beforeEach(async () => {
      await supply.issueAsset(asset.id, asset.owner, asset.co2, asset.cert, asset.metadata);
    });

    describe("transfer ownership of asset", () => {
      it("should change the owner of the asset", async () => {
        await supply.transferOwnership(asset.id, asset.owner + 1);

        const [owner, _, __, ___, ____] = await supply.assetInfo(asset.id);
        expect(owner).to.equal(asset.owner + 1);
      });
    });

    describe("add movement to asset", () => {
      const move = {
        id: 1010,
        lat: "80000",
        lng: "-4000",
        co2: 500,
        cert: 100,
        metadata: JSON.stringify({ issuer: 100 }),
      };

      beforeEach(async () => {
        await supply.issueMovement(move.id, move.lat, move.lng, move.co2, move.cert, move.metadata);
        await supply.addMovements(asset.id, [move.id]);
      });

      it("should affect asset co2 emission", async () => {
        const [_, __, co2, ___, ____] = await supply.assetInfo(asset.id);
        expect(co2).to.equal(asset.co2 + move.co2);
      });

      it("should be present in asset traceability", async () => {
        const [[id], [lat], [lng]] = await supply.assetTraceability(asset.id);
        expect(id).to.equal(move.id);
        expect(lat).to.equal(move.lat);
        expect(lng).to.equal(move.lng);
      });

      it("should fails if same movement is added to the same asset", async () => {
        await expect(supply.addMovements(asset.id, [move.id])).to.be.revertedWith(
          "Error: movement already added"
        );
      });

      it("should fail when added inexistent movement", async () => {
        await expect(supply.addMovements(asset.id, [2002])).to.be.revertedWith(
          "Error: can not add inexistent movement"
        );
      });
    });

    describe("add composition to asset", () => {
      const parent = {
        id: 1012,
        owner: 200,
        co2: 500,
        cert: 100,
        metadata: JSON.stringify({ issuer: 100 }),
      };

      beforeEach(async () => {
        await supply.issueAsset(parent.id, parent.owner, parent.co2, parent.cert, parent.metadata);
        await supply.addParents(asset.id, [parent.id], [100]);
      });

      it("should affect asset co2 emission", async () => {
        const [_, __, co2, ___, ____] = await supply.assetInfo(asset.id);
        expect(co2).to.equal(asset.co2 + parent.co2);
      });

      it("should be present in asset traceability", async () => {
        const [[id], [comp]] = await supply.assetComposition(asset.id);
        expect(id).to.equal(parent.id);
        expect(comp).to.equal(100);
      });

      it("should fail if same parent is added to same asset", async () => {
        await expect(supply.addParents(asset.id, [parent.id], [100])).to.be.revertedWith(
          "Error: parent already present"
        );
      });

      it("should fail when added inexistent parent", async () => {
        await expect(supply.addParents(asset.id, [2002], [100])).to.be.revertedWith(
          "Error: can not add inexistent asset"
        );
      });
    });

    describe("alteration to children", () => {
      const parent = {
        id: 1012,
        owner: 202,
        co2: 502,
        cert: 100,
        metadata: JSON.stringify({ issuer: 100 }),
      };

      beforeEach(async () => {
        await supply.issueAsset(parent.id, parent.owner, parent.co2, parent.cert, parent.metadata);
        await supply.addParents(asset.id, [parent.id], [100]);
      });

      describe("when movement added to parent", () => {
        const move = {
          id: 1010,
          lat: "80000",
          lng: "-4000",
          co2: 500,
          cert: 100,
          metadata: JSON.stringify({ issuer: 100 }),
        };

        beforeEach(async () => {
          await supply.issueMovement(
            move.id,
            move.lat,
            move.lng,
            move.co2,
            move.cert,
            move.metadata
          );
          await supply.addMovements(parent.id, [move.id]);
        });

        it("should affect parent co2 emission", async () => {
          const [_, __, co2, ___, ____] = await supply.assetInfo(parent.id);
          expect(co2).to.equal(move.co2 + parent.co2);
        });

        it("should affect asset co2 emission", async () => {
          const [_, __, co2, ___, ____] = await supply.assetInfo(asset.id);
          expect(co2).to.equal(move.co2 + parent.co2 + asset.co2);
        });
      });

      describe("when composition added to parent", () => {
        const grand = {
          id: 1014,
          owner: 200,
          co2: 500,
          cert: 100,
          metadata: JSON.stringify({ issuer: 100 }),
        };

        beforeEach(async () => {
          await supply.issueAsset(grand.id, grand.owner, grand.co2, grand.cert, grand.metadata);
          await supply.addParents(parent.id, [grand.id], [100]);
        });

        it("should affect parent co2 emission", async () => {
          const [_, __, co2, ___, ____] = await supply.assetInfo(parent.id);
          expect(co2).to.equal(grand.co2 + parent.co2);
        });

        it("should affect asset co2 emission", async () => {
          const [_, __, co2, ___, ____] = await supply.assetInfo(asset.id);
          expect(co2).to.equal(grand.co2 + parent.co2 + asset.co2);
        });
      });
    });
  });
});
