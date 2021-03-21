// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package faaslevel

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FaaSLevelABI is the input ABI used to generate the binding from.
const FaaSLevelABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_core\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_mem\",\"type\":\"uint256\"}],\"name\":\"addFaaSLevelEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_core\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mem\",\"type\":\"uint256\"}],\"name\":\"addFaaSLevel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_faasLevelID\",\"type\":\"uint256\"}],\"name\":\"getFaaSLevel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numFaaSLevel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FaaSLevelBin is the compiled bytecode used for deploying new contracts.
var FaaSLevelBin = "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600060018190555061006c60016102006100bd60201b60201c565b61007f60016104006100bd60201b60201c565b61009260016108006100bd60201b60201c565b6100a560026104006100bd60201b60201c565b6100b860046108006100bd60201b60201c565b61031b565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461014b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161014290610216565b60405180910390fd5b6000600160008154809291906101609061027a565b919050559050604051806040016040528084815260200183815250600260008381526020019081526020016000206000820151816000015560208201518160010155905050807f901fa3fc598633460c6bb1e971f1a36b52ae0240ec8e04d15251e8f092cd7ec184846040516101d7929190610236565b60405180910390a2505050565b60006101f160198361025f565b91506101fc826102f2565b602082019050919050565b61021081610270565b82525050565b6000602082019050818103600083015261022f816101e4565b9050919050565b600060408201905061024b6000830185610207565b6102586020830184610207565b9392505050565b600082825260208201905092915050565b6000819050919050565b600061028582610270565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156102b8576102b76102c3565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4f6e6c79206f776e65722063616e2063616c6c20746869732e00000000000000600082015250565b6105808061032a6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80638da5cb5b146100515780639a5992011461006f578063c0b717c01461008d578063c86fdf7d146100a9575b600080fd5b6100596100da565b604051610066919061037d565b60405180910390f35b6100776100fe565b60405161008491906103d8565b60405180910390f35b6100a760048036038101906100a291906102dd565b610104565b005b6100c360048036038101906100be91906102b4565b61022b565b6040516100d19291906103f3565b60405180910390f35b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610192576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161018990610398565b60405180910390fd5b6000600160008154809291906101a790610469565b919050559050604051806040016040528084815260200183815250600260008381526020019081526020016000206000820151816000015560208201518160010155905050807f901fa3fc598633460c6bb1e971f1a36b52ae0240ec8e04d15251e8f092cd7ec1848460405161021e9291906103f3565b60405180910390a2505050565b600080826001548110610273576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026a906103b8565b60405180910390fd5b600060026000868152602001908152602001600020905080600001548160010154935093505050915091565b6000813590506102ae81610533565b92915050565b6000602082840312156102c657600080fd5b60006102d48482850161029f565b91505092915050565b600080604083850312156102f057600080fd5b60006102fe8582860161029f565b925050602061030f8582860161029f565b9150509250929050565b6103228161042d565b82525050565b600061033560198361041c565b9150610340826104e1565b602082019050919050565b6000610358601f8361041c565b91506103638261050a565b602082019050919050565b6103778161045f565b82525050565b60006020820190506103926000830184610319565b92915050565b600060208201905081810360008301526103b181610328565b9050919050565b600060208201905081810360008301526103d18161034b565b9050919050565b60006020820190506103ed600083018461036e565b92915050565b6000604082019050610408600083018561036e565b610415602083018461036e565b9392505050565b600082825260208201905092915050565b60006104388261043f565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006104748261045f565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156104a7576104a66104b2565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4f6e6c79206f776e65722063616e2063616c6c20746869732e00000000000000600082015250565b7f466161534c6576656c3a20696e76616c696420666161736c6576656c20494400600082015250565b61053c8161045f565b811461054757600080fd5b5056fea26469706673582212207502392c15f26c43c7120cb47743882540afa45907259c76d0035e749fa4ad8764736f6c63430008010033"

