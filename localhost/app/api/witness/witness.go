package witness

import (
	"defaas/localhost/app/api/response"

	"github.com/gogf/gf/net/ghttp"
)

type WitnessnAPI struct{}

func (w *WitnessnAPI) Login(r *ghttp.Request) {
	response.JSONExit(r, 0, "ok", nil)
}

func (w *WitnessnAPI) Logout(r *ghttp.Request) {
	response.JSONExit(r, 0, "ok", nil)
}

func (w *WitnessnAPI) TurnOn(r *ghttp.Request) {
	response.JSONExit(r, 0, "ok", nil)
}

func (w *WitnessnAPI) TurnOff(r *ghttp.Request) {
	response.JSONExit(r, 0, "ok", nil)
}
