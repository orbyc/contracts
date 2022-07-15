// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sdk

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SupplyChainMetaData contains all meta data concerning the SupplyChain contract.
var SupplyChainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"AgentDefined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"AgentRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"certId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"AssetCertificateAssigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"AssetIssued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"certId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"CertIssued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"moveId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"MovementAssigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"certId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"moveId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"MovementCertificateAssigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"moveId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"MovementIssued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"parentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"ParentAssigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"RoleDefined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"SignatureReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"TransferAsset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ASSET_ISSUER_ROLE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CERT_ISSUER_ROLE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MOVEMENT_ISSUER_ROLE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORBYC_ACCOUNT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROVIDER_ROLE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"accountInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"accountOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"certificates\",\"type\":\"uint256[]\"}],\"name\":\"addAssetCertificates\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"moveId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"certificates\",\"type\":\"uint256[]\"}],\"name\":\"addMovementCertificates\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"movements\",\"type\":\"uint256[]\"}],\"name\":\"addMovements\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"parents\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"composition\",\"type\":\"uint16[]\"}],\"name\":\"addParents\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"assetCertificates\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"assetComposition\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"\",\"type\":\"uint16[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"assetInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"assetTraceability\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"certId\",\"type\":\"uint256\"}],\"name\":\"certificateInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadata_\",\"type\":\"string\"}],\"name\":\"defineAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata_\",\"type\":\"string\"}],\"name\":\"defineRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"grantRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"co2e\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"certId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata_\",\"type\":\"string\"}],\"name\":\"issueAsset\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"certId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata_\",\"type\":\"string\"}],\"name\":\"issueCertificate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"moveId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"co2e\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"certId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata_\",\"type\":\"string\"}],\"name\":\"issueMovement\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"moveId\",\"type\":\"uint256\"}],\"name\":\"movementCertificates\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"moveId\",\"type\":\"uint256\"}],\"name\":\"movementInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"removeAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"revokeRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"roleInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SupplyChainABI is the input ABI used to generate the binding from.
// Deprecated: Use SupplyChainMetaData.ABI instead.
var SupplyChainABI = SupplyChainMetaData.ABI

// SupplyChain is an auto generated Go binding around an Ethereum contract.
type SupplyChain struct {
	SupplyChainCaller     // Read-only binding to the contract
	SupplyChainTransactor // Write-only binding to the contract
	SupplyChainFilterer   // Log filterer for contract events
}

// SupplyChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type SupplyChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupplyChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SupplyChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupplyChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SupplyChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupplyChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SupplyChainSession struct {
	Contract     *SupplyChain      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SupplyChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SupplyChainCallerSession struct {
	Contract *SupplyChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SupplyChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SupplyChainTransactorSession struct {
	Contract     *SupplyChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SupplyChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type SupplyChainRaw struct {
	Contract *SupplyChain // Generic contract binding to access the raw methods on
}

// SupplyChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SupplyChainCallerRaw struct {
	Contract *SupplyChainCaller // Generic read-only contract binding to access the raw methods on
}

// SupplyChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SupplyChainTransactorRaw struct {
	Contract *SupplyChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSupplyChain creates a new instance of SupplyChain, bound to a specific deployed contract.
func NewSupplyChain(address common.Address, backend bind.ContractBackend) (*SupplyChain, error) {
	contract, err := bindSupplyChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SupplyChain{SupplyChainCaller: SupplyChainCaller{contract: contract}, SupplyChainTransactor: SupplyChainTransactor{contract: contract}, SupplyChainFilterer: SupplyChainFilterer{contract: contract}}, nil
}

// NewSupplyChainCaller creates a new read-only instance of SupplyChain, bound to a specific deployed contract.
func NewSupplyChainCaller(address common.Address, caller bind.ContractCaller) (*SupplyChainCaller, error) {
	contract, err := bindSupplyChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SupplyChainCaller{contract: contract}, nil
}

// NewSupplyChainTransactor creates a new write-only instance of SupplyChain, bound to a specific deployed contract.
func NewSupplyChainTransactor(address common.Address, transactor bind.ContractTransactor) (*SupplyChainTransactor, error) {
	contract, err := bindSupplyChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SupplyChainTransactor{contract: contract}, nil
}

// NewSupplyChainFilterer creates a new log filterer instance of SupplyChain, bound to a specific deployed contract.
func NewSupplyChainFilterer(address common.Address, filterer bind.ContractFilterer) (*SupplyChainFilterer, error) {
	contract, err := bindSupplyChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SupplyChainFilterer{contract: contract}, nil
}

// bindSupplyChain binds a generic wrapper to an already deployed contract.
func bindSupplyChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SupplyChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SupplyChain *SupplyChainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SupplyChain.Contract.SupplyChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SupplyChain *SupplyChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SupplyChain.Contract.SupplyChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SupplyChain *SupplyChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SupplyChain.Contract.SupplyChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SupplyChain *SupplyChainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SupplyChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SupplyChain *SupplyChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SupplyChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SupplyChain *SupplyChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SupplyChain.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(uint64)
func (_SupplyChain *SupplyChainCaller) ADMINROLE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SupplyChain.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(uint64)
func (_SupplyChain *SupplyChainSession) ADMINROLE() (uint64, error) {
	return _SupplyChain.Contract.ADMINROLE(&_SupplyChain.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(uint64)
func (_SupplyChain *SupplyChainCallerSession) ADMINROLE() (uint64, error) {
	return _SupplyChain.Contract.ADMINROLE(&_SupplyChain.CallOpts)
}

// ASSETISSUERROLE is a free data retrieval call binding the contract method 0xd8f9ce24.
//
// Solidity: function ASSET_ISSUER_ROLE() view returns(uint64)
func (_SupplyChain *SupplyChainCaller) ASSETISSUERROLE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SupplyChain.contract.Call(opts, &out, "ASSET_ISSUER_ROLE")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ASSETISSUERROLE is a free data retrieval call binding the contract method 0xd8f9ce24.
//
// Solidity: function ASSET_ISSUER_ROLE() view returns(uint64)
func (_SupplyChain *SupplyChainSession) ASSETISSUERROLE() (uint64, error) {
	return _SupplyChain.Contract.ASSETISSUERROLE(&_SupplyChain.CallOpts)
}

// ASSETISSUERROLE is a free data retrieval call binding the contract method 0xd8f9ce24.
//
// Solidity: function ASSET_ISSUER_ROLE() view returns(uint64)
func (_SupplyChain *SupplyChainCallerSession) ASSETISSUERROLE() (uint64, error) {
	return _SupplyChain.Contract.ASSETISSUERROLE(&_SupplyChain.CallOpts)
}

// CERTISSUERROLE is a free data retrieval call binding the contract method 0x714eb854.
//
// Solidity: function CERT_ISSUER_ROLE() view returns(uint64)
func (_SupplyChain *SupplyChainCaller) CERTISSUERROLE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SupplyChain.contract.Call(opts, &out, "CERT_ISSUER_ROLE")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CERTISSUERROLE is a free data retrieval call binding the contract method 0x714eb854.
//
// Solidity: function CERT_ISSUER_ROLE() view returns(uint64)
func (_SupplyChain *SupplyChainSession) CERTISSUERROLE() (uint64, error) {
	return _SupplyChain.Contract.CERTISSUERROLE(&_SupplyChain.CallOpts)
}

// CERTISSUERROLE is a free data retrieval call binding the contract method 0x714eb854.
//
// Solidity: function CERT_ISSUER_ROLE() view returns(uint64)
func (_SupplyChain *SupplyChainCallerSession) CERTISSUERROLE() (uint64, error) {
	return _SupplyChain.Contract.CERTISSUERROLE(&_SupplyChain.CallOpts)
}

// MOVEMENTISSUERROLE is a free data retrieval call binding the contract method 0x53dcbaaf.
//
// Solidity: function MOVEMENT_ISSUER_ROLE() view returns(uint64)
func (_SupplyChain *SupplyChainCaller) MOVEMENTISSUERROLE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SupplyChain.contract.Call(opts, &out, "MOVEMENT_ISSUER_ROLE")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MOVEMENTISSUERROLE is a free data retrieval call binding the contract method 0x53dcbaaf.
//
// Solidity: function MOVEMENT_ISSUER_ROLE() view returns(uint64)
func (_SupplyChain *SupplyChainSession) MOVEMENTISSUERROLE() (uint64, error) {
	return _SupplyChain.Contract.MOVEMENTISSUERROLE(&_SupplyChain.CallOpts)
}

// MOVEMENTISSUERROLE is a free data retrieval call binding the contract method 0x53dcbaaf.
//
// Solidity: function MOVEMENT_ISSUER_ROLE() view returns(uint64)
func (_SupplyChain *SupplyChainCallerSession) MOVEMENTISSUERROLE() (uint64, error) {
	return _SupplyChain.Contract.MOVEMENTISSUERROLE(&_SupplyChain.CallOpts)
}

// ORBYCACCOUNT is a free data retrieval call binding the contract method 0x8fa07302.
//
// Solidity: function ORBYC_ACCOUNT() view returns(address)
func (_SupplyChain *SupplyChainCaller) ORBYCACCOUNT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SupplyChain.contract.Call(opts, &out, "ORBYC_ACCOUNT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ORBYCACCOUNT is a free data retrieval call binding the contract method 0x8fa07302.
//
// Solidity: function ORBYC_ACCOUNT() view returns(address)
func (_SupplyChain *SupplyChainSession) ORBYCACCOUNT() (common.Address, error) {
	return _SupplyChain.Contract.ORBYCACCOUNT(&_SupplyChain.CallOpts)
}

// ORBYCACCOUNT is a free data retrieval call binding the contract method 0x8fa07302.
//
// Solidity: function ORBYC_ACCOUNT() view returns(address)
func (_SupplyChain *SupplyChainCallerSession) ORBYCACCOUNT() (common.Address, error) {
	return _SupplyChain.Contract.ORBYCACCOUNT(&_SupplyChain.CallOpts)
}

// PROVIDERROLE is a free data retrieval call binding the contract method 0x24c20a34.
//
// Solidity: function PROVIDER_ROLE() view returns(uint64)
func (_SupplyChain *SupplyChainCaller) PROVIDERROLE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SupplyChain.contract.Call(opts, &out, "PROVIDER_ROLE")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PROVIDERROLE is a free data retrieval call binding the contract method 0x24c20a34.
//
// Solidity: function PROVIDER_ROLE() view returns(uint64)
func (_SupplyChain *SupplyChainSession) PROVIDERROLE() (uint64, error) {
	return _SupplyChain.Contract.PROVIDERROLE(&_SupplyChain.CallOpts)
}

// PROVIDERROLE is a free data retrieval call binding the contract method 0x24c20a34.
//
// Solidity: function PROVIDER_ROLE() view returns(uint64)
func (_SupplyChain *SupplyChainCallerSession) PROVIDERROLE() (uint64, error) {
	return _SupplyChain.Contract.PROVIDERROLE(&_SupplyChain.CallOpts)
}

// AccountInfo is a free data retrieval call binding the contract method 0xa7310b58.
//
// Solidity: function accountInfo(address account) view returns(string)
func (_SupplyChain *SupplyChainCaller) AccountInfo(opts *bind.CallOpts, account common.Address) (string, error) {
	var out []interface{}
	err := _SupplyChain.contract.Call(opts, &out, "accountInfo", account)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AccountInfo is a free data retrieval call binding the contract method 0xa7310b58.
//
// Solidity: function accountInfo(address account) view returns(string)
func (_SupplyChain *SupplyChainSession) AccountInfo(account common.Address) (string, error) {
	return _SupplyChain.Contract.AccountInfo(&_SupplyChain.CallOpts, account)
}

// AccountInfo is a free data retrieval call binding the contract method 0xa7310b58.
//
// Solidity: function accountInfo(address account) view returns(string)
func (_SupplyChain *SupplyChainCallerSession) AccountInfo(account common.Address) (string, error) {
	return _SupplyChain.Contract.AccountInfo(&_SupplyChain.CallOpts, account)
}

// AccountOf is a free data retrieval call binding the contract method 0x8086b8ba.
//
// Solidity: function accountOf(address agent) view returns(address)
func (_SupplyChain *SupplyChainCaller) AccountOf(opts *bind.CallOpts, agent common.Address) (common.Address, error) {
	var out []interface{}
	err := _SupplyChain.contract.Call(opts, &out, "accountOf", agent)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountOf is a free data retrieval call binding the contract method 0x8086b8ba.
//
// Solidity: function accountOf(address agent) view returns(address)
func (_SupplyChain *SupplyChainSession) AccountOf(agent common.Address) (common.Address, error) {
	return _SupplyChain.Contract.AccountOf(&_SupplyChain.CallOpts, agent)
}

// AccountOf is a free data retrieval call binding the contract method 0x8086b8ba.
//
// Solidity: function accountOf(address agent) view returns(address)
func (_SupplyChain *SupplyChainCallerSession) AccountOf(agent common.Address) (common.Address, error) {
	return _SupplyChain.Contract.AccountOf(&_SupplyChain.CallOpts, agent)
}

// AssetCertificates is a free data retrieval call binding the contract method 0xc22d9980.
//
// Solidity: function assetCertificates(uint256 assetId) view returns(uint256[])
func (_SupplyChain *SupplyChainCaller) AssetCertificates(opts *bind.CallOpts, assetId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _SupplyChain.contract.Call(opts, &out, "assetCertificates", assetId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AssetCertificates is a free data retrieval call binding the contract method 0xc22d9980.
//
// Solidity: function assetCertificates(uint256 assetId) view returns(uint256[])
func (_SupplyChain *SupplyChainSession) AssetCertificates(assetId *big.Int) ([]*big.Int, error) {
	return _SupplyChain.Contract.AssetCertificates(&_SupplyChain.CallOpts, assetId)
}

// AssetCertificates is a free data retrieval call binding the contract method 0xc22d9980.
//
// Solidity: function assetCertificates(uint256 assetId) view returns(uint256[])
func (_SupplyChain *SupplyChainCallerSession) AssetCertificates(assetId *big.Int) ([]*big.Int, error) {
	return _SupplyChain.Contract.AssetCertificates(&_SupplyChain.CallOpts, assetId)
}

// AssetComposition is a free data retrieval call binding the contract method 0x2192bc2a.
//
// Solidity: function assetComposition(uint256 assetId) view returns(uint256[], uint16[])
func (_SupplyChain *SupplyChainCaller) AssetComposition(opts *bind.CallOpts, assetId *big.Int) ([]*big.Int, []uint16, error) {
	var out []interface{}
	err := _SupplyChain.contract.Call(opts, &out, "assetComposition", assetId)

	if err != nil {
		return *new([]*big.Int), *new([]uint16), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]uint16)).(*[]uint16)

	return out0, out1, err

}

// AssetComposition is a free data retrieval call binding the contract method 0x2192bc2a.
//
// Solidity: function assetComposition(uint256 assetId) view returns(uint256[], uint16[])
func (_SupplyChain *SupplyChainSession) AssetComposition(assetId *big.Int) ([]*big.Int, []uint16, error) {
	return _SupplyChain.Contract.AssetComposition(&_SupplyChain.CallOpts, assetId)
}

// AssetComposition is a free data retrieval call binding the contract method 0x2192bc2a.
//
// Solidity: function assetComposition(uint256 assetId) view returns(uint256[], uint16[])
func (_SupplyChain *SupplyChainCallerSession) AssetComposition(assetId *big.Int) ([]*big.Int, []uint16, error) {
	return _SupplyChain.Contract.AssetComposition(&_SupplyChain.CallOpts, assetId)
}

// AssetInfo is a free data retrieval call binding the contract method 0xa879fcbb.
//
// Solidity: function assetInfo(uint256 assetId) view returns(address, address, uint64, uint256, string)
func (_SupplyChain *SupplyChainCaller) AssetInfo(opts *bind.CallOpts, assetId *big.Int) (common.Address, common.Address, uint64, *big.Int, string, error) {
	var out []interface{}
	err := _SupplyChain.contract.Call(opts, &out, "assetInfo", assetId)

	if err != nil {
		return *new(common.Address), *new(common.Address), *new(uint64), *new(*big.Int), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(string)).(*string)

	return out0, out1, out2, out3, out4, err

}

// AssetInfo is a free data retrieval call binding the contract method 0xa879fcbb.
//
// Solidity: function assetInfo(uint256 assetId) view returns(address, address, uint64, uint256, string)
func (_SupplyChain *SupplyChainSession) AssetInfo(assetId *big.Int) (common.Address, common.Address, uint64, *big.Int, string, error) {
	return _SupplyChain.Contract.AssetInfo(&_SupplyChain.CallOpts, assetId)
}

// AssetInfo is a free data retrieval call binding the contract method 0xa879fcbb.
//
// Solidity: function assetInfo(uint256 assetId) view returns(address, address, uint64, uint256, string)
func (_SupplyChain *SupplyChainCallerSession) AssetInfo(assetId *big.Int) (common.Address, common.Address, uint64, *big.Int, string, error) {
	return _SupplyChain.Contract.AssetInfo(&_SupplyChain.CallOpts, assetId)
}

// AssetTraceability is a free data retrieval call binding the contract method 0xf6863e83.
//
// Solidity: function assetTraceability(uint256 assetId) view returns(uint256[])
func (_SupplyChain *SupplyChainCaller) AssetTraceability(opts *bind.CallOpts, assetId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _SupplyChain.contract.Call(opts, &out, "assetTraceability", assetId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AssetTraceability is a free data retrieval call binding the contract method 0xf6863e83.
//
// Solidity: function assetTraceability(uint256 assetId) view returns(uint256[])
func (_SupplyChain *SupplyChainSession) AssetTraceability(assetId *big.Int) ([]*big.Int, error) {
	return _SupplyChain.Contract.AssetTraceability(&_SupplyChain.CallOpts, assetId)
}

// AssetTraceability is a free data retrieval call binding the contract method 0xf6863e83.
//
// Solidity: function assetTraceability(uint256 assetId) view returns(uint256[])
func (_SupplyChain *SupplyChainCallerSession) AssetTraceability(assetId *big.Int) ([]*big.Int, error) {
	return _SupplyChain.Contract.AssetTraceability(&_SupplyChain.CallOpts, assetId)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SupplyChain *SupplyChainCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SupplyChain.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SupplyChain *SupplyChainSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SupplyChain.Contract.BalanceOf(&_SupplyChain.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SupplyChain *SupplyChainCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SupplyChain.Contract.BalanceOf(&_SupplyChain.CallOpts, owner)
}

// CertificateInfo is a free data retrieval call binding the contract method 0xba0a9f50.
//
// Solidity: function certificateInfo(uint256 certId) view returns(address, string)
func (_SupplyChain *SupplyChainCaller) CertificateInfo(opts *bind.CallOpts, certId *big.Int) (common.Address, string, error) {
	var out []interface{}
	err := _SupplyChain.contract.Call(opts, &out, "certificateInfo", certId)

	if err != nil {
		return *new(common.Address), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// CertificateInfo is a free data retrieval call binding the contract method 0xba0a9f50.
//
// Solidity: function certificateInfo(uint256 certId) view returns(address, string)
func (_SupplyChain *SupplyChainSession) CertificateInfo(certId *big.Int) (common.Address, string, error) {
	return _SupplyChain.Contract.CertificateInfo(&_SupplyChain.CallOpts, certId)
}

// CertificateInfo is a free data retrieval call binding the contract method 0xba0a9f50.
//
// Solidity: function certificateInfo(uint256 certId) view returns(address, string)
func (_SupplyChain *SupplyChainCallerSession) CertificateInfo(certId *big.Int) (common.Address, string, error) {
	return _SupplyChain.Contract.CertificateInfo(&_SupplyChain.CallOpts, certId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SupplyChain *SupplyChainCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SupplyChain.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SupplyChain *SupplyChainSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SupplyChain.Contract.GetApproved(&_SupplyChain.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SupplyChain *SupplyChainCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SupplyChain.Contract.GetApproved(&_SupplyChain.CallOpts, tokenId)
}

// HasRole is a free data retrieval call binding the contract method 0x5c97f4a2.
//
// Solidity: function hasRole(address account, uint256 role) view returns(bool)
func (_SupplyChain *SupplyChainCaller) HasRole(opts *bind.CallOpts, account common.Address, role *big.Int) (bool, error) {
	var out []interface{}
	err := _SupplyChain.contract.Call(opts, &out, "hasRole", account, role)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x5c97f4a2.
//
// Solidity: function hasRole(address account, uint256 role) view returns(bool)
func (_SupplyChain *SupplyChainSession) HasRole(account common.Address, role *big.Int) (bool, error) {
	return _SupplyChain.Contract.HasRole(&_SupplyChain.CallOpts, account, role)
}

// HasRole is a free data retrieval call binding the contract method 0x5c97f4a2.
//
// Solidity: function hasRole(address account, uint256 role) view returns(bool)
func (_SupplyChain *SupplyChainCallerSession) HasRole(account common.Address, role *big.Int) (bool, error) {
	return _SupplyChain.Contract.HasRole(&_SupplyChain.CallOpts, account, role)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SupplyChain *SupplyChainCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _SupplyChain.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SupplyChain *SupplyChainSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SupplyChain.Contract.IsApprovedForAll(&_SupplyChain.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SupplyChain *SupplyChainCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SupplyChain.Contract.IsApprovedForAll(&_SupplyChain.CallOpts, owner, operator)
}

// MovementCertificates is a free data retrieval call binding the contract method 0xd1ac4d65.
//
// Solidity: function movementCertificates(uint256 moveId) view returns(uint256[])
func (_SupplyChain *SupplyChainCaller) MovementCertificates(opts *bind.CallOpts, moveId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _SupplyChain.contract.Call(opts, &out, "movementCertificates", moveId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// MovementCertificates is a free data retrieval call binding the contract method 0xd1ac4d65.
//
// Solidity: function movementCertificates(uint256 moveId) view returns(uint256[])
func (_SupplyChain *SupplyChainSession) MovementCertificates(moveId *big.Int) ([]*big.Int, error) {
	return _SupplyChain.Contract.MovementCertificates(&_SupplyChain.CallOpts, moveId)
}

// MovementCertificates is a free data retrieval call binding the contract method 0xd1ac4d65.
//
// Solidity: function movementCertificates(uint256 moveId) view returns(uint256[])
func (_SupplyChain *SupplyChainCallerSession) MovementCertificates(moveId *big.Int) ([]*big.Int, error) {
	return _SupplyChain.Contract.MovementCertificates(&_SupplyChain.CallOpts, moveId)
}

// MovementInfo is a free data retrieval call binding the contract method 0x3586ed2f.
//
// Solidity: function movementInfo(uint256 moveId) view returns(address, uint64, uint256, string)
func (_SupplyChain *SupplyChainCaller) MovementInfo(opts *bind.CallOpts, moveId *big.Int) (common.Address, uint64, *big.Int, string, error) {
	var out []interface{}
	err := _SupplyChain.contract.Call(opts, &out, "movementInfo", moveId)

	if err != nil {
		return *new(common.Address), *new(uint64), *new(*big.Int), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)

	return out0, out1, out2, out3, err

}

// MovementInfo is a free data retrieval call binding the contract method 0x3586ed2f.
//
// Solidity: function movementInfo(uint256 moveId) view returns(address, uint64, uint256, string)
func (_SupplyChain *SupplyChainSession) MovementInfo(moveId *big.Int) (common.Address, uint64, *big.Int, string, error) {
	return _SupplyChain.Contract.MovementInfo(&_SupplyChain.CallOpts, moveId)
}

// MovementInfo is a free data retrieval call binding the contract method 0x3586ed2f.
//
// Solidity: function movementInfo(uint256 moveId) view returns(address, uint64, uint256, string)
func (_SupplyChain *SupplyChainCallerSession) MovementInfo(moveId *big.Int) (common.Address, uint64, *big.Int, string, error) {
	return _SupplyChain.Contract.MovementInfo(&_SupplyChain.CallOpts, moveId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SupplyChain *SupplyChainCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SupplyChain.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SupplyChain *SupplyChainSession) Name() (string, error) {
	return _SupplyChain.Contract.Name(&_SupplyChain.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SupplyChain *SupplyChainCallerSession) Name() (string, error) {
	return _SupplyChain.Contract.Name(&_SupplyChain.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SupplyChain *SupplyChainCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SupplyChain.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SupplyChain *SupplyChainSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SupplyChain.Contract.OwnerOf(&_SupplyChain.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SupplyChain *SupplyChainCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SupplyChain.Contract.OwnerOf(&_SupplyChain.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SupplyChain *SupplyChainCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SupplyChain.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SupplyChain *SupplyChainSession) Paused() (bool, error) {
	return _SupplyChain.Contract.Paused(&_SupplyChain.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SupplyChain *SupplyChainCallerSession) Paused() (bool, error) {
	return _SupplyChain.Contract.Paused(&_SupplyChain.CallOpts)
}

// RoleInfo is a free data retrieval call binding the contract method 0x93123e0b.
//
// Solidity: function roleInfo(uint256 role) view returns(string)
func (_SupplyChain *SupplyChainCaller) RoleInfo(opts *bind.CallOpts, role *big.Int) (string, error) {
	var out []interface{}
	err := _SupplyChain.contract.Call(opts, &out, "roleInfo", role)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RoleInfo is a free data retrieval call binding the contract method 0x93123e0b.
//
// Solidity: function roleInfo(uint256 role) view returns(string)
func (_SupplyChain *SupplyChainSession) RoleInfo(role *big.Int) (string, error) {
	return _SupplyChain.Contract.RoleInfo(&_SupplyChain.CallOpts, role)
}

// RoleInfo is a free data retrieval call binding the contract method 0x93123e0b.
//
// Solidity: function roleInfo(uint256 role) view returns(string)
func (_SupplyChain *SupplyChainCallerSession) RoleInfo(role *big.Int) (string, error) {
	return _SupplyChain.Contract.RoleInfo(&_SupplyChain.CallOpts, role)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SupplyChain *SupplyChainCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SupplyChain.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SupplyChain *SupplyChainSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SupplyChain.Contract.SupportsInterface(&_SupplyChain.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SupplyChain *SupplyChainCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SupplyChain.Contract.SupportsInterface(&_SupplyChain.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SupplyChain *SupplyChainCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SupplyChain.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SupplyChain *SupplyChainSession) Symbol() (string, error) {
	return _SupplyChain.Contract.Symbol(&_SupplyChain.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SupplyChain *SupplyChainCallerSession) Symbol() (string, error) {
	return _SupplyChain.Contract.Symbol(&_SupplyChain.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_SupplyChain *SupplyChainCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _SupplyChain.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_SupplyChain *SupplyChainSession) TokenURI(tokenId *big.Int) (string, error) {
	return _SupplyChain.Contract.TokenURI(&_SupplyChain.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_SupplyChain *SupplyChainCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _SupplyChain.Contract.TokenURI(&_SupplyChain.CallOpts, tokenId)
}

// Pause is a paid mutator transaction binding the contract method 0x6985a022.
//
// Solidity: function Pause() returns()
func (_SupplyChain *SupplyChainTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SupplyChain.contract.Transact(opts, "Pause")
}

// Pause is a paid mutator transaction binding the contract method 0x6985a022.
//
// Solidity: function Pause() returns()
func (_SupplyChain *SupplyChainSession) Pause() (*types.Transaction, error) {
	return _SupplyChain.Contract.Pause(&_SupplyChain.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x6985a022.
//
// Solidity: function Pause() returns()
func (_SupplyChain *SupplyChainTransactorSession) Pause() (*types.Transaction, error) {
	return _SupplyChain.Contract.Pause(&_SupplyChain.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x7805862f.
//
// Solidity: function Unpause() returns()
func (_SupplyChain *SupplyChainTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SupplyChain.contract.Transact(opts, "Unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x7805862f.
//
// Solidity: function Unpause() returns()
func (_SupplyChain *SupplyChainSession) Unpause() (*types.Transaction, error) {
	return _SupplyChain.Contract.Unpause(&_SupplyChain.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x7805862f.
//
// Solidity: function Unpause() returns()
func (_SupplyChain *SupplyChainTransactorSession) Unpause() (*types.Transaction, error) {
	return _SupplyChain.Contract.Unpause(&_SupplyChain.TransactOpts)
}

// AddAssetCertificates is a paid mutator transaction binding the contract method 0x4c40065e.
//
// Solidity: function addAssetCertificates(uint256 assetId, uint256[] certificates) returns(bool)
func (_SupplyChain *SupplyChainTransactor) AddAssetCertificates(opts *bind.TransactOpts, assetId *big.Int, certificates []*big.Int) (*types.Transaction, error) {
	return _SupplyChain.contract.Transact(opts, "addAssetCertificates", assetId, certificates)
}

// AddAssetCertificates is a paid mutator transaction binding the contract method 0x4c40065e.
//
// Solidity: function addAssetCertificates(uint256 assetId, uint256[] certificates) returns(bool)
func (_SupplyChain *SupplyChainSession) AddAssetCertificates(assetId *big.Int, certificates []*big.Int) (*types.Transaction, error) {
	return _SupplyChain.Contract.AddAssetCertificates(&_SupplyChain.TransactOpts, assetId, certificates)
}

// AddAssetCertificates is a paid mutator transaction binding the contract method 0x4c40065e.
//
// Solidity: function addAssetCertificates(uint256 assetId, uint256[] certificates) returns(bool)
func (_SupplyChain *SupplyChainTransactorSession) AddAssetCertificates(assetId *big.Int, certificates []*big.Int) (*types.Transaction, error) {
	return _SupplyChain.Contract.AddAssetCertificates(&_SupplyChain.TransactOpts, assetId, certificates)
}

// AddMovementCertificates is a paid mutator transaction binding the contract method 0xaa2d16f1.
//
// Solidity: function addMovementCertificates(uint256 moveId, uint256[] certificates) returns(bool)
func (_SupplyChain *SupplyChainTransactor) AddMovementCertificates(opts *bind.TransactOpts, moveId *big.Int, certificates []*big.Int) (*types.Transaction, error) {
	return _SupplyChain.contract.Transact(opts, "addMovementCertificates", moveId, certificates)
}

// AddMovementCertificates is a paid mutator transaction binding the contract method 0xaa2d16f1.
//
// Solidity: function addMovementCertificates(uint256 moveId, uint256[] certificates) returns(bool)
func (_SupplyChain *SupplyChainSession) AddMovementCertificates(moveId *big.Int, certificates []*big.Int) (*types.Transaction, error) {
	return _SupplyChain.Contract.AddMovementCertificates(&_SupplyChain.TransactOpts, moveId, certificates)
}

// AddMovementCertificates is a paid mutator transaction binding the contract method 0xaa2d16f1.
//
// Solidity: function addMovementCertificates(uint256 moveId, uint256[] certificates) returns(bool)
func (_SupplyChain *SupplyChainTransactorSession) AddMovementCertificates(moveId *big.Int, certificates []*big.Int) (*types.Transaction, error) {
	return _SupplyChain.Contract.AddMovementCertificates(&_SupplyChain.TransactOpts, moveId, certificates)
}

// AddMovements is a paid mutator transaction binding the contract method 0xe4eb771c.
//
// Solidity: function addMovements(uint256 assetId, uint256[] movements) returns(bool)
func (_SupplyChain *SupplyChainTransactor) AddMovements(opts *bind.TransactOpts, assetId *big.Int, movements []*big.Int) (*types.Transaction, error) {
	return _SupplyChain.contract.Transact(opts, "addMovements", assetId, movements)
}

// AddMovements is a paid mutator transaction binding the contract method 0xe4eb771c.
//
// Solidity: function addMovements(uint256 assetId, uint256[] movements) returns(bool)
func (_SupplyChain *SupplyChainSession) AddMovements(assetId *big.Int, movements []*big.Int) (*types.Transaction, error) {
	return _SupplyChain.Contract.AddMovements(&_SupplyChain.TransactOpts, assetId, movements)
}

// AddMovements is a paid mutator transaction binding the contract method 0xe4eb771c.
//
// Solidity: function addMovements(uint256 assetId, uint256[] movements) returns(bool)
func (_SupplyChain *SupplyChainTransactorSession) AddMovements(assetId *big.Int, movements []*big.Int) (*types.Transaction, error) {
	return _SupplyChain.Contract.AddMovements(&_SupplyChain.TransactOpts, assetId, movements)
}

// AddParents is a paid mutator transaction binding the contract method 0xbad6a0ab.
//
// Solidity: function addParents(uint256 assetId, uint256[] parents, uint16[] composition) returns(bool)
func (_SupplyChain *SupplyChainTransactor) AddParents(opts *bind.TransactOpts, assetId *big.Int, parents []*big.Int, composition []uint16) (*types.Transaction, error) {
	return _SupplyChain.contract.Transact(opts, "addParents", assetId, parents, composition)
}

// AddParents is a paid mutator transaction binding the contract method 0xbad6a0ab.
//
// Solidity: function addParents(uint256 assetId, uint256[] parents, uint16[] composition) returns(bool)
func (_SupplyChain *SupplyChainSession) AddParents(assetId *big.Int, parents []*big.Int, composition []uint16) (*types.Transaction, error) {
	return _SupplyChain.Contract.AddParents(&_SupplyChain.TransactOpts, assetId, parents, composition)
}

// AddParents is a paid mutator transaction binding the contract method 0xbad6a0ab.
//
// Solidity: function addParents(uint256 assetId, uint256[] parents, uint16[] composition) returns(bool)
func (_SupplyChain *SupplyChainTransactorSession) AddParents(assetId *big.Int, parents []*big.Int, composition []uint16) (*types.Transaction, error) {
	return _SupplyChain.Contract.AddParents(&_SupplyChain.TransactOpts, assetId, parents, composition)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SupplyChain *SupplyChainTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SupplyChain.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SupplyChain *SupplyChainSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SupplyChain.Contract.Approve(&_SupplyChain.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SupplyChain *SupplyChainTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SupplyChain.Contract.Approve(&_SupplyChain.TransactOpts, to, tokenId)
}

// DefineAgent is a paid mutator transaction binding the contract method 0x5ee3bc2a.
//
// Solidity: function defineAgent(address agent, address account, string metadata_) returns(bool)
func (_SupplyChain *SupplyChainTransactor) DefineAgent(opts *bind.TransactOpts, agent common.Address, account common.Address, metadata_ string) (*types.Transaction, error) {
	return _SupplyChain.contract.Transact(opts, "defineAgent", agent, account, metadata_)
}

// DefineAgent is a paid mutator transaction binding the contract method 0x5ee3bc2a.
//
// Solidity: function defineAgent(address agent, address account, string metadata_) returns(bool)
func (_SupplyChain *SupplyChainSession) DefineAgent(agent common.Address, account common.Address, metadata_ string) (*types.Transaction, error) {
	return _SupplyChain.Contract.DefineAgent(&_SupplyChain.TransactOpts, agent, account, metadata_)
}

// DefineAgent is a paid mutator transaction binding the contract method 0x5ee3bc2a.
//
// Solidity: function defineAgent(address agent, address account, string metadata_) returns(bool)
func (_SupplyChain *SupplyChainTransactorSession) DefineAgent(agent common.Address, account common.Address, metadata_ string) (*types.Transaction, error) {
	return _SupplyChain.Contract.DefineAgent(&_SupplyChain.TransactOpts, agent, account, metadata_)
}

// DefineRole is a paid mutator transaction binding the contract method 0xb77e76c2.
//
// Solidity: function defineRole(uint256 role, string metadata_) returns(bool)
func (_SupplyChain *SupplyChainTransactor) DefineRole(opts *bind.TransactOpts, role *big.Int, metadata_ string) (*types.Transaction, error) {
	return _SupplyChain.contract.Transact(opts, "defineRole", role, metadata_)
}

// DefineRole is a paid mutator transaction binding the contract method 0xb77e76c2.
//
// Solidity: function defineRole(uint256 role, string metadata_) returns(bool)
func (_SupplyChain *SupplyChainSession) DefineRole(role *big.Int, metadata_ string) (*types.Transaction, error) {
	return _SupplyChain.Contract.DefineRole(&_SupplyChain.TransactOpts, role, metadata_)
}

// DefineRole is a paid mutator transaction binding the contract method 0xb77e76c2.
//
// Solidity: function defineRole(uint256 role, string metadata_) returns(bool)
func (_SupplyChain *SupplyChainTransactorSession) DefineRole(role *big.Int, metadata_ string) (*types.Transaction, error) {
	return _SupplyChain.Contract.DefineRole(&_SupplyChain.TransactOpts, role, metadata_)
}

// GrantRole is a paid mutator transaction binding the contract method 0x3c09e2fd.
//
// Solidity: function grantRole(address account, uint256 role) returns(bool)
func (_SupplyChain *SupplyChainTransactor) GrantRole(opts *bind.TransactOpts, account common.Address, role *big.Int) (*types.Transaction, error) {
	return _SupplyChain.contract.Transact(opts, "grantRole", account, role)
}

// GrantRole is a paid mutator transaction binding the contract method 0x3c09e2fd.
//
// Solidity: function grantRole(address account, uint256 role) returns(bool)
func (_SupplyChain *SupplyChainSession) GrantRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _SupplyChain.Contract.GrantRole(&_SupplyChain.TransactOpts, account, role)
}

// GrantRole is a paid mutator transaction binding the contract method 0x3c09e2fd.
//
// Solidity: function grantRole(address account, uint256 role) returns(bool)
func (_SupplyChain *SupplyChainTransactorSession) GrantRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _SupplyChain.Contract.GrantRole(&_SupplyChain.TransactOpts, account, role)
}

// IssueAsset is a paid mutator transaction binding the contract method 0xf17011c6.
//
// Solidity: function issueAsset(uint256 assetId, address owner, uint64 co2e, uint256 certId, string metadata_) returns(bool)
func (_SupplyChain *SupplyChainTransactor) IssueAsset(opts *bind.TransactOpts, assetId *big.Int, owner common.Address, co2e uint64, certId *big.Int, metadata_ string) (*types.Transaction, error) {
	return _SupplyChain.contract.Transact(opts, "issueAsset", assetId, owner, co2e, certId, metadata_)
}

// IssueAsset is a paid mutator transaction binding the contract method 0xf17011c6.
//
// Solidity: function issueAsset(uint256 assetId, address owner, uint64 co2e, uint256 certId, string metadata_) returns(bool)
func (_SupplyChain *SupplyChainSession) IssueAsset(assetId *big.Int, owner common.Address, co2e uint64, certId *big.Int, metadata_ string) (*types.Transaction, error) {
	return _SupplyChain.Contract.IssueAsset(&_SupplyChain.TransactOpts, assetId, owner, co2e, certId, metadata_)
}

// IssueAsset is a paid mutator transaction binding the contract method 0xf17011c6.
//
// Solidity: function issueAsset(uint256 assetId, address owner, uint64 co2e, uint256 certId, string metadata_) returns(bool)
func (_SupplyChain *SupplyChainTransactorSession) IssueAsset(assetId *big.Int, owner common.Address, co2e uint64, certId *big.Int, metadata_ string) (*types.Transaction, error) {
	return _SupplyChain.Contract.IssueAsset(&_SupplyChain.TransactOpts, assetId, owner, co2e, certId, metadata_)
}

// IssueCertificate is a paid mutator transaction binding the contract method 0x0c816de4.
//
// Solidity: function issueCertificate(uint256 certId, string metadata_) returns(bool)
func (_SupplyChain *SupplyChainTransactor) IssueCertificate(opts *bind.TransactOpts, certId *big.Int, metadata_ string) (*types.Transaction, error) {
	return _SupplyChain.contract.Transact(opts, "issueCertificate", certId, metadata_)
}

// IssueCertificate is a paid mutator transaction binding the contract method 0x0c816de4.
//
// Solidity: function issueCertificate(uint256 certId, string metadata_) returns(bool)
func (_SupplyChain *SupplyChainSession) IssueCertificate(certId *big.Int, metadata_ string) (*types.Transaction, error) {
	return _SupplyChain.Contract.IssueCertificate(&_SupplyChain.TransactOpts, certId, metadata_)
}

// IssueCertificate is a paid mutator transaction binding the contract method 0x0c816de4.
//
// Solidity: function issueCertificate(uint256 certId, string metadata_) returns(bool)
func (_SupplyChain *SupplyChainTransactorSession) IssueCertificate(certId *big.Int, metadata_ string) (*types.Transaction, error) {
	return _SupplyChain.Contract.IssueCertificate(&_SupplyChain.TransactOpts, certId, metadata_)
}

// IssueMovement is a paid mutator transaction binding the contract method 0xe842eafa.
//
// Solidity: function issueMovement(uint256 moveId, uint64 co2e, uint256 certId, string metadata_) returns(bool)
func (_SupplyChain *SupplyChainTransactor) IssueMovement(opts *bind.TransactOpts, moveId *big.Int, co2e uint64, certId *big.Int, metadata_ string) (*types.Transaction, error) {
	return _SupplyChain.contract.Transact(opts, "issueMovement", moveId, co2e, certId, metadata_)
}

// IssueMovement is a paid mutator transaction binding the contract method 0xe842eafa.
//
// Solidity: function issueMovement(uint256 moveId, uint64 co2e, uint256 certId, string metadata_) returns(bool)
func (_SupplyChain *SupplyChainSession) IssueMovement(moveId *big.Int, co2e uint64, certId *big.Int, metadata_ string) (*types.Transaction, error) {
	return _SupplyChain.Contract.IssueMovement(&_SupplyChain.TransactOpts, moveId, co2e, certId, metadata_)
}

// IssueMovement is a paid mutator transaction binding the contract method 0xe842eafa.
//
// Solidity: function issueMovement(uint256 moveId, uint64 co2e, uint256 certId, string metadata_) returns(bool)
func (_SupplyChain *SupplyChainTransactorSession) IssueMovement(moveId *big.Int, co2e uint64, certId *big.Int, metadata_ string) (*types.Transaction, error) {
	return _SupplyChain.Contract.IssueMovement(&_SupplyChain.TransactOpts, moveId, co2e, certId, metadata_)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0x97a6278e.
//
// Solidity: function removeAgent(address agent) returns(bool)
func (_SupplyChain *SupplyChainTransactor) RemoveAgent(opts *bind.TransactOpts, agent common.Address) (*types.Transaction, error) {
	return _SupplyChain.contract.Transact(opts, "removeAgent", agent)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0x97a6278e.
//
// Solidity: function removeAgent(address agent) returns(bool)
func (_SupplyChain *SupplyChainSession) RemoveAgent(agent common.Address) (*types.Transaction, error) {
	return _SupplyChain.Contract.RemoveAgent(&_SupplyChain.TransactOpts, agent)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0x97a6278e.
//
// Solidity: function removeAgent(address agent) returns(bool)
func (_SupplyChain *SupplyChainTransactorSession) RemoveAgent(agent common.Address) (*types.Transaction, error) {
	return _SupplyChain.Contract.RemoveAgent(&_SupplyChain.TransactOpts, agent)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x0912ed77.
//
// Solidity: function revokeRole(address account, uint256 role) returns(bool)
func (_SupplyChain *SupplyChainTransactor) RevokeRole(opts *bind.TransactOpts, account common.Address, role *big.Int) (*types.Transaction, error) {
	return _SupplyChain.contract.Transact(opts, "revokeRole", account, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x0912ed77.
//
// Solidity: function revokeRole(address account, uint256 role) returns(bool)
func (_SupplyChain *SupplyChainSession) RevokeRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _SupplyChain.Contract.RevokeRole(&_SupplyChain.TransactOpts, account, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x0912ed77.
//
// Solidity: function revokeRole(address account, uint256 role) returns(bool)
func (_SupplyChain *SupplyChainTransactorSession) RevokeRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _SupplyChain.Contract.RevokeRole(&_SupplyChain.TransactOpts, account, role)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SupplyChain *SupplyChainTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SupplyChain.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SupplyChain *SupplyChainSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SupplyChain.Contract.SafeTransferFrom(&_SupplyChain.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SupplyChain *SupplyChainTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SupplyChain.Contract.SafeTransferFrom(&_SupplyChain.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_SupplyChain *SupplyChainTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _SupplyChain.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_SupplyChain *SupplyChainSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _SupplyChain.Contract.SafeTransferFrom0(&_SupplyChain.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_SupplyChain *SupplyChainTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _SupplyChain.Contract.SafeTransferFrom0(&_SupplyChain.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SupplyChain *SupplyChainTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _SupplyChain.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SupplyChain *SupplyChainSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SupplyChain.Contract.SetApprovalForAll(&_SupplyChain.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SupplyChain *SupplyChainTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SupplyChain.Contract.SetApprovalForAll(&_SupplyChain.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SupplyChain *SupplyChainTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SupplyChain.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SupplyChain *SupplyChainSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SupplyChain.Contract.TransferFrom(&_SupplyChain.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SupplyChain *SupplyChainTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SupplyChain.Contract.TransferFrom(&_SupplyChain.TransactOpts, from, to, tokenId)
}

// SupplyChainAgentDefinedIterator is returned from FilterAgentDefined and is used to iterate over the raw logs and unpacked data for AgentDefined events raised by the SupplyChain contract.
type SupplyChainAgentDefinedIterator struct {
	Event *SupplyChainAgentDefined // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SupplyChainAgentDefinedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyChainAgentDefined)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SupplyChainAgentDefined)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SupplyChainAgentDefinedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyChainAgentDefinedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyChainAgentDefined represents a AgentDefined event raised by the SupplyChain contract.
type SupplyChainAgentDefined struct {
	Agent   common.Address
	Account common.Address
	Issuer  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentDefined is a free log retrieval operation binding the contract event 0xb61ddfe33d9ed54d8bdb62fe6fe16b84a72a4bae657baea7d2ae636fd1714b89.
//
// Solidity: event AgentDefined(address agent, address account, address issuer)
func (_SupplyChain *SupplyChainFilterer) FilterAgentDefined(opts *bind.FilterOpts) (*SupplyChainAgentDefinedIterator, error) {

	logs, sub, err := _SupplyChain.contract.FilterLogs(opts, "AgentDefined")
	if err != nil {
		return nil, err
	}
	return &SupplyChainAgentDefinedIterator{contract: _SupplyChain.contract, event: "AgentDefined", logs: logs, sub: sub}, nil
}

// WatchAgentDefined is a free log subscription operation binding the contract event 0xb61ddfe33d9ed54d8bdb62fe6fe16b84a72a4bae657baea7d2ae636fd1714b89.
//
// Solidity: event AgentDefined(address agent, address account, address issuer)
func (_SupplyChain *SupplyChainFilterer) WatchAgentDefined(opts *bind.WatchOpts, sink chan<- *SupplyChainAgentDefined) (event.Subscription, error) {

	logs, sub, err := _SupplyChain.contract.WatchLogs(opts, "AgentDefined")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyChainAgentDefined)
				if err := _SupplyChain.contract.UnpackLog(event, "AgentDefined", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAgentDefined is a log parse operation binding the contract event 0xb61ddfe33d9ed54d8bdb62fe6fe16b84a72a4bae657baea7d2ae636fd1714b89.
//
// Solidity: event AgentDefined(address agent, address account, address issuer)
func (_SupplyChain *SupplyChainFilterer) ParseAgentDefined(log types.Log) (*SupplyChainAgentDefined, error) {
	event := new(SupplyChainAgentDefined)
	if err := _SupplyChain.contract.UnpackLog(event, "AgentDefined", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SupplyChainAgentRemovedIterator is returned from FilterAgentRemoved and is used to iterate over the raw logs and unpacked data for AgentRemoved events raised by the SupplyChain contract.
type SupplyChainAgentRemovedIterator struct {
	Event *SupplyChainAgentRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SupplyChainAgentRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyChainAgentRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SupplyChainAgentRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SupplyChainAgentRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyChainAgentRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyChainAgentRemoved represents a AgentRemoved event raised by the SupplyChain contract.
type SupplyChainAgentRemoved struct {
	Agent  common.Address
	Issuer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAgentRemoved is a free log retrieval operation binding the contract event 0x4ce8b3fe8e0bf47a6fefbecdfd7c799755cede6061655521b10fc2e4b1fcc6b3.
//
// Solidity: event AgentRemoved(address agent, address issuer)
func (_SupplyChain *SupplyChainFilterer) FilterAgentRemoved(opts *bind.FilterOpts) (*SupplyChainAgentRemovedIterator, error) {

	logs, sub, err := _SupplyChain.contract.FilterLogs(opts, "AgentRemoved")
	if err != nil {
		return nil, err
	}
	return &SupplyChainAgentRemovedIterator{contract: _SupplyChain.contract, event: "AgentRemoved", logs: logs, sub: sub}, nil
}

// WatchAgentRemoved is a free log subscription operation binding the contract event 0x4ce8b3fe8e0bf47a6fefbecdfd7c799755cede6061655521b10fc2e4b1fcc6b3.
//
// Solidity: event AgentRemoved(address agent, address issuer)
func (_SupplyChain *SupplyChainFilterer) WatchAgentRemoved(opts *bind.WatchOpts, sink chan<- *SupplyChainAgentRemoved) (event.Subscription, error) {

	logs, sub, err := _SupplyChain.contract.WatchLogs(opts, "AgentRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyChainAgentRemoved)
				if err := _SupplyChain.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAgentRemoved is a log parse operation binding the contract event 0x4ce8b3fe8e0bf47a6fefbecdfd7c799755cede6061655521b10fc2e4b1fcc6b3.
//
// Solidity: event AgentRemoved(address agent, address issuer)
func (_SupplyChain *SupplyChainFilterer) ParseAgentRemoved(log types.Log) (*SupplyChainAgentRemoved, error) {
	event := new(SupplyChainAgentRemoved)
	if err := _SupplyChain.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SupplyChainApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SupplyChain contract.
type SupplyChainApprovalIterator struct {
	Event *SupplyChainApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SupplyChainApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyChainApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SupplyChainApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SupplyChainApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyChainApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyChainApproval represents a Approval event raised by the SupplyChain contract.
type SupplyChainApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SupplyChain *SupplyChainFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*SupplyChainApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SupplyChain.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SupplyChainApprovalIterator{contract: _SupplyChain.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SupplyChain *SupplyChainFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SupplyChainApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SupplyChain.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyChainApproval)
				if err := _SupplyChain.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SupplyChain *SupplyChainFilterer) ParseApproval(log types.Log) (*SupplyChainApproval, error) {
	event := new(SupplyChainApproval)
	if err := _SupplyChain.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SupplyChainApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the SupplyChain contract.
type SupplyChainApprovalForAllIterator struct {
	Event *SupplyChainApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SupplyChainApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyChainApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SupplyChainApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SupplyChainApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyChainApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyChainApprovalForAll represents a ApprovalForAll event raised by the SupplyChain contract.
type SupplyChainApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SupplyChain *SupplyChainFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*SupplyChainApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SupplyChain.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &SupplyChainApprovalForAllIterator{contract: _SupplyChain.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SupplyChain *SupplyChainFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *SupplyChainApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SupplyChain.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyChainApprovalForAll)
				if err := _SupplyChain.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SupplyChain *SupplyChainFilterer) ParseApprovalForAll(log types.Log) (*SupplyChainApprovalForAll, error) {
	event := new(SupplyChainApprovalForAll)
	if err := _SupplyChain.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SupplyChainAssetCertificateAssignedIterator is returned from FilterAssetCertificateAssigned and is used to iterate over the raw logs and unpacked data for AssetCertificateAssigned events raised by the SupplyChain contract.
type SupplyChainAssetCertificateAssignedIterator struct {
	Event *SupplyChainAssetCertificateAssigned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SupplyChainAssetCertificateAssignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyChainAssetCertificateAssigned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SupplyChainAssetCertificateAssigned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SupplyChainAssetCertificateAssignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyChainAssetCertificateAssignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyChainAssetCertificateAssigned represents a AssetCertificateAssigned event raised by the SupplyChain contract.
type SupplyChainAssetCertificateAssigned struct {
	CertId  *big.Int
	AssetId *big.Int
	Signer  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAssetCertificateAssigned is a free log retrieval operation binding the contract event 0x4fc952580cfc5ff048a6b9111d8facdf029038c2a88b0ea0c42da67a1a5e00d1.
//
// Solidity: event AssetCertificateAssigned(uint256 certId, uint256 assetId, address signer)
func (_SupplyChain *SupplyChainFilterer) FilterAssetCertificateAssigned(opts *bind.FilterOpts) (*SupplyChainAssetCertificateAssignedIterator, error) {

	logs, sub, err := _SupplyChain.contract.FilterLogs(opts, "AssetCertificateAssigned")
	if err != nil {
		return nil, err
	}
	return &SupplyChainAssetCertificateAssignedIterator{contract: _SupplyChain.contract, event: "AssetCertificateAssigned", logs: logs, sub: sub}, nil
}

// WatchAssetCertificateAssigned is a free log subscription operation binding the contract event 0x4fc952580cfc5ff048a6b9111d8facdf029038c2a88b0ea0c42da67a1a5e00d1.
//
// Solidity: event AssetCertificateAssigned(uint256 certId, uint256 assetId, address signer)
func (_SupplyChain *SupplyChainFilterer) WatchAssetCertificateAssigned(opts *bind.WatchOpts, sink chan<- *SupplyChainAssetCertificateAssigned) (event.Subscription, error) {

	logs, sub, err := _SupplyChain.contract.WatchLogs(opts, "AssetCertificateAssigned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyChainAssetCertificateAssigned)
				if err := _SupplyChain.contract.UnpackLog(event, "AssetCertificateAssigned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAssetCertificateAssigned is a log parse operation binding the contract event 0x4fc952580cfc5ff048a6b9111d8facdf029038c2a88b0ea0c42da67a1a5e00d1.
//
// Solidity: event AssetCertificateAssigned(uint256 certId, uint256 assetId, address signer)
func (_SupplyChain *SupplyChainFilterer) ParseAssetCertificateAssigned(log types.Log) (*SupplyChainAssetCertificateAssigned, error) {
	event := new(SupplyChainAssetCertificateAssigned)
	if err := _SupplyChain.contract.UnpackLog(event, "AssetCertificateAssigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SupplyChainAssetIssuedIterator is returned from FilterAssetIssued and is used to iterate over the raw logs and unpacked data for AssetIssued events raised by the SupplyChain contract.
type SupplyChainAssetIssuedIterator struct {
	Event *SupplyChainAssetIssued // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SupplyChainAssetIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyChainAssetIssued)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SupplyChainAssetIssued)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SupplyChainAssetIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyChainAssetIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyChainAssetIssued represents a AssetIssued event raised by the SupplyChain contract.
type SupplyChainAssetIssued struct {
	AssetId *big.Int
	Signer  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAssetIssued is a free log retrieval operation binding the contract event 0xf37f680ac83ec9f2a328ebff310dfe6d49043380dcb563e89ef4ca09eaa38d51.
//
// Solidity: event AssetIssued(uint256 assetId, address indexed signer)
func (_SupplyChain *SupplyChainFilterer) FilterAssetIssued(opts *bind.FilterOpts, signer []common.Address) (*SupplyChainAssetIssuedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _SupplyChain.contract.FilterLogs(opts, "AssetIssued", signerRule)
	if err != nil {
		return nil, err
	}
	return &SupplyChainAssetIssuedIterator{contract: _SupplyChain.contract, event: "AssetIssued", logs: logs, sub: sub}, nil
}

// WatchAssetIssued is a free log subscription operation binding the contract event 0xf37f680ac83ec9f2a328ebff310dfe6d49043380dcb563e89ef4ca09eaa38d51.
//
// Solidity: event AssetIssued(uint256 assetId, address indexed signer)
func (_SupplyChain *SupplyChainFilterer) WatchAssetIssued(opts *bind.WatchOpts, sink chan<- *SupplyChainAssetIssued, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _SupplyChain.contract.WatchLogs(opts, "AssetIssued", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyChainAssetIssued)
				if err := _SupplyChain.contract.UnpackLog(event, "AssetIssued", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAssetIssued is a log parse operation binding the contract event 0xf37f680ac83ec9f2a328ebff310dfe6d49043380dcb563e89ef4ca09eaa38d51.
//
// Solidity: event AssetIssued(uint256 assetId, address indexed signer)
func (_SupplyChain *SupplyChainFilterer) ParseAssetIssued(log types.Log) (*SupplyChainAssetIssued, error) {
	event := new(SupplyChainAssetIssued)
	if err := _SupplyChain.contract.UnpackLog(event, "AssetIssued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SupplyChainCertIssuedIterator is returned from FilterCertIssued and is used to iterate over the raw logs and unpacked data for CertIssued events raised by the SupplyChain contract.
type SupplyChainCertIssuedIterator struct {
	Event *SupplyChainCertIssued // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SupplyChainCertIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyChainCertIssued)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SupplyChainCertIssued)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SupplyChainCertIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyChainCertIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyChainCertIssued represents a CertIssued event raised by the SupplyChain contract.
type SupplyChainCertIssued struct {
	CertId *big.Int
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCertIssued is a free log retrieval operation binding the contract event 0x0ac0bb5418edafe119eaae67d0ee0bc6f39a3efafc3305a56b617c171d0d835f.
//
// Solidity: event CertIssued(uint256 certId, address indexed signer)
func (_SupplyChain *SupplyChainFilterer) FilterCertIssued(opts *bind.FilterOpts, signer []common.Address) (*SupplyChainCertIssuedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _SupplyChain.contract.FilterLogs(opts, "CertIssued", signerRule)
	if err != nil {
		return nil, err
	}
	return &SupplyChainCertIssuedIterator{contract: _SupplyChain.contract, event: "CertIssued", logs: logs, sub: sub}, nil
}

// WatchCertIssued is a free log subscription operation binding the contract event 0x0ac0bb5418edafe119eaae67d0ee0bc6f39a3efafc3305a56b617c171d0d835f.
//
// Solidity: event CertIssued(uint256 certId, address indexed signer)
func (_SupplyChain *SupplyChainFilterer) WatchCertIssued(opts *bind.WatchOpts, sink chan<- *SupplyChainCertIssued, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _SupplyChain.contract.WatchLogs(opts, "CertIssued", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyChainCertIssued)
				if err := _SupplyChain.contract.UnpackLog(event, "CertIssued", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCertIssued is a log parse operation binding the contract event 0x0ac0bb5418edafe119eaae67d0ee0bc6f39a3efafc3305a56b617c171d0d835f.
//
// Solidity: event CertIssued(uint256 certId, address indexed signer)
func (_SupplyChain *SupplyChainFilterer) ParseCertIssued(log types.Log) (*SupplyChainCertIssued, error) {
	event := new(SupplyChainCertIssued)
	if err := _SupplyChain.contract.UnpackLog(event, "CertIssued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SupplyChainMovementAssignedIterator is returned from FilterMovementAssigned and is used to iterate over the raw logs and unpacked data for MovementAssigned events raised by the SupplyChain contract.
type SupplyChainMovementAssignedIterator struct {
	Event *SupplyChainMovementAssigned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SupplyChainMovementAssignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyChainMovementAssigned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SupplyChainMovementAssigned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SupplyChainMovementAssignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyChainMovementAssignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyChainMovementAssigned represents a MovementAssigned event raised by the SupplyChain contract.
type SupplyChainMovementAssigned struct {
	MoveId  *big.Int
	AssetId *big.Int
	Signer  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMovementAssigned is a free log retrieval operation binding the contract event 0x1fb5e37b2fee9c771f5a7ac82a15698c9f978e0e6cef7ac48ceb8b9792caaf16.
//
// Solidity: event MovementAssigned(uint256 moveId, uint256 assetId, address signer)
func (_SupplyChain *SupplyChainFilterer) FilterMovementAssigned(opts *bind.FilterOpts) (*SupplyChainMovementAssignedIterator, error) {

	logs, sub, err := _SupplyChain.contract.FilterLogs(opts, "MovementAssigned")
	if err != nil {
		return nil, err
	}
	return &SupplyChainMovementAssignedIterator{contract: _SupplyChain.contract, event: "MovementAssigned", logs: logs, sub: sub}, nil
}

// WatchMovementAssigned is a free log subscription operation binding the contract event 0x1fb5e37b2fee9c771f5a7ac82a15698c9f978e0e6cef7ac48ceb8b9792caaf16.
//
// Solidity: event MovementAssigned(uint256 moveId, uint256 assetId, address signer)
func (_SupplyChain *SupplyChainFilterer) WatchMovementAssigned(opts *bind.WatchOpts, sink chan<- *SupplyChainMovementAssigned) (event.Subscription, error) {

	logs, sub, err := _SupplyChain.contract.WatchLogs(opts, "MovementAssigned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyChainMovementAssigned)
				if err := _SupplyChain.contract.UnpackLog(event, "MovementAssigned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMovementAssigned is a log parse operation binding the contract event 0x1fb5e37b2fee9c771f5a7ac82a15698c9f978e0e6cef7ac48ceb8b9792caaf16.
//
// Solidity: event MovementAssigned(uint256 moveId, uint256 assetId, address signer)
func (_SupplyChain *SupplyChainFilterer) ParseMovementAssigned(log types.Log) (*SupplyChainMovementAssigned, error) {
	event := new(SupplyChainMovementAssigned)
	if err := _SupplyChain.contract.UnpackLog(event, "MovementAssigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SupplyChainMovementCertificateAssignedIterator is returned from FilterMovementCertificateAssigned and is used to iterate over the raw logs and unpacked data for MovementCertificateAssigned events raised by the SupplyChain contract.
type SupplyChainMovementCertificateAssignedIterator struct {
	Event *SupplyChainMovementCertificateAssigned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SupplyChainMovementCertificateAssignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyChainMovementCertificateAssigned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SupplyChainMovementCertificateAssigned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SupplyChainMovementCertificateAssignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyChainMovementCertificateAssignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyChainMovementCertificateAssigned represents a MovementCertificateAssigned event raised by the SupplyChain contract.
type SupplyChainMovementCertificateAssigned struct {
	CertId *big.Int
	MoveId *big.Int
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMovementCertificateAssigned is a free log retrieval operation binding the contract event 0xb8730f75259bf8d0baccdacc5bc1580e45d23db9e5a88f49ba8bc829c14f64f8.
//
// Solidity: event MovementCertificateAssigned(uint256 certId, uint256 moveId, address signer)
func (_SupplyChain *SupplyChainFilterer) FilterMovementCertificateAssigned(opts *bind.FilterOpts) (*SupplyChainMovementCertificateAssignedIterator, error) {

	logs, sub, err := _SupplyChain.contract.FilterLogs(opts, "MovementCertificateAssigned")
	if err != nil {
		return nil, err
	}
	return &SupplyChainMovementCertificateAssignedIterator{contract: _SupplyChain.contract, event: "MovementCertificateAssigned", logs: logs, sub: sub}, nil
}

// WatchMovementCertificateAssigned is a free log subscription operation binding the contract event 0xb8730f75259bf8d0baccdacc5bc1580e45d23db9e5a88f49ba8bc829c14f64f8.
//
// Solidity: event MovementCertificateAssigned(uint256 certId, uint256 moveId, address signer)
func (_SupplyChain *SupplyChainFilterer) WatchMovementCertificateAssigned(opts *bind.WatchOpts, sink chan<- *SupplyChainMovementCertificateAssigned) (event.Subscription, error) {

	logs, sub, err := _SupplyChain.contract.WatchLogs(opts, "MovementCertificateAssigned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyChainMovementCertificateAssigned)
				if err := _SupplyChain.contract.UnpackLog(event, "MovementCertificateAssigned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMovementCertificateAssigned is a log parse operation binding the contract event 0xb8730f75259bf8d0baccdacc5bc1580e45d23db9e5a88f49ba8bc829c14f64f8.
//
// Solidity: event MovementCertificateAssigned(uint256 certId, uint256 moveId, address signer)
func (_SupplyChain *SupplyChainFilterer) ParseMovementCertificateAssigned(log types.Log) (*SupplyChainMovementCertificateAssigned, error) {
	event := new(SupplyChainMovementCertificateAssigned)
	if err := _SupplyChain.contract.UnpackLog(event, "MovementCertificateAssigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SupplyChainMovementIssuedIterator is returned from FilterMovementIssued and is used to iterate over the raw logs and unpacked data for MovementIssued events raised by the SupplyChain contract.
type SupplyChainMovementIssuedIterator struct {
	Event *SupplyChainMovementIssued // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SupplyChainMovementIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyChainMovementIssued)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SupplyChainMovementIssued)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SupplyChainMovementIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyChainMovementIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyChainMovementIssued represents a MovementIssued event raised by the SupplyChain contract.
type SupplyChainMovementIssued struct {
	MoveId *big.Int
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMovementIssued is a free log retrieval operation binding the contract event 0x741b29ae377b3381e554e394b7d2e2fefb4db0c40f73b9e0ac4cd843ef56171d.
//
// Solidity: event MovementIssued(uint256 moveId, address indexed signer)
func (_SupplyChain *SupplyChainFilterer) FilterMovementIssued(opts *bind.FilterOpts, signer []common.Address) (*SupplyChainMovementIssuedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _SupplyChain.contract.FilterLogs(opts, "MovementIssued", signerRule)
	if err != nil {
		return nil, err
	}
	return &SupplyChainMovementIssuedIterator{contract: _SupplyChain.contract, event: "MovementIssued", logs: logs, sub: sub}, nil
}

// WatchMovementIssued is a free log subscription operation binding the contract event 0x741b29ae377b3381e554e394b7d2e2fefb4db0c40f73b9e0ac4cd843ef56171d.
//
// Solidity: event MovementIssued(uint256 moveId, address indexed signer)
func (_SupplyChain *SupplyChainFilterer) WatchMovementIssued(opts *bind.WatchOpts, sink chan<- *SupplyChainMovementIssued, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _SupplyChain.contract.WatchLogs(opts, "MovementIssued", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyChainMovementIssued)
				if err := _SupplyChain.contract.UnpackLog(event, "MovementIssued", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMovementIssued is a log parse operation binding the contract event 0x741b29ae377b3381e554e394b7d2e2fefb4db0c40f73b9e0ac4cd843ef56171d.
//
// Solidity: event MovementIssued(uint256 moveId, address indexed signer)
func (_SupplyChain *SupplyChainFilterer) ParseMovementIssued(log types.Log) (*SupplyChainMovementIssued, error) {
	event := new(SupplyChainMovementIssued)
	if err := _SupplyChain.contract.UnpackLog(event, "MovementIssued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SupplyChainParentAssignedIterator is returned from FilterParentAssigned and is used to iterate over the raw logs and unpacked data for ParentAssigned events raised by the SupplyChain contract.
type SupplyChainParentAssignedIterator struct {
	Event *SupplyChainParentAssigned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SupplyChainParentAssignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyChainParentAssigned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SupplyChainParentAssigned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SupplyChainParentAssignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyChainParentAssignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyChainParentAssigned represents a ParentAssigned event raised by the SupplyChain contract.
type SupplyChainParentAssigned struct {
	ParentId *big.Int
	AssetId  *big.Int
	Signer   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterParentAssigned is a free log retrieval operation binding the contract event 0x4154f07b27a742714e27b3894179c9cb5aba0e2a7fa26cb0f7513abecf1b051b.
//
// Solidity: event ParentAssigned(uint256 parentId, uint256 assetId, address signer)
func (_SupplyChain *SupplyChainFilterer) FilterParentAssigned(opts *bind.FilterOpts) (*SupplyChainParentAssignedIterator, error) {

	logs, sub, err := _SupplyChain.contract.FilterLogs(opts, "ParentAssigned")
	if err != nil {
		return nil, err
	}
	return &SupplyChainParentAssignedIterator{contract: _SupplyChain.contract, event: "ParentAssigned", logs: logs, sub: sub}, nil
}

// WatchParentAssigned is a free log subscription operation binding the contract event 0x4154f07b27a742714e27b3894179c9cb5aba0e2a7fa26cb0f7513abecf1b051b.
//
// Solidity: event ParentAssigned(uint256 parentId, uint256 assetId, address signer)
func (_SupplyChain *SupplyChainFilterer) WatchParentAssigned(opts *bind.WatchOpts, sink chan<- *SupplyChainParentAssigned) (event.Subscription, error) {

	logs, sub, err := _SupplyChain.contract.WatchLogs(opts, "ParentAssigned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyChainParentAssigned)
				if err := _SupplyChain.contract.UnpackLog(event, "ParentAssigned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseParentAssigned is a log parse operation binding the contract event 0x4154f07b27a742714e27b3894179c9cb5aba0e2a7fa26cb0f7513abecf1b051b.
//
// Solidity: event ParentAssigned(uint256 parentId, uint256 assetId, address signer)
func (_SupplyChain *SupplyChainFilterer) ParseParentAssigned(log types.Log) (*SupplyChainParentAssigned, error) {
	event := new(SupplyChainParentAssigned)
	if err := _SupplyChain.contract.UnpackLog(event, "ParentAssigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SupplyChainPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the SupplyChain contract.
type SupplyChainPausedIterator struct {
	Event *SupplyChainPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SupplyChainPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyChainPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SupplyChainPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SupplyChainPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyChainPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyChainPaused represents a Paused event raised by the SupplyChain contract.
type SupplyChainPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SupplyChain *SupplyChainFilterer) FilterPaused(opts *bind.FilterOpts) (*SupplyChainPausedIterator, error) {

	logs, sub, err := _SupplyChain.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SupplyChainPausedIterator{contract: _SupplyChain.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SupplyChain *SupplyChainFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SupplyChainPaused) (event.Subscription, error) {

	logs, sub, err := _SupplyChain.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyChainPaused)
				if err := _SupplyChain.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SupplyChain *SupplyChainFilterer) ParsePaused(log types.Log) (*SupplyChainPaused, error) {
	event := new(SupplyChainPaused)
	if err := _SupplyChain.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SupplyChainRoleDefinedIterator is returned from FilterRoleDefined and is used to iterate over the raw logs and unpacked data for RoleDefined events raised by the SupplyChain contract.
type SupplyChainRoleDefinedIterator struct {
	Event *SupplyChainRoleDefined // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SupplyChainRoleDefinedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyChainRoleDefined)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SupplyChainRoleDefined)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SupplyChainRoleDefinedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyChainRoleDefinedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyChainRoleDefined represents a RoleDefined event raised by the SupplyChain contract.
type SupplyChainRoleDefined struct {
	Role   *big.Int
	Issuer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRoleDefined is a free log retrieval operation binding the contract event 0x2e82a3308270f66933807098cf62441247b1de86244636b03b0565b2c3dbc62b.
//
// Solidity: event RoleDefined(uint256 role, address issuer)
func (_SupplyChain *SupplyChainFilterer) FilterRoleDefined(opts *bind.FilterOpts) (*SupplyChainRoleDefinedIterator, error) {

	logs, sub, err := _SupplyChain.contract.FilterLogs(opts, "RoleDefined")
	if err != nil {
		return nil, err
	}
	return &SupplyChainRoleDefinedIterator{contract: _SupplyChain.contract, event: "RoleDefined", logs: logs, sub: sub}, nil
}

// WatchRoleDefined is a free log subscription operation binding the contract event 0x2e82a3308270f66933807098cf62441247b1de86244636b03b0565b2c3dbc62b.
//
// Solidity: event RoleDefined(uint256 role, address issuer)
func (_SupplyChain *SupplyChainFilterer) WatchRoleDefined(opts *bind.WatchOpts, sink chan<- *SupplyChainRoleDefined) (event.Subscription, error) {

	logs, sub, err := _SupplyChain.contract.WatchLogs(opts, "RoleDefined")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyChainRoleDefined)
				if err := _SupplyChain.contract.UnpackLog(event, "RoleDefined", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleDefined is a log parse operation binding the contract event 0x2e82a3308270f66933807098cf62441247b1de86244636b03b0565b2c3dbc62b.
//
// Solidity: event RoleDefined(uint256 role, address issuer)
func (_SupplyChain *SupplyChainFilterer) ParseRoleDefined(log types.Log) (*SupplyChainRoleDefined, error) {
	event := new(SupplyChainRoleDefined)
	if err := _SupplyChain.contract.UnpackLog(event, "RoleDefined", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SupplyChainRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the SupplyChain contract.
type SupplyChainRoleGrantedIterator struct {
	Event *SupplyChainRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SupplyChainRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyChainRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SupplyChainRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SupplyChainRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyChainRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyChainRoleGranted represents a RoleGranted event raised by the SupplyChain contract.
type SupplyChainRoleGranted struct {
	Role    *big.Int
	Account common.Address
	Issuer  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x1ec1667fba5e43c5c76fff54e76d7a4a20a4fecf7b49724ad8d52a3f726e9dbd.
//
// Solidity: event RoleGranted(uint256 role, address account, address issuer)
func (_SupplyChain *SupplyChainFilterer) FilterRoleGranted(opts *bind.FilterOpts) (*SupplyChainRoleGrantedIterator, error) {

	logs, sub, err := _SupplyChain.contract.FilterLogs(opts, "RoleGranted")
	if err != nil {
		return nil, err
	}
	return &SupplyChainRoleGrantedIterator{contract: _SupplyChain.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x1ec1667fba5e43c5c76fff54e76d7a4a20a4fecf7b49724ad8d52a3f726e9dbd.
//
// Solidity: event RoleGranted(uint256 role, address account, address issuer)
func (_SupplyChain *SupplyChainFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SupplyChainRoleGranted) (event.Subscription, error) {

	logs, sub, err := _SupplyChain.contract.WatchLogs(opts, "RoleGranted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyChainRoleGranted)
				if err := _SupplyChain.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x1ec1667fba5e43c5c76fff54e76d7a4a20a4fecf7b49724ad8d52a3f726e9dbd.
//
// Solidity: event RoleGranted(uint256 role, address account, address issuer)
func (_SupplyChain *SupplyChainFilterer) ParseRoleGranted(log types.Log) (*SupplyChainRoleGranted, error) {
	event := new(SupplyChainRoleGranted)
	if err := _SupplyChain.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SupplyChainRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the SupplyChain contract.
type SupplyChainRoleRevokedIterator struct {
	Event *SupplyChainRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SupplyChainRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyChainRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SupplyChainRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SupplyChainRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyChainRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyChainRoleRevoked represents a RoleRevoked event raised by the SupplyChain contract.
type SupplyChainRoleRevoked struct {
	Role    *big.Int
	Account common.Address
	Issuer  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xe0df21b65c73c27081b8f042a012b124085b41d78d27b7e3c4780f5650f5ebb8.
//
// Solidity: event RoleRevoked(uint256 role, address account, address issuer)
func (_SupplyChain *SupplyChainFilterer) FilterRoleRevoked(opts *bind.FilterOpts) (*SupplyChainRoleRevokedIterator, error) {

	logs, sub, err := _SupplyChain.contract.FilterLogs(opts, "RoleRevoked")
	if err != nil {
		return nil, err
	}
	return &SupplyChainRoleRevokedIterator{contract: _SupplyChain.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xe0df21b65c73c27081b8f042a012b124085b41d78d27b7e3c4780f5650f5ebb8.
//
// Solidity: event RoleRevoked(uint256 role, address account, address issuer)
func (_SupplyChain *SupplyChainFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SupplyChainRoleRevoked) (event.Subscription, error) {

	logs, sub, err := _SupplyChain.contract.WatchLogs(opts, "RoleRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyChainRoleRevoked)
				if err := _SupplyChain.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xe0df21b65c73c27081b8f042a012b124085b41d78d27b7e3c4780f5650f5ebb8.
//
// Solidity: event RoleRevoked(uint256 role, address account, address issuer)
func (_SupplyChain *SupplyChainFilterer) ParseRoleRevoked(log types.Log) (*SupplyChainRoleRevoked, error) {
	event := new(SupplyChainRoleRevoked)
	if err := _SupplyChain.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SupplyChainSignatureReceivedIterator is returned from FilterSignatureReceived and is used to iterate over the raw logs and unpacked data for SignatureReceived events raised by the SupplyChain contract.
type SupplyChainSignatureReceivedIterator struct {
	Event *SupplyChainSignatureReceived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SupplyChainSignatureReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyChainSignatureReceived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SupplyChainSignatureReceived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SupplyChainSignatureReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyChainSignatureReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyChainSignatureReceived represents a SignatureReceived event raised by the SupplyChain contract.
type SupplyChainSignatureReceived struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignatureReceived is a free log retrieval operation binding the contract event 0x45daca243d56fe35e7e5efbf15f8ecd64486be00404d4d2d84aa6c7103127673.
//
// Solidity: event SignatureReceived(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SupplyChain *SupplyChainFilterer) FilterSignatureReceived(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*SupplyChainSignatureReceivedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SupplyChain.contract.FilterLogs(opts, "SignatureReceived", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SupplyChainSignatureReceivedIterator{contract: _SupplyChain.contract, event: "SignatureReceived", logs: logs, sub: sub}, nil
}

// WatchSignatureReceived is a free log subscription operation binding the contract event 0x45daca243d56fe35e7e5efbf15f8ecd64486be00404d4d2d84aa6c7103127673.
//
// Solidity: event SignatureReceived(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SupplyChain *SupplyChainFilterer) WatchSignatureReceived(opts *bind.WatchOpts, sink chan<- *SupplyChainSignatureReceived, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SupplyChain.contract.WatchLogs(opts, "SignatureReceived", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyChainSignatureReceived)
				if err := _SupplyChain.contract.UnpackLog(event, "SignatureReceived", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignatureReceived is a log parse operation binding the contract event 0x45daca243d56fe35e7e5efbf15f8ecd64486be00404d4d2d84aa6c7103127673.
//
// Solidity: event SignatureReceived(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SupplyChain *SupplyChainFilterer) ParseSignatureReceived(log types.Log) (*SupplyChainSignatureReceived, error) {
	event := new(SupplyChainSignatureReceived)
	if err := _SupplyChain.contract.UnpackLog(event, "SignatureReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SupplyChainTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SupplyChain contract.
type SupplyChainTransferIterator struct {
	Event *SupplyChainTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SupplyChainTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyChainTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SupplyChainTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SupplyChainTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyChainTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyChainTransfer represents a Transfer event raised by the SupplyChain contract.
type SupplyChainTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SupplyChain *SupplyChainFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*SupplyChainTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SupplyChain.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SupplyChainTransferIterator{contract: _SupplyChain.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SupplyChain *SupplyChainFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SupplyChainTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SupplyChain.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyChainTransfer)
				if err := _SupplyChain.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SupplyChain *SupplyChainFilterer) ParseTransfer(log types.Log) (*SupplyChainTransfer, error) {
	event := new(SupplyChainTransfer)
	if err := _SupplyChain.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SupplyChainTransferAssetIterator is returned from FilterTransferAsset and is used to iterate over the raw logs and unpacked data for TransferAsset events raised by the SupplyChain contract.
type SupplyChainTransferAssetIterator struct {
	Event *SupplyChainTransferAsset // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SupplyChainTransferAssetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyChainTransferAsset)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SupplyChainTransferAsset)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SupplyChainTransferAssetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyChainTransferAssetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyChainTransferAsset represents a TransferAsset event raised by the SupplyChain contract.
type SupplyChainTransferAsset struct {
	From    common.Address
	To      common.Address
	AssetId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransferAsset is a free log retrieval operation binding the contract event 0x428ab6cd67836d0fe93626e98108a1a2593b0bc06a1a87b59e437a17cbaf430c.
//
// Solidity: event TransferAsset(address indexed from, address indexed to, uint256 indexed assetId)
func (_SupplyChain *SupplyChainFilterer) FilterTransferAsset(opts *bind.FilterOpts, from []common.Address, to []common.Address, assetId []*big.Int) (*SupplyChainTransferAssetIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _SupplyChain.contract.FilterLogs(opts, "TransferAsset", fromRule, toRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &SupplyChainTransferAssetIterator{contract: _SupplyChain.contract, event: "TransferAsset", logs: logs, sub: sub}, nil
}

// WatchTransferAsset is a free log subscription operation binding the contract event 0x428ab6cd67836d0fe93626e98108a1a2593b0bc06a1a87b59e437a17cbaf430c.
//
// Solidity: event TransferAsset(address indexed from, address indexed to, uint256 indexed assetId)
func (_SupplyChain *SupplyChainFilterer) WatchTransferAsset(opts *bind.WatchOpts, sink chan<- *SupplyChainTransferAsset, from []common.Address, to []common.Address, assetId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _SupplyChain.contract.WatchLogs(opts, "TransferAsset", fromRule, toRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyChainTransferAsset)
				if err := _SupplyChain.contract.UnpackLog(event, "TransferAsset", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferAsset is a log parse operation binding the contract event 0x428ab6cd67836d0fe93626e98108a1a2593b0bc06a1a87b59e437a17cbaf430c.
//
// Solidity: event TransferAsset(address indexed from, address indexed to, uint256 indexed assetId)
func (_SupplyChain *SupplyChainFilterer) ParseTransferAsset(log types.Log) (*SupplyChainTransferAsset, error) {
	event := new(SupplyChainTransferAsset)
	if err := _SupplyChain.contract.UnpackLog(event, "TransferAsset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SupplyChainUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the SupplyChain contract.
type SupplyChainUnpausedIterator struct {
	Event *SupplyChainUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SupplyChainUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyChainUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SupplyChainUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SupplyChainUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyChainUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyChainUnpaused represents a Unpaused event raised by the SupplyChain contract.
type SupplyChainUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SupplyChain *SupplyChainFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SupplyChainUnpausedIterator, error) {

	logs, sub, err := _SupplyChain.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SupplyChainUnpausedIterator{contract: _SupplyChain.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SupplyChain *SupplyChainFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SupplyChainUnpaused) (event.Subscription, error) {

	logs, sub, err := _SupplyChain.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyChainUnpaused)
				if err := _SupplyChain.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SupplyChain *SupplyChainFilterer) ParseUnpaused(log types.Log) (*SupplyChainUnpaused, error) {
	event := new(SupplyChainUnpaused)
	if err := _SupplyChain.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
