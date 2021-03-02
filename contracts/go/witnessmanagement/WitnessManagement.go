// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package witnessmanagement

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

// WitnessManagementABI is the input ABI used to generate the binding from.
const WitnessManagementABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContractAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_witness\",\"type\":\"address\"},{\"internalType\":\"enumWitnessManagement.WStates\",\"name\":\"_state\",\"type\":\"uint8\"}],\"name\":\"isAtWState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_witness\",\"type\":\"address\"}],\"name\":\"isWitnessQualified\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_witness\",\"type\":\"address\"}],\"name\":\"isWitnessRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stdWitnessDepoist\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenContract\",\"outputs\":[{\"internalType\":\"contractFaaSToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wintessTurnOn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"witenessLogout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"witnessDrawReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"witnessLogin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"witnessReputationInit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"witnessReputationQualified\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"witnessTurnOff\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// WitnessManagementBin is the compiled bytecode used for deploying new contracts.
var WitnessManagementBin = "0x608060405234801561001057600080fd5b506040516113053803806113058339818101604052602081101561003357600080fd5b810190808051906020019092919050505080806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505060006006819055506064600181905550606460028190555060016003819055505061124f806100b66000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80639faa79e3116100715780639faa79e3146101bc578063a35da30c146101c6578063cf617762146101e4578063d446fcdd14610240578063ebe471101461029c578063ecc7e64b146102a6576100b4565b806316c3734d146100b95780632fba9d1a146100d7578063426f93b1146100e157806347ff3aba146100eb57806355a373d61461015457806389640b581461019e575b600080fd5b6100c16102b0565b6040518082815260200191505060405180910390f35b6100df6102b6565b005b6100e96105e5565b005b61013a6004803603604081101561010157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560ff169060200190929190505050610938565b604051808215151515815260200191505060405180910390f35b61015c6109aa565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101a66109cf565b6040518082815260200191505060405180910390f35b6101c46109d5565b005b6101ce610d09565b6040518082815260200191505060405180910390f35b610226600480360360208110156101fa57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d0f565b604051808215151515815260200191505060405180910390f35b6102826004803603602081101561025657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d60565b604051808215151515815260200191505060405180910390f35b6102a4610db9565b005b6102ae610f66565b005b60025481565b600015156102c333610d60565b15151461031b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602c81526020018061113a602c913960400191505060405180910390fd5b6001151561032833610d0f565b151514610380576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260278152602001806110af6027913960400191505060405180910390fd5b6000600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600401541161041b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180611194602a913960400191505060405180910390fd5b6000600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206004015490506000600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600401819055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561055257600080fd5b505af1158015610566573d6000803e3d6000fd5b505050506040513d602081101561057c57600080fd5b81019080805190602001909291905050506105e2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e815260200180611166602e913960400191505060405180910390fd5b50565b600115156105f233610d60565b15151461064a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806110d66030913960400191505060405180910390fd5b6001151561065733610d0f565b1515146106af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260278152602001806110af6027913960400191505060405180910390fd5b6000600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206003015490506000600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548160ff0219169083600381111561075557fe5b02179055506000600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160016101000a81548160ff0219169083151502179055506000600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600301819055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156108a557600080fd5b505af11580156108b9573d6000803e3d6000fd5b505050506040513d60208110156108cf57600080fd5b8101908080519060200190929190505050610935576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260318152602001806111be6031913960400191505060405180910390fd5b50565b600081600381111561094657fe5b600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff1660038111156109a157fe5b14905092915050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035481565b600015156109e233610d60565b151514610a3a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602c81526020018061113a602c913960400191505060405180910390fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd33306001546040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b158015610b1857600080fd5b505af1158015610b2c573d6000803e3d6000fd5b505050506040513d6020811015610b4257600080fd5b8101908080519060200190929190505050610ba8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602b8152602001806111ef602b913960400191505060405180910390fd5b6005339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040518060c0016040528060006003811115610c2357fe5b81526020016001151581526020016001600580549050038152602001600254815260200160015481526020016000815250600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690836003811115610cb757fe5b021790555060208201518160000160016101000a81548160ff02191690831515021790555060408201518160010155606082015181600201556080820151816003015560a08201518160040155905050565b60015481565b6000600354600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002015410159050919050565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160019054906101000a900460ff169050919050565b60011515610dc633610d60565b151514610e1e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806110d66030913960400191505060405180910390fd5b60011515610e2b33610d0f565b151514610e83576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260278152602001806110af6027913960400191505060405180910390fd5b33600060011515610e948383610938565b151514610eec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260348152602001806111066034913960400191505060405180910390fd5b6001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548160ff02191690836003811115610f4b57fe5b02179055506006600081548092919060010191905055505050565b60011515610f7333610d60565b151514610fcb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806110d66030913960400191505060405180910390fd5b336001801515610fdb8383610938565b151514611033576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260348152602001806111066034913960400191505060405180910390fd5b6000600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548160ff0219169083600381111561109257fe5b021790555060066000815480929190600190039190505550505056fe5769746e657373506f6f6c3a20746865207769746e65737320697320756e7175616c69666965645769746e657373506f6f6c3a20746865206164647265737320686164206e6f74206265656e2072656769737465726564576974656e7373506f6f6c3a2066756e6374696f6e2063616e6e6f742062652063616c6c656420617420746869732073746174655769746e657373506f6f6c3a20746865206164647265737320686164206265656e20726567697374657265645769746e657373506f6f6c3a206661696c656420746f206472617720746865207769746e657373207265776172645769746e657373506f6f6c3a2074686520726577617264206f66207769746e657373206973207a65726f5769746e657373506f6f6c3a206661696c656420746f20726566756e6420746865207769746e657373206465706f7369745769746e657373506f6f6c3a206661696c656420746f206c6f636b207769746e657373206465706f697374a26469706673582212208abc40090b7c7b127944b3b1e26bc2459a52a9763cb0973e8df1d41f73564e5364736f6c634300060a0033"

