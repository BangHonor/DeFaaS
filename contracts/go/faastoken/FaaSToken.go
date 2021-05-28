// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package faastoken

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

// FaaSTokenABI is the input ABI used to generate the binding from.
const FaaSTokenABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FaaSTokenBin is the compiled bytecode used for deploying new contracts.
var FaaSTokenBin = "0x60806040523480156200001157600080fd5b506040518060400160405280600981526020017f46616153546f6b656e00000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f4653540000000000000000000000000000000000000000000000000000000000815250816003908051906020019062000096929190620000d4565b508060049080519060200190620000af929190620000d4565b506012600560006101000a81548160ff021916908360ff1602179055505050620001e9565b828054620000e29062000184565b90600052602060002090601f01602090048101928262000106576000855562000152565b82601f106200012157805160ff191683800117855562000152565b8280016001018555821562000152579182015b828111156200015157825182559160200191906001019062000134565b5b50905062000161919062000165565b5090565b5b808211156200018057600081600090555060010162000166565b5090565b600060028204905060018216806200019d57607f821691505b60208210811415620001b457620001b3620001ba565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b61164b80620001f96000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063449a52f81161008c578063a0712d6811610066578063a0712d6814610228578063a457c2d714610244578063a9059cbb14610274578063dd62ed3e146102a4576100cf565b8063449a52f8146101be57806370a08231146101da57806395d89b411461020a576100cf565b806306fdde03146100d4578063095ea7b3146100f257806318160ddd1461012257806323b872dd14610140578063313ce56714610170578063395093511461018e575b600080fd5b6100dc6102d4565b6040516100e9919061106e565b60405180910390f35b61010c60048036038101906101079190610e70565b610366565b6040516101199190611053565b60405180910390f35b61012a610384565b6040516101379190611190565b60405180910390f35b61015a60048036038101906101559190610e21565b61038e565b6040516101679190611053565b60405180910390f35b61017861048f565b60405161018591906111ab565b60405180910390f35b6101a860048036038101906101a39190610e70565b6104a6565b6040516101b59190611053565b60405180910390f35b6101d860048036038101906101d39190610e70565b610552565b005b6101f460048036038101906101ef9190610dbc565b610560565b6040516102019190611190565b60405180910390f35b6102126105a8565b60405161021f919061106e565b60405180910390f35b610242600480360381019061023d9190610eac565b61063a565b005b61025e60048036038101906102599190610e70565b61064e565b60405161026b9190611053565b60405180910390f35b61028e60048036038101906102899190610e70565b610742565b60405161029b9190611053565b60405180910390f35b6102be60048036038101906102b99190610de5565b610760565b6040516102cb9190611190565b60405180910390f35b6060600380546102e3906112f4565b80601f016020809104026020016040519081016040528092919081815260200182805461030f906112f4565b801561035c5780601f106103315761010080835404028352916020019161035c565b820191906000526020600020905b81548152906001019060200180831161033f57829003601f168201915b5050505050905090565b600061037a6103736107e7565b84846107ef565b6001905092915050565b6000600254905090565b600061039b8484846109ba565b6000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006103e66107e7565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905082811015610466576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045d906110f0565b60405180910390fd5b610483856104726107e7565b858461047e9190611238565b6107ef565b60019150509392505050565b6000600560009054906101000a900460ff16905090565b60006105486104b36107e7565b8484600160006104c16107e7565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461054391906111e2565b6107ef565b6001905092915050565b61055c8282610c39565b5050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6060600480546105b7906112f4565b80601f01602080910402602001604051908101604052809291908181526020018280546105e3906112f4565b80156106305780601f1061060557610100808354040283529160200191610630565b820191906000526020600020905b81548152906001019060200180831161061357829003601f168201915b5050505050905090565b61064b6106456107e7565b82610c39565b50565b6000806001600061065d6107e7565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508281101561071a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161071190611150565b60405180910390fd5b6107376107256107e7565b8585846107329190611238565b6107ef565b600191505092915050565b600061075661074f6107e7565b84846109ba565b6001905092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561085f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161085690611130565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156108cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108c6906110b0565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516109ad9190611190565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610a2a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2190611110565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610a9a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a9190611090565b60405180910390fd5b610aa5838383610d8d565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610b2b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b22906110d0565b60405180910390fd5b8181610b379190611238565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610bc791906111e2565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610c2b9190611190565b60405180910390a350505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610ca9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ca090611170565b60405180910390fd5b610cb560008383610d8d565b8060026000828254610cc791906111e2565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610d1c91906111e2565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610d819190611190565b60405180910390a35050565b505050565b600081359050610da1816115e7565b92915050565b600081359050610db6816115fe565b92915050565b600060208284031215610dce57600080fd5b6000610ddc84828501610d92565b91505092915050565b60008060408385031215610df857600080fd5b6000610e0685828601610d92565b9250506020610e1785828601610d92565b9150509250929050565b600080600060608486031215610e3657600080fd5b6000610e4486828701610d92565b9350506020610e5586828701610d92565b9250506040610e6686828701610da7565b9150509250925092565b60008060408385031215610e8357600080fd5b6000610e9185828601610d92565b9250506020610ea285828601610da7565b9150509250929050565b600060208284031215610ebe57600080fd5b6000610ecc84828501610da7565b91505092915050565b610ede8161127e565b82525050565b6000610eef826111c6565b610ef981856111d1565b9350610f098185602086016112c1565b610f1281611384565b840191505092915050565b6000610f2a6023836111d1565b9150610f3582611395565b604082019050919050565b6000610f4d6022836111d1565b9150610f58826113e4565b604082019050919050565b6000610f706026836111d1565b9150610f7b82611433565b604082019050919050565b6000610f936028836111d1565b9150610f9e82611482565b604082019050919050565b6000610fb66025836111d1565b9150610fc1826114d1565b604082019050919050565b6000610fd96024836111d1565b9150610fe482611520565b604082019050919050565b6000610ffc6025836111d1565b91506110078261156f565b604082019050919050565b600061101f601f836111d1565b915061102a826115be565b602082019050919050565b61103e816112aa565b82525050565b61104d816112b4565b82525050565b60006020820190506110686000830184610ed5565b92915050565b600060208201905081810360008301526110888184610ee4565b905092915050565b600060208201905081810360008301526110a981610f1d565b9050919050565b600060208201905081810360008301526110c981610f40565b9050919050565b600060208201905081810360008301526110e981610f63565b9050919050565b6000602082019050818103600083015261110981610f86565b9050919050565b6000602082019050818103600083015261112981610fa9565b9050919050565b6000602082019050818103600083015261114981610fcc565b9050919050565b6000602082019050818103600083015261116981610fef565b9050919050565b6000602082019050818103600083015261118981611012565b9050919050565b60006020820190506111a56000830184611035565b92915050565b60006020820190506111c06000830184611044565b92915050565b600081519050919050565b600082825260208201905092915050565b60006111ed826112aa565b91506111f8836112aa565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561122d5761122c611326565b5b828201905092915050565b6000611243826112aa565b915061124e836112aa565b92508282101561126157611260611326565b5b828203905092915050565b60006112778261128a565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b838110156112df5780820151818401526020810190506112c4565b838111156112ee576000848401525b50505050565b6000600282049050600182168061130c57607f821691505b602082108114156113205761131f611355565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000601f19601f8301169050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206160008201527f6c6c6f77616e6365000000000000000000000000000000000000000000000000602082015250565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b6115f08161126c565b81146115fb57600080fd5b50565b611607816112aa565b811461161257600080fd5b5056fea26469706673582212200650552aec5d2265e03c420f43529a56a8fa7b408a90df71ce7633eb60c9d43264736f6c63430008010033"

