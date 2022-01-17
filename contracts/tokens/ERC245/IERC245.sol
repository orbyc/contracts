// SPDX-License-Identifier: MIT
// Orbyc Contracts v1.0.0 (tokens/ERC245/IERC245.sol)

pragma solidity ^0.8.0;

/**
 * @dev Interface of the ERC245 standard as defined in the EIP.
 */
interface IERC245 {
    /**
     * @dev Returns the name of the smart contract issuer
     */
    function name() external view returns (string memory);

    /**
     * @dev Returns the id for the provided `signer`
     */
    function idOf(address signer) external view returns (uint64);

    /**
     * @dev Returns the information present in the certificate (`issuer`, `metadata`)
     */
    function certificateInfo(uint256 id) external view returns (
        uint64, 
        string memory
    );

    /**
     * @dev Returns the information present in the asset (`owner`, `issuer`, `co2`, `cert`, `metadata`)
     */
    function assetInfo(uint256 id) external view returns (
        uint64,
        uint64,
        uint256,
        uint256,
        string memory
    );

    /**
     * @dev Returns the asset composition in two same-sized lists (`asset`, `portion`)
     * 
     * `asset`   -> asset id
     * `portion` -> portion that `asset` is present in the child asset
     */
    function assetComposition(uint256 id) external view returns (
        uint256[] memory, 
        uint16[] memory
    );

    /**
     * @dev Returns the asset traceability in three same-sized lists (`movement`, `lat`, `lng`)
     *
     * `movement` -> movement id
     * `lat`      -> geographical latitude
     * `lng`      -> geographical longitude
     */
    function assetTraceability(uint256 id) external view  returns (
        uint256[] memory,
        string[] memory,
        string[] memory
    );

    /**
     * @dev Returns the information presented in the movement (`issuer`, `lat`, `lng`, `co2`, `cert`, `metadata`)
     */
    function movementInfo(uint256 id) external view returns (
        uint64,
        string memory,
        string memory,
        uint256,
        uint256,
        string memory
    );

    /**
     * @dev Transfers ownership of the `asset` to the `recipient`.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {Transfer} event.
     */
    function transferOwnership(uint256 asset, uint64 recipient) external returns (bool);

    /**
     * @dev Issues a certificate in the supply chain.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {CertIssued} event.
     */
    function issueCertificate(
        uint256 id, 
        string memory metadata
    ) external returns (bool);

    /**
     * @dev Issues an asset in the supply chain.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {AssetIssued} event.
     */
    function issueAsset(
        uint256 id,
        uint64 owner,
        uint256 co2,
        uint256 cert,
        string memory metadata
    ) external returns (bool);

    /**
     * @dev Issues a movement in the supply chain.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {MovementIssued} event.
     */
    function issueMovement(
        uint256 id,
        string memory lat,
        string memory lng,
        uint256 co2,
        uint256 cert,
        string memory metadata
    ) external returns (bool);

    /**
     * @dev Append existing movements to asset traceability.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     */
    function addMovements(
        uint256 id, 
        uint256[] memory movements
    ) external returns (bool);

    /**
     * @dev Add existing assets to asset composition.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     */
    function addParents(
        uint256 id,
        uint256[] memory parents,
        uint16[] memory composition
    ) external returns (bool);

    /**
     * @dev Emitted when certificate with `id` was issued by `issuer` using `signer`
     */
    event CertIssued(uint256 id, uint64 issuer, address signer);
    /**
     * @dev Emitted when asset with `id` was issued by `issuer` using `signer`
     */
    event AssetIssued(uint256 id, uint64 issuer, address signer);
    /**
     * @dev Emitted when a movement with `id` was issued by `issuer` using `signer`
     */
    event MovementIssued(uint256 id, uint64 issuer, address signer);
     
    /**
     * @dev Emitted when asset with `id` has change ownership to `recipient` using `signer`.
     */
    event TransferOwnership(uint256 id, uint64 sender, uint64 recipient, address signer);
}
