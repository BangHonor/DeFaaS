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
const FaaSTokenABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_tokenTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"safeAdd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"safeDiv\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"safeMul\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"safeSub\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FaaSTokenBin is the compiled bytecode used for deploying new contracts.
var FaaSTokenBin = "0x60806040523480156200001157600080fd5b506040518060400160405280600581526020017f4641415354000000000000000000000000000000000000000000000000000000815250600090805190602001906200005f929190620001aa565b506040518060400160405280600a81526020017f4661615320546f6b656e0000000000000000000000000000000000000000000081525060019080519060200190620000ad929190620001aa565b506000600260006101000a81548160ff021916908360ff1602179055506a52b7d2dcc80cd2e40000006003819055506000735a86f0cafd4ef3ba4f0344c138afcc84bd1ed2229050600354600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef6003546040518082815260200191505060405180910390a35062000259565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001ed57805160ff19168380011785556200021e565b828001600101855582156200021e579182015b828111156200021d57825182559160200191906001019062000200565b5b5090506200022d919062000231565b5090565b6200025691905b808211156200025257600081600090555060010162000238565b5090565b90565b61111480620002696000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806395d89b4111610097578063cae9ca5111610066578063cae9ca51146104a2578063d05c78da1461059f578063dd62ed3e146105eb578063e6cb901314610663576100f5565b806395d89b4114610321578063a293d1e8146103a4578063a9059cbb146103f0578063b5931f7c14610456576100f5565b806323b872dd116100d357806323b872dd14610201578063313ce5671461028757806370a08231146102ab578063953c151c14610303576100f5565b806306fdde03146100fa578063095ea7b31461017d57806318160ddd146101e3575b600080fd5b6101026106af565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610142578082015181840152602081019050610127565b50505050905090810190601f16801561016f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101c96004803603604081101561019357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061074d565b604051808215151515815260200191505060405180910390f35b6101eb61083f565b6040518082815260200191505060405180910390f35b61026d6004803603606081101561021757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061088a565b604051808215151515815260200191505060405180910390f35b61028f610b1a565b604051808260ff1660ff16815260200191505060405180910390f35b6102ed600480360360208110156102c157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b2d565b6040518082815260200191505060405180910390f35b61030b610b76565b6040518082815260200191505060405180910390f35b610329610b7c565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561036957808201518184015260208101905061034e565b50505050905090810190601f1680156103965780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6103da600480360360408110156103ba57600080fd5b810190808035906020019092919080359060200190929190505050610c1a565b6040518082815260200191505060405180910390f35b61043c6004803603604081101561040657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610c34565b604051808215151515815260200191505060405180910390f35b61048c6004803603604081101561046c57600080fd5b810190808035906020019092919080359060200190929190505050610dbd565b6040518082815260200191505060405180910390f35b610585600480360360608110156104b857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001906401000000008111156104ff57600080fd5b82018360208201111561051157600080fd5b8035906020019184600183028401116401000000008311171561053357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610ddd565b604051808215151515815260200191505060405180910390f35b6105d5600480360360408110156105b557600080fd5b810190808035906020019092919080359060200190929190505050611010565b6040518082815260200191505060405180910390f35b61064d6004803603604081101561060157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061103d565b6040518082815260200191505060405180910390f35b6106996004803603604081101561067957600080fd5b8101908080359060200190929190803590602001909291905050506110c4565b6040518082815260200191505060405180910390f35b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107455780601f1061071a57610100808354040283529160200191610745565b820191906000526020600020905b81548152906001019060200180831161072857829003601f168201915b505050505081565b600081600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a36001905092915050565b6000600460008073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205460035403905090565b60006108d5600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205483610c1a565b600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061099e600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205483610c1a565b600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610a67600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054836110c4565b600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a3600190509392505050565b600260009054906101000a900460ff1681565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60035481565b60008054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c125780601f10610be757610100808354040283529160200191610c12565b820191906000526020600020905b815481529060010190602001808311610bf557829003601f168201915b505050505081565b600082821115610c2957600080fd5b818303905092915050565b6000610c7f600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205483610c1a565b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610d0b600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054836110c4565b600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a36001905092915050565b6000808211610dcb57600080fd5b818381610dd457fe5b04905092915050565b600082600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925856040518082815260200191505060405180910390a38373ffffffffffffffffffffffffffffffffffffffff16638f4ffcb1338530866040518563ffffffff1660e01b8152600401808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610f9e578082015181840152602081019050610f83565b50505050905090810190601f168015610fcb5780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b158015610fed57600080fd5b505af1158015611001573d6000803e3d6000fd5b50505050600190509392505050565b60008183029050600083148061102e57508183828161102b57fe5b04145b61103757600080fd5b92915050565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b60008183019050828110156110d857600080fd5b9291505056fea2646970667358221220a6774395e3bcb142649232faf0b39da550bf47003894a1615d1cbfac963e267064736f6c634300060a0033"

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