// DeployWitnessManagement deploys a new contract, binding an instance of WitnessManagement to it.
func DeployWitnessManagement(auth *bind.TransactOpts, backend bind.ContractBackend, _tokenContractAddress common.Address) (common.Address, *types.Transaction, *WitnessManagement, error) {
	parsed, err := abi.JSON(strings.NewReader(WitnessManagementABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WitnessManagementBin), backend, _tokenContractAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WitnessManagement{WitnessManagementCaller: WitnessManagementCaller{contract: contract}, WitnessManagementTransactor: WitnessManagementTransactor{contract: contract}, WitnessManagementFilterer: WitnessManagementFilterer{contract: contract}}, nil
}

func AsyncDeployWitnessManagement(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend, _tokenContractAddress common.Address) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(WitnessManagementABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(WitnessManagementBin), backend, _tokenContractAddress)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// WitnessManagement is an auto generated Go binding around a Solidity contract.
type WitnessManagement struct {
	WitnessManagementCaller     // Read-only binding to the contract
	WitnessManagementTransactor // Write-only binding to the contract
	WitnessManagementFilterer   // Log filterer for contract events
}

// WitnessManagementCaller is an auto generated read-only Go binding around a Solidity contract.
type WitnessManagementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WitnessManagementTransactor is an auto generated write-only Go binding around a Solidity contract.
type WitnessManagementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WitnessManagementFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type WitnessManagementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WitnessManagementSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type WitnessManagementSession struct {
	Contract     *WitnessManagement // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// WitnessManagementCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type WitnessManagementCallerSession struct {
	Contract *WitnessManagementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// WitnessManagementTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type WitnessManagementTransactorSession struct {
	Contract     *WitnessManagementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// WitnessManagementRaw is an auto generated low-level Go binding around a Solidity contract.
type WitnessManagementRaw struct {
	Contract *WitnessManagement // Generic contract binding to access the raw methods on
}

// WitnessManagementCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type WitnessManagementCallerRaw struct {
	Contract *WitnessManagementCaller // Generic read-only contract binding to access the raw methods on
}

// WitnessManagementTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type WitnessManagementTransactorRaw struct {
	Contract *WitnessManagementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWitnessManagement creates a new instance of WitnessManagement, bound to a specific deployed contract.
func NewWitnessManagement(address common.Address, backend bind.ContractBackend) (*WitnessManagement, error) {
	contract, err := bindWitnessManagement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WitnessManagement{WitnessManagementCaller: WitnessManagementCaller{contract: contract}, WitnessManagementTransactor: WitnessManagementTransactor{contract: contract}, WitnessManagementFilterer: WitnessManagementFilterer{contract: contract}}, nil
}

// NewWitnessManagementCaller creates a new read-only instance of WitnessManagement, bound to a specific deployed contract.
func NewWitnessManagementCaller(address common.Address, caller bind.ContractCaller) (*WitnessManagementCaller, error) {
	contract, err := bindWitnessManagement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WitnessManagementCaller{contract: contract}, nil
}

// NewWitnessManagementTransactor creates a new write-only instance of WitnessManagement, bound to a specific deployed contract.
func NewWitnessManagementTransactor(address common.Address, transactor bind.ContractTransactor) (*WitnessManagementTransactor, error) {
	contract, err := bindWitnessManagement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WitnessManagementTransactor{contract: contract}, nil
}

// NewWitnessManagementFilterer creates a new log filterer instance of WitnessManagement, bound to a specific deployed contract.
func NewWitnessManagementFilterer(address common.Address, filterer bind.ContractFilterer) (*WitnessManagementFilterer, error) {
	contract, err := bindWitnessManagement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WitnessManagementFilterer{contract: contract}, nil
}

// bindWitnessManagement binds a generic wrapper to an already deployed contract.
func bindWitnessManagement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WitnessManagementABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WitnessManagement *WitnessManagementRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WitnessManagement.Contract.WitnessManagementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WitnessManagement *WitnessManagementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _WitnessManagement.Contract.WitnessManagementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WitnessManagement *WitnessManagementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _WitnessManagement.Contract.WitnessManagementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WitnessManagement *WitnessManagementCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WitnessManagement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WitnessManagement *WitnessManagementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _WitnessManagement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WitnessManagement *WitnessManagementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _WitnessManagement.Contract.contract.Transact(opts, method, params...)
}

