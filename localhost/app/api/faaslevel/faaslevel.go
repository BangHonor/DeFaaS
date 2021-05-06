package faaslevel

import (
	"defaas/localhost/app/api/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type FaaSLevelAPI struct{}

func (a *FaaSLevelAPI) List(r *ghttp.Request) {
	response.JSONExit(r, 0, "ok", g.Map{
		"levels": []g.Map{
			{
				"id":  "0",
				"cpu": "1",
				"mem": "512",
			},
			{
				"id":  "1",
				"cpu": "2",
				"mem": "1024",
			},
		},
	})
}

func (a *FaaSLevelAPI) Add(r *ghttp.Request) {
	response.JSONExit(r, 0, "ok", nil)
}
