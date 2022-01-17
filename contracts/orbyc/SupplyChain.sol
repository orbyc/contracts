// SPDX-License-Identifier: MIT
// Orbyc Contracts v1.0.0 (Supply.sol)

pragma solidity ^0.8.0;

import {Pausable} from "@openzeppelin/contracts/security/Pausable.sol";
import {Context} from "@openzeppelin/contracts/utils/Context.sol";
import {ERC423} from "../security/ERC423/ERC423.sol";
import {Array} from "../utils/Collections.sol";
import {ERC245} from "../tokens/ERC245/ERC245.sol";
import {Chain} from "../tokens/ERC245/schema/IERC245Schema.sol";

library Multi {
    struct Transfer {
        // asset id
        uint256 asset;
        // recipient agent id
        uint64 recipient;
        // agent id => hasSigned
        mapping(uint64 => bool) signatures;
    }
}

contract SupplyChain is ERC245, ERC423, Pausable {
    /**
     * Roles
     */

    uint64 public ADMIN_ROLE = 0x01 << 1;
    uint64 public EDITOR_ROLE = 0x01 << 2;
    uint64 public PROVIDER_ROLE = 0x01 << 3;
    uint64 public CERT_ISSUER_ROLE = 0x01 << 4;
    uint64 public ASSET_ISSUER_ROLE = 0x01 << 5;
    uint64 public MOVEMENT_ISSUER_ROLE = 0x01 << 6;

    /**
     * Agents
     */

    uint64 public ORBYC_AGENT = 0x01;
    uint64 public NULL_AGENT = 0xfffffffffff;

    /**
     * Privates
     */

    // mapping from asset id to ownership multitransfer
    mapping(uint256 => Multi.Transfer) private _transfers;

    /**
     * Modifiers
     */

    modifier nonNull(uint64 agent) {
        require(agent != NULL_AGENT, "Error: can not operate over null agent");
        _;
    }

    modifier validAddress(address signer) {
        require(
            idOf(signer) != NULL_AGENT,
            "Error: agent is null address, can not operate"
        );
        _;
    }

    modifier onlyIssuer(uint256 id) {
        uint64 issuer;
        (, issuer, , , ) = assetInfo(id);
        uint64 agent = idOf(_msgSender());
        require(agent == issuer, "Error: agent is not issuer");
        _;
    }

    /**
     * Constructor
     */

    constructor(string memory name_) ERC245(name_) ERC423(name_) {
        // define roles
        _defineRole(
            ADMIN_ROLE,
            '{"name":"ADMIN", "description":"allowed to issue and assign roles"}'
        );
        _defineRole(
            EDITOR_ROLE,
            '{"name":"EDITOR", "description":"allowed to register addresses as agents"}'
        );
        _defineRole(
            PROVIDER_ROLE,
            '{"name":"PROVIDER", "description":"allowed to hold assets"}'
        );
        _defineRole(
            CERT_ISSUER_ROLE,
            '{"name":"CERT_ISSUER", "description":"allowed to issue certificates"}'
        );
        _defineRole(
            ASSET_ISSUER_ROLE,
            '{"name":"ASSET_ISSUER", "description":"allowed to issue assets"}'
        );
        _defineRole(
            MOVEMENT_ISSUER_ROLE,
            '{"name":"MOVEMENT_ISSUER", "description":"allowed to issue movements"}'
        );

        // define Orbyc agent
        _defineAgent(
            _msgSender(),
            ORBYC_AGENT,
            '{"name":"Orbyc Default Admin"}'
        );

        // grant roles to orbyc agent
        _grantRole(ORBYC_AGENT, ADMIN_ROLE);
        _grantRole(ORBYC_AGENT, EDITOR_ROLE);
        _grantRole(ORBYC_AGENT, PROVIDER_ROLE);
        _grantRole(ORBYC_AGENT, CERT_ISSUER_ROLE);
        _grantRole(ORBYC_AGENT, ASSET_ISSUER_ROLE);
        _grantRole(ORBYC_AGENT, MOVEMENT_ISSUER_ROLE);
    }

    /**
     * ERC245
     */

    function name()
        public
        view
        override(ERC245, ERC423)
        returns (string memory)
    {
        return ERC245.name();
    }

    function idOf(address signer)
        public
        view
        override(ERC245, ERC423)
        returns (uint64)
    {
        return ERC423.idOf(signer);
    }

    /**
     * @dev Multisig implementation helper for the transfer method.
     *
     * Each ownership transfer must be approved for the `owner` and
     * the `issuer` of the address.
     */
    function getSigners(uint256 asset)
        internal
        virtual
        returns (uint64[] memory)
    {
        uint64 owner;
        uint64 issuer;
        (owner, issuer, , , ) = assetInfo(asset);

        uint64[] memory signers = new uint64[](2);
        signers[0] = owner;
        signers[1] = issuer;

        return signers;
    }

    /**
     * @dev Multisig implementation for the transfer method.
     *
     * This method will use the virtual `getSigners` to create a transaction proposal.
     * All the signers must perform a transfer in order to asset ownership change.
     */
    function transferOwnership(uint256 asset, uint64 recipient)
        public
        override
        whenNotPaused
        returns (bool)
    {
        // get signer id
        uint64 signer = idOf(_msgSender());

        // get transaction valid signers
        uint64[] memory signers = getSigners(asset);

        // validate signer
        require(
            Array.contains(getSigners(asset), signer),
            "Error: signer not valid for this asset"
        );

        // get transaction from storage
        Multi.Transfer storage transaction = _transfers[asset];

        // validate signer has not sign
        require(
            !transaction.signatures[signer],
            "Error: signer already sign this transfer"
        );

        // reset transfer if recipient has change
        if (transaction.recipient != recipient) {
            transaction.recipient = recipient;
            transaction.asset = asset;

            for (uint256 i = 0; i < signers.length; i++) {
                delete transaction.signatures[signers[i]];
            }
        }

        // update signers signature as approval
        transaction.signatures[signer] = true;

        // return if there are missing signatures
        for (uint256 i = 0; i < signers.length; i++) {
            if (!transaction.signatures[signers[i]]) {
                // emit {Transfer} event
                emit TransferOwnership(asset, signer, recipient, _msgSender());
                // return sign succeed
                return true;
            }
        }

        // clear transaction recipient
        transaction.recipient = 0;

        // return the parent transfer execution
        return super.transferOwnership(asset, recipient);
    }

    function issueCertificate(uint256 id, string memory metadata)
        public
        override
        onlyRole(CERT_ISSUER_ROLE)
        whenNotPaused
        returns (bool)
    {
        return super.issueCertificate(id, metadata);
    }

    function issueAsset(
        uint256 id,
        uint64 owner,
        uint256 co2,
        uint256 cert,
        string memory metadata
    ) public override onlyRole(ASSET_ISSUER_ROLE) whenNotPaused returns (bool) {
        return super.issueAsset(id, owner, co2, cert, metadata);
    }

    function issueMovement(
        uint256 id,
        string memory lat,
        string memory lng,
        uint256 co2,
        uint256 cert,
        string memory metadata
    )
        public
        override
        onlyRole(MOVEMENT_ISSUER_ROLE)
        whenNotPaused
        returns (bool)
    {
        return super.issueMovement(id, lat, lng, co2, cert, metadata);
    }

    function addMovements(uint256 id, uint256[] memory transports)
        public
        virtual
        override
        onlyIssuer(id)
        whenNotPaused
        returns (bool)
    {
        return super.addMovements(id, transports);
    }

    function addParents(
        uint256 id,
        uint256[] memory parents,
        uint16[] memory composition
    ) public override onlyIssuer(id) whenNotPaused returns (bool) {
        return super.addParents(id, parents, composition);
    }

    /**
     * ERC423
     */

    function defineRole(uint64 role, string memory info)
        public
        override
        onlyRole(ADMIN_ROLE)
        returns (bool)
    {
        return super.defineRole(role, info);
    }

    function defineAgent(
        address agent,
        uint64 id,
        string memory info
    ) public override onlyRole(EDITOR_ROLE) validAddress(agent) returns (bool) {
        return super.defineAgent(agent, id, info);
    }

    function grantRole(uint64 id, uint64 role)
        public
        override
        onlyRole(ADMIN_ROLE)
        nonNull(id)
        returns (bool)
    {
        return super.grantRole(id, role);
    }

    function revokeRole(uint64 id, uint64 role)
        public
        override
        onlyRole(ADMIN_ROLE)
        nonNull(id)
        returns (bool)
    {
        return super.revokeRole(id, role);
    }

    /**
     * @dev Banns an address making it impossible to opperate again
     */
    function bannAddress(address agent)
        public
        onlyRole(ADMIN_ROLE)
        returns (bool)
    {
        return _defineAgent(agent, NULL_AGENT, "...");
    }

    /**
     * Pausable
     */

    function Pause() public onlyRole(ADMIN_ROLE) {
        _pause();
    }

    function Unpause() public onlyRole(ADMIN_ROLE) {
        _unpause();
    }
}
