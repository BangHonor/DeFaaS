// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package providerpool

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

// ProviderPoolABI is the input ABI used to generate the binding from.
const ProviderPoolABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContractAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"}],\"name\":\"getProviderDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"}],\"name\":\"getProviderReputation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStdProviderDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"}],\"name\":\"isProviderQualified\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"}],\"name\":\"isProviderRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerLogin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerLogout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenContract\",\"outputs\":[{\"internalType\":\"contractFaaSToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ProviderPoolBin is the compiled bytecode used for deploying new contracts.
var ProviderPoolBin = "0x608060405234801561001057600080fd5b50604051610baa380380610baa8339818101604052602081101561003357600080fd5b810190808051906020019092919050505080806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505060646001819055506005600281905550600560038190555050610afc806100ae6000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806361bb20f01161005b57806361bb20f01461015b578063729d98e9146101b3578063884eacb81461020b578063b82964711461021557610088565b80631c0d3ebd1461008d5780631ead914e146100ab5780632bf355701461010757806355a373d614610111575b600080fd5b610095610271565b6040518082815260200191505060405180910390f35b6100ed600480360360208110156100c157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061027b565b604051808215151515815260200191505060405180910390f35b61010f6102c9565b005b61011961053e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61019d6004803603602081101561017157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610563565b6040518082815260200191505060405180910390f35b6101f5600480360360208110156101c957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610616565b6040518082815260200191505060405180910390f35b6102136106c9565b005b6102576004803603602081101561022b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506109a9565b604051808215151515815260200191505060405180910390f35b6000600154905090565b600080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154119050919050565b33600015156102d78261027b565b15151461032f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526027815260200180610a236027913960400191505060405180910390fd5b600154600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010181905550600254600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000181905550600115156000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd333061040c610271565b6040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b1580156104a857600080fd5b505af11580156104bc573d6000803e3d6000fd5b505050506040513d60208110156104d257600080fd5b810190808051906020019092919050505015151461053b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260288152602001806109fb6028913960400191505060405180910390fd5b50565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600081600115156105738261027b565b1515146105cb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602b815260200180610a6f602b913960400191505060405180910390fd5b600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010154915050919050565b600081600115156106268261027b565b15151461067e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602b815260200180610a6f602b913960400191505060405180910390fd5b600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154915050919050565b33600115156106d78261027b565b15151461072f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602b815260200180610a6f602b913960400191505060405180910390fd5b33610739816109a9565b61078e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180610a4a6025913960400191505060405180910390fd5b6000600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015490506000600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055506000600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000181905550600115156000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33846040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561091157600080fd5b505af1158015610925573d6000803e3d6000fd5b505050506040513d602081101561093b57600080fd5b81019080805190602001909291905050501515146109a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602d815260200180610a9a602d913960400191505060405180910390fd5b505050565b6000600354600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001541015905091905056fe4d61726b65743a206661696c656420746f207061792061207265676973746572206465706f7369744d61726b65743a20746865206164647265737320686164206265656e20726567697374657265644d61726b65743a207468652070726f7669646572206973206e6f74207175616c69666965644d61726b65743a20746865206164647265737320686164206e6f74206265656e20726567697374657265644d61726b65743a206661696c656420746f20726566756e64207468652070726f7669646572206465706f736974a26469706673582212207f4301c5a6cafbd86b0cbb52e07e80014d6af9a879e478633d7f10d1fffd052964736f6c634300060a0033"

