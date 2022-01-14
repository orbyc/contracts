// SPDX-License-Identifier: MIT
// Orbyc Contracts v1.0.0 (security/ERC423/IERC423.sol)

pragma solidity ^0.8.0;

/**
 * @dev Interface of the ERC423 standard as defined in the EIP.
 */
interface IERC423 {
    /**
     * @dev Returns the `id` associated to the given `address`.
     */
    function idOf(address agent) external view returns (uint64);

    /**
     * @dev Returns the if given `agent` has `role`.
     */
    function hasRole(uint64 id, uint64 role) external view returns (bool);

    /**
     * @dev Returns info of the given `role`.
     */
    function roleInfo(uint64 role) external view returns (string memory);

    /**
     * @dev Returns info of agent with the given `id`.
     */
    function agentInfo(uint64 id) external view returns (string memory);

    /**
     * @dev Define `role` with the given `info`.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {RoleDefined} event.
     */
    function defineRole(uint64 role, string memory info) external returns (bool);

    /**
     * @dev Define an `agent` holder of the provided address with the specified `id` and `info`.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {AgentDefined} event.
     */
    function defineAgent(address agent, uint64 id, string memory info) external returns (bool);

    /**
     * @dev Grants `role` to the agent with the given `id`.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     */
    function grantRole(uint64 id, uint64 role) external returns (bool);

    /**
     * @dev Revokes `role` from the agent with the given `id`.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     */
    function revokeRole(uint64 id, uint64 role) external returns (bool);

    /**
     * @dev Emitted when `agent` is defined with the given `id`.
     */
    event RoleDefined(uint64 role, address issuer);

    /**
     * @dev Emitted when `agent` is defined with the given `id`.
     */
    event AgentDefined(address agent, uint64 id, address issuer);
}
