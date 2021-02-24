// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package market

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

// MarketABI is the input ABI used to generate the binding from.
const MarketABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContractAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_delpoymentOrderID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_auctionAddress\",\"type\":\"address\"}],\"name\":\"newDeploymentOrderEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctions\",\"outputs\":[{\"internalType\":\"contractSimpleAuction\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"delpoymentOrders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"customer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"faaSLevelID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"faaSDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highestUnitPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_delpoymentOrderID\",\"type\":\"uint256\"}],\"name\":\"getDeploymentOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProviderDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"getProviderReputation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProviderReputationQualified\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"leases\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"customer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"faaSLevelID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"faaSDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unitPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"providerServiceFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"customerCompensationFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"witnessFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maintainerFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_faaSLevelID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_highestUnitPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_faaSDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_biddingDuration\",\"type\":\"uint256\"}],\"name\":\"newDeploymentOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_delpoymentOrderID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_unitPrice\",\"type\":\"uint256\"}],\"name\":\"newLease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numDelpoymentOrders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numLeases\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerLogin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerLogout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerReputationInit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerReputationQualified\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MarketBin is the compiled bytecode used for deploying new contracts.
var MarketBin = "0x608060405234801561001057600080fd5b5060405161205a38038061205a8339818101604052602081101561003357600080fd5b8101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050611fc6806100946000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c8063884eacb811610097578063dbd4314f11610066578063dbd4314f14610436578063dbeefdd2146104c9578063f0611e38146104e7578063f843dcc41461050557610100565b8063884eacb8146102b95780638927a106146102c3578063cddd5df714610395578063ce67ca03146103b357610100565b80634192b663116100d35780634192b66314610169578063571a26a0146101d5578063729d98e914610243578063757f4a971461029b57610100565b806306263993146101055780632152189214610123578063230544a7146101415780632bf355701461015f575b600080fd5b61010d610588565b6040518082815260200191505060405180910390f35b61012b61058e565b6040518082815260200191505060405180910390f35b610149610594565b6040518082815260200191505060405180910390f35b61016761059a565b005b6101bf6004803603606081101561017f57600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610835565b6040518082815260200191505060405180910390f35b610201600480360360208110156101eb57600080fd5b8101908080359060200190929190505050610b32565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102856004803603602081101561025957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b65565b6040518082815260200191505060405180910390f35b6102a3610bae565b6040518082815260200191505060405180910390f35b6102c1610bb8565b005b6102ef600480360360208110156102d957600080fd5b8101908080359060200190929190505050610f3c565b604051808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001888152602001878152602001868152602001858152602001848152602001838152602001828152602001995050505050505050505060405180910390f35b61039d610fca565b6040518082815260200191505060405180910390f35b6103df600480360360208110156103c957600080fd5b8101908080359060200190929190505050610fd0565b604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200184815260200183815260200182815260200194505050505060405180910390f35b6104806004803603608081101561044c57600080fd5b8101908080359060200190929190803590602001909291908035906020019092919080359060200190929190505050611096565b604051808381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390f35b6104d16113f1565b6040518082815260200191505060405180910390f35b6104ef6113fb565b6040518082815260200191505060405180910390f35b6105316004803603602081101561051b57600080fd5b8101908080359060200190929190505050611401565b604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200184815260200183815260200182815260200194505050505060405180910390f35b60065481565b60045481565b60015481565b6000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414610632576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526027815260200180611e916027913960400191505060405180910390fd5b600115156000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd33306001546040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b15801561071457600080fd5b505af1158015610728573d6000803e3d6000fd5b505050506040513d602081101561073e57600080fd5b81019080805190602001909291905050501515146107a7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526028815260200180611e696028913960400191505060405180910390fd5b600154600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600354600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550565b60006008600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108ee576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526033815260200180611eb86033913960400191505060405180910390fd5b6108f6611451565b600760008681526020019081526020016000206040518060800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820154815260200160028201548152602001600382015481525050905060006009600081548092919060010191905055905060008483604001510290506000819050600080905060008090506000809050604051806101200160405280886000015173ffffffffffffffffffffffffffffffffffffffff1681526020018b73ffffffffffffffffffffffffffffffffffffffff16815260200188602001518152602001886040015181526020018a815260200185815260200184815260200183815260200182815250600a600088815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e082015181600701556101008201518160080155905050859750505050505050509392505050565b60086020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000600454905090565b6000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205411610c50576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602b815260200180611f10602b913960400191505060405180910390fd5b600454600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205411610ce9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180611eeb6025913960400191505060405180910390fd5b6000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541115610f3a57600115156000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015610e6157600080fd5b505af1158015610e75573d6000803e3d6000fd5b505050506040513d6020811015610e8b57600080fd5b8101908080519060200190929190505050151514610ef4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602d815260200180611f64602d913960400191505060405180910390fd5b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b565b600a6020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154908060030154908060040154908060050154908060060154908060070154908060080154905089565b60095481565b600080600080610fde611451565b600760008781526020019081526020016000206040518060800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820154815260200160028201548152602001600382015481525050905080600001518160200151826040015183606001519450945094509450509193509193565b60008060008486029050600115156000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b15801561118057600080fd5b505af1158015611194573d6000803e3d6000fd5b505050506040513d60208110156111aa57600080fd5b8101908080519060200190929190505050151514611213576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526029815260200180611f3b6029913960400191505060405180910390fd5b60006006600081548092919060010191905055905060405180608001604052803373ffffffffffffffffffffffffffffffffffffffff168152602001878152602001878152602001888152506007600083815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201556060820151816003015590505060008188876040516112eb9061148f565b808481526020018381526020018281526020019350505050604051809103906000f08015801561131f573d6000803e3d6000fd5b509050806008600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f93b403b917e94d112528f49b59221bab32fb8fd897cbc154bf092ef7796d5dd08282604051808381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a181819450945050505094509492505050565b6000600154905090565b60035481565b60076020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154908060030154905084565b6040518060800160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001600081525090565b6109cc8061149d8339019056fe608060405234801561001057600080fd5b506040516109cc3803806109cc8339818101604052606081101561003357600080fd5b8101908080519060200190929190805190602001909291908051906020019092919050505060008060006101000a81548160ff0219169083600281111561007657fe5b021790555082600181905550816002819055504260038190555080600481905550816005819055506000600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600660146101000a81548160ff0219169083151502179055505050506108bf8061010d6000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806361dcd7ab1161005b57806361dcd7ab14610101578063c19d93fb1461011f578063c8e7b31d1461014b578063de2927891461016957610088565b80632a24f46c1461008d578063454a2ab3146100975780634840f2f0146100c55780635d06978f146100e3575b600080fd5b6100956101cc565b005b6100c3600480360360208110156100ad57600080fd5b81019080803590602001909291905050506103b8565b005b6100cd610617565b6040518082815260200191505060405180910390f35b6100eb61061d565b6040518082815260200191505060405180910390f35b610109610623565b6040518082815260200191505060405180910390f35b610127610629565b6040518082600281111561013757fe5b60ff16815260200191505060405180910390f35b61015361063b565b6040518082815260200191505060405180910390f35b610171610641565b60405180851515151581526020018481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200194505050505060405180910390f35b600060028111156101d957fe5b6000809054906101000a900460ff1660028111156101f357fe5b14801561020557506004546003540142115b156102135761021261070d565b5b600180600281111561022157fe5b6000809054906101000a900460ff16600281111561023b57fe5b14610291576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260288152602001806108626028913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610303576001600660146101000a81548160ff0219169083151502179055505b7f522abce2186043ff800e5e4b54ae810b7041225964a29cc14bb825bdc073f935600660149054906101000a900460ff16600154600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660405180841515151581526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001935050505060405180910390a16103b561070d565b50565b600060028111156103c557fe5b6000809054906101000a900460ff1660028111156103df57fe5b1480156103f157506004546003540142115b156103ff576103fe61070d565b5b600080600281111561040d57fe5b6000809054906101000a900460ff16600281111561042757fe5b1461047d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260288152602001806108626028913960400191505060405180910390fd5b6002548211156104d8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603c815260200180610802603c913960400191505060405180910390fd5b6005548210610532576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602481526020018061083e6024913960400191505060405180910390fd5b8160058190555033600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f61a89883ddb4daf2c81411c80011c699535d5386345b29246cff5b493a6dc161600154600554600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051808481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001935050505060405180910390a15050565b60025481565b60015481565b60035481565b6000809054906101000a900460ff1681565b60045481565b600080600080600280600281111561065557fe5b6000809054906101000a900460ff16600281111561066f57fe5b146106c5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260288152602001806108626028913960400191505060405180910390fd5b600660149054906101000a900460ff16600154600554600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1694509450945094505090919293565b6000600281111561071a57fe5b6000809054906101000a900460ff16600281111561073457fe5b141561075e5760016000806101000a81548160ff0219169083600281111561075857fe5b02179055505b6001600281111561076b57fe5b6000809054906101000a900460ff16600281111561078557fe5b14156107af5760026000806101000a81548160ff021916908360028111156107a957fe5b02179055505b6002808111156107bb57fe5b6000809054906101000a900460ff1660028111156107d557fe5b14156107ff5760026000806101000a81548160ff021916908360028111156107f957fe5b02179055505b56fe54686520756e69742d707269636520697320686967686572207468616e20686967686573742074686520637573746f6d65722061636365707465642e546865726520616c72656164792069732061206c6f7765722075696e742d70726963652e46756e6374696f6e2063616e6e6f742062652063616c6c656420617420746869732073746174652ea2646970667358221220797865f27e5f5a4d0ae52778bfa3056aacdb3bc7bc171218842af78f73a28cce64736f6c634300060a00334d61726b65743a206661696c656420746f207061792061207265676973746572206465706f7369744d61726b65743a20746865206164647265737320686164206265656e20726567697374657265644f6e6c792073706563696669632061756374696f6e20736d61727420636f6e74726163742063616e206e6577206c656173652e4d61726b65743a207468652070726f7669646572206973206e6f74207175616c69666965644d61726b65743a20746865206164647265737320686164206e6f74206265656e20726567697374657265646c6f636b20666565206265666f7265206372656174696e67206465706c6f796d656e74206f726465724d61726b65743a206661696c656420746f20726566756e64207468652070726f7669646572206465706f736974a2646970667358221220f7445eb5240c4d063eac3ba11b637524f8cfaebb6a52ac09f7e2d3209f51a6a464736f6c634300060a0033"

