// SPDX-License-Identifier: MIT
// Orbyc Contracts v1.0.0 (security/ERC423/schema/IERC423Schema.sol)

pragma solidity ^0.8.0;

library Identity {
    struct Agent {
        // agent unique identifier
        uint64 id;
        // agent roles
        uint64 roles;
        // agent metadata
        string info;
    }

    struct Role {
        // role unique identifier
        uint64 role;
        // role metadata
        string info;
    }
}
