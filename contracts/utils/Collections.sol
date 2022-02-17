// SPDX-License-Identifier: MIT
// Orbyc Utils v1.0.0 (utils/Collection.sol)

pragma solidity ^0.8.0;

library Array {
    function contains(address[] memory collection, address element)
        public
        pure
        returns (bool)
    {
        for (uint256 i = 0; i < collection.length; i++) {
            if (collection[i] == element) {
                return true;
            }
        }
        return false;
    }

    function contains(uint256[] memory collection, uint256 element)
        public
        pure
        returns (bool)
    {
        for (uint256 i = 0; i < collection.length; i++) {
            if (collection[i] == element) {
                return true;
            }
        }
        return false;
    }
}
