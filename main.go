/*
@Time : 2020/6/28 21:24
@Author : xuyiqing
@File : main.py
*/

package main

import (
	"gin-template/conf"
	"gin-template/models"
	"gin-template/routers"
)

func init()  {
	conf.SetUp()
	models.SetUp()
}

func main()  {
	router := routers.InitRouter()
	err := router.Run()
	if err != nil {

	}
}
