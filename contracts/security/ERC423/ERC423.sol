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

    // mapping from agent Address to account Address
    mapping(address => address) _agentAccounts;

    // mapping from agent Address remove status
    mapping(address => bool) _agentRemoved;

    // mapping from agent ID to agent Info
    mapping(address => Identity.Account) private _accounts;

    // mapping from role ID to role Info
    mapping(uint256 => Identity.Role) private _roles;

    modifier onlyRole(uint256 indexed role) {
        require(hasRole(accountOf(_msgSender()), role), "Error: agent has not the required role");
        _;
    }

    constructor(string memory name_) {
        _name = name_;
    }

    /**
     * @dev Returns the accounts collection name.
     */
    function name() public view virtual override returns (string memory) {
        return _name;
    }

    /**
     * @dev Returns the `account` associated to the given `agent`.
     */
    function accountOf(address agent) public view virtual override returns (address) {
        require(!_agentRemoved[agent], "ERC423: agent has been removed");
        return _agentAccounts[agent];
    }

     /**
     * @dev Returns info of the `account`.
     */
    function accountInfo(address account) public view virtual override returns (string memory) {
        return _accounts[account].metadata;
    }

    /**
     * @dev Returns true if given `account` has `role`.
     */
    function hasRole(address account, uint256 role) public view virtual override returns (bool) {
        return _accounts[account].roles & role != 0;
    }

    /**
     * @dev Returns info of the given `role`.
     */
    function roleInfo(uint256 role) public view virtual override returns (string memory) {
        return _roles[role].metadata;
    }

    /**
     * @dev Define the `agent` with the given `account`
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {AgentDefined} event.
     */
    function defineAgent(
        address agent,
        address account,
        string calldata metadata_
    ) public virtual override returns (bool) {
        return _defineAgent(agent, account, metadata_);
    }

    /**
     * @dev Removes an agent from the suite. Removed agents can no longer be associated with accounts.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {AgentRemoved} event. 
     */
    function removeAgent(address agent) public virtual override returns (bool){
        return _removeAgent(agent);
    }

    /**
     * @dev Define `role`.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {RoleDefined} event.
     */
    function defineRole(uint256 role, string calldata metadata_) public virtual override returns (bool) {
        return _defineRole(role, metadata_);
    }

    /**
     * @dev Grants `role` to the given `account`.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {RoleGranted} event.
     */
    function  grantRole(address account, uint256 role)
        public
        virtual
        override
        returns (bool)
    {
        return _grantRole(account, role);
    }

     /**
     * @dev Revokes `role` from the given `account`.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {RoleRevoked} event.
     */
    function revokeRole(address account, uint256 role)
        public
        virtual
        override
        returns (bool)
    {
        return _revokeRole(account, role);
    }

    /**
     * @dev See {IERC423-defineAgent}
     */
    function _defineAgent(address agent, address account, string calldata metadata_) internal virtual returns (bool) {
        _accounts[account].metadata = metadata_;
        _agentAccounts[agent] = account;

        emit AgentDefined(agent, account, _msgSender());
        return true;
    }

    /**
     * @dev See {IERC423-removeAgent}
     */
    function _removeAgent(address agent) internal virtual returns (bool) {
         _agentRemoved[agent] = true;

        emit AgentRemoved(agent, _msgSender());
        return true;
    }

    /**
     * @dev See {IERC423-defineRole}
     */
    function _defineRole(uint256 role, string calldata metadata_)
        internal
        virtual
        returns (bool)
    {
        _roles[role] = Identity.Role(role, metadata_);

        emit RoleDefined(role, _msgSender());
        return true;
    }

    /**
     * @dev See {IERC423-grantRole}
     */
    function _grantRole(address account, uint256 role) internal virtual returns (bool) {
        _accounts[account].roles |= role;

        emit RoleGranted(role, account, _msgSender());
        return true;
    }

    /**
     * @dev See {IERC423-revokeRole}
     */
    function _revokeRole(address account, uint256 role) internal virtual returns (bool) {
        _accounts[account].roles ^= role;

        emit RoleRevoked(role, account, _msgSender());
        return true;
    }
}
