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
const MarketABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContractAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_delpoymentOrderID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_unitPrice\",\"type\":\"uint256\"}],\"name\":\"auctionEndEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_delpoymentOrderID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_auctionAddress\",\"type\":\"address\"}],\"name\":\"newDeploymentOrderEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_customer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_leaseID\",\"type\":\"uint256\"}],\"name\":\"newLeaseEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctions\",\"outputs\":[{\"internalType\":\"contractSimpleAuction\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_delpoymentOrderID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_unitPrice\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"delpoymentOrders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"customer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"faaSLevelID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"faaSDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highestUnitPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_highestUnitPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_faaSDuration\",\"type\":\"uint256\"}],\"name\":\"getDelpoymentOrderLockFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_delpoymentOrderID\",\"type\":\"uint256\"}],\"name\":\"getDeploymentOrder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProviderDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"getProviderReputation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProviderReputationQualified\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"isProviderQualified\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"isProviderRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"leases\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"customer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"faaSLevelID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"faaSDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unitPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"providerServiceFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"customerWithdrawFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"witnessFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maintainerFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_delpoymentOrderID\",\"type\":\"uint256\"}],\"name\":\"matchDeploymentOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_faaSLevelID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_highestUnitPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_faaSDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_biddingDuration\",\"type\":\"uint256\"}],\"name\":\"newDeploymentOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numDelpoymentOrders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerLogin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerLogout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerReputationInit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerReputationQualified\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_delpoymentOrderID\",\"type\":\"uint256\"}],\"name\":\"settleDeploymemtOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MarketBin is the compiled bytecode used for deploying new contracts.
var MarketBin = "0x608060405234801561001057600080fd5b506040516123a43803806123a48339818101604052602081101561003357600080fd5b8101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506064600181905550600560038190555060056004819055506001600681905550506122f0806100b46000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c8063757f4a97116100ad578063d2d615ad11610071578063d2d615ad14610543578063dbd4314f1461058f578063dbeefdd214610622578063f0611e3814610640578063f843dcc41461065e5761012c565b8063757f4a971461036a578063884eacb8146103885780638927a10614610392578063b829647114610464578063ce67ca03146104c05761012c565b8063230544a7116100f4578063230544a7146102445780632bf3557014610262578063571a26a01461026c578063598647f8146102da578063729d98e9146103125761012c565b80630626399314610131578063106647241461014f57806318f645fa1461017d5780631ead914e146101ca5780632152189214610226575b600080fd5b6101396106e1565b6040518082815260200191505060405180910390f35b61017b6004803603602081101561016557600080fd5b81019080803590602001909291905050506106e7565b005b6101a96004803603602081101561019357600080fd5b81019080803590602001909291905050506106ea565b60405180831515151581526020018281526020019250505060405180910390f35b61020c600480360360208110156101e057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610886565b604051808215151515815260200191505060405180910390f35b61022e6108d1565b6040518082815260200191505060405180910390f35b61024c6108d7565b6040518082815260200191505060405180910390f35b61026a6108dd565b005b6102986004803603602081101561028257600080fd5b8101908080359060200190929190505050610b45565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610310600480360360408110156102f057600080fd5b810190808035906020019092919080359060200190929190505050610b78565b005b6103546004803603602081101561032857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d11565b6040518082815260200191505060405180910390f35b610372610d5a565b6040518082815260200191505060405180910390f35b610390610d64565b005b6103be600480360360208110156103a857600080fd5b8101908080359060200190929190505050611037565b604051808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001888152602001878152602001868152602001858152602001848152602001838152602001828152602001995050505050505050505060405180910390f35b6104a66004803603602081101561047a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506110c5565b604051808215151515815260200191505060405180910390f35b6104ec600480360360208110156104d657600080fd5b8101908080359060200190929190505050611113565b604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200184815260200183815260200182815260200194505050505060405180910390f35b6105796004803603604081101561055957600080fd5b8101908080359060200190929190803590602001909291905050506111d9565b6040518082815260200191505060405180910390f35b6105d9600480360360808110156105a557600080fd5b81019080803590602001909291908035906020019092919080359060200190929190803590602001909291905050506111e6565b604051808381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390f35b61062a61157c565b6040518082815260200191505060405180910390f35b610648611586565b6040518082815260200191505060405180910390f35b61068a6004803603602081101561067457600080fd5b810190808035906020019092919050505061158c565b604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200184815260200183815260200182815260200194505050505060405180910390f35b60065481565b50565b60008060008060006008600087815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632a24f46c6040518163ffffffff1660e01b8152600401606060405180830381600087803b15801561076d57600080fd5b505af1158015610781573d6000803e3d6000fd5b505050506040513d606081101561079757600080fd5b810190808051906020019092919080519060200190929190805190602001909291905050508093508194508295505050507ff19639f9d5b84e723a190dcd994b11629429fd9d86ec12313709c578534eb03f8684848460405180858152602001841515151581526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200194505050505060405180910390a16000151583151514156108665760008080905094509450505050610881565b60006108738784846115dc565b905060018195509550505050505b915091565b600080600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054119050919050565b60045481565b60015481565b600015156108ea33610886565b151514610942576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260278152602001806121ee6027913960400191505060405180910390fd5b600115156000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd33306001546040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b158015610a2457600080fd5b505af1158015610a38573d6000803e3d6000fd5b505050506040513d6020811015610a4e57600080fd5b8101908080519060200190929190505050151514610ab7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260288152602001806121c66028913960400191505060405180910390fd5b600154600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600354600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550565b60086020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60011515610b8533610886565b151514610bdd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602b81526020018061223a602b913960400191505060405180910390fd5b610be6336110c5565b610c3b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806122156025913960400191505060405180910390fd5b6008600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166359d667a533836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015610cf557600080fd5b505af1158015610d09573d6000803e3d6000fd5b505050505050565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000600454905090565b60011515610d7133610886565b151514610dc9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602b81526020018061223a602b913960400191505060405180910390fd5b610dd2336110c5565b610e27576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806122156025913960400191505060405180910390fd5b6000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600115156000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33846040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015610fa157600080fd5b505af1158015610fb5573d6000803e3d6000fd5b505050506040513d6020811015610fcb57600080fd5b8101908080519060200190929190505050151514611034576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602d81526020018061228e602d913960400191505060405180910390fd5b50565b60096020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154908060030154908060040154908060050154908060060154908060070154908060080154905089565b6000600454600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410159050919050565b6000806000806111216118b4565b600760008781526020019081526020016000206040518060800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820154815260200160028201548152602001600382015481525050905080600001518160200151826040015183606001519450945094509450509193509193565b6000818302905092915050565b60008060006111f586866111d9565b9050600115156000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b1580156112d757600080fd5b505af11580156112eb573d6000803e3d6000fd5b505050506040513d602081101561130157600080fd5b810190808051906020019092919050505015151461136a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001806122656029913960400191505060405180910390fd5b60006006600081548092919060010191905055905060405180608001604052803373ffffffffffffffffffffffffffffffffffffffff168152602001878152602001878152602001888152506007600083815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101556040820151816002015560608201518160030155905050600030828988604051611443906118f2565b808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001848152602001838152602001828152602001945050505050604051809103906000f0801580156114aa573d6000803e3d6000fd5b509050806008600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f93b403b917e94d112528f49b59221bab32fb8fd897cbc154bf092ef7796d5dd08282604051808381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a181819450945050505094509492505050565b6000600154905090565b60035481565b60076020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154908060030154905084565b6000808490506115ea6118b4565b600760008781526020019081526020016000206040518060800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820154815260200160028201548152602001600382015481525050905060008482604001510290506000819050600080905060008090506000809050604051806101200160405280876000015173ffffffffffffffffffffffffffffffffffffffff1681526020018b73ffffffffffffffffffffffffffffffffffffffff16815260200187602001518152602001876040015181526020018a8152602001858152602001848152602001838152602001828152506009600089815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e0820151816007015561010082015181600801559050507f42ecce389f5f2ecbd7ab89ee93ed8201c66dce8685258d5b48bb6b2017e46f0986600001518b89604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a1869750505050505050509392505050565b6040518060800160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001600081525090565b6108c6806119008339019056fe608060405234801561001057600080fd5b506040516108c63803806108c68339818101604052608081101561003357600080fd5b8101908080519060200190929190805190602001909291908051906020019092919080519060200190929190505050836000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008060146101000a81548160ff021916908360018111156100c057fe5b021790555082600181905550816002819055504260038190555080600481905550816005819055506000600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050506107898061013d6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80635d06978f1161005b5780635d06978f1461014a57806361dcd7ab14610168578063c19d93fb14610186578063c8e7b31d146101b25761007d565b80632a24f46c146100825780634840f2f0146100de57806359d667a5146100fc575b600080fd5b61008a6101d0565b60405180841515151581526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390f35b6100e6610325565b6040518082815260200191505060405180910390f35b6101486004803603604081101561011257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061032b565b005b610152610599565b6040518082815260200191505060405180910390f35b61017061059f565b6040518082815260200191505060405180910390f35b61018e6105a5565b6040518082600181111561019e57fe5b60ff16815260200191505060405180910390f35b6101ba6105b8565b6040518082815260200191505060405180910390f35b60008060008060018111156101e157fe5b600060149054906101000a900460ff1660018111156101fc57fe5b14801561020e57506004546003540142115b1561021c5761021b6105be565b5b600180600181111561022a57fe5b600060149054906101000a900460ff16600181111561024557fe5b1461029b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603681526020018061071e6036913960400191505060405180910390fd5b60008073ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415905080600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166005549450945094505050909192565b60025481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103d0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603b8152602001806106e3603b913960400191505060405180910390fd5b600060018111156103dd57fe5b600060149054906101000a900460ff1660018111156103f857fe5b14801561040a57506004546003540142115b15610418576104176105be565b5b600080600181111561042657fe5b600060149054906101000a900460ff16600181111561044157fe5b14610497576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603681526020018061071e6036913960400191505060405180910390fd5b6002548211156104f2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252604a815260200180610699604a913960600191505060405180910390fd5b600554821061054c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260338152602001806106666033913960400191505060405180910390fd5b8160058190555082600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b60015481565b60035481565b600060149054906101000a900460ff1681565b60045481565b600060018111156105cb57fe5b600060149054906101000a900460ff1660018111156105e657fe5b1415610611576001600060146101000a81548160ff0219169083600181111561060b57fe5b02179055505b60018081111561061d57fe5b600060149054906101000a900460ff16600181111561063857fe5b1415610663576001600060146101000a81548160ff0219169083600181111561065d57fe5b02179055505b56fe53696d706c6541756374696f6e3a20746865726520616c72656164792069732061206c6f7765722075696e742d70726963652e53696d706c6541756374696f6e3a2074686520756e69742d707269636520697320686967686572207468616e20686967686573742074686520637573746f6d657220616363657074656453696d706c6541756374696f6e3a206f6e6c79206d61726b657420636f6e6374726163742063616e2063616c6c20746869732066756e6369746f6e53696d706c6541756374696f6e3a2066756e6374696f6e2063616e6e6f742062652063616c6c65642061742074686973207374617465a264697066735822122037c5104697dc4bd98622042da04bb857fbf968fb41ed542054430086eca1f0e964736f6c634300060a00334d61726b65743a206661696c656420746f207061792061207265676973746572206465706f7369744d61726b65743a20746865206164647265737320686164206265656e20726567697374657265644d61726b65743a207468652070726f7669646572206973206e6f74207175616c69666965644d61726b65743a20746865206164647265737320686164206e6f74206265656e20726567697374657265646c6f636b20666565206265666f7265206372656174696e67206465706c6f796d656e74206f726465724d61726b65743a206661696c656420746f20726566756e64207468652070726f7669646572206465706f736974a264697066735822122011663aa03c2f34cbd9efe9e6217dc352c3ad031f4600ebf4eecf6b7202f76be864736f6c634300060a0033"

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

