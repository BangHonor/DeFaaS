package account

import (
	"defaas/localhost/app/api/response"
	"defaas/localhost/app/model"
	"defaas/localhost/app/service/accountsvc"

	"github.com/gogf/gf/net/ghttp"
)

type AccountAPI struct{}

// ----------------------------------------------------------------------------------------------------------------

type AccountCreateReq struct {
	Password string `param:"password" v:"required"`
}

type AccountCreateRes model.AccountItem

func (a *AccountAPI) Create(r *ghttp.Request) {

	var (
		apiReq AccountCreateReq
		apiRes AccountCreateRes
	)

	if err := r.Parse(&apiReq); err != nil {
		response.JSONExit(r, 1, "wrong param "+err.Error())
	}

	item, err := accountsvc.Service().Create(apiReq.Password)
	if err != nil {
		response.JSONExit(r, 1, err.Error())
	}

	apiRes = AccountCreateRes(*item)
	response.JSONExit(r, 0, "ok", apiRes)
}

// ----------------------------------------------------------------------------------------------------------------

func (a *AccountAPI) List(r *ghttp.Request) {
	response.JSONExit(r, 0, "ok", nil)
}

func (a *AccountAPI) WitnessLogin(r *ghttp.Request) {
	response.JSONExit(r, 0, "ok", nil)
}

func (a *AccountAPI) WitnessLogout(r *ghttp.Request) {
	response.JSONExit(r, 0, "ok", nil)
}

func (a *AccountAPI) WitnessTurnOn(r *ghttp.Request) {
	response.JSONExit(r, 0, "ok", nil)
}

func (a *AccountAPI) WitnessTurnOff(r *ghttp.Request) {
	response.JSONExit(r, 0, "ok", nil)
}

func (a *AccountAPI) ProviderLogin(r *ghttp.Request) {
	response.JSONExit(r, 0, "ok", nil)
}

func (a *AccountAPI) ProviderLogout(r *ghttp.Request) {
	response.JSONExit(r, 0, "ok", nil)
}
