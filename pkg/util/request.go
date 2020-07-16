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
	"strings"
)

// 反序列化request.body中的json数据为map
func GetBodyData(ctx *gin.Context) (map[string]interface{}, error) {
	bdata := make([]byte, 1024)
	length, err := ctx.Request.Body.Read(bdata)
	if err != nil && err != io.EOF {
		return nil, err
	}
	var data map[string]interface{}
	str := string(bdata[:length])
	decoder := json.NewDecoder(strings.NewReader(str))
	decoder.UseNumber()
	err1 := decoder.Decode(&data)
	if err1 != nil {
		return nil, err
	}
	return data, nil
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
