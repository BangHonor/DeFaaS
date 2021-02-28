package customer

import (
	"defaas/contracts/go/faastoken"
	"defaas/contracts/go/market"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"

	sdkClient "github.com/FISCO-BCOS/go-sdk/client"
)

const (
	// ToleranceDuration 与区块链时间的相差的容忍时间间隔
	ToleranceDuration = 10 * time.Minute
)

// DeploymentOrder 部署订单
type DeploymentOrder struct {
	Customer         common.Address
	FaaSLevelID      *big.Int
	HighestUnitPrice *big.Int
	FaaSDuration     time.Duration
	BiddingDuration  time.Duration
}

// Lease 租约
type Lease struct {
	Customer common.Address
	Provider common.Address

	FaaSLevelID  *big.Int
	FaaSDuration *big.Int
	UnitPrice    *big.Int // 成交价格

	ProviderServiceFee  *big.Int
	CustomerWithdrawFee *big.Int
	WitnessFee          *big.Int
	MaintainerFee       *big.Int
}

// Function 是 FaaS 函数的表示
type Function struct{}

// Client 是租户客户端
type Client struct {
	_client   *sdkClient.Client
	faastoken *faastoken.FaaSTokenSession
	market    *market.MarketSession

	FaaSTokenContractAddress common.Address
	MarketContractAddress    common.Address
}

// NewClient 新建租户客户端
func NewClient() *Client {
	client := &Client{}
	// TODO
	return client
}

func timeDurationToSoliditySecond(d time.Duration) (seconds *big.Int) {
	return big.NewInt(int64(d.Seconds()))
}

// Deploy 部署一个 FaaS 任务
func (client *Client) Deploy(order *DeploymentOrder, function *Function) error {

	deploymentOrderID, _ := client.newDeploymentOrder(order)
	success, _, _ := client.matchDeploymentOrder(deploymentOrderID)

	// 等待竞价结束
	<-time.After(order.BiddingDuration + ToleranceDuration)

	// 订单流拍，返回
	if !success {
		return errors.New("deploymentOrder is missed")
	}

	// 订单匹配成功
	// 开始部署
	// client.deployToProvider()
	// client.market.FulfillDeploymentOrder()

	// 等待履行结束
	<-time.After(order.FaaSDuration + ToleranceDuration)

	// 结算订单
	// client.market.SettleDeploymentOrder()

	// 结束订单
	// client.market.FinishDeploymemtOrder()

	return nil
}

func (client *Client) newDeploymentOrder(order *DeploymentOrder) (deploymentOrderID *big.Int, err error) {

	lockFee, err := client.market.GetDeploymentOrderLockFee(
		order.HighestUnitPrice,
		timeDurationToSoliditySecond(order.FaaSDuration))

	_, _, err = client.faastoken.Approve(client.MarketContractAddress, lockFee)

	_, _, err = client.market.NewDeploymentOrder(
		order.FaaSLevelID,
		order.HighestUnitPrice,
		timeDurationToSoliditySecond(order.FaaSDuration),
		timeDurationToSoliditySecond(order.BiddingDuration))

	return deploymentOrderID, err
}

func (client *Client) matchDeploymentOrder(deploymentOrderID *big.Int) (success bool, leaseID *big.Int, err error) {

	_, _, err = client.market.MatchDeploymentOrder(deploymentOrderID)

	return success, leaseID, err
}
