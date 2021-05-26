package funccode

import (
	"defaas/localhost/app/api/response"
	"defaas/localhost/app/model"
	"defaas/localhost/app/service/funccodesvc"

	"github.com/gogf/gf/net/ghttp"
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
	Name  string
	Tag   string
	Files []struct {
		Filename string `param:"filename" v:"required"`
		Language string `param:"language" v:"required"`
		Code     string `param:"code" v:"required"`
	}
	//读入前端传递的数据
}

type FunccodeAddRes model.FunccodeItem

//后台数据格式

func (a *FunccodeAPI) Add(r *ghttp.Request) {
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

		Name: apiReq.Name,
		Tag:  apiReq.Tag,
		Files: []struct {
			Filename string "json:\"filename\""
			Language string "json:\"language\""
			Code     string "json:\"code\""
		}(apiReq.Files),
	}

	item, err := funccodesvc.Service().Add(item)
	if err != nil {
		response.JSONExit(r, 1, err.Error())
	}

	apiRes = FunccodeAddRes(item)

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

		Name: apiReq.Name,
		Tag:  apiReq.Tag,
		Files: []struct {
			Filename string "json:\"filename\""
			Language string "json:\"language\""
			Code     string "json:\"code\""
		}(apiReq.Files),
	}

	item, err := funccodesvc.Service().Delete(item)
	if err != nil {
		response.JSONExit(r, 1, err.Error())
	}

	apiRes = FunccodeAddRes(item)

	response.JSONExit(r, 0, "ok", apiRes)
}