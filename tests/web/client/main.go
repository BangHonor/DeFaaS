package main

import (
	"defaas/core/data"
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/net/ghttp"
)

func main() {

	url := "http://" + "127.0.0.1:60066" + "/deploy"

	requestData := &data.DeployToProviderRequest{
		CustomerAddress:   common.HexToAddress("0x01"),
		DeploymentOrderID: big.NewInt(666),
		AccessKey:         "accessKey",
		FulfillSecretKey:  data.GenerateFulfillSecretKey(),
		Adapter:           "docker",
		AdapterData:       []byte(`hello`),
	}

	reqJson, err := json.Marshal(requestData)
	if err != nil {
		log.Fatal(err)
	}

	client := ghttp.NewClient()
	httpRes, err := client.ContentJson().Post(url, reqJson)
	if err != nil {
		log.Fatal(err)
	}
	defer httpRes.Close()

	log.Println("http code", httpRes.Status)

	responseData := &data.DeployToProviderResponce{}
	if err := json.Unmarshal([]byte(httpRes.ReadAllString()), responseData); err != nil {
		log.Fatal(err)
	}
	fmt.Println(responseData)
}
