package account

import (
	"defaas/localhost/app/api/response"
	"defaas/localhost/app/service/accountsvc"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type AccountAPI struct{}

func (a *AccountAPI) Create(r *ghttp.Request) {

	password, ok := r.Get("password").(string)
	if !ok {
		response.JSONExit(r, 1, "wrong [password] param")
	}

	item, err := accountsvc.Service().CreateAccount(password)
	if err != nil {
		response.JSONExit(r, 2, err.Error())
	}

	response.JSONExit(r, 0, "ok", g.Map{
		"address":  item.Address,
		"password": item.Password,
	})
}

func (a *AccountAPI) List(r *ghttp.Request) {
	response.JSONExit(r, 0, "ok", nil)
}
