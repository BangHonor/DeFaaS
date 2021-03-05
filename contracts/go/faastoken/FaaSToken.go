// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package faastoken

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

// FaaSTokenABI is the input ABI used to generate the binding from.
const FaaSTokenABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FaaSTokenBin is the compiled bytecode used for deploying new contracts.
var FaaSTokenBin = "0x60806040523480156200001157600080fd5b506040518060400160405280600981526020017f46616153546f6b656e00000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f4653540000000000000000000000000000000000000000000000000000000000815250816003908051906020019062000096929190620000d4565b508060049080519060200190620000af929190620000d4565b506012600560006101000a81548160ff021916908360ff160217905550505062000183565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200011757805160ff191683800117855562000148565b8280016001018555821562000148579182015b82811115620001475782518255916020019190600101906200012a565b5b5090506200015791906200015b565b5090565b6200018091905b808211156200017c57600081600090555060010162000162565b5090565b90565b61122980620001936000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063449a52f81161008c578063a0712d6811610066578063a0712d6814610414578063a457c2d714610442578063a9059cbb146104a8578063dd62ed3e1461050e576100cf565b8063449a52f8146102eb57806370a082311461033957806395d89b4114610391576100cf565b806306fdde03146100d4578063095ea7b31461015757806318160ddd146101bd57806323b872dd146101db578063313ce567146102615780633950935114610285575b600080fd5b6100dc610586565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561011c578082015181840152602081019050610101565b50505050905090810190601f1680156101495780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101a36004803603604081101561016d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610628565b604051808215151515815260200191505060405180910390f35b6101c5610646565b6040518082815260200191505060405180910390f35b610247600480360360608110156101f157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610650565b604051808215151515815260200191505060405180910390f35b61026961075e565b604051808260ff1660ff16815260200191505060405180910390f35b6102d16004803603604081101561029b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610775565b604051808215151515815260200191505060405180910390f35b6103376004803603604081101561030157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610818565b005b61037b6004803603602081101561034f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610826565b6040518082815260200191505060405180910390f35b61039961086e565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103d95780820151818401526020810190506103be565b50505050905090810190601f1680156104065780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6104406004803603602081101561042a57600080fd5b8101908080359060200190929190505050610910565b005b61048e6004803603604081101561045857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610924565b604051808215151515815260200191505060405180910390f35b6104f4600480360360408110156104be57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a25565b604051808215151515815260200191505060405180910390f35b6105706004803603604081101561052457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610a43565b6040518082815260200191505060405180910390f35b606060038054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561061e5780601f106105f35761010080835404028352916020019161061e565b820191906000526020600020905b81548152906001019060200180831161060157829003601f168201915b5050505050905090565b600061063c610635610aca565b8484610ad2565b6001905092915050565b6000600254905090565b600061065d848484610cc9565b6000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006106a8610aca565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508281101561073e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602881526020018061115e6028913960400191505060405180910390fd5b6107528561074a610aca565b858403610ad2565b60019150509392505050565b6000600560009054906101000a900460ff16905090565b600061080e610782610aca565b848460016000610790610aca565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205401610ad2565b6001905092915050565b6108228282610f78565b5050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b606060048054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109065780601f106108db57610100808354040283529160200191610906565b820191906000526020600020905b8154815290600101906020018083116108e957829003601f168201915b5050505050905090565b61092161091b610aca565b82610f78565b50565b60008060016000610933610aca565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905082811015610a06576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806111cf6025913960400191505060405180910390fd5b610a1a610a11610aca565b85858403610ad2565b600191505092915050565b6000610a39610a32610aca565b8484610cc9565b6001905092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610b58576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806111ab6024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610bde576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806111166022913960400191505060405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610d4f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806111866025913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610dd5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001806110f36023913960400191505060405180910390fd5b610de08383836110ed565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610e7c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806111386026913960400191505060405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a350505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561101b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45524332303a206d696e7420746f20746865207a65726f20616464726573730081525060200191505060405180910390fd5b611027600083836110ed565b80600260008282540192505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b50505056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa26469706673582212209a628a0797ce9174edb8572238bbc7dda131be07fb811a731e43feb919c9ff4764736f6c634300060a0033"

