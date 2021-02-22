// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package faaslevel

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/event"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FaaSLevelABI is the input ABI used to generate the binding from.
const FaaSLevelABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"core\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mem\",\"type\":\"uint256\"}],\"name\":\"newFaaSLevelEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_levelID\",\"type\":\"uint256\"}],\"name\":\"getFaaSLevel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFaaSLevelNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_core\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mem\",\"type\":\"uint256\"}],\"name\":\"newFaaSLevel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FaaSLevelBin is the compiled bytecode used for deploying new contracts.
var FaaSLevelBin = "0x608060405234801561001057600080fd5b506040516106293803806106298339818101604052602081101561003357600080fd5b810190808051906020019092919050505080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000808190555061009f60016102006100f660201b60201c565b506100b360016104006100f660201b60201c565b506100c760016108006100f660201b60201c565b506100db60026104006100f660201b60201c565b506100ef60046108006100f660201b60201c565b505061025f565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f4f6e6c79206f776e65722063616e2063616c6c20746869732e0000000000000081525060200191505060405180910390fd5b60008060008154809291906001019190505590506040518060400160405280858152602001848152506001600083815260200190815260200160002060008201518160000155602082015181600101559050507f0f7fd73bc5b00f997cac82624b411f100d076b6990053a1e4956f92f69fbf2f781858560405180848152602001838152602001828152602001935050505060405180910390a18091505092915050565b6103bb8061026e6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806304d15fb4146100515780638da5cb5b1461009d578063c7807478146100e7578063c86fdf7d14610105575b600080fd5b6100876004803603604081101561006757600080fd5b810190808035906020019092919080359060200190929190505050610159565b6040518082815260200191505060405180910390f35b6100a56102c2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6100ef6102e8565b6040518082815260200191505060405180910390f35b6101316004803603602081101561011b57600080fd5b81019080803590602001909291905050506102f1565b6040518084151515158152602001838152602001828152602001935050505060405180910390f35b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461021e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f4f6e6c79206f776e65722063616e2063616c6c20746869732e0000000000000081525060200191505060405180910390fd5b60008060008154809291906001019190505590506040518060400160405280858152602001848152506001600083815260200190815260200160002060008201518160000155602082015181600101559050507f0f7fd73bc5b00f997cac82624b411f100d076b6990053a1e4956f92f69fbf2f781858560405180848152602001838152602001828152602001935050505060405180910390a18091505092915050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054905090565b600080600080548410610314576000806000819150809050925092509250610364565b61031c61036b565b60016000868152602001908152602001600020604051806040016040529081600082015481526020016001820154815250509050600181600001518260200151935093509350505b9193909250565b60405180604001604052806000815260200160008152509056fea26469706673582212204484facd1493063c11cf9c9b78263eaaf7a59de8822e68f0f6fb21e54da3114364736f6c634300060a0033"

