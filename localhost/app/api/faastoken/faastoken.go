package faastoken

import (
	"defaas/localhost/app/api/response"

	"github.com/gogf/gf/net/ghttp"
)

type FaaSTokenAPI struct{}

func (a *FaaSTokenAPI) BalanceOf(r *ghttp.Request) {
	response.JSONExit(r, 0, "ok", nil)
}

func (a *FaaSTokenAPI) Mind(r *ghttp.Request) {
	response.JSONExit(r, 0, "ok", nil)
}

func (a *FaaSTokenAPI) Transfer(r *ghttp.Request) {
	response.JSONExit(r, 0, "ok", nil)
}
