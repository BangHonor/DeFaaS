// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package witnesspool

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

// WitnessPoolABI is the input ABI used to generate the binding from.
const WitnessPoolABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContractAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_witness\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_slaID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_monitoringBeginTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_monitoringDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_funcToMonitor\",\"type\":\"string\"}],\"name\":\"WitnessSelectedEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slaID\",\"type\":\"uint256\"},{\"internalType\":\"enumWitnessPool.SLAStates\",\"name\":\"_state\",\"type\":\"uint8\"}],\"name\":\"isAtSLAState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_witness\",\"type\":\"address\"},{\"internalType\":\"enumWitnessManagement.WStates\",\"name\":\"_state\",\"type\":\"uint8\"}],\"name\":\"isAtWState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slaID\",\"type\":\"uint256\"}],\"name\":\"isViolatedSLA\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slaID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_witness\",\"type\":\"address\"}],\"name\":\"isWitnessCommitteeMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_witness\",\"type\":\"address\"}],\"name\":\"isWitnessQualified\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_witness\",\"type\":\"address\"}],\"name\":\"isWitnessRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_curBlockNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_customer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_slaID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_funcToMonitor\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_monitoringDuration\",\"type\":\"uint256\"}],\"name\":\"newSLA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slaID\",\"type\":\"uint256\"}],\"name\":\"reportViolation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reputationDishonestReduced\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardNoviolationReport\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardNoviolationSilence\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardViolationReport\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardViolationSilence\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stdNumReportRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stdNumWitness\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stdWitnessDepoist\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stdblockNeed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenContract\",\"outputs\":[{\"internalType\":\"contractFaaSToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wintessTurnOn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"witenessLogout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"witnessDrawReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"witnessLogin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"witnessReputationInit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"witnessReputationQualified\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"witnessTurnOff\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// WitnessPoolBin is the compiled bytecode used for deploying new contracts.
var WitnessPoolBin = "0x6080604052600a6009556000600a556000600b556001600c556001600d556003600e556002600f55600260105534801561003857600080fd5b50604051612c60380380612c608339818101604052602081101561005b57600080fd5b81019080805190602001909291905050508080336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505060006007819055506064600281905550606460038190555060016004819055505050612b3f806101216000396000f3fe608060405234801561001057600080fd5b506004361061018e5760003560e01c80639faa79e3116100de578063e0a5a2aa11610097578063ecc7e64b11610071578063ecc7e64b1461064e578063ed06256714610658578063f11ad7f0146106be578063fc634e8f146106dc5761018e565b8063e0a5a2aa146105e0578063e81bc992146105fe578063ebe47110146106445761018e565b80639faa79e314610394578063a35da30c1461039e578063b65f6706146103bc578063cda28864146104d5578063cf61776214610528578063d446fcdd146105845761018e565b806355a373d61161014b5780637d635a7f116101255780637d635a7f146102f057806389640b581461030e57806389be6f461461032c5780638da5cb5b1461034a5761018e565b806355a373d61461026a578063744c2e33146102b4578063746ea9c6146102d25761018e565b806316c3734d146101935780632fba9d1a146101b157806337edd5a4146101bb5780633f5307c4146101d9578063426f93b1146101f757806347ff3aba14610201575b600080fd5b61019b61070a565b6040518082815260200191505060405180910390f35b6101b9610710565b005b6101c3610a40565b6040518082815260200191505060405180910390f35b6101e1610a46565b6040518082815260200191505060405180910390f35b6101ff610a4c565b005b6102506004803603604081101561021757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560ff169060200190929190505050610da0565b604051808215151515815260200191505060405180910390f35b610272610e12565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102bc610e38565b6040518082815260200191505060405180910390f35b6102da610e3e565b6040518082815260200191505060405180910390f35b6102f8610e44565b6040518082815260200191505060405180910390f35b610316610e4a565b6040518082815260200191505060405180910390f35b610334610e50565b6040518082815260200191505060405180910390f35b610352610e56565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61039c610e7b565b005b6103a66111b0565b6040518082815260200191505060405180910390f35b6104d3600480360360c08110156103d257600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561044357600080fd5b82018360208201111561045557600080fd5b8035906020019184600183028401116401000000008311171561047757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001909291905050506111b6565b005b61050e600480360360408110156104eb57600080fd5b8101908080359060200190929190803560ff16906020019092919050505061154f565b604051808215151515815260200191505060405180910390f35b61056a6004803603602081101561053e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611595565b604051808215151515815260200191505060405180910390f35b6105c66004803603602081101561059a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506115e6565b604051808215151515815260200191505060405180910390f35b6105e861163f565b6040518082815260200191505060405180910390f35b61062a6004803603602081101561061457600080fd5b8101908080359060200190929190505050611645565b604051808215151515815260200191505060405180910390f35b61064c6117f2565b005b61065661199f565b005b6106a46004803603604081101561066e57600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611ae7565b604051808215151515815260200191505060405180910390f35b6106c6611b5c565b6040518082815260200191505060405180910390f35b610708600480360360208110156106f257600080fd5b8101908080359060200190929190505050611b62565b005b60035481565b6000151561071d336115e6565b151514610775576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602c815260200180612994602c913960400191505060405180910390fd5b6001151561078233611595565b1515146107da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260278152602001806128b46027913960400191505060405180910390fd5b6000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206004015411610875576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180612a5d602a913960400191505060405180910390fd5b6000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206004015490506000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040181905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156109ad57600080fd5b505af11580156109c1573d6000803e3d6000fd5b505050506040513d60208110156109d757600080fd5b8101908080519060200190929190505050610a3d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e8152602001806129f3602e913960400191505060405180910390fd5b50565b600a5481565b600e5481565b60011515610a59336115e6565b151514610ab1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806129006030913960400191505060405180910390fd5b60011515610abe33611595565b151514610b16576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260278152602001806128b46027913960400191505060405180910390fd5b6000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206003015490506000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548160ff02191690836003811115610bbc57fe5b02179055506000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160016101000a81548160ff0219169083151502179055506000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030181905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015610d0d57600080fd5b505af1158015610d21573d6000803e3d6000fd5b505050506040513d6020811015610d3757600080fd5b8101908080519060200190929190505050610d9d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526031815260200180612aae6031913960400191505060405180910390fd5b50565b6000816003811115610dae57fe5b600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff166003811115610e0957fe5b14905092915050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600c5481565b600f5481565b600b5481565b60045481565b600d5481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60001515610e88336115e6565b151514610ee0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602c815260200180612994602c913960400191505060405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd33306002546040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b158015610fbf57600080fd5b505af1158015610fd3573d6000803e3d6000fd5b505050506040513d6020811015610fe957600080fd5b810190808051906020019092919050505061104f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602b815260200180612adf602b913960400191505060405180910390fd5b6006339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040518060c00160405280600060038111156110ca57fe5b81526020016001151581526020016001600680549050038152602001600354815260200160025481526020016000815250600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff0219169083600381111561115e57fe5b021790555060208201518160000160016101000a81548160ff02191690831515021790555060408201518160010155606082015181600201556080820151816003015560a08201518160040155905050565b60025481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611278576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f4f6e6c79206f776e65722063616e2063616c6c20746869732e0000000000000081525060200191505060405180910390fd5b606061128b600e54886010548989611e6e565b9050600060086000868152602001908152602001600020905060018160000160006101000a81548160ff021916908360028111156112c557fe5b021790555060018160000160016101000a81548160ff02191690831515021790555060008160000160026101000a81548160ff0219169083151502179055508381600101908051906020019061131c929190612741565b50603c42018160020181905550828160030181905550600f548160040181905550818160050190805190602001906113559291906127c1565b5060008090505b816005018054905081101561154457600082600501828154811061137c57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060018360060160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548160ff02191690831515021790555060008360060160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160016101000a81548160ff0219169083151502179055508073ffffffffffffffffffffffffffffffffffffffff167fb75dc87e2f9d1c0a563b98a1148d3f71d4137a45be4148a6a35a0b045a3c910188856002015486600301548a6040518085815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156114f95780820151818401526020810190506114de565b50505050905090810190601f1680156115265780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a250808060010191505061135c565b505050505050505050565b600081600281111561155d57fe5b6008600085815260200190815260200160002060000160009054906101000a900460ff16600281111561158c57fe5b14905092915050565b6000600454600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002015410159050919050565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160019054906101000a900460ff169050919050565b60105481565b6000816116518161215e565b6116c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f5769746e657373506f6f6c3a20696e76616c696420736c61000000000000000081525060200191505060405180910390fd5b826000600860008381526020019081526020016000209050600160028111156116e857fe5b8160000160009054906101000a900460ff16600281111561170557fe5b14801561171b5750806003015481600201540142115b156117595761172982612192565b61173282612642565b60028160000160006101000a81548160ff0219169083600281111561175357fe5b02179055505b8460026001151561176a838361154f565b1515146117c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260348152602001806129306034913960400191505060405180910390fd5b6008600088815260200190815260200160002060000160029054906101000a900460ff1695505050505050919050565b600115156117ff336115e6565b151514611857576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806129006030913960400191505060405180910390fd5b6001151561186433611595565b1515146118bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260278152602001806128b46027913960400191505060405180910390fd5b336000600115156118cd8383610da0565b151514611925576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260348152602001806129306034913960400191505060405180910390fd5b6001600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548160ff0219169083600381111561198457fe5b02179055506007600081548092919060010191905055505050565b600115156119ac336115e6565b151514611a04576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806129006030913960400191505060405180910390fd5b336001801515611a148383610da0565b151514611a6c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260348152602001806129306034913960400191505060405180910390fd5b6000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548160ff02191690836003811115611acb57fe5b0217905550600760008154809291906001900391905055505050565b6000600115156008600085815260200190815260200160002060060160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff16151514905092915050565b60095481565b60011515611b6f336115e6565b151514611bc7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806129006030913960400191505060405180910390fd5b60011515611bd433611595565b151514611c2c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260278152602001806128b46027913960400191505060405180910390fd5b80611c368161215e565b611ca8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f5769746e657373506f6f6c3a20696e76616c696420736c61000000000000000081525060200191505060405180910390fd5b81611cb38133611ae7565b611d08576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260338152602001806129c06033913960400191505060405180910390fd5b826001801515611d18838361154f565b151514611d70576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260348152602001806129306034913960400191505060405180910390fd5b8460086000828152602001908152602001600020600301546008600083815260200190815260200160002060020154014210611df7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806128db6025913960400191505060405180910390fd5b60016008600088815260200190815260200160002060060160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160016101000a81548160ff021916908315150217905550505050505050565b60608086600a0260075411611ece576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526027815260200180612a876027913960400191505060405180910390fd5b60ff86014310611f29576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603c815260200180612a21603c913960400191505060405180910390fd5b8486014311611f83576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806129646030913960400191505060405180910390fd5b600080905060008090505b86811015611fb1576001818901014060001c820191508080600101915050611f8e565b5060008090505b8881101561214f576000600680805490508481611fd157fe5b0681548110611fdc57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050612014816001610da0565b8015612025575061202481611595565b5b801561205d57508673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b801561209557508573ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b1561211a576003600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548160ff021916908360038111156120f957fe5b02179055506007600081548092919060019003919050555081806001019250505b8260001b604051602001808281526020019150506040516020818303038152906040528051906020012060001c925050611fb8565b82935050505095945050505050565b6000600115156008600084815260200190815260200160002060000160019054906101000a900460ff161515149050919050565b8060018015156121a2838361154f565b1515146121fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260348152602001806129306034913960400191505060405180910390fd5b6000600860008581526020019081526020016000209050600080905060008090505b82600501805490508110156122de57600083600501828154811061223c57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600115158460060160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160019054906101000a900460ff16151514156122d05782806001019350505b50808060010191505061221c565b50816004015481118260000160026101000a81548160ff02191690831515021790555060008090505b826005018054905081101561263a57600083600501828154811061232757fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600115158460000160029054906101000a900460ff16151514156124d157600115158460060160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160019054906101000a900460ff161515141561242757600954600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600401600082825401925050819055506124cc565b600a54600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160008282540192505081905550600d54600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201600082825403925050819055505b61262c565b600115158460060160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160019054906101000a900460ff16151514156125d857600b54600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160008282540192505081905550600d54600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002016000828254039250508190555061262b565b600c54600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600401600082825401925050819055505b5b508080600101915050612307565b505050505050565b60008090505b600860008381526020019081526020016000206005018054905081101561273d57600060086000848152602001908152602001600020600501828154811061268c57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506001600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548160ff0219169083600381111561271857fe5b0217905550600760008154809291906001019190505550508080600101915050612648565b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061278257805160ff19168380011785556127b0565b828001600101855582156127b0579182015b828111156127af578251825591602001919060010190612794565b5b5090506127bd919061284b565b5090565b82805482825590600052602060002090810192821561283a579160200282015b828111156128395782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550916020019190600101906127e1565b5b5090506128479190612870565b5090565b61286d91905b80821115612869576000816000905550600101612851565b5090565b90565b6128b091905b808211156128ac57600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101612876565b5090565b9056fe5769746e657373506f6f6c3a20746865207769746e65737320697320756e7175616c69666965645769746e657373506f6f6c3a206d6f6e69746f72696e672074696d652065786365656465645769746e657373506f6f6c3a20746865206164647265737320686164206e6f74206265656e2072656769737465726564576974656e7373506f6f6c3a2066756e6374696f6e2063616e6e6f742062652063616c6c6564206174207468697320737461746557696e746e657373506f6f6c3a206e6f206d6f726520626c6f636b4e65656420626c6f636b732067656e6572617465645769746e657373506f6f6c3a20746865206164647265737320686164206265656e20726567697374657265645769746e657373506f6f6c3a20746865207769746e657373206973206e6f74206120636f6d6d6974746565206d656d626572205769746e657373506f6f6c3a206661696c656420746f206472617720746865207769746e657373207265776172645769746e657373506f6f6c3a20626c6f636b686173682063616e206f6e6c792062652061636365737365642077697468696e203235352064657074685769746e657373506f6f6c3a2074686520726577617264206f66207769746e657373206973207a65726f57696e746e657373506f6f6c3a206e6f7420656e6f756768206f6e6c696e65207769746e6573735769746e657373506f6f6c3a206661696c656420746f20726566756e6420746865207769746e657373206465706f7369745769746e657373506f6f6c3a206661696c656420746f206c6f636b207769746e657373206465706f697374a2646970667358221220b3d209b31a671d47c1971c1ee01fab50d92ef1946345148907a499e2bb3153d964736f6c634300060a0033"

