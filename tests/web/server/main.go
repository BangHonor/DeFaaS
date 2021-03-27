package main

import (
	"defaas/core/data"
	"encoding/json"
	"fmt"

	"github.com/gogf/gf/net/ghttp"
)

func main() {

	serverAddr := "127.0.0.1:60066"

	serverEntry := "/deploy"

	server := ghttp.GetServer("deploy-server")
	server.SetAddr(serverAddr)

	server.BindHandler(serverEntry, func(r *ghttp.Request) {

		req := &data.DeployToProviderRequest{}

		reqJson := r.GetBody()

		if err := json.Unmarshal(reqJson, req); err != nil {

			r.Response.WriteJsonExit(data.DeployToProviderResponce{
				Code:  1,
				Error: err.Error(),
				Data:  nil,
			})

		}

		fmt.Println(req)

		r.Response.WriteJsonExit(data.DeployToProviderResponce{
			Code:  0,
			Error: "",
			Data:  nil,
		})
	})

	server.Run()
}
