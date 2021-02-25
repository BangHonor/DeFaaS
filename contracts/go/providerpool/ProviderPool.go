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
const ProviderPoolABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getProviderDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"getProviderReputation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProviderReputationQualified\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"isProviderQualified\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"isProviderRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerLogin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerLogout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerReputationInit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerReputationQualified\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenContract\",\"outputs\":[{\"internalType\":\"contractFaaSToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ProviderPoolBin is the compiled bytecode used for deploying new contracts.
var ProviderPoolBin = "0x608060405234801561001057600080fd5b506000806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050606460018190555060056003819055506005600481905550610a1c8061007b6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063729d98e911610071578063729d98e91461019a578063757f4a97146101f2578063884eacb814610210578063b82964711461021a578063dbeefdd214610276578063f0611e3814610294576100a9565b80631ead914e146100ae578063215218921461010a578063230544a7146101285780632bf355701461014657806355a373d614610150575b600080fd5b6100f0600480360360208110156100c457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506102b2565b604051808215151515815260200191505060405180910390f35b6101126102fd565b6040518082815260200191505060405180910390f35b610130610303565b6040518082815260200191505060405180910390f35b61014e610309565b005b610158610571565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101dc600480360360208110156101b057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610596565b6040518082815260200191505060405180910390f35b6101fa6105df565b6040518082815260200191505060405180910390f35b6102186105e9565b005b61025c6004803603602081101561023057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506108bc565b604051808215151515815260200191505060405180910390f35b61027e61090a565b6040518082815260200191505060405180910390f35b61029c610914565b6040518082815260200191505060405180910390f35b600080600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054119050919050565b60045481565b60015481565b60001515610316336102b2565b15151461036e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260278152602001806109436027913960400191505060405180910390fd5b600115156000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd33306001546040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b15801561045057600080fd5b505af1158015610464573d6000803e3d6000fd5b505050506040513d602081101561047a57600080fd5b81019080805190602001909291905050501515146104e3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602881526020018061091b6028913960400191505060405180910390fd5b600154600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600354600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000600454905090565b600115156105f6336102b2565b15151461064e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602b81526020018061098f602b913960400191505060405180910390fd5b610657336108bc565b6106ac576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602581526020018061096a6025913960400191505060405180910390fd5b6000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600115156000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33846040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561082657600080fd5b505af115801561083a573d6000803e3d6000fd5b505050506040513d602081101561085057600080fd5b81019080805190602001909291905050501515146108b9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602d8152602001806109ba602d913960400191505060405180910390fd5b50565b6000600454600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410159050919050565b6000600154905090565b6003548156fe4d61726b65743a206661696c656420746f207061792061207265676973746572206465706f7369744d61726b65743a20746865206164647265737320686164206265656e20726567697374657265644d61726b65743a207468652070726f7669646572206973206e6f74207175616c69666965644d61726b65743a20746865206164647265737320686164206e6f74206265656e20726567697374657265644d61726b65743a206661696c656420746f20726566756e64207468652070726f7669646572206465706f736974a26469706673582212203f67fa0ad357258208dfb34b06233b6a04a0dd05a291c99cf7ab4ec61bd7182e64736f6c634300060a0033"

