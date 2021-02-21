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
const FaaSLevelABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_core\",\"type\":\"uint256\"},{\"name\":\"_mem\",\"type\":\"uint256\"}],\"name\":\"newFaaSLevel\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFaaSLevelNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_levelID\",\"type\":\"uint256\"}],\"name\":\"getFaaSLevel\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"core\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"mem\",\"type\":\"uint256\"}],\"name\":\"newFaaSLevelEvent\",\"type\":\"event\"}]"

// FaaSLevelBin is the compiled bytecode used for deploying new contracts.
var FaaSLevelBin = "0x608060405234801561001057600080fd5b506040516020806105fa8339810180604052810190808051906020019092919050505080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600080819055506100976001610200610112640100000000026401000000009004565b506100b46001610400610112640100000000026401000000009004565b506100d16001610800610112640100000000026401000000009004565b506100ee6002610400610112640100000000026401000000009004565b5061010b6004610800610112640100000000026401000000009004565b505061027c565b600080600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156101da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f4f6e6c79206f776e65722063616e2063616c6c20746869732e0000000000000081525060200191505060405180910390fd5b6000808154809291906001019190505590506040805190810160405280858152602001848152506001600083815260200190815260200160002060008201518160000155602082015181600101559050507f0f7fd73bc5b00f997cac82624b411f100d076b6990053a1e4956f92f69fbf2f781858560405180848152602001838152602001828152602001935050505060405180910390a18091505092915050565b61036f8061028b6000396000f300608060405260043610610062576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806304d15fb4146100675780638da5cb5b146100b2578063c780747814610109578063c86fdf7d14610134575b600080fd5b34801561007357600080fd5b5061009c600480360381019080803590602001909291908035906020019092919050505061017c565b6040518082815260200191505060405180910390f35b3480156100be57600080fd5b506100c76102e6565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561011557600080fd5b5061011e61030c565b6040518082815260200191505060405180910390f35b34801561014057600080fd5b5061015f60048036038101908080359060200190929190505050610315565b604051808381526020018281526020019250505060405180910390f35b600080600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610244576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f4f6e6c79206f776e65722063616e2063616c6c20746869732e0000000000000081525060200191505060405180910390fd5b6000808154809291906001019190505590506040805190810160405280858152602001848152506001600083815260200190815260200160002060008201518160000155602082015181600101559050507f0f7fd73bc5b00f997cac82624b411f100d076b6990053a1e4956f92f69fbf2f781858560405180848152602001838152602001828152602001935050505060405180910390a18091505092915050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054905090565b60008060006001600085815260200190815260200160002090508060000154816001015492509250509150915600a165627a7a723058201968e3ab2bf52687633262bbd8c39229d4855de4a849f6e61dcee012210852980029"

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
// Solidity: function getFaaSLevel(uint256 _levelID) constant returns(uint256, uint256)
func (_FaaSLevel *FaaSLevelCaller) GetFaaSLevel(opts *bind.CallOpts, _levelID *big.Int) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _FaaSLevel.contract.Call(opts, out, "getFaaSLevel", _levelID)
	return *ret0, *ret1, err
}

// GetFaaSLevel is a free data retrieval call binding the contract method 0xc86fdf7d.
//
// Solidity: function getFaaSLevel(uint256 _levelID) constant returns(uint256, uint256)
func (_FaaSLevel *FaaSLevelSession) GetFaaSLevel(_levelID *big.Int) (*big.Int, *big.Int, error) {
	return _FaaSLevel.Contract.GetFaaSLevel(&_FaaSLevel.CallOpts, _levelID)
}

// GetFaaSLevel is a free data retrieval call binding the contract method 0xc86fdf7d.
//
// Solidity: function getFaaSLevel(uint256 _levelID) constant returns(uint256, uint256)
func (_FaaSLevel *FaaSLevelCallerSession) GetFaaSLevel(_levelID *big.Int) (*big.Int, *big.Int, error) {
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