// DeployWitnessPool deploys a new contract, binding an instance of WitnessPool to it.
func DeployWitnessPool(auth *bind.TransactOpts, backend bind.ContractBackend, _tokenContractAddress common.Address) (common.Address, *types.Transaction, *WitnessPool, error) {
	parsed, err := abi.JSON(strings.NewReader(WitnessPoolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WitnessPoolBin), backend, _tokenContractAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WitnessPool{WitnessPoolCaller: WitnessPoolCaller{contract: contract}, WitnessPoolTransactor: WitnessPoolTransactor{contract: contract}, WitnessPoolFilterer: WitnessPoolFilterer{contract: contract}}, nil
}

func AsyncDeployWitnessPool(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend, _tokenContractAddress common.Address) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(WitnessPoolABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(WitnessPoolBin), backend, _tokenContractAddress)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// WitnessPool is an auto generated Go binding around a Solidity contract.
type WitnessPool struct {
	WitnessPoolCaller     // Read-only binding to the contract
	WitnessPoolTransactor // Write-only binding to the contract
	WitnessPoolFilterer   // Log filterer for contract events
}

// WitnessPoolCaller is an auto generated read-only Go binding around a Solidity contract.
type WitnessPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WitnessPoolTransactor is an auto generated write-only Go binding around a Solidity contract.
type WitnessPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WitnessPoolFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type WitnessPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WitnessPoolSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type WitnessPoolSession struct {
	Contract     *WitnessPool      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WitnessPoolCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type WitnessPoolCallerSession struct {
	Contract *WitnessPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// WitnessPoolTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type WitnessPoolTransactorSession struct {
	Contract     *WitnessPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// WitnessPoolRaw is an auto generated low-level Go binding around a Solidity contract.
type WitnessPoolRaw struct {
	Contract *WitnessPool // Generic contract binding to access the raw methods on
}

// WitnessPoolCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type WitnessPoolCallerRaw struct {
	Contract *WitnessPoolCaller // Generic read-only contract binding to access the raw methods on
}

// WitnessPoolTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type WitnessPoolTransactorRaw struct {
	Contract *WitnessPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWitnessPool creates a new instance of WitnessPool, bound to a specific deployed contract.
func NewWitnessPool(address common.Address, backend bind.ContractBackend) (*WitnessPool, error) {
	contract, err := bindWitnessPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WitnessPool{WitnessPoolCaller: WitnessPoolCaller{contract: contract}, WitnessPoolTransactor: WitnessPoolTransactor{contract: contract}, WitnessPoolFilterer: WitnessPoolFilterer{contract: contract}}, nil
}

// NewWitnessPoolCaller creates a new read-only instance of WitnessPool, bound to a specific deployed contract.
func NewWitnessPoolCaller(address common.Address, caller bind.ContractCaller) (*WitnessPoolCaller, error) {
	contract, err := bindWitnessPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WitnessPoolCaller{contract: contract}, nil
}

// NewWitnessPoolTransactor creates a new write-only instance of WitnessPool, bound to a specific deployed contract.
func NewWitnessPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*WitnessPoolTransactor, error) {
	contract, err := bindWitnessPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WitnessPoolTransactor{contract: contract}, nil
}

// NewWitnessPoolFilterer creates a new log filterer instance of WitnessPool, bound to a specific deployed contract.
func NewWitnessPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*WitnessPoolFilterer, error) {
	contract, err := bindWitnessPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WitnessPoolFilterer{contract: contract}, nil
}

