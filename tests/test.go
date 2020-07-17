/*
@Time : 2020/7/17 11:47
@Author : xuyiqing
@File : test.py
*/

package main

import (
	"flag"
	"fmt"
)

func main()  {
	var port string
	var mode bool
	var errMsg bool
	var ormDebug bool

	flag.StringVar(&port, "p", "", "web端口")
	flag.BoolVar(&mode, "debug", false, "是否开启debug")
	flag.BoolVar(&errMsg, "err", false, "是否返回错误信息")
	flag.BoolVar(&ormDebug, "orm", false, "是否开启gorm的debug信息")
	flag.Parse()
	fmt.Println(port, mode, errMsg, ormDebug)
}