// DeployMarket deploys a new contract, binding an instance of Market to it.
func DeployMarket(auth *bind.TransactOpts, backend bind.ContractBackend, tokenContractAddress common.Address) (common.Address, *types.Transaction, *Market, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MarketBin), backend, tokenContractAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Market{MarketCaller: MarketCaller{contract: contract}, MarketTransactor: MarketTransactor{contract: contract}, MarketFilterer: MarketFilterer{contract: contract}}, nil
}

func AsyncDeployMarket(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend, tokenContractAddress common.Address) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(MarketBin), backend, tokenContractAddress)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Market is an auto generated Go binding around a Solidity contract.
type Market struct {
	MarketCaller     // Read-only binding to the contract
	MarketTransactor // Write-only binding to the contract
	MarketFilterer   // Log filterer for contract events
}

// MarketCaller is an auto generated read-only Go binding around a Solidity contract.
type MarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketTransactor is an auto generated write-only Go binding around a Solidity contract.
type MarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type MarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type MarketSession struct {
	Contract     *Market           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type MarketCallerSession struct {
	Contract *MarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MarketTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type MarketTransactorSession struct {
	Contract     *MarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketRaw is an auto generated low-level Go binding around a Solidity contract.
type MarketRaw struct {
	Contract *Market // Generic contract binding to access the raw methods on
}

// MarketCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type MarketCallerRaw struct {
	Contract *MarketCaller // Generic read-only contract binding to access the raw methods on
}

// MarketTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type MarketTransactorRaw struct {
	Contract *MarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarket creates a new instance of Market, bound to a specific deployed contract.
func NewMarket(address common.Address, backend bind.ContractBackend) (*Market, error) {
	contract, err := bindMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Market{MarketCaller: MarketCaller{contract: contract}, MarketTransactor: MarketTransactor{contract: contract}, MarketFilterer: MarketFilterer{contract: contract}}, nil
}

// NewMarketCaller creates a new read-only instance of Market, bound to a specific deployed contract.
func NewMarketCaller(address common.Address, caller bind.ContractCaller) (*MarketCaller, error) {
	contract, err := bindMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketCaller{contract: contract}, nil
}

// NewMarketTransactor creates a new write-only instance of Market, bound to a specific deployed contract.
func NewMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketTransactor, error) {
	contract, err := bindMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketTransactor{contract: contract}, nil
}

// NewMarketFilterer creates a new log filterer instance of Market, bound to a specific deployed contract.
func NewMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketFilterer, error) {
	contract, err := bindMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketFilterer{contract: contract}, nil
}

