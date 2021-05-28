package response

import "github.com/gogf/gf/net/ghttp"

// api HTTP 响应的 JSON 数据结构标准

// JSONResponse 数据返回通用JSON 数据结构
type JSONResponse struct {
	Code    int         `json:"code"` // 错误码((0:成功, >0:错误码))
	Message string      `json:"msg"`  // 提示信息
	Data    interface{} `json:"data"` // 返回数据(业务接口定义具体数据结构)
}

// JSON 标准返回结果数据结构封装。
func JSON(r *ghttp.Request, code int, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.Response.WriteJson(JSONResponse{
		Code:    code,
		Message: message,
		Data:    responseData,
	})
}

// JSONExit 返回JSON数据并退出当前HTTP执行函数。
// 仅退出当前执行的逻辑方法，不退出后续的请求流程，可用于替代return。
// 数据返回-Exit控制 https://goframe.org/pages/viewpage.action?pageId=1114204
func JSONExit(r *ghttp.Request, err int, msg string, data ...interface{}) {
	JSON(r, err, msg, data...)

	r.Exit()
}
