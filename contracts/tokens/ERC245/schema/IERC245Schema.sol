// SPDX-License-Identifier: MIT
// Orbyc Contracts v1.0.0 (tokens/ERC245/schema/IERC245Schema.sol)

pragma solidity ^0.8.0;

library Chain {
    struct Certificate {
        // certificate unique identifier
        uint256 id;
        // certificate issuer id
        uint64 issuer;
        // certificate metadata
        string metadata;
    }

    struct Asset {
        // asset unique identifier
        uint256 id;
        // asset owner id
        uint64 owner;
        // asset issuer id
        uint64 issuer;
        // co2 emission (maybe on creation)
        uint256 co2;
        // creation co2 emission certificate
        uint256 cert;
        // list of assets that composes this asset
        uint256[] parents;
        // mapping of asset id to composition (sum represents 100%)
        mapping(uint256 => uint16) composition;
        // list of locations of the asset in order
        uint256[] traceability;
        // asset metadata
        string metadata;
    }

    struct Movement {
        // movement unique identifier
        uint256 id;
        // movement issuer id
        uint64 issuer;
        // latitude
        string lat;
        // longitude
        string lng;
        // co2 emission
        uint256 co2;
        // co2 emission certificate
        uint256 cert;
        // movement metadata
        string metadata;
    }
}