// GetDelpoymentOrderLockFee is a free data retrieval call binding the contract method 0xd2d615ad.
//
// Solidity: function getDelpoymentOrderLockFee(uint256 _highestUnitPrice, uint256 _faaSDuration) constant returns(uint256)
func (_Market *MarketCaller) GetDelpoymentOrderLockFee(opts *bind.CallOpts, _highestUnitPrice *big.Int, _faaSDuration *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "getDelpoymentOrderLockFee", _highestUnitPrice, _faaSDuration)
	return *ret0, err
}

// GetDelpoymentOrderLockFee is a free data retrieval call binding the contract method 0xd2d615ad.
//
// Solidity: function getDelpoymentOrderLockFee(uint256 _highestUnitPrice, uint256 _faaSDuration) constant returns(uint256)
func (_Market *MarketSession) GetDelpoymentOrderLockFee(_highestUnitPrice *big.Int, _faaSDuration *big.Int) (*big.Int, error) {
	return _Market.Contract.GetDelpoymentOrderLockFee(&_Market.CallOpts, _highestUnitPrice, _faaSDuration)
}

// GetDelpoymentOrderLockFee is a free data retrieval call binding the contract method 0xd2d615ad.
//
// Solidity: function getDelpoymentOrderLockFee(uint256 _highestUnitPrice, uint256 _faaSDuration) constant returns(uint256)
func (_Market *MarketCallerSession) GetDelpoymentOrderLockFee(_highestUnitPrice *big.Int, _faaSDuration *big.Int) (*big.Int, error) {
	return _Market.Contract.GetDelpoymentOrderLockFee(&_Market.CallOpts, _highestUnitPrice, _faaSDuration)
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

// IsProviderQualified is a free data retrieval call binding the contract method 0xb8296471.
//
// Solidity: function isProviderQualified(address provider) constant returns(bool)
func (_Market *MarketCaller) IsProviderQualified(opts *bind.CallOpts, provider common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "isProviderQualified", provider)
	return *ret0, err
}