// DeployProviderPool deploys a new contract, binding an instance of ProviderPool to it.
func DeployProviderPool(auth *bind.TransactOpts, backend bind.ContractBackend, _tokenContractAddress common.Address) (common.Address, *types.Transaction, *ProviderPool, error) {
	parsed, err := abi.JSON(strings.NewReader(ProviderPoolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProviderPoolBin), backend, _tokenContractAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProviderPool{ProviderPoolCaller: ProviderPoolCaller{contract: contract}, ProviderPoolTransactor: ProviderPoolTransactor{contract: contract}, ProviderPoolFilterer: ProviderPoolFilterer{contract: contract}}, nil
}

func AsyncDeployProviderPool(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend, _tokenContractAddress common.Address) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(ProviderPoolABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(ProviderPoolBin), backend, _tokenContractAddress)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// ProviderPool is an auto generated Go binding around a Solidity contract.
type ProviderPool struct {
	ProviderPoolCaller     // Read-only binding to the contract
	ProviderPoolTransactor // Write-only binding to the contract
	ProviderPoolFilterer   // Log filterer for contract events
}

// ProviderPoolCaller is an auto generated read-only Go binding around a Solidity contract.
type ProviderPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProviderPoolTransactor is an auto generated write-only Go binding around a Solidity contract.
type ProviderPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProviderPoolFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type ProviderPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProviderPoolSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type ProviderPoolSession struct {
	Contract     *ProviderPool     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProviderPoolCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type ProviderPoolCallerSession struct {
	Contract *ProviderPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ProviderPoolTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type ProviderPoolTransactorSession struct {
	Contract     *ProviderPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ProviderPoolRaw is an auto generated low-level Go binding around a Solidity contract.
type ProviderPoolRaw struct {
	Contract *ProviderPool // Generic contract binding to access the raw methods on
}

// ProviderPoolCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type ProviderPoolCallerRaw struct {
	Contract *ProviderPoolCaller // Generic read-only contract binding to access the raw methods on
}

// ProviderPoolTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type ProviderPoolTransactorRaw struct {
	Contract *ProviderPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProviderPool creates a new instance of ProviderPool, bound to a specific deployed contract.
func NewProviderPool(address common.Address, backend bind.ContractBackend) (*ProviderPool, error) {
	contract, err := bindProviderPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProviderPool{ProviderPoolCaller: ProviderPoolCaller{contract: contract}, ProviderPoolTransactor: ProviderPoolTransactor{contract: contract}, ProviderPoolFilterer: ProviderPoolFilterer{contract: contract}}, nil
}

// NewProviderPoolCaller creates a new read-only instance of ProviderPool, bound to a specific deployed contract.
func NewProviderPoolCaller(address common.Address, caller bind.ContractCaller) (*ProviderPoolCaller, error) {
	contract, err := bindProviderPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProviderPoolCaller{contract: contract}, nil
}

// NewProviderPoolTransactor creates a new write-only instance of ProviderPool, bound to a specific deployed contract.
func NewProviderPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*ProviderPoolTransactor, error) {
	contract, err := bindProviderPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProviderPoolTransactor{contract: contract}, nil
}

// NewProviderPoolFilterer creates a new log filterer instance of ProviderPool, bound to a specific deployed contract.
func NewProviderPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*ProviderPoolFilterer, error) {
	contract, err := bindProviderPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProviderPoolFilterer{contract: contract}, nil
}

// bindProviderPool binds a generic wrapper to an already deployed contract.
func bindProviderPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProviderPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProviderPool *ProviderPoolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProviderPool.Contract.ProviderPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProviderPool *ProviderPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ProviderPool.Contract.ProviderPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProviderPool *ProviderPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _ProviderPool.Contract.ProviderPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProviderPool *ProviderPoolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProviderPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProviderPool *ProviderPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ProviderPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProviderPool *ProviderPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _ProviderPool.Contract.contract.Transact(opts, method, params...)
}

// GetProviderDeposit is a free data retrieval call binding the contract method 0x61bb20f0.
//
// Solidity: function getProviderDeposit(address _provider) constant returns(uint256)
func (_ProviderPool *ProviderPoolCaller) GetProviderDeposit(opts *bind.CallOpts, _provider common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ProviderPool.contract.Call(opts, out, "getProviderDeposit", _provider)
	return *ret0, err
}

// GetProviderDeposit is a free data retrieval call binding the contract method 0x61bb20f0.
//
// Solidity: function getProviderDeposit(address _provider) constant returns(uint256)
func (_ProviderPool *ProviderPoolSession) GetProviderDeposit(_provider common.Address) (*big.Int, error) {
	return _ProviderPool.Contract.GetProviderDeposit(&_ProviderPool.CallOpts, _provider)
}

// GetProviderDeposit is a free data retrieval call binding the contract method 0x61bb20f0.
//
// Solidity: function getProviderDeposit(address _provider) constant returns(uint256)
func (_ProviderPool *ProviderPoolCallerSession) GetProviderDeposit(_provider common.Address) (*big.Int, error) {
	return _ProviderPool.Contract.GetProviderDeposit(&_ProviderPool.CallOpts, _provider)
}

