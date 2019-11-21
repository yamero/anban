package main

import (
	_ "anban/models"
	_ "anban/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/astaxie/beego/session/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	logs.Async(1e3) // 开启异步写日志
	logs.EnableFuncCallDepth(true) // 日志包含文件名和行号
	// 设置日志引擎为file(文件)
	logs.SetLogger("file",`{"filename":"logs/app.log","level":7,"maxlines":1000,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)

	beego.Run()
}