// TokenTotalSupply is a free data retrieval call binding the contract method 0x953c151c.
//
// Solidity: function _tokenTotalSupply() constant returns(uint256)
func (_FaaSToken *FaaSTokenCaller) TokenTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FaaSToken.contract.Call(opts, out, "_tokenTotalSupply")
	return *ret0, err
}

// TokenTotalSupply is a free data retrieval call binding the contract method 0x953c151c.
//
// Solidity: function _tokenTotalSupply() constant returns(uint256)
func (_FaaSToken *FaaSTokenSession) TokenTotalSupply() (*big.Int, error) {
	return _FaaSToken.Contract.TokenTotalSupply(&_FaaSToken.CallOpts)
}

// TokenTotalSupply is a free data retrieval call binding the contract method 0x953c151c.
//
// Solidity: function _tokenTotalSupply() constant returns(uint256)
func (_FaaSToken *FaaSTokenCallerSession) TokenTotalSupply() (*big.Int, error) {
	return _FaaSToken.Contract.TokenTotalSupply(&_FaaSToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) constant returns(uint256 remaining)
func (_FaaSToken *FaaSTokenCaller) Allowance(opts *bind.CallOpts, tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FaaSToken.contract.Call(opts, out, "allowance", tokenOwner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) constant returns(uint256 remaining)
func (_FaaSToken *FaaSTokenSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _FaaSToken.Contract.Allowance(&_FaaSToken.CallOpts, tokenOwner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) constant returns(uint256 remaining)
func (_FaaSToken *FaaSTokenCallerSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _FaaSToken.Contract.Allowance(&_FaaSToken.CallOpts, tokenOwner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) constant returns(uint256 balance)
func (_FaaSToken *FaaSTokenCaller) BalanceOf(opts *bind.CallOpts, tokenOwner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FaaSToken.contract.Call(opts, out, "balanceOf", tokenOwner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) constant returns(uint256 balance)
func (_FaaSToken *FaaSTokenSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _FaaSToken.Contract.BalanceOf(&_FaaSToken.CallOpts, tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) constant returns(uint256 balance)
func (_FaaSToken *FaaSTokenCallerSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _FaaSToken.Contract.BalanceOf(&_FaaSToken.CallOpts, tokenOwner)
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

// SafeAdd is a free data retrieval call binding the contract method 0xe6cb9013.
//
// Solidity: function safeAdd(uint256 a, uint256 b) constant returns(uint256 c)
func (_FaaSToken *FaaSTokenCaller) SafeAdd(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FaaSToken.contract.Call(opts, out, "safeAdd", a, b)
	return *ret0, err
}

// SafeAdd is a free data retrieval call binding the contract method 0xe6cb9013.
//
// Solidity: function safeAdd(uint256 a, uint256 b) constant returns(uint256 c)
func (_FaaSToken *FaaSTokenSession) SafeAdd(a *big.Int, b *big.Int) (*big.Int, error) {
	return _FaaSToken.Contract.SafeAdd(&_FaaSToken.CallOpts, a, b)
}

// SafeAdd is a free data retrieval call binding the contract method 0xe6cb9013.
//
// Solidity: function safeAdd(uint256 a, uint256 b) constant returns(uint256 c)
func (_FaaSToken *FaaSTokenCallerSession) SafeAdd(a *big.Int, b *big.Int) (*big.Int, error) {
	return _FaaSToken.Contract.SafeAdd(&_FaaSToken.CallOpts, a, b)
}

// SafeDiv is a free data retrieval call binding the contract method 0xb5931f7c.
//
// Solidity: function safeDiv(uint256 a, uint256 b) constant returns(uint256 c)
func (_FaaSToken *FaaSTokenCaller) SafeDiv(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FaaSToken.contract.Call(opts, out, "safeDiv", a, b)
	return *ret0, err
}

// SafeDiv is a free data retrieval call binding the contract method 0xb5931f7c.
//
// Solidity: function safeDiv(uint256 a, uint256 b) constant returns(uint256 c)
func (_FaaSToken *FaaSTokenSession) SafeDiv(a *big.Int, b *big.Int) (*big.Int, error) {
	return _FaaSToken.Contract.SafeDiv(&_FaaSToken.CallOpts, a, b)
}

// SafeDiv is a free data retrieval call binding the contract method 0xb5931f7c.
//
// Solidity: function safeDiv(uint256 a, uint256 b) constant returns(uint256 c)
func (_FaaSToken *FaaSTokenCallerSession) SafeDiv(a *big.Int, b *big.Int) (*big.Int, error) {
	return _FaaSToken.Contract.SafeDiv(&_FaaSToken.CallOpts, a, b)
}

// SafeMul is a free data retrieval call binding the contract method 0xd05c78da.
//
// Solidity: function safeMul(uint256 a, uint256 b) constant returns(uint256 c)
func (_FaaSToken *FaaSTokenCaller) SafeMul(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FaaSToken.contract.Call(opts, out, "safeMul", a, b)
	return *ret0, err
}

// SafeMul is a free data retrieval call binding the contract method 0xd05c78da.
//
// Solidity: function safeMul(uint256 a, uint256 b) constant returns(uint256 c)
func (_FaaSToken *FaaSTokenSession) SafeMul(a *big.Int, b *big.Int) (*big.Int, error) {
	return _FaaSToken.Contract.SafeMul(&_FaaSToken.CallOpts, a, b)
}

// SafeMul is a free data retrieval call binding the contract method 0xd05c78da.
//
// Solidity: function safeMul(uint256 a, uint256 b) constant returns(uint256 c)
func (_FaaSToken *FaaSTokenCallerSession) SafeMul(a *big.Int, b *big.Int) (*big.Int, error) {
	return _FaaSToken.Contract.SafeMul(&_FaaSToken.CallOpts, a, b)
}

// SafeSub is a free data retrieval call binding the contract method 0xa293d1e8.
//
// Solidity: function safeSub(uint256 a, uint256 b) constant returns(uint256 c)
func (_FaaSToken *FaaSTokenCaller) SafeSub(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FaaSToken.contract.Call(opts, out, "safeSub", a, b)
	return *ret0, err
}

// SafeSub is a free data retrieval call binding the contract method 0xa293d1e8.
//
// Solidity: function safeSub(uint256 a, uint256 b) constant returns(uint256 c)
func (_FaaSToken *FaaSTokenSession) SafeSub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _FaaSToken.Contract.SafeSub(&_FaaSToken.CallOpts, a, b)
}

// SafeSub is a free data retrieval call binding the contract method 0xa293d1e8.
//
// Solidity: function safeSub(uint256 a, uint256 b) constant returns(uint256 c)
func (_FaaSToken *FaaSTokenCallerSession) SafeSub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _FaaSToken.Contract.SafeSub(&_FaaSToken.CallOpts, a, b)
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
// Solidity: function approve(address spender, uint256 tokens) returns(bool success)
func (_FaaSToken *FaaSTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, tokens *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.contract.Transact(opts, "approve", spender, tokens)
}

func (_FaaSToken *FaaSTokenTransactor) AsyncApprove(handler func(*types.Receipt, error), opts *bind.TransactOpts, spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _FaaSToken.contract.AsyncTransact(opts, handler, "approve", spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 tokens) returns(bool success)
func (_FaaSToken *FaaSTokenSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.Contract.Approve(&_FaaSToken.TransactOpts, spender, tokens)
}

func (_FaaSToken *FaaSTokenSession) AsyncApprove(handler func(*types.Receipt, error), spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.AsyncApprove(handler, &_FaaSToken.TransactOpts, spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 tokens) returns(bool success)
func (_FaaSToken *FaaSTokenTransactorSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.Contract.Approve(&_FaaSToken.TransactOpts, spender, tokens)
}

func (_FaaSToken *FaaSTokenTransactorSession) AsyncApprove(handler func(*types.Receipt, error), spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.AsyncApprove(handler, &_FaaSToken.TransactOpts, spender, tokens)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 tokens, bytes data) returns(bool success)
func (_FaaSToken *FaaSTokenTransactor) ApproveAndCall(opts *bind.TransactOpts, spender common.Address, tokens *big.Int, data []byte) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.contract.Transact(opts, "approveAndCall", spender, tokens, data)
}

func (_FaaSToken *FaaSTokenTransactor) AsyncApproveAndCall(handler func(*types.Receipt, error), opts *bind.TransactOpts, spender common.Address, tokens *big.Int, data []byte) (*types.Transaction, error) {
	return _FaaSToken.contract.AsyncTransact(opts, handler, "approveAndCall", spender, tokens, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 tokens, bytes data) returns(bool success)
func (_FaaSToken *FaaSTokenSession) ApproveAndCall(spender common.Address, tokens *big.Int, data []byte) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.Contract.ApproveAndCall(&_FaaSToken.TransactOpts, spender, tokens, data)
}

func (_FaaSToken *FaaSTokenSession) AsyncApproveAndCall(handler func(*types.Receipt, error), spender common.Address, tokens *big.Int, data []byte) (*types.Transaction, error) {
	return _FaaSToken.Contract.AsyncApproveAndCall(handler, &_FaaSToken.TransactOpts, spender, tokens, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 tokens, bytes data) returns(bool success)
func (_FaaSToken *FaaSTokenTransactorSession) ApproveAndCall(spender common.Address, tokens *big.Int, data []byte) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.Contract.ApproveAndCall(&_FaaSToken.TransactOpts, spender, tokens, data)
}

func (_FaaSToken *FaaSTokenTransactorSession) AsyncApproveAndCall(handler func(*types.Receipt, error), spender common.Address, tokens *big.Int, data []byte) (*types.Transaction, error) {
	return _FaaSToken.Contract.AsyncApproveAndCall(handler, &_FaaSToken.TransactOpts, spender, tokens, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 tokens) returns(bool success)
func (_FaaSToken *FaaSTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, tokens *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.contract.Transact(opts, "transfer", to, tokens)
}

func (_FaaSToken *FaaSTokenTransactor) AsyncTransfer(handler func(*types.Receipt, error), opts *bind.TransactOpts, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _FaaSToken.contract.AsyncTransact(opts, handler, "transfer", to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 tokens) returns(bool success)
func (_FaaSToken *FaaSTokenSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.Contract.Transfer(&_FaaSToken.TransactOpts, to, tokens)
}

func (_FaaSToken *FaaSTokenSession) AsyncTransfer(handler func(*types.Receipt, error), to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.AsyncTransfer(handler, &_FaaSToken.TransactOpts, to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 tokens) returns(bool success)
func (_FaaSToken *FaaSTokenTransactorSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.Contract.Transfer(&_FaaSToken.TransactOpts, to, tokens)
}

func (_FaaSToken *FaaSTokenTransactorSession) AsyncTransfer(handler func(*types.Receipt, error), to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.AsyncTransfer(handler, &_FaaSToken.TransactOpts, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokens) returns(bool success)
func (_FaaSToken *FaaSTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.contract.Transact(opts, "transferFrom", from, to, tokens)
}

func (_FaaSToken *FaaSTokenTransactor) AsyncTransferFrom(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _FaaSToken.contract.AsyncTransact(opts, handler, "transferFrom", from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokens) returns(bool success)
func (_FaaSToken *FaaSTokenSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.Contract.TransferFrom(&_FaaSToken.TransactOpts, from, to, tokens)
}

func (_FaaSToken *FaaSTokenSession) AsyncTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.AsyncTransferFrom(handler, &_FaaSToken.TransactOpts, from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokens) returns(bool success)
func (_FaaSToken *FaaSTokenTransactorSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _FaaSToken.Contract.TransferFrom(&_FaaSToken.TransactOpts, from, to, tokens)
}

func (_FaaSToken *FaaSTokenTransactorSession) AsyncTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.AsyncTransferFrom(handler, &_FaaSToken.TransactOpts, from, to, tokens)
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