// bindWitnessPool binds a generic wrapper to an already deployed contract.
func bindWitnessPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WitnessPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WitnessPool *WitnessPoolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WitnessPool.Contract.WitnessPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WitnessPool *WitnessPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.Contract.WitnessPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WitnessPool *WitnessPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.Contract.WitnessPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WitnessPool *WitnessPoolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WitnessPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WitnessPool *WitnessPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WitnessPool *WitnessPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.Contract.contract.Transact(opts, method, params...)
}

// IsAtSLAState is a free data retrieval call binding the contract method 0xcda28864.
//
// Solidity: function isAtSLAState(uint256 _slaID, uint8 _state) constant returns(bool)
func (_WitnessPool *WitnessPoolCaller) IsAtSLAState(opts *bind.CallOpts, _slaID *big.Int, _state uint8) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WitnessPool.contract.Call(opts, out, "isAtSLAState", _slaID, _state)
	return *ret0, err
}

// IsAtSLAState is a free data retrieval call binding the contract method 0xcda28864.
//
// Solidity: function isAtSLAState(uint256 _slaID, uint8 _state) constant returns(bool)
func (_WitnessPool *WitnessPoolSession) IsAtSLAState(_slaID *big.Int, _state uint8) (bool, error) {
	return _WitnessPool.Contract.IsAtSLAState(&_WitnessPool.CallOpts, _slaID, _state)
}

// IsAtSLAState is a free data retrieval call binding the contract method 0xcda28864.
//
// Solidity: function isAtSLAState(uint256 _slaID, uint8 _state) constant returns(bool)
func (_WitnessPool *WitnessPoolCallerSession) IsAtSLAState(_slaID *big.Int, _state uint8) (bool, error) {
	return _WitnessPool.Contract.IsAtSLAState(&_WitnessPool.CallOpts, _slaID, _state)
}

// IsAtWState is a free data retrieval call binding the contract method 0x47ff3aba.
//
// Solidity: function isAtWState(address _witness, uint8 _state) constant returns(bool)
func (_WitnessPool *WitnessPoolCaller) IsAtWState(opts *bind.CallOpts, _witness common.Address, _state uint8) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WitnessPool.contract.Call(opts, out, "isAtWState", _witness, _state)
	return *ret0, err
}

// IsAtWState is a free data retrieval call binding the contract method 0x47ff3aba.
//
// Solidity: function isAtWState(address _witness, uint8 _state) constant returns(bool)
func (_WitnessPool *WitnessPoolSession) IsAtWState(_witness common.Address, _state uint8) (bool, error) {
	return _WitnessPool.Contract.IsAtWState(&_WitnessPool.CallOpts, _witness, _state)
}

// IsAtWState is a free data retrieval call binding the contract method 0x47ff3aba.
//
// Solidity: function isAtWState(address _witness, uint8 _state) constant returns(bool)
func (_WitnessPool *WitnessPoolCallerSession) IsAtWState(_witness common.Address, _state uint8) (bool, error) {
	return _WitnessPool.Contract.IsAtWState(&_WitnessPool.CallOpts, _witness, _state)
}

