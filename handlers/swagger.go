/*
@Time : 2020/6/29 15:56
@Author : xuyiqing
@File : swagger.py
*/

package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CustomSwaggerIndexHandler(ctx *gin.Context) {
	doc := ctx.DefaultQuery("doc", "swagger")
	path := fmt.Sprintf("/assets/docs/%s.yaml", doc)
	ctx.HTML(200, "index.html", gin.H{
		"doc": path,
	})
}
