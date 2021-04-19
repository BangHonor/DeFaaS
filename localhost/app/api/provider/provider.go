package provider

import (
	"defaas/localhost/app/api/response"

	"github.com/gogf/gf/net/ghttp"
)

type ProviderAPI struct{}

func (w *ProviderAPI) Login(r *ghttp.Request) {
	response.JSONExit(r, 0, "ok", nil)
}

func (w *ProviderAPI) Logout(r *ghttp.Request) {
	response.JSONExit(r, 0, "ok", nil)
}