// IsAtWState is a free data retrieval call binding the contract method 0x47ff3aba.
//
// Solidity: function isAtWState(address _witness, uint8 _state) constant returns(bool)
func (_WitnessManagement *WitnessManagementCaller) IsAtWState(opts *bind.CallOpts, _witness common.Address, _state uint8) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WitnessManagement.contract.Call(opts, out, "isAtWState", _witness, _state)
	return *ret0, err
}

// IsAtWState is a free data retrieval call binding the contract method 0x47ff3aba.
//
// Solidity: function isAtWState(address _witness, uint8 _state) constant returns(bool)
func (_WitnessManagement *WitnessManagementSession) IsAtWState(_witness common.Address, _state uint8) (bool, error) {
	return _WitnessManagement.Contract.IsAtWState(&_WitnessManagement.CallOpts, _witness, _state)
}

// IsAtWState is a free data retrieval call binding the contract method 0x47ff3aba.
//
// Solidity: function isAtWState(address _witness, uint8 _state) constant returns(bool)
func (_WitnessManagement *WitnessManagementCallerSession) IsAtWState(_witness common.Address, _state uint8) (bool, error) {
	return _WitnessManagement.Contract.IsAtWState(&_WitnessManagement.CallOpts, _witness, _state)
}

// IsWitnessQualified is a free data retrieval call binding the contract method 0xcf617762.
//
// Solidity: function isWitnessQualified(address _witness) constant returns(bool)
func (_WitnessManagement *WitnessManagementCaller) IsWitnessQualified(opts *bind.CallOpts, _witness common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WitnessManagement.contract.Call(opts, out, "isWitnessQualified", _witness)
	return *ret0, err
}

// IsWitnessQualified is a free data retrieval call binding the contract method 0xcf617762.
//
// Solidity: function isWitnessQualified(address _witness) constant returns(bool)
func (_WitnessManagement *WitnessManagementSession) IsWitnessQualified(_witness common.Address) (bool, error) {
	return _WitnessManagement.Contract.IsWitnessQualified(&_WitnessManagement.CallOpts, _witness)
}

// IsWitnessQualified is a free data retrieval call binding the contract method 0xcf617762.
//
// Solidity: function isWitnessQualified(address _witness) constant returns(bool)
func (_WitnessManagement *WitnessManagementCallerSession) IsWitnessQualified(_witness common.Address) (bool, error) {
	return _WitnessManagement.Contract.IsWitnessQualified(&_WitnessManagement.CallOpts, _witness)
}

// IsWitnessRegistered is a free data retrieval call binding the contract method 0xd446fcdd.
//
// Solidity: function isWitnessRegistered(address _witness) constant returns(bool)
func (_WitnessManagement *WitnessManagementCaller) IsWitnessRegistered(opts *bind.CallOpts, _witness common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WitnessManagement.contract.Call(opts, out, "isWitnessRegistered", _witness)
	return *ret0, err
}

