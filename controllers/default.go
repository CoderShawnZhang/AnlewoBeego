package controllers

import (
	
)

type MainController struct {
	BaseController
}

//定义菜单结构
type Menu struct {
    Title string;
    Url  string;
};

 
func (c *MainController) Index() {
	c.Data["Title"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.render()
}