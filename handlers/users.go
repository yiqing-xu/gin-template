/*
@Time : 2020/6/28 21:40
@Author : xuyiqing
@File : users.py
*/

package handlers

import (
	"encoding/json"
	"fmt"
	"gin-template/models"
	"gin-template/pkg/jwt"
	"gin-template/pkg/util"
	"github.com/gin-gonic/gin"
)

func UsersLoginHandler(ctx *gin.Context) {
	response := Response{Ctx: ctx}
	var user models.Account
	if err := ctx.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		panic(err)
	}
	loginUser, isLoginUser := user.CheckPassword(user.Password)
	if !isLoginUser {
		response.BadRequest("密码错误")
		return
	}
	token, err := jwt.GenToken(loginUser.ID, loginUser.Username)
	if err != nil {
		panic(err)
	}
	var data map[string]interface{}
	userData, _ := json.Marshal(loginUser)
	if err := json.Unmarshal(userData, &data); err != nil {
		panic(err)
	}
	data["token"] = token
	response.Response(data)
	return
}

func UsersRegisterHandler(ctx *gin.Context) {
	response := Response{Ctx: ctx}
	var user models.Account
	if err := ctx.ShouldBind(&user); err != nil {
		panic(err)
	}
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

func UsersSetInfoHandler(ctx *gin.Context) {
	response := Response{Ctx: ctx}
	jsonData := util.GetBodyData(ctx)
	if jsonData == nil {
		response.BadRequest("获取不到参数")
		return
	}
	currentUser := jwt.AssertUser(ctx)
	if currentUser != nil {
		models.DB.Model(currentUser).Updates(jsonData)
		response.Response(currentUser)
	}
}
