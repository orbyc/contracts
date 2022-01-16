// SPDX-License-Identifier: MIT
// Orbyc Contracts v1.0.0 (security/ERC423/ERC423.sol)

pragma solidity ^0.8.0;

import {Context} from "@openzeppelin/contracts/utils/Context.sol";
import {IERC423} from "./IERC423.sol";
import {IERC423Metadata} from "./extensions/IERC423Metadata.sol";
import {Identity} from "./schema/IERC423Schema.sol";

contract ERC423 is Context, IERC423, IERC423Metadata {
    // Suite Name
    string private _name;

    // mapping from agent Address to agent ID
    mapping(address => uint64) _agents;

    // mapping from agent ID to agent Info
    mapping(uint64 => Identity.Agent) private _agentsInfo;

    // mapping from role ID to role Info
    mapping(uint64 => Identity.Role) private _rolesInfo;

    modifier onlyRole(uint64 indexed role) {
        require(hasRole(idOf(_msgSender()), role), "Error: agent has not the required role");
        _;
    }

    constructor(string memory name_) {
        _name = name_;
    }

    /**
     * @dev Returns the agents collection name.
     */
    function name() public view virtual override returns (string memory) {
        return _name;
    }

    /**
     * @dev Returns the `id` associated to the given `address`.
     */
    function idOf(address agent) public view virtual override returns (uint64) {
        return _idOf(agent);
    }

    /**
     * @dev Returns the if given `agent` has `role`.
     */
    function hasRole(uint64 id, uint64 role)
        public
        view
        virtual
        override
        returns (bool)
    {
        return _hasRole(id, role);
    }

    /**
     * @dev Returns info of the given `role`.
     */
    function roleInfo(uint64 role)
        public
        view
        virtual
        override
        returns (string memory)
    {
        return _roleInfo(role);
    }

    /**
     * @dev Returns info of agent with the given `id`.
     */
    function agentInfo(uint64 id)
        public
        view
        virtual
        override
        returns (string memory)
    {
        return _agentInfo(id);
    }

    /**
     * @dev Define `role` with the given `info`.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {RoleDefined} event.
     */
    function defineRole(uint64 role, string memory info)
        public
        virtual
        override
        returns (bool)
    {
        return _defineRole(role, info);
    }

    /**
     * @dev Define an `agent` holder of the provided address with the specified `id` and `info`.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {AgentDefined} event.
     */
    function defineAgent(
        address agent,
        uint64 id,
        string memory info
    ) public virtual override returns (bool) {
        return _defineAgent(agent, id, info);
    }

    /**
     * @dev Grants `role` to the agent with the given `id`.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     */
    function grantRole(uint64 id, uint64 role)
        public
        virtual
        override
        returns (bool)
    {
        return _grantRole(id, role);
    }

    /**
     * @dev Revokes `role` from the agent with the given `id`.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     */
    function revokeRole(uint64 id, uint64 role)
        public
        virtual
        override
        returns (bool)
    {
        return _revokeRole(id, role);
    }

    /**
     * @dev See {IERC423-idOf}
     */
    function _idOf(address agent) internal view virtual returns (uint64) {
        return _agents[agent];
    }

    /**
     * @dev See {IERC423-hasRole}
     */
    function _hasRole(uint64 id, uint64 role)
        internal
        view
        virtual
        returns (bool)
    { 
        return  _agentsInfo[id].roles & role != 0;
    }

    /**
     * @dev See {IERC423-roleInfo}
     */
    function _roleInfo(uint64 id)
        internal
        view
        virtual
        returns (string memory)
    { 
        return _rolesInfo[id].info;
    }

    /**
     * @dev See {IERC423-agentInfo}
     */
    function _agentInfo(uint64 id)
        internal
        view
        virtual
        returns (string memory)
    {
        return _agentsInfo[id].info;
    }

    /**
     * @dev See {IERC423-defineRole}
     */
    function _defineRole(uint64 role, string memory info)
        internal
        virtual
        returns (bool)
    {
        _rolesInfo[role] = Identity.Role(role, info);
        emit RoleDefined(role, _msgSender());

        return true;
    }

    /**
     * @dev See {IERC423-defineAgent}
     */
    function _defineAgent(
        address agent,
        uint64 id,
        string memory info
    ) internal virtual returns (bool) {
        _agentsInfo[id].info = info;
        _agents[agent] = id;

        emit AgentDefined(agent, id, _msgSender());
        return true;
    }

    /**
     * @dev See {IERC423-grantRole}
     */
    function _grantRole(uint64 id, uint64 role) internal virtual returns (bool) {
        _agentsInfo[id].roles |= role;
        return true;
    }

    /**
     * @dev See {IERC423-revokeRole}
     */
    function _revokeRole(uint64 id, uint64 role) internal virtual returns (bool) {
        _agentsInfo[id].roles ^= role;
        return true;
    }
}
