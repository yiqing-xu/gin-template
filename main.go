/*
@Time : 2020/6/28 21:24
@Author : xuyiqing
@File : main.py
*/

package main

import (
	"flag"
	"fmt"
	"gin-template/conf"
	"gin-template/models"
	"gin-template/routers"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

var port string
var isDebugMode bool
var isErrMsg bool
var isOrmDebug bool


func init()  {
	flag.StringVar(&port, "p", "", "监听端口")
	flag.BoolVar(&isDebugMode, "debug", true, "是否开启debug")
	flag.BoolVar(&isErrMsg, "err", true, "是否返回错误信息")
	flag.BoolVar(&isOrmDebug, "orm", true, "是否开启gorm的debug信息")
	flag.Parse()

	conf.SetUp()
	models.SetUp(isOrmDebug)
}

func main() {
	if isDebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	var f *os.File
	if err := os.Mkdir("logs", os.ModePerm); err != nil {
		f, _ = os.OpenFile("logs/gin.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	} else {
		f, _ = os.Create("logs/gin.log")
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	defer models.DB.Close()
	defer f.Close()

	router := routers.InitRouter(isErrMsg)

	if len([]rune(port)) < 4 {
		port = conf.HttpServer.Port
	}
	panic(router.Run(fmt.Sprintf(":%s", port)))

	//server := http.Server{
	//	Addr:           fmt.Sprintf(":%d", 7890+i),
	//	Handler:        router,
	//	ReadTimeout:    conf.HttpServer.ReadTimeout * time.Second,
	//	WriteTimeout:   conf.HttpServer.WriteTimeout * time.Second,
	//}
	//return server.ListenAndServe()

}
