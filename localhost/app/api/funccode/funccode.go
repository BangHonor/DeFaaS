package funccode

import (
	devutils "defaas/dev-cmd/utils"
	"defaas/localhost/app/api/response"
	"defaas/localhost/app/model"
	"defaas/localhost/app/service/funccodesvc"
	"github.com/go-cmd/cmd"
	"github.com/gogf/gf/net/ghttp"
	"io/ioutil"
)

type FunccodeAPI struct{}

// ----------------------------------------------------------------------------------------------------------------

type FunccodeListRes []model.FunccodeItem

func (a *FunccodeAPI) List(r *ghttp.Request) {
	//获取Service数据
	levels, err := funccodesvc.Service().List()
	if err != nil {
		response.JSONExit(r, 1, err.Error())
	}

	apiRes := FunccodeListRes(levels)

	response.JSONExit(r, 0, "ok", apiRes)
}

// ----------------------------------------------------------------------------------------------------------------

type FunccodeAddReq struct {
	Name     string `param:"name"`
	Tag      string `param:"tag"`
	Filename string `param:"filename" v:"required"`
	Language string `param:"language" v:"required"`
	Code     string `param:"code" v:"required"`

	//读入前端传递的数据
}

type FunccodeAddRes model.FunccodeItem

//后台数据格式

func (a *FunccodeAPI) Add(r *ghttp.Request) {

	Map := map[string]string{
		"python": "py",
		"go":     "go",
		"java":   "jsp",
	}

	//请求添加数据req
	var (
		apiReq FunccodeAddReq
		apiRes FunccodeAddRes
	)
	if err := r.Parse(&apiReq); err != nil {
		response.JSONExit(r, 1, "wrong param "+err.Error())
	}

	//openfaas放在前面，不然获得数据后直接就中断了，轮不到faas-cli

	faasnew := cmd.NewCmd(
		"faas-cli",
		"new", apiReq.Name,
		"--handler", "../../func/"+apiReq.Name, //文件相对比较大，转移地方
		"--lang", apiReq.Language,
		"-p", "honorbang")

	chmod := cmd.NewCmd(
		"chmod", "-R", "777", "/home/dds/kitchen/defaas/func/"+apiReq.Name+"/handler."+Map[apiReq.Language])
	//修改handler文件内容，改为自己的函数
	editFunc := cmd.NewCmd(
		"sed", "-i", "1,10d", "/home/dds/kitchen/defaas/func/"+apiReq.Name+"/handler."+Map[apiReq.Language])

	go func() {
		devutils.RunCmd(faasnew)
		devutils.RunCmd(chmod)
		devutils.RunCmd(editFunc)
		name := "/home/dds/kitchen/defaas/func/" + apiReq.Name + "/handler." + Map[apiReq.Language]
		content := []byte(apiReq.Code)
		ioutil.WriteFile(name, content, 0644)
	}()

	item := model.FunccodeItem{
		//从前端存储到后端

		Name:     apiReq.Name,
		Tag:      apiReq.Tag,
		Filename: apiReq.Filename,
		Language: apiReq.Language,
		Code:     apiReq.Code,
	}

	res, err := funccodesvc.Service().Add(item)
	if err != nil {
		response.JSONExit(r, 1, err.Error())
	}

	apiRes = FunccodeAddRes(res)
	response.JSONExit(r, 0, "ok", apiRes)
}

func (a *FunccodeAPI) Delete(r *ghttp.Request) {
	//请求添加数据req
	var (
		apiReq FunccodeAddReq
		apiRes FunccodeAddRes
	)

	if err := r.Parse(&apiReq); err != nil {
		response.JSONExit(r, 1, "wrong param "+err.Error())
	}

	item := model.FunccodeItem{
		//从前端存储到后端

		Name:     apiReq.Name,
		Tag:      apiReq.Tag,
		Filename: apiReq.Filename,
		Language: apiReq.Language,
		Code:     apiReq.Code,
	}

	item, err := funccodesvc.Service().Delete(item)
	if err != nil {
		response.JSONExit(r, 1, err.Error())
	}

	apiRes = FunccodeAddRes(item)

	response.JSONExit(r, 0, "ok", apiRes)
}