// bindMarket binds a generic wrapper to an already deployed contract.
func bindMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Market.Contract.MarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Market.Contract.MarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Market.Contract.MarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Market.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Market.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Market.Contract.contract.Transact(opts, method, params...)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) constant returns(address)
func (_Market *MarketCaller) Auctions(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "auctions", arg0)
	return *ret0, err
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) constant returns(address)
func (_Market *MarketSession) Auctions(arg0 *big.Int) (common.Address, error) {
	return _Market.Contract.Auctions(&_Market.CallOpts, arg0)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) constant returns(address)
func (_Market *MarketCallerSession) Auctions(arg0 *big.Int) (common.Address, error) {
	return _Market.Contract.Auctions(&_Market.CallOpts, arg0)
}

// DelpoymentOrders is a free data retrieval call binding the contract method 0xf843dcc4.
//
// Solidity: function delpoymentOrders(uint256 ) constant returns(address customer, uint256 faaSLevelID, uint256 faaSDuration, uint256 highestUnitPrice)
func (_Market *MarketCaller) DelpoymentOrders(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Customer         common.Address
	FaaSLevelID      *big.Int
	FaaSDuration     *big.Int
	HighestUnitPrice *big.Int
}, error) {
	ret := new(struct {
		Customer         common.Address
		FaaSLevelID      *big.Int
		FaaSDuration     *big.Int
		HighestUnitPrice *big.Int
	})
	out := ret
	err := _Market.contract.Call(opts, out, "delpoymentOrders", arg0)
	return *ret, err
}

