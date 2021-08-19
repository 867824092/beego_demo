package main

import (
	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/core/logs"
	_ "go_web/routers"
)

func main() {
	//设置格式
	formatter := &logs.PatternLogFormatter{
		Pattern:    "\r %w | [%T] | %m",
		WhenFormat: "2006-01-02 15:04:05",
	}
	name := "pattern"
	logs.RegisterFormatter(name, formatter)
	logs.SetGlobalFormatter(name)
	logs.SetLogger(logs.AdapterConsole, `{"level":7,"color":true}`)
	logs.SetLogger(logs.AdapterFile,
		`{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	beego.Run()
}