// IsWitnessCommitteeMember is a free data retrieval call binding the contract method 0xed062567.
//
// Solidity: function isWitnessCommitteeMember(uint256 _slaID, address _witness) constant returns(bool)
func (_WitnessPool *WitnessPoolCaller) IsWitnessCommitteeMember(opts *bind.CallOpts, _slaID *big.Int, _witness common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WitnessPool.contract.Call(opts, out, "isWitnessCommitteeMember", _slaID, _witness)
	return *ret0, err
}

// IsWitnessCommitteeMember is a free data retrieval call binding the contract method 0xed062567.
//
// Solidity: function isWitnessCommitteeMember(uint256 _slaID, address _witness) constant returns(bool)
func (_WitnessPool *WitnessPoolSession) IsWitnessCommitteeMember(_slaID *big.Int, _witness common.Address) (bool, error) {
	return _WitnessPool.Contract.IsWitnessCommitteeMember(&_WitnessPool.CallOpts, _slaID, _witness)
}

// IsWitnessCommitteeMember is a free data retrieval call binding the contract method 0xed062567.
//
// Solidity: function isWitnessCommitteeMember(uint256 _slaID, address _witness) constant returns(bool)
func (_WitnessPool *WitnessPoolCallerSession) IsWitnessCommitteeMember(_slaID *big.Int, _witness common.Address) (bool, error) {
	return _WitnessPool.Contract.IsWitnessCommitteeMember(&_WitnessPool.CallOpts, _slaID, _witness)
}

// IsWitnessQualified is a free data retrieval call binding the contract method 0xcf617762.
//
// Solidity: function isWitnessQualified(address _witness) constant returns(bool)
func (_WitnessPool *WitnessPoolCaller) IsWitnessQualified(opts *bind.CallOpts, _witness common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WitnessPool.contract.Call(opts, out, "isWitnessQualified", _witness)
	return *ret0, err
}

// IsWitnessQualified is a free data retrieval call binding the contract method 0xcf617762.
//
// Solidity: function isWitnessQualified(address _witness) constant returns(bool)
func (_WitnessPool *WitnessPoolSession) IsWitnessQualified(_witness common.Address) (bool, error) {
	return _WitnessPool.Contract.IsWitnessQualified(&_WitnessPool.CallOpts, _witness)
}

// IsWitnessQualified is a free data retrieval call binding the contract method 0xcf617762.
//
// Solidity: function isWitnessQualified(address _witness) constant returns(bool)
func (_WitnessPool *WitnessPoolCallerSession) IsWitnessQualified(_witness common.Address) (bool, error) {
	return _WitnessPool.Contract.IsWitnessQualified(&_WitnessPool.CallOpts, _witness)
}

// IsWitnessRegistered is a free data retrieval call binding the contract method 0xd446fcdd.
//
// Solidity: function isWitnessRegistered(address _witness) constant returns(bool)
func (_WitnessPool *WitnessPoolCaller) IsWitnessRegistered(opts *bind.CallOpts, _witness common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WitnessPool.contract.Call(opts, out, "isWitnessRegistered", _witness)
	return *ret0, err
}

// IsWitnessRegistered is a free data retrieval call binding the contract method 0xd446fcdd.
//
// Solidity: function isWitnessRegistered(address _witness) constant returns(bool)
func (_WitnessPool *WitnessPoolSession) IsWitnessRegistered(_witness common.Address) (bool, error) {
	return _WitnessPool.Contract.IsWitnessRegistered(&_WitnessPool.CallOpts, _witness)
}

