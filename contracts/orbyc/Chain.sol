// SPDX-License-Identifier: MIT
// Orbyc Contracts v1.0.0 (orbyc/Chain.sol)

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/token/common/ERC2981.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

import "closedzeppelin-contracts/access/AccessControl.sol";
import "closedzeppelin-contracts/access/IAccountControl.sol";
import "closedzeppelin-contracts/utils/Multisig.sol";

import "../token/ERC721/extensions/ERC721Certificable.sol";
import "../token/ERC721/extensions/ERC721Compound.sol";
import "../token/ERC721/extensions/ERC721EmbedMetadata.sol";
import "../token/ERC721/extensions/ERC721Traceable.sol";

/**
 * @title OrbycChain
 * @dev The OrbycChain smart contract is a versatile contract that allows for a variety of actions related to non-fungible tokens (NFTs) on the Ethereum
 * blockchain. It includes a range of interfaces and functions that enable NFT ownership tracking, metadata management, and certification, as well as
 * access control and multisig capabilities for enhanced security. Additionally, it supports the ERC2981 royalty standard for NFTs, providing a way for
 * creators to receive royalties on secondary sales of their work. Overall, the OrbycChain contract is a powerful tool for creating and managing NFTs on
 * the Ethereum network.
 */
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
    /**
     * @dev The OrbycChain contract has a single variable that contains an IAccountControl instance, which is used to manage accounts.
     */
    IAccountControl private _accounts;

    /**
     * @dev Roles to manage access control.
     */
    uint64 public constant ASSET_ISSUER_ROLE = 0x1 << 1;
    uint64 public constant CERTIFICATE_ISSUER_ROLE = 0x1 << 2;
    uint64 public constant COMPOSITION_ISSUER_ROLE = 0x1 << 3;
    uint64 public constant TRACEABILITY_ISSUER_ROLE = 0x1 << 4;

    /**
     * @dev Check if a given interface is supported by OrbycChain.
     * @param interfaceId ID of the interface.
     * @return True if the interface is supported by OrbycChain, otherwise false.
     */
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

    /**
     * @dev Modifier that checks if the signature comes from the owner of the token.
     * @param tokenId ID of the token.
     */
    modifier signatureFromOwner(uint256 tokenId) {
        (address owner, ) = royaltyInfo(tokenId, 0);
        require(
            _accounts.accountOf(signers(0)) == owner,
            "Orbyc: signer not owner"
        );
        _;
    }

    /**
     * @dev Constructor function for OrbycChain.
     * @param name_ Name of the NFT contract.
     * @param symbol_ Symbol of the NFT contract.
     * @param accounts_ An instance of the IAccountControl contract.
     */
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

        // _grantRole(DEFAULT_ADMIN_ROLE, _msgSender());
    }

    /**
     * @dev Returns the account address of the message sender.
     * @return The account address of the message sender.
     */
    function _msgSender() internal view override returns (address) {
        return _accounts.accountOf(super._msgSender());
    }

    /**
     * @notice Returns the Uniform Resource Identifier (URI) for a token.
     * @param tokenId The token to fetch the URI for.
     * @return A string representing the URI.
     */
    function tokenURI(uint256 tokenId)
        public
        pure
        override
        returns (string memory)
    {
        return
            string.concat(
                "https://wallet.orbyc.com/metadata/",
                Strings.toString(tokenId)
            );
    }

    /**
     * @notice Mint a new token with the given ID and metadata.
     * @dev Only users with the ASSET_ISSUER_ROLE can mint tokens.
     * @param tokenId The ID of the token to mint.
     * @param data The metadata for the token.
     */
    function mint(uint256 tokenId, string memory data)
        public
        onlyRole(ASSET_ISSUER_ROLE)
    {
        _mint(_msgSender(), tokenId);
        _setTokenMetadata(tokenId, data);
        _setTokenRoyalty(tokenId, _msgSender(), _feeDenominator() / 10);
    }

    /**
     * @notice Appends new certificates to a given token.
     * @dev Only users with the CERTIFICATE_ISSUER_ROLE can append certificates.
     * @param tokenId The ID of the token to append certificates to.
     * @param certificates An array of Certificate structs containing the new certificates to append.
     */
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

    /**
     * @notice Resets the certificates for a given token.
     * @dev Only users with the DEFAULT_ADMIN_ROLE can reset certifications.
     * @param tokenId The ID of the token to reset certifications for.
     */
    function resetCertifications(uint256 tokenId)
        public
        requireSignatures(1)
        signatureFromOwner(tokenId)
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        _resetCertifications(tokenId);
    }

    /**
     * @notice Appends new components to a given token.
     * @dev Only users with the COMPOSITION_ISSUER_ROLE can append components.
     * @param tokenId The ID of the token to append components to.
     * @param components An array of Component structs containing the new components to append.
     */
    function appendComponents(uint256 tokenId, Component[] memory components)
        public
        override
        requireSignatures(1)
        signatureFromOwner(tokenId)
        onlyRole(COMPOSITION_ISSUER_ROLE)
    {
        _appendComponents(tokenId, components);
    }

    /**
     * @notice Resets the components for a given token.
     * @dev Only users with the DEFAULT_ADMIN_ROLE can reset composition.
     * @param tokenId The ID of the token to reset composition for.
     */
    function resetComposition(uint256 tokenId)
        public
        requireSignatures(1)
        signatureFromOwner(tokenId)
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        _resetComposition(tokenId);
    }

    /**
     * @notice Appends new traces to a given token.
     * @dev Only users with the TRACEABILITY_ISSUER_ROLE can append traces.
     * @param tokenId The ID of the token to append traces to.
     * @param traces An array of Trace structs containing the new traces to append.
     */
    function appendTraces(uint256 tokenId, Trace[] memory traces)
        public
        override
        requireSignatures(1)
        signatureFromOwner(tokenId)
        onlyRole(TRACEABILITY_ISSUER_ROLE)
    {
        _appendTraces(tokenId, traces);
    }

    /**
     * @notice Resets the traceability for a given token.
     * @dev Only users with the DEFAULT_ADMIN_ROLE can reset traceability.
     * @param tokenId The ID of the token to reset traceability for.
     */
    function resetTraceability(uint256 tokenId)
        public
        requireSignatures(1)
        signatureFromOwner(tokenId)
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        _resetTraceability(tokenId);
    }

    /**
     * @notice Resets the metadata for a given token.
     * @dev Only users with the DEFAULT_ADMIN_ROLE can reset metadata.
     * @param tokenId The ID of the token to reset metadata for.
     */
    function resetTokenMetadata(uint256 tokenId, string memory metadata)
        public
        requireSignatures(1)
        signatureFromOwner(tokenId)
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        _setTokenMetadata(tokenId, metadata);
    }
}
