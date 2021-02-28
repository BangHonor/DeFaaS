// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package simpleauction

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

// SimpleAuctionABI is the input ABI used to generate the binding from.
const SimpleAuctionABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marketContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_delpoymentOrderID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_highestUnitPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_biddingDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"auctionEnd\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_unitPrice\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"biddingDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delpoymentOrderID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highestUnitPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumSimpleAuction.States\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SimpleAuctionBin is the compiled bytecode used for deploying new contracts.
var SimpleAuctionBin = "0x608060405234801561001057600080fd5b506040516108cb3803806108cb8339818101604052608081101561003357600080fd5b8101908080519060200190929190805190602001909291908051906020019092919080519060200190929190505050836000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008060146101000a81548160ff021916908360018111156100c057fe5b021790555082600181905550816002819055504260038190555080600481905550816005819055506000600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505061078e8061013d6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80635d06978f1161005b5780635d06978f1461014a57806361dcd7ab14610168578063c19d93fb14610186578063c8e7b31d146101b25761007d565b80632a24f46c146100825780634840f2f0146100de57806359d667a5146100fc575b600080fd5b61008a6101d0565b60405180841515151581526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390f35b6100e6610325565b6040518082815260200191505060405180910390f35b6101486004803603604081101561011257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061032b565b005b610152610599565b6040518082815260200191505060405180910390f35b61017061059f565b6040518082815260200191505060405180910390f35b61018e6105a5565b6040518082600181111561019e57fe5b60ff16815260200191505060405180910390f35b6101ba6105b8565b6040518082815260200191505060405180910390f35b60008060008060018111156101e157fe5b600060149054906101000a900460ff1660018111156101fc57fe5b14801561020e57506004546003540142115b1561021c5761021b6105be565b5b600180600181111561022a57fe5b600060149054906101000a900460ff16600181111561024557fe5b1461029b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260368152602001806107236036913960400191505060405180910390fd5b60008073ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415905080600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166005549450945094505050909192565b60025481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103d0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603b8152602001806106e8603b913960400191505060405180910390fd5b600060018111156103dd57fe5b600060149054906101000a900460ff1660018111156103f857fe5b14801561040a57506004546003540142115b15610418576104176105be565b5b600080600181111561042657fe5b600060149054906101000a900460ff16600181111561044157fe5b14610497576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260368152602001806107236036913960400191505060405180910390fd5b6002548211156104f2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252604a81526020018061069e604a913960600191505060405180910390fd5b600554821061054c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603381526020018061066b6033913960400191505060405180910390fd5b8160058190555082600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b60015481565b60035481565b600060149054906101000a900460ff1681565b60045481565b600060018111156105cb57fe5b600060149054906101000a900460ff1660018111156105e657fe5b1415610615576001600060146101000a81548160ff0219169083600181111561060b57fe5b0217905550610668565b60018081111561062157fe5b600060149054906101000a900460ff16600181111561063c57fe5b1415610667576001600060146101000a81548160ff0219169083600181111561066157fe5b02179055505b5b56fe53696d706c6541756374696f6e3a20746865726520616c72656164792069732061206c6f7765722075696e742d70726963652e53696d706c6541756374696f6e3a2074686520756e69742d707269636520697320686967686572207468616e20686967686573742074686520637573746f6d657220616363657074656453696d706c6541756374696f6e3a206f6e6c79206d61726b657420636f6e6374726163742063616e2063616c6c20746869732066756e6369746f6e53696d706c6541756374696f6e3a2066756e6374696f6e2063616e6e6f742062652063616c6c65642061742074686973207374617465a264697066735822122040d46e20005969c88558fb747bdc09e13e96184d64c033c0d995fcd76ca30b0464736f6c634300060a0033"

