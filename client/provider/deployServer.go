package provider

import (
	"fmt"
	"log"

	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/net/ghttp"
)

type DeployServer struct {
	//
	DeployingPool gmap.AnyAnyMap
	//
	adapter string // deployment adapter
	//
	serverAddr  string // serverAddr is like '0.0.0.0:80', '127.0.0.1:80', '180.18.99.10:80', etc
	serverEntry string // serverEntry is the entry path of server, like '127.0.0.1:80/{serverEntry}'
	//
	server *ghttp.Server
}

func ParseDeployPath(deployPath string) (string, string, string) {

	// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	// Note that it is a simplification in DEV, NOW
	adapter := "docker"
	serverAddr := "127.0.0.1:60666"
	serverEntry := "/deploy"
	// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

	return adapter, serverAddr, serverEntry

	// TODO
	// deployPath format
	// [adpater:tag][serverAddr/serverEntry]
	// like: [docker:tag][http:127.0.0.1:80/deploy]

}

func NewDeployServer(deployPath string, serverID int) (*DeployServer, error) {

	ds := &DeployServer{}

	adapter, serverAddr, serverEntry := ParseDeployPath(deployPath)

	ds.adapter = adapter
	ds.serverAddr = serverAddr
	ds.serverEntry = serverEntry

	ds.server = ghttp.GetServer("deploy-server-" + fmt.Sprint(serverID))
	ds.server.SetAddr(ds.serverAddr)                               // https://pkg.go.dev/github.com/gogf/gf/net/ghttp#Server.SetAddr
	ds.server.BindHandler(ds.serverEntry, ds.DeployRequestHandler) // register handler

	return ds, nil
}

func (ds *DeployServer) Start() error {

	go ds.server.Run()

	log.Printf("[provider/deployServer] run deploy server [%s]\n", ds.serverAddr+ds.serverEntry)

	return nil
}

func (ds *DeployServer) Stop() error {

	err := ds.server.Shutdown()
	if err != nil {
		log.Printf("[provider/deployServer] failed to shutdown deploy server [%s]\n", ds.serverAddr+ds.serverEntry)
		return err
	}

	log.Printf("[provider/deployServer] shutdown deploy server [%s]\n", ds.serverAddr+ds.serverEntry)

	return nil
}

func (ds *DeployServer) DeployRequestHandler(r *ghttp.Request) {

	r.Response.Write("Hello")

	// TODO

	// 1. parse parameter from request
	// 2. check is at pool
	// 3. auth
	// 4. use adpater
	// 4.1 docker adapter
}

func (ds *DeployServer) NewListeningDeployTask(funcPath string, accessKey string, notify chan<- bool) error {

	// add to accessKey pool

	// add to norify pool

	return nil
}
