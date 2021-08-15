package controllers

import (
	"github.com/astaxie/beego"
	"time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "867824092@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Detail() {
	//c.Ctx.ResponseWriter.Write([]byte("name:"+name))
	id := c.Ctx.Input.Query("id")
	c.Ctx.ResponseWriter.Write([]byte("接收到参数:" + id))
	c.Ctx.ResponseWriter.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))
}
