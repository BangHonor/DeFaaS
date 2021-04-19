package faaslevel

import (
	"defaas/localhost/app/api/response"

	"github.com/gogf/gf/net/ghttp"
)

type FaaSLevelAPI struct{}

func (a *FaaSLevelAPI) List(r *ghttp.Request) {
	response.JSONExit(r, 0, "ok", nil)
}

func (a *FaaSLevelAPI) Add(r *ghttp.Request) {
	response.JSONExit(r, 0, "ok", nil)
}