// DeployProviderPool deploys a new contract, binding an instance of ProviderPool to it.
func DeployProviderPool(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ProviderPool, error) {
	parsed, err := abi.JSON(strings.NewReader(ProviderPoolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProviderPoolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProviderPool{ProviderPoolCaller: ProviderPoolCaller{contract: contract}, ProviderPoolTransactor: ProviderPoolTransactor{contract: contract}, ProviderPoolFilterer: ProviderPoolFilterer{contract: contract}}, nil
}

func AsyncDeployProviderPool(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(ProviderPoolABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(ProviderPoolBin), backend)
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

// GetProviderDeposit is a free data retrieval call binding the contract method 0xdbeefdd2.
//
// Solidity: function getProviderDeposit() constant returns(uint256)
func (_ProviderPool *ProviderPoolCaller) GetProviderDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ProviderPool.contract.Call(opts, out, "getProviderDeposit")
	return *ret0, err
}

// GetProviderDeposit is a free data retrieval call binding the contract method 0xdbeefdd2.
//
// Solidity: function getProviderDeposit() constant returns(uint256)
func (_ProviderPool *ProviderPoolSession) GetProviderDeposit() (*big.Int, error) {
	return _ProviderPool.Contract.GetProviderDeposit(&_ProviderPool.CallOpts)
}

// GetProviderDeposit is a free data retrieval call binding the contract method 0xdbeefdd2.
//
// Solidity: function getProviderDeposit() constant returns(uint256)
func (_ProviderPool *ProviderPoolCallerSession) GetProviderDeposit() (*big.Int, error) {
	return _ProviderPool.Contract.GetProviderDeposit(&_ProviderPool.CallOpts)
}

// GetProviderReputation is a free data retrieval call binding the contract method 0x729d98e9.
//
// Solidity: function getProviderReputation(address provider) constant returns(uint256)
func (_ProviderPool *ProviderPoolCaller) GetProviderReputation(opts *bind.CallOpts, provider common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ProviderPool.contract.Call(opts, out, "getProviderReputation", provider)
	return *ret0, err
}

// GetProviderReputation is a free data retrieval call binding the contract method 0x729d98e9.
//
// Solidity: function getProviderReputation(address provider) constant returns(uint256)
func (_ProviderPool *ProviderPoolSession) GetProviderReputation(provider common.Address) (*big.Int, error) {
	return _ProviderPool.Contract.GetProviderReputation(&_ProviderPool.CallOpts, provider)
}

// GetProviderReputation is a free data retrieval call binding the contract method 0x729d98e9.
//
// Solidity: function getProviderReputation(address provider) constant returns(uint256)
func (_ProviderPool *ProviderPoolCallerSession) GetProviderReputation(provider common.Address) (*big.Int, error) {
	return _ProviderPool.Contract.GetProviderReputation(&_ProviderPool.CallOpts, provider)
}

// GetProviderReputationQualified is a free data retrieval call binding the contract method 0x757f4a97.
//
// Solidity: function getProviderReputationQualified() constant returns(uint256)
func (_ProviderPool *ProviderPoolCaller) GetProviderReputationQualified(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ProviderPool.contract.Call(opts, out, "getProviderReputationQualified")
	return *ret0, err
}

// GetProviderReputationQualified is a free data retrieval call binding the contract method 0x757f4a97.
//
// Solidity: function getProviderReputationQualified() constant returns(uint256)
func (_ProviderPool *ProviderPoolSession) GetProviderReputationQualified() (*big.Int, error) {
	return _ProviderPool.Contract.GetProviderReputationQualified(&_ProviderPool.CallOpts)
}

// GetProviderReputationQualified is a free data retrieval call binding the contract method 0x757f4a97.
//
// Solidity: function getProviderReputationQualified() constant returns(uint256)
func (_ProviderPool *ProviderPoolCallerSession) GetProviderReputationQualified() (*big.Int, error) {
	return _ProviderPool.Contract.GetProviderReputationQualified(&_ProviderPool.CallOpts)
}

// IsProviderQualified is a free data retrieval call binding the contract method 0xb8296471.
//
// Solidity: function isProviderQualified(address provider) constant returns(bool)
func (_ProviderPool *ProviderPoolCaller) IsProviderQualified(opts *bind.CallOpts, provider common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ProviderPool.contract.Call(opts, out, "isProviderQualified", provider)
	return *ret0, err
}

// IsProviderQualified is a free data retrieval call binding the contract method 0xb8296471.
//
// Solidity: function isProviderQualified(address provider) constant returns(bool)
func (_ProviderPool *ProviderPoolSession) IsProviderQualified(provider common.Address) (bool, error) {
	return _ProviderPool.Contract.IsProviderQualified(&_ProviderPool.CallOpts, provider)
}

// IsProviderQualified is a free data retrieval call binding the contract method 0xb8296471.
//
// Solidity: function isProviderQualified(address provider) constant returns(bool)
func (_ProviderPool *ProviderPoolCallerSession) IsProviderQualified(provider common.Address) (bool, error) {
	return _ProviderPool.Contract.IsProviderQualified(&_ProviderPool.CallOpts, provider)
}

// IsProviderRegistered is a free data retrieval call binding the contract method 0x1ead914e.
//
// Solidity: function isProviderRegistered(address provider) constant returns(bool)
func (_ProviderPool *ProviderPoolCaller) IsProviderRegistered(opts *bind.CallOpts, provider common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ProviderPool.contract.Call(opts, out, "isProviderRegistered", provider)
	return *ret0, err
}

// IsProviderRegistered is a free data retrieval call binding the contract method 0x1ead914e.
//
// Solidity: function isProviderRegistered(address provider) constant returns(bool)
func (_ProviderPool *ProviderPoolSession) IsProviderRegistered(provider common.Address) (bool, error) {
	return _ProviderPool.Contract.IsProviderRegistered(&_ProviderPool.CallOpts, provider)
}

// IsProviderRegistered is a free data retrieval call binding the contract method 0x1ead914e.
//
// Solidity: function isProviderRegistered(address provider) constant returns(bool)
func (_ProviderPool *ProviderPoolCallerSession) IsProviderRegistered(provider common.Address) (bool, error) {
	return _ProviderPool.Contract.IsProviderRegistered(&_ProviderPool.CallOpts, provider)
}

// ProviderDeposit is a free data retrieval call binding the contract method 0x230544a7.
//
// Solidity: function providerDeposit() constant returns(uint256)
func (_ProviderPool *ProviderPoolCaller) ProviderDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ProviderPool.contract.Call(opts, out, "providerDeposit")
	return *ret0, err
}

// ProviderDeposit is a free data retrieval call binding the contract method 0x230544a7.
//
// Solidity: function providerDeposit() constant returns(uint256)
func (_ProviderPool *ProviderPoolSession) ProviderDeposit() (*big.Int, error) {
	return _ProviderPool.Contract.ProviderDeposit(&_ProviderPool.CallOpts)
}

// ProviderDeposit is a free data retrieval call binding the contract method 0x230544a7.
//
// Solidity: function providerDeposit() constant returns(uint256)
func (_ProviderPool *ProviderPoolCallerSession) ProviderDeposit() (*big.Int, error) {
	return _ProviderPool.Contract.ProviderDeposit(&_ProviderPool.CallOpts)
}

// ProviderReputationInit is a free data retrieval call binding the contract method 0xf0611e38.
//
// Solidity: function providerReputationInit() constant returns(uint256)
func (_ProviderPool *ProviderPoolCaller) ProviderReputationInit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ProviderPool.contract.Call(opts, out, "providerReputationInit")
	return *ret0, err
}

