// SPDX-License-Identifier: MIT
// Orbyc Contracts v1.0.0 (security/ERC423/schema/IERC423Schema.sol)

pragma solidity ^0.8.0;

library Identity {
    struct Account {
        // account address
        address account;
        // account roles
        uint256 roles;
        // account metadata
        string metadata;
    }

    struct Role {
        // role unique identifier
        uint256 role;
        // role metadata
        string metadata;
    }
}
