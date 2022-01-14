// SPDX-License-Identifier: MIT
// Orbyc Contracts v1.0.0 (security/ERC423/extensions/IERC423Metadata.sol)

pragma solidity ^0.8.0;

import "../IERC423.sol";

/**
 * @title ERC-423 Agents Control Standard, optional metadata extension
 * @dev See https://eips.ethereum.org/EIPS/eip-423
 */
interface IERC423Metadata is IERC423 {
    /**
     * @dev Returns the agents collection name.
     */
    function name() external view returns (string memory);
}
