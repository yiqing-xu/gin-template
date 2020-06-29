/*
@Time : 2020/6/28 21:35
@Author : xuyiqing
@File : users.py
*/

package routers

import (
	"gin-template/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterUsersRouter(group *gin.RouterGroup)  {
	group.POST("/login", handlers.UsersLoginHandler)
	group.POST("/register", handlers.UsersRegisterHandler)
}

func RegisterUsersRouterWithAuth(group *gin.RouterGroup) {
	group.PUT("/userinfo", handlers.UsersSetInfoHandler)
}