// GetProviderReputation is a free data retrieval call binding the contract method 0x729d98e9.
//
// Solidity: function getProviderReputation(address _provider) constant returns(uint256)
func (_ProviderPool *ProviderPoolCaller) GetProviderReputation(opts *bind.CallOpts, _provider common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ProviderPool.contract.Call(opts, out, "getProviderReputation", _provider)
	return *ret0, err
}

// GetProviderReputation is a free data retrieval call binding the contract method 0x729d98e9.
//
// Solidity: function getProviderReputation(address _provider) constant returns(uint256)
func (_ProviderPool *ProviderPoolSession) GetProviderReputation(_provider common.Address) (*big.Int, error) {
	return _ProviderPool.Contract.GetProviderReputation(&_ProviderPool.CallOpts, _provider)
}

// GetProviderReputation is a free data retrieval call binding the contract method 0x729d98e9.
//
// Solidity: function getProviderReputation(address _provider) constant returns(uint256)
func (_ProviderPool *ProviderPoolCallerSession) GetProviderReputation(_provider common.Address) (*big.Int, error) {
	return _ProviderPool.Contract.GetProviderReputation(&_ProviderPool.CallOpts, _provider)
}

// GetStdProviderDeposit is a free data retrieval call binding the contract method 0x1c0d3ebd.
//
// Solidity: function getStdProviderDeposit() constant returns(uint256)
func (_ProviderPool *ProviderPoolCaller) GetStdProviderDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ProviderPool.contract.Call(opts, out, "getStdProviderDeposit")
	return *ret0, err
}

// GetStdProviderDeposit is a free data retrieval call binding the contract method 0x1c0d3ebd.
//
// Solidity: function getStdProviderDeposit() constant returns(uint256)
func (_ProviderPool *ProviderPoolSession) GetStdProviderDeposit() (*big.Int, error) {
	return _ProviderPool.Contract.GetStdProviderDeposit(&_ProviderPool.CallOpts)
}

// GetStdProviderDeposit is a free data retrieval call binding the contract method 0x1c0d3ebd.
//
// Solidity: function getStdProviderDeposit() constant returns(uint256)
func (_ProviderPool *ProviderPoolCallerSession) GetStdProviderDeposit() (*big.Int, error) {
	return _ProviderPool.Contract.GetStdProviderDeposit(&_ProviderPool.CallOpts)
}

// IsProviderQualified is a free data retrieval call binding the contract method 0xb8296471.
//
// Solidity: function isProviderQualified(address _provider) constant returns(bool)
func (_ProviderPool *ProviderPoolCaller) IsProviderQualified(opts *bind.CallOpts, _provider common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ProviderPool.contract.Call(opts, out, "isProviderQualified", _provider)
	return *ret0, err
}

// IsProviderQualified is a free data retrieval call binding the contract method 0xb8296471.
//
// Solidity: function isProviderQualified(address _provider) constant returns(bool)
func (_ProviderPool *ProviderPoolSession) IsProviderQualified(_provider common.Address) (bool, error) {
	return _ProviderPool.Contract.IsProviderQualified(&_ProviderPool.CallOpts, _provider)
}

// IsProviderQualified is a free data retrieval call binding the contract method 0xb8296471.
//
// Solidity: function isProviderQualified(address _provider) constant returns(bool)
func (_ProviderPool *ProviderPoolCallerSession) IsProviderQualified(_provider common.Address) (bool, error) {
	return _ProviderPool.Contract.IsProviderQualified(&_ProviderPool.CallOpts, _provider)
}

// IsProviderRegistered is a free data retrieval call binding the contract method 0x1ead914e.
//
// Solidity: function isProviderRegistered(address _provider) constant returns(bool)
func (_ProviderPool *ProviderPoolCaller) IsProviderRegistered(opts *bind.CallOpts, _provider common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ProviderPool.contract.Call(opts, out, "isProviderRegistered", _provider)
	return *ret0, err
}

// IsProviderRegistered is a free data retrieval call binding the contract method 0x1ead914e.
//
// Solidity: function isProviderRegistered(address _provider) constant returns(bool)
func (_ProviderPool *ProviderPoolSession) IsProviderRegistered(_provider common.Address) (bool, error) {
	return _ProviderPool.Contract.IsProviderRegistered(&_ProviderPool.CallOpts, _provider)
}

