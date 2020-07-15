/*
@Time : 2020/6/28 21:24
@Author : xuyiqing
@File : main.py
*/

package main

import (
	"fmt"
	"gin-template/conf"
	"gin-template/models"
	"gin-template/routers"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func init()  {
	conf.SetUp()
	models.SetUp()
}

func main()  {
	gin.SetMode(gin.DebugMode)
	var f *os.File
	if err := os.Mkdir("logs", os.ModePerm); err != nil {
		f, _ = os.OpenFile("logs/gin.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	} else {
		f, _ = os.Create("logs/gin.log")
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	defer models.DB.Close()
	defer f.Close()

	router := routers.InitRouter()
	panic(router.Run(fmt.Sprintf(":%s", conf.HttpServer.Port)))

	//server := http.Server{
	//	Addr:           fmt.Sprintf(":%d", 7890+i),
	//	Handler:        router,
	//	ReadTimeout:    conf.HttpServer.ReadTimeout * time.Second,
	//	WriteTimeout:   conf.HttpServer.WriteTimeout * time.Second,
	//}
	//return server.ListenAndServe()

}