// DeployFaaSLevel deploys a new contract, binding an instance of FaaSLevel to it.
func DeployFaaSLevel(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *FaaSLevel, error) {
	parsed, err := abi.JSON(strings.NewReader(FaaSLevelABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FaaSLevelBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FaaSLevel{FaaSLevelCaller: FaaSLevelCaller{contract: contract}, FaaSLevelTransactor: FaaSLevelTransactor{contract: contract}, FaaSLevelFilterer: FaaSLevelFilterer{contract: contract}}, nil
}

func AsyncDeployFaaSLevel(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend, _owner common.Address) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(FaaSLevelABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(FaaSLevelBin), backend, _owner)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// FaaSLevel is an auto generated Go binding around a Solidity contract.
type FaaSLevel struct {
	FaaSLevelCaller     // Read-only binding to the contract
	FaaSLevelTransactor // Write-only binding to the contract
	FaaSLevelFilterer   // Log filterer for contract events
}

// FaaSLevelCaller is an auto generated read-only Go binding around a Solidity contract.
type FaaSLevelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaaSLevelTransactor is an auto generated write-only Go binding around a Solidity contract.
type FaaSLevelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaaSLevelFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type FaaSLevelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaaSLevelSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type FaaSLevelSession struct {
	Contract     *FaaSLevel        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FaaSLevelCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type FaaSLevelCallerSession struct {
	Contract *FaaSLevelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// FaaSLevelTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type FaaSLevelTransactorSession struct {
	Contract     *FaaSLevelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FaaSLevelRaw is an auto generated low-level Go binding around a Solidity contract.
type FaaSLevelRaw struct {
	Contract *FaaSLevel // Generic contract binding to access the raw methods on
}

// FaaSLevelCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type FaaSLevelCallerRaw struct {
	Contract *FaaSLevelCaller // Generic read-only contract binding to access the raw methods on
}

// FaaSLevelTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type FaaSLevelTransactorRaw struct {
	Contract *FaaSLevelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFaaSLevel creates a new instance of FaaSLevel, bound to a specific deployed contract.
func NewFaaSLevel(address common.Address, backend bind.ContractBackend) (*FaaSLevel, error) {
	contract, err := bindFaaSLevel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FaaSLevel{FaaSLevelCaller: FaaSLevelCaller{contract: contract}, FaaSLevelTransactor: FaaSLevelTransactor{contract: contract}, FaaSLevelFilterer: FaaSLevelFilterer{contract: contract}}, nil
}

// NewFaaSLevelCaller creates a new read-only instance of FaaSLevel, bound to a specific deployed contract.
func NewFaaSLevelCaller(address common.Address, caller bind.ContractCaller) (*FaaSLevelCaller, error) {
	contract, err := bindFaaSLevel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FaaSLevelCaller{contract: contract}, nil
}

// NewFaaSLevelTransactor creates a new write-only instance of FaaSLevel, bound to a specific deployed contract.
func NewFaaSLevelTransactor(address common.Address, transactor bind.ContractTransactor) (*FaaSLevelTransactor, error) {
	contract, err := bindFaaSLevel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FaaSLevelTransactor{contract: contract}, nil
}

// NewFaaSLevelFilterer creates a new log filterer instance of FaaSLevel, bound to a specific deployed contract.
func NewFaaSLevelFilterer(address common.Address, filterer bind.ContractFilterer) (*FaaSLevelFilterer, error) {
	contract, err := bindFaaSLevel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FaaSLevelFilterer{contract: contract}, nil
}

// bindFaaSLevel binds a generic wrapper to an already deployed contract.
func bindFaaSLevel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FaaSLevelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FaaSLevel *FaaSLevelRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FaaSLevel.Contract.FaaSLevelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FaaSLevel *FaaSLevelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _FaaSLevel.Contract.FaaSLevelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FaaSLevel *FaaSLevelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _FaaSLevel.Contract.FaaSLevelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FaaSLevel *FaaSLevelCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FaaSLevel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FaaSLevel *FaaSLevelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _FaaSLevel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FaaSLevel *FaaSLevelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _FaaSLevel.Contract.contract.Transact(opts, method, params...)
}

// GetFaaSLevel is a free data retrieval call binding the contract method 0xc86fdf7d.
//
// Solidity: function getFaaSLevel(uint256 _levelID) constant returns(bool, uint256, uint256)
func (_FaaSLevel *FaaSLevelCaller) GetFaaSLevel(opts *bind.CallOpts, _levelID *big.Int) (bool, *big.Int, *big.Int, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _FaaSLevel.contract.Call(opts, out, "getFaaSLevel", _levelID)
	return *ret0, *ret1, *ret2, err
}

// GetFaaSLevel is a free data retrieval call binding the contract method 0xc86fdf7d.
//
// Solidity: function getFaaSLevel(uint256 _levelID) constant returns(bool, uint256, uint256)
func (_FaaSLevel *FaaSLevelSession) GetFaaSLevel(_levelID *big.Int) (bool, *big.Int, *big.Int, error) {
	return _FaaSLevel.Contract.GetFaaSLevel(&_FaaSLevel.CallOpts, _levelID)
}

// GetFaaSLevel is a free data retrieval call binding the contract method 0xc86fdf7d.
//
// Solidity: function getFaaSLevel(uint256 _levelID) constant returns(bool, uint256, uint256)
func (_FaaSLevel *FaaSLevelCallerSession) GetFaaSLevel(_levelID *big.Int) (bool, *big.Int, *big.Int, error) {
	return _FaaSLevel.Contract.GetFaaSLevel(&_FaaSLevel.CallOpts, _levelID)
}

// GetFaaSLevelNumber is a free data retrieval call binding the contract method 0xc7807478.
//
// Solidity: function getFaaSLevelNumber() constant returns(uint256)
func (_FaaSLevel *FaaSLevelCaller) GetFaaSLevelNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FaaSLevel.contract.Call(opts, out, "getFaaSLevelNumber")
	return *ret0, err
}

// GetFaaSLevelNumber is a free data retrieval call binding the contract method 0xc7807478.
//
// Solidity: function getFaaSLevelNumber() constant returns(uint256)
func (_FaaSLevel *FaaSLevelSession) GetFaaSLevelNumber() (*big.Int, error) {
	return _FaaSLevel.Contract.GetFaaSLevelNumber(&_FaaSLevel.CallOpts)
}

// GetFaaSLevelNumber is a free data retrieval call binding the contract method 0xc7807478.
//
// Solidity: function getFaaSLevelNumber() constant returns(uint256)
func (_FaaSLevel *FaaSLevelCallerSession) GetFaaSLevelNumber() (*big.Int, error) {
	return _FaaSLevel.Contract.GetFaaSLevelNumber(&_FaaSLevel.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FaaSLevel *FaaSLevelCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FaaSLevel.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FaaSLevel *FaaSLevelSession) Owner() (common.Address, error) {
	return _FaaSLevel.Contract.Owner(&_FaaSLevel.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FaaSLevel *FaaSLevelCallerSession) Owner() (common.Address, error) {
	return _FaaSLevel.Contract.Owner(&_FaaSLevel.CallOpts)
}

// NewFaaSLevel is a paid mutator transaction binding the contract method 0x04d15fb4.
//
// Solidity: function newFaaSLevel(uint256 _core, uint256 _mem) returns(uint256)
func (_FaaSLevel *FaaSLevelTransactor) NewFaaSLevel(opts *bind.TransactOpts, _core *big.Int, _mem *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSLevel.contract.Transact(opts, "newFaaSLevel", _core, _mem)
}

func (_FaaSLevel *FaaSLevelTransactor) AsyncNewFaaSLevel(handler func(*types.Receipt, error), opts *bind.TransactOpts, _core *big.Int, _mem *big.Int) (*types.Transaction, error) {
	return _FaaSLevel.contract.AsyncTransact(opts, handler, "newFaaSLevel", _core, _mem)
}

// NewFaaSLevel is a paid mutator transaction binding the contract method 0x04d15fb4.
//
// Solidity: function newFaaSLevel(uint256 _core, uint256 _mem) returns(uint256)
func (_FaaSLevel *FaaSLevelSession) NewFaaSLevel(_core *big.Int, _mem *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSLevel.Contract.NewFaaSLevel(&_FaaSLevel.TransactOpts, _core, _mem)
}

func (_FaaSLevel *FaaSLevelSession) AsyncNewFaaSLevel(handler func(*types.Receipt, error), _core *big.Int, _mem *big.Int) (*types.Transaction, error) {
	return _FaaSLevel.Contract.AsyncNewFaaSLevel(handler, &_FaaSLevel.TransactOpts, _core, _mem)
}

// NewFaaSLevel is a paid mutator transaction binding the contract method 0x04d15fb4.
//
// Solidity: function newFaaSLevel(uint256 _core, uint256 _mem) returns(uint256)
func (_FaaSLevel *FaaSLevelTransactorSession) NewFaaSLevel(_core *big.Int, _mem *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSLevel.Contract.NewFaaSLevel(&_FaaSLevel.TransactOpts, _core, _mem)
}

func (_FaaSLevel *FaaSLevelTransactorSession) AsyncNewFaaSLevel(handler func(*types.Receipt, error), _core *big.Int, _mem *big.Int) (*types.Transaction, error) {
	return _FaaSLevel.Contract.AsyncNewFaaSLevel(handler, &_FaaSLevel.TransactOpts, _core, _mem)
}

// FaaSLevelNewFaaSLevelEventIterator is returned from FilterNewFaaSLevelEvent and is used to iterate over the raw logs and unpacked data for NewFaaSLevelEvent events raised by the FaaSLevel contract.
type FaaSLevelNewFaaSLevelEventIterator struct {
	Event *FaaSLevelNewFaaSLevelEvent // Event containing the contract specifics and raw log

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
func (it *FaaSLevelNewFaaSLevelEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FaaSLevelNewFaaSLevelEvent)
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
		it.Event = new(FaaSLevelNewFaaSLevelEvent)
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
func (it *FaaSLevelNewFaaSLevelEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FaaSLevelNewFaaSLevelEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FaaSLevelNewFaaSLevelEvent represents a NewFaaSLevelEvent event raised by the FaaSLevel contract.
type FaaSLevelNewFaaSLevelEvent struct {
	Index *big.Int
	Core  *big.Int
	Mem   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewFaaSLevelEvent is a free log retrieval operation binding the contract event 0x0f7fd73bc5b00f997cac82624b411f100d076b6990053a1e4956f92f69fbf2f7.
//
// Solidity: event newFaaSLevelEvent(uint256 index, uint256 core, uint256 mem)
func (_FaaSLevel *FaaSLevelFilterer) FilterNewFaaSLevelEvent(opts *bind.FilterOpts) (*FaaSLevelNewFaaSLevelEventIterator, error) {

	logs, sub, err := _FaaSLevel.contract.FilterLogs(opts, "newFaaSLevelEvent")
	if err != nil {
		return nil, err
	}
	return &FaaSLevelNewFaaSLevelEventIterator{contract: _FaaSLevel.contract, event: "newFaaSLevelEvent", logs: logs, sub: sub}, nil
}

// WatchNewFaaSLevelEvent is a free log subscription operation binding the contract event 0x0f7fd73bc5b00f997cac82624b411f100d076b6990053a1e4956f92f69fbf2f7.
//
// Solidity: event newFaaSLevelEvent(uint256 index, uint256 core, uint256 mem)
func (_FaaSLevel *FaaSLevelFilterer) WatchNewFaaSLevelEvent(opts *bind.WatchOpts, sink chan<- *FaaSLevelNewFaaSLevelEvent) (event.Subscription, error) {

	logs, sub, err := _FaaSLevel.contract.WatchLogs(opts, "newFaaSLevelEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FaaSLevelNewFaaSLevelEvent)
				if err := _FaaSLevel.contract.UnpackLog(event, "newFaaSLevelEvent", log); err != nil {
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

// ParseNewFaaSLevelEvent is a log parse operation binding the contract event 0x0f7fd73bc5b00f997cac82624b411f100d076b6990053a1e4956f92f69fbf2f7.
//
// Solidity: event newFaaSLevelEvent(uint256 index, uint256 core, uint256 mem)
func (_FaaSLevel *FaaSLevelFilterer) ParseNewFaaSLevelEvent(log types.Log) (*FaaSLevelNewFaaSLevelEvent, error) {
	event := new(FaaSLevelNewFaaSLevelEvent)
	if err := _FaaSLevel.contract.UnpackLog(event, "newFaaSLevelEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}
