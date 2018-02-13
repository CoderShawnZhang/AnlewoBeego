package controllers

import (
	// "github.com/astaxie/beego"
	// "strconv"
	// "hello/models"
)

type OrdersController struct {
	BaseController
}

func (c *OrdersController) Index() {
	c.render()
}

func (c *OrdersController) Detail() {
	c.render()
}
//hello:<Ormer.QueryTable> table name: `hello/models.Orders` not exists.  需要注册