// IsProviderQualified is a free data retrieval call binding the contract method 0xb8296471.
//
// Solidity: function isProviderQualified(address provider) constant returns(bool)
func (_Market *MarketSession) IsProviderQualified(provider common.Address) (bool, error) {
	return _Market.Contract.IsProviderQualified(&_Market.CallOpts, provider)
}

// IsProviderQualified is a free data retrieval call binding the contract method 0xb8296471.
//
// Solidity: function isProviderQualified(address provider) constant returns(bool)
func (_Market *MarketCallerSession) IsProviderQualified(provider common.Address) (bool, error) {
	return _Market.Contract.IsProviderQualified(&_Market.CallOpts, provider)
}

// IsProviderRegistered is a free data retrieval call binding the contract method 0x1ead914e.
//
// Solidity: function isProviderRegistered(address provider) constant returns(bool)
func (_Market *MarketCaller) IsProviderRegistered(opts *bind.CallOpts, provider common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Market.contract.Call(opts, out, "isProviderRegistered", provider)
	return *ret0, err
}

// IsProviderRegistered is a free data retrieval call binding the contract method 0x1ead914e.
//
// Solidity: function isProviderRegistered(address provider) constant returns(bool)
func (_Market *MarketSession) IsProviderRegistered(provider common.Address) (bool, error) {
	return _Market.Contract.IsProviderRegistered(&_Market.CallOpts, provider)
}

