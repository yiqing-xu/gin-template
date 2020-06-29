/*
@Time : 2020/6/28 22:01
@Author : xuyiqing
@File : users.py
*/

package models

import (
	"golang.org/x/crypto/bcrypt"
)

const PasswordCryptLevel = 12

type Account struct {
	BaseModel
	Username string `gorm:"column:username;not null;unique_index;comment:'用户名'" json:"username" form:"username"`
	Password string `gorm:"column:password;comment:'密码'" form:"password" json:"-"`
	Name string `form:"name" json:"name"`
	IsActive bool `json:"-"`
}

func (a *Account) TableName() string {
	return "users_account"
}

func (a *Account) GetUserByID(id uint) *Account {
	DB.Model(&Account{}).First(a, id)
	if a.ID > 0 {
		return a
	} else {
		return nil
	}
}

// 设置密码加密
func (a *Account) SetPassword(password string) error {
	p, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCryptLevel)
	if err != nil {
		return err
	}
	a.Password = string(p)
	return nil
}

// 验证帐户密码合法性
func (a *Account) CheckPassword(password string) (*Account, bool) {
	DB.Model(&Account{}).Where("username = ?", a.Username).First(a)
	err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password))
	return a, err == nil
}

// 验证用户民重复
func (a *Account) CheckDuplicateUsername() bool {
	var count int
	if DB.Model(&Account{}).Where("username=?", a.Username).Count(&count); count > 0 {
		return false
	} else {
		return true
	}
}