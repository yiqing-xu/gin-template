/*
@Time : 2020/6/28 22:16
@Author : xuyiqing
@File : users.py
*/

package serializers

import "gin-template/models"

type Account struct {
	Username string `form:"username"json:"username"`
	Password string `form:"password"json:"-"`
}

func (a *Account) GenAccount() models.Account {
	return models.Account{
		Username: a.Username,
		Password: a.Password,
	}
}
