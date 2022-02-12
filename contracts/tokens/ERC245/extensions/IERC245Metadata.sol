// SPDX-License-Identifier: MIT
// Orbyc Contracts v1.0.0 (tokens/ERC245/extensions/IERC245Metadata.sol)

pragma solidity ^0.8.0;

import "../IERC245.sol";

/**
 * @title ERC-245 CO2e Supply Chain, optional metadata extension
 * @dev See https://eips.ethereum.org/EIPS/eip-245
 */
interface IERC245Metadata is IERC245 {
    /**
     * @dev Returns the supply chain name.
     */
    function name() external view returns (string memory);
}
