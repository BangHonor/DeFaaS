package data

import (
	"encoding/json"
	"errors"
	"log"
	"math/big"
	"regexp"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

type FaaSLevel struct {
	ID   int
	Core *big.Int
	Mem  *big.Int
}

type DeploymentOrderState int

const (
	InitState       DeploymentOrderState = 0
	BiddingState    DeploymentOrderState = 1
	ConfirmingState DeploymentOrderState = 2
	DeployingState  DeploymentOrderState = 3
	FulfillingState DeploymentOrderState = 4
	FinishedState   DeploymentOrderState = 5
)

var (
	ErrWrongState = errors.New("wrong state")
)

type DeploymentOrder struct {
	Customer         common.Address
	Nonce            *big.Int
	FaaSLevelID      *big.Int
	HighestUnitPrice *big.Int
	FaaSDuration     *big.Int
	BiddingDuration  *big.Int
	PublicKey        string
	FulfillSecretKey [32]byte
	FulfillKey       [32]byte // FulfillKey = sha256.Sum256(FulfillSecretKey)
}

type DeploymentInfo struct {
	Provider   common.Address
	UnitPrice  *big.Int
	FuncPath   string
	DeployPath string
	AccessKey  string
}

type Lease struct {
	IsProviderViolated bool
}

type DeploymentItem struct {
	sync.Mutex

	ID    *big.Int
	State DeploymentOrderState
	Order DeploymentOrder
	Info  DeploymentInfo
}

// --------------------------------------------------------------------

type DeployToProviderRequest struct {
	CustomerAddress   common.Address `json:"customer_address_hex"`
	DeploymentOrderID *big.Int       `json:"deployment_order_id"`
	AccessKey         string         `json:"access_key"`
	FulfillSecretKey  [32]byte       `json:"fulfill_secret_key"`
	Adapter           string         `json:"adapter"`
	AdapterData       []byte         `json:"adapter_data"`
}

type DeployToProviderResponce struct {
	Code  int         `json:"code"`  // 错误码((0:成功, 1:失败, >1:错误码))
	Error string      `json:"error"` // 错误信息
	Data  interface{} `json:"data"`  // 返回数据
}

func (r DeployToProviderRequest) String() string {

	j, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	return string(j)
}

func (r DeployToProviderResponce) String() string {

	j, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	return string(j)
}

// ------------------------------------------------------------

func GetDeployPath(adapter, serverAddr, serverEntry string) string {

	deployPath := strings.Join([]string{adapter, serverAddr, serverEntry}, "|")

	return deployPath
}

func ParseDeployPath(deployPath string) (string, string, string) {
	ss := strings.Split(deployPath, "|")
	return ss[0], ss[1], ss[2]
}

func ServerAddrTrim(serverAddr string) string {

	rep := regexp.MustCompile(`http(|s)://`)

	return rep.ReplaceAllString(serverAddr, "")
}
