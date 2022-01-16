// SPDX-License-Identifier: MIT
// Orbyc Contracts v1.0.0 (Suite.sol)

pragma solidity ^0.8.0;

import {ERC423} from "./security/ERC423/ERC423.sol";

contract Suite is ERC423 {
    /**
     * Roles
     */

    uint64 public ADMIN_ROLE = 0x01 << 1;
    uint64 public EDITOR_ROLE = 0x01 << 2;

    /**
     * Agents
     */

    uint64 public ORBYC_AGENT = 0x01;
    uint64 public NULL_AGENT = 0xfffffffffff;

    /**
     * Modifiers
     */

    modifier nonNull(uint64 agent) {
        require(agent != NULL_AGENT, "Error: can not operate over null agent");
        _;
    }

    modifier validAddress(address signer) {
        require(
            idOf(signer) != NULL_AGENT,
            "Error: agent is null address, can not operate"
        );
        _;
    }

    /**
     * Constructor
     */

    constructor() ERC423("Orbyc Agents Suite") {
        // define roles
        _defineRole(
            ADMIN_ROLE,
            '{"name":"ADMIN", "description":"allowed to issue and assign roles"}'
        );
        _defineRole(
            EDITOR_ROLE,
            '{"name":"EDITOR", "description":"allowed to register addresses as agents"}'
        );

        // define orbyc agent
        _defineAgent(
            _msgSender(),
            ORBYC_AGENT,
            '{"name":"Orbyc Default Admin"}'
        );

        // add defined roles to orbyc agent
        _grantRole(ORBYC_AGENT, ADMIN_ROLE);
        _grantRole(ORBYC_AGENT, EDITOR_ROLE);
    }

    /**
     * ERC423
     */

    function defineRole(uint64 role, string memory info)
        public
        override
        onlyRole(ADMIN_ROLE)
        returns (bool)
    {
        return super.defineRole(role, info);
    }

    function defineAgent(
        address agent,
        uint64 id,
        string memory info
    ) public override onlyRole(EDITOR_ROLE) validAddress(agent) returns (bool) {
        return super.defineAgent(agent, id, info);
    }

    function grantRole(uint64 id, uint64 role)
        public
        override
        onlyRole(ADMIN_ROLE)
        nonNull(id)
        returns (bool)
    {
        return super.grantRole(id, role);
    }

    function revokeRole(uint64 id, uint64 role)
        public
        override
        onlyRole(ADMIN_ROLE)
        nonNull(id)
        returns (bool)
    {
        return super.revokeRole(id, role);
    }

    /**
     * @dev Banns an address making it impossible to opperate again
     */
    function bannAddress(address agent)
        public
        onlyRole(ADMIN_ROLE)
        returns (bool)
    {
        return _defineAgent(agent, NULL_AGENT, "");
    }
}
