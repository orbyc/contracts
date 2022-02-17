// const { expect } = require("chai");
// const { ethers } = require("hardhat");

// describe("Orbyc SupplyChain ERC245", function () {
//   let supplyChain;
//   let owner, addr1, addr2, addr3;

//   let adminRole,
//     editorRole,
//     assetIssuerRole,
//     certIssuerRole,
//     movementIssuerRole,
//     orbycAgent,
//     nullAgent;

//   const agent1 = 1001;
//   const agent2 = 1002;
//   const agent3 = 1003;

//   beforeEach(async () => {
//     const Array = await ethers.getContractFactory("Array");
//     const array = await Array.deploy();
//     await array.deployed();

//     const SupplyChain = await ethers.getContractFactory("SupplyChain", {
//       libraries: {
//         Array: array.address,
//       },
//     });
//     supplyChain = await SupplyChain.deploy("Orbyc Supply Chain");
//     await supplyChain.deployed();

//     [owner, addr1, addr2, addr3] = await ethers.getSigners();

//     adminRole = await supplyChain.ADMIN_ROLE();
//     editorRole = await supplyChain.EDITOR_ROLE();
//     assetIssuerRole = await supplyChain.ASSET_ISSUER_ROLE();
//     certIssuerRole = await supplyChain.CERT_ISSUER_ROLE();
//     movementIssuerRole = await supplyChain.MOVEMENT_ISSUER_ROLE();
//     orbycAgent = await supplyChain.ORBYC_AGENT();
//     nullAgent = await supplyChain.NULL_AGENT();

//     await supplyChain.defineAgent(addr1.address, agent1, "");
//     await supplyChain.defineAgent(addr2.address, agent2, "");
//     await supplyChain.defineAgent(addr3.address, agent3, "");
//   });

//   it("initializes the smart contract", async () => {
//     const agent = await supplyChain.idOf(owner.address);
//     expect(agent).to.equal(orbycAgent);

//     const isAdmin = await supplyChain.hasRole(agent, adminRole);
//     const isEditor = await supplyChain.hasRole(agent, editorRole);
//     const isAssetIssuer = await supplyChain.hasRole(agent, assetIssuerRole);
//     const isCertIssuer = await supplyChain.hasRole(agent, certIssuerRole);
//     const isMovementIssuer = await supplyChain.hasRole(agent, movementIssuerRole);
//     expect(isAdmin && isEditor && isAssetIssuer && isCertIssuer && isMovementIssuer).to.equal(true);
//   });

//   describe("certificate issuing", () => {
//     const cert = {
//       id: 1,
//       metadata: JSON.stringify({ data: "" }),
//     };

//     it("should fail if issuer has not certificate issue role", async () => {
//       await expect(
//         supplyChain.connect(addr1).issueCertificate(cert.id, cert.metadata)
//       ).to.be.revertedWith("Error: agent has not the required role");
//     });

//     it("should succeed if issuer has correct role", async () => {
//       await supplyChain.grantRole(agent1, certIssuerRole);

//       await supplyChain.connect(addr1).issueCertificate(cert.id, cert.metadata);

//       const [issuer, _] = await supplyChain.certificateInfo(cert.id);
//       expect(issuer).to.equal(agent1);
//     });
//   });

//   describe("assets issuing", () => {
//     const asset = {
//       id: 1,
//       owner: 1,
//       co2: 5000,
//       cert: 1,
//       metadata: JSON.stringify({ data: "" }),
//     };

//     it("should fail if issuer has not asset issue role", async () => {
//       await expect(
//         supplyChain
//           .connect(addr1)
//           .issueAsset(asset.id, asset.owner, asset.co2, asset.cert, asset.metadata)
//       ).to.be.revertedWith("Error: agent has not the required role");
//     });

//     it("should succeed if issuer has correct role", async () => {
//       await supplyChain.grantRole(agent1, assetIssuerRole);

//       await supplyChain.issueCertificate(asset.cert, "");
//       await supplyChain
//         .connect(addr1)
//         .issueAsset(asset.id, asset.owner, asset.co2, asset.cert, asset.metadata);

//       const [_, issuer, __, ___, ____] = await supplyChain.assetInfo(asset.id);
//       expect(issuer).to.equal(agent1);
//     });
//   });

//   describe("movement issuing", () => {
//     const move = {
//       id: 1,
//       lat: "80000",
//       lng: "-4000",
//       co2: 500,
//       cert: 100,
//       metadata: JSON.stringify({ data: "" }),
//     };

