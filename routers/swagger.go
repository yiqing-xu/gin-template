/*
@Time : 2020/6/29 19:11
@Author : xuyiqing
@File : swagger.py
*/

package routers

import (
	"gin-template/handlers"
	"github.com/gin-gonic/gin"

)

func RegisterSwaggerRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/doc", handlers.CustomSwaggerIndexHandler)
}
