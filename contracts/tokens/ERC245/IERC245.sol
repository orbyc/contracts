// SPDX-License-Identifier: MIT
// Orbyc Contracts v1.0.0 (tokens/ERC245/IERC245.sol)

pragma solidity ^0.8.0;

/**
 * @dev Interface of the ERC245 standard as defined in the EIP.
 */
interface IERC245 {
    /**
     * @dev Emitted when `asset` is transfered by `from` sender `to` recipient
     */
    event TransferAsset(address indexed from, address indexed to, uint256 indexed assetId);

    /**
     * @dev Emitted when `asset` was issued by `signer`
     */
    event AssetIssued(uint256 assetId, address signer);
    /**
     * @dev Emitted when `certificate` was issued by `signer`
     */
    event CertIssued(uint256 certId, address signer);
    /**
     * @dev Emitted when `movement` was issued by `signer`
     */
    event MovementIssued(uint256 moveId, address signer);
     
    /**
     * @dev Emitted when `certificate` is assigned to `asset` by `signer`
     */
    event CertificateAssigned(uint256 certId, uint256 assetId, address signer);
    /**
     * @dev Emitted when `movement` is assigned to `asset` by `signer`
     */
    event MovementAssigned(uint256 moveId, uint256 assetId, address signer);
     /**
     * @dev Emitted when `parent` is assigned to `asset` by `signer`
     */
    event ParentAssigned(uint256 parentId, uint256 assetId, address signer);

    /**
     * @dev Returns the information present in the asset (`owner`, `issuer`, `co2e`, `certId`, `metadata`)
     */
    function assetInfo(uint256 assetId) external view returns (
        address,
        address,
        uint64,
        uint256,
        string memory
    );

    /**
     * @dev Returns the asset certificates in a `certificate` list
     * 
     * `certificate` -> certificate id
     */
    function assetCertificates(uint256 assetId) external view returns (uint256[] memory);

    /**
     * @dev Returns the asset composition in two same-sized lists (`asset`, `portion`)
     * 
     * `asset`   -> asset id
     * `portion` -> portion of the parent `asset` that conforms the current `asset`
     */
    function assetComposition(uint256 assetId) external view returns (
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
    function assetTraceability(uint256 assetId) external view  returns (
        uint256[] memory,
        string[] memory,
        string[] memory
    );

    /**
     * @dev Returns the information present in the certificate (`issuer`, `metadata`)
     */
    function certificateInfo(uint256 certId) external view returns (
        address, 
        string memory
    );

    /**
     * @dev Returns the information presented in the movement (`issuer`, `lat`, `lng`, `co2e`, `certId`, `metadata`)
     */
    function movementInfo(uint256 moveId) external view returns (
        address,
        string memory,
        string memory,
        uint64,
        uint256,
        string memory
    );

    /**
     * @dev Transfers `asset` ownership.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {TransferAsset} event.
     */
    function transferFrom(
        address from,
        address to,
        uint256 assetId
    ) external;

    /**
     * @dev Issues an asset in the supply chain.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {AssetIssued} event.
     */
    function issueAsset(
        uint256 assetId,
        address owner,
        uint64 co2e,
        uint256 certId,
        string memory metadata_
    ) external returns (bool);

    /**
     * @dev Issues a certificate in the supply chain.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {CertIssued} event.
     */
    function issueCertificate(
        uint256 certId, 
        string calldata metadata_
    ) external returns (bool);

    /**
     * @dev Issues a movement in the supply chain.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {MovementIssued} event.
     */
    function issueMovement(
        uint256 moveId,
        string memory lat,
        string memory lng,
        uint64 co2e,
        uint256 certId,
        string memory metadata_
    ) external returns (bool);

    /**
     * @dev Append existing movements to asset traceability.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {CertificateAssigned} event.
     */
    function addCertificates(
        uint256 assetId, 
        uint256[] memory certificates
    ) external returns (bool);

    /**
     * @dev Append existing movements to asset traceability.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {MovementAssigned} event.
     */
    function addMovements(
        uint256 assetId, 
        uint256[] memory movements
    ) external returns (bool);

    /**
     * @dev Add existing assets to asset composition.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {ParentAssigned} event.
     */
    function addParents(
        uint256 assetId,
        uint256[] memory parents,
        uint16[] memory composition
    ) external returns (bool);
}
