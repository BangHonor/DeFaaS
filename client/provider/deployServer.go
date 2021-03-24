package provider

import (
	"defaas/core/data"
	"math/big"
	"strings"

	"defaas/adapter"

	"github.com/gogf/gf/net/ghttp"
)

func GetDeployPathFromProviderConfig(pc *ProviderConfig) string {

	return GetDeployPath(pc.Adapter, pc.ServerAddr, pc.ServerEntry)

}

func GetDeployPath(adapter, serverAddr, serverEntry string) string {

	deployPath := strings.Join([]string{adapter, serverAddr, serverEntry}, "|")

	return deployPath
}

func ParseDeployPath(deployPath string) (string, string, string) {
	ss := strings.Split(deployPath, "|")
	return ss[0], ss[1], ss[2]
}

func (client *ProviderClient) DeployServerRequestHandler(r *ghttp.Request) {

	// 1. parse parameter from request
	// 2. check is at pool
	// 3. auth
	// 4. use adpater
	// 4.1 docker adapter

	id, success := big.NewInt(0).SetString(r.GetString("id"), 10)
	if !success {
		r.Response.Write("Hello")
	}
	accessKey := r.GetString("accessKey")
	adapterData := r.GetString("adapterData")

	item := client.itemPool.Get(id).(*data.DeploymentItem)
	if data.DeployingState != item.State {
		r.Response.Write("Hello")
	}
	if accessKey != item.Info.AccessKey {
		r.Response.Write("Hello")
	}

	adapter.New(client.providerConfig.Adapter).DeployFrom(item, adapterData)

	r.Response.Write("Hello")
}

func (client *ProviderClient) DeployServerRegisterDeploying(item *data.DeploymentItem) {

	notifier := make(chan [32]byte, 1)

	client.deployedNotifierPool.Set(item.ID, notifier)
}

func (client *ProviderClient) DeployServerWaitForDeploying(item *data.DeploymentItem) (fulfillSecretKey [32]byte) {

	// wait
	fulfillSecretKey = <-client.deployedNotifierPool.Get(item.ID).(chan [32]byte)

	// delete notifier
	client.deployedNotifierPool.Remove(item.ID)

	return fulfillSecretKey
}
