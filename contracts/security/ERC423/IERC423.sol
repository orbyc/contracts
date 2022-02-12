// SPDX-License-Identifier: MIT
// Orbyc Contracts v1.0.0 (security/ERC423/IERC423.sol)

pragma solidity ^0.8.0;

/**
 * @dev Interface of the ERC423 standard as defined in the EIP.
 */
interface IERC423 {
    /**
     * @dev Emitted when `agent` is defined with the given `account`.
     */
    event AgentDefined(address agent, address account, address issuer);
    /**
     * @dev Emitted when `agent` removed.
     */
    event AgentRemoved(address agent, address issuer);

    /**
     * @dev Emitted when `role` is defined.
     */
    event RoleDefined(uint256 role, address issuer);
    /**
     * @dev Emitted when `role` is granted to the given `account`.
     */
    event RoleGranted(uint256 role, address account, address issuer);
    /**
     * @dev Emitted when `role` is revoked from the given `account`.
     */
    event RoleRevoked(uint256 role, address account, address issuer);
    

    /**
     * @dev Returns the `account` associated to the given `agent`.
     */
    function accountOf(address agent) external view returns (address);

    /**
     * @dev Returns info of the `account`.
     */
    function accountInfo(address account) external view returns (string memory);

    /**
     * @dev Returns true if given `account` has `role`.
     */
    function hasRole(address account, uint256 role) external view returns (bool);

    /**
     * @dev Returns info of the given `role`.
     */
    function roleInfo(uint256 role) external view returns (string memory);

    /**
     * @dev Define the `agent` with the given `account`
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {AgentDefined} event.
     */
    function defineAgent(address agent, address account, string memory metadata_) external returns (bool);

    /**
     * @dev Removes an agent from the suite. Removed agents can no longer be associated with accounts.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {AgentRemoved} event. 
     */
    function removeAgent(address agent) external returns (bool);

    /**
     * @dev Define `role`.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {RoleDefined} event.
     */
    function defineRole(uint256 role, string memory metadata_) external returns (bool);

    /**
     * @dev Grants `role` to the given `account`.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {RoleGranted} event.
     */
    function grantRole(address account, uint256 role) external returns (bool);

    /**
     * @dev Revokes `role` from the given `account`.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {RoleRevoked} event.
     */
    function revokeRole(address account, uint256 role) external returns (bool);
}
