package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"hello/models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Detail() {
	id := c.Ctx.Input.Param(":id")
	tid, _ := strconv.Atoi(id)
	ok, user := models.FindUserByUserName(tid)
	if ok {
		c.Data["user"] = user
		c.Data["CurrentUserInfo"] = "99999999"
		c.Data["PageTitle"] = "个人主页"
	}
	c.TplName = "user.tpl"
}

func (c *UserController) Search() {
	id := c.Ctx.Input.Param(":id")
	tid, _ := strconv.Atoi(id)
	ok, list := models.FindUserByUserId(tid)
	if ok {
		c.Data["user"] = list
	}
	c.TplName = "index.tpl"
}

func (c *UserController) Update() {
	id := c.Ctx.Input.Param(":id")
	tid, _ := strconv.Atoi(id)
	ok, topic := models.FindUserByUserId(tid)
	if ok {
		topic.Password = "123456"
		topic.Remark ="33333"
		models.UpdateUser(&topic)
		c.Data["user"] = topic
		c.Data["title"] = "layui"
	}
	c.TplName = "404.tpl"
}

func (c *UserController) Case() {
	c.TplName = "case/case.html"
}
func (c *UserController) Insert() {
	// id := c.Ctx.Input.Param(":id")
	// uid, _  := strconv.Atoi(id)
	user := models.User{Name: "111",Password: "00000", Remark:"77777799999"}
	newId := models.InsertUser(&user)
	c.Redirect("/detail/" + strconv.FormatInt(newId,10), 302)
}