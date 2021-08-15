package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"

	//usermodel "go_web/common"
	"go_web/common"
)

type UserController struct {
	beego.Controller
}

var userArray = make([]*common.User, 0)

func init() {
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("第%d个名字", i+1)
		userArray = append(userArray, common.NewUser(i+1, name, 20+i))
	}
}

func (c *UserController) Get() {
	c.TplName = "user/index.html"
	//c.Layout = "shard/_Layout.html"
	//c.LayoutSections = make(map[string]string)
	//c.LayoutSections["Scripts"] = "user.tpl"
}

// Post 新增编辑
func (c *UserController) Post() {
	model := common.User{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &model); err != nil {
		c.Data["json"] = common.NewResponse(false, "解析参数失败："+err.Error())
	}
	var entity *common.User
	for _, user := range userArray {
		if user.Id == model.Id { //更改
			entity = user
		}
	}
	if entity != nil { //更改
		entity.Name = model.Name
		entity.Age = model.Age
	} else { //新增
		userArray = append(userArray, &model)
	}
	c.Data["json"] = common.NewResponse(true, "操作成功")
	c.ServeJSON()
}

func (c *UserController) Delete() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	index := -1
	for i, user := range userArray {
		if user.Id == id { //更改
			index = i
		}
	}
	if index != -1 {
		userArray = append(userArray[:index], userArray[index+1:]...)
	}
	c.Data["json"] = common.NewResponse(true, "删除成功")
	c.ServeJSON()

}

func (c *UserController) GetList() {
	c.Data["json"] = &userArray
	c.ServeJSON()
}
