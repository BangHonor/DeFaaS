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

	levels, err := funccodesvc.Service().List()
	if err != nil {
		response.JSONExit(r, 1, err.Error())
	}

	apiRes := FunccodeListRes(levels)

	response.JSONExit(r, 0, "ok", apiRes)
}

// ----------------------------------------------------------------------------------------------------------------

type FunccodeAddReq struct {
	Files struct {
		Filename string `param:"filename" v:"required"`
		Language string `param:"language" v:"required"`
		Code     string `param:"code" v:"required"`
	}
	Name string `json:"name" v:"required"`
	Tag  string `json:"tag" v:"required"`
}

type FunccodeAddRes model.FunccodeItem

func (a *FunccodeAPI) Add(r *ghttp.Request) {

	var (
		apiReq FunccodeAddReq
		apiRes FunccodeAddRes
	)

	if err := r.Parse(&apiReq); err != nil {
		response.JSONExit(r, 1, "wrong param "+err.Error())
	}

	item := model.FunccodeItem{
		Name:  apiRes.Name,
		Tag:   apiRes.Tag,
		Files: apiRes.Files,
	}

	item, err := funccodesvc.Service().Add(item)
	if err != nil {
		response.JSONExit(r, 1, err.Error())
	}

	apiRes = FunccodeAddRes(item)

	response.JSONExit(r, 0, "ok", apiRes)
}