// IsWitnessRegistered is a free data retrieval call binding the contract method 0xd446fcdd.
//
// Solidity: function isWitnessRegistered(address _witness) constant returns(bool)
func (_WitnessPool *WitnessPoolCallerSession) IsWitnessRegistered(_witness common.Address) (bool, error) {
	return _WitnessPool.Contract.IsWitnessRegistered(&_WitnessPool.CallOpts, _witness)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_WitnessPool *WitnessPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WitnessPool.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_WitnessPool *WitnessPoolSession) Owner() (common.Address, error) {
	return _WitnessPool.Contract.Owner(&_WitnessPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_WitnessPool *WitnessPoolCallerSession) Owner() (common.Address, error) {
	return _WitnessPool.Contract.Owner(&_WitnessPool.CallOpts)
}

// ReputationDishonestReduced is a free data retrieval call binding the contract method 0x89be6f46.
//
// Solidity: function reputationDishonestReduced() constant returns(uint256)
func (_WitnessPool *WitnessPoolCaller) ReputationDishonestReduced(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WitnessPool.contract.Call(opts, out, "reputationDishonestReduced")
	return *ret0, err
}

// ReputationDishonestReduced is a free data retrieval call binding the contract method 0x89be6f46.
//
// Solidity: function reputationDishonestReduced() constant returns(uint256)
func (_WitnessPool *WitnessPoolSession) ReputationDishonestReduced() (*big.Int, error) {
	return _WitnessPool.Contract.ReputationDishonestReduced(&_WitnessPool.CallOpts)
}

// ReputationDishonestReduced is a free data retrieval call binding the contract method 0x89be6f46.
//
// Solidity: function reputationDishonestReduced() constant returns(uint256)
func (_WitnessPool *WitnessPoolCallerSession) ReputationDishonestReduced() (*big.Int, error) {
	return _WitnessPool.Contract.ReputationDishonestReduced(&_WitnessPool.CallOpts)
}

// RewardNoviolationReport is a free data retrieval call binding the contract method 0x7d635a7f.
//
// Solidity: function rewardNoviolationReport() constant returns(uint256)
func (_WitnessPool *WitnessPoolCaller) RewardNoviolationReport(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WitnessPool.contract.Call(opts, out, "rewardNoviolationReport")
	return *ret0, err
}

// RewardNoviolationReport is a free data retrieval call binding the contract method 0x7d635a7f.
//
// Solidity: function rewardNoviolationReport() constant returns(uint256)
func (_WitnessPool *WitnessPoolSession) RewardNoviolationReport() (*big.Int, error) {
	return _WitnessPool.Contract.RewardNoviolationReport(&_WitnessPool.CallOpts)
}

// RewardNoviolationReport is a free data retrieval call binding the contract method 0x7d635a7f.
//
// Solidity: function rewardNoviolationReport() constant returns(uint256)
func (_WitnessPool *WitnessPoolCallerSession) RewardNoviolationReport() (*big.Int, error) {
	return _WitnessPool.Contract.RewardNoviolationReport(&_WitnessPool.CallOpts)
}

// RewardNoviolationSilence is a free data retrieval call binding the contract method 0x744c2e33.
//
// Solidity: function rewardNoviolationSilence() constant returns(uint256)
func (_WitnessPool *WitnessPoolCaller) RewardNoviolationSilence(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WitnessPool.contract.Call(opts, out, "rewardNoviolationSilence")
	return *ret0, err
}

// RewardNoviolationSilence is a free data retrieval call binding the contract method 0x744c2e33.
//
// Solidity: function rewardNoviolationSilence() constant returns(uint256)
func (_WitnessPool *WitnessPoolSession) RewardNoviolationSilence() (*big.Int, error) {
	return _WitnessPool.Contract.RewardNoviolationSilence(&_WitnessPool.CallOpts)
}

// RewardNoviolationSilence is a free data retrieval call binding the contract method 0x744c2e33.
//
// Solidity: function rewardNoviolationSilence() constant returns(uint256)
func (_WitnessPool *WitnessPoolCallerSession) RewardNoviolationSilence() (*big.Int, error) {
	return _WitnessPool.Contract.RewardNoviolationSilence(&_WitnessPool.CallOpts)
}

// RewardViolationReport is a free data retrieval call binding the contract method 0xf11ad7f0.
//
// Solidity: function rewardViolationReport() constant returns(uint256)
func (_WitnessPool *WitnessPoolCaller) RewardViolationReport(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WitnessPool.contract.Call(opts, out, "rewardViolationReport")
	return *ret0, err
}

// RewardViolationReport is a free data retrieval call binding the contract method 0xf11ad7f0.
//
// Solidity: function rewardViolationReport() constant returns(uint256)
func (_WitnessPool *WitnessPoolSession) RewardViolationReport() (*big.Int, error) {
	return _WitnessPool.Contract.RewardViolationReport(&_WitnessPool.CallOpts)
}

// RewardViolationReport is a free data retrieval call binding the contract method 0xf11ad7f0.
//
// Solidity: function rewardViolationReport() constant returns(uint256)
func (_WitnessPool *WitnessPoolCallerSession) RewardViolationReport() (*big.Int, error) {
	return _WitnessPool.Contract.RewardViolationReport(&_WitnessPool.CallOpts)
}

// RewardViolationSilence is a free data retrieval call binding the contract method 0x37edd5a4.
//
// Solidity: function rewardViolationSilence() constant returns(uint256)
func (_WitnessPool *WitnessPoolCaller) RewardViolationSilence(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WitnessPool.contract.Call(opts, out, "rewardViolationSilence")
	return *ret0, err
}

// RewardViolationSilence is a free data retrieval call binding the contract method 0x37edd5a4.
//
// Solidity: function rewardViolationSilence() constant returns(uint256)
func (_WitnessPool *WitnessPoolSession) RewardViolationSilence() (*big.Int, error) {
	return _WitnessPool.Contract.RewardViolationSilence(&_WitnessPool.CallOpts)
}

// RewardViolationSilence is a free data retrieval call binding the contract method 0x37edd5a4.
//
// Solidity: function rewardViolationSilence() constant returns(uint256)
func (_WitnessPool *WitnessPoolCallerSession) RewardViolationSilence() (*big.Int, error) {
	return _WitnessPool.Contract.RewardViolationSilence(&_WitnessPool.CallOpts)
}

// StdNumReportRequired is a free data retrieval call binding the contract method 0x746ea9c6.
//
// Solidity: function stdNumReportRequired() constant returns(uint256)
func (_WitnessPool *WitnessPoolCaller) StdNumReportRequired(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WitnessPool.contract.Call(opts, out, "stdNumReportRequired")
	return *ret0, err
}

// StdNumReportRequired is a free data retrieval call binding the contract method 0x746ea9c6.
//
// Solidity: function stdNumReportRequired() constant returns(uint256)
func (_WitnessPool *WitnessPoolSession) StdNumReportRequired() (*big.Int, error) {
	return _WitnessPool.Contract.StdNumReportRequired(&_WitnessPool.CallOpts)
}

// StdNumReportRequired is a free data retrieval call binding the contract method 0x746ea9c6.
//
// Solidity: function stdNumReportRequired() constant returns(uint256)
func (_WitnessPool *WitnessPoolCallerSession) StdNumReportRequired() (*big.Int, error) {
	return _WitnessPool.Contract.StdNumReportRequired(&_WitnessPool.CallOpts)
}

// StdNumWitness is a free data retrieval call binding the contract method 0x3f5307c4.
//
// Solidity: function stdNumWitness() constant returns(uint256)
func (_WitnessPool *WitnessPoolCaller) StdNumWitness(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WitnessPool.contract.Call(opts, out, "stdNumWitness")
	return *ret0, err
}

// StdNumWitness is a free data retrieval call binding the contract method 0x3f5307c4.
//
// Solidity: function stdNumWitness() constant returns(uint256)
func (_WitnessPool *WitnessPoolSession) StdNumWitness() (*big.Int, error) {
	return _WitnessPool.Contract.StdNumWitness(&_WitnessPool.CallOpts)
}

// StdNumWitness is a free data retrieval call binding the contract method 0x3f5307c4.
//
// Solidity: function stdNumWitness() constant returns(uint256)
func (_WitnessPool *WitnessPoolCallerSession) StdNumWitness() (*big.Int, error) {
	return _WitnessPool.Contract.StdNumWitness(&_WitnessPool.CallOpts)
}

// StdWitnessDepoist is a free data retrieval call binding the contract method 0xa35da30c.
//
// Solidity: function stdWitnessDepoist() constant returns(uint256)
func (_WitnessPool *WitnessPoolCaller) StdWitnessDepoist(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WitnessPool.contract.Call(opts, out, "stdWitnessDepoist")
	return *ret0, err
}

// StdWitnessDepoist is a free data retrieval call binding the contract method 0xa35da30c.
//
// Solidity: function stdWitnessDepoist() constant returns(uint256)
func (_WitnessPool *WitnessPoolSession) StdWitnessDepoist() (*big.Int, error) {
	return _WitnessPool.Contract.StdWitnessDepoist(&_WitnessPool.CallOpts)
}

// StdWitnessDepoist is a free data retrieval call binding the contract method 0xa35da30c.
//
// Solidity: function stdWitnessDepoist() constant returns(uint256)
func (_WitnessPool *WitnessPoolCallerSession) StdWitnessDepoist() (*big.Int, error) {
	return _WitnessPool.Contract.StdWitnessDepoist(&_WitnessPool.CallOpts)
}

// StdblockNeed is a free data retrieval call binding the contract method 0xe0a5a2aa.
//
// Solidity: function stdblockNeed() constant returns(uint256)
func (_WitnessPool *WitnessPoolCaller) StdblockNeed(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WitnessPool.contract.Call(opts, out, "stdblockNeed")
	return *ret0, err
}

// StdblockNeed is a free data retrieval call binding the contract method 0xe0a5a2aa.
//
// Solidity: function stdblockNeed() constant returns(uint256)
func (_WitnessPool *WitnessPoolSession) StdblockNeed() (*big.Int, error) {
	return _WitnessPool.Contract.StdblockNeed(&_WitnessPool.CallOpts)
}

// StdblockNeed is a free data retrieval call binding the contract method 0xe0a5a2aa.
//
// Solidity: function stdblockNeed() constant returns(uint256)
func (_WitnessPool *WitnessPoolCallerSession) StdblockNeed() (*big.Int, error) {
	return _WitnessPool.Contract.StdblockNeed(&_WitnessPool.CallOpts)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_WitnessPool *WitnessPoolCaller) TokenContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WitnessPool.contract.Call(opts, out, "tokenContract")
	return *ret0, err
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_WitnessPool *WitnessPoolSession) TokenContract() (common.Address, error) {
	return _WitnessPool.Contract.TokenContract(&_WitnessPool.CallOpts)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_WitnessPool *WitnessPoolCallerSession) TokenContract() (common.Address, error) {
	return _WitnessPool.Contract.TokenContract(&_WitnessPool.CallOpts)
}

