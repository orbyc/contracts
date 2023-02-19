// SPDX-License-Identifier: MIT
// Orbyc Contracts v1.0.0 (orbyc/Chain.sol)

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/token/common/ERC2981.sol";

import "closedzeppelin-contracts/access/AccessControl.sol";
import "closedzeppelin-contracts/access/IAccountControl.sol";
import "closedzeppelin-contracts/utils/Multisig.sol";

import "../token/ERC721/extensions/ERC721Certificable.sol";
import "../token/ERC721/extensions/ERC721Compound.sol";
import "../token/ERC721/extensions/ERC721EmbedMetadata.sol";
import "../token/ERC721/extensions/ERC721Traceable.sol";

contract OrbycChain is
    AccessControl,
    ERC721Enumerable,
    ERC721Certificable,
    ERC721Compound,
    ERC721EmbedMetadata,
    ERC721Traceable,
    ERC2981,
    Multisig
{
    IAccountControl private _accounts;

    uint64 public constant ASSET_ISSUER_ROLE = 0x1 << 1;
    uint64 public constant CERTIFICATE_ISSUER_ROLE = 0x1 << 2;
    uint64 public constant COMPOSITION_ISSUER_ROLE = 0x1 << 3;
    uint64 public constant TRACEABILITY_ISSUER_ROLE = 0x1 << 4;

    function supportsInterface(bytes4 interfaceId)
        public
        view
        virtual
        override(AccessControl, ERC721Enumerable, ERC2981, Multisig)
        returns (bool)
    {
        return
            AccessControl.supportsInterface(interfaceId) ||
            ERC721Enumerable.supportsInterface(interfaceId) ||
            ERC2981.supportsInterface(interfaceId) ||
            Multisig.supportsInterface(interfaceId);
    }

    modifier signatureFromOwner(uint256 tokenId) {
        (address owner, ) = royaltyInfo(tokenId, 0);
        require(
            _accounts.accountOf(signers(0)) == owner,
            "Orbyc: signer not owner"
        );
        _;
    }

    constructor(
        string memory name_,
        string memory symbol_,
        IAccountControl accounts_
    ) ERC721(name_, symbol_) Multisig(name_) {
        _accounts = accounts_;

        _setRoleAdmin(ASSET_ISSUER_ROLE, DEFAULT_ADMIN_ROLE);
        _setRoleAdmin(CERTIFICATE_ISSUER_ROLE, DEFAULT_ADMIN_ROLE);
        _setRoleAdmin(COMPOSITION_ISSUER_ROLE, DEFAULT_ADMIN_ROLE);
        _setRoleAdmin(TRACEABILITY_ISSUER_ROLE, DEFAULT_ADMIN_ROLE);

        _grantRole(DEFAULT_ADMIN_ROLE, _msgSender());
    }

    function _msgSender() internal view override returns (address) {
        return _accounts.accountOf(super._msgSender());
    }

    function tokenURI(uint256 tokenId)
        public
        view
        override
        returns (string memory)
    {
        return "https://orbyc.com/nft/" + tokenId;
    }

    function mint(uint256 tokenId, string memory data)
        public
        onlyRole(ASSET_ISSUER_ROLE)
    {
        _mint(_msgSender(), tokenId);
        _setTokenMetadata(tokenId, data);
        _setTokenRoyalty(tokenId, _msgSender(), _feeDenominator());
    }

    function appendCertificates(
        uint256 tokenId,
        Certificate[] memory certificates
    )
        public
        override
        requireSignatures(1)
        signatureFromOwner(tokenId)
        onlyRole(CERTIFICATE_ISSUER_ROLE)
    {
        _appendCertificates(tokenId, certificates);
    }

    function resetCertifications(uint256 tokenId)
        public
        requireSignatures(1)
        signatureFromOwner(tokenId)
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        _resetCertifications(tokenId);
    }

    function appendComponents(uint256 tokenId, Component[] memory components)
        public
        override
        requireSignatures(1)
        signatureFromOwner(tokenId)
        onlyRole(COMPOSITION_ISSUER_ROLE)
    {
        _appendComponents(tokenId, components);
    }

    function resetComposition(uint256 tokenId)
        public
        requireSignatures(1)
        signatureFromOwner(tokenId)
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        _resetComposition(tokenId);
    }

    function appendTraces(uint256 tokenId, Trace[] memory traces)
        public
        override
        requireSignatures(1)
        signatureFromOwner(tokenId)
        onlyRole(TRACEABILITY_ISSUER_ROLE)
    {
        _appendTraces(tokenId, traces);
    }

    function resetTraceability(uint256 tokenId)
        public
        requireSignatures(1)
        signatureFromOwner(tokenId)
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        _resetTraceability(tokenId);
    }

    function resetTokenMetadata(uint256 tokenId, string memory metadata)
        public
        requireSignatures(1)
        signatureFromOwner(tokenId)
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        _setTokenMetadata(tokenId, metadata);
    }
}