// DeployFaaSToken deploys a new Ethereum contract, binding an instance of FaaSToken to it.
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

// FaaSToken is an auto generated Go binding around an Ethereum contract.
type FaaSToken struct {
	FaaSTokenCaller     // Read-only binding to the contract
	FaaSTokenTransactor // Write-only binding to the contract
	FaaSTokenFilterer   // Log filterer for contract events
}

// FaaSTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type FaaSTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaaSTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FaaSTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaaSTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FaaSTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaaSTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FaaSTokenSession struct {
	Contract     *FaaSToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FaaSTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FaaSTokenCallerSession struct {
	Contract *FaaSTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// FaaSTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FaaSTokenTransactorSession struct {
	Contract     *FaaSTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FaaSTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type FaaSTokenRaw struct {
	Contract *FaaSToken // Generic contract binding to access the raw methods on
}

// FaaSTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FaaSTokenCallerRaw struct {
	Contract *FaaSTokenCaller // Generic read-only contract binding to access the raw methods on
}

// FaaSTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
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
func (_FaaSToken *FaaSTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FaaSToken.Contract.FaaSTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FaaSToken *FaaSTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FaaSToken.Contract.FaaSTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FaaSToken *FaaSTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FaaSToken.Contract.FaaSTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FaaSToken *FaaSTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FaaSToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FaaSToken *FaaSTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FaaSToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FaaSToken *FaaSTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FaaSToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_FaaSToken *FaaSTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FaaSToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_FaaSToken *FaaSTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _FaaSToken.Contract.Allowance(&_FaaSToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_FaaSToken *FaaSTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _FaaSToken.Contract.Allowance(&_FaaSToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_FaaSToken *FaaSTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FaaSToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_FaaSToken *FaaSTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _FaaSToken.Contract.BalanceOf(&_FaaSToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_FaaSToken *FaaSTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _FaaSToken.Contract.BalanceOf(&_FaaSToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FaaSToken *FaaSTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FaaSToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FaaSToken *FaaSTokenSession) Decimals() (uint8, error) {
	return _FaaSToken.Contract.Decimals(&_FaaSToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FaaSToken *FaaSTokenCallerSession) Decimals() (uint8, error) {
	return _FaaSToken.Contract.Decimals(&_FaaSToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FaaSToken *FaaSTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FaaSToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FaaSToken *FaaSTokenSession) Name() (string, error) {
	return _FaaSToken.Contract.Name(&_FaaSToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FaaSToken *FaaSTokenCallerSession) Name() (string, error) {
	return _FaaSToken.Contract.Name(&_FaaSToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FaaSToken *FaaSTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FaaSToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FaaSToken *FaaSTokenSession) Symbol() (string, error) {
	return _FaaSToken.Contract.Symbol(&_FaaSToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FaaSToken *FaaSTokenCallerSession) Symbol() (string, error) {
	return _FaaSToken.Contract.Symbol(&_FaaSToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FaaSToken *FaaSTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FaaSToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FaaSToken *FaaSTokenSession) TotalSupply() (*big.Int, error) {
	return _FaaSToken.Contract.TotalSupply(&_FaaSToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FaaSToken *FaaSTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _FaaSToken.Contract.TotalSupply(&_FaaSToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_FaaSToken *FaaSTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_FaaSToken *FaaSTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.Approve(&_FaaSToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_FaaSToken *FaaSTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.Approve(&_FaaSToken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_FaaSToken *FaaSTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _FaaSToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_FaaSToken *FaaSTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.DecreaseAllowance(&_FaaSToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_FaaSToken *FaaSTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.DecreaseAllowance(&_FaaSToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_FaaSToken *FaaSTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _FaaSToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_FaaSToken *FaaSTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.IncreaseAllowance(&_FaaSToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_FaaSToken *FaaSTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.IncreaseAllowance(&_FaaSToken.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_FaaSToken *FaaSTokenTransactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_FaaSToken *FaaSTokenSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.Mint(&_FaaSToken.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_FaaSToken *FaaSTokenTransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.Mint(&_FaaSToken.TransactOpts, amount)
}

// MintTo is a paid mutator transaction binding the contract method 0x449a52f8.
//
// Solidity: function mintTo(address account, uint256 amount) returns()
func (_FaaSToken *FaaSTokenTransactor) MintTo(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.contract.Transact(opts, "mintTo", account, amount)
}

// MintTo is a paid mutator transaction binding the contract method 0x449a52f8.
//
// Solidity: function mintTo(address account, uint256 amount) returns()
func (_FaaSToken *FaaSTokenSession) MintTo(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.MintTo(&_FaaSToken.TransactOpts, account, amount)
}

// MintTo is a paid mutator transaction binding the contract method 0x449a52f8.
//
// Solidity: function mintTo(address account, uint256 amount) returns()
func (_FaaSToken *FaaSTokenTransactorSession) MintTo(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.MintTo(&_FaaSToken.TransactOpts, account, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_FaaSToken *FaaSTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_FaaSToken *FaaSTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.Transfer(&_FaaSToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_FaaSToken *FaaSTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.Transfer(&_FaaSToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_FaaSToken *FaaSTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_FaaSToken *FaaSTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.TransferFrom(&_FaaSToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_FaaSToken *FaaSTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FaaSToken.Contract.TransferFrom(&_FaaSToken.TransactOpts, sender, recipient, amount)
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
	event.Raw = log
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
	event.Raw = log
	return event, nil
}
