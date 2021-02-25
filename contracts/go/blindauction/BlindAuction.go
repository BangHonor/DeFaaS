// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blindauction

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

// BlindAuctionABI is the input ABI used to generate the binding from.
const BlindAuctionABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"auctionEnd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumBlindAuction.States\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BlindAuctionBin is the compiled bytecode used for deploying new contracts.
var BlindAuctionBin = "0x608060405234801561001057600080fd5b5060008060006101000a81548160ff0219169083600381111561002f57fe5b0217905550610627806100436000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80631998aeef1461005c5780632a24f46c14610066578063a475b5dd14610070578063c19d93fb1461007a578063de292789146100a6575b600080fd5b6100646100b0565b005b61006e6101c3565b005b6100786102de565b005b6100826103f1565b6040518082600381111561009257fe5b60ff16815260200191505060405180910390f35b6100ae610403565b005b600060038111156100bd57fe5b6000809054906101000a900460ff1660038111156100d757fe5b1480156100e957506002546001540142115b156100f7576100f6610484565b5b6001600381111561010457fe5b6000809054906101000a900460ff16600381111561011e57fe5b1480156101345750600354600254600154010142115b1561014257610141610484565b5b600080600381111561015057fe5b6000809054906101000a900460ff16600381111561016a57fe5b146101c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260288152602001806105ca6028913960400191505060405180910390fd5b50565b600060038111156101d057fe5b6000809054906101000a900460ff1660038111156101ea57fe5b1480156101fc57506002546001540142115b1561020a57610209610484565b5b6001600381111561021757fe5b6000809054906101000a900460ff16600381111561023157fe5b1480156102475750600354600254600154010142115b1561025557610254610484565b5b600280600381111561026357fe5b6000809054906101000a900460ff16600381111561027d57fe5b146102d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260288152602001806105ca6028913960400191505060405180910390fd5b6102db610484565b50565b600060038111156102eb57fe5b6000809054906101000a900460ff16600381111561030557fe5b14801561031757506002546001540142115b1561032557610324610484565b5b6001600381111561033257fe5b6000809054906101000a900460ff16600381111561034c57fe5b1480156103625750600354600254600154010142115b156103705761036f610484565b5b600180600381111561037e57fe5b6000809054906101000a900460ff16600381111561039857fe5b146103ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260288152602001806105ca6028913960400191505060405180910390fd5b50565b6000809054906101000a900460ff1681565b600380600381111561041157fe5b6000809054906101000a900460ff16600381111561042b57fe5b14610481576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260288152602001806105ca6028913960400191505060405180910390fd5b50565b6000600381111561049157fe5b6000809054906101000a900460ff1660038111156104ab57fe5b14156104d55760016000806101000a81548160ff021916908360038111156104cf57fe5b02179055505b600160038111156104e257fe5b6000809054906101000a900460ff1660038111156104fc57fe5b14156105265760026000806101000a81548160ff0219169083600381111561052057fe5b02179055505b6002600381111561053357fe5b6000809054906101000a900460ff16600381111561054d57fe5b14156105775760036000806101000a81548160ff0219169083600381111561057157fe5b02179055505b60038081111561058357fe5b6000809054906101000a900460ff16600381111561059d57fe5b14156105c75760036000806101000a81548160ff021916908360038111156105c157fe5b02179055505b56fe46756e6374696f6e2063616e6e6f742062652063616c6c656420617420746869732073746174652ea2646970667358221220b1240eef1272a3715f68ed980fa30c9fddb0187ebd1f2cb479c734980a6383b864736f6c634300060a0033"

