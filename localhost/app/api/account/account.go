package account

import (
	"defaas/localhost/app/api/response"

	"github.com/gogf/gf/net/ghttp"
)

type AccountAPI struct{}

func (a *AccountAPI) New(r *ghttp.Request) {
	response.JSONExit(r, 0, "ok", nil)
}

func (a *AccountAPI) List(r *ghttp.Request) {
	response.JSONExit(r, 0, "ok", nil)
}
