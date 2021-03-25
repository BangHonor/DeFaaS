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

	{
		// deploy function service ...
		adapter.New(client.providerConfig.Adapter).DeployFrom(item, adapterData)
	}

	r.Response.Write("Hello")

	item.Order.FulfillSecretKey = [32]byte{} // TODO
	client.DeployServerNotify(item)
}

func (client *ProviderClient) DeployServerRegisterWait(item *data.DeploymentItem) {

	notifier := make(chan struct{}, 1)

	// register notifier
	client.deployedNotifierPool.Set(item.ID, notifier)

	// wait
	<-client.deployedNotifierPool.Get(item.ID).(chan struct{})
	// TODO select with a timeout and ctx

	// unregister notifier
	client.deployedNotifierPool.Remove(item.ID)
}

func (client *ProviderClient) DeployServerNotify(item *data.DeploymentItem) {

	notifier := client.deployedNotifierPool.Get(item.ID)

	if notifier != nil {
		notifier.(chan struct{}) <- struct{}{}
		return
	}
	// _ = errors.New("invalid notifier...")
}
