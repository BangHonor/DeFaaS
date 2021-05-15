package provider

import (
	"defaas/core/data"
	"encoding/json"
	"math/big"

	"github.com/gogf/gf/net/ghttp"
)

func (client *ProviderClient) DeployServerRequestHandler(r *ghttp.Request) {

	// 1
	// parse parameter from request
	// 2
	// check is at pool
	// 3
	// auth
	// 4
	// deploy on provider

	// parse parameter from request
	req := &data.DeployToProviderRequest{}
	reqJson := r.GetBody()
	if err := json.Unmarshal(reqJson, req); err != nil {

		r.Response.WriteJsonExit(data.DeployToProviderResponce{
			Code:  1,
			Error: err.Error(),
			Data:  nil,
		})

	}

	// check is at pool
	if !client.DeployServerIsRegistered(req.DeploymentOrderID) {
		r.Response.WriteJsonExit(data.DeployToProviderResponce{
			Code:  2,
			Error: "invalid deployment order id",
			Data:  nil,
		})
	}

	item := client.itemPool.Get(req.DeploymentOrderID).(*data.DeploymentItem)
	if item == nil && item.State != data.DeployingState {
		r.Response.WriteJsonExit(data.DeployToProviderResponce{
			Code:  3,
			Error: "the deployment order is not at deploying state",
			Data:  nil,
		})
	}

	// auth
	if req.AccessKey != item.Info.AccessKey {
		r.Response.WriteJsonExit(data.DeployToProviderResponce{
			Code:  4,
			Error: "wrong access key",
			Data:  nil,
		})
	}

	// deploy on provider
	// TODO
	// adapter

	r.Response.WriteJsonExit(data.DeployToProviderResponce{
		Code:  0,
		Error: "",
		Data:  nil,
	})
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
