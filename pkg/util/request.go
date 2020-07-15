/*
@Time : 2020/6/29 14:40
@Author : xuyiqing
@File : body.py
*/

package util

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
)

// 反序列化request.body中的json数据为map
func GetBodyData(ctx *gin.Context) map[string]interface{} {
	bdata := make([]byte, 1024)
	length, err := ctx.Request.Body.Read(bdata)
	if err != nil && err != io.EOF {
		return nil
	}
	var data map[string]interface{}
	if err := json.Unmarshal(bdata[:length], &data); err != nil {
		return nil
	}
	return data
}

// 构建文件url连接主机端口全链接 "https://192.168.11.121:7889/meida/upload/..."
func BuildAbsoluteUri(ctx *gin.Context, filePath string) string {
	host := ctx.Request.Host
	schema := ctx.Request.Header.Get("X-Forwarded-Proto")
	if schema == "https" {
		return fmt.Sprintf("https://%s/%s", host, filePath)
	} else {
		return fmt.Sprintf("http://%s/%s", host, filePath)
	}
}
