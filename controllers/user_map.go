package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"go_web/common"
	"sort"
	"strconv"
)

type UserMapController struct {
	beego.Controller
}

var userMap = make(map[int]*common.User, 0)

func init() {
	for i := 10; i >= 1; i-- {
		name := fmt.Sprintf("第%d个名字", i)
		userMap[i] = common.NewUser(i, name, 20+i)
	}
}

func (c *UserMapController) Get() {
	c.TplName = "user/index.html"
}
func (c *UserMapController) GetList() {
	userList := make([]*common.User, 0)
	for _, value := range userMap {
		userList = append(userList, value)
	}
	sort.Slice(userList, func(i, j int) bool {

		return userList[i].Id > userList[j].Id
	})
	c.Data["json"] = &userList
	c.ServeJSON()
}

// Post 新增编辑
func (c *UserMapController) Post() {
	model := common.User{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &model); err != nil {
		c.Data["json"] = common.NewResponse(false, "解析参数失败："+err.Error())
	}
	entity := userMap[model.Id]

	if entity != nil { //更改
		entity.Name = model.Name
		entity.Age = model.Age
	} else { //新增
		userMap[model.Id] = &model
	}
	c.Data["json"] = common.NewResponse(true, "操作成功")
	c.ServeJSON()
}

func (c *UserMapController) Delete() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	entity := userMap[id]
	res := common.Response{}
	if entity != nil {
		delete(userMap, id)
		res = *common.NewResponse(true, "删除成功")
	} else {
		res = *common.NewResponse(false, "所选数据已被清空，请刷新数据")
	}

	c.Data["json"] = res
	c.ServeJSON()

}
