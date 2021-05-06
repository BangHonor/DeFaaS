package main

import (
	_ "defaas/localhost/boot"
	_ "defaas/localhost/router"

	"github.com/gogf/gf/frame/g"
)

func main() {

	// 开启debug模式,打印更多调试信息,建议在开发环境开启
	g.SetDebug(true)

	s := g.Server()

	// 设置配置
	// https://goframe.org/pages/viewpage.action?pageId=1114489
	s.SetConfigWithMap(g.Map{
		"address":          ":55555",
		"accessLogEnabled": true,
		"errorLogEnabled":  true,
		"pprofEnabled":     true,
	})

	s.Run()
}
