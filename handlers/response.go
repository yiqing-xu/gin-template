/*
@Time : 2020/6/28 22:37
@Author : xuyiqing
@File : response.py
*/

package handlers

import (
	"github.com/gin-gonic/gin"
)

const (
	Success		 	= 200
	BadRequest 		= 400
	Unauthenticated = 401
	NoPermisson 	= 403
	NotFund 		= 404
	ServerError 	= 500
)

type Response struct {
	Ctx *gin.Context
}

// 定义基础返回结构体
type JsonResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data,omitempty"`
	Pager interface{} `json:"pager,omitempty"`
}

func (resp *Response) Response(data interface{}, pager interface{}) {
	resp.Ctx.JSON(Success, JsonResponse{
		Code: Success,
		Msg: "返回成功",
		Data: data,
		Pager: pager,
	})
}

// 400错误请求
func (resp *Response) BadRequest(msg string) {
	resp.Ctx.AbortWithStatusJSON(Success, JsonResponse{
		Code: BadRequest,
		Msg: msg,
	})
}


// 401未登录验证
func (resp *Response) Unauthenticated(msg string) {
	resp.Ctx.AbortWithStatusJSON(Success, JsonResponse{
		Code: Unauthenticated,
		Msg: msg,
	})
}

// 403没有权限
func (resp *Response) NoPermisson(msg string) {
	resp.Ctx.AbortWithStatusJSON(Success, JsonResponse{
		Code: NoPermisson,
		Msg: msg,
	})
}

// 404资源不存在
func (resp *Response) NotFund(msg string) {
	resp.Ctx.AbortWithStatusJSON(Success, JsonResponse{
		Code: NotFund,
		Msg: msg,
	})
}

// 500服务器出错
func (resp *Response) ServerError(msg string) {
	resp.Ctx.AbortWithStatusJSON(200, JsonResponse{
		Code: ServerError,
		Msg: msg,
	})
}
