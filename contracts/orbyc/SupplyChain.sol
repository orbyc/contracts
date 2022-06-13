// SPDX-License-Identifier: MIT
// Orbyc Contracts v1.0.0 (Supply.sol)

pragma solidity ^0.8.0;

import {Pausable} from "@openzeppelin/contracts/security/Pausable.sol";
import {Context} from "@openzeppelin/contracts/utils/Context.sol";
import {ERC721} from "@openzeppelin/contracts/token/ERC721/ERC721.sol";

import {ERC423} from "../security/ERC423/ERC423.sol";
import {ERC245} from "../tokens/ERC245/ERC245.sol";
import {Array} from "../utils/Collections.sol";

struct MultiTransfer {
    // asset id
    uint256 assetId;
    // receiver address
    address receiver;
    // signer address => hasSigned
    mapping(address => bool) signatures;
}

contract SupplyChain is ERC245, ERC423, ERC721, Pausable {
    /**
     * Roles
     */
    uint64 public ADMIN_ROLE = 0x01 << 1;
    uint64 public PROVIDER_ROLE = 0x01 << 2;
    uint64 public ASSET_ISSUER_ROLE = 0x01 << 3;
    uint64 public CERT_ISSUER_ROLE = 0x01 << 4;
    uint64 public MOVEMENT_ISSUER_ROLE = 0x01 << 5;

    /**
     * Agents
     */
    address public ORBYC_ACCOUNT = 0x1001bdC1076D31806222A2e826FB6819FFb3b809;

    /**
     * Privates
     */
    // mapping from asset id to ownership multitransfer
    mapping(uint256 => MultiTransfer) private _transfers;

    /**
     * Modifiers
     */
    modifier onlyAssetIssuer(uint256 assetId) {
        (, address issuer, , , ) = assetInfo(assetId);
        address account = accountOf(_msgSender());
        require(account == issuer, "Error: account is not the asset issuer");
        _;
    }

    modifier onlyMovementIssuer(uint256 moveId) {
        (address issuer, , , ) = movementInfo(moveId);
        address account = accountOf(_msgSender());
        require(account == issuer, "Error: account is not the asset issuer");
        _;
    }

    /**
     * Events
     */
    /**
     * @dev Emitted when `asset` was sent by `signer`
     */
    event SignatureReceived(
        address indexed from,
        address indexed to,
        uint256 indexed tokenId
    );

    /**
     * Constructor
     */
    constructor(string memory name_, string memory symbol_)
        ERC721(name_, symbol_)
        ERC245(name_)
        ERC423(name_)
    {
        // define roles
        _defineRole(
            ADMIN_ROLE,
            '{"name":"ADMIN", "description":"allowed to issue and assign roles"}'
        );
        _defineRole(
            PROVIDER_ROLE,
            '{"name":"PROVIDER", "description":"allowed to issue agents with client role"}'
        );

        _defineRole(
            ASSET_ISSUER_ROLE,
            '{"name":"ASSET_ISSUER", "description":"allowed to issue assets"}'
        );
        _defineRole(
            CERT_ISSUER_ROLE,
            '{"name":"CERT_ISSUER", "description":"allowed to issue certificates"}'
        );
        _defineRole(
            MOVEMENT_ISSUER_ROLE,
            '{"name":"MOVEMENT_ISSUER", "description":"allowed to issue movements"}'
        );

        // define Orbyc agent
        _defineAgent(
            _msgSender(),
            ORBYC_ACCOUNT,
            '{"name":"Orbyc Default Admin"}'
        );

        // grant roles to orbyc agent
        _grantRole(ORBYC_ACCOUNT, ADMIN_ROLE);
        _grantRole(ORBYC_ACCOUNT, PROVIDER_ROLE);
        _grantRole(ORBYC_ACCOUNT, ASSET_ISSUER_ROLE);
        _grantRole(ORBYC_ACCOUNT, CERT_ISSUER_ROLE);
        _grantRole(ORBYC_ACCOUNT, MOVEMENT_ISSUER_ROLE);
    }

    function name()
        public
        view
        override(ERC245, ERC423, ERC721)
        returns (string memory)
    {
        return ERC245.name();
    }

    /**
     * ERC721
     */
    function balanceOf(address owner)
        public
        view
        virtual
        override
        returns (uint256)
    {
        address account = accountOf(owner);
        return super.balanceOf(account);
    }

    function _baseURI() internal view virtual override returns (string memory) {
        return "https://orbyc.com/#/";
    }

    function transferFrom(
        address from,
        address to,
        uint256 tokenId
    ) public virtual override(ERC245, ERC721) whenNotPaused {
        address sender = accountOf(from);
        address receiver = accountOf(to);

        if (!_requireMultisig(sender, receiver, tokenId)) {
            return;
        }

        ERC245.transferFrom(sender, receiver, tokenId);
        ERC721.transferFrom(sender, receiver, tokenId);
    }

    function _isApprovedOrOwner(address spender, uint256 tokenId)
        internal
        view
        virtual
        override
        returns (bool)
    {
        address account = accountOf(spender);
        return super._isApprovedOrOwner(account, tokenId);
    }

    /**
     * @dev Multisig implementation helper for the transfer method.
     *
     * Each ownership transfer must be approved for the `owner` and
     * the `issuer` of the address.
     */
    function _getSigners(uint256 assetId)
        internal
        virtual
        returns (address[] memory)
    {
        (address owner, address issuer, , , ) = assetInfo(assetId);

        address[] memory signers = new address[](2);
        signers[0] = owner;
        signers[1] = issuer;

        return signers;
    }

    /**
     * @dev Multisig implementation for the transfer method.
     *
     * This method will use the virtual `getSigners` to create a transaction proposal.
     * All the signers must perform a transfer in order to asset ownership change.
     *
     * emits {SignatureReceived} when transfer is performed but not completed
     */
    function _requireMultisig(
        address sender,
        address receiver,
        uint256 assetId
    ) internal virtual returns (bool) {
        address[] memory signers = _getSigners(assetId);

        // check sender is one of the signers
        require(
            Array.contains(signers, sender),
            "Error: signer not valid for this asset"
        );

        // get transaction from storage
        MultiTransfer storage transaction = _transfers[assetId];

        // validate signer has not sign yet
        require(
            !transaction.signatures[sender],
            "Error: signer already sign this transfer"
        );

        // reset transfer if receiver has change
        if (transaction.receiver != receiver) {
            transaction.receiver = receiver;
            transaction.assetId = assetId;
            // reset all signatures
            for (uint256 i = 0; i < signers.length; i++) {
                delete transaction.signatures[signers[i]];
            }
        }

        // update signers signature as approval
        transaction.signatures[sender] = true;

        // return if there are missing signatures
        for (uint256 i = 0; i < signers.length; i++) {
            if (!transaction.signatures[signers[i]]) {
                emit SignatureReceived(sender, receiver, assetId);

                // return false as multisig is not complete
                return false;
            }
        }

        // return true as multisig is complete
        return true;
    }

    /**
     * ERC245
     */
    function issueAsset(
        uint256 assetId,
        address owner,
        uint64 co2e,
        uint256 certId,
        string memory metadata_
    ) public override onlyRole(ASSET_ISSUER_ROLE) whenNotPaused returns (bool) {
        address issuer = accountOf(_msgSender());
        _mint(owner, assetId);
        return
            super._issueAsset(assetId, owner, issuer, co2e, certId, metadata_);
    }

    function issueCertificate(uint256 certId, string calldata metadata_)
        public
        override
        onlyRole(CERT_ISSUER_ROLE)
        whenNotPaused
        returns (bool)
    {
        address issuer = accountOf(_msgSender());
        return super._issueCertificate(certId, issuer, metadata_);
    }

    function issueMovement(
        uint256 moveId,
        uint64 co2e,
        uint256 certId,
        string memory metadata_
    )
        public
        override
        onlyRole(MOVEMENT_ISSUER_ROLE)
        whenNotPaused
        returns (bool)
    {
        address issuer = accountOf(_msgSender());
        return super._issueMovement(moveId, issuer, co2e, certId, metadata_);
    }

    function addAssetCertificates(
        uint256 assetId,
        uint256[] memory certificates
    ) public virtual override onlyAssetIssuer(assetId) returns (bool) {
        return super.addAssetCertificates(assetId, certificates);
    }

    function addMovementCertificates(
        uint256 moveId,
        uint256[] memory certificates
    ) public virtual override onlyMovementIssuer(moveId) returns (bool) {
        return super.addMovementCertificates(moveId, certificates);
    }

    function addMovements(uint256 assetId, uint256[] memory movements)
        public
        virtual
        override
        onlyAssetIssuer(assetId)
        whenNotPaused
        returns (bool)
    {
        return super.addMovements(assetId, movements);
    }

    function addParents(
        uint256 assetId,
        uint256[] memory parents,
        uint16[] memory composition
    ) public override onlyAssetIssuer(assetId) whenNotPaused returns (bool) {
        return super.addParents(assetId, parents, composition);
    }

    /**
     * ERC423
     */
    function defineAgent(
        address agent,
        address account,
        string memory metadata_
    ) public override onlyRole(ADMIN_ROLE) validAgent(agent) returns (bool) {
        return super.defineAgent(agent, account, metadata_);
    }

    function removeAgent(address agent)
        public
        override
        onlyRole(ADMIN_ROLE)
        returns (bool)
    {
        return super.removeAgent(agent);
    }

    function defineRole(uint256 role, string memory metadata_)
        public
        override
        onlyRole(ADMIN_ROLE)
        returns (bool)
    {
        return super.defineRole(role, metadata_);
    }

    function grantRole(address account, uint256 role)
        public
        override
        onlyRole(ADMIN_ROLE)
        returns (bool)
    {
        return super.grantRole(account, role);
    }

    function revokeRole(address account, uint256 role)
        public
        override
        onlyRole(ADMIN_ROLE)
        returns (bool)
    {
        return super.revokeRole(account, role);
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