// DeploySimpleAuction deploys a new contract, binding an instance of SimpleAuction to it.
func DeploySimpleAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _marketContractAddress common.Address, _delpoymentOrderID *big.Int, _highestUnitPrice *big.Int, _biddingDuration *big.Int) (common.Address, *types.Transaction, *SimpleAuction, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleAuctionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SimpleAuctionBin), backend, _marketContractAddress, _delpoymentOrderID, _highestUnitPrice, _biddingDuration)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleAuction{SimpleAuctionCaller: SimpleAuctionCaller{contract: contract}, SimpleAuctionTransactor: SimpleAuctionTransactor{contract: contract}, SimpleAuctionFilterer: SimpleAuctionFilterer{contract: contract}}, nil
}

func AsyncDeploySimpleAuction(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend, _marketContractAddress common.Address, _delpoymentOrderID *big.Int, _highestUnitPrice *big.Int, _biddingDuration *big.Int) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleAuctionABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(SimpleAuctionBin), backend, _marketContractAddress, _delpoymentOrderID, _highestUnitPrice, _biddingDuration)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// SimpleAuction is an auto generated Go binding around a Solidity contract.
type SimpleAuction struct {
	SimpleAuctionCaller     // Read-only binding to the contract
	SimpleAuctionTransactor // Write-only binding to the contract
	SimpleAuctionFilterer   // Log filterer for contract events
}

// SimpleAuctionCaller is an auto generated read-only Go binding around a Solidity contract.
type SimpleAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleAuctionTransactor is an auto generated write-only Go binding around a Solidity contract.
type SimpleAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleAuctionFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type SimpleAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleAuctionSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type SimpleAuctionSession struct {
	Contract     *SimpleAuction    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleAuctionCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type SimpleAuctionCallerSession struct {
	Contract *SimpleAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SimpleAuctionTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type SimpleAuctionTransactorSession struct {
	Contract     *SimpleAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SimpleAuctionRaw is an auto generated low-level Go binding around a Solidity contract.
type SimpleAuctionRaw struct {
	Contract *SimpleAuction // Generic contract binding to access the raw methods on
}

// SimpleAuctionCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type SimpleAuctionCallerRaw struct {
	Contract *SimpleAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleAuctionTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type SimpleAuctionTransactorRaw struct {
	Contract *SimpleAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleAuction creates a new instance of SimpleAuction, bound to a specific deployed contract.
func NewSimpleAuction(address common.Address, backend bind.ContractBackend) (*SimpleAuction, error) {
	contract, err := bindSimpleAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleAuction{SimpleAuctionCaller: SimpleAuctionCaller{contract: contract}, SimpleAuctionTransactor: SimpleAuctionTransactor{contract: contract}, SimpleAuctionFilterer: SimpleAuctionFilterer{contract: contract}}, nil
}

// NewSimpleAuctionCaller creates a new read-only instance of SimpleAuction, bound to a specific deployed contract.
func NewSimpleAuctionCaller(address common.Address, caller bind.ContractCaller) (*SimpleAuctionCaller, error) {
	contract, err := bindSimpleAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleAuctionCaller{contract: contract}, nil
}

// NewSimpleAuctionTransactor creates a new write-only instance of SimpleAuction, bound to a specific deployed contract.
func NewSimpleAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleAuctionTransactor, error) {
	contract, err := bindSimpleAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleAuctionTransactor{contract: contract}, nil
}

// NewSimpleAuctionFilterer creates a new log filterer instance of SimpleAuction, bound to a specific deployed contract.
func NewSimpleAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleAuctionFilterer, error) {
	contract, err := bindSimpleAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleAuctionFilterer{contract: contract}, nil
}

// bindSimpleAuction binds a generic wrapper to an already deployed contract.
func bindSimpleAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleAuction *SimpleAuctionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleAuction.Contract.SimpleAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleAuction *SimpleAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _SimpleAuction.Contract.SimpleAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleAuction *SimpleAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _SimpleAuction.Contract.SimpleAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleAuction *SimpleAuctionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleAuction *SimpleAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _SimpleAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleAuction *SimpleAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _SimpleAuction.Contract.contract.Transact(opts, method, params...)
}

// BiddingDuration is a free data retrieval call binding the contract method 0xc8e7b31d.
//
// Solidity: function biddingDuration() constant returns(uint256)
func (_SimpleAuction *SimpleAuctionCaller) BiddingDuration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleAuction.contract.Call(opts, out, "biddingDuration")
	return *ret0, err
}

