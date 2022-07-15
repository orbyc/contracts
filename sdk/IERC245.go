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

// IERC245MetaData contains all meta data concerning the IERC245 contract.
var IERC245MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"certId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"AssetCertificateAssigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"AssetIssued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"certId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"CertIssued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"moveId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"MovementAssigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"certId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"moveId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"MovementCertificateAssigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"moveId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"MovementIssued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"parentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"ParentAssigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"TransferAsset\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"certificates\",\"type\":\"uint256[]\"}],\"name\":\"addAssetCertificates\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"moveId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"certificates\",\"type\":\"uint256[]\"}],\"name\":\"addMovementCertificates\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"movements\",\"type\":\"uint256[]\"}],\"name\":\"addMovements\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"parents\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"composition\",\"type\":\"uint16[]\"}],\"name\":\"addParents\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"assetCertificates\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"assetComposition\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"\",\"type\":\"uint16[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"assetInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"assetTraceability\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"certId\",\"type\":\"uint256\"}],\"name\":\"certificateInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"co2e\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"certId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata_\",\"type\":\"string\"}],\"name\":\"issueAsset\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"certId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata_\",\"type\":\"string\"}],\"name\":\"issueCertificate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"moveId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"co2e\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"certId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata_\",\"type\":\"string\"}],\"name\":\"issueMovement\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"moveId\",\"type\":\"uint256\"}],\"name\":\"movementCertificates\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"moveId\",\"type\":\"uint256\"}],\"name\":\"movementInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC245ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC245MetaData.ABI instead.
var IERC245ABI = IERC245MetaData.ABI

// IERC245 is an auto generated Go binding around an Ethereum contract.
type IERC245 struct {
	IERC245Caller     // Read-only binding to the contract
	IERC245Transactor // Write-only binding to the contract
	IERC245Filterer   // Log filterer for contract events
}