// DelpoymentOrders is a free data retrieval call binding the contract method 0xf843dcc4.
//
// Solidity: function delpoymentOrders(uint256 ) constant returns(address customer, uint256 faaSLevelID, uint256 faaSDuration, uint256 highestUnitPrice)
func (_Market *MarketSession) DelpoymentOrders(arg0 *big.Int) (struct {
	Customer         common.Address
	FaaSLevelID      *big.Int
	FaaSDuration     *big.Int
	HighestUnitPrice *big.Int
}, error) {
	return _Market.Contract.DelpoymentOrders(&_Market.CallOpts, arg0)
}

// DelpoymentOrders is a free data retrieval call binding the contract method 0xf843dcc4.
//
// Solidity: function delpoymentOrders(uint256 ) constant returns(address customer, uint256 faaSLevelID, uint256 faaSDuration, uint256 highestUnitPrice)
func (_Market *MarketCallerSession) DelpoymentOrders(arg0 *big.Int) (struct {
	Customer         common.Address
	FaaSLevelID      *big.Int
	FaaSDuration     *big.Int
	HighestUnitPrice *big.Int
}, error) {
	return _Market.Contract.DelpoymentOrders(&_Market.CallOpts, arg0)
}

// GetDeploymentOrder is a free data retrieval call binding the contract method 0xce67ca03.
//
// Solidity: function getDeploymentOrder(uint256 _delpoymentOrderID) constant returns(address, uint256, uint256, uint256)
func (_Market *MarketCaller) GetDeploymentOrder(opts *bind.CallOpts, _delpoymentOrderID *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Market.contract.Call(opts, out, "getDeploymentOrder", _delpoymentOrderID)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetDeploymentOrder is a free data retrieval call binding the contract method 0xce67ca03.
//
// Solidity: function getDeploymentOrder(uint256 _delpoymentOrderID) constant returns(address, uint256, uint256, uint256)
func (_Market *MarketSession) GetDeploymentOrder(_delpoymentOrderID *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _Market.Contract.GetDeploymentOrder(&_Market.CallOpts, _delpoymentOrderID)
}

// GetDeploymentOrder is a free data retrieval call binding the contract method 0xce67ca03.
//
// Solidity: function getDeploymentOrder(uint256 _delpoymentOrderID) constant returns(address, uint256, uint256, uint256)
func (_Market *MarketCallerSession) GetDeploymentOrder(_delpoymentOrderID *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _Market.Contract.GetDeploymentOrder(&_Market.CallOpts, _delpoymentOrderID)
}

// GetProviderDeposit is a free data retrieval call binding the contract method 0xdbeefdd2.
//
// Solidity: function getProviderDeposit() constant returns(uint256)
func (_Market *MarketCaller) GetProviderDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getProviderDeposit")
	return *ret0, err
}

// GetProviderDeposit is a free data retrieval call binding the contract method 0xdbeefdd2.
//
// Solidity: function getProviderDeposit() constant returns(uint256)
func (_Market *MarketSession) GetProviderDeposit() (*big.Int, error) {
	return _Market.Contract.GetProviderDeposit(&_Market.CallOpts)
}

// GetProviderDeposit is a free data retrieval call binding the contract method 0xdbeefdd2.
//
// Solidity: function getProviderDeposit() constant returns(uint256)
func (_Market *MarketCallerSession) GetProviderDeposit() (*big.Int, error) {
	return _Market.Contract.GetProviderDeposit(&_Market.CallOpts)
}

// GetProviderReputation is a free data retrieval call binding the contract method 0x729d98e9.
//
// Solidity: function getProviderReputation(address provider) constant returns(uint256)
func (_Market *MarketCaller) GetProviderReputation(opts *bind.CallOpts, provider common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getProviderReputation", provider)
	return *ret0, err
}

// GetProviderReputation is a free data retrieval call binding the contract method 0x729d98e9.
//
// Solidity: function getProviderReputation(address provider) constant returns(uint256)
func (_Market *MarketSession) GetProviderReputation(provider common.Address) (*big.Int, error) {
	return _Market.Contract.GetProviderReputation(&_Market.CallOpts, provider)
}

// GetProviderReputation is a free data retrieval call binding the contract method 0x729d98e9.
//
// Solidity: function getProviderReputation(address provider) constant returns(uint256)
func (_Market *MarketCallerSession) GetProviderReputation(provider common.Address) (*big.Int, error) {
	return _Market.Contract.GetProviderReputation(&_Market.CallOpts, provider)
}

// GetProviderReputationQualified is a free data retrieval call binding the contract method 0x757f4a97.
//
// Solidity: function getProviderReputationQualified() constant returns(uint256)
func (_Market *MarketCaller) GetProviderReputationQualified(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getProviderReputationQualified")
	return *ret0, err
}

