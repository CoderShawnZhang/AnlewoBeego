package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

const (
	VIEW_EXTENSION_NAME = ".html"
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	userId         int
	userName       string
	pageSize       int
}

func (c *BaseController) Prepare() {
	controllerName,actionName := c.GetControllerAndAction()
	c.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	c.actionName = strings.ToLower(actionName);
	c.Data["curRoute"] = c.controllerName + "." + c.actionName
}

func (c *BaseController) render(view ...string) {
	var viewName string
	if len(view) > 0 {
		viewName = view[0] + ".html"
	} else {
		viewName = c.controllerName + "/" + c.actionName + ".html"
	}

	c.Layout = "public/layout.html"
	c.TplName = viewName
}