// DeployBlindAuction deploys a new contract, binding an instance of BlindAuction to it.
func DeployBlindAuction(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BlindAuction, error) {
	parsed, err := abi.JSON(strings.NewReader(BlindAuctionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BlindAuctionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BlindAuction{BlindAuctionCaller: BlindAuctionCaller{contract: contract}, BlindAuctionTransactor: BlindAuctionTransactor{contract: contract}, BlindAuctionFilterer: BlindAuctionFilterer{contract: contract}}, nil
}

func AsyncDeployBlindAuction(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(BlindAuctionABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(BlindAuctionBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// BlindAuction is an auto generated Go binding around a Solidity contract.
type BlindAuction struct {
	BlindAuctionCaller     // Read-only binding to the contract
	BlindAuctionTransactor // Write-only binding to the contract
	BlindAuctionFilterer   // Log filterer for contract events
}

// BlindAuctionCaller is an auto generated read-only Go binding around a Solidity contract.
type BlindAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlindAuctionTransactor is an auto generated write-only Go binding around a Solidity contract.
type BlindAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlindAuctionFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type BlindAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlindAuctionSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type BlindAuctionSession struct {
	Contract     *BlindAuction     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlindAuctionCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type BlindAuctionCallerSession struct {
	Contract *BlindAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BlindAuctionTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type BlindAuctionTransactorSession struct {
	Contract     *BlindAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BlindAuctionRaw is an auto generated low-level Go binding around a Solidity contract.
type BlindAuctionRaw struct {
	Contract *BlindAuction // Generic contract binding to access the raw methods on
}

// BlindAuctionCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type BlindAuctionCallerRaw struct {
	Contract *BlindAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// BlindAuctionTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type BlindAuctionTransactorRaw struct {
	Contract *BlindAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlindAuction creates a new instance of BlindAuction, bound to a specific deployed contract.
func NewBlindAuction(address common.Address, backend bind.ContractBackend) (*BlindAuction, error) {
	contract, err := bindBlindAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlindAuction{BlindAuctionCaller: BlindAuctionCaller{contract: contract}, BlindAuctionTransactor: BlindAuctionTransactor{contract: contract}, BlindAuctionFilterer: BlindAuctionFilterer{contract: contract}}, nil
}

// NewBlindAuctionCaller creates a new read-only instance of BlindAuction, bound to a specific deployed contract.
func NewBlindAuctionCaller(address common.Address, caller bind.ContractCaller) (*BlindAuctionCaller, error) {
	contract, err := bindBlindAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlindAuctionCaller{contract: contract}, nil
}

// NewBlindAuctionTransactor creates a new write-only instance of BlindAuction, bound to a specific deployed contract.
func NewBlindAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*BlindAuctionTransactor, error) {
	contract, err := bindBlindAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlindAuctionTransactor{contract: contract}, nil
}

// NewBlindAuctionFilterer creates a new log filterer instance of BlindAuction, bound to a specific deployed contract.
func NewBlindAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*BlindAuctionFilterer, error) {
	contract, err := bindBlindAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlindAuctionFilterer{contract: contract}, nil
}

// bindBlindAuction binds a generic wrapper to an already deployed contract.
func bindBlindAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlindAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlindAuction *BlindAuctionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BlindAuction.Contract.BlindAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlindAuction *BlindAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _BlindAuction.Contract.BlindAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlindAuction *BlindAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _BlindAuction.Contract.BlindAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlindAuction *BlindAuctionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BlindAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlindAuction *BlindAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _BlindAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlindAuction *BlindAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _BlindAuction.Contract.contract.Transact(opts, method, params...)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_BlindAuction *BlindAuctionCaller) State(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _BlindAuction.contract.Call(opts, out, "state")
	return *ret0, err
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_BlindAuction *BlindAuctionSession) State() (uint8, error) {
	return _BlindAuction.Contract.State(&_BlindAuction.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_BlindAuction *BlindAuctionCallerSession) State() (uint8, error) {
	return _BlindAuction.Contract.State(&_BlindAuction.CallOpts)
}

// AuctionEnd is a paid mutator transaction binding the contract method 0x2a24f46c.
//
// Solidity: function auctionEnd() returns()
func (_BlindAuction *BlindAuctionTransactor) AuctionEnd(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _BlindAuction.contract.Transact(opts, "auctionEnd")
}

func (_BlindAuction *BlindAuctionTransactor) AsyncAuctionEnd(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlindAuction.contract.AsyncTransact(opts, handler, "auctionEnd")
}

// AuctionEnd is a paid mutator transaction binding the contract method 0x2a24f46c.
//
// Solidity: function auctionEnd() returns()
func (_BlindAuction *BlindAuctionSession) AuctionEnd() (*types.Transaction, *types.Receipt, error) {
	return _BlindAuction.Contract.AuctionEnd(&_BlindAuction.TransactOpts)
}

func (_BlindAuction *BlindAuctionSession) AsyncAuctionEnd(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _BlindAuction.Contract.AsyncAuctionEnd(handler, &_BlindAuction.TransactOpts)
}

// AuctionEnd is a paid mutator transaction binding the contract method 0x2a24f46c.
//
// Solidity: function auctionEnd() returns()
func (_BlindAuction *BlindAuctionTransactorSession) AuctionEnd() (*types.Transaction, *types.Receipt, error) {
	return _BlindAuction.Contract.AuctionEnd(&_BlindAuction.TransactOpts)
}

func (_BlindAuction *BlindAuctionTransactorSession) AsyncAuctionEnd(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _BlindAuction.Contract.AsyncAuctionEnd(handler, &_BlindAuction.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() returns()
func (_BlindAuction *BlindAuctionTransactor) Bid(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _BlindAuction.contract.Transact(opts, "bid")
}

func (_BlindAuction *BlindAuctionTransactor) AsyncBid(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlindAuction.contract.AsyncTransact(opts, handler, "bid")
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() returns()
func (_BlindAuction *BlindAuctionSession) Bid() (*types.Transaction, *types.Receipt, error) {
	return _BlindAuction.Contract.Bid(&_BlindAuction.TransactOpts)
}

func (_BlindAuction *BlindAuctionSession) AsyncBid(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _BlindAuction.Contract.AsyncBid(handler, &_BlindAuction.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() returns()
func (_BlindAuction *BlindAuctionTransactorSession) Bid() (*types.Transaction, *types.Receipt, error) {
	return _BlindAuction.Contract.Bid(&_BlindAuction.TransactOpts)
}

func (_BlindAuction *BlindAuctionTransactorSession) AsyncBid(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _BlindAuction.Contract.AsyncBid(handler, &_BlindAuction.TransactOpts)
}

// GetResult is a paid mutator transaction binding the contract method 0xde292789.
//
// Solidity: function getResult() returns()
func (_BlindAuction *BlindAuctionTransactor) GetResult(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _BlindAuction.contract.Transact(opts, "getResult")
}

func (_BlindAuction *BlindAuctionTransactor) AsyncGetResult(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlindAuction.contract.AsyncTransact(opts, handler, "getResult")
}

// GetResult is a paid mutator transaction binding the contract method 0xde292789.
//
// Solidity: function getResult() returns()
func (_BlindAuction *BlindAuctionSession) GetResult() (*types.Transaction, *types.Receipt, error) {
	return _BlindAuction.Contract.GetResult(&_BlindAuction.TransactOpts)
}

func (_BlindAuction *BlindAuctionSession) AsyncGetResult(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _BlindAuction.Contract.AsyncGetResult(handler, &_BlindAuction.TransactOpts)
}

// GetResult is a paid mutator transaction binding the contract method 0xde292789.
//
// Solidity: function getResult() returns()
func (_BlindAuction *BlindAuctionTransactorSession) GetResult() (*types.Transaction, *types.Receipt, error) {
	return _BlindAuction.Contract.GetResult(&_BlindAuction.TransactOpts)
}

func (_BlindAuction *BlindAuctionTransactorSession) AsyncGetResult(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _BlindAuction.Contract.AsyncGetResult(handler, &_BlindAuction.TransactOpts)
}

// Reveal is a paid mutator transaction binding the contract method 0xa475b5dd.
//
// Solidity: function reveal() returns()
func (_BlindAuction *BlindAuctionTransactor) Reveal(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _BlindAuction.contract.Transact(opts, "reveal")
}

func (_BlindAuction *BlindAuctionTransactor) AsyncReveal(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlindAuction.contract.AsyncTransact(opts, handler, "reveal")
}

// Reveal is a paid mutator transaction binding the contract method 0xa475b5dd.
//
// Solidity: function reveal() returns()
func (_BlindAuction *BlindAuctionSession) Reveal() (*types.Transaction, *types.Receipt, error) {
	return _BlindAuction.Contract.Reveal(&_BlindAuction.TransactOpts)
}

func (_BlindAuction *BlindAuctionSession) AsyncReveal(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _BlindAuction.Contract.AsyncReveal(handler, &_BlindAuction.TransactOpts)
}

// Reveal is a paid mutator transaction binding the contract method 0xa475b5dd.
//
// Solidity: function reveal() returns()
func (_BlindAuction *BlindAuctionTransactorSession) Reveal() (*types.Transaction, *types.Receipt, error) {
	return _BlindAuction.Contract.Reveal(&_BlindAuction.TransactOpts)
}

func (_BlindAuction *BlindAuctionTransactorSession) AsyncReveal(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _BlindAuction.Contract.AsyncReveal(handler, &_BlindAuction.TransactOpts)
}
