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

// IERC423MetaData contains all meta data concerning the IERC423 contract.
var IERC423MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"AgentDefined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"AgentRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"RoleDefined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"accountInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"accountOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadata_\",\"type\":\"string\"}],\"name\":\"defineAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata_\",\"type\":\"string\"}],\"name\":\"defineRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"grantRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"removeAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"revokeRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"roleInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IERC423ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC423MetaData.ABI instead.
var IERC423ABI = IERC423MetaData.ABI

// IERC423 is an auto generated Go binding around an Ethereum contract.
type IERC423 struct {
	IERC423Caller     // Read-only binding to the contract
	IERC423Transactor // Write-only binding to the contract
	IERC423Filterer   // Log filterer for contract events
}

// IERC423Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC423Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC423Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC423Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC423Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC423Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC423Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC423Session struct {
	Contract     *IERC423          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC423CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC423CallerSession struct {
	Contract *IERC423Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC423TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC423TransactorSession struct {
	Contract     *IERC423Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC423Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC423Raw struct {
	Contract *IERC423 // Generic contract binding to access the raw methods on
}

// IERC423CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC423CallerRaw struct {
	Contract *IERC423Caller // Generic read-only contract binding to access the raw methods on
}

// IERC423TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC423TransactorRaw struct {
	Contract *IERC423Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC423 creates a new instance of IERC423, bound to a specific deployed contract.
func NewIERC423(address common.Address, backend bind.ContractBackend) (*IERC423, error) {
	contract, err := bindIERC423(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC423{IERC423Caller: IERC423Caller{contract: contract}, IERC423Transactor: IERC423Transactor{contract: contract}, IERC423Filterer: IERC423Filterer{contract: contract}}, nil
}

// NewIERC423Caller creates a new read-only instance of IERC423, bound to a specific deployed contract.
func NewIERC423Caller(address common.Address, caller bind.ContractCaller) (*IERC423Caller, error) {
	contract, err := bindIERC423(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC423Caller{contract: contract}, nil
}

// NewIERC423Transactor creates a new write-only instance of IERC423, bound to a specific deployed contract.
func NewIERC423Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC423Transactor, error) {
	contract, err := bindIERC423(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC423Transactor{contract: contract}, nil
}

// NewIERC423Filterer creates a new log filterer instance of IERC423, bound to a specific deployed contract.
func NewIERC423Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC423Filterer, error) {
	contract, err := bindIERC423(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC423Filterer{contract: contract}, nil
}

// bindIERC423 binds a generic wrapper to an already deployed contract.
func bindIERC423(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC423ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC423 *IERC423Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC423.Contract.IERC423Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC423 *IERC423Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC423.Contract.IERC423Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC423 *IERC423Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC423.Contract.IERC423Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC423 *IERC423CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC423.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC423 *IERC423TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC423.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC423 *IERC423TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC423.Contract.contract.Transact(opts, method, params...)
}

// AccountInfo is a free data retrieval call binding the contract method 0xa7310b58.
//
// Solidity: function accountInfo(address account) view returns(string)
func (_IERC423 *IERC423Caller) AccountInfo(opts *bind.CallOpts, account common.Address) (string, error) {
	var out []interface{}
	err := _IERC423.contract.Call(opts, &out, "accountInfo", account)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AccountInfo is a free data retrieval call binding the contract method 0xa7310b58.
//
// Solidity: function accountInfo(address account) view returns(string)
func (_IERC423 *IERC423Session) AccountInfo(account common.Address) (string, error) {
	return _IERC423.Contract.AccountInfo(&_IERC423.CallOpts, account)
}

// AccountInfo is a free data retrieval call binding the contract method 0xa7310b58.
//
// Solidity: function accountInfo(address account) view returns(string)
func (_IERC423 *IERC423CallerSession) AccountInfo(account common.Address) (string, error) {
	return _IERC423.Contract.AccountInfo(&_IERC423.CallOpts, account)
}

// AccountOf is a free data retrieval call binding the contract method 0x8086b8ba.
//
// Solidity: function accountOf(address agent) view returns(address)
func (_IERC423 *IERC423Caller) AccountOf(opts *bind.CallOpts, agent common.Address) (common.Address, error) {
	var out []interface{}
	err := _IERC423.contract.Call(opts, &out, "accountOf", agent)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountOf is a free data retrieval call binding the contract method 0x8086b8ba.
//
// Solidity: function accountOf(address agent) view returns(address)
func (_IERC423 *IERC423Session) AccountOf(agent common.Address) (common.Address, error) {
	return _IERC423.Contract.AccountOf(&_IERC423.CallOpts, agent)
}

// AccountOf is a free data retrieval call binding the contract method 0x8086b8ba.
//
// Solidity: function accountOf(address agent) view returns(address)
func (_IERC423 *IERC423CallerSession) AccountOf(agent common.Address) (common.Address, error) {
	return _IERC423.Contract.AccountOf(&_IERC423.CallOpts, agent)
}

// HasRole is a free data retrieval call binding the contract method 0x5c97f4a2.
//
// Solidity: function hasRole(address account, uint256 role) view returns(bool)
func (_IERC423 *IERC423Caller) HasRole(opts *bind.CallOpts, account common.Address, role *big.Int) (bool, error) {
	var out []interface{}
	err := _IERC423.contract.Call(opts, &out, "hasRole", account, role)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x5c97f4a2.
//
// Solidity: function hasRole(address account, uint256 role) view returns(bool)
func (_IERC423 *IERC423Session) HasRole(account common.Address, role *big.Int) (bool, error) {
	return _IERC423.Contract.HasRole(&_IERC423.CallOpts, account, role)
}

// HasRole is a free data retrieval call binding the contract method 0x5c97f4a2.
//
// Solidity: function hasRole(address account, uint256 role) view returns(bool)
func (_IERC423 *IERC423CallerSession) HasRole(account common.Address, role *big.Int) (bool, error) {
	return _IERC423.Contract.HasRole(&_IERC423.CallOpts, account, role)
}

// RoleInfo is a free data retrieval call binding the contract method 0x93123e0b.
//
// Solidity: function roleInfo(uint256 role) view returns(string)
func (_IERC423 *IERC423Caller) RoleInfo(opts *bind.CallOpts, role *big.Int) (string, error) {
	var out []interface{}
	err := _IERC423.contract.Call(opts, &out, "roleInfo", role)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RoleInfo is a free data retrieval call binding the contract method 0x93123e0b.
//
// Solidity: function roleInfo(uint256 role) view returns(string)
func (_IERC423 *IERC423Session) RoleInfo(role *big.Int) (string, error) {
	return _IERC423.Contract.RoleInfo(&_IERC423.CallOpts, role)
}

// RoleInfo is a free data retrieval call binding the contract method 0x93123e0b.
//
// Solidity: function roleInfo(uint256 role) view returns(string)
func (_IERC423 *IERC423CallerSession) RoleInfo(role *big.Int) (string, error) {
	return _IERC423.Contract.RoleInfo(&_IERC423.CallOpts, role)
}

// DefineAgent is a paid mutator transaction binding the contract method 0x5ee3bc2a.
//
// Solidity: function defineAgent(address agent, address account, string metadata_) returns(bool)
func (_IERC423 *IERC423Transactor) DefineAgent(opts *bind.TransactOpts, agent common.Address, account common.Address, metadata_ string) (*types.Transaction, error) {
	return _IERC423.contract.Transact(opts, "defineAgent", agent, account, metadata_)
}

// DefineAgent is a paid mutator transaction binding the contract method 0x5ee3bc2a.
//
// Solidity: function defineAgent(address agent, address account, string metadata_) returns(bool)
func (_IERC423 *IERC423Session) DefineAgent(agent common.Address, account common.Address, metadata_ string) (*types.Transaction, error) {
	return _IERC423.Contract.DefineAgent(&_IERC423.TransactOpts, agent, account, metadata_)
}

// DefineAgent is a paid mutator transaction binding the contract method 0x5ee3bc2a.
//
// Solidity: function defineAgent(address agent, address account, string metadata_) returns(bool)
func (_IERC423 *IERC423TransactorSession) DefineAgent(agent common.Address, account common.Address, metadata_ string) (*types.Transaction, error) {
	return _IERC423.Contract.DefineAgent(&_IERC423.TransactOpts, agent, account, metadata_)
}

// DefineRole is a paid mutator transaction binding the contract method 0xb77e76c2.
//
// Solidity: function defineRole(uint256 role, string metadata_) returns(bool)
func (_IERC423 *IERC423Transactor) DefineRole(opts *bind.TransactOpts, role *big.Int, metadata_ string) (*types.Transaction, error) {
	return _IERC423.contract.Transact(opts, "defineRole", role, metadata_)
}

// DefineRole is a paid mutator transaction binding the contract method 0xb77e76c2.
//
// Solidity: function defineRole(uint256 role, string metadata_) returns(bool)
func (_IERC423 *IERC423Session) DefineRole(role *big.Int, metadata_ string) (*types.Transaction, error) {
	return _IERC423.Contract.DefineRole(&_IERC423.TransactOpts, role, metadata_)
}

// DefineRole is a paid mutator transaction binding the contract method 0xb77e76c2.
//
// Solidity: function defineRole(uint256 role, string metadata_) returns(bool)
func (_IERC423 *IERC423TransactorSession) DefineRole(role *big.Int, metadata_ string) (*types.Transaction, error) {
	return _IERC423.Contract.DefineRole(&_IERC423.TransactOpts, role, metadata_)
}

// GrantRole is a paid mutator transaction binding the contract method 0x3c09e2fd.
//
// Solidity: function grantRole(address account, uint256 role) returns(bool)
func (_IERC423 *IERC423Transactor) GrantRole(opts *bind.TransactOpts, account common.Address, role *big.Int) (*types.Transaction, error) {
	return _IERC423.contract.Transact(opts, "grantRole", account, role)
}

// GrantRole is a paid mutator transaction binding the contract method 0x3c09e2fd.
//
// Solidity: function grantRole(address account, uint256 role) returns(bool)
func (_IERC423 *IERC423Session) GrantRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _IERC423.Contract.GrantRole(&_IERC423.TransactOpts, account, role)
}

// GrantRole is a paid mutator transaction binding the contract method 0x3c09e2fd.
//
// Solidity: function grantRole(address account, uint256 role) returns(bool)
func (_IERC423 *IERC423TransactorSession) GrantRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _IERC423.Contract.GrantRole(&_IERC423.TransactOpts, account, role)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0x97a6278e.
//
// Solidity: function removeAgent(address agent) returns(bool)
func (_IERC423 *IERC423Transactor) RemoveAgent(opts *bind.TransactOpts, agent common.Address) (*types.Transaction, error) {
	return _IERC423.contract.Transact(opts, "removeAgent", agent)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0x97a6278e.
//
// Solidity: function removeAgent(address agent) returns(bool)
func (_IERC423 *IERC423Session) RemoveAgent(agent common.Address) (*types.Transaction, error) {
	return _IERC423.Contract.RemoveAgent(&_IERC423.TransactOpts, agent)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0x97a6278e.
//
// Solidity: function removeAgent(address agent) returns(bool)
func (_IERC423 *IERC423TransactorSession) RemoveAgent(agent common.Address) (*types.Transaction, error) {
	return _IERC423.Contract.RemoveAgent(&_IERC423.TransactOpts, agent)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x0912ed77.
//
// Solidity: function revokeRole(address account, uint256 role) returns(bool)
func (_IERC423 *IERC423Transactor) RevokeRole(opts *bind.TransactOpts, account common.Address, role *big.Int) (*types.Transaction, error) {
	return _IERC423.contract.Transact(opts, "revokeRole", account, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x0912ed77.
//
// Solidity: function revokeRole(address account, uint256 role) returns(bool)
func (_IERC423 *IERC423Session) RevokeRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _IERC423.Contract.RevokeRole(&_IERC423.TransactOpts, account, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x0912ed77.
//
// Solidity: function revokeRole(address account, uint256 role) returns(bool)
func (_IERC423 *IERC423TransactorSession) RevokeRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _IERC423.Contract.RevokeRole(&_IERC423.TransactOpts, account, role)
}

// IERC423AgentDefinedIterator is returned from FilterAgentDefined and is used to iterate over the raw logs and unpacked data for AgentDefined events raised by the IERC423 contract.
type IERC423AgentDefinedIterator struct {
	Event *IERC423AgentDefined // Event containing the contract specifics and raw log

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
func (it *IERC423AgentDefinedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC423AgentDefined)
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
		it.Event = new(IERC423AgentDefined)
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
func (it *IERC423AgentDefinedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC423AgentDefinedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC423AgentDefined represents a AgentDefined event raised by the IERC423 contract.
type IERC423AgentDefined struct {
	Agent   common.Address
	Account common.Address
	Issuer  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentDefined is a free log retrieval operation binding the contract event 0xb61ddfe33d9ed54d8bdb62fe6fe16b84a72a4bae657baea7d2ae636fd1714b89.
//
// Solidity: event AgentDefined(address agent, address account, address issuer)
func (_IERC423 *IERC423Filterer) FilterAgentDefined(opts *bind.FilterOpts) (*IERC423AgentDefinedIterator, error) {

	logs, sub, err := _IERC423.contract.FilterLogs(opts, "AgentDefined")
	if err != nil {
		return nil, err
	}
	return &IERC423AgentDefinedIterator{contract: _IERC423.contract, event: "AgentDefined", logs: logs, sub: sub}, nil
}

// WatchAgentDefined is a free log subscription operation binding the contract event 0xb61ddfe33d9ed54d8bdb62fe6fe16b84a72a4bae657baea7d2ae636fd1714b89.
//
// Solidity: event AgentDefined(address agent, address account, address issuer)
func (_IERC423 *IERC423Filterer) WatchAgentDefined(opts *bind.WatchOpts, sink chan<- *IERC423AgentDefined) (event.Subscription, error) {

	logs, sub, err := _IERC423.contract.WatchLogs(opts, "AgentDefined")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC423AgentDefined)
				if err := _IERC423.contract.UnpackLog(event, "AgentDefined", log); err != nil {
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
func (_IERC423 *IERC423Filterer) ParseAgentDefined(log types.Log) (*IERC423AgentDefined, error) {
	event := new(IERC423AgentDefined)
	if err := _IERC423.contract.UnpackLog(event, "AgentDefined", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC423AgentRemovedIterator is returned from FilterAgentRemoved and is used to iterate over the raw logs and unpacked data for AgentRemoved events raised by the IERC423 contract.
type IERC423AgentRemovedIterator struct {
	Event *IERC423AgentRemoved // Event containing the contract specifics and raw log

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
func (it *IERC423AgentRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC423AgentRemoved)
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
		it.Event = new(IERC423AgentRemoved)
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
func (it *IERC423AgentRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC423AgentRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC423AgentRemoved represents a AgentRemoved event raised by the IERC423 contract.
type IERC423AgentRemoved struct {
	Agent  common.Address
	Issuer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAgentRemoved is a free log retrieval operation binding the contract event 0x4ce8b3fe8e0bf47a6fefbecdfd7c799755cede6061655521b10fc2e4b1fcc6b3.
//
// Solidity: event AgentRemoved(address agent, address issuer)
func (_IERC423 *IERC423Filterer) FilterAgentRemoved(opts *bind.FilterOpts) (*IERC423AgentRemovedIterator, error) {

	logs, sub, err := _IERC423.contract.FilterLogs(opts, "AgentRemoved")
	if err != nil {
		return nil, err
	}
	return &IERC423AgentRemovedIterator{contract: _IERC423.contract, event: "AgentRemoved", logs: logs, sub: sub}, nil
}

// WatchAgentRemoved is a free log subscription operation binding the contract event 0x4ce8b3fe8e0bf47a6fefbecdfd7c799755cede6061655521b10fc2e4b1fcc6b3.
//
// Solidity: event AgentRemoved(address agent, address issuer)
func (_IERC423 *IERC423Filterer) WatchAgentRemoved(opts *bind.WatchOpts, sink chan<- *IERC423AgentRemoved) (event.Subscription, error) {

	logs, sub, err := _IERC423.contract.WatchLogs(opts, "AgentRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC423AgentRemoved)
				if err := _IERC423.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
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
func (_IERC423 *IERC423Filterer) ParseAgentRemoved(log types.Log) (*IERC423AgentRemoved, error) {
	event := new(IERC423AgentRemoved)
	if err := _IERC423.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC423RoleDefinedIterator is returned from FilterRoleDefined and is used to iterate over the raw logs and unpacked data for RoleDefined events raised by the IERC423 contract.
type IERC423RoleDefinedIterator struct {
	Event *IERC423RoleDefined // Event containing the contract specifics and raw log

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
func (it *IERC423RoleDefinedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC423RoleDefined)
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
		it.Event = new(IERC423RoleDefined)
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
func (it *IERC423RoleDefinedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC423RoleDefinedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC423RoleDefined represents a RoleDefined event raised by the IERC423 contract.
type IERC423RoleDefined struct {
	Role   *big.Int
	Issuer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRoleDefined is a free log retrieval operation binding the contract event 0x2e82a3308270f66933807098cf62441247b1de86244636b03b0565b2c3dbc62b.
//
// Solidity: event RoleDefined(uint256 role, address issuer)
func (_IERC423 *IERC423Filterer) FilterRoleDefined(opts *bind.FilterOpts) (*IERC423RoleDefinedIterator, error) {

	logs, sub, err := _IERC423.contract.FilterLogs(opts, "RoleDefined")
	if err != nil {
		return nil, err
	}
	return &IERC423RoleDefinedIterator{contract: _IERC423.contract, event: "RoleDefined", logs: logs, sub: sub}, nil
}

// WatchRoleDefined is a free log subscription operation binding the contract event 0x2e82a3308270f66933807098cf62441247b1de86244636b03b0565b2c3dbc62b.
//
// Solidity: event RoleDefined(uint256 role, address issuer)
func (_IERC423 *IERC423Filterer) WatchRoleDefined(opts *bind.WatchOpts, sink chan<- *IERC423RoleDefined) (event.Subscription, error) {

	logs, sub, err := _IERC423.contract.WatchLogs(opts, "RoleDefined")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC423RoleDefined)
				if err := _IERC423.contract.UnpackLog(event, "RoleDefined", log); err != nil {
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
func (_IERC423 *IERC423Filterer) ParseRoleDefined(log types.Log) (*IERC423RoleDefined, error) {
	event := new(IERC423RoleDefined)
	if err := _IERC423.contract.UnpackLog(event, "RoleDefined", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC423RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the IERC423 contract.
type IERC423RoleGrantedIterator struct {
	Event *IERC423RoleGranted // Event containing the contract specifics and raw log

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
func (it *IERC423RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC423RoleGranted)
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
		it.Event = new(IERC423RoleGranted)
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
func (it *IERC423RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC423RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC423RoleGranted represents a RoleGranted event raised by the IERC423 contract.
type IERC423RoleGranted struct {
	Role    *big.Int
	Account common.Address
	Issuer  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x1ec1667fba5e43c5c76fff54e76d7a4a20a4fecf7b49724ad8d52a3f726e9dbd.
//
// Solidity: event RoleGranted(uint256 role, address account, address issuer)
func (_IERC423 *IERC423Filterer) FilterRoleGranted(opts *bind.FilterOpts) (*IERC423RoleGrantedIterator, error) {

	logs, sub, err := _IERC423.contract.FilterLogs(opts, "RoleGranted")
	if err != nil {
		return nil, err
	}
	return &IERC423RoleGrantedIterator{contract: _IERC423.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x1ec1667fba5e43c5c76fff54e76d7a4a20a4fecf7b49724ad8d52a3f726e9dbd.
//
// Solidity: event RoleGranted(uint256 role, address account, address issuer)
func (_IERC423 *IERC423Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *IERC423RoleGranted) (event.Subscription, error) {

	logs, sub, err := _IERC423.contract.WatchLogs(opts, "RoleGranted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC423RoleGranted)
				if err := _IERC423.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_IERC423 *IERC423Filterer) ParseRoleGranted(log types.Log) (*IERC423RoleGranted, error) {
	event := new(IERC423RoleGranted)
	if err := _IERC423.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC423RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the IERC423 contract.
type IERC423RoleRevokedIterator struct {
	Event *IERC423RoleRevoked // Event containing the contract specifics and raw log

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
func (it *IERC423RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC423RoleRevoked)
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
		it.Event = new(IERC423RoleRevoked)
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
func (it *IERC423RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC423RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC423RoleRevoked represents a RoleRevoked event raised by the IERC423 contract.
type IERC423RoleRevoked struct {
	Role    *big.Int
	Account common.Address
	Issuer  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xe0df21b65c73c27081b8f042a012b124085b41d78d27b7e3c4780f5650f5ebb8.
//
// Solidity: event RoleRevoked(uint256 role, address account, address issuer)
func (_IERC423 *IERC423Filterer) FilterRoleRevoked(opts *bind.FilterOpts) (*IERC423RoleRevokedIterator, error) {

	logs, sub, err := _IERC423.contract.FilterLogs(opts, "RoleRevoked")
	if err != nil {
		return nil, err
	}
	return &IERC423RoleRevokedIterator{contract: _IERC423.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xe0df21b65c73c27081b8f042a012b124085b41d78d27b7e3c4780f5650f5ebb8.
//
// Solidity: event RoleRevoked(uint256 role, address account, address issuer)
func (_IERC423 *IERC423Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *IERC423RoleRevoked) (event.Subscription, error) {

	logs, sub, err := _IERC423.contract.WatchLogs(opts, "RoleRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC423RoleRevoked)
				if err := _IERC423.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_IERC423 *IERC423Filterer) ParseRoleRevoked(log types.Log) (*IERC423RoleRevoked, error) {
	event := new(IERC423RoleRevoked)
	if err := _IERC423.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