// DeployFaaSToken deploys a new contract, binding an instance of FaaSToken to it.
func DeployFaaSToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FaaSToken, error) {
	parsed, err := abi.JSON(strings.NewReader(FaaSTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FaaSTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FaaSToken{FaaSTokenCaller: FaaSTokenCaller{contract: contract}, FaaSTokenTransactor: FaaSTokenTransactor{contract: contract}, FaaSTokenFilterer: FaaSTokenFilterer{contract: contract}}, nil
}

func AsyncDeployFaaSToken(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(FaaSTokenABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(FaaSTokenBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// FaaSToken is an auto generated Go binding around a Solidity contract.
type FaaSToken struct {
	FaaSTokenCaller     // Read-only binding to the contract
	FaaSTokenTransactor // Write-only binding to the contract
	FaaSTokenFilterer   // Log filterer for contract events
}

// FaaSTokenCaller is an auto generated read-only Go binding around a Solidity contract.
type FaaSTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaaSTokenTransactor is an auto generated write-only Go binding around a Solidity contract.
type FaaSTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaaSTokenFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type FaaSTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaaSTokenSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type FaaSTokenSession struct {
	Contract     *FaaSToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FaaSTokenCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type FaaSTokenCallerSession struct {
	Contract *FaaSTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// FaaSTokenTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type FaaSTokenTransactorSession struct {
	Contract     *FaaSTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FaaSTokenRaw is an auto generated low-level Go binding around a Solidity contract.
type FaaSTokenRaw struct {
	Contract *FaaSToken // Generic contract binding to access the raw methods on
}

// FaaSTokenCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type FaaSTokenCallerRaw struct {
	Contract *FaaSTokenCaller // Generic read-only contract binding to access the raw methods on
}

// FaaSTokenTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type FaaSTokenTransactorRaw struct {
	Contract *FaaSTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFaaSToken creates a new instance of FaaSToken, bound to a specific deployed contract.
func NewFaaSToken(address common.Address, backend bind.ContractBackend) (*FaaSToken, error) {
	contract, err := bindFaaSToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FaaSToken{FaaSTokenCaller: FaaSTokenCaller{contract: contract}, FaaSTokenTransactor: FaaSTokenTransactor{contract: contract}, FaaSTokenFilterer: FaaSTokenFilterer{contract: contract}}, nil
}

// NewFaaSTokenCaller creates a new read-only instance of FaaSToken, bound to a specific deployed contract.
func NewFaaSTokenCaller(address common.Address, caller bind.ContractCaller) (*FaaSTokenCaller, error) {
	contract, err := bindFaaSToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FaaSTokenCaller{contract: contract}, nil
}

// NewFaaSTokenTransactor creates a new write-only instance of FaaSToken, bound to a specific deployed contract.
func NewFaaSTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*FaaSTokenTransactor, error) {
	contract, err := bindFaaSToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FaaSTokenTransactor{contract: contract}, nil
}

// NewFaaSTokenFilterer creates a new log filterer instance of FaaSToken, bound to a specific deployed contract.
func NewFaaSTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*FaaSTokenFilterer, error) {
	contract, err := bindFaaSToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FaaSTokenFilterer{contract: contract}, nil
}

// bindFaaSToken binds a generic wrapper to an already deployed contract.
func bindFaaSToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FaaSTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FaaSToken *FaaSTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FaaSToken.Contract.FaaSTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FaaSToken *FaaSTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.Contract.FaaSTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FaaSToken *FaaSTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.Contract.FaaSTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FaaSToken *FaaSTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FaaSToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FaaSToken *FaaSTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FaaSToken *FaaSTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_FaaSToken *FaaSTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FaaSToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_FaaSToken *FaaSTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _FaaSToken.Contract.Allowance(&_FaaSToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_FaaSToken *FaaSTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _FaaSToken.Contract.Allowance(&_FaaSToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_FaaSToken *FaaSTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FaaSToken.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_FaaSToken *FaaSTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _FaaSToken.Contract.BalanceOf(&_FaaSToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_FaaSToken *FaaSTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _FaaSToken.Contract.BalanceOf(&_FaaSToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_FaaSToken *FaaSTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _FaaSToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_FaaSToken *FaaSTokenSession) Decimals() (uint8, error) {
	return _FaaSToken.Contract.Decimals(&_FaaSToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_FaaSToken *FaaSTokenCallerSession) Decimals() (uint8, error) {
	return _FaaSToken.Contract.Decimals(&_FaaSToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_FaaSToken *FaaSTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FaaSToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_FaaSToken *FaaSTokenSession) Name() (string, error) {
	return _FaaSToken.Contract.Name(&_FaaSToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_FaaSToken *FaaSTokenCallerSession) Name() (string, error) {
	return _FaaSToken.Contract.Name(&_FaaSToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_FaaSToken *FaaSTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FaaSToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_FaaSToken *FaaSTokenSession) Symbol() (string, error) {
	return _FaaSToken.Contract.Symbol(&_FaaSToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_FaaSToken *FaaSTokenCallerSession) Symbol() (string, error) {
	return _FaaSToken.Contract.Symbol(&_FaaSToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_FaaSToken *FaaSTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FaaSToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_FaaSToken *FaaSTokenSession) TotalSupply() (*big.Int, error) {
	return _FaaSToken.Contract.TotalSupply(&_FaaSToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_FaaSToken *FaaSTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _FaaSToken.Contract.TotalSupply(&_FaaSToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_FaaSToken *FaaSTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.contract.Transact(opts, "approve", spender, amount)
}

func (_FaaSToken *FaaSTokenTransactor) AsyncApprove(handler func(*types.Receipt, error), opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.contract.AsyncTransact(opts, handler, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_FaaSToken *FaaSTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.Contract.Approve(&_FaaSToken.TransactOpts, spender, amount)
}

func (_FaaSToken *FaaSTokenSession) AsyncApprove(handler func(*types.Receipt, error), spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.AsyncApprove(handler, &_FaaSToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_FaaSToken *FaaSTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.Contract.Approve(&_FaaSToken.TransactOpts, spender, amount)
}

func (_FaaSToken *FaaSTokenTransactorSession) AsyncApprove(handler func(*types.Receipt, error), spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.AsyncApprove(handler, &_FaaSToken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_FaaSToken *FaaSTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

func (_FaaSToken *FaaSTokenTransactor) AsyncDecreaseAllowance(handler func(*types.Receipt, error), opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _FaaSToken.contract.AsyncTransact(opts, handler, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_FaaSToken *FaaSTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.Contract.DecreaseAllowance(&_FaaSToken.TransactOpts, spender, subtractedValue)
}

func (_FaaSToken *FaaSTokenSession) AsyncDecreaseAllowance(handler func(*types.Receipt, error), spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.AsyncDecreaseAllowance(handler, &_FaaSToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_FaaSToken *FaaSTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.Contract.DecreaseAllowance(&_FaaSToken.TransactOpts, spender, subtractedValue)
}

func (_FaaSToken *FaaSTokenTransactorSession) AsyncDecreaseAllowance(handler func(*types.Receipt, error), spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.AsyncDecreaseAllowance(handler, &_FaaSToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_FaaSToken *FaaSTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

func (_FaaSToken *FaaSTokenTransactor) AsyncIncreaseAllowance(handler func(*types.Receipt, error), opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _FaaSToken.contract.AsyncTransact(opts, handler, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_FaaSToken *FaaSTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.Contract.IncreaseAllowance(&_FaaSToken.TransactOpts, spender, addedValue)
}

func (_FaaSToken *FaaSTokenSession) AsyncIncreaseAllowance(handler func(*types.Receipt, error), spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.AsyncIncreaseAllowance(handler, &_FaaSToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_FaaSToken *FaaSTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.Contract.IncreaseAllowance(&_FaaSToken.TransactOpts, spender, addedValue)
}

func (_FaaSToken *FaaSTokenTransactorSession) AsyncIncreaseAllowance(handler func(*types.Receipt, error), spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.AsyncIncreaseAllowance(handler, &_FaaSToken.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_FaaSToken *FaaSTokenTransactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.contract.Transact(opts, "mint", amount)
}

func (_FaaSToken *FaaSTokenTransactor) AsyncMint(handler func(*types.Receipt, error), opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.contract.AsyncTransact(opts, handler, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_FaaSToken *FaaSTokenSession) Mint(amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.Contract.Mint(&_FaaSToken.TransactOpts, amount)
}

func (_FaaSToken *FaaSTokenSession) AsyncMint(handler func(*types.Receipt, error), amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.AsyncMint(handler, &_FaaSToken.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_FaaSToken *FaaSTokenTransactorSession) Mint(amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.Contract.Mint(&_FaaSToken.TransactOpts, amount)
}

func (_FaaSToken *FaaSTokenTransactorSession) AsyncMint(handler func(*types.Receipt, error), amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.AsyncMint(handler, &_FaaSToken.TransactOpts, amount)
}

// MintTo is a paid mutator transaction binding the contract method 0x449a52f8.
//
// Solidity: function mintTo(address account, uint256 amount) returns()
func (_FaaSToken *FaaSTokenTransactor) MintTo(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.contract.Transact(opts, "mintTo", account, amount)
}

func (_FaaSToken *FaaSTokenTransactor) AsyncMintTo(handler func(*types.Receipt, error), opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.contract.AsyncTransact(opts, handler, "mintTo", account, amount)
}

// MintTo is a paid mutator transaction binding the contract method 0x449a52f8.
//
// Solidity: function mintTo(address account, uint256 amount) returns()
func (_FaaSToken *FaaSTokenSession) MintTo(account common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.Contract.MintTo(&_FaaSToken.TransactOpts, account, amount)
}

func (_FaaSToken *FaaSTokenSession) AsyncMintTo(handler func(*types.Receipt, error), account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.AsyncMintTo(handler, &_FaaSToken.TransactOpts, account, amount)
}

// MintTo is a paid mutator transaction binding the contract method 0x449a52f8.
//
// Solidity: function mintTo(address account, uint256 amount) returns()
func (_FaaSToken *FaaSTokenTransactorSession) MintTo(account common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.Contract.MintTo(&_FaaSToken.TransactOpts, account, amount)
}

func (_FaaSToken *FaaSTokenTransactorSession) AsyncMintTo(handler func(*types.Receipt, error), account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.AsyncMintTo(handler, &_FaaSToken.TransactOpts, account, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_FaaSToken *FaaSTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.contract.Transact(opts, "transfer", recipient, amount)
}

func (_FaaSToken *FaaSTokenTransactor) AsyncTransfer(handler func(*types.Receipt, error), opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.contract.AsyncTransact(opts, handler, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_FaaSToken *FaaSTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.Contract.Transfer(&_FaaSToken.TransactOpts, recipient, amount)
}

func (_FaaSToken *FaaSTokenSession) AsyncTransfer(handler func(*types.Receipt, error), recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.AsyncTransfer(handler, &_FaaSToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_FaaSToken *FaaSTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.Contract.Transfer(&_FaaSToken.TransactOpts, recipient, amount)
}

func (_FaaSToken *FaaSTokenTransactorSession) AsyncTransfer(handler func(*types.Receipt, error), recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.AsyncTransfer(handler, &_FaaSToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_FaaSToken *FaaSTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

func (_FaaSToken *FaaSTokenTransactor) AsyncTransferFrom(handler func(*types.Receipt, error), opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.contract.AsyncTransact(opts, handler, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_FaaSToken *FaaSTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.Contract.TransferFrom(&_FaaSToken.TransactOpts, sender, recipient, amount)
}

func (_FaaSToken *FaaSTokenSession) AsyncTransferFrom(handler func(*types.Receipt, error), sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.AsyncTransferFrom(handler, &_FaaSToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_FaaSToken *FaaSTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.Contract.TransferFrom(&_FaaSToken.TransactOpts, sender, recipient, amount)
}

func (_FaaSToken *FaaSTokenTransactorSession) AsyncTransferFrom(handler func(*types.Receipt, error), sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.AsyncTransferFrom(handler, &_FaaSToken.TransactOpts, sender, recipient, amount)
}

// FaaSTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the FaaSToken contract.
type FaaSTokenApprovalIterator struct {
	Event *FaaSTokenApproval // Event containing the contract specifics and raw log

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
func (it *FaaSTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FaaSTokenApproval)
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
		it.Event = new(FaaSTokenApproval)
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
func (it *FaaSTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FaaSTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FaaSTokenApproval represents a Approval event raised by the FaaSToken contract.
type FaaSTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_FaaSToken *FaaSTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*FaaSTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FaaSToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &FaaSTokenApprovalIterator{contract: _FaaSToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_FaaSToken *FaaSTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FaaSTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FaaSToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FaaSTokenApproval)
				if err := _FaaSToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_FaaSToken *FaaSTokenFilterer) ParseApproval(log types.Log) (*FaaSTokenApproval, error) {
	event := new(FaaSTokenApproval)
	if err := _FaaSToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// FaaSTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the FaaSToken contract.
type FaaSTokenTransferIterator struct {
	Event *FaaSTokenTransfer // Event containing the contract specifics and raw log

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
func (it *FaaSTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FaaSTokenTransfer)
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
		it.Event = new(FaaSTokenTransfer)
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
func (it *FaaSTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FaaSTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FaaSTokenTransfer represents a Transfer event raised by the FaaSToken contract.
type FaaSTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FaaSToken *FaaSTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FaaSTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FaaSToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FaaSTokenTransferIterator{contract: _FaaSToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FaaSToken *FaaSTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FaaSTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FaaSToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FaaSTokenTransfer)
				if err := _FaaSToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FaaSToken *FaaSTokenFilterer) ParseTransfer(log types.Log) (*FaaSTokenTransfer, error) {
	event := new(FaaSTokenTransfer)
	if err := _FaaSToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