// ProviderReputationInit is a free data retrieval call binding the contract method 0xf0611e38.
//
// Solidity: function providerReputationInit() constant returns(uint256)
func (_ProviderPool *ProviderPoolSession) ProviderReputationInit() (*big.Int, error) {
	return _ProviderPool.Contract.ProviderReputationInit(&_ProviderPool.CallOpts)
}

// ProviderReputationInit is a free data retrieval call binding the contract method 0xf0611e38.
//
// Solidity: function providerReputationInit() constant returns(uint256)
func (_ProviderPool *ProviderPoolCallerSession) ProviderReputationInit() (*big.Int, error) {
	return _ProviderPool.Contract.ProviderReputationInit(&_ProviderPool.CallOpts)
}

// ProviderReputationQualified is a free data retrieval call binding the contract method 0x21521892.
//
// Solidity: function providerReputationQualified() constant returns(uint256)
func (_ProviderPool *ProviderPoolCaller) ProviderReputationQualified(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ProviderPool.contract.Call(opts, out, "providerReputationQualified")
	return *ret0, err
}

// ProviderReputationQualified is a free data retrieval call binding the contract method 0x21521892.
//
// Solidity: function providerReputationQualified() constant returns(uint256)
func (_ProviderPool *ProviderPoolSession) ProviderReputationQualified() (*big.Int, error) {
	return _ProviderPool.Contract.ProviderReputationQualified(&_ProviderPool.CallOpts)
}

// ProviderReputationQualified is a free data retrieval call binding the contract method 0x21521892.
//
// Solidity: function providerReputationQualified() constant returns(uint256)
func (_ProviderPool *ProviderPoolCallerSession) ProviderReputationQualified() (*big.Int, error) {
	return _ProviderPool.Contract.ProviderReputationQualified(&_ProviderPool.CallOpts)
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
