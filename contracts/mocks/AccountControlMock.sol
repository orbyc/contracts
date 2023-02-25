// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "closedzeppelin-contracts/access/IAccountControl.sol";

contract AccountControlMock is IAccountControl {
    function accountOf(address signer) public pure returns (address) {
        return signer;
    }

    function addSigner(
        address,
        address,
        string memory
    ) public pure returns (bool) {
        return true;
    }

    function removeSigner(address, string memory) public pure returns (bool) {
        return true;
    }
}