//     it("should fail if issuer has not certificate issue role", async () => {
//       await expect(
//         supplyChain
//           .connect(addr1)
//           .issueMovement(move.id, move.lat, move.lng, move.co2, move.cert, move.metadata)
//       ).to.be.revertedWith("Error: agent has not the required role");
//     });

//     it("should succeed if issuer has correct role", async () => {
//       await supplyChain.grantRole(agent1, movementIssuerRole);

//       await supplyChain.issueCertificate(move.cert, "");
//       await supplyChain
//         .connect(addr1)
//         .issueMovement(move.id, move.lat, move.lng, move.co2, move.cert, move.metadata);

//       const [issuer, _, __, ___, ____, _____] = await supplyChain.movementInfo(move.id);
//       expect(issuer).to.equal(agent1);
//     });
//   });

//   describe("asset manipulation", () => {
//     const asset = {
//       id: 1,
//       owner: agent2,
//       co2: 5000,
//       cert: 1,
//       metadata: JSON.stringify({ data: "" }),
//     };

//     beforeEach(async () => {
//       await supplyChain.grantRole(agent1, assetIssuerRole);
//       await supplyChain.grantRole(agent2, assetIssuerRole);

//       await supplyChain.issueCertificate(asset.cert, "");
//       await supplyChain
//         .connect(addr1)
//         .issueAsset(asset.id, asset.owner, asset.co2, asset.cert, asset.metadata);
//     });

//     describe("transfer ownership", () => {
//       it("should fail if neither the owner or the issuer transferOwnership", async () => {
//         await expect(supplyChain.transferOwnership(asset.id, agent3)).to.be.revertedWith(
//           "Error: signer not valid for this asset"
//         );
//       });

//       it("should succeed when correct signers", async () => {
//         await supplyChain.connect(addr1).transferOwnership(asset.id, agent3);

//         // transfer signed by the issuer: should not change data
//         let [owner, issuer, _, __, ___] = await supplyChain.assetInfo(asset.id);
//         expect(issuer).to.equal(agent1);
//         expect(owner).to.equal(agent2);

//         await supplyChain.connect(addr2).transferOwnership(asset.id, agent3);

//         // transfer signed by the owner: should update data
//         [owner, issuer, _, __, ___] = await supplyChain.assetInfo(asset.id);
//         expect(issuer).to.equal(agent1);
//         expect(owner).to.equal(agent3);
//       });
//     });

//     describe("adding traceability", () => {
//       const move = {
//         id: 1,
//         lat: "80000",
//         lng: "-40000",
//         co2: asset.co2,
//         cert: asset.cert,
//         metadata: asset.metadata,
//       };

//       beforeEach(async () => {
//         await supplyChain.issueMovement(
//           move.id,
//           move.lat,
//           move.lng,
//           move.co2,
//           move.cert,
//           move.metadata
//         );
//       });

//       it("should fail adding traceability to asset using different issuer", async () => {
//         await expect(
//           supplyChain.connect(addr2).addMovements(asset.id, [move.id])
//         ).to.be.revertedWith("Error: agent is not issuer");
//       });

//       it("should succeed adding traceability to asset using it's issuer", async () => {
//         await supplyChain.connect(addr1).addMovements(asset.id, [move.id]);

//         const [[id], [lat], [lng]] = await supplyChain.assetTraceability(asset.id);

//         expect(id).to.equal(move.id);
//         expect(lat).to.equal(move.lat);
//         expect(lng).to.equal(move.lng);
//       });
//     });

//     describe("adding composition", () => {
//       const parent = {
//         id: 2,
//         owner: 1,
//         co2: 5000,
//         cert: asset.cert,
//         metadata: JSON.stringify({ data: "" }),
//       };

//       beforeEach(async () => {
//         await supplyChain
//           .connect(addr1)
//           .issueAsset(parent.id, parent.owner, parent.co2, parent.cert, parent.metadata);
//       });

//       it("should fail adding traceability to asset using different issuer", async () => {
//         await expect(
//           supplyChain.connect(addr2).addParents(asset.id, [parent.id], [100])
//         ).to.be.revertedWith("Error: agent is not issuer");
//       });

//       it("should succeed adding traceability to asset using it's issuer", async () => {
//         await supplyChain.connect(addr1).addParents(asset.id, [parent.id], [100]);

//         const [[id], [value]] = await supplyChain.assetComposition(asset.id);

//         expect(id).to.equal(parent.id);
//         expect(value).to.equal(100);
//       });
//     });
//   });
// });
