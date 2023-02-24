// SPDX-License-Identifier: MIT
// Orbyc Contracts v1.0.0 (token/ERC721/extensions/ERC721Certificable.sol)

pragma solidity ^0.8.0;

/**
 * @title ERC-721 Non-Fungible Token Standard, optional append-only certificates extension
 */
abstract contract ERC721Certificable {
    struct Certificate {
        uint256 id;
        string metadata;
    }

    mapping(uint256 => Certificate[]) private _certifications;

    /**
     * @dev Returns the certification records of a non-fungible token with the specified token ID
     */
    function getCertifications(uint256 tokenId)
        external
        view
        returns (Certificate[] memory)
    {
        return _certifications[tokenId];
    }

    /**
     * @dev Adds new certification records to a non-fungible token
     */
    function appendCertificates(
        uint256 tokenId,
        Certificate[] memory certificates
    ) external virtual {
        _appendCertificates(tokenId, certificates);
    }

    /**
     * @dev Adds new certification records to a non-fungible token
     */
    function _appendCertificates(
        uint256 tokenId,
        Certificate[] memory certificates
    ) internal virtual {
        for (uint256 i = 0; i < certificates.length; i++) {
            _certifications[tokenId].push(certificates[i]);
        }
    }

    /**
     * @dev Resets the certification records of a non-fungible token with the specified token ID
     */
    function _resetCertifications(uint256 tokenId) internal virtual {
        delete _certifications[tokenId];
    }
}