// IsProviderRegistered is a free data retrieval call binding the contract method 0x1ead914e.
//
// Solidity: function isProviderRegistered(address provider) constant returns(bool)
func (_Market *MarketCallerSession) IsProviderRegistered(provider common.Address) (bool, error) {
	return _Market.Contract.IsProviderRegistered(&_Market.CallOpts, provider)
}

// Leases is a free data retrieval call binding the contract method 0x8927a106.
//
// Solidity: function leases(uint256 ) constant returns(address customer, address provider, uint256 faaSLevelID, uint256 faaSDuration, uint256 unitPrice, uint256 providerServiceFee, uint256 customerWithdrawFee, uint256 witnessFee, uint256 maintainerFee)
func (_Market *MarketCaller) Leases(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Customer            common.Address
	Provider            common.Address
	FaaSLevelID         *big.Int
	FaaSDuration        *big.Int
	UnitPrice           *big.Int
	ProviderServiceFee  *big.Int
	CustomerWithdrawFee *big.Int
	WitnessFee          *big.Int
	MaintainerFee       *big.Int
}, error) {
	ret := new(struct {
		Customer            common.Address
		Provider            common.Address
		FaaSLevelID         *big.Int
		FaaSDuration        *big.Int
		UnitPrice           *big.Int
		ProviderServiceFee  *big.Int
		CustomerWithdrawFee *big.Int
		WitnessFee          *big.Int
		MaintainerFee       *big.Int
	})
	out := ret
	err := _Market.contract.Call(opts, out, "leases", arg0)
	return *ret, err
}

// Leases is a free data retrieval call binding the contract method 0x8927a106.
//
// Solidity: function leases(uint256 ) constant returns(address customer, address provider, uint256 faaSLevelID, uint256 faaSDuration, uint256 unitPrice, uint256 providerServiceFee, uint256 customerWithdrawFee, uint256 witnessFee, uint256 maintainerFee)
func (_Market *MarketSession) Leases(arg0 *big.Int) (struct {
	Customer            common.Address
	Provider            common.Address
	FaaSLevelID         *big.Int
	FaaSDuration        *big.Int
	UnitPrice           *big.Int
	ProviderServiceFee  *big.Int
	CustomerWithdrawFee *big.Int
	WitnessFee          *big.Int
	MaintainerFee       *big.Int
}, error) {
	return _Market.Contract.Leases(&_Market.CallOpts, arg0)
}