// WitnessReputationInit is a free data retrieval call binding the contract method 0x16c3734d.
//
// Solidity: function witnessReputationInit() constant returns(uint256)
func (_WitnessPool *WitnessPoolCaller) WitnessReputationInit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WitnessPool.contract.Call(opts, out, "witnessReputationInit")
	return *ret0, err
}

// WitnessReputationInit is a free data retrieval call binding the contract method 0x16c3734d.
//
// Solidity: function witnessReputationInit() constant returns(uint256)
func (_WitnessPool *WitnessPoolSession) WitnessReputationInit() (*big.Int, error) {
	return _WitnessPool.Contract.WitnessReputationInit(&_WitnessPool.CallOpts)
}

// WitnessReputationInit is a free data retrieval call binding the contract method 0x16c3734d.
//
// Solidity: function witnessReputationInit() constant returns(uint256)
func (_WitnessPool *WitnessPoolCallerSession) WitnessReputationInit() (*big.Int, error) {
	return _WitnessPool.Contract.WitnessReputationInit(&_WitnessPool.CallOpts)
}

// WitnessReputationQualified is a free data retrieval call binding the contract method 0x89640b58.
//
// Solidity: function witnessReputationQualified() constant returns(uint256)
func (_WitnessPool *WitnessPoolCaller) WitnessReputationQualified(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WitnessPool.contract.Call(opts, out, "witnessReputationQualified")
	return *ret0, err
}

// WitnessReputationQualified is a free data retrieval call binding the contract method 0x89640b58.
//
// Solidity: function witnessReputationQualified() constant returns(uint256)
func (_WitnessPool *WitnessPoolSession) WitnessReputationQualified() (*big.Int, error) {
	return _WitnessPool.Contract.WitnessReputationQualified(&_WitnessPool.CallOpts)
}

// WitnessReputationQualified is a free data retrieval call binding the contract method 0x89640b58.
//
// Solidity: function witnessReputationQualified() constant returns(uint256)
func (_WitnessPool *WitnessPoolCallerSession) WitnessReputationQualified() (*big.Int, error) {
	return _WitnessPool.Contract.WitnessReputationQualified(&_WitnessPool.CallOpts)
}

// IsViolatedSLA is a paid mutator transaction binding the contract method 0xe81bc992.
//
// Solidity: function isViolatedSLA(uint256 _slaID) returns(bool)
func (_WitnessPool *WitnessPoolTransactor) IsViolatedSLA(opts *bind.TransactOpts, _slaID *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.contract.Transact(opts, "isViolatedSLA", _slaID)
}

func (_WitnessPool *WitnessPoolTransactor) AsyncIsViolatedSLA(handler func(*types.Receipt, error), opts *bind.TransactOpts, _slaID *big.Int) (*types.Transaction, error) {
	return _WitnessPool.contract.AsyncTransact(opts, handler, "isViolatedSLA", _slaID)
}

// IsViolatedSLA is a paid mutator transaction binding the contract method 0xe81bc992.
//
// Solidity: function isViolatedSLA(uint256 _slaID) returns(bool)
func (_WitnessPool *WitnessPoolSession) IsViolatedSLA(_slaID *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.Contract.IsViolatedSLA(&_WitnessPool.TransactOpts, _slaID)
}

func (_WitnessPool *WitnessPoolSession) AsyncIsViolatedSLA(handler func(*types.Receipt, error), _slaID *big.Int) (*types.Transaction, error) {
	return _WitnessPool.Contract.AsyncIsViolatedSLA(handler, &_WitnessPool.TransactOpts, _slaID)
}

// IsViolatedSLA is a paid mutator transaction binding the contract method 0xe81bc992.
//
// Solidity: function isViolatedSLA(uint256 _slaID) returns(bool)
func (_WitnessPool *WitnessPoolTransactorSession) IsViolatedSLA(_slaID *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.Contract.IsViolatedSLA(&_WitnessPool.TransactOpts, _slaID)
}

func (_WitnessPool *WitnessPoolTransactorSession) AsyncIsViolatedSLA(handler func(*types.Receipt, error), _slaID *big.Int) (*types.Transaction, error) {
	return _WitnessPool.Contract.AsyncIsViolatedSLA(handler, &_WitnessPool.TransactOpts, _slaID)
}

// NewSLA is a paid mutator transaction binding the contract method 0xb65f6706.
//
// Solidity: function newSLA(uint256 _curBlockNum, address _provider, address _customer, uint256 _slaID, string _funcToMonitor, uint256 _monitoringDuration) returns()
func (_WitnessPool *WitnessPoolTransactor) NewSLA(opts *bind.TransactOpts, _curBlockNum *big.Int, _provider common.Address, _customer common.Address, _slaID *big.Int, _funcToMonitor string, _monitoringDuration *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.contract.Transact(opts, "newSLA", _curBlockNum, _provider, _customer, _slaID, _funcToMonitor, _monitoringDuration)
}

func (_WitnessPool *WitnessPoolTransactor) AsyncNewSLA(handler func(*types.Receipt, error), opts *bind.TransactOpts, _curBlockNum *big.Int, _provider common.Address, _customer common.Address, _slaID *big.Int, _funcToMonitor string, _monitoringDuration *big.Int) (*types.Transaction, error) {
	return _WitnessPool.contract.AsyncTransact(opts, handler, "newSLA", _curBlockNum, _provider, _customer, _slaID, _funcToMonitor, _monitoringDuration)
}

// NewSLA is a paid mutator transaction binding the contract method 0xb65f6706.
//
// Solidity: function newSLA(uint256 _curBlockNum, address _provider, address _customer, uint256 _slaID, string _funcToMonitor, uint256 _monitoringDuration) returns()
func (_WitnessPool *WitnessPoolSession) NewSLA(_curBlockNum *big.Int, _provider common.Address, _customer common.Address, _slaID *big.Int, _funcToMonitor string, _monitoringDuration *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.Contract.NewSLA(&_WitnessPool.TransactOpts, _curBlockNum, _provider, _customer, _slaID, _funcToMonitor, _monitoringDuration)
}

func (_WitnessPool *WitnessPoolSession) AsyncNewSLA(handler func(*types.Receipt, error), _curBlockNum *big.Int, _provider common.Address, _customer common.Address, _slaID *big.Int, _funcToMonitor string, _monitoringDuration *big.Int) (*types.Transaction, error) {
	return _WitnessPool.Contract.AsyncNewSLA(handler, &_WitnessPool.TransactOpts, _curBlockNum, _provider, _customer, _slaID, _funcToMonitor, _monitoringDuration)
}