// GetProviderReputationQualified is a free data retrieval call binding the contract method 0x757f4a97.
//
// Solidity: function getProviderReputationQualified() constant returns(uint256)
func (_Market *MarketSession) GetProviderReputationQualified() (*big.Int, error) {
	return _Market.Contract.GetProviderReputationQualified(&_Market.CallOpts)
}

// GetProviderReputationQualified is a free data retrieval call binding the contract method 0x757f4a97.
//
// Solidity: function getProviderReputationQualified() constant returns(uint256)
func (_Market *MarketCallerSession) GetProviderReputationQualified() (*big.Int, error) {
	return _Market.Contract.GetProviderReputationQualified(&_Market.CallOpts)
}

// Leases is a free data retrieval call binding the contract method 0x8927a106.
//
// Solidity: function leases(uint256 ) constant returns(address customer, address provider, uint256 faaSLevelID, uint256 faaSDuration, uint256 unitPrice, uint256 providerServiceFee, uint256 customerCompensationFee, uint256 witnessFee, uint256 maintainerFee)
func (_Market *MarketCaller) Leases(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Customer                common.Address
	Provider                common.Address
	FaaSLevelID             *big.Int
	FaaSDuration            *big.Int
	UnitPrice               *big.Int
	ProviderServiceFee      *big.Int
	CustomerCompensationFee *big.Int
	WitnessFee              *big.Int
	MaintainerFee           *big.Int
}, error) {
	ret := new(struct {
		Customer                common.Address
		Provider                common.Address
		FaaSLevelID             *big.Int
		FaaSDuration            *big.Int
		UnitPrice               *big.Int
		ProviderServiceFee      *big.Int
		CustomerCompensationFee *big.Int
		WitnessFee              *big.Int
		MaintainerFee           *big.Int
	})
	out := ret
	err := _Market.contract.Call(opts, out, "leases", arg0)
	return *ret, err
}

// Leases is a free data retrieval call binding the contract method 0x8927a106.
//
// Solidity: function leases(uint256 ) constant returns(address customer, address provider, uint256 faaSLevelID, uint256 faaSDuration, uint256 unitPrice, uint256 providerServiceFee, uint256 customerCompensationFee, uint256 witnessFee, uint256 maintainerFee)
func (_Market *MarketSession) Leases(arg0 *big.Int) (struct {
	Customer                common.Address
	Provider                common.Address
	FaaSLevelID             *big.Int
	FaaSDuration            *big.Int
	UnitPrice               *big.Int
	ProviderServiceFee      *big.Int
	CustomerCompensationFee *big.Int
	WitnessFee              *big.Int
	MaintainerFee           *big.Int
}, error) {
	return _Market.Contract.Leases(&_Market.CallOpts, arg0)
}

// Leases is a free data retrieval call binding the contract method 0x8927a106.
//
// Solidity: function leases(uint256 ) constant returns(address customer, address provider, uint256 faaSLevelID, uint256 faaSDuration, uint256 unitPrice, uint256 providerServiceFee, uint256 customerCompensationFee, uint256 witnessFee, uint256 maintainerFee)
func (_Market *MarketCallerSession) Leases(arg0 *big.Int) (struct {
	Customer                common.Address
	Provider                common.Address
	FaaSLevelID             *big.Int
	FaaSDuration            *big.Int
	UnitPrice               *big.Int
	ProviderServiceFee      *big.Int
	CustomerCompensationFee *big.Int
	WitnessFee              *big.Int
	MaintainerFee           *big.Int
}, error) {
	return _Market.Contract.Leases(&_Market.CallOpts, arg0)
}

// NumDelpoymentOrders is a free data retrieval call binding the contract method 0x06263993.
//
// Solidity: function numDelpoymentOrders() constant returns(uint256)
func (_Market *MarketCaller) NumDelpoymentOrders(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "numDelpoymentOrders")
	return *ret0, err
}

// NumDelpoymentOrders is a free data retrieval call binding the contract method 0x06263993.
//
// Solidity: function numDelpoymentOrders() constant returns(uint256)
func (_Market *MarketSession) NumDelpoymentOrders() (*big.Int, error) {
	return _Market.Contract.NumDelpoymentOrders(&_Market.CallOpts)
}

// NumDelpoymentOrders is a free data retrieval call binding the contract method 0x06263993.
//
// Solidity: function numDelpoymentOrders() constant returns(uint256)
func (_Market *MarketCallerSession) NumDelpoymentOrders() (*big.Int, error) {
	return _Market.Contract.NumDelpoymentOrders(&_Market.CallOpts)
}

// NumLeases is a free data retrieval call binding the contract method 0xcddd5df7.
//
// Solidity: function numLeases() constant returns(uint256)
func (_Market *MarketCaller) NumLeases(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "numLeases")
	return *ret0, err
}

// NumLeases is a free data retrieval call binding the contract method 0xcddd5df7.
//
// Solidity: function numLeases() constant returns(uint256)
func (_Market *MarketSession) NumLeases() (*big.Int, error) {
	return _Market.Contract.NumLeases(&_Market.CallOpts)
}

