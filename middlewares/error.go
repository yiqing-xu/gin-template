/*
@Time : 2020/6/28 22:50
@Author : xuyiqing
@File : error.py
*/

package middlewares

import (
	"gin-template/handlers"
	"github.com/gin-gonic/gin"
	"runtime/debug"
)

// 捕获程序报错异常栈
func ErrorHandleMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				resp := handlers.Response{Ctx: ctx}
				resp.ServerError(string(debug.Stack()))
			}
		}()
		ctx.Next()
	}
}

