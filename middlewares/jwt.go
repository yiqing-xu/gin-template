/*
@Time : 2020/6/29 11:16
@Author : xuyiqing
@File : jwt.py
*/

package middlewares

import (
	"gin-template/handlers"
	"gin-template/pkg/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

// jwt token验证有效性，并赋值ctx当前会话用户
func AuthJwtTokenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//method := ctx.Request.Method
		resp := handlers.Response{Ctx: ctx}
		tokenStr := ctx.Request.Header.Get("Authentication")
		if len([]rune(tokenStr)) == 0 {
			resp.Unauthenticated("未登录验证")
			return
		}
		strs := strings.Split(tokenStr, " ")
		if strs[0] != "Bearer" {
			resp.BadRequest("token格式不正确，${Bearer token}")
		}
		claims, err := jwt.ValidateJwtToken(strs[1])
		if err != nil {
			resp.Unauthenticated("未登录验证，" + err.Error())
			return
		} else {
			CurrentUser := claims.GetUserByID()
			if CurrentUser != nil {
				ctx.Set("CurrentUser", CurrentUser)
			} else {
				resp.Unauthenticated("查询不到token对应用户")
				return
			}
		}
		ctx.Next()
	}
}
