// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package shatest

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

// ShaTestABI is the input ABI used to generate the binding from.
const ShaTestABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getData\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_memory\",\"type\":\"bytes\"}],\"name\":\"getSha256\",\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_memory\",\"type\":\"bytes\"}],\"name\":\"getKeccak256\",\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ShaTestBin is the compiled bytecode used for deploying new contracts.
var ShaTestBin = "0x60806040526040805190810160405280600e81526020017f48656c6c6f2c20536861546573740000000000000000000000000000000000008152506000908051906020019061004f929190610062565b5034801561005c57600080fd5b50610107565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100a357805160ff19168380011785556100d1565b828001600101855582156100d1579182015b828111156100d05782518255916020019190600101906100b5565b5b5090506100de91906100e2565b5090565b61010491905b808211156101005760008160009055506001016100e8565b5090565b90565b6103db806101166000396000f300608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680633bc5de301461005c5780638b053758146100ec57806393730bbe14610171575b600080fd5b34801561006857600080fd5b506100716101f6565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100b1578082015181840152602081019050610096565b50505050905090810190601f1680156100de5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156100f857600080fd5b50610153600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610298565b60405180826000191660001916815260200191505060405180910390f35b34801561017d57600080fd5b506101d8600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610343565b60405180826000191660001916815260200191505060405180910390f35b606060008054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561028e5780601f106102635761010080835404028352916020019161028e565b820191906000526020600020905b81548152906001019060200180831161027157829003601f168201915b5050505050905090565b60006002826040518082805190602001908083835b6020831015156102d257805182526020820191506020810190506020830392506102ad565b6001836020036101000a0380198251168184511680821785525050505050509050019150506020604051808303816000865af1158015610316573d6000803e3d6000fd5b5050506040513d602081101561032b57600080fd5b81019080805190602001909291905050509050919050565b6000816040518082805190602001908083835b60208310151561037b5780518252602082019150602081019050602083039250610356565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902090509190505600a165627a7a72305820a8a8c65eb1bc251ed022e678e6bc35abc47d42379ba68ae0a256557f694456e60029"