// NewSLA is a paid mutator transaction binding the contract method 0xb65f6706.
//
// Solidity: function newSLA(uint256 _curBlockNum, address _provider, address _customer, uint256 _slaID, string _funcToMonitor, uint256 _monitoringDuration) returns()
func (_WitnessPool *WitnessPoolTransactorSession) NewSLA(_curBlockNum *big.Int, _provider common.Address, _customer common.Address, _slaID *big.Int, _funcToMonitor string, _monitoringDuration *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.Contract.NewSLA(&_WitnessPool.TransactOpts, _curBlockNum, _provider, _customer, _slaID, _funcToMonitor, _monitoringDuration)
}

func (_WitnessPool *WitnessPoolTransactorSession) AsyncNewSLA(handler func(*types.Receipt, error), _curBlockNum *big.Int, _provider common.Address, _customer common.Address, _slaID *big.Int, _funcToMonitor string, _monitoringDuration *big.Int) (*types.Transaction, error) {
	return _WitnessPool.Contract.AsyncNewSLA(handler, &_WitnessPool.TransactOpts, _curBlockNum, _provider, _customer, _slaID, _funcToMonitor, _monitoringDuration)
}

// ReportViolation is a paid mutator transaction binding the contract method 0xfc634e8f.
//
// Solidity: function reportViolation(uint256 _slaID) returns()
func (_WitnessPool *WitnessPoolTransactor) ReportViolation(opts *bind.TransactOpts, _slaID *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.contract.Transact(opts, "reportViolation", _slaID)
}

func (_WitnessPool *WitnessPoolTransactor) AsyncReportViolation(handler func(*types.Receipt, error), opts *bind.TransactOpts, _slaID *big.Int) (*types.Transaction, error) {
	return _WitnessPool.contract.AsyncTransact(opts, handler, "reportViolation", _slaID)
}

// ReportViolation is a paid mutator transaction binding the contract method 0xfc634e8f.
//
// Solidity: function reportViolation(uint256 _slaID) returns()
func (_WitnessPool *WitnessPoolSession) ReportViolation(_slaID *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.Contract.ReportViolation(&_WitnessPool.TransactOpts, _slaID)
}

func (_WitnessPool *WitnessPoolSession) AsyncReportViolation(handler func(*types.Receipt, error), _slaID *big.Int) (*types.Transaction, error) {
	return _WitnessPool.Contract.AsyncReportViolation(handler, &_WitnessPool.TransactOpts, _slaID)
}

// ReportViolation is a paid mutator transaction binding the contract method 0xfc634e8f.
//
// Solidity: function reportViolation(uint256 _slaID) returns()
func (_WitnessPool *WitnessPoolTransactorSession) ReportViolation(_slaID *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.Contract.ReportViolation(&_WitnessPool.TransactOpts, _slaID)
}

func (_WitnessPool *WitnessPoolTransactorSession) AsyncReportViolation(handler func(*types.Receipt, error), _slaID *big.Int) (*types.Transaction, error) {
	return _WitnessPool.Contract.AsyncReportViolation(handler, &_WitnessPool.TransactOpts, _slaID)
}

// WintessTurnOn is a paid mutator transaction binding the contract method 0xebe47110.
//
// Solidity: function wintessTurnOn() returns()
func (_WitnessPool *WitnessPoolTransactor) WintessTurnOn(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.contract.Transact(opts, "wintessTurnOn")
}

func (_WitnessPool *WitnessPoolTransactor) AsyncWintessTurnOn(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessPool.contract.AsyncTransact(opts, handler, "wintessTurnOn")
}

// WintessTurnOn is a paid mutator transaction binding the contract method 0xebe47110.
//
// Solidity: function wintessTurnOn() returns()
func (_WitnessPool *WitnessPoolSession) WintessTurnOn() (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.Contract.WintessTurnOn(&_WitnessPool.TransactOpts)
}

func (_WitnessPool *WitnessPoolSession) AsyncWintessTurnOn(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _WitnessPool.Contract.AsyncWintessTurnOn(handler, &_WitnessPool.TransactOpts)
}

// WintessTurnOn is a paid mutator transaction binding the contract method 0xebe47110.
//
// Solidity: function wintessTurnOn() returns()
func (_WitnessPool *WitnessPoolTransactorSession) WintessTurnOn() (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.Contract.WintessTurnOn(&_WitnessPool.TransactOpts)
}

func (_WitnessPool *WitnessPoolTransactorSession) AsyncWintessTurnOn(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _WitnessPool.Contract.AsyncWintessTurnOn(handler, &_WitnessPool.TransactOpts)
}

// WitenessLogout is a paid mutator transaction binding the contract method 0x426f93b1.
//
// Solidity: function witenessLogout() returns()
func (_WitnessPool *WitnessPoolTransactor) WitenessLogout(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.contract.Transact(opts, "witenessLogout")
}

func (_WitnessPool *WitnessPoolTransactor) AsyncWitenessLogout(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessPool.contract.AsyncTransact(opts, handler, "witenessLogout")
}

// WitenessLogout is a paid mutator transaction binding the contract method 0x426f93b1.
//
// Solidity: function witenessLogout() returns()
func (_WitnessPool *WitnessPoolSession) WitenessLogout() (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.Contract.WitenessLogout(&_WitnessPool.TransactOpts)
}

func (_WitnessPool *WitnessPoolSession) AsyncWitenessLogout(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _WitnessPool.Contract.AsyncWitenessLogout(handler, &_WitnessPool.TransactOpts)
}

// WitenessLogout is a paid mutator transaction binding the contract method 0x426f93b1.
//
// Solidity: function witenessLogout() returns()
func (_WitnessPool *WitnessPoolTransactorSession) WitenessLogout() (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.Contract.WitenessLogout(&_WitnessPool.TransactOpts)
}

func (_WitnessPool *WitnessPoolTransactorSession) AsyncWitenessLogout(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _WitnessPool.Contract.AsyncWitenessLogout(handler, &_WitnessPool.TransactOpts)
}

// WitnessDrawReward is a paid mutator transaction binding the contract method 0x2fba9d1a.
//
// Solidity: function witnessDrawReward() returns()
func (_WitnessPool *WitnessPoolTransactor) WitnessDrawReward(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.contract.Transact(opts, "witnessDrawReward")
}

func (_WitnessPool *WitnessPoolTransactor) AsyncWitnessDrawReward(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessPool.contract.AsyncTransact(opts, handler, "witnessDrawReward")
}

// WitnessDrawReward is a paid mutator transaction binding the contract method 0x2fba9d1a.
//
// Solidity: function witnessDrawReward() returns()
func (_WitnessPool *WitnessPoolSession) WitnessDrawReward() (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.Contract.WitnessDrawReward(&_WitnessPool.TransactOpts)
}

func (_WitnessPool *WitnessPoolSession) AsyncWitnessDrawReward(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _WitnessPool.Contract.AsyncWitnessDrawReward(handler, &_WitnessPool.TransactOpts)
}

// WitnessDrawReward is a paid mutator transaction binding the contract method 0x2fba9d1a.
//
// Solidity: function witnessDrawReward() returns()
func (_WitnessPool *WitnessPoolTransactorSession) WitnessDrawReward() (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.Contract.WitnessDrawReward(&_WitnessPool.TransactOpts)
}

