package faaslevel

import (
	"defaas/localhost/app/api/response"
	"defaas/localhost/app/model"
	"defaas/localhost/app/service/faaslevelsvc"

	"github.com/gogf/gf/net/ghttp"
)

type FaaSLevelAPI struct{}

// ----------------------------------------------------------------------------------------------------------------

type FaaslevelListRes []model.FaaslevelItem

func (a *FaaSLevelAPI) List(r *ghttp.Request) {

	levels, err := faaslevelsvc.Service().List()
	if err != nil {
		response.JSONExit(r, 1, err.Error())
	}

	apiRes := FaaslevelListRes(levels)

	response.JSONExit(r, 0, "ok", apiRes)
}

// ----------------------------------------------------------------------------------------------------------------

type FaaslevelAddReq struct {
	CPU string `param:"cpu" v:"required"`
	Mem string `param:"mem" v:"required"`
}

type FaaslevelAddRes model.FaaslevelItem

func (a *FaaSLevelAPI) Add(r *ghttp.Request) {

	var (
		apiReq FaaslevelAddReq
		apiRes FaaslevelAddRes
	)

	if err := r.Parse(&apiReq); err != nil {
		response.JSONExit(r, 1, "wrong param "+err.Error())
	}

	item := model.FaaslevelItem{
		ID:  "",
		CPU: apiReq.CPU,
		Mem: apiReq.Mem,
	}

	item, err := faaslevelsvc.Service().Add(item)
	if err != nil {
		response.JSONExit(r, 1, err.Error())
	}

	apiRes = FaaslevelAddRes(item)

	response.JSONExit(r, 0, "ok", apiRes)
}
