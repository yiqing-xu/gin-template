/*
@Time : 2020/7/16 12:49
@Author : xuyiqing
@File : ws.py
*/

package routers

import (
	"gin-template/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterWsHandler(routerGroup *gin.RouterGroup) {
	routerGroup.GET("ws", handlers.WsMessageHandler)
}

func RegisterWsAuthHandler(routerGroup *gin.RouterGroup) {
	routerGroup.GET("ws/test", handlers.TestWsMessageHandler)
	routerGroup.GET("ws/messages/:id", handlers.GetWsMessageHandler)
}