// NumLeases is a free data retrieval call binding the contract method 0xcddd5df7.
//
// Solidity: function numLeases() constant returns(uint256)
func (_Market *MarketCallerSession) NumLeases() (*big.Int, error) {
	return _Market.Contract.NumLeases(&_Market.CallOpts)
}

// ProviderDeposit is a free data retrieval call binding the contract method 0x230544a7.
//
// Solidity: function providerDeposit() constant returns(uint256)
func (_Market *MarketCaller) ProviderDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "providerDeposit")
	return *ret0, err
}

// ProviderDeposit is a free data retrieval call binding the contract method 0x230544a7.
//
// Solidity: function providerDeposit() constant returns(uint256)
func (_Market *MarketSession) ProviderDeposit() (*big.Int, error) {
	return _Market.Contract.ProviderDeposit(&_Market.CallOpts)
}

// ProviderDeposit is a free data retrieval call binding the contract method 0x230544a7.
//
// Solidity: function providerDeposit() constant returns(uint256)
func (_Market *MarketCallerSession) ProviderDeposit() (*big.Int, error) {
	return _Market.Contract.ProviderDeposit(&_Market.CallOpts)
}

// ProviderReputationInit is a free data retrieval call binding the contract method 0xf0611e38.
//
// Solidity: function providerReputationInit() constant returns(uint256)
func (_Market *MarketCaller) ProviderReputationInit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "providerReputationInit")
	return *ret0, err
}

// ProviderReputationInit is a free data retrieval call binding the contract method 0xf0611e38.
//
// Solidity: function providerReputationInit() constant returns(uint256)
func (_Market *MarketSession) ProviderReputationInit() (*big.Int, error) {
	return _Market.Contract.ProviderReputationInit(&_Market.CallOpts)
}

// ProviderReputationInit is a free data retrieval call binding the contract method 0xf0611e38.
//
// Solidity: function providerReputationInit() constant returns(uint256)
func (_Market *MarketCallerSession) ProviderReputationInit() (*big.Int, error) {
	return _Market.Contract.ProviderReputationInit(&_Market.CallOpts)
}

// ProviderReputationQualified is a free data retrieval call binding the contract method 0x21521892.
//
// Solidity: function providerReputationQualified() constant returns(uint256)
func (_Market *MarketCaller) ProviderReputationQualified(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "providerReputationQualified")
	return *ret0, err
}

// ProviderReputationQualified is a free data retrieval call binding the contract method 0x21521892.
//
// Solidity: function providerReputationQualified() constant returns(uint256)
func (_Market *MarketSession) ProviderReputationQualified() (*big.Int, error) {
	return _Market.Contract.ProviderReputationQualified(&_Market.CallOpts)
}

// ProviderReputationQualified is a free data retrieval call binding the contract method 0x21521892.
//
// Solidity: function providerReputationQualified() constant returns(uint256)
func (_Market *MarketCallerSession) ProviderReputationQualified() (*big.Int, error) {
	return _Market.Contract.ProviderReputationQualified(&_Market.CallOpts)
}

// NewDeploymentOrder is a paid mutator transaction binding the contract method 0xdbd4314f.
//
// Solidity: function newDeploymentOrder(uint256 _faaSLevelID, uint256 _highestUnitPrice, uint256 _faaSDuration, uint256 _biddingDuration) returns(uint256, address)
func (_Market *MarketTransactor) NewDeploymentOrder(opts *bind.TransactOpts, _faaSLevelID *big.Int, _highestUnitPrice *big.Int, _faaSDuration *big.Int, _biddingDuration *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Market.contract.Transact(opts, "newDeploymentOrder", _faaSLevelID, _highestUnitPrice, _faaSDuration, _biddingDuration)
}

func (_Market *MarketTransactor) AsyncNewDeploymentOrder(handler func(*types.Receipt, error), opts *bind.TransactOpts, _faaSLevelID *big.Int, _highestUnitPrice *big.Int, _faaSDuration *big.Int, _biddingDuration *big.Int) (*types.Transaction, error) {
	return _Market.contract.AsyncTransact(opts, handler, "newDeploymentOrder", _faaSLevelID, _highestUnitPrice, _faaSDuration, _biddingDuration)
}

// NewDeploymentOrder is a paid mutator transaction binding the contract method 0xdbd4314f.
//
// Solidity: function newDeploymentOrder(uint256 _faaSLevelID, uint256 _highestUnitPrice, uint256 _faaSDuration, uint256 _biddingDuration) returns(uint256, address)
func (_Market *MarketSession) NewDeploymentOrder(_faaSLevelID *big.Int, _highestUnitPrice *big.Int, _faaSDuration *big.Int, _biddingDuration *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Market.Contract.NewDeploymentOrder(&_Market.TransactOpts, _faaSLevelID, _highestUnitPrice, _faaSDuration, _biddingDuration)
}

