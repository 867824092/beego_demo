package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
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

	//#region session
	beego.Router("/session", &controllers.SessionController{}, "get:Get;post:Post")
	//#endregion

	//region filter
	beego.InsertFilter("/filter", beego.BeforeExec, func(context *context.Context) {
		fmt.Println("ActionExec Filter Execute")
	})
	beego.Router("/filter", &controllers.FilterController{})
	//endregion

	//region orm
	orm := beego.NewNamespace("orm",
		beego.NSRouter("/", &controllers.UserOrmController{}, "get:Index;post:Post;put:Put"),
		beego.NSRouter("/List", &controllers.UserOrmController{}, "get:List"),
		beego.NSRouter("/:id", &controllers.UserOrmController{}, "get:Get;delete:Delete"),
		beego.NSRouter("/migration", &controllers.UserOrmController{}, "get:Migrate"),
	)
	beego.AddNamespace(orm)
	//endregion
}
