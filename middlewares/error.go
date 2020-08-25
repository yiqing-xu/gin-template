/*
@Time : 2020/6/28 22:50
@Author : xuyiqing
@File : error.go
*/

package middlewares

import (
	"gin-template/handlers"
	"github.com/gin-gonic/gin"
	"runtime/debug"
)

// 捕获程序报错异常栈
func ErrorHandleMiddleware(isErrMsg bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				resp := handlers.Response{Ctx: ctx}
				if isErrMsg {
					resp.ServerError(string(debug.Stack()))
				} else {
					resp.ServerError("系统出错")
				}

			}
		}()
		ctx.Next()
	}
}
