// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package faastokenpay

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

// FaaSTokenPayABI is the input ABI used to generate the binding from.
const FaaSTokenPayABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContractAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"tokenContract\",\"outputs\":[{\"internalType\":\"contractFaaSToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FaaSTokenPayBin is the compiled bytecode used for deploying new contracts.
var FaaSTokenPayBin = "0x608060405234801561001057600080fd5b506040516101633803806101638339818101604052602081101561003357600080fd5b8101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505060d0806100936000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c806355a373d614602d575b600080fd5b60336075565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff168156fea26469706673582212205d2686846bb0adc3ea8fb116cb5f2c07aedf9f9584abaf99440c86587ec26c0664736f6c634300060a0033"

// DeployFaaSTokenPay deploys a new contract, binding an instance of FaaSTokenPay to it.
func DeployFaaSTokenPay(auth *bind.TransactOpts, backend bind.ContractBackend, _tokenContractAddress common.Address) (common.Address, *types.Transaction, *FaaSTokenPay, error) {
	parsed, err := abi.JSON(strings.NewReader(FaaSTokenPayABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FaaSTokenPayBin), backend, _tokenContractAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FaaSTokenPay{FaaSTokenPayCaller: FaaSTokenPayCaller{contract: contract}, FaaSTokenPayTransactor: FaaSTokenPayTransactor{contract: contract}, FaaSTokenPayFilterer: FaaSTokenPayFilterer{contract: contract}}, nil
}

func AsyncDeployFaaSTokenPay(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend, _tokenContractAddress common.Address) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(FaaSTokenPayABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(FaaSTokenPayBin), backend, _tokenContractAddress)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// FaaSTokenPay is an auto generated Go binding around a Solidity contract.
type FaaSTokenPay struct {
	FaaSTokenPayCaller     // Read-only binding to the contract
	FaaSTokenPayTransactor // Write-only binding to the contract
	FaaSTokenPayFilterer   // Log filterer for contract events
}

// FaaSTokenPayCaller is an auto generated read-only Go binding around a Solidity contract.
type FaaSTokenPayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaaSTokenPayTransactor is an auto generated write-only Go binding around a Solidity contract.
type FaaSTokenPayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaaSTokenPayFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type FaaSTokenPayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaaSTokenPaySession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type FaaSTokenPaySession struct {
	Contract     *FaaSTokenPay     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FaaSTokenPayCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type FaaSTokenPayCallerSession struct {
	Contract *FaaSTokenPayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// FaaSTokenPayTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type FaaSTokenPayTransactorSession struct {
	Contract     *FaaSTokenPayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// FaaSTokenPayRaw is an auto generated low-level Go binding around a Solidity contract.
type FaaSTokenPayRaw struct {
	Contract *FaaSTokenPay // Generic contract binding to access the raw methods on
}

// FaaSTokenPayCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type FaaSTokenPayCallerRaw struct {
	Contract *FaaSTokenPayCaller // Generic read-only contract binding to access the raw methods on
}

// FaaSTokenPayTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type FaaSTokenPayTransactorRaw struct {
	Contract *FaaSTokenPayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFaaSTokenPay creates a new instance of FaaSTokenPay, bound to a specific deployed contract.
func NewFaaSTokenPay(address common.Address, backend bind.ContractBackend) (*FaaSTokenPay, error) {
	contract, err := bindFaaSTokenPay(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FaaSTokenPay{FaaSTokenPayCaller: FaaSTokenPayCaller{contract: contract}, FaaSTokenPayTransactor: FaaSTokenPayTransactor{contract: contract}, FaaSTokenPayFilterer: FaaSTokenPayFilterer{contract: contract}}, nil
}

// NewFaaSTokenPayCaller creates a new read-only instance of FaaSTokenPay, bound to a specific deployed contract.
func NewFaaSTokenPayCaller(address common.Address, caller bind.ContractCaller) (*FaaSTokenPayCaller, error) {
	contract, err := bindFaaSTokenPay(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FaaSTokenPayCaller{contract: contract}, nil
}

// NewFaaSTokenPayTransactor creates a new write-only instance of FaaSTokenPay, bound to a specific deployed contract.
func NewFaaSTokenPayTransactor(address common.Address, transactor bind.ContractTransactor) (*FaaSTokenPayTransactor, error) {
	contract, err := bindFaaSTokenPay(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FaaSTokenPayTransactor{contract: contract}, nil
}

// NewFaaSTokenPayFilterer creates a new log filterer instance of FaaSTokenPay, bound to a specific deployed contract.
func NewFaaSTokenPayFilterer(address common.Address, filterer bind.ContractFilterer) (*FaaSTokenPayFilterer, error) {
	contract, err := bindFaaSTokenPay(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FaaSTokenPayFilterer{contract: contract}, nil
}

// bindFaaSTokenPay binds a generic wrapper to an already deployed contract.
func bindFaaSTokenPay(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FaaSTokenPayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FaaSTokenPay *FaaSTokenPayRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FaaSTokenPay.Contract.FaaSTokenPayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FaaSTokenPay *FaaSTokenPayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _FaaSTokenPay.Contract.FaaSTokenPayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FaaSTokenPay *FaaSTokenPayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _FaaSTokenPay.Contract.FaaSTokenPayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FaaSTokenPay *FaaSTokenPayCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FaaSTokenPay.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FaaSTokenPay *FaaSTokenPayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _FaaSTokenPay.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FaaSTokenPay *FaaSTokenPayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _FaaSTokenPay.Contract.contract.Transact(opts, method, params...)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_FaaSTokenPay *FaaSTokenPayCaller) TokenContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FaaSTokenPay.contract.Call(opts, out, "tokenContract")
	return *ret0, err
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_FaaSTokenPay *FaaSTokenPaySession) TokenContract() (common.Address, error) {
	return _FaaSTokenPay.Contract.TokenContract(&_FaaSTokenPay.CallOpts)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_FaaSTokenPay *FaaSTokenPayCallerSession) TokenContract() (common.Address, error) {
	return _FaaSTokenPay.Contract.TokenContract(&_FaaSTokenPay.CallOpts)
}