// DeployFaaSLevel deploys a new Ethereum contract, binding an instance of FaaSLevel to it.
func DeployFaaSLevel(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FaaSLevel, error) {
	parsed, err := abi.JSON(strings.NewReader(FaaSLevelABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FaaSLevelBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FaaSLevel{FaaSLevelCaller: FaaSLevelCaller{contract: contract}, FaaSLevelTransactor: FaaSLevelTransactor{contract: contract}, FaaSLevelFilterer: FaaSLevelFilterer{contract: contract}}, nil
}

// FaaSLevel is an auto generated Go binding around an Ethereum contract.
type FaaSLevel struct {
	FaaSLevelCaller     // Read-only binding to the contract
	FaaSLevelTransactor // Write-only binding to the contract
	FaaSLevelFilterer   // Log filterer for contract events
}

// FaaSLevelCaller is an auto generated read-only Go binding around an Ethereum contract.
type FaaSLevelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaaSLevelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FaaSLevelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaaSLevelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FaaSLevelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaaSLevelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FaaSLevelSession struct {
	Contract     *FaaSLevel        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FaaSLevelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FaaSLevelCallerSession struct {
	Contract *FaaSLevelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// FaaSLevelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FaaSLevelTransactorSession struct {
	Contract     *FaaSLevelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FaaSLevelRaw is an auto generated low-level Go binding around an Ethereum contract.
type FaaSLevelRaw struct {
	Contract *FaaSLevel // Generic contract binding to access the raw methods on
}

// FaaSLevelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FaaSLevelCallerRaw struct {
	Contract *FaaSLevelCaller // Generic read-only contract binding to access the raw methods on
}

// FaaSLevelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
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
func (_FaaSLevel *FaaSLevelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FaaSLevel.Contract.FaaSLevelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FaaSLevel *FaaSLevelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FaaSLevel.Contract.FaaSLevelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FaaSLevel *FaaSLevelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FaaSLevel.Contract.FaaSLevelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FaaSLevel *FaaSLevelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FaaSLevel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FaaSLevel *FaaSLevelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FaaSLevel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FaaSLevel *FaaSLevelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FaaSLevel.Contract.contract.Transact(opts, method, params...)
}

// GetFaaSLevel is a free data retrieval call binding the contract method 0xc86fdf7d.
//
// Solidity: function getFaaSLevel(uint256 _faasLevelID) view returns(uint256, uint256)
func (_FaaSLevel *FaaSLevelCaller) GetFaaSLevel(opts *bind.CallOpts, _faasLevelID *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _FaaSLevel.contract.Call(opts, &out, "getFaaSLevel", _faasLevelID)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetFaaSLevel is a free data retrieval call binding the contract method 0xc86fdf7d.
//
// Solidity: function getFaaSLevel(uint256 _faasLevelID) view returns(uint256, uint256)
func (_FaaSLevel *FaaSLevelSession) GetFaaSLevel(_faasLevelID *big.Int) (*big.Int, *big.Int, error) {
	return _FaaSLevel.Contract.GetFaaSLevel(&_FaaSLevel.CallOpts, _faasLevelID)
}

// GetFaaSLevel is a free data retrieval call binding the contract method 0xc86fdf7d.
//
// Solidity: function getFaaSLevel(uint256 _faasLevelID) view returns(uint256, uint256)
func (_FaaSLevel *FaaSLevelCallerSession) GetFaaSLevel(_faasLevelID *big.Int) (*big.Int, *big.Int, error) {
	return _FaaSLevel.Contract.GetFaaSLevel(&_FaaSLevel.CallOpts, _faasLevelID)
}

// NumFaaSLevel is a free data retrieval call binding the contract method 0x9a599201.
//
// Solidity: function numFaaSLevel() view returns(uint256)
func (_FaaSLevel *FaaSLevelCaller) NumFaaSLevel(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FaaSLevel.contract.Call(opts, &out, "numFaaSLevel")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumFaaSLevel is a free data retrieval call binding the contract method 0x9a599201.
//
// Solidity: function numFaaSLevel() view returns(uint256)
func (_FaaSLevel *FaaSLevelSession) NumFaaSLevel() (*big.Int, error) {
	return _FaaSLevel.Contract.NumFaaSLevel(&_FaaSLevel.CallOpts)
}

// NumFaaSLevel is a free data retrieval call binding the contract method 0x9a599201.
//
// Solidity: function numFaaSLevel() view returns(uint256)
func (_FaaSLevel *FaaSLevelCallerSession) NumFaaSLevel() (*big.Int, error) {
	return _FaaSLevel.Contract.NumFaaSLevel(&_FaaSLevel.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FaaSLevel *FaaSLevelCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FaaSLevel.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FaaSLevel *FaaSLevelSession) Owner() (common.Address, error) {
	return _FaaSLevel.Contract.Owner(&_FaaSLevel.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FaaSLevel *FaaSLevelCallerSession) Owner() (common.Address, error) {
	return _FaaSLevel.Contract.Owner(&_FaaSLevel.CallOpts)
}

// AddFaaSLevel is a paid mutator transaction binding the contract method 0xc0b717c0.
//
// Solidity: function addFaaSLevel(uint256 _core, uint256 _mem) returns()
func (_FaaSLevel *FaaSLevelTransactor) AddFaaSLevel(opts *bind.TransactOpts, _core *big.Int, _mem *big.Int) (*types.Transaction, error) {
	return _FaaSLevel.contract.Transact(opts, "addFaaSLevel", _core, _mem)
}

// AddFaaSLevel is a paid mutator transaction binding the contract method 0xc0b717c0.
//
// Solidity: function addFaaSLevel(uint256 _core, uint256 _mem) returns()
func (_FaaSLevel *FaaSLevelSession) AddFaaSLevel(_core *big.Int, _mem *big.Int) (*types.Transaction, error) {
	return _FaaSLevel.Contract.AddFaaSLevel(&_FaaSLevel.TransactOpts, _core, _mem)
}

// AddFaaSLevel is a paid mutator transaction binding the contract method 0xc0b717c0.
//
// Solidity: function addFaaSLevel(uint256 _core, uint256 _mem) returns()
func (_FaaSLevel *FaaSLevelTransactorSession) AddFaaSLevel(_core *big.Int, _mem *big.Int) (*types.Transaction, error) {
	return _FaaSLevel.Contract.AddFaaSLevel(&_FaaSLevel.TransactOpts, _core, _mem)
}

// FaaSLevelAddFaaSLevelEventIterator is returned from FilterAddFaaSLevelEvent and is used to iterate over the raw logs and unpacked data for AddFaaSLevelEvent events raised by the FaaSLevel contract.
type FaaSLevelAddFaaSLevelEventIterator struct {
	Event *FaaSLevelAddFaaSLevelEvent // Event containing the contract specifics and raw log

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
func (it *FaaSLevelAddFaaSLevelEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FaaSLevelAddFaaSLevelEvent)
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
		it.Event = new(FaaSLevelAddFaaSLevelEvent)
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
func (it *FaaSLevelAddFaaSLevelEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FaaSLevelAddFaaSLevelEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FaaSLevelAddFaaSLevelEvent represents a AddFaaSLevelEvent event raised by the FaaSLevel contract.
type FaaSLevelAddFaaSLevelEvent struct {
	Index *big.Int
	Core  *big.Int
	Mem   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddFaaSLevelEvent is a free log retrieval operation binding the contract event 0x901fa3fc598633460c6bb1e971f1a36b52ae0240ec8e04d15251e8f092cd7ec1.
//
// Solidity: event addFaaSLevelEvent(uint256 indexed _index, uint256 _core, uint256 _mem)
func (_FaaSLevel *FaaSLevelFilterer) FilterAddFaaSLevelEvent(opts *bind.FilterOpts, _index []*big.Int) (*FaaSLevelAddFaaSLevelEventIterator, error) {

	var _indexRule []interface{}
	for _, _indexItem := range _index {
		_indexRule = append(_indexRule, _indexItem)
	}

	logs, sub, err := _FaaSLevel.contract.FilterLogs(opts, "addFaaSLevelEvent", _indexRule)
	if err != nil {
		return nil, err
	}
	return &FaaSLevelAddFaaSLevelEventIterator{contract: _FaaSLevel.contract, event: "addFaaSLevelEvent", logs: logs, sub: sub}, nil
}

// WatchAddFaaSLevelEvent is a free log subscription operation binding the contract event 0x901fa3fc598633460c6bb1e971f1a36b52ae0240ec8e04d15251e8f092cd7ec1.
//
// Solidity: event addFaaSLevelEvent(uint256 indexed _index, uint256 _core, uint256 _mem)
func (_FaaSLevel *FaaSLevelFilterer) WatchAddFaaSLevelEvent(opts *bind.WatchOpts, sink chan<- *FaaSLevelAddFaaSLevelEvent, _index []*big.Int) (event.Subscription, error) {

	var _indexRule []interface{}
	for _, _indexItem := range _index {
		_indexRule = append(_indexRule, _indexItem)
	}

	logs, sub, err := _FaaSLevel.contract.WatchLogs(opts, "addFaaSLevelEvent", _indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FaaSLevelAddFaaSLevelEvent)
				if err := _FaaSLevel.contract.UnpackLog(event, "addFaaSLevelEvent", log); err != nil {
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

// ParseAddFaaSLevelEvent is a log parse operation binding the contract event 0x901fa3fc598633460c6bb1e971f1a36b52ae0240ec8e04d15251e8f092cd7ec1.
//
// Solidity: event addFaaSLevelEvent(uint256 indexed _index, uint256 _core, uint256 _mem)
func (_FaaSLevel *FaaSLevelFilterer) ParseAddFaaSLevelEvent(log types.Log) (*FaaSLevelAddFaaSLevelEvent, error) {
	event := new(FaaSLevelAddFaaSLevelEvent)
	if err := _FaaSLevel.contract.UnpackLog(event, "addFaaSLevelEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