// IsWitnessRegistered is a free data retrieval call binding the contract method 0xd446fcdd.
//
// Solidity: function isWitnessRegistered(address _witness) constant returns(bool)
func (_WitnessManagement *WitnessManagementSession) IsWitnessRegistered(_witness common.Address) (bool, error) {
	return _WitnessManagement.Contract.IsWitnessRegistered(&_WitnessManagement.CallOpts, _witness)
}

// IsWitnessRegistered is a free data retrieval call binding the contract method 0xd446fcdd.
//
// Solidity: function isWitnessRegistered(address _witness) constant returns(bool)
func (_WitnessManagement *WitnessManagementCallerSession) IsWitnessRegistered(_witness common.Address) (bool, error) {
	return _WitnessManagement.Contract.IsWitnessRegistered(&_WitnessManagement.CallOpts, _witness)
}

// StdWitnessDepoist is a free data retrieval call binding the contract method 0xa35da30c.
//
// Solidity: function stdWitnessDepoist() constant returns(uint256)
func (_WitnessManagement *WitnessManagementCaller) StdWitnessDepoist(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WitnessManagement.contract.Call(opts, out, "stdWitnessDepoist")
	return *ret0, err
}

// StdWitnessDepoist is a free data retrieval call binding the contract method 0xa35da30c.
//
// Solidity: function stdWitnessDepoist() constant returns(uint256)
func (_WitnessManagement *WitnessManagementSession) StdWitnessDepoist() (*big.Int, error) {
	return _WitnessManagement.Contract.StdWitnessDepoist(&_WitnessManagement.CallOpts)
}

// StdWitnessDepoist is a free data retrieval call binding the contract method 0xa35da30c.
//
// Solidity: function stdWitnessDepoist() constant returns(uint256)
func (_WitnessManagement *WitnessManagementCallerSession) StdWitnessDepoist() (*big.Int, error) {
	return _WitnessManagement.Contract.StdWitnessDepoist(&_WitnessManagement.CallOpts)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_WitnessManagement *WitnessManagementCaller) TokenContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WitnessManagement.contract.Call(opts, out, "tokenContract")
	return *ret0, err
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_WitnessManagement *WitnessManagementSession) TokenContract() (common.Address, error) {
	return _WitnessManagement.Contract.TokenContract(&_WitnessManagement.CallOpts)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_WitnessManagement *WitnessManagementCallerSession) TokenContract() (common.Address, error) {
	return _WitnessManagement.Contract.TokenContract(&_WitnessManagement.CallOpts)
}

// WitnessReputationInit is a free data retrieval call binding the contract method 0x16c3734d.
//
// Solidity: function witnessReputationInit() constant returns(uint256)
func (_WitnessManagement *WitnessManagementCaller) WitnessReputationInit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WitnessManagement.contract.Call(opts, out, "witnessReputationInit")
	return *ret0, err
}

// WitnessReputationInit is a free data retrieval call binding the contract method 0x16c3734d.
//
// Solidity: function witnessReputationInit() constant returns(uint256)
func (_WitnessManagement *WitnessManagementSession) WitnessReputationInit() (*big.Int, error) {
	return _WitnessManagement.Contract.WitnessReputationInit(&_WitnessManagement.CallOpts)
}

// WitnessReputationInit is a free data retrieval call binding the contract method 0x16c3734d.
//
// Solidity: function witnessReputationInit() constant returns(uint256)
func (_WitnessManagement *WitnessManagementCallerSession) WitnessReputationInit() (*big.Int, error) {
	return _WitnessManagement.Contract.WitnessReputationInit(&_WitnessManagement.CallOpts)
}

// WitnessReputationQualified is a free data retrieval call binding the contract method 0x89640b58.
//
// Solidity: function witnessReputationQualified() constant returns(uint256)
func (_WitnessManagement *WitnessManagementCaller) WitnessReputationQualified(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WitnessManagement.contract.Call(opts, out, "witnessReputationQualified")
	return *ret0, err
}

// WitnessReputationQualified is a free data retrieval call binding the contract method 0x89640b58.
//
// Solidity: function witnessReputationQualified() constant returns(uint256)
func (_WitnessManagement *WitnessManagementSession) WitnessReputationQualified() (*big.Int, error) {
	return _WitnessManagement.Contract.WitnessReputationQualified(&_WitnessManagement.CallOpts)
}

