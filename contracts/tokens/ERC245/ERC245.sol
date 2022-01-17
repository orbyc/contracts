// SPDX-License-Identifier: MIT
// Orbyc Contracts v1.0.0 (tokens/ERC245/ERC245.sol)

pragma solidity ^0.8.0;

import {Context} from "@openzeppelin/contracts/utils/Context.sol";
import {IERC245} from "./IERC245.sol";
import {Chain} from "./schema/IERC245Schema.sol";
import {Array} from "../../utils/Collections.sol";

abstract contract ERC245 is Context, IERC245 {
    string private _name;

    mapping(uint256 => Chain.Movement) private _movements;
    mapping(uint256 => Chain.Certificate) private _certs;
    mapping(uint256 => Chain.Asset) private _assets;

    constructor(string memory name_) {
        // assign contract metadata
        _name = name_;
    }

    /**
     * @dev Returns the name of the smart contract issuer
     */
    function name() public view virtual override returns (string memory) {
        return _name;
    }

    /**
     * @dev Returns the id for the provided `signer`
     */
    function idOf(address signer) public view virtual override returns (uint64);

    /**
     * @dev Returns the information present in the certificate (`issuer`, `metadata`)
     */
    function certificateInfo(uint256 id)
        public
        view
        virtual
        override
        returns (uint64, string memory)
    {
        Chain.Certificate storage cert = _certs[id];
        return (cert.issuer, cert.metadata);
    }

    /**
     * @dev Returns the information present in the asset (`owner`, `issuer`, `co2`, `cert`, `metadata`)
     */
    function assetInfo(uint256 id)
        public
        view
        virtual
        override
        returns (
            uint64,
            uint64,
            uint256,
            uint256,
            string memory
        )
    {
        Chain.Asset storage asset = _assets[id];
        return (
            asset.owner,
            asset.issuer,
            asset.co2,
            asset.cert,
            asset.metadata
        );
    }

    /**
     * @dev Returns the asset composition in two same-sized lists (`asset`, `portion`)
     *
     * `asset`   -> asset id
     * `portion` -> portion that `asset` is present in the child asset
     */
    function assetComposition(uint256 id)
        public
        view
        virtual
        override
        returns (uint256[] memory, uint16[] memory)
    {
        Chain.Asset storage asset = _assets[id];
        uint16[] memory percent = new uint16[](asset.parents.length);

        for (uint256 i = 0; i < asset.parents.length; i++) {
            percent[i] = asset.composition[asset.parents[i]];
        }

        return (asset.parents, percent);
    }

    /**
     * @dev Returns the asset traceability in three same-sized lists (`movement`, `lat`, `lng`)
     *
     * `movement` -> movement id
     * `lat`      -> geographical latitude
     * `lng`      -> geographical longitude
     */
    function assetTraceability(uint256 id)
        public
        view
        virtual
        override
        returns (
            uint256[] memory,
            string[] memory,
            string[] memory
        )
    {
        Chain.Asset storage asset = _assets[id];
        string[] memory lats = new string[](asset.traceability.length);
        string[] memory lngs = new string[](asset.traceability.length);

        for (uint256 i = 0; i < asset.traceability.length; i++) {
            lats[i] = _movements[asset.traceability[i]].lat;
            lngs[i] = _movements[asset.traceability[i]].lng;
        }

        return (asset.traceability, lats, lngs);
    }

    /**
     * @dev Returns the information presented in the movement (`lat`, `lng`, `co2`, `cert`, `metadata`)
     */
    function movementInfo(uint256 id)
        public
        view
        virtual
        override
        returns (
            uint64,
            string memory,
            string memory,
            uint256,
            uint256,
            string memory
        )
    {
        Chain.Movement storage movement = _movements[id];
        return (
            movement.issuer,
            movement.lat,
            movement.lng,
            movement.co2,
            movement.cert,
            movement.metadata
        );
    }

    /**
     * @dev See {IERC245-transferOwnership}
     */
    function transferOwnership(uint256 asset, uint64 recipient)
        public
        virtual
        override
        returns (bool)
    {
        return _transferOwnership(asset, recipient);
    }

    /**
     * @dev See {IERC245-issueCertificate}
     */
    function issueCertificate(uint256 id, string memory metadata)
        public
        virtual
        override
        returns (bool)
    {
        return _issueCertificate(id, metadata);
    }

    /**
     * @dev See {IERC245-issueAsset}
     */
    function issueAsset(
        uint256 id,
        uint64 owner,
        uint256 co2,
        uint256 cert,
        string memory metadata
    ) public virtual override returns (bool) {
        return _issueAsset(id, owner, co2, cert, metadata);
    }

    /**
     * @dev See {IERC245-issueMovement}
     */
    function issueMovement(
        uint256 id,
        string memory lat,
        string memory lng,
        uint256 co2,
        uint256 cert,
        string memory metadata
    ) public virtual override returns (bool) {
        return _issueMovement(id, lat, lng, co2, cert, metadata);
    }

    /**
     * @dev See {IERC245-addMovements}
     */
    function addMovements(uint256 id, uint256[] memory movements)
        public
        virtual
        override
        returns (bool)
    {
        return _addMovements(id, movements);
    }

    /**
     * @dev See {IERC245-addParents}
     */
    function addParents(
        uint256 id,
        uint256[] memory parents,
        uint16[] memory composition
    ) public virtual override returns (bool) {
        return _addParents(id, parents, composition);
    }

    /**
     * @dev Transfers ownership of the `asset` to the `recipient`.
     *
     * NOTE: any signer can make a transaction on behalf of the owner.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {Transfer} event.
     */
    function _transferOwnership(uint256 id, uint64 recipient)
        internal
        virtual
        returns (bool)
    {
        Chain.Asset storage asset = _assets[id];
        require(asset.id != 0, "Error: asset does not exists");

        uint64 sender = asset.owner;
        asset.owner = recipient;
        emit TransferOwnership(asset.id, sender, recipient, _msgSender());

        return true;
    }

    /**
     * @dev Issues a certificate in the supply chain.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {CertIssued} event.
     */
    function _issueCertificate(uint256 id, string memory metadata)
        internal
        virtual
        returns (bool)
    {
        uint64 agent = idOf(_msgSender());

        Chain.Certificate storage cert = _certs[id];
        require(cert.id == 0, "Error: certificate already exists");

        cert.id = id;
        cert.issuer = agent;
        cert.metadata = metadata;

        emit CertIssued(cert.id, agent, _msgSender());

        return true;
    }

    /**
     * @dev Issues an asset in the supply chain.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {AssetIssued} event.
     */
    function _issueAsset(
        uint256 id,
        uint64 owner,
        uint256 co2,
        uint256 cert,
        string memory metadata
    ) internal virtual returns (bool) {
        uint64 issuer = idOf(_msgSender());

        Chain.Asset storage asset = _assets[id];
        require(asset.id == 0, "Error: asset already exists");

        asset.id = id;
        asset.owner = owner;
        asset.issuer = issuer;
        asset.co2 = co2;
        asset.cert = cert;
        asset.metadata = metadata;

        emit AssetIssued(asset.id, issuer, _msgSender());

        return true;
    }

    /**
     * @dev Issues a movement in the supply chain.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {MovementIssued} event.
     */
    function _issueMovement(
        uint256 id,
        string memory lat,
        string memory lng,
        uint256 co2,
        uint256 cert,
        string memory metadata
    ) internal virtual returns (bool) {
        uint64 issuer = idOf(_msgSender());

        Chain.Movement storage movement = _movements[id];
        require(movement.id == 0, "Error: movement already exists");

        movement.id = id;
        movement.issuer = issuer;
        movement.lat = lat;
        movement.lng = lng;
        movement.co2 = co2;
        movement.cert = cert;
        movement.metadata = metadata;

        emit MovementIssued(movement.id, issuer, _msgSender());

        return true;
    }

    /**
     * @dev Append existing movement to asset traceability.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {AssetCo2Update} event.
     */
    function _addMovements(uint256 id, uint256[] memory movements)
        internal
        virtual
        returns (bool)
    {
        Chain.Asset storage asset = _assets[id];

        uint256 co2;
        for (uint256 i = 0; i < movements.length; i++) {
            require(
                !Array.contains(asset.traceability, movements[i]),
                "Error: movement already added"
            );

            Chain.Movement storage movement = _movements[movements[i]];
            require(movement.id != 0, "Error: can not add inexistent movement");

            asset.traceability.push(movement.id);
            co2 += movement.co2;
        }
        _updateCo2(asset.id, co2);
        return true;
    }

    /**
     * @dev Add existing assets to asset composition.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {AssetCo2Update} event.
     */
    function _addParents(
        uint256 id,
        uint256[] memory parents,
        uint16[] memory composition
    ) internal virtual returns (bool) {
        Chain.Asset storage asset = _assets[id];

        uint256 co2;
        for (uint256 i = 0; i < parents.length; i++) {
            require(
                asset.composition[parents[i]] == 0,
                "Error: parent already present"
            );

            Chain.Asset storage parent = _assets[parents[i]];
            require(parent.id != 0, "Error: can not add inexistent asset");

            asset.composition[parent.id] = composition[i];
            asset.parents.push(parent.id);
            parent.children.push(asset.id);
            co2 += parent.co2;
        }

        _updateCo2(asset.id, co2);
        return true;
    }

    /**
     * @dev Recursively updates the children's co2 value with the provided increment
     *
     * Emits a {AssetCo2Update} event.
     */
    function _updateCo2(uint256 id, uint256 co2) private {
        Chain.Asset storage asset = _assets[id];

        asset.co2 += co2;
        for (uint256 i = 0; i < asset.children.length; i++) {
            _updateCo2(asset.children[i], co2);
        }

        emit AssetCo2Update(asset.id, co2);
    }
}

/**
 * @dev Example of ERC245 full implementation.
 *
 * NOTE: Can be used in a private network with same-role actors.
 */
contract ERC245Private is ERC245 {
    constructor(string memory name_) ERC245(name_) {}

    function idOf(address) public pure override returns (uint64) {
        // by mapping every signer to the same `id` assigns every `signer` the same role.
        return 1;
    }
}
