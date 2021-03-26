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
		r.Response.Write("Err")
	}
	if !client.DeployServerIsRegistered(id) {
		r.Response.Write("Err")
	}

	item := client.itemPool.Get(id).(*data.DeploymentItem)
	if data.DeployingState != item.State {
		r.Response.Write("Err")
	}

	accessKey := r.GetString("accessKey")
	if accessKey != item.Info.AccessKey {
		r.Response.Write("Err")
	}

	{
		// deploy function service ...
		adapterData := r.GetString("adapterData")
		adapter.New(client.providerConfig.Adapter).DeployFrom(item, adapterData)

		// notify
		client.DeployServerNotify(id)
	}

	r.Response.Write("Ok")
}

func (client *ProviderClient) DeployServerIsRegistered(id *big.Int) bool {

	notifier := client.deployedNotifierPool.Get(id)

	return (notifier != nil)
}

func (client *ProviderClient) DeployServerRegisterWait(id *big.Int) [32]byte {

	notifier := make(chan struct{}, 1)

	// register notifier
	client.deployedNotifierPool.Set(id, notifier)

	// wait
	<-client.deployedNotifierPool.Get(id).(chan struct{})
	// TODO select with a timeout and ctx

	// unregister notifier
	client.deployedNotifierPool.Remove(id)

	// TODO
	fulfillSecretKey := [32]byte{}
	return fulfillSecretKey
}

func (client *ProviderClient) DeployServerNotify(id *big.Int) {

	if client.DeployServerIsRegistered(id) {
		notifier := client.deployedNotifierPool.Get(id)
		notifier.(chan struct{}) <- struct{}{}
		return
	}
}