// BiddingDuration is a free data retrieval call binding the contract method 0xc8e7b31d.
//
// Solidity: function biddingDuration() constant returns(uint256)
func (_SimpleAuction *SimpleAuctionSession) BiddingDuration() (*big.Int, error) {
	return _SimpleAuction.Contract.BiddingDuration(&_SimpleAuction.CallOpts)
}

// BiddingDuration is a free data retrieval call binding the contract method 0xc8e7b31d.
//
// Solidity: function biddingDuration() constant returns(uint256)
func (_SimpleAuction *SimpleAuctionCallerSession) BiddingDuration() (*big.Int, error) {
	return _SimpleAuction.Contract.BiddingDuration(&_SimpleAuction.CallOpts)
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_SimpleAuction *SimpleAuctionCaller) CreateTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleAuction.contract.Call(opts, out, "createTime")
	return *ret0, err
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_SimpleAuction *SimpleAuctionSession) CreateTime() (*big.Int, error) {
	return _SimpleAuction.Contract.CreateTime(&_SimpleAuction.CallOpts)
}

// CreateTime is a free data retrieval call binding the contract method 0x61dcd7ab.
//
// Solidity: function createTime() constant returns(uint256)
func (_SimpleAuction *SimpleAuctionCallerSession) CreateTime() (*big.Int, error) {
	return _SimpleAuction.Contract.CreateTime(&_SimpleAuction.CallOpts)
}

// DelpoymentOrderID is a free data retrieval call binding the contract method 0x5d06978f.
//
// Solidity: function delpoymentOrderID() constant returns(uint256)
func (_SimpleAuction *SimpleAuctionCaller) DelpoymentOrderID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleAuction.contract.Call(opts, out, "delpoymentOrderID")
	return *ret0, err
}

// DelpoymentOrderID is a free data retrieval call binding the contract method 0x5d06978f.
//
// Solidity: function delpoymentOrderID() constant returns(uint256)
func (_SimpleAuction *SimpleAuctionSession) DelpoymentOrderID() (*big.Int, error) {
	return _SimpleAuction.Contract.DelpoymentOrderID(&_SimpleAuction.CallOpts)
}

// DelpoymentOrderID is a free data retrieval call binding the contract method 0x5d06978f.
//
// Solidity: function delpoymentOrderID() constant returns(uint256)
func (_SimpleAuction *SimpleAuctionCallerSession) DelpoymentOrderID() (*big.Int, error) {
	return _SimpleAuction.Contract.DelpoymentOrderID(&_SimpleAuction.CallOpts)
}

// HighestUnitPrice is a free data retrieval call binding the contract method 0x4840f2f0.
//
// Solidity: function highestUnitPrice() constant returns(uint256)
func (_SimpleAuction *SimpleAuctionCaller) HighestUnitPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleAuction.contract.Call(opts, out, "highestUnitPrice")
	return *ret0, err
}

// HighestUnitPrice is a free data retrieval call binding the contract method 0x4840f2f0.
//
// Solidity: function highestUnitPrice() constant returns(uint256)
func (_SimpleAuction *SimpleAuctionSession) HighestUnitPrice() (*big.Int, error) {
	return _SimpleAuction.Contract.HighestUnitPrice(&_SimpleAuction.CallOpts)
}

