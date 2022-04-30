// SPDX-License-Identifier: MIT
// Orbyc Contracts v1.0.0 (tokens/ERC245/ERC245.sol)

pragma solidity ^0.8.0;

import {Context} from "@openzeppelin/contracts/utils/Context.sol";
import {IERC245Metadata} from "./extensions/IERC245Metadata.sol";
import {IERC245} from "./IERC245.sol";
import {Chain} from "./schema/IERC245Schema.sol";
import {Array} from "../../utils/Collections.sol";

contract ERC245 is Context, IERC245, IERC245Metadata {
    string private _name;

    mapping(uint256 => Chain.Movement) private _movements;
    mapping(uint256 => Chain.Certificate) private _certs;
    mapping(uint256 => Chain.Asset) private _assets;

    constructor(string memory name_) {
        _name = name_;
    }

    /**
     * @dev Returns the supply chain name.
     */
    function name() public view virtual override returns (string memory) {
        return _name;
    }

    /**
     * @dev Returns the information present in the asset (`owner`, `issuer`, `co2e`, `certId`, `metadata`)
     */
    function assetInfo(uint256 assetId)
        public
        view
        virtual
        override
        returns (
            address,
            address,
            uint64,
            uint256,
            string memory
        )
    {
        Chain.Asset storage asset = _assets[assetId];
        return (
            asset.owner,
            asset.issuer,
            asset.co2e,
            asset.certId,
            asset.metadata
        );
    }

    /**
     * @dev Returns the asset certificates in a `certificate` list
     *
     * `certificate` -> certificate id
     */
    function assetCertificates(uint256 assetId)
        public
        view
        virtual
        override
        returns (uint256[] memory)
    {
        Chain.Asset storage asset = _assets[assetId];

        return asset.certificates;
    }

    /**
     * @dev Returns the asset composition in two same-sized lists (`asset`, `portion`)
     *
     * `asset`   -> asset id
     * `portion` -> portion of the parent `asset` that conforms the current `asset`
     */
    function assetComposition(uint256 assetId)
        public
        view
        virtual
        override
        returns (uint256[] memory, uint16[] memory)
    {
        Chain.Asset storage asset = _assets[assetId];
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
    function assetTraceability(uint256 assetId)
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
        Chain.Asset storage asset = _assets[assetId];
        string[] memory lats = new string[](asset.traceability.length);
        string[] memory lngs = new string[](asset.traceability.length);

        for (uint256 i = 0; i < asset.traceability.length; i++) {
            lats[i] = _movements[asset.traceability[i]].lat;
            lngs[i] = _movements[asset.traceability[i]].lng;
        }

        return (asset.traceability, lats, lngs);
    }

    /**
     * @dev Returns the information present in the certificate (`issuer`, `metadata`)
     */
    function certificateInfo(uint256 certId)
        public
        view
        virtual
        override
        returns (address, string memory)
    {
        Chain.Certificate storage cert = _certs[certId];
        return (cert.issuer, cert.metadata);
    }

    /**
     * @dev Returns the information presented in the movement (`issuer`, `lat`, `lng`, `co2e`, `certId`, `metadata`)
     */
    function movementInfo(uint256 moveId)
        public
        view
        virtual
        override
        returns (
            address,
            string memory,
            string memory,
            uint64,
            uint256,
            string memory
        )
    {
        Chain.Movement storage movement = _movements[moveId];
        return (
            movement.issuer,
            movement.lat,
            movement.lng,
            movement.co2e,
            movement.certId,
            movement.metadata
        );
    }

    /**
     * @dev Transfers `asset` ownership.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {Transfer} event.
     */
    function transferFrom(
        address from,
        address to,
        uint256 assetId
    ) public virtual override {
        _transferFrom(from, to, assetId);
    }

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
    ) public virtual override returns (bool) {
        return
            _issueAsset(assetId, owner, _msgSender(), co2e, certId, metadata_);
    }

    /**
     * @dev Issues a certificate in the supply chain.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {CertIssued} event.
     */
    function issueCertificate(uint256 certId, string calldata metadata_)
        public
        virtual
        override
        returns (bool)
    {
        return _issueCertificate(certId, _msgSender(), metadata_);
    }

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
    ) public virtual override returns (bool) {
        return
            _issueMovement(
                moveId,
                _msgSender(),
                lat,
                lng,
                co2e,
                certId,
                metadata_
            );
    }

    /**
     * @dev Append existing movements to asset traceability.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {CertificateAssigned} event.
     */
    function addCertificates(uint256 assetId, uint256[] memory certificates)
        public
        virtual
        override
        returns (bool)
    {
        return _addCertificates(assetId, certificates);
    }

    /**
     * @dev Append existing movements to asset traceability.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {MovementAssigned} event.
     */
    function addMovements(uint256 assetId, uint256[] memory movements)
        public
        virtual
        override
        returns (bool)
    {
        return _addMovements(assetId, movements);
    }

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
    ) public virtual override returns (bool) {
        return _addParents(assetId, parents, composition);
    }

    /**
     * @dev See {IERC245-issueAsset}
     */
    function _issueAsset(
        uint256 assetId,
        address owner,
        address issuer,
        uint64 co2e,
        uint256 certId,
        string memory metadata_
    ) internal virtual returns (bool) {
        Chain.Asset storage asset = _assets[assetId];
        require(asset.id == 0, "Error: asset already exists");

        Chain.Certificate storage certificate = _certs[certId];
        require(certificate.id == certId, "Error: certificate does not exists");

        asset.id = assetId;
        asset.owner = owner;
        asset.issuer = issuer;
        asset.co2e = co2e;
        asset.certId = certificate.id;
        asset.metadata = metadata_;

        emit AssetIssued(asset.id, _msgSender());
        return true;
    }

    /**
     * @dev See {IERC245-issueCertificate}
     */
    function _issueCertificate(
        uint256 certId,
        address issuer,
        string memory metadata_
    ) internal virtual returns (bool) {
        Chain.Certificate storage cert = _certs[certId];
        require(cert.id == 0, "Error: certificate already exists");

        cert.id = certId;
        cert.issuer = issuer;
        cert.metadata = metadata_;

        emit CertIssued(cert.id, _msgSender());
        return true;
    }

    /**
     * @dev See {IERC245-issueMovement}
     */
    function _issueMovement(
        uint256 moveId,
        address issuer,
        string memory lat,
        string memory lng,
        uint64 co2e,
        uint256 certId,
        string memory metadata_
    ) internal virtual returns (bool) {
        Chain.Movement storage movement = _movements[moveId];
        require(movement.id == 0, "Error: movement already exists");

        Chain.Certificate storage certificate = _certs[certId];
        require(certificate.id == certId, "Error: certificate does not exists");

        movement.id = moveId;
        movement.issuer = issuer;
        movement.lat = lat;
        movement.lng = lng;
        movement.co2e = co2e;
        movement.certId = certificate.id;
        movement.metadata = metadata_;

        emit MovementIssued(movement.id, _msgSender());
        return true;
    }

    function _addCertificates(uint256 assetId, uint256[] memory certificates)
        internal
        virtual
        returns (bool)
    {
        Chain.Asset storage asset = _assets[assetId];

        for (uint256 i = 0; i < certificates.length; i++) {
            require(
                !Array.contains(asset.certificates, certificates[i]),
                "Error: certificate already added"
            );

            Chain.Certificate storage certificate = _certs[certificates[i]];
            require(
                certificate.id != 0,
                "Error: can not add inexistent certificate"
            );

            asset.certificates.push(certificate.id);

            emit CertificateAssigned(certificate.id, assetId, _msgSender());
        }

        return true;
    }

    /**
     * @dev See {IERC245-addMovements}
     */
    function _addMovements(uint256 assetId, uint256[] memory movements)
        internal
        virtual
        returns (bool)
    {
        Chain.Asset storage asset = _assets[assetId];

        for (uint256 i = 0; i < movements.length; i++) {
            require(
                !Array.contains(asset.traceability, movements[i]),
                "Error: movement already added"
            );

            Chain.Movement storage movement = _movements[movements[i]];
            require(movement.id != 0, "Error: can not add inexistent movement");

            asset.traceability.push(movement.id);

            emit MovementAssigned(movement.id, assetId, _msgSender());
        }

        return true;
    }

    /**
     * @dev See {IERC245-addParents}
     */
    function _addParents(
        uint256 assetId,
        uint256[] memory parents,
        uint16[] memory composition
    ) internal virtual returns (bool) {
        Chain.Asset storage asset = _assets[assetId];

        for (uint256 i = 0; i < parents.length; i++) {
            require(
                asset.composition[parents[i]] == 0,
                "Error: parent already present"
            );

            Chain.Asset storage parent = _assets[parents[i]];
            require(parent.id != 0, "Error: can not add inexistent asset");

            asset.composition[parent.id] = composition[i];
            asset.parents.push(parent.id);

            emit ParentAssigned(parent.id, assetId, _msgSender());
        }

        return true;
    }

    /**
     * @dev See {IERC245-transferFrom}
     */
    function _transferFrom(
        address from,
        address to,
        uint256 assetId
    ) internal virtual {
        Chain.Asset storage asset = _assets[assetId];
        require(asset.owner == from, "ERC245: sender is not the asset owner");

        asset.owner = to;

        emit TransferAsset(from, to, assetId);
    }
}
