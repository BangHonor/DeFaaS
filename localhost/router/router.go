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
		s.Group("/api", func(group *ghttp.RouterGroup) {

			group.Group("/account", func(accountsGroup *ghttp.RouterGroup) {

				accountAPI := new(account.AccountAPI)

				accountsGroup.POST("/create", accountAPI.Create)
				accountsGroup.GET("/list", accountAPI.List)

				accountsGroup.Group("/witness", func(witnessGroup *ghttp.RouterGroup) {
					wintessAPI := new(witness.WitnessnAPI)
					witnessGroup.POST("/login", wintessAPI.Login)
					witnessGroup.POST("/logout", wintessAPI.Logout)
					witnessGroup.POST("/turn-on", wintessAPI.TurnOn)
					witnessGroup.POST("/turn-off", wintessAPI.TurnOff)
				})

				accountsGroup.Group("/provider", func(providerGroup *ghttp.RouterGroup) {
					providerAPI := new(provider.ProviderAPI)
					providerGroup.POST("/login", providerAPI.Login)
					providerGroup.POST("/logout", providerAPI.Logout)
				})

			})

			group.Group("/faastoken", func(faastokenGroup *ghttp.RouterGroup) {

				faastokenAPI := new(faastoken.FaaSTokenAPI)

				faastokenGroup.GET("/balanceOf", faastokenAPI.BalanceOf)
				faastokenGroup.POST("/mind", faastokenAPI.BalanceOf)
				faastokenGroup.POST("/transfer", faastokenAPI.BalanceOf)
			})

			group.Group("/faaslevel", func(faaslevelGroup *ghttp.RouterGroup) {

				api := new(faaslevel.FaaSLevelAPI)

				faaslevelGroup.POST("/create", api.Add)
				faaslevelGroup.GET("/list", api.List)
			})

		})
	}
}
