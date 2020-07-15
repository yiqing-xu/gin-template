/*
@Time : 2020/7/15 22:18
@Author : xuyiqing
@File : cms.py
*/

package routers

import (
	"gin-template/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterCmsRouter(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/file", handlers.UploadFileHandler)
	routerGroup.GET("/file/:id", handlers.DownloadFileHandler)
}
