/*
@Time : 2020/6/28 21:46
@Author : xuyiqing
@File : init.py
*/

package models

import (
	"fmt"
	"gin-template/conf"
	"gin-template/pkg/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

func SetUp() {
	conUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		                  conf.DataBase.User,
		                  conf.DataBase.Password,
		                  conf.DataBase.Host,
		                  conf.DataBase.Port,
		                  conf.DataBase.DB,
		                  conf.DataBase.Charset)
	db, err := gorm.Open(conf.DataBase.Type, conUri)
	if err != nil {
		panic(err)
	}
	DB = db

	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return conf.DataBase.Prefix + defaultTableName
	}

	DB.AutoMigrate(&Account{})
	DB.AutoMigrate(&FileModel{})
	DB.AutoMigrate(&Message{})

}

type BaseModel struct {
	ID        uint64 `gorm:"primary_key'" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

// 生成全局唯一ID
func (m *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	if m.ID == 0 {
		m.ID = util.GenSonyFlakeId()
	}
	return nil
}
