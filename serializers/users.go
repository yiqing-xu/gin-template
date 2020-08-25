/*
@Time : 2020/6/28 22:16
@Author : xuyiqing
@File : users.go
*/

package serializers

import "gin-template/models"

type Login struct {
	Username string `form:"usernmae"; json:"username"`
	Password string `form:"password"; json:"password"`
}

func (l *Login) GetUser() *models.Account {
	return &models.Account{
		Username: l.Username,
		Password: l.Password,
	}
}

type Account struct {
	Username string `form:"username" json:"username"`
	OldPwd string `form:"oldPwd" json:"oldPwd"`
	NewPwd string `form:"newPwd"json:"newPwd"`
}
