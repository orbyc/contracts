// SPDX-License-Identifier: MIT
// Orbyc Contracts v1.0.0 (tokens/ERC245/schema/IERC245Schema.sol)

pragma solidity ^0.8.0;

library Chain {
    struct Certificate {
        // certificate unique identifier
        uint256 id;
        // certificate issuer
        address issuer;
        // certificate metadata
        string metadata;
    }

    struct Asset {
        // asset unique identifier
        uint256 id;
        // asset owner
        address owner;
        // asset issuer
        address issuer;
        // co2 emission (maybe on creation)
        uint64 co2e;
        // creation co2 emission certificate
        uint256 certId;
        // list of other certificates
        uint256[] certificates;
        // list of locations of the asset in order
        uint256[] traceability;
        // list of assets that composes this asset
        uint256[] parents;
        // mapping of asset id to composition representing the portion of the father that composes the asset.
        mapping(uint256 => uint16) composition;
        // asset metadata
        string metadata;
    }

    struct Movement {
        // movement unique identifier
        uint256 id;
        // movement issuer id
        address issuer;
        // latitude
        string lat;
        // longitude
        string lng;
        // co2 emission
        uint64 co2e;
        // co2 emission certificate
        uint256 certId;
        // movement metadata
        string metadata;
    }
}