// IsProviderRegistered is a free data retrieval call binding the contract method 0x1ead914e.
//
// Solidity: function isProviderRegistered(address _provider) constant returns(bool)
func (_ProviderPool *ProviderPoolCallerSession) IsProviderRegistered(_provider common.Address) (bool, error) {
	return _ProviderPool.Contract.IsProviderRegistered(&_ProviderPool.CallOpts, _provider)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_ProviderPool *ProviderPoolCaller) TokenContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ProviderPool.contract.Call(opts, out, "tokenContract")
	return *ret0, err
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_ProviderPool *ProviderPoolSession) TokenContract() (common.Address, error) {
	return _ProviderPool.Contract.TokenContract(&_ProviderPool.CallOpts)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_ProviderPool *ProviderPoolCallerSession) TokenContract() (common.Address, error) {
	return _ProviderPool.Contract.TokenContract(&_ProviderPool.CallOpts)
}

// ProviderLogin is a paid mutator transaction binding the contract method 0x2bf35570.
//
// Solidity: function providerLogin() returns()
func (_ProviderPool *ProviderPoolTransactor) ProviderLogin(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ProviderPool.contract.Transact(opts, "providerLogin")
}

func (_ProviderPool *ProviderPoolTransactor) AsyncProviderLogin(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProviderPool.contract.AsyncTransact(opts, handler, "providerLogin")
}

// ProviderLogin is a paid mutator transaction binding the contract method 0x2bf35570.
//
// Solidity: function providerLogin() returns()
func (_ProviderPool *ProviderPoolSession) ProviderLogin() (*types.Transaction, *types.Receipt, error) {
	return _ProviderPool.Contract.ProviderLogin(&_ProviderPool.TransactOpts)
}

func (_ProviderPool *ProviderPoolSession) AsyncProviderLogin(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _ProviderPool.Contract.AsyncProviderLogin(handler, &_ProviderPool.TransactOpts)
}

// ProviderLogin is a paid mutator transaction binding the contract method 0x2bf35570.
//
// Solidity: function providerLogin() returns()
func (_ProviderPool *ProviderPoolTransactorSession) ProviderLogin() (*types.Transaction, *types.Receipt, error) {
	return _ProviderPool.Contract.ProviderLogin(&_ProviderPool.TransactOpts)
}

func (_ProviderPool *ProviderPoolTransactorSession) AsyncProviderLogin(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _ProviderPool.Contract.AsyncProviderLogin(handler, &_ProviderPool.TransactOpts)
}

// ProviderLogout is a paid mutator transaction binding the contract method 0x884eacb8.
//
// Solidity: function providerLogout() returns()
func (_ProviderPool *ProviderPoolTransactor) ProviderLogout(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ProviderPool.contract.Transact(opts, "providerLogout")
}

func (_ProviderPool *ProviderPoolTransactor) AsyncProviderLogout(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProviderPool.contract.AsyncTransact(opts, handler, "providerLogout")
}

// ProviderLogout is a paid mutator transaction binding the contract method 0x884eacb8.
//
// Solidity: function providerLogout() returns()
func (_ProviderPool *ProviderPoolSession) ProviderLogout() (*types.Transaction, *types.Receipt, error) {
	return _ProviderPool.Contract.ProviderLogout(&_ProviderPool.TransactOpts)
}

func (_ProviderPool *ProviderPoolSession) AsyncProviderLogout(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _ProviderPool.Contract.AsyncProviderLogout(handler, &_ProviderPool.TransactOpts)
}

// ProviderLogout is a paid mutator transaction binding the contract method 0x884eacb8.
//
// Solidity: function providerLogout() returns()
func (_ProviderPool *ProviderPoolTransactorSession) ProviderLogout() (*types.Transaction, *types.Receipt, error) {
	return _ProviderPool.Contract.ProviderLogout(&_ProviderPool.TransactOpts)
}

func (_ProviderPool *ProviderPoolTransactorSession) AsyncProviderLogout(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _ProviderPool.Contract.AsyncProviderLogout(handler, &_ProviderPool.TransactOpts)
}
