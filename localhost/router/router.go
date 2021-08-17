package router

import (
	"defaas/localhost/app/api/account"
	"defaas/localhost/app/api/faaslevel"
	"defaas/localhost/app/api/faastoken"
	"defaas/localhost/app/api/funccode"

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
					witnessGroup.POST("/login", accountAPI.WitnessLogin)
					witnessGroup.POST("/logout", accountAPI.WitnessLogout)
					witnessGroup.POST("/turn-on", accountAPI.WitnessTurnOn)
					witnessGroup.POST("/turn-off", accountAPI.WitnessTurnOff)
				})

				accountsGroup.Group("/provider", func(providerGroup *ghttp.RouterGroup) {
					providerGroup.POST("/login", accountAPI.ProviderLogin)
					providerGroup.POST("/logout", accountAPI.ProviderLogout)
				})

			})

			group.Group("/faastoken", func(faastokenGroup *ghttp.RouterGroup) {

				faastokenAPI := new(faastoken.FaaSTokenAPI)

				faastokenGroup.GET("/balanceOf", faastokenAPI.BalanceOf)
				faastokenGroup.POST("/mind", faastokenAPI.BalanceOf)
				faastokenGroup.POST("/transfer", faastokenAPI.BalanceOf)
			})

			group.Group("/faaslevel", func(faaslevelGroup *ghttp.RouterGroup) {

				faaslevelAPI := new(faaslevel.FaaSLevelAPI)

				faaslevelGroup.POST("/add", faaslevelAPI.Add)
				faaslevelGroup.GET("/list", faaslevelAPI.List)
			})
			
			group.Group("/funccode", func(funccodeGroup *ghttp.RouterGroup) {

				funccodeAPI := new(funccode.FunccodeAPI)

				funccodeGroup.POST("/add", funccodeAPI.Add)
				funccodeGroup.POST("/delete", funccodeAPI.Delete)
				funccodeGroup.GET("/list", funccodeAPI.List)
			})

		})
	}
}
