// SPDX-License-Identifier: MIT
// Orbyc Contracts v1.0.0 (orbyc/Chain.sol)

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
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
 * @dev The OrbycChain smart contract is a multi-faceted contract that integrates several standard interfaces, such as AccessControl,
 * ERC1155, ERC2981, and Multisig. It is responsible for managing accounts, issuing assets, and managing access control
 * through various roles. The contract also includes functions for appending certificates and components to tokens and
 * resetting these values, which are limited to users with the appropriate role.
 */
contract OrbycChain is
    AccessControl,
    ERC1155,
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
        override(AccessControl, ERC1155, ERC2981, Multisig)
        returns (bool)
    {
        return
            AccessControl.supportsInterface(interfaceId) ||
            ERC1155.supportsInterface(interfaceId) ||
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
     * @param accounts_ An instance of the IAccountControl contract.
     */
    constructor(IAccountControl accounts_)
        ERC1155("https://wallet.orbyc.com/metadata")
        Multisig("orbyc Chain")
    {
        _accounts = accounts_;

        _setRoleAdmin(ASSET_ISSUER_ROLE, DEFAULT_ADMIN_ROLE);
        _setRoleAdmin(CERTIFICATE_ISSUER_ROLE, DEFAULT_ADMIN_ROLE);
        _setRoleAdmin(COMPOSITION_ISSUER_ROLE, DEFAULT_ADMIN_ROLE);
        _setRoleAdmin(TRACEABILITY_ISSUER_ROLE, DEFAULT_ADMIN_ROLE);

        _grantRole(DEFAULT_ADMIN_ROLE, _msgSender());
    }

    /**
     * @dev Returns the account address of the message sender.
     * @return The account address of the message sender.
     */
    function _msgSender() internal view override returns (address) {
        return _accounts.accountOf(super._msgSender());
    }

    /**
    @dev Mints a new token with the given ID, amount and metadata, and assigns it to the specified address.
    Only the asset issuer role can perform this action.
    @param to The address that will receive the minted tokens.
    @param id The ID of the token to mint.
    @param amount The amount of tokens to mint.
    @param data Additional metadata to be associated with the token.
    */
    function mint(
        address to,
        uint256 id,
        uint256 amount,
        bytes memory data
    ) public onlyRole(ASSET_ISSUER_ROLE) {
        _mint(to, id, amount, data);
        _setTokenMetadata(id, data);
        _setTokenRoyalty(id, to, _feeDenominator() / 10);
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
    function resetTokenMetadata(uint256 tokenId, bytes memory metadata)
        public
        requireSignatures(1)
        signatureFromOwner(tokenId)
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        _setTokenMetadata(tokenId, metadata);
    }
}