func (_Market *MarketSession) AsyncNewDeploymentOrder(handler func(*types.Receipt, error), _faaSLevelID *big.Int, _highestUnitPrice *big.Int, _faaSDuration *big.Int, _biddingDuration *big.Int) (*types.Transaction, error) {
	return _Market.Contract.AsyncNewDeploymentOrder(handler, &_Market.TransactOpts, _faaSLevelID, _highestUnitPrice, _faaSDuration, _biddingDuration)
}

// NewDeploymentOrder is a paid mutator transaction binding the contract method 0xdbd4314f.
//
// Solidity: function newDeploymentOrder(uint256 _faaSLevelID, uint256 _highestUnitPrice, uint256 _faaSDuration, uint256 _biddingDuration) returns(uint256, address)
func (_Market *MarketTransactorSession) NewDeploymentOrder(_faaSLevelID *big.Int, _highestUnitPrice *big.Int, _faaSDuration *big.Int, _biddingDuration *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Market.Contract.NewDeploymentOrder(&_Market.TransactOpts, _faaSLevelID, _highestUnitPrice, _faaSDuration, _biddingDuration)
}

func (_Market *MarketTransactorSession) AsyncNewDeploymentOrder(handler func(*types.Receipt, error), _faaSLevelID *big.Int, _highestUnitPrice *big.Int, _faaSDuration *big.Int, _biddingDuration *big.Int) (*types.Transaction, error) {
	return _Market.Contract.AsyncNewDeploymentOrder(handler, &_Market.TransactOpts, _faaSLevelID, _highestUnitPrice, _faaSDuration, _biddingDuration)
}

// NewLease is a paid mutator transaction binding the contract method 0x4192b663.
//
// Solidity: function newLease(uint256 _delpoymentOrderID, address _provider, uint256 _unitPrice) returns(uint256)
func (_Market *MarketTransactor) NewLease(opts *bind.TransactOpts, _delpoymentOrderID *big.Int, _provider common.Address, _unitPrice *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Market.contract.Transact(opts, "newLease", _delpoymentOrderID, _provider, _unitPrice)
}

func (_Market *MarketTransactor) AsyncNewLease(handler func(*types.Receipt, error), opts *bind.TransactOpts, _delpoymentOrderID *big.Int, _provider common.Address, _unitPrice *big.Int) (*types.Transaction, error) {
	return _Market.contract.AsyncTransact(opts, handler, "newLease", _delpoymentOrderID, _provider, _unitPrice)
}

// NewLease is a paid mutator transaction binding the contract method 0x4192b663.
//
// Solidity: function newLease(uint256 _delpoymentOrderID, address _provider, uint256 _unitPrice) returns(uint256)
func (_Market *MarketSession) NewLease(_delpoymentOrderID *big.Int, _provider common.Address, _unitPrice *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Market.Contract.NewLease(&_Market.TransactOpts, _delpoymentOrderID, _provider, _unitPrice)
}

func (_Market *MarketSession) AsyncNewLease(handler func(*types.Receipt, error), _delpoymentOrderID *big.Int, _provider common.Address, _unitPrice *big.Int) (*types.Transaction, error) {
	return _Market.Contract.AsyncNewLease(handler, &_Market.TransactOpts, _delpoymentOrderID, _provider, _unitPrice)
}

// NewLease is a paid mutator transaction binding the contract method 0x4192b663.
//
// Solidity: function newLease(uint256 _delpoymentOrderID, address _provider, uint256 _unitPrice) returns(uint256)
func (_Market *MarketTransactorSession) NewLease(_delpoymentOrderID *big.Int, _provider common.Address, _unitPrice *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Market.Contract.NewLease(&_Market.TransactOpts, _delpoymentOrderID, _provider, _unitPrice)
}

func (_Market *MarketTransactorSession) AsyncNewLease(handler func(*types.Receipt, error), _delpoymentOrderID *big.Int, _provider common.Address, _unitPrice *big.Int) (*types.Transaction, error) {
	return _Market.Contract.AsyncNewLease(handler, &_Market.TransactOpts, _delpoymentOrderID, _provider, _unitPrice)
}

// ProviderLogin is a paid mutator transaction binding the contract method 0x2bf35570.
//
// Solidity: function providerLogin() returns()
func (_Market *MarketTransactor) ProviderLogin(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Market.contract.Transact(opts, "providerLogin")
}

func (_Market *MarketTransactor) AsyncProviderLogin(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.contract.AsyncTransact(opts, handler, "providerLogin")
}

// ProviderLogin is a paid mutator transaction binding the contract method 0x2bf35570.
//
// Solidity: function providerLogin() returns()
func (_Market *MarketSession) ProviderLogin() (*types.Transaction, *types.Receipt, error) {
	return _Market.Contract.ProviderLogin(&_Market.TransactOpts)
}

func (_Market *MarketSession) AsyncProviderLogin(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _Market.Contract.AsyncProviderLogin(handler, &_Market.TransactOpts)
}

