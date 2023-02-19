// SPDX-License-Identifier: MIT
// Orbyc Contracts v1.0.0 (token/ERC721/extensions/ERC721Compound.sol)

pragma solidity ^0.8.0;

/**
 * @title ERC-721 Non-Fungible Token Standard, optional append-only compound extension
 */
abstract contract ERC721Compound {
    struct Component {
        uint256 tokenId;
        string metadata;
    }

    mapping(uint256 => Component[]) private _components;

    function getComposition(uint256 tokenId)
        external
        view
        returns (Component[] memory)
    {
        return _components[tokenId];
    }

    function appendComponents(uint256 tokenId, Component[] memory components)
        external
        virtual
    {
        _appendComponents(tokenId, components);
    }

    function _appendComponents(uint256 tokenId, Component[] memory components)
        internal
        virtual
    {
        for (uint256 i = 0; i < components.length; i++) {
            _components[tokenId].push(components[i]);
        }
    }

    function _resetComposition(uint256 tokenId) internal virtual {
        delete _components[tokenId];
    }
}
