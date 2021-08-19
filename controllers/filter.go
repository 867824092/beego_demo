package controllers

import "github.com/astaxie/beego"

type FilterController struct {
	beego.Controller
}

func (c *FilterController) Get() {
	c.Ctx.ResponseWriter.Write([]byte("过滤器访问"))
}
