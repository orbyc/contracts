// SPDX-License-Identifier: MIT
// Orbyc Contracts v1.0.0 (token/ERC721/extensions/ERC721EmbedMetadata.sol)

pragma solidity ^0.8.0;

/**
 * @title ERC-721 Non-Fungible Token Standard, optional embeded metadata extension
 */
abstract contract ERC721EmbedMetadata {
    mapping(uint256 => string) private _metadata;

    function getMetadata(uint256 tokenId)
        external
        view
        returns (string memory)
    {
        return _metadata[tokenId];
    }

    function _setTokenMetadata(uint256 tokenId, string memory _data)
        internal
        virtual
    {
        _metadata[tokenId] = _data;
    }
}