// WitnessReputationQualified is a free data retrieval call binding the contract method 0x89640b58.
//
// Solidity: function witnessReputationQualified() constant returns(uint256)
func (_WitnessManagement *WitnessManagementCallerSession) WitnessReputationQualified() (*big.Int, error) {
	return _WitnessManagement.Contract.WitnessReputationQualified(&_WitnessManagement.CallOpts)
}

// WintessTurnOn is a paid mutator transaction binding the contract method 0xebe47110.
//
// Solidity: function wintessTurnOn() returns()
func (_WitnessManagement *WitnessManagementTransactor) WintessTurnOn(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _WitnessManagement.contract.Transact(opts, "wintessTurnOn")
}

func (_WitnessManagement *WitnessManagementTransactor) AsyncWintessTurnOn(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessManagement.contract.AsyncTransact(opts, handler, "wintessTurnOn")
}

// WintessTurnOn is a paid mutator transaction binding the contract method 0xebe47110.
//
// Solidity: function wintessTurnOn() returns()
func (_WitnessManagement *WitnessManagementSession) WintessTurnOn() (*types.Transaction, *types.Receipt, error) {
	return _WitnessManagement.Contract.WintessTurnOn(&_WitnessManagement.TransactOpts)
}

func (_WitnessManagement *WitnessManagementSession) AsyncWintessTurnOn(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _WitnessManagement.Contract.AsyncWintessTurnOn(handler, &_WitnessManagement.TransactOpts)
}

// WintessTurnOn is a paid mutator transaction binding the contract method 0xebe47110.
//
// Solidity: function wintessTurnOn() returns()
func (_WitnessManagement *WitnessManagementTransactorSession) WintessTurnOn() (*types.Transaction, *types.Receipt, error) {
	return _WitnessManagement.Contract.WintessTurnOn(&_WitnessManagement.TransactOpts)
}

func (_WitnessManagement *WitnessManagementTransactorSession) AsyncWintessTurnOn(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _WitnessManagement.Contract.AsyncWintessTurnOn(handler, &_WitnessManagement.TransactOpts)
}

// WitenessLogout is a paid mutator transaction binding the contract method 0x426f93b1.
//
// Solidity: function witenessLogout() returns()
func (_WitnessManagement *WitnessManagementTransactor) WitenessLogout(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _WitnessManagement.contract.Transact(opts, "witenessLogout")
}

func (_WitnessManagement *WitnessManagementTransactor) AsyncWitenessLogout(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessManagement.contract.AsyncTransact(opts, handler, "witenessLogout")
}

// WitenessLogout is a paid mutator transaction binding the contract method 0x426f93b1.
//
// Solidity: function witenessLogout() returns()
func (_WitnessManagement *WitnessManagementSession) WitenessLogout() (*types.Transaction, *types.Receipt, error) {
	return _WitnessManagement.Contract.WitenessLogout(&_WitnessManagement.TransactOpts)
}

func (_WitnessManagement *WitnessManagementSession) AsyncWitenessLogout(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _WitnessManagement.Contract.AsyncWitenessLogout(handler, &_WitnessManagement.TransactOpts)
}

// WitenessLogout is a paid mutator transaction binding the contract method 0x426f93b1.
//
// Solidity: function witenessLogout() returns()
func (_WitnessManagement *WitnessManagementTransactorSession) WitenessLogout() (*types.Transaction, *types.Receipt, error) {
	return _WitnessManagement.Contract.WitenessLogout(&_WitnessManagement.TransactOpts)
}

func (_WitnessManagement *WitnessManagementTransactorSession) AsyncWitenessLogout(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _WitnessManagement.Contract.AsyncWitenessLogout(handler, &_WitnessManagement.TransactOpts)
}

// WitnessDrawReward is a paid mutator transaction binding the contract method 0x2fba9d1a.
//
// Solidity: function witnessDrawReward() returns()
func (_WitnessManagement *WitnessManagementTransactor) WitnessDrawReward(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _WitnessManagement.contract.Transact(opts, "witnessDrawReward")
}

func (_WitnessManagement *WitnessManagementTransactor) AsyncWitnessDrawReward(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessManagement.contract.AsyncTransact(opts, handler, "witnessDrawReward")
}

// WitnessDrawReward is a paid mutator transaction binding the contract method 0x2fba9d1a.
//
// Solidity: function witnessDrawReward() returns()
func (_WitnessManagement *WitnessManagementSession) WitnessDrawReward() (*types.Transaction, *types.Receipt, error) {
	return _WitnessManagement.Contract.WitnessDrawReward(&_WitnessManagement.TransactOpts)
}