// IERC245Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC245Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC245Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC245Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC245Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC245Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC245Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC245Session struct {
	Contract     *IERC245          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC245CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC245CallerSession struct {
	Contract *IERC245Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC245TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC245TransactorSession struct {
	Contract     *IERC245Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC245Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC245Raw struct {
	Contract *IERC245 // Generic contract binding to access the raw methods on
}

// IERC245CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC245CallerRaw struct {
	Contract *IERC245Caller // Generic read-only contract binding to access the raw methods on
}

// IERC245TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC245TransactorRaw struct {
	Contract *IERC245Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC245 creates a new instance of IERC245, bound to a specific deployed contract.
func NewIERC245(address common.Address, backend bind.ContractBackend) (*IERC245, error) {
	contract, err := bindIERC245(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC245{IERC245Caller: IERC245Caller{contract: contract}, IERC245Transactor: IERC245Transactor{contract: contract}, IERC245Filterer: IERC245Filterer{contract: contract}}, nil
}

// NewIERC245Caller creates a new read-only instance of IERC245, bound to a specific deployed contract.
func NewIERC245Caller(address common.Address, caller bind.ContractCaller) (*IERC245Caller, error) {
	contract, err := bindIERC245(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC245Caller{contract: contract}, nil
}

// NewIERC245Transactor creates a new write-only instance of IERC245, bound to a specific deployed contract.
func NewIERC245Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC245Transactor, error) {
	contract, err := bindIERC245(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC245Transactor{contract: contract}, nil
}

// NewIERC245Filterer creates a new log filterer instance of IERC245, bound to a specific deployed contract.
func NewIERC245Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC245Filterer, error) {
	contract, err := bindIERC245(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC245Filterer{contract: contract}, nil
}

// bindIERC245 binds a generic wrapper to an already deployed contract.
func bindIERC245(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC245ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC245 *IERC245Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC245.Contract.IERC245Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC245 *IERC245Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC245.Contract.IERC245Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC245 *IERC245Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC245.Contract.IERC245Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC245 *IERC245CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC245.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC245 *IERC245TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC245.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC245 *IERC245TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC245.Contract.contract.Transact(opts, method, params...)
}

// AssetCertificates is a free data retrieval call binding the contract method 0xc22d9980.
//
// Solidity: function assetCertificates(uint256 assetId) view returns(uint256[])
func (_IERC245 *IERC245Caller) AssetCertificates(opts *bind.CallOpts, assetId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _IERC245.contract.Call(opts, &out, "assetCertificates", assetId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AssetCertificates is a free data retrieval call binding the contract method 0xc22d9980.
//
// Solidity: function assetCertificates(uint256 assetId) view returns(uint256[])
func (_IERC245 *IERC245Session) AssetCertificates(assetId *big.Int) ([]*big.Int, error) {
	return _IERC245.Contract.AssetCertificates(&_IERC245.CallOpts, assetId)
}

// AssetCertificates is a free data retrieval call binding the contract method 0xc22d9980.
//
// Solidity: function assetCertificates(uint256 assetId) view returns(uint256[])
func (_IERC245 *IERC245CallerSession) AssetCertificates(assetId *big.Int) ([]*big.Int, error) {
	return _IERC245.Contract.AssetCertificates(&_IERC245.CallOpts, assetId)
}

// AssetComposition is a free data retrieval call binding the contract method 0x2192bc2a.
//
// Solidity: function assetComposition(uint256 assetId) view returns(uint256[], uint16[])
func (_IERC245 *IERC245Caller) AssetComposition(opts *bind.CallOpts, assetId *big.Int) ([]*big.Int, []uint16, error) {
	var out []interface{}
	err := _IERC245.contract.Call(opts, &out, "assetComposition", assetId)

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
func (_IERC245 *IERC245Session) AssetComposition(assetId *big.Int) ([]*big.Int, []uint16, error) {
	return _IERC245.Contract.AssetComposition(&_IERC245.CallOpts, assetId)
}

// AssetComposition is a free data retrieval call binding the contract method 0x2192bc2a.
//
// Solidity: function assetComposition(uint256 assetId) view returns(uint256[], uint16[])
func (_IERC245 *IERC245CallerSession) AssetComposition(assetId *big.Int) ([]*big.Int, []uint16, error) {
	return _IERC245.Contract.AssetComposition(&_IERC245.CallOpts, assetId)
}

// AssetInfo is a free data retrieval call binding the contract method 0xa879fcbb.
//
// Solidity: function assetInfo(uint256 assetId) view returns(address, address, uint64, uint256, string)
func (_IERC245 *IERC245Caller) AssetInfo(opts *bind.CallOpts, assetId *big.Int) (common.Address, common.Address, uint64, *big.Int, string, error) {
	var out []interface{}
	err := _IERC245.contract.Call(opts, &out, "assetInfo", assetId)

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
func (_IERC245 *IERC245Session) AssetInfo(assetId *big.Int) (common.Address, common.Address, uint64, *big.Int, string, error) {
	return _IERC245.Contract.AssetInfo(&_IERC245.CallOpts, assetId)
}

// AssetInfo is a free data retrieval call binding the contract method 0xa879fcbb.
//
// Solidity: function assetInfo(uint256 assetId) view returns(address, address, uint64, uint256, string)
func (_IERC245 *IERC245CallerSession) AssetInfo(assetId *big.Int) (common.Address, common.Address, uint64, *big.Int, string, error) {
	return _IERC245.Contract.AssetInfo(&_IERC245.CallOpts, assetId)
}

// AssetTraceability is a free data retrieval call binding the contract method 0xf6863e83.
//
// Solidity: function assetTraceability(uint256 assetId) view returns(uint256[])
func (_IERC245 *IERC245Caller) AssetTraceability(opts *bind.CallOpts, assetId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _IERC245.contract.Call(opts, &out, "assetTraceability", assetId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// AssetTraceability is a free data retrieval call binding the contract method 0xf6863e83.
//
// Solidity: function assetTraceability(uint256 assetId) view returns(uint256[])
func (_IERC245 *IERC245Session) AssetTraceability(assetId *big.Int) ([]*big.Int, error) {
	return _IERC245.Contract.AssetTraceability(&_IERC245.CallOpts, assetId)
}

// AssetTraceability is a free data retrieval call binding the contract method 0xf6863e83.
//
// Solidity: function assetTraceability(uint256 assetId) view returns(uint256[])
func (_IERC245 *IERC245CallerSession) AssetTraceability(assetId *big.Int) ([]*big.Int, error) {
	return _IERC245.Contract.AssetTraceability(&_IERC245.CallOpts, assetId)
}

// CertificateInfo is a free data retrieval call binding the contract method 0xba0a9f50.
//
// Solidity: function certificateInfo(uint256 certId) view returns(address, string)
func (_IERC245 *IERC245Caller) CertificateInfo(opts *bind.CallOpts, certId *big.Int) (common.Address, string, error) {
	var out []interface{}
	err := _IERC245.contract.Call(opts, &out, "certificateInfo", certId)

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
func (_IERC245 *IERC245Session) CertificateInfo(certId *big.Int) (common.Address, string, error) {
	return _IERC245.Contract.CertificateInfo(&_IERC245.CallOpts, certId)
}

// CertificateInfo is a free data retrieval call binding the contract method 0xba0a9f50.
//
// Solidity: function certificateInfo(uint256 certId) view returns(address, string)
func (_IERC245 *IERC245CallerSession) CertificateInfo(certId *big.Int) (common.Address, string, error) {
	return _IERC245.Contract.CertificateInfo(&_IERC245.CallOpts, certId)
}

// MovementCertificates is a free data retrieval call binding the contract method 0xd1ac4d65.
//
// Solidity: function movementCertificates(uint256 moveId) view returns(uint256[])
func (_IERC245 *IERC245Caller) MovementCertificates(opts *bind.CallOpts, moveId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _IERC245.contract.Call(opts, &out, "movementCertificates", moveId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// MovementCertificates is a free data retrieval call binding the contract method 0xd1ac4d65.
//
// Solidity: function movementCertificates(uint256 moveId) view returns(uint256[])
func (_IERC245 *IERC245Session) MovementCertificates(moveId *big.Int) ([]*big.Int, error) {
	return _IERC245.Contract.MovementCertificates(&_IERC245.CallOpts, moveId)
}

// MovementCertificates is a free data retrieval call binding the contract method 0xd1ac4d65.
//
// Solidity: function movementCertificates(uint256 moveId) view returns(uint256[])
func (_IERC245 *IERC245CallerSession) MovementCertificates(moveId *big.Int) ([]*big.Int, error) {
	return _IERC245.Contract.MovementCertificates(&_IERC245.CallOpts, moveId)
}

// MovementInfo is a free data retrieval call binding the contract method 0x3586ed2f.
//
// Solidity: function movementInfo(uint256 moveId) view returns(address, uint64, uint256, string)
func (_IERC245 *IERC245Caller) MovementInfo(opts *bind.CallOpts, moveId *big.Int) (common.Address, uint64, *big.Int, string, error) {
	var out []interface{}
	err := _IERC245.contract.Call(opts, &out, "movementInfo", moveId)

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
func (_IERC245 *IERC245Session) MovementInfo(moveId *big.Int) (common.Address, uint64, *big.Int, string, error) {
	return _IERC245.Contract.MovementInfo(&_IERC245.CallOpts, moveId)
}

// MovementInfo is a free data retrieval call binding the contract method 0x3586ed2f.
//
// Solidity: function movementInfo(uint256 moveId) view returns(address, uint64, uint256, string)
func (_IERC245 *IERC245CallerSession) MovementInfo(moveId *big.Int) (common.Address, uint64, *big.Int, string, error) {
	return _IERC245.Contract.MovementInfo(&_IERC245.CallOpts, moveId)
}

// AddAssetCertificates is a paid mutator transaction binding the contract method 0x4c40065e.
//
// Solidity: function addAssetCertificates(uint256 assetId, uint256[] certificates) returns(bool)
func (_IERC245 *IERC245Transactor) AddAssetCertificates(opts *bind.TransactOpts, assetId *big.Int, certificates []*big.Int) (*types.Transaction, error) {
	return _IERC245.contract.Transact(opts, "addAssetCertificates", assetId, certificates)
}

// AddAssetCertificates is a paid mutator transaction binding the contract method 0x4c40065e.
//
// Solidity: function addAssetCertificates(uint256 assetId, uint256[] certificates) returns(bool)
func (_IERC245 *IERC245Session) AddAssetCertificates(assetId *big.Int, certificates []*big.Int) (*types.Transaction, error) {
	return _IERC245.Contract.AddAssetCertificates(&_IERC245.TransactOpts, assetId, certificates)
}

// AddAssetCertificates is a paid mutator transaction binding the contract method 0x4c40065e.
//
// Solidity: function addAssetCertificates(uint256 assetId, uint256[] certificates) returns(bool)
func (_IERC245 *IERC245TransactorSession) AddAssetCertificates(assetId *big.Int, certificates []*big.Int) (*types.Transaction, error) {
	return _IERC245.Contract.AddAssetCertificates(&_IERC245.TransactOpts, assetId, certificates)
}

// AddMovementCertificates is a paid mutator transaction binding the contract method 0xaa2d16f1.
//
// Solidity: function addMovementCertificates(uint256 moveId, uint256[] certificates) returns(bool)
func (_IERC245 *IERC245Transactor) AddMovementCertificates(opts *bind.TransactOpts, moveId *big.Int, certificates []*big.Int) (*types.Transaction, error) {
	return _IERC245.contract.Transact(opts, "addMovementCertificates", moveId, certificates)
}

// AddMovementCertificates is a paid mutator transaction binding the contract method 0xaa2d16f1.
//
// Solidity: function addMovementCertificates(uint256 moveId, uint256[] certificates) returns(bool)
func (_IERC245 *IERC245Session) AddMovementCertificates(moveId *big.Int, certificates []*big.Int) (*types.Transaction, error) {
	return _IERC245.Contract.AddMovementCertificates(&_IERC245.TransactOpts, moveId, certificates)
}

// AddMovementCertificates is a paid mutator transaction binding the contract method 0xaa2d16f1.
//
// Solidity: function addMovementCertificates(uint256 moveId, uint256[] certificates) returns(bool)
func (_IERC245 *IERC245TransactorSession) AddMovementCertificates(moveId *big.Int, certificates []*big.Int) (*types.Transaction, error) {
	return _IERC245.Contract.AddMovementCertificates(&_IERC245.TransactOpts, moveId, certificates)
}

// AddMovements is a paid mutator transaction binding the contract method 0xe4eb771c.
//
// Solidity: function addMovements(uint256 assetId, uint256[] movements) returns(bool)
func (_IERC245 *IERC245Transactor) AddMovements(opts *bind.TransactOpts, assetId *big.Int, movements []*big.Int) (*types.Transaction, error) {
	return _IERC245.contract.Transact(opts, "addMovements", assetId, movements)
}

// AddMovements is a paid mutator transaction binding the contract method 0xe4eb771c.
//
// Solidity: function addMovements(uint256 assetId, uint256[] movements) returns(bool)
func (_IERC245 *IERC245Session) AddMovements(assetId *big.Int, movements []*big.Int) (*types.Transaction, error) {
	return _IERC245.Contract.AddMovements(&_IERC245.TransactOpts, assetId, movements)
}

// AddMovements is a paid mutator transaction binding the contract method 0xe4eb771c.
//
// Solidity: function addMovements(uint256 assetId, uint256[] movements) returns(bool)
func (_IERC245 *IERC245TransactorSession) AddMovements(assetId *big.Int, movements []*big.Int) (*types.Transaction, error) {
	return _IERC245.Contract.AddMovements(&_IERC245.TransactOpts, assetId, movements)
}

// AddParents is a paid mutator transaction binding the contract method 0xbad6a0ab.
//
// Solidity: function addParents(uint256 assetId, uint256[] parents, uint16[] composition) returns(bool)
func (_IERC245 *IERC245Transactor) AddParents(opts *bind.TransactOpts, assetId *big.Int, parents []*big.Int, composition []uint16) (*types.Transaction, error) {
	return _IERC245.contract.Transact(opts, "addParents", assetId, parents, composition)
}

// AddParents is a paid mutator transaction binding the contract method 0xbad6a0ab.
//
// Solidity: function addParents(uint256 assetId, uint256[] parents, uint16[] composition) returns(bool)
func (_IERC245 *IERC245Session) AddParents(assetId *big.Int, parents []*big.Int, composition []uint16) (*types.Transaction, error) {
	return _IERC245.Contract.AddParents(&_IERC245.TransactOpts, assetId, parents, composition)
}

// AddParents is a paid mutator transaction binding the contract method 0xbad6a0ab.
//
// Solidity: function addParents(uint256 assetId, uint256[] parents, uint16[] composition) returns(bool)
func (_IERC245 *IERC245TransactorSession) AddParents(assetId *big.Int, parents []*big.Int, composition []uint16) (*types.Transaction, error) {
	return _IERC245.Contract.AddParents(&_IERC245.TransactOpts, assetId, parents, composition)
}

// IssueAsset is a paid mutator transaction binding the contract method 0xf17011c6.
//
// Solidity: function issueAsset(uint256 assetId, address owner, uint64 co2e, uint256 certId, string metadata_) returns(bool)
func (_IERC245 *IERC245Transactor) IssueAsset(opts *bind.TransactOpts, assetId *big.Int, owner common.Address, co2e uint64, certId *big.Int, metadata_ string) (*types.Transaction, error) {
	return _IERC245.contract.Transact(opts, "issueAsset", assetId, owner, co2e, certId, metadata_)
}

// IssueAsset is a paid mutator transaction binding the contract method 0xf17011c6.
//
// Solidity: function issueAsset(uint256 assetId, address owner, uint64 co2e, uint256 certId, string metadata_) returns(bool)
func (_IERC245 *IERC245Session) IssueAsset(assetId *big.Int, owner common.Address, co2e uint64, certId *big.Int, metadata_ string) (*types.Transaction, error) {
	return _IERC245.Contract.IssueAsset(&_IERC245.TransactOpts, assetId, owner, co2e, certId, metadata_)
}

// IssueAsset is a paid mutator transaction binding the contract method 0xf17011c6.
//
// Solidity: function issueAsset(uint256 assetId, address owner, uint64 co2e, uint256 certId, string metadata_) returns(bool)
func (_IERC245 *IERC245TransactorSession) IssueAsset(assetId *big.Int, owner common.Address, co2e uint64, certId *big.Int, metadata_ string) (*types.Transaction, error) {
	return _IERC245.Contract.IssueAsset(&_IERC245.TransactOpts, assetId, owner, co2e, certId, metadata_)
}

// IssueCertificate is a paid mutator transaction binding the contract method 0x0c816de4.
//
// Solidity: function issueCertificate(uint256 certId, string metadata_) returns(bool)
func (_IERC245 *IERC245Transactor) IssueCertificate(opts *bind.TransactOpts, certId *big.Int, metadata_ string) (*types.Transaction, error) {
	return _IERC245.contract.Transact(opts, "issueCertificate", certId, metadata_)
}

// IssueCertificate is a paid mutator transaction binding the contract method 0x0c816de4.
//
// Solidity: function issueCertificate(uint256 certId, string metadata_) returns(bool)
func (_IERC245 *IERC245Session) IssueCertificate(certId *big.Int, metadata_ string) (*types.Transaction, error) {
	return _IERC245.Contract.IssueCertificate(&_IERC245.TransactOpts, certId, metadata_)
}

// IssueCertificate is a paid mutator transaction binding the contract method 0x0c816de4.
//
// Solidity: function issueCertificate(uint256 certId, string metadata_) returns(bool)
func (_IERC245 *IERC245TransactorSession) IssueCertificate(certId *big.Int, metadata_ string) (*types.Transaction, error) {
	return _IERC245.Contract.IssueCertificate(&_IERC245.TransactOpts, certId, metadata_)
}

// IssueMovement is a paid mutator transaction binding the contract method 0xe842eafa.
//
// Solidity: function issueMovement(uint256 moveId, uint64 co2e, uint256 certId, string metadata_) returns(bool)
func (_IERC245 *IERC245Transactor) IssueMovement(opts *bind.TransactOpts, moveId *big.Int, co2e uint64, certId *big.Int, metadata_ string) (*types.Transaction, error) {
	return _IERC245.contract.Transact(opts, "issueMovement", moveId, co2e, certId, metadata_)
}

// IssueMovement is a paid mutator transaction binding the contract method 0xe842eafa.
//
// Solidity: function issueMovement(uint256 moveId, uint64 co2e, uint256 certId, string metadata_) returns(bool)
func (_IERC245 *IERC245Session) IssueMovement(moveId *big.Int, co2e uint64, certId *big.Int, metadata_ string) (*types.Transaction, error) {
	return _IERC245.Contract.IssueMovement(&_IERC245.TransactOpts, moveId, co2e, certId, metadata_)
}

// IssueMovement is a paid mutator transaction binding the contract method 0xe842eafa.
//
// Solidity: function issueMovement(uint256 moveId, uint64 co2e, uint256 certId, string metadata_) returns(bool)
func (_IERC245 *IERC245TransactorSession) IssueMovement(moveId *big.Int, co2e uint64, certId *big.Int, metadata_ string) (*types.Transaction, error) {
	return _IERC245.Contract.IssueMovement(&_IERC245.TransactOpts, moveId, co2e, certId, metadata_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 assetId) returns()
func (_IERC245 *IERC245Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, assetId *big.Int) (*types.Transaction, error) {
	return _IERC245.contract.Transact(opts, "transferFrom", from, to, assetId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 assetId) returns()
func (_IERC245 *IERC245Session) TransferFrom(from common.Address, to common.Address, assetId *big.Int) (*types.Transaction, error) {
	return _IERC245.Contract.TransferFrom(&_IERC245.TransactOpts, from, to, assetId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 assetId) returns()
func (_IERC245 *IERC245TransactorSession) TransferFrom(from common.Address, to common.Address, assetId *big.Int) (*types.Transaction, error) {
	return _IERC245.Contract.TransferFrom(&_IERC245.TransactOpts, from, to, assetId)
}

// IERC245AssetCertificateAssignedIterator is returned from FilterAssetCertificateAssigned and is used to iterate over the raw logs and unpacked data for AssetCertificateAssigned events raised by the IERC245 contract.
type IERC245AssetCertificateAssignedIterator struct {
	Event *IERC245AssetCertificateAssigned // Event containing the contract specifics and raw log

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
func (it *IERC245AssetCertificateAssignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC245AssetCertificateAssigned)
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
		it.Event = new(IERC245AssetCertificateAssigned)
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
func (it *IERC245AssetCertificateAssignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC245AssetCertificateAssignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC245AssetCertificateAssigned represents a AssetCertificateAssigned event raised by the IERC245 contract.
type IERC245AssetCertificateAssigned struct {
	CertId  *big.Int
	AssetId *big.Int
	Signer  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAssetCertificateAssigned is a free log retrieval operation binding the contract event 0x4fc952580cfc5ff048a6b9111d8facdf029038c2a88b0ea0c42da67a1a5e00d1.
//
// Solidity: event AssetCertificateAssigned(uint256 certId, uint256 assetId, address signer)
func (_IERC245 *IERC245Filterer) FilterAssetCertificateAssigned(opts *bind.FilterOpts) (*IERC245AssetCertificateAssignedIterator, error) {

	logs, sub, err := _IERC245.contract.FilterLogs(opts, "AssetCertificateAssigned")
	if err != nil {
		return nil, err
	}
	return &IERC245AssetCertificateAssignedIterator{contract: _IERC245.contract, event: "AssetCertificateAssigned", logs: logs, sub: sub}, nil
}

// WatchAssetCertificateAssigned is a free log subscription operation binding the contract event 0x4fc952580cfc5ff048a6b9111d8facdf029038c2a88b0ea0c42da67a1a5e00d1.
//
// Solidity: event AssetCertificateAssigned(uint256 certId, uint256 assetId, address signer)
func (_IERC245 *IERC245Filterer) WatchAssetCertificateAssigned(opts *bind.WatchOpts, sink chan<- *IERC245AssetCertificateAssigned) (event.Subscription, error) {

	logs, sub, err := _IERC245.contract.WatchLogs(opts, "AssetCertificateAssigned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC245AssetCertificateAssigned)
				if err := _IERC245.contract.UnpackLog(event, "AssetCertificateAssigned", log); err != nil {
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
func (_IERC245 *IERC245Filterer) ParseAssetCertificateAssigned(log types.Log) (*IERC245AssetCertificateAssigned, error) {
	event := new(IERC245AssetCertificateAssigned)
	if err := _IERC245.contract.UnpackLog(event, "AssetCertificateAssigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC245AssetIssuedIterator is returned from FilterAssetIssued and is used to iterate over the raw logs and unpacked data for AssetIssued events raised by the IERC245 contract.
type IERC245AssetIssuedIterator struct {
	Event *IERC245AssetIssued // Event containing the contract specifics and raw log

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
func (it *IERC245AssetIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC245AssetIssued)
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
		it.Event = new(IERC245AssetIssued)
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
func (it *IERC245AssetIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC245AssetIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC245AssetIssued represents a AssetIssued event raised by the IERC245 contract.
type IERC245AssetIssued struct {
	AssetId *big.Int
	Signer  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAssetIssued is a free log retrieval operation binding the contract event 0xf37f680ac83ec9f2a328ebff310dfe6d49043380dcb563e89ef4ca09eaa38d51.
//
// Solidity: event AssetIssued(uint256 assetId, address indexed signer)
func (_IERC245 *IERC245Filterer) FilterAssetIssued(opts *bind.FilterOpts, signer []common.Address) (*IERC245AssetIssuedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _IERC245.contract.FilterLogs(opts, "AssetIssued", signerRule)
	if err != nil {
		return nil, err
	}
	return &IERC245AssetIssuedIterator{contract: _IERC245.contract, event: "AssetIssued", logs: logs, sub: sub}, nil
}

// WatchAssetIssued is a free log subscription operation binding the contract event 0xf37f680ac83ec9f2a328ebff310dfe6d49043380dcb563e89ef4ca09eaa38d51.
//
// Solidity: event AssetIssued(uint256 assetId, address indexed signer)
func (_IERC245 *IERC245Filterer) WatchAssetIssued(opts *bind.WatchOpts, sink chan<- *IERC245AssetIssued, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _IERC245.contract.WatchLogs(opts, "AssetIssued", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC245AssetIssued)
				if err := _IERC245.contract.UnpackLog(event, "AssetIssued", log); err != nil {
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
func (_IERC245 *IERC245Filterer) ParseAssetIssued(log types.Log) (*IERC245AssetIssued, error) {
	event := new(IERC245AssetIssued)
	if err := _IERC245.contract.UnpackLog(event, "AssetIssued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC245CertIssuedIterator is returned from FilterCertIssued and is used to iterate over the raw logs and unpacked data for CertIssued events raised by the IERC245 contract.
type IERC245CertIssuedIterator struct {
	Event *IERC245CertIssued // Event containing the contract specifics and raw log

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
func (it *IERC245CertIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC245CertIssued)
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
		it.Event = new(IERC245CertIssued)
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
func (it *IERC245CertIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC245CertIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC245CertIssued represents a CertIssued event raised by the IERC245 contract.
type IERC245CertIssued struct {
	CertId *big.Int
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCertIssued is a free log retrieval operation binding the contract event 0x0ac0bb5418edafe119eaae67d0ee0bc6f39a3efafc3305a56b617c171d0d835f.
//
// Solidity: event CertIssued(uint256 certId, address indexed signer)
func (_IERC245 *IERC245Filterer) FilterCertIssued(opts *bind.FilterOpts, signer []common.Address) (*IERC245CertIssuedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _IERC245.contract.FilterLogs(opts, "CertIssued", signerRule)
	if err != nil {
		return nil, err
	}
	return &IERC245CertIssuedIterator{contract: _IERC245.contract, event: "CertIssued", logs: logs, sub: sub}, nil
}

// WatchCertIssued is a free log subscription operation binding the contract event 0x0ac0bb5418edafe119eaae67d0ee0bc6f39a3efafc3305a56b617c171d0d835f.
//
// Solidity: event CertIssued(uint256 certId, address indexed signer)
func (_IERC245 *IERC245Filterer) WatchCertIssued(opts *bind.WatchOpts, sink chan<- *IERC245CertIssued, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _IERC245.contract.WatchLogs(opts, "CertIssued", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC245CertIssued)
				if err := _IERC245.contract.UnpackLog(event, "CertIssued", log); err != nil {
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
func (_IERC245 *IERC245Filterer) ParseCertIssued(log types.Log) (*IERC245CertIssued, error) {
	event := new(IERC245CertIssued)
	if err := _IERC245.contract.UnpackLog(event, "CertIssued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC245MovementAssignedIterator is returned from FilterMovementAssigned and is used to iterate over the raw logs and unpacked data for MovementAssigned events raised by the IERC245 contract.
type IERC245MovementAssignedIterator struct {
	Event *IERC245MovementAssigned // Event containing the contract specifics and raw log

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
func (it *IERC245MovementAssignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC245MovementAssigned)
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
		it.Event = new(IERC245MovementAssigned)
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
func (it *IERC245MovementAssignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC245MovementAssignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC245MovementAssigned represents a MovementAssigned event raised by the IERC245 contract.
type IERC245MovementAssigned struct {
	MoveId  *big.Int
	AssetId *big.Int
	Signer  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMovementAssigned is a free log retrieval operation binding the contract event 0x1fb5e37b2fee9c771f5a7ac82a15698c9f978e0e6cef7ac48ceb8b9792caaf16.
//
// Solidity: event MovementAssigned(uint256 moveId, uint256 assetId, address signer)
func (_IERC245 *IERC245Filterer) FilterMovementAssigned(opts *bind.FilterOpts) (*IERC245MovementAssignedIterator, error) {

	logs, sub, err := _IERC245.contract.FilterLogs(opts, "MovementAssigned")
	if err != nil {
		return nil, err
	}
	return &IERC245MovementAssignedIterator{contract: _IERC245.contract, event: "MovementAssigned", logs: logs, sub: sub}, nil
}

// WatchMovementAssigned is a free log subscription operation binding the contract event 0x1fb5e37b2fee9c771f5a7ac82a15698c9f978e0e6cef7ac48ceb8b9792caaf16.
//
// Solidity: event MovementAssigned(uint256 moveId, uint256 assetId, address signer)
func (_IERC245 *IERC245Filterer) WatchMovementAssigned(opts *bind.WatchOpts, sink chan<- *IERC245MovementAssigned) (event.Subscription, error) {

	logs, sub, err := _IERC245.contract.WatchLogs(opts, "MovementAssigned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC245MovementAssigned)
				if err := _IERC245.contract.UnpackLog(event, "MovementAssigned", log); err != nil {
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
func (_IERC245 *IERC245Filterer) ParseMovementAssigned(log types.Log) (*IERC245MovementAssigned, error) {
	event := new(IERC245MovementAssigned)
	if err := _IERC245.contract.UnpackLog(event, "MovementAssigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC245MovementCertificateAssignedIterator is returned from FilterMovementCertificateAssigned and is used to iterate over the raw logs and unpacked data for MovementCertificateAssigned events raised by the IERC245 contract.
type IERC245MovementCertificateAssignedIterator struct {
	Event *IERC245MovementCertificateAssigned // Event containing the contract specifics and raw log

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
func (it *IERC245MovementCertificateAssignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC245MovementCertificateAssigned)
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
		it.Event = new(IERC245MovementCertificateAssigned)
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
func (it *IERC245MovementCertificateAssignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC245MovementCertificateAssignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC245MovementCertificateAssigned represents a MovementCertificateAssigned event raised by the IERC245 contract.
type IERC245MovementCertificateAssigned struct {
	CertId *big.Int
	MoveId *big.Int
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMovementCertificateAssigned is a free log retrieval operation binding the contract event 0xb8730f75259bf8d0baccdacc5bc1580e45d23db9e5a88f49ba8bc829c14f64f8.
//
// Solidity: event MovementCertificateAssigned(uint256 certId, uint256 moveId, address signer)
func (_IERC245 *IERC245Filterer) FilterMovementCertificateAssigned(opts *bind.FilterOpts) (*IERC245MovementCertificateAssignedIterator, error) {

	logs, sub, err := _IERC245.contract.FilterLogs(opts, "MovementCertificateAssigned")
	if err != nil {
		return nil, err
	}
	return &IERC245MovementCertificateAssignedIterator{contract: _IERC245.contract, event: "MovementCertificateAssigned", logs: logs, sub: sub}, nil
}

// WatchMovementCertificateAssigned is a free log subscription operation binding the contract event 0xb8730f75259bf8d0baccdacc5bc1580e45d23db9e5a88f49ba8bc829c14f64f8.
//
// Solidity: event MovementCertificateAssigned(uint256 certId, uint256 moveId, address signer)
func (_IERC245 *IERC245Filterer) WatchMovementCertificateAssigned(opts *bind.WatchOpts, sink chan<- *IERC245MovementCertificateAssigned) (event.Subscription, error) {

	logs, sub, err := _IERC245.contract.WatchLogs(opts, "MovementCertificateAssigned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC245MovementCertificateAssigned)
				if err := _IERC245.contract.UnpackLog(event, "MovementCertificateAssigned", log); err != nil {
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
func (_IERC245 *IERC245Filterer) ParseMovementCertificateAssigned(log types.Log) (*IERC245MovementCertificateAssigned, error) {
	event := new(IERC245MovementCertificateAssigned)
	if err := _IERC245.contract.UnpackLog(event, "MovementCertificateAssigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC245MovementIssuedIterator is returned from FilterMovementIssued and is used to iterate over the raw logs and unpacked data for MovementIssued events raised by the IERC245 contract.
type IERC245MovementIssuedIterator struct {
	Event *IERC245MovementIssued // Event containing the contract specifics and raw log

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
func (it *IERC245MovementIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC245MovementIssued)
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
		it.Event = new(IERC245MovementIssued)
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
func (it *IERC245MovementIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC245MovementIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC245MovementIssued represents a MovementIssued event raised by the IERC245 contract.
type IERC245MovementIssued struct {
	MoveId *big.Int
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMovementIssued is a free log retrieval operation binding the contract event 0x741b29ae377b3381e554e394b7d2e2fefb4db0c40f73b9e0ac4cd843ef56171d.
//
// Solidity: event MovementIssued(uint256 moveId, address indexed signer)
func (_IERC245 *IERC245Filterer) FilterMovementIssued(opts *bind.FilterOpts, signer []common.Address) (*IERC245MovementIssuedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _IERC245.contract.FilterLogs(opts, "MovementIssued", signerRule)
	if err != nil {
		return nil, err
	}
	return &IERC245MovementIssuedIterator{contract: _IERC245.contract, event: "MovementIssued", logs: logs, sub: sub}, nil
}

// WatchMovementIssued is a free log subscription operation binding the contract event 0x741b29ae377b3381e554e394b7d2e2fefb4db0c40f73b9e0ac4cd843ef56171d.
//
// Solidity: event MovementIssued(uint256 moveId, address indexed signer)
func (_IERC245 *IERC245Filterer) WatchMovementIssued(opts *bind.WatchOpts, sink chan<- *IERC245MovementIssued, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _IERC245.contract.WatchLogs(opts, "MovementIssued", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC245MovementIssued)
				if err := _IERC245.contract.UnpackLog(event, "MovementIssued", log); err != nil {
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
func (_IERC245 *IERC245Filterer) ParseMovementIssued(log types.Log) (*IERC245MovementIssued, error) {
	event := new(IERC245MovementIssued)
	if err := _IERC245.contract.UnpackLog(event, "MovementIssued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC245ParentAssignedIterator is returned from FilterParentAssigned and is used to iterate over the raw logs and unpacked data for ParentAssigned events raised by the IERC245 contract.
type IERC245ParentAssignedIterator struct {
	Event *IERC245ParentAssigned // Event containing the contract specifics and raw log

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
func (it *IERC245ParentAssignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC245ParentAssigned)
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
		it.Event = new(IERC245ParentAssigned)
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
func (it *IERC245ParentAssignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC245ParentAssignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC245ParentAssigned represents a ParentAssigned event raised by the IERC245 contract.
type IERC245ParentAssigned struct {
	ParentId *big.Int
	AssetId  *big.Int
	Signer   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterParentAssigned is a free log retrieval operation binding the contract event 0x4154f07b27a742714e27b3894179c9cb5aba0e2a7fa26cb0f7513abecf1b051b.
//
// Solidity: event ParentAssigned(uint256 parentId, uint256 assetId, address signer)
func (_IERC245 *IERC245Filterer) FilterParentAssigned(opts *bind.FilterOpts) (*IERC245ParentAssignedIterator, error) {

	logs, sub, err := _IERC245.contract.FilterLogs(opts, "ParentAssigned")
	if err != nil {
		return nil, err
	}
	return &IERC245ParentAssignedIterator{contract: _IERC245.contract, event: "ParentAssigned", logs: logs, sub: sub}, nil
}

// WatchParentAssigned is a free log subscription operation binding the contract event 0x4154f07b27a742714e27b3894179c9cb5aba0e2a7fa26cb0f7513abecf1b051b.
//
// Solidity: event ParentAssigned(uint256 parentId, uint256 assetId, address signer)
func (_IERC245 *IERC245Filterer) WatchParentAssigned(opts *bind.WatchOpts, sink chan<- *IERC245ParentAssigned) (event.Subscription, error) {

	logs, sub, err := _IERC245.contract.WatchLogs(opts, "ParentAssigned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC245ParentAssigned)
				if err := _IERC245.contract.UnpackLog(event, "ParentAssigned", log); err != nil {
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
func (_IERC245 *IERC245Filterer) ParseParentAssigned(log types.Log) (*IERC245ParentAssigned, error) {
	event := new(IERC245ParentAssigned)
	if err := _IERC245.contract.UnpackLog(event, "ParentAssigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC245TransferAssetIterator is returned from FilterTransferAsset and is used to iterate over the raw logs and unpacked data for TransferAsset events raised by the IERC245 contract.
type IERC245TransferAssetIterator struct {
	Event *IERC245TransferAsset // Event containing the contract specifics and raw log

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
func (it *IERC245TransferAssetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC245TransferAsset)
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
		it.Event = new(IERC245TransferAsset)
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
func (it *IERC245TransferAssetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC245TransferAssetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC245TransferAsset represents a TransferAsset event raised by the IERC245 contract.
type IERC245TransferAsset struct {
	From    common.Address
	To      common.Address
	AssetId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransferAsset is a free log retrieval operation binding the contract event 0x428ab6cd67836d0fe93626e98108a1a2593b0bc06a1a87b59e437a17cbaf430c.
//
// Solidity: event TransferAsset(address indexed from, address indexed to, uint256 indexed assetId)
func (_IERC245 *IERC245Filterer) FilterTransferAsset(opts *bind.FilterOpts, from []common.Address, to []common.Address, assetId []*big.Int) (*IERC245TransferAssetIterator, error) {

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

	logs, sub, err := _IERC245.contract.FilterLogs(opts, "TransferAsset", fromRule, toRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC245TransferAssetIterator{contract: _IERC245.contract, event: "TransferAsset", logs: logs, sub: sub}, nil
}

// WatchTransferAsset is a free log subscription operation binding the contract event 0x428ab6cd67836d0fe93626e98108a1a2593b0bc06a1a87b59e437a17cbaf430c.
//
// Solidity: event TransferAsset(address indexed from, address indexed to, uint256 indexed assetId)
func (_IERC245 *IERC245Filterer) WatchTransferAsset(opts *bind.WatchOpts, sink chan<- *IERC245TransferAsset, from []common.Address, to []common.Address, assetId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _IERC245.contract.WatchLogs(opts, "TransferAsset", fromRule, toRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC245TransferAsset)
				if err := _IERC245.contract.UnpackLog(event, "TransferAsset", log); err != nil {
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
func (_IERC245 *IERC245Filterer) ParseTransferAsset(log types.Log) (*IERC245TransferAsset, error) {
	event := new(IERC245TransferAsset)
	if err := _IERC245.contract.UnpackLog(event, "TransferAsset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
