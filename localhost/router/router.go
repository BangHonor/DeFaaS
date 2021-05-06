package router

import (
	"defaas/localhost/app/api/account"
	"defaas/localhost/app/api/faaslevel"
	"defaas/localhost/app/api/faastoken"
	"defaas/localhost/app/api/provider"
	"defaas/localhost/app/api/witness"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {

	s := g.Server()

	// 中间件
	// s.Use(middleware.AccessLog, middleware.ErrorHandler)

	{
		s.Group("/v1/api", func(group *ghttp.RouterGroup) {

			group.Group("/account", func(accountsGroup *ghttp.RouterGroup) {
				api := new(account.AccountAPI)
				accountsGroup.POST("/create", api.Create)
				accountsGroup.GET("/list", api.List)
			})

			group.Group("/faastoken", func(faastokenGroup *ghttp.RouterGroup) {
				api := new(faastoken.FaaSTokenAPI)
				faastokenGroup.GET("/balanceOf", api.BalanceOf)
			})

			group.Group("/faaslevel", func(faaslevelGroup *ghttp.RouterGroup) {
				api := new(faaslevel.FaaSLevelAPI)
				faaslevelGroup.GET("/list", api.List)
				faaslevelGroup.POST("/add", api.Add)

			})

			group.Group("/witness", func(witnessGroup *ghttp.RouterGroup) {
				api := new(witness.WitnessnAPI)
				witnessGroup.POST("/login", api.Login)
				witnessGroup.POST("/logout", api.Logout)
				witnessGroup.POST("/turn-on", api.TurnOn)
				witnessGroup.POST("/turn-off", api.TurnOff)
			})

			group.Group("/provider", func(providerGroup *ghttp.RouterGroup) {
				api := new(provider.ProviderAPI)
				providerGroup.POST("/login", api.Login)
				providerGroup.POST("/logout", api.Logout)
			})

		})
	}
}