// DeployShaTest deploys a new contract, binding an instance of ShaTest to it.
func DeployShaTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ShaTest, error) {
	parsed, err := abi.JSON(strings.NewReader(ShaTestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ShaTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ShaTest{ShaTestCaller: ShaTestCaller{contract: contract}, ShaTestTransactor: ShaTestTransactor{contract: contract}, ShaTestFilterer: ShaTestFilterer{contract: contract}}, nil
}

func AsyncDeployShaTest(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(ShaTestABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(ShaTestBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// ShaTest is an auto generated Go binding around a Solidity contract.
type ShaTest struct {
	ShaTestCaller     // Read-only binding to the contract
	ShaTestTransactor // Write-only binding to the contract
	ShaTestFilterer   // Log filterer for contract events
}

// ShaTestCaller is an auto generated read-only Go binding around a Solidity contract.
type ShaTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShaTestTransactor is an auto generated write-only Go binding around a Solidity contract.
type ShaTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShaTestFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type ShaTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShaTestSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type ShaTestSession struct {
	Contract     *ShaTest          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ShaTestCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type ShaTestCallerSession struct {
	Contract *ShaTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ShaTestTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type ShaTestTransactorSession struct {
	Contract     *ShaTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ShaTestRaw is an auto generated low-level Go binding around a Solidity contract.
type ShaTestRaw struct {
	Contract *ShaTest // Generic contract binding to access the raw methods on
}

// ShaTestCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type ShaTestCallerRaw struct {
	Contract *ShaTestCaller // Generic read-only contract binding to access the raw methods on
}

// ShaTestTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type ShaTestTransactorRaw struct {
	Contract *ShaTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShaTest creates a new instance of ShaTest, bound to a specific deployed contract.
func NewShaTest(address common.Address, backend bind.ContractBackend) (*ShaTest, error) {
	contract, err := bindShaTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ShaTest{ShaTestCaller: ShaTestCaller{contract: contract}, ShaTestTransactor: ShaTestTransactor{contract: contract}, ShaTestFilterer: ShaTestFilterer{contract: contract}}, nil
}

// NewShaTestCaller creates a new read-only instance of ShaTest, bound to a specific deployed contract.
func NewShaTestCaller(address common.Address, caller bind.ContractCaller) (*ShaTestCaller, error) {
	contract, err := bindShaTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShaTestCaller{contract: contract}, nil
}

// NewShaTestTransactor creates a new write-only instance of ShaTest, bound to a specific deployed contract.
func NewShaTestTransactor(address common.Address, transactor bind.ContractTransactor) (*ShaTestTransactor, error) {
	contract, err := bindShaTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShaTestTransactor{contract: contract}, nil
}

// NewShaTestFilterer creates a new log filterer instance of ShaTest, bound to a specific deployed contract.
func NewShaTestFilterer(address common.Address, filterer bind.ContractFilterer) (*ShaTestFilterer, error) {
	contract, err := bindShaTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShaTestFilterer{contract: contract}, nil
}

// bindShaTest binds a generic wrapper to an already deployed contract.
func bindShaTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ShaTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShaTest *ShaTestRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ShaTest.Contract.ShaTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShaTest *ShaTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ShaTest.Contract.ShaTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShaTest *ShaTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _ShaTest.Contract.ShaTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShaTest *ShaTestCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ShaTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShaTest *ShaTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ShaTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShaTest *ShaTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _ShaTest.Contract.contract.Transact(opts, method, params...)
}

// GetData is a free data retrieval call binding the contract method 0x3bc5de30.
//
// Solidity: function getData() constant returns(bytes)
func (_ShaTest *ShaTestCaller) GetData(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _ShaTest.contract.Call(opts, out, "getData")
	return *ret0, err
}

// GetData is a free data retrieval call binding the contract method 0x3bc5de30.
//
// Solidity: function getData() constant returns(bytes)
func (_ShaTest *ShaTestSession) GetData() ([]byte, error) {
	return _ShaTest.Contract.GetData(&_ShaTest.CallOpts)
}

// GetData is a free data retrieval call binding the contract method 0x3bc5de30.
//
// Solidity: function getData() constant returns(bytes)
func (_ShaTest *ShaTestCallerSession) GetData() ([]byte, error) {
	return _ShaTest.Contract.GetData(&_ShaTest.CallOpts)
}

// GetKeccak256 is a paid mutator transaction binding the contract method 0x93730bbe.
//
// Solidity: function getKeccak256(bytes _memory) returns(bytes32 result)
func (_ShaTest *ShaTestTransactor) GetKeccak256(opts *bind.TransactOpts, _memory []byte) (*types.Transaction, *types.Receipt, error) {
	return _ShaTest.contract.Transact(opts, "getKeccak256", _memory)
}

func (_ShaTest *ShaTestTransactor) AsyncGetKeccak256(handler func(*types.Receipt, error), opts *bind.TransactOpts, _memory []byte) (*types.Transaction, error) {
	return _ShaTest.contract.AsyncTransact(opts, handler, "getKeccak256", _memory)
}

// GetKeccak256 is a paid mutator transaction binding the contract method 0x93730bbe.
//
// Solidity: function getKeccak256(bytes _memory) returns(bytes32 result)
func (_ShaTest *ShaTestSession) GetKeccak256(_memory []byte) (*types.Transaction, *types.Receipt, error) {
	return _ShaTest.Contract.GetKeccak256(&_ShaTest.TransactOpts, _memory)
}

func (_ShaTest *ShaTestSession) AsyncGetKeccak256(handler func(*types.Receipt, error), _memory []byte) (*types.Transaction, error) {
	return _ShaTest.Contract.AsyncGetKeccak256(handler, &_ShaTest.TransactOpts, _memory)
}

// GetKeccak256 is a paid mutator transaction binding the contract method 0x93730bbe.
//
// Solidity: function getKeccak256(bytes _memory) returns(bytes32 result)
func (_ShaTest *ShaTestTransactorSession) GetKeccak256(_memory []byte) (*types.Transaction, *types.Receipt, error) {
	return _ShaTest.Contract.GetKeccak256(&_ShaTest.TransactOpts, _memory)
}

func (_ShaTest *ShaTestTransactorSession) AsyncGetKeccak256(handler func(*types.Receipt, error), _memory []byte) (*types.Transaction, error) {
	return _ShaTest.Contract.AsyncGetKeccak256(handler, &_ShaTest.TransactOpts, _memory)
}

// GetSha256 is a paid mutator transaction binding the contract method 0x8b053758.
//
// Solidity: function getSha256(bytes _memory) returns(bytes32 result)
func (_ShaTest *ShaTestTransactor) GetSha256(opts *bind.TransactOpts, _memory []byte) (*types.Transaction, *types.Receipt, error) {
	return _ShaTest.contract.Transact(opts, "getSha256", _memory)
}

func (_ShaTest *ShaTestTransactor) AsyncGetSha256(handler func(*types.Receipt, error), opts *bind.TransactOpts, _memory []byte) (*types.Transaction, error) {
	return _ShaTest.contract.AsyncTransact(opts, handler, "getSha256", _memory)
}

// GetSha256 is a paid mutator transaction binding the contract method 0x8b053758.
//
// Solidity: function getSha256(bytes _memory) returns(bytes32 result)
func (_ShaTest *ShaTestSession) GetSha256(_memory []byte) (*types.Transaction, *types.Receipt, error) {
	return _ShaTest.Contract.GetSha256(&_ShaTest.TransactOpts, _memory)
}

func (_ShaTest *ShaTestSession) AsyncGetSha256(handler func(*types.Receipt, error), _memory []byte) (*types.Transaction, error) {
	return _ShaTest.Contract.AsyncGetSha256(handler, &_ShaTest.TransactOpts, _memory)
}

// GetSha256 is a paid mutator transaction binding the contract method 0x8b053758.
//
// Solidity: function getSha256(bytes _memory) returns(bytes32 result)
func (_ShaTest *ShaTestTransactorSession) GetSha256(_memory []byte) (*types.Transaction, *types.Receipt, error) {
	return _ShaTest.Contract.GetSha256(&_ShaTest.TransactOpts, _memory)
}

func (_ShaTest *ShaTestTransactorSession) AsyncGetSha256(handler func(*types.Receipt, error), _memory []byte) (*types.Transaction, error) {
	return _ShaTest.Contract.AsyncGetSha256(handler, &_ShaTest.TransactOpts, _memory)
}
