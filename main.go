/*
@Time : 2020/6/28 21:24
@Author : xuyiqing
@File : main.go
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

var (
	f *os.File
	host string
	port string
	isDebugMode bool
	isErrMsg bool
	isOrmDebug bool
)

func init() {
	flag.StringVar(&host, "h", "127.0.0.1", "主机")
	flag.StringVar(&port, "p", "", "监听端口")
	flag.BoolVar(&isDebugMode, "debug", true, "是否开启debug")
	flag.BoolVar(&isErrMsg, "err", true, "是否返回错误信息")
	flag.BoolVar(&isOrmDebug, "orm", true, "是否开启gorm的debug信息")
	flag.Parse()

	conf.SetUp()
	models.SetUp(isOrmDebug)

	if isDebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	if err := os.Mkdir("logs", os.ModePerm); err != nil {
		f, _ = os.OpenFile("logs/gin.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	} else {
		f, _ = os.Create("logs/gin.log")
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	defer models.DB.Close()
	defer f.Close()
	router := routers.InitRouter(isErrMsg, isDebugMode)

	if len([]rune(port)) < 4 {
		port = conf.HttpServer.Port
	}
	panic(router.Run(fmt.Sprintf("%s:%s", host, port)))

	//server := http.Server{
	//	Addr:           "127.0.0.1:7890",
	//	Handler:        router,
	//	ReadTimeout:    conf.HttpServer.ReadTimeout * time.Second,
	//	WriteTimeout:   conf.HttpServer.WriteTimeout * time.Second,
	//}
	//panic(server.ListenAndServe())

}
