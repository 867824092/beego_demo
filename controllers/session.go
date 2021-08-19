package controllers

import (
	"github.com/astaxie/beego"
	"go_web/common"
)

type SessionController struct {
	beego.Controller
}

func (c *SessionController) Get() {
	info := c.GetSession("user")
	if info == nil {
		c.Data["json"] = "当前session为null"
		c.ServeJSON()
		return
	}
	c.Data["json"] = info
	c.ServeJSON()
}

func (c *SessionController) Post() {
	user := common.NewUser(1, "张三", 20)
	c.SetSession("user", user)
	c.Data["json"] = "处理成功"
	c.ServeJSON()
}
