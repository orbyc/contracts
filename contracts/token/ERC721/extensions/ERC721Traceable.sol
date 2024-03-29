// SPDX-License-Identifier: MIT
// Orbyc Contracts v1.0.0 (token/ERC721/extensions/ERC721Traceable.sol)

pragma solidity ^0.8.0;

/**
 * @title ERC-721 Non-Fungible Token Standard, optional append-only traceability extension
 */
abstract contract ERC721Traceable {
    struct Trace {
        string metadata;
    }

    mapping(uint256 => Trace[]) private _traces;

    /**
     * @dev Returns the traceability records of a non-fungible token with the specified token ID
     */
    function getTraceability(uint256 tokenId)
        external
        view
        returns (Trace[] memory)
    {
        return _traces[tokenId];
    }

    /**
     * @dev Adds new traceability records to a non-fungible token's history
     */
    function appendTraces(uint256 tokenId, Trace[] memory traces)
        external
        virtual
    {
        _appendTraces(tokenId, traces);
    }

    /**
     * @dev Adds new traceability records to a non-fungible token's history
     */
    function _appendTraces(uint256 tokenId, Trace[] memory traces)
        internal
        virtual
    {
        for (uint256 i = 0; i < traces.length; i++) {
            _traces[tokenId].push(traces[i]);
        }
    }

    /**
     * @dev Resets the traceability records of a non-fungible token with the specified token ID
     */
    function _resetTraceability(uint256 tokenId) internal virtual {
        delete _traces[tokenId];
    }
}
