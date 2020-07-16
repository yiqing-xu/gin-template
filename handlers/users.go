/*
@Time : 2020/6/28 21:40
@Author : xuyiqing
@File : users.py
*/

package handlers

import (
	"encoding/json"
	"gin-template/models"
	"gin-template/pkg/jwt"
	"gin-template/pkg/util"
	"gin-template/serializers"
	"github.com/gin-gonic/gin"
)

// 登录
func UsersLoginHandler(ctx *gin.Context) {
	response := Response{Ctx: ctx}
	var loginUser serializers.Login
	if err := ctx.ShouldBindJSON(&loginUser); err != nil {
		panic(err)
	}
	user := loginUser.GetUser()
	isLoginUser := user.CheckPassword()
	if !isLoginUser {
		response.BadRequest("密码错误")
		return
	}
	token, err := jwt.GenToken(user.ID, user.Username)
	if err != nil {
		panic(err)
	}
	var data map[string]interface{}
	userData, _ := json.Marshal(user)
	if err := json.Unmarshal(userData, &data); err != nil {
		panic(err)
	}
	data["token"] = token
	response.Response(data)
	return
}

// 注册
func UsersRegisterHandler(ctx *gin.Context) {
	response := Response{Ctx: ctx}
	var registerUser serializers.Login
	if err := ctx.ShouldBind(&registerUser); err != nil {
		panic(err)
	}
	user := registerUser.GetUser()
	status := user.CheckDuplicateUsername()
	if status == false {
		response.BadRequest("用户名已存在")
		return
	}
	if err := user.SetPassword(user.Password); err != nil {
		panic(err)
	}
	user.IsActive = true
	models.DB.Create(&user)
	response.Response(nil)
}

// 修改用户信息
func UsersSetInfoHandler(ctx *gin.Context) {
	response := Response{Ctx: ctx}
	jsonData := util.GetBodyData(ctx)
	if jsonData == nil {
		response.BadRequest("获取不到参数")
		return
	}
	currentUser := jwt.AssertUser(ctx)
	if currentUser != nil {
		models.DB.Model(&currentUser).Updates(jsonData)
		response.Response(currentUser)
		return
	}
}

// 修改密码
func UsersSetPwdHandler(ctx *gin.Context) {
	response := Response{Ctx: ctx}
	currentUser := jwt.AssertUser(ctx)
	if currentUser == nil {
		response.Unauthenticated("未验证登录")
		return
	}
	var user serializers.Account
	if err := ctx.ShouldBindJSON(&user); err != nil {
		response.BadRequest(err.Error())
		return
	}
	if user.Username != currentUser.Username {
		response.BadRequest("当前登录用户用户名与输入用户名不符")
		return
	}
	if user.OldPwd == user.NewPwd {
		response.BadRequest("两次输入的密码相同")
		return
	}
	if isPwd := currentUser.IsPasswordEqual(user.OldPwd); !isPwd {
		response.BadRequest("原密码错误")
		return
	}
	if err := currentUser.SetPassword(user.NewPwd); err != nil {
		response.BadRequest(err.Error())
		return
	}
	models.DB.Save(&currentUser)
	response.Response(nil)
}