// Leases is a free data retrieval call binding the contract method 0x8927a106.
//
// Solidity: function leases(uint256 ) constant returns(address customer, address provider, uint256 faaSLevelID, uint256 faaSDuration, uint256 unitPrice, uint256 providerServiceFee, uint256 customerWithdrawFee, uint256 witnessFee, uint256 maintainerFee)
func (_Market *MarketCallerSession) Leases(arg0 *big.Int) (struct {
	Customer            common.Address
	Provider            common.Address
	FaaSLevelID         *big.Int
	FaaSDuration        *big.Int
	UnitPrice           *big.Int
	ProviderServiceFee  *big.Int
	CustomerWithdrawFee *big.Int
	WitnessFee          *big.Int
	MaintainerFee       *big.Int
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

// Bid is a paid mutator transaction binding the contract method 0x598647f8.
//
// Solidity: function bid(uint256 _delpoymentOrderID, uint256 _unitPrice) returns()
func (_Market *MarketTransactor) Bid(opts *bind.TransactOpts, _delpoymentOrderID *big.Int, _unitPrice *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Market.contract.Transact(opts, "bid", _delpoymentOrderID, _unitPrice)
}

func (_Market *MarketTransactor) AsyncBid(handler func(*types.Receipt, error), opts *bind.TransactOpts, _delpoymentOrderID *big.Int, _unitPrice *big.Int) (*types.Transaction, error) {
	return _Market.contract.AsyncTransact(opts, handler, "bid", _delpoymentOrderID, _unitPrice)
}

// Bid is a paid mutator transaction binding the contract method 0x598647f8.
//
// Solidity: function bid(uint256 _delpoymentOrderID, uint256 _unitPrice) returns()
func (_Market *MarketSession) Bid(_delpoymentOrderID *big.Int, _unitPrice *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Market.Contract.Bid(&_Market.TransactOpts, _delpoymentOrderID, _unitPrice)
}

func (_Market *MarketSession) AsyncBid(handler func(*types.Receipt, error), _delpoymentOrderID *big.Int, _unitPrice *big.Int) (*types.Transaction, error) {
	return _Market.Contract.AsyncBid(handler, &_Market.TransactOpts, _delpoymentOrderID, _unitPrice)
}

// Bid is a paid mutator transaction binding the contract method 0x598647f8.
//
// Solidity: function bid(uint256 _delpoymentOrderID, uint256 _unitPrice) returns()
func (_Market *MarketTransactorSession) Bid(_delpoymentOrderID *big.Int, _unitPrice *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Market.Contract.Bid(&_Market.TransactOpts, _delpoymentOrderID, _unitPrice)
}

func (_Market *MarketTransactorSession) AsyncBid(handler func(*types.Receipt, error), _delpoymentOrderID *big.Int, _unitPrice *big.Int) (*types.Transaction, error) {
	return _Market.Contract.AsyncBid(handler, &_Market.TransactOpts, _delpoymentOrderID, _unitPrice)
}

// MatchDeploymentOrder is a paid mutator transaction binding the contract method 0x18f645fa.
//
// Solidity: function matchDeploymentOrder(uint256 _delpoymentOrderID) returns(bool, uint256)
func (_Market *MarketTransactor) MatchDeploymentOrder(opts *bind.TransactOpts, _delpoymentOrderID *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Market.contract.Transact(opts, "matchDeploymentOrder", _delpoymentOrderID)
}

func (_Market *MarketTransactor) AsyncMatchDeploymentOrder(handler func(*types.Receipt, error), opts *bind.TransactOpts, _delpoymentOrderID *big.Int) (*types.Transaction, error) {
	return _Market.contract.AsyncTransact(opts, handler, "matchDeploymentOrder", _delpoymentOrderID)
}

// MatchDeploymentOrder is a paid mutator transaction binding the contract method 0x18f645fa.
//
// Solidity: function matchDeploymentOrder(uint256 _delpoymentOrderID) returns(bool, uint256)
func (_Market *MarketSession) MatchDeploymentOrder(_delpoymentOrderID *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Market.Contract.MatchDeploymentOrder(&_Market.TransactOpts, _delpoymentOrderID)
}

func (_Market *MarketSession) AsyncMatchDeploymentOrder(handler func(*types.Receipt, error), _delpoymentOrderID *big.Int) (*types.Transaction, error) {
	return _Market.Contract.AsyncMatchDeploymentOrder(handler, &_Market.TransactOpts, _delpoymentOrderID)
}

// MatchDeploymentOrder is a paid mutator transaction binding the contract method 0x18f645fa.
//
// Solidity: function matchDeploymentOrder(uint256 _delpoymentOrderID) returns(bool, uint256)
func (_Market *MarketTransactorSession) MatchDeploymentOrder(_delpoymentOrderID *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Market.Contract.MatchDeploymentOrder(&_Market.TransactOpts, _delpoymentOrderID)
}

func (_Market *MarketTransactorSession) AsyncMatchDeploymentOrder(handler func(*types.Receipt, error), _delpoymentOrderID *big.Int) (*types.Transaction, error) {
	return _Market.Contract.AsyncMatchDeploymentOrder(handler, &_Market.TransactOpts, _delpoymentOrderID)
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

// SettleDeploymemtOrder is a paid mutator transaction binding the contract method 0x10664724.
//
// Solidity: function settleDeploymemtOrder(uint256 _delpoymentOrderID) returns()
func (_Market *MarketTransactor) SettleDeploymemtOrder(opts *bind.TransactOpts, _delpoymentOrderID *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Market.contract.Transact(opts, "settleDeploymemtOrder", _delpoymentOrderID)
}

func (_Market *MarketTransactor) AsyncSettleDeploymemtOrder(handler func(*types.Receipt, error), opts *bind.TransactOpts, _delpoymentOrderID *big.Int) (*types.Transaction, error) {
	return _Market.contract.AsyncTransact(opts, handler, "settleDeploymemtOrder", _delpoymentOrderID)
}

// SettleDeploymemtOrder is a paid mutator transaction binding the contract method 0x10664724.
//
// Solidity: function settleDeploymemtOrder(uint256 _delpoymentOrderID) returns()
func (_Market *MarketSession) SettleDeploymemtOrder(_delpoymentOrderID *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Market.Contract.SettleDeploymemtOrder(&_Market.TransactOpts, _delpoymentOrderID)
}

func (_Market *MarketSession) AsyncSettleDeploymemtOrder(handler func(*types.Receipt, error), _delpoymentOrderID *big.Int) (*types.Transaction, error) {
	return _Market.Contract.AsyncSettleDeploymemtOrder(handler, &_Market.TransactOpts, _delpoymentOrderID)
}

// SettleDeploymemtOrder is a paid mutator transaction binding the contract method 0x10664724.
//
// Solidity: function settleDeploymemtOrder(uint256 _delpoymentOrderID) returns()
func (_Market *MarketTransactorSession) SettleDeploymemtOrder(_delpoymentOrderID *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Market.Contract.SettleDeploymemtOrder(&_Market.TransactOpts, _delpoymentOrderID)
}

func (_Market *MarketTransactorSession) AsyncSettleDeploymemtOrder(handler func(*types.Receipt, error), _delpoymentOrderID *big.Int) (*types.Transaction, error) {
	return _Market.Contract.AsyncSettleDeploymemtOrder(handler, &_Market.TransactOpts, _delpoymentOrderID)
}

// MarketAuctionEndEventIterator is returned from FilterAuctionEndEvent and is used to iterate over the raw logs and unpacked data for AuctionEndEvent events raised by the Market contract.
type MarketAuctionEndEventIterator struct {
	Event *MarketAuctionEndEvent // Event containing the contract specifics and raw log

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
func (it *MarketAuctionEndEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketAuctionEndEvent)
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
		it.Event = new(MarketAuctionEndEvent)
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
func (it *MarketAuctionEndEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketAuctionEndEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketAuctionEndEvent represents a AuctionEndEvent event raised by the Market contract.
type MarketAuctionEndEvent struct {
	DelpoymentOrderID *big.Int
	Success           bool
	Provider          common.Address
	UnitPrice         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAuctionEndEvent is a free log retrieval operation binding the contract event 0xf19639f9d5b84e723a190dcd994b11629429fd9d86ec12313709c578534eb03f.
//
// Solidity: event auctionEndEvent(uint256 _delpoymentOrderID, bool _success, address _provider, uint256 _unitPrice)
func (_Market *MarketFilterer) FilterAuctionEndEvent(opts *bind.FilterOpts) (*MarketAuctionEndEventIterator, error) {

	logs, sub, err := _Market.contract.FilterLogs(opts, "auctionEndEvent")
	if err != nil {
		return nil, err
	}
	return &MarketAuctionEndEventIterator{contract: _Market.contract, event: "auctionEndEvent", logs: logs, sub: sub}, nil
}

// WatchAuctionEndEvent is a free log subscription operation binding the contract event 0xf19639f9d5b84e723a190dcd994b11629429fd9d86ec12313709c578534eb03f.
//
// Solidity: event auctionEndEvent(uint256 _delpoymentOrderID, bool _success, address _provider, uint256 _unitPrice)
func (_Market *MarketFilterer) WatchAuctionEndEvent(opts *bind.WatchOpts, sink chan<- *MarketAuctionEndEvent) (event.Subscription, error) {

	logs, sub, err := _Market.contract.WatchLogs(opts, "auctionEndEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketAuctionEndEvent)
				if err := _Market.contract.UnpackLog(event, "auctionEndEvent", log); err != nil {
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

// ParseAuctionEndEvent is a log parse operation binding the contract event 0xf19639f9d5b84e723a190dcd994b11629429fd9d86ec12313709c578534eb03f.
//
// Solidity: event auctionEndEvent(uint256 _delpoymentOrderID, bool _success, address _provider, uint256 _unitPrice)
func (_Market *MarketFilterer) ParseAuctionEndEvent(log types.Log) (*MarketAuctionEndEvent, error) {
	event := new(MarketAuctionEndEvent)
	if err := _Market.contract.UnpackLog(event, "auctionEndEvent", log); err != nil {
		return nil, err
	}
	return event, nil
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

// MarketNewLeaseEventIterator is returned from FilterNewLeaseEvent and is used to iterate over the raw logs and unpacked data for NewLeaseEvent events raised by the Market contract.
type MarketNewLeaseEventIterator struct {
	Event *MarketNewLeaseEvent // Event containing the contract specifics and raw log

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
func (it *MarketNewLeaseEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketNewLeaseEvent)
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
		it.Event = new(MarketNewLeaseEvent)
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
func (it *MarketNewLeaseEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketNewLeaseEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketNewLeaseEvent represents a NewLeaseEvent event raised by the Market contract.
type MarketNewLeaseEvent struct {
	Customer common.Address
	Provider common.Address
	LeaseID  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewLeaseEvent is a free log retrieval operation binding the contract event 0x42ecce389f5f2ecbd7ab89ee93ed8201c66dce8685258d5b48bb6b2017e46f09.
//
// Solidity: event newLeaseEvent(address _customer, address _provider, uint256 _leaseID)
func (_Market *MarketFilterer) FilterNewLeaseEvent(opts *bind.FilterOpts) (*MarketNewLeaseEventIterator, error) {

	logs, sub, err := _Market.contract.FilterLogs(opts, "newLeaseEvent")
	if err != nil {
		return nil, err
	}
	return &MarketNewLeaseEventIterator{contract: _Market.contract, event: "newLeaseEvent", logs: logs, sub: sub}, nil
}

// WatchNewLeaseEvent is a free log subscription operation binding the contract event 0x42ecce389f5f2ecbd7ab89ee93ed8201c66dce8685258d5b48bb6b2017e46f09.
//
// Solidity: event newLeaseEvent(address _customer, address _provider, uint256 _leaseID)
func (_Market *MarketFilterer) WatchNewLeaseEvent(opts *bind.WatchOpts, sink chan<- *MarketNewLeaseEvent) (event.Subscription, error) {

	logs, sub, err := _Market.contract.WatchLogs(opts, "newLeaseEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketNewLeaseEvent)
				if err := _Market.contract.UnpackLog(event, "newLeaseEvent", log); err != nil {
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

// ParseNewLeaseEvent is a log parse operation binding the contract event 0x42ecce389f5f2ecbd7ab89ee93ed8201c66dce8685258d5b48bb6b2017e46f09.
//
// Solidity: event newLeaseEvent(address _customer, address _provider, uint256 _leaseID)
func (_Market *MarketFilterer) ParseNewLeaseEvent(log types.Log) (*MarketNewLeaseEvent, error) {
	event := new(MarketNewLeaseEvent)
	if err := _Market.contract.UnpackLog(event, "newLeaseEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}
