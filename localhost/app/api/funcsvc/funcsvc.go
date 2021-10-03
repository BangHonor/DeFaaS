package funcsvc

import (
	devutils "defaas/dev-cmd/utils"
	"defaas/localhost/app/api/response"
	"defaas/localhost/app/model"
	"defaas/localhost/app/service/funcsvcsvc"

	"bufio"
	"fmt"
	"github.com/go-cmd/cmd"
	"github.com/gogf/gf/net/ghttp"
	"os"
)

type FuncsvcAPI struct{}

// ----------------------------------------------------------------------------------------------------------------

type FuncsvcListRes []model.FuncsvcItem

// ----------------------------------------------------------------------------------------------------------------

type FuncsvcAddReq struct {
	Name                 string `param:"name"`
	AccountAddress       string `param:"accountAddress"`
	FunccodeName         string `param:"funccodeName"` //知道funncodename直接获取相关信息
	FaaslevelID          string `param:"faaslevelID"`
	BidDuration          string `param:"bidDuration"`
	ServiceDuration      string `param:"serviceDuration"`
	HighestUnitPrice     string `param:"highestUnitPrice"`
	UnitPrice            string `param:"unitPrice"`
	ServiceFee           string `param:"serviceFee"`
	DeploymentOrderID    string `param:"deploymentOrderID"`
	DeploymentOrderState string `param:"deploymentOrderState"`
	//读入前端传递的数据
}

type FuncsvcAddRes model.FuncsvcItem

//后台数据格式
func (a *FuncsvcAPI) List(r *ghttp.Request) {
	//获取Service数据
	levels, err := funcsvcsvc.Service().List()
	if err != nil {
		response.JSONExit(r, 1, err.Error())
	}

	apiRes := FuncsvcListRes(levels)

	response.JSONExit(r, 0, "ok", apiRes)
}

func (a *FuncsvcAPI) Add(r *ghttp.Request) {
	//请求添加数据req
	var (
		apiReq FuncsvcAddReq
		apiRes FuncsvcAddRes
	)

	if err := r.Parse(&apiReq); err != nil {
		response.JSONExit(r, 1, "wrong param "+err.Error())
	}

	item := model.FuncsvcItem{
		//从前端存储到后端
		Name:                 apiReq.Name,
		AccountAddress:       apiReq.AccountAddress,
		FunccodeName:         apiReq.FunccodeName, //知道funncodename直接获取相关信息
		FaaslevelID:          apiReq.FaaslevelID,
		BidDuration:          apiReq.BidDuration,
		ServiceDuration:      apiReq.ServiceDuration,
		HighestUnitPrice:     apiReq.Name,
		UnitPrice:            apiReq.UnitPrice,
		ServiceFee:           apiReq.ServiceFee,
		DeploymentOrderID:    apiReq.DeploymentOrderID,
		DeploymentOrderState: apiReq.DeploymentOrderState,
	}

	item, err := funcsvcsvc.Service().Add(item)
	if err != nil {
		response.JSONExit(r, 1, err.Error())
	}

	apiRes = FuncsvcAddRes(item)

	//openfaas

	faasbuild := cmd.NewCmd(
		"faas-cli",
		"build",
		"-f",
		"./"+apiReq.FunccodeName+".yml")

	dockerpull := cmd.NewCmd(
		"docker",
		"push",
		"honorbang/"+apiReq.FunccodeName)

	chmod := cmd.NewCmd(
		"chmod", "-R", "777", "/home/dds/kitchen/defaas/localhost/server/"+apiReq.FunccodeName+".yml")

	attachLabel := cmd.NewCmd(
		"kubectl", "label", "node/lanjing5", "--overwrite",
		apiReq.FunccodeName+"=true")

	unattachLabel := cmd.NewCmd(
		"kubectl", "label", "node/lanjing5", "--overwrite",
		apiReq.FunccodeName+"=false")

	//根据实际情况调用机器

	faasdeploy := cmd.NewCmd(
		"faas-cli",
		"deploy",
		"-f",
		"./"+apiReq.FunccodeName+".yml",
		"--gateway",
		"http://10.186.133.126:31112")

	go func() {
		devutils.RunCmd(faasbuild)
		devutils.RunCmd(dockerpull)
		devutils.RunCmd(chmod)
		devutils.RunCmd(attachLabel)

		name := "/home/dds/kitchen/defaas/localhost/server/" + apiReq.FunccodeName + ".yml"
		//选择node
		file, err := os.OpenFile(name, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("文件打开失败", err)
		}
		//及时关闭file句柄
		defer file.Close()
		//写入文件时，使用带缓存的 *Writer
		write := bufio.NewWriter(file)
		write.WriteString("    constraints:\n")
		write.WriteString(`    - "` + apiReq.FunccodeName + `=true"` + "\n")
		write.WriteString("    limits:\n")
		write.WriteString("      memory: 40Mi\n")
		write.WriteString("      cpu: 100m")
		//Flush将缓存的文件真正写入到文件中
		write.Flush()

		devutils.RunCmd(faasdeploy)
		devutils.RunCmd(unattachLabel)
		//在root状态下,先实现登录faas-cli login，然后export OPENFAAS_URL=http://10.186.133.126:31112实现默认接口
	}()
	response.JSONExit(r, 0, "ok", apiRes)
}

func (a *FuncsvcAPI) Delete(r *ghttp.Request) {
	//请求添加数据req
	var (
		apiReq FuncsvcAddReq
		apiRes FuncsvcAddRes
	)

	if err := r.Parse(&apiReq); err != nil {
		response.JSONExit(r, 1, "wrong param "+err.Error())
	}

	item := model.FuncsvcItem{
		//从前端存储到后端
		Name:                 apiReq.Name,
		AccountAddress:       apiReq.AccountAddress,
		FunccodeName:         apiReq.FunccodeName, //知道funncodename直接获取相关信息
		FaaslevelID:          apiReq.FaaslevelID,
		BidDuration:          apiReq.BidDuration,
		ServiceDuration:      apiReq.ServiceDuration,
		HighestUnitPrice:     apiReq.Name,
		UnitPrice:            apiReq.UnitPrice,
		ServiceFee:           apiReq.ServiceFee,
		DeploymentOrderID:    apiReq.DeploymentOrderID,
		DeploymentOrderState: apiReq.DeploymentOrderState,
	}

	item, err := funcsvcsvc.Service().Delete(item)
	if err != nil {
		response.JSONExit(r, 1, err.Error())
	}

	apiRes = FuncsvcAddRes(item)

	response.JSONExit(r, 0, "ok", apiRes)
}
