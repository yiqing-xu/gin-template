/*
@Time : 2020/6/29 14:40
@Author : xuyiqing
@File : body.py
*/

package util

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
)

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
