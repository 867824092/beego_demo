package routers

import (
	"github.com/astaxie/beego"
	"go_web/controllers"
)

func init() {
	beego.BConfig.WebConfig.TemplateLeft = "<@"
	beego.BConfig.WebConfig.TemplateRight = "@>"
	beego.Router("/", &controllers.MainController{})

	beego.Router("/detail", &controllers.MainController{}, "get:Detail")

	//region 用户控制器
	//region 利用数组来实现
	//ns := beego.NewNamespace("user",
	//	beego.NSRouter("/",&controllers.UserController{}),
	//	beego.NSRouter("/getList",&controllers.UserController{},"get:GetList"),
	//	beego.NSRouter("/?:id",&controllers.UserController{}),
	//)
	//endregion
	//region 利用map来实现
	ns := beego.NewNamespace("user",
		beego.NSRouter("/", &controllers.UserMapController{}),
		beego.NSRouter("/getList", &controllers.UserMapController{}, "get:GetList"),
		beego.NSRouter("/?:id", &controllers.UserMapController{}),
	)
	//endregion
	//注册 namespace
	beego.AddNamespace(ns)
	//endregion

	//
}