func (_WitnessPool *WitnessPoolTransactorSession) AsyncWitnessDrawReward(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _WitnessPool.Contract.AsyncWitnessDrawReward(handler, &_WitnessPool.TransactOpts)
}

// WitnessLogin is a paid mutator transaction binding the contract method 0x9faa79e3.
//
// Solidity: function witnessLogin() returns()
func (_WitnessPool *WitnessPoolTransactor) WitnessLogin(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.contract.Transact(opts, "witnessLogin")
}

func (_WitnessPool *WitnessPoolTransactor) AsyncWitnessLogin(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessPool.contract.AsyncTransact(opts, handler, "witnessLogin")
}

// WitnessLogin is a paid mutator transaction binding the contract method 0x9faa79e3.
//
// Solidity: function witnessLogin() returns()
func (_WitnessPool *WitnessPoolSession) WitnessLogin() (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.Contract.WitnessLogin(&_WitnessPool.TransactOpts)
}

func (_WitnessPool *WitnessPoolSession) AsyncWitnessLogin(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _WitnessPool.Contract.AsyncWitnessLogin(handler, &_WitnessPool.TransactOpts)
}

// WitnessLogin is a paid mutator transaction binding the contract method 0x9faa79e3.
//
// Solidity: function witnessLogin() returns()
func (_WitnessPool *WitnessPoolTransactorSession) WitnessLogin() (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.Contract.WitnessLogin(&_WitnessPool.TransactOpts)
}

func (_WitnessPool *WitnessPoolTransactorSession) AsyncWitnessLogin(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _WitnessPool.Contract.AsyncWitnessLogin(handler, &_WitnessPool.TransactOpts)
}

// WitnessTurnOff is a paid mutator transaction binding the contract method 0xecc7e64b.
//
// Solidity: function witnessTurnOff() returns()
func (_WitnessPool *WitnessPoolTransactor) WitnessTurnOff(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.contract.Transact(opts, "witnessTurnOff")
}

func (_WitnessPool *WitnessPoolTransactor) AsyncWitnessTurnOff(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessPool.contract.AsyncTransact(opts, handler, "witnessTurnOff")
}

// WitnessTurnOff is a paid mutator transaction binding the contract method 0xecc7e64b.
//
// Solidity: function witnessTurnOff() returns()
func (_WitnessPool *WitnessPoolSession) WitnessTurnOff() (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.Contract.WitnessTurnOff(&_WitnessPool.TransactOpts)
}

func (_WitnessPool *WitnessPoolSession) AsyncWitnessTurnOff(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _WitnessPool.Contract.AsyncWitnessTurnOff(handler, &_WitnessPool.TransactOpts)
}

// WitnessTurnOff is a paid mutator transaction binding the contract method 0xecc7e64b.
//
// Solidity: function witnessTurnOff() returns()
func (_WitnessPool *WitnessPoolTransactorSession) WitnessTurnOff() (*types.Transaction, *types.Receipt, error) {
	return _WitnessPool.Contract.WitnessTurnOff(&_WitnessPool.TransactOpts)
}

func (_WitnessPool *WitnessPoolTransactorSession) AsyncWitnessTurnOff(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _WitnessPool.Contract.AsyncWitnessTurnOff(handler, &_WitnessPool.TransactOpts)
}

// WitnessPoolWitnessSelectedEventIterator is returned from FilterWitnessSelectedEvent and is used to iterate over the raw logs and unpacked data for WitnessSelectedEvent events raised by the WitnessPool contract.
type WitnessPoolWitnessSelectedEventIterator struct {
	Event *WitnessPoolWitnessSelectedEvent // Event containing the contract specifics and raw log

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
func (it *WitnessPoolWitnessSelectedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessPoolWitnessSelectedEvent)
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
		it.Event = new(WitnessPoolWitnessSelectedEvent)
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
func (it *WitnessPoolWitnessSelectedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessPoolWitnessSelectedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessPoolWitnessSelectedEvent represents a WitnessSelectedEvent event raised by the WitnessPool contract.
type WitnessPoolWitnessSelectedEvent struct {
	Witness             common.Address
	SlaID               *big.Int
	MonitoringBeginTime *big.Int
	MonitoringDuration  *big.Int
	FuncToMonitor       string
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterWitnessSelectedEvent is a free log retrieval operation binding the contract event 0xb75dc87e2f9d1c0a563b98a1148d3f71d4137a45be4148a6a35a0b045a3c9101.
//
// Solidity: event WitnessSelectedEvent(address indexed _witness, uint256 _slaID, uint256 _monitoringBeginTime, uint256 _monitoringDuration, string _funcToMonitor)
func (_WitnessPool *WitnessPoolFilterer) FilterWitnessSelectedEvent(opts *bind.FilterOpts, _witness []common.Address) (*WitnessPoolWitnessSelectedEventIterator, error) {

	var _witnessRule []interface{}
	for _, _witnessItem := range _witness {
		_witnessRule = append(_witnessRule, _witnessItem)
	}

	logs, sub, err := _WitnessPool.contract.FilterLogs(opts, "WitnessSelectedEvent", _witnessRule)
	if err != nil {
		return nil, err
	}
	return &WitnessPoolWitnessSelectedEventIterator{contract: _WitnessPool.contract, event: "WitnessSelectedEvent", logs: logs, sub: sub}, nil
}

// WatchWitnessSelectedEvent is a free log subscription operation binding the contract event 0xb75dc87e2f9d1c0a563b98a1148d3f71d4137a45be4148a6a35a0b045a3c9101.
//
// Solidity: event WitnessSelectedEvent(address indexed _witness, uint256 _slaID, uint256 _monitoringBeginTime, uint256 _monitoringDuration, string _funcToMonitor)
func (_WitnessPool *WitnessPoolFilterer) WatchWitnessSelectedEvent(opts *bind.WatchOpts, sink chan<- *WitnessPoolWitnessSelectedEvent, _witness []common.Address) (event.Subscription, error) {

	var _witnessRule []interface{}
	for _, _witnessItem := range _witness {
		_witnessRule = append(_witnessRule, _witnessItem)
	}

	logs, sub, err := _WitnessPool.contract.WatchLogs(opts, "WitnessSelectedEvent", _witnessRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessPoolWitnessSelectedEvent)
				if err := _WitnessPool.contract.UnpackLog(event, "WitnessSelectedEvent", log); err != nil {
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

// ParseWitnessSelectedEvent is a log parse operation binding the contract event 0xb75dc87e2f9d1c0a563b98a1148d3f71d4137a45be4148a6a35a0b045a3c9101.
//
// Solidity: event WitnessSelectedEvent(address indexed _witness, uint256 _slaID, uint256 _monitoringBeginTime, uint256 _monitoringDuration, string _funcToMonitor)
func (_WitnessPool *WitnessPoolFilterer) ParseWitnessSelectedEvent(log types.Log) (*WitnessPoolWitnessSelectedEvent, error) {
	event := new(WitnessPoolWitnessSelectedEvent)
	if err := _WitnessPool.contract.UnpackLog(event, "WitnessSelectedEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}