// HighestUnitPrice is a free data retrieval call binding the contract method 0x4840f2f0.
//
// Solidity: function highestUnitPrice() constant returns(uint256)
func (_SimpleAuction *SimpleAuctionCallerSession) HighestUnitPrice() (*big.Int, error) {
	return _SimpleAuction.Contract.HighestUnitPrice(&_SimpleAuction.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_SimpleAuction *SimpleAuctionCaller) State(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _SimpleAuction.contract.Call(opts, out, "state")
	return *ret0, err
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_SimpleAuction *SimpleAuctionSession) State() (uint8, error) {
	return _SimpleAuction.Contract.State(&_SimpleAuction.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_SimpleAuction *SimpleAuctionCallerSession) State() (uint8, error) {
	return _SimpleAuction.Contract.State(&_SimpleAuction.CallOpts)
}

// AuctionEnd is a paid mutator transaction binding the contract method 0x2a24f46c.
//
// Solidity: function auctionEnd() returns(bool, address, uint256)
func (_SimpleAuction *SimpleAuctionTransactor) AuctionEnd(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _SimpleAuction.contract.Transact(opts, "auctionEnd")
}

func (_SimpleAuction *SimpleAuctionTransactor) AsyncAuctionEnd(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleAuction.contract.AsyncTransact(opts, handler, "auctionEnd")
}

// AuctionEnd is a paid mutator transaction binding the contract method 0x2a24f46c.
//
// Solidity: function auctionEnd() returns(bool, address, uint256)
func (_SimpleAuction *SimpleAuctionSession) AuctionEnd() (*types.Transaction, *types.Receipt, error) {
	return _SimpleAuction.Contract.AuctionEnd(&_SimpleAuction.TransactOpts)
}

func (_SimpleAuction *SimpleAuctionSession) AsyncAuctionEnd(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _SimpleAuction.Contract.AsyncAuctionEnd(handler, &_SimpleAuction.TransactOpts)
}

// AuctionEnd is a paid mutator transaction binding the contract method 0x2a24f46c.
//
// Solidity: function auctionEnd() returns(bool, address, uint256)
func (_SimpleAuction *SimpleAuctionTransactorSession) AuctionEnd() (*types.Transaction, *types.Receipt, error) {
	return _SimpleAuction.Contract.AuctionEnd(&_SimpleAuction.TransactOpts)
}

func (_SimpleAuction *SimpleAuctionTransactorSession) AsyncAuctionEnd(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _SimpleAuction.Contract.AsyncAuctionEnd(handler, &_SimpleAuction.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x59d667a5.
//
// Solidity: function bid(address _provider, uint256 _unitPrice) returns()
func (_SimpleAuction *SimpleAuctionTransactor) Bid(opts *bind.TransactOpts, _provider common.Address, _unitPrice *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _SimpleAuction.contract.Transact(opts, "bid", _provider, _unitPrice)
}

func (_SimpleAuction *SimpleAuctionTransactor) AsyncBid(handler func(*types.Receipt, error), opts *bind.TransactOpts, _provider common.Address, _unitPrice *big.Int) (*types.Transaction, error) {
	return _SimpleAuction.contract.AsyncTransact(opts, handler, "bid", _provider, _unitPrice)
}

// Bid is a paid mutator transaction binding the contract method 0x59d667a5.
//
// Solidity: function bid(address _provider, uint256 _unitPrice) returns()
func (_SimpleAuction *SimpleAuctionSession) Bid(_provider common.Address, _unitPrice *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _SimpleAuction.Contract.Bid(&_SimpleAuction.TransactOpts, _provider, _unitPrice)
}

func (_SimpleAuction *SimpleAuctionSession) AsyncBid(handler func(*types.Receipt, error), _provider common.Address, _unitPrice *big.Int) (*types.Transaction, error) {
	return _SimpleAuction.Contract.AsyncBid(handler, &_SimpleAuction.TransactOpts, _provider, _unitPrice)
}

// Bid is a paid mutator transaction binding the contract method 0x59d667a5.
//
// Solidity: function bid(address _provider, uint256 _unitPrice) returns()
func (_SimpleAuction *SimpleAuctionTransactorSession) Bid(_provider common.Address, _unitPrice *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _SimpleAuction.Contract.Bid(&_SimpleAuction.TransactOpts, _provider, _unitPrice)
}

func (_SimpleAuction *SimpleAuctionTransactorSession) AsyncBid(handler func(*types.Receipt, error), _provider common.Address, _unitPrice *big.Int) (*types.Transaction, error) {
	return _SimpleAuction.Contract.AsyncBid(handler, &_SimpleAuction.TransactOpts, _provider, _unitPrice)
}