// ProviderLogin is a paid mutator transaction binding the contract method 0x2bf35570.
//
// Solidity: function providerLogin() returns()
func (_Market *MarketTransactorSession) ProviderLogin() (*types.Transaction, *types.Receipt, error) {
	return _Market.Contract.ProviderLogin(&_Market.TransactOpts)
}

func (_Market *MarketTransactorSession) AsyncProviderLogin(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _Market.Contract.AsyncProviderLogin(handler, &_Market.TransactOpts)
}

// ProviderLogout is a paid mutator transaction binding the contract method 0x884eacb8.
//
// Solidity: function providerLogout() returns()
func (_Market *MarketTransactor) ProviderLogout(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Market.contract.Transact(opts, "providerLogout")
}

func (_Market *MarketTransactor) AsyncProviderLogout(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.contract.AsyncTransact(opts, handler, "providerLogout")
}

// ProviderLogout is a paid mutator transaction binding the contract method 0x884eacb8.
//
// Solidity: function providerLogout() returns()
func (_Market *MarketSession) ProviderLogout() (*types.Transaction, *types.Receipt, error) {
	return _Market.Contract.ProviderLogout(&_Market.TransactOpts)
}

func (_Market *MarketSession) AsyncProviderLogout(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _Market.Contract.AsyncProviderLogout(handler, &_Market.TransactOpts)
}

// ProviderLogout is a paid mutator transaction binding the contract method 0x884eacb8.
//
// Solidity: function providerLogout() returns()
func (_Market *MarketTransactorSession) ProviderLogout() (*types.Transaction, *types.Receipt, error) {
	return _Market.Contract.ProviderLogout(&_Market.TransactOpts)
}

func (_Market *MarketTransactorSession) AsyncProviderLogout(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _Market.Contract.AsyncProviderLogout(handler, &_Market.TransactOpts)
}

// MarketNewDeploymentOrderEventIterator is returned from FilterNewDeploymentOrderEvent and is used to iterate over the raw logs and unpacked data for NewDeploymentOrderEvent events raised by the Market contract.
type MarketNewDeploymentOrderEventIterator struct {
	Event *MarketNewDeploymentOrderEvent // Event containing the contract specifics and raw log

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
func (it *MarketNewDeploymentOrderEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketNewDeploymentOrderEvent)
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
		it.Event = new(MarketNewDeploymentOrderEvent)
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
func (it *MarketNewDeploymentOrderEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketNewDeploymentOrderEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketNewDeploymentOrderEvent represents a NewDeploymentOrderEvent event raised by the Market contract.
type MarketNewDeploymentOrderEvent struct {
	DelpoymentOrderID *big.Int
	AuctionAddress    common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewDeploymentOrderEvent is a free log retrieval operation binding the contract event 0x93b403b917e94d112528f49b59221bab32fb8fd897cbc154bf092ef7796d5dd0.
//
// Solidity: event newDeploymentOrderEvent(uint256 _delpoymentOrderID, address _auctionAddress)
func (_Market *MarketFilterer) FilterNewDeploymentOrderEvent(opts *bind.FilterOpts) (*MarketNewDeploymentOrderEventIterator, error) {

	logs, sub, err := _Market.contract.FilterLogs(opts, "newDeploymentOrderEvent")
	if err != nil {
		return nil, err
	}
	return &MarketNewDeploymentOrderEventIterator{contract: _Market.contract, event: "newDeploymentOrderEvent", logs: logs, sub: sub}, nil
}

// WatchNewDeploymentOrderEvent is a free log subscription operation binding the contract event 0x93b403b917e94d112528f49b59221bab32fb8fd897cbc154bf092ef7796d5dd0.
//
// Solidity: event newDeploymentOrderEvent(uint256 _delpoymentOrderID, address _auctionAddress)
func (_Market *MarketFilterer) WatchNewDeploymentOrderEvent(opts *bind.WatchOpts, sink chan<- *MarketNewDeploymentOrderEvent) (event.Subscription, error) {

	logs, sub, err := _Market.contract.WatchLogs(opts, "newDeploymentOrderEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketNewDeploymentOrderEvent)
				if err := _Market.contract.UnpackLog(event, "newDeploymentOrderEvent", log); err != nil {
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

// ParseNewDeploymentOrderEvent is a log parse operation binding the contract event 0x93b403b917e94d112528f49b59221bab32fb8fd897cbc154bf092ef7796d5dd0.
//
// Solidity: event newDeploymentOrderEvent(uint256 _delpoymentOrderID, address _auctionAddress)
func (_Market *MarketFilterer) ParseNewDeploymentOrderEvent(log types.Log) (*MarketNewDeploymentOrderEvent, error) {
	event := new(MarketNewDeploymentOrderEvent)
	if err := _Market.contract.UnpackLog(event, "newDeploymentOrderEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}