func (_WitnessManagement *WitnessManagementSession) AsyncWitnessDrawReward(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _WitnessManagement.Contract.AsyncWitnessDrawReward(handler, &_WitnessManagement.TransactOpts)
}

// WitnessDrawReward is a paid mutator transaction binding the contract method 0x2fba9d1a.
//
// Solidity: function witnessDrawReward() returns()
func (_WitnessManagement *WitnessManagementTransactorSession) WitnessDrawReward() (*types.Transaction, *types.Receipt, error) {
	return _WitnessManagement.Contract.WitnessDrawReward(&_WitnessManagement.TransactOpts)
}

func (_WitnessManagement *WitnessManagementTransactorSession) AsyncWitnessDrawReward(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _WitnessManagement.Contract.AsyncWitnessDrawReward(handler, &_WitnessManagement.TransactOpts)
}

// WitnessLogin is a paid mutator transaction binding the contract method 0x9faa79e3.
//
// Solidity: function witnessLogin() returns()
func (_WitnessManagement *WitnessManagementTransactor) WitnessLogin(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _WitnessManagement.contract.Transact(opts, "witnessLogin")
}

func (_WitnessManagement *WitnessManagementTransactor) AsyncWitnessLogin(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessManagement.contract.AsyncTransact(opts, handler, "witnessLogin")
}

// WitnessLogin is a paid mutator transaction binding the contract method 0x9faa79e3.
//
// Solidity: function witnessLogin() returns()
func (_WitnessManagement *WitnessManagementSession) WitnessLogin() (*types.Transaction, *types.Receipt, error) {
	return _WitnessManagement.Contract.WitnessLogin(&_WitnessManagement.TransactOpts)
}

func (_WitnessManagement *WitnessManagementSession) AsyncWitnessLogin(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _WitnessManagement.Contract.AsyncWitnessLogin(handler, &_WitnessManagement.TransactOpts)
}

// WitnessLogin is a paid mutator transaction binding the contract method 0x9faa79e3.
//
// Solidity: function witnessLogin() returns()
func (_WitnessManagement *WitnessManagementTransactorSession) WitnessLogin() (*types.Transaction, *types.Receipt, error) {
	return _WitnessManagement.Contract.WitnessLogin(&_WitnessManagement.TransactOpts)
}

func (_WitnessManagement *WitnessManagementTransactorSession) AsyncWitnessLogin(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _WitnessManagement.Contract.AsyncWitnessLogin(handler, &_WitnessManagement.TransactOpts)
}

// WitnessTurnOff is a paid mutator transaction binding the contract method 0xecc7e64b.
//
// Solidity: function witnessTurnOff() returns()
func (_WitnessManagement *WitnessManagementTransactor) WitnessTurnOff(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _WitnessManagement.contract.Transact(opts, "witnessTurnOff")
}

func (_WitnessManagement *WitnessManagementTransactor) AsyncWitnessTurnOff(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessManagement.contract.AsyncTransact(opts, handler, "witnessTurnOff")
}

// WitnessTurnOff is a paid mutator transaction binding the contract method 0xecc7e64b.
//
// Solidity: function witnessTurnOff() returns()
func (_WitnessManagement *WitnessManagementSession) WitnessTurnOff() (*types.Transaction, *types.Receipt, error) {
	return _WitnessManagement.Contract.WitnessTurnOff(&_WitnessManagement.TransactOpts)
}

func (_WitnessManagement *WitnessManagementSession) AsyncWitnessTurnOff(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _WitnessManagement.Contract.AsyncWitnessTurnOff(handler, &_WitnessManagement.TransactOpts)
}

// WitnessTurnOff is a paid mutator transaction binding the contract method 0xecc7e64b.
//
// Solidity: function witnessTurnOff() returns()
func (_WitnessManagement *WitnessManagementTransactorSession) WitnessTurnOff() (*types.Transaction, *types.Receipt, error) {
	return _WitnessManagement.Contract.WitnessTurnOff(&_WitnessManagement.TransactOpts)
}

func (_WitnessManagement *WitnessManagementTransactorSession) AsyncWitnessTurnOff(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _WitnessManagement.Contract.AsyncWitnessTurnOff(handler, &_WitnessManagement.TransactOpts